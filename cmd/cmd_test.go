package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	cve "github.com/scagogogo/cve-skills"
)

// TestReadInputsFromArgs 覆盖 len(args) > 0 分支：直接返回 args，不读 stdin。
func TestReadInputsFromArgs(t *testing.T) {
	args := []string{"CVE-2022-1", "CVE-2022-2"}
	got := readInputs(args)
	if len(got) != 2 || got[0] != "CVE-2022-1" || got[1] != "CVE-2022-2" {
		t.Fatalf("readInputs(args) = %v, want %v", got, args)
	}
}

// TestReadInputsFromStdin 覆盖 stdin 管道分支：逐行读取，跳过空行。
func TestReadInputsFromStdin(t *testing.T) {
	originalStdin := os.Stdin
	defer func() { os.Stdin = originalStdin }()

	input := "CVE-2022-1\n\nCVE-2022-2\n"
	tmp, err := os.CreateTemp("", "stdin-*")
	if err != nil {
		t.Fatalf("CreateTemp: %v", err)
	}
	defer os.Remove(tmp.Name())
	if _, err := tmp.WriteString(input); err != nil {
		t.Fatalf("WriteString: %v", err)
	}
	tmp.Close()

	f, err := os.Open(tmp.Name())
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	defer f.Close()
	os.Stdin = f

	got := readInputs(nil)
	want := []string{"CVE-2022-1", "CVE-2022-2"}
	if len(got) != len(want) {
		t.Fatalf("readInputs(nil) from stdin = %v, want %v (empty lines skipped)", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("readInputs(nil)[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}

// TestReadInputsCharDevice 覆盖字符设备分支（helpers.go:17）：
// /dev/null 是字符设备，(stat.Mode() & os.ModeCharDevice) != 0 为真 → return nil。
func TestReadInputsCharDevice(t *testing.T) {
	originalStdin := os.Stdin
	defer func() { os.Stdin = originalStdin }()

	f, err := os.Open("/dev/null")
	if err != nil {
		t.Skipf("/dev/null not available: %v", err)
	}
	defer f.Close()
	os.Stdin = f

	got := readInputs(nil)
	if got != nil {
		t.Fatalf("readInputs(nil) with /dev/null (char device) = %v, want nil", got)
	}
}

// TestReadInputsEmptyPipe 覆盖 stdin 管道 EOF 分支：scanner 零迭代，lines 为 nil。
func TestReadInputsEmptyPipe(t *testing.T) {
	originalStdin := os.Stdin
	defer func() { os.Stdin = originalStdin }()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Pipe: %v", err)
	}
	os.Stdin = r
	w.Close() // 立即关闭写端，读端读到 EOF，scanner 无数据

	got := readInputs(nil)
	if got != nil && len(got) != 0 {
		t.Fatalf("readInputs(nil) with empty pipe = %v, want nil or empty", got)
	}
}

var (
	coverBinaryOnce sync.Once
	coverBinaryPath string
	coverBinaryErr  error

	// coverBinaryDir 是 buildCoveredBinary 创建的持久目录，TestMain 结束时清理。
	coverBinaryDir string

	// collectedCoverDirs 收集所有 runCve 子进程的 GOCOVERDIR，供 TestMain 合并为单一 coverage profile。
	// go test -coverprofile 只收集测试进程内覆盖，不收集 go build -cover 子进程二进制的覆盖；
	// 子进程把覆盖数据写 GOCOVERDIR，TestMain 用 go tool covdata merge + textfmt 导出标准 profile。
	collectedCoverDirs []string
	coverMu            sync.Mutex
)

var (
	moduleRootOnce sync.Once
	moduleRoot     string
	moduleRootErr  error
)

// findModuleRoot 通过 `go list -m -f {{.Dir}}` 取模块根目录，避免硬编码绝对路径。
func findModuleRoot() string {
	moduleRootOnce.Do(func() {
		out, err := exec.Command("go", "list", "-m", "-f", "{{.Dir}}").Output()
		if err != nil {
			moduleRootErr = err
			return
		}
		moduleRoot = strings.TrimSpace(string(out))
	})
	return moduleRoot
}

// buildCoveredBinary 编译带覆盖率插桩的 cve CLI 二进制（main 包），缓存复用。
// go build -cover（Go 1.20+）编译出带插桩的二进制，运行时若设置 GOCOVERDIR 则写覆盖率数据。
// 目标是 ./cmd/cve（main 包入口），而非测试二进制，这样子进程运行的是真正的 cobra CLI。
// 注意：二进制路径不能用 t.TempDir()，否则 Once.Do 首次调用的 t 结束后会清理掉目录，
// 后续测试找不到二进制。改用 os.MkdirTemp 创建独立于任何 t 的持久目录，由 TestMain 统一清理。
func buildCoveredBinary(t *testing.T) string {
	coverBinaryOnce.Do(func() {
		root := findModuleRoot()
		if moduleRootErr != nil {
			coverBinaryErr = fmt.Errorf("find module root: %v", moduleRootErr)
			return
		}
		dir, err := os.MkdirTemp("", "cve-cover-bin-*")
		if err != nil {
			coverBinaryErr = fmt.Errorf("MkdirTemp: %v", err)
			return
		}
		coverBinaryDir = dir
		path := filepath.Join(dir, "cve")
		cmd := exec.Command("go", "build", "-cover", "-o", path, "./cmd/cve")
		cmd.Dir = root
		out, err := cmd.CombinedOutput()
		if err != nil {
			coverBinaryErr = fmt.Errorf("go build -cover failed: %v\n%s", err, out)
			return
		}
		coverBinaryPath = path
	})
	if coverBinaryErr != nil {
		t.Fatalf("build covered binary: %v", coverBinaryErr)
	}
	return coverBinaryPath
}

// runCve 执行带覆盖率的二进制，返回 stdout、stderr、exitCode。
// 每个 t 有独立 coverDir，避免并发写冲突；coverDir 注册到 collectedCoverDirs 供 TestMain 合并。
// 注意：coverDir 用 os.MkdirTemp（非 t.TempDir）——t.TempDir 在 t 结束时清理，
// 而 TestMain 在所有 t 结束后才合并覆盖数据，那时 t.TempDir 已被删，covdata 会读不到。
// 这些持久目录由 TestMain 统一清理。
func runCve(t *testing.T, args ...string) (stdout, stderr string, exitCode int) {
	t.Helper()
	bin := buildCoveredBinary(t)
	coverDir, err := os.MkdirTemp("", "cve-cover-run-*")
	if err != nil {
		t.Fatalf("MkdirTemp coverDir: %v", err)
	}
	coverMu.Lock()
	collectedCoverDirs = append(collectedCoverDirs, coverDir)
	coverMu.Unlock()
	cmd := exec.Command(bin, args...)
	cmd.Env = append(os.Environ(), "GOCOVERDIR="+coverDir)
	var out, errBuf bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	err = cmd.Run()
	stdout = out.String()
	stderr = errBuf.String()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		} else {
			t.Fatalf("run cve %v: %v (stderr: %s)", args, err, stderr)
		}
	}
	return stdout, stderr, exitCode
}

// TestCLIVersion 覆盖 version 命令的 Run 闭包 + Execute 无 error 路径。
func TestCLIVersion(t *testing.T) {
	stdout, _, code := runCve(t, "version")
	if code != 0 {
		t.Fatalf("cve version exit code = %d, want 0", code)
	}
	// cve.Version 默认 "dev"（源码构建无 ldflags），版本号即此值
	want := cve.Version + "\n"
	if stdout != want {
		t.Fatalf("cve version stdout = %q, want %q", stdout, want)
	}
}

// TestCLIRootHelp 覆盖 rootCmd --help 路径：打印帮助内容。
func TestCLIRootHelp(t *testing.T) {
	stdout, _, code := runCve(t, "--help")
	if code != 0 {
		t.Fatalf("cve --help exit code = %d, want 0", code)
	}
	if !strings.Contains(stdout, "cve") || !strings.Contains(stdout, "Available Commands") {
		t.Fatalf("cve --help output missing help content: %q", stdout)
	}
}

// TestCLIUnknownCommand 覆盖 Execute 的 error 路径：未知子命令 → cobra 返回 error → Execute 打印 stderr 并 os.Exit(1)。
func TestCLIUnknownCommand(t *testing.T) {
	_, _, code := runCve(t, "no-such-command")
	if code != 1 {
		t.Fatalf("cve no-such-command exit code = %d, want 1 (Execute error path)", code)
	}
}

// TestMain 合并子进程覆盖数据并清理持久二进制目录。
//
// go test -coverprofile 只收集测试进程内的覆盖（readInputs 等），不收集 go build -cover
// 子进程二进制的覆盖（Execute、version 的 Run 闭包等）。子进程运行时设置 GOCOVERDIR，
// 把覆盖数据写入该目录；runCve 把每个 coverDir 注册到 collectedCoverDirs，这里统一合并：
//
//	go tool covdata merge -i=<dir1>,<dir2>,... -o=<mergeddir>   // 合并多个 GOCOVERDIR
//	go tool covdata textfmt -i=<mergeddir> -o=<profile.out>     // 转标准 mode: set profile
//
// 最终 profile 路径取自 GOCOVER_SUBPROCESS_OUT 环境变量，与 go test -coverprofile 互补。
// 注意 covdata 的 -i 用逗号分隔目录（非 os.PathListSeparator 冒号）。
func TestMain(m *testing.M) {
	code := m.Run()

	// 合并子进程覆盖数据为单一 profile。coverDir 用 os.MkdirTemp 持久化（非 t.TempDir，
	// 否则 t 结束即清理，此处合并时目录已不存在）。covdata 的 -i 用逗号分隔目录。
	if finalOut := os.Getenv("GOCOVER_SUBPROCESS_OUT"); finalOut != "" && len(collectedCoverDirs) > 0 {
		merged, err := os.MkdirTemp("", "cve-cover-merged-*")
		if err == nil {
			defer os.RemoveAll(merged)
			dirs := strings.Join(collectedCoverDirs, ",")
			mergeCmd := exec.Command("go", "tool", "covdata", "merge", "-i="+dirs, "-o="+merged)
			if mergeCmd.Run() == nil {
				textCmd := exec.Command("go", "tool", "covdata", "textfmt", "-i="+merged, "-o="+finalOut)
				textCmd.Run()
			}
		}
	}

	// 清理 runCve 的持久 coverDir 与编译出的覆盖率二进制目录
	for _, d := range collectedCoverDirs {
		os.RemoveAll(d)
	}
	if coverBinaryDir != "" {
		os.RemoveAll(coverBinaryDir)
	}

	os.Exit(code)
}
