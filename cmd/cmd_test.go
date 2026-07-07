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
)

// buildCoveredBinary 编译带覆盖率插桩的 cve CLI 二进制（main 包），缓存复用。
// go build -cover（Go 1.20+）编译出带插桩的二进制，运行时若设置 GOCOVERDIR 则写覆盖率数据。
// 目标是 ./cmd/cve（main 包入口），而非测试二进制，这样子进程运行的是真正的 cobra CLI。
// 注意：二进制路径不能用 t.TempDir()，否则 Once.Do 首次调用的 t 结束后会清理掉目录，
// 后续测试找不到二进制。改用 os.MkdirTemp 创建独立于任何 t 的持久目录。
func buildCoveredBinary(t *testing.T) string {
	coverBinaryOnce.Do(func() {
		dir, err := os.MkdirTemp("", "cve-cover-bin-*")
		if err != nil {
			coverBinaryErr = fmt.Errorf("MkdirTemp: %w", err)
			return
		}
		path := filepath.Join(dir, "cve")
		cmd := exec.Command("go", "build", "-cover", "-o", path, "./cmd/cve")
		cmd.Dir = "/home/cc11001100/github/scagogogo/cve-skills"
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

// runCve 执行带覆盖率的二进制，返回 stdout、exitCode。
// 每个 t 有独立 coverDir，避免并发写冲突。
func runCve(t *testing.T, args ...string) (stdout string, exitCode int) {
	t.Helper()
	bin := buildCoveredBinary(t)
	coverDir := t.TempDir()
	cmd := exec.Command(bin, args...)
	cmd.Env = append(os.Environ(), "GOCOVERDIR="+coverDir)
	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	stdout = out.String()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		} else {
			t.Fatalf("run cve %v: %v (stderr: %s)", args, err, stderr.String())
		}
	}
	return stdout, exitCode
}

// TestCLIVersion 覆盖 version 命令的 Run 闭包 + Execute 无 error 路径。
func TestCLIVersion(t *testing.T) {
	stdout, code := runCve(t, "version")
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
	stdout, code := runCve(t, "--help")
	if code != 0 {
		t.Fatalf("cve --help exit code = %d, want 0", code)
	}
	if !strings.Contains(stdout, "cve") || !strings.Contains(stdout, "Available Commands") {
		t.Fatalf("cve --help output missing help content: %q", stdout)
	}
}

// TestCLIUnknownCommand 覆盖 Execute 的 error 路径：未知子命令 → cobra 返回 error → Execute 打印 stderr 并 os.Exit(1)。
func TestCLIUnknownCommand(t *testing.T) {
	_, code := runCve(t, "no-such-command")
	if code != 1 {
		t.Fatalf("cve no-such-command exit code = %d, want 1 (Execute error path)", code)
	}
}
