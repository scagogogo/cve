package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"testing"
	"time"

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
// 子进程二进制的覆盖（Execute、version 的 Run 闭包等）。子进程运行时设置 GOCOVERDIR，
// 把覆盖数据写入该目录；runCve 把每个 coverDir 注册到 collectedCoverDirs，这里统一合并：
//
//	go tool covdata merge -i=<dir1>,<dir2>,... -o=<mergeddir>   // 合并多个 GOCOVERDIR
//	go tool covdata textfmt -i=<mergeddir> -o=<profile.out>     // 转标准 mode: set profile
//
// 两条互斥路径：
//   - GOCOVER_SUBPROCESS_OUT=<file>：旧路径，TestMain 内部 merge+textfmt 输出单一 profile，删各 coverDir
//   - GOCOVER_SUBPROCESS_DIRS=<file>：新路径，把 collectedCoverDirs 路径列表写入该文件，不合并不删，
//     供外部 Makefile 与进程内 covdir 一并用 covdata merge -pcombine 统一合并成全仓库单一 profile
func TestMain(m *testing.M) {
	code := m.Run()

	// 持久 coverDir 列表优先交给外部合并（纯 Go covdata merge 路径）
	if dirsFile := os.Getenv("GOCOVER_SUBPROCESS_DIRS"); dirsFile != "" && len(collectedCoverDirs) > 0 {
		content := []byte(strings.Join(collectedCoverDirs, "\n") + "\n")
		if err := os.WriteFile(dirsFile, content, 0o644); err != nil {
			fmt.Fprintf(os.Stderr, "write GOCOVER_SUBPROCESS_DIRS: %v\n", err)
		}
		// 保留 coverDir 供外部 merge，不清理
		// 仅清理二进制目录
		if coverBinaryDir != "" {
			os.RemoveAll(coverBinaryDir)
		}
		os.Exit(code)
	}

	// 旧路径：内部 merge+textfmt 输出单一 profile
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

// TestCLIFormat 覆盖 format 命令：有 args → 逐个格式化打印（Run 闭包 for 循环）。
func TestCLIFormat(t *testing.T) {
	stdout, _, code := runCve(t, "format", "cve-2022-12345", " cve-2021-44228 ")
	if code != 0 {
		t.Fatalf("cve format exit code = %d, want 0", code)
	}
	want := "CVE-2022-12345\nCVE-2021-44228\n"
	if stdout != want {
		t.Fatalf("cve format stdout = %q, want %q", stdout, want)
	}
}

// TestCLIFormatNoInput 覆盖 format 命令空输入分支：len(inputs)==0 → os.Exit(1)。
func TestCLIFormatNoInput(t *testing.T) {
	_, _, code := runCve(t, "format")
	if code != 1 {
		t.Fatalf("cve format (no input) exit code = %d, want 1", code)
	}
}

// TestCLIExtract 覆盖 extract 命令：从文本提取多个 CVE 并逐行打印。
func TestCLIExtract(t *testing.T) {
	stdout, _, code := runCve(t, "extract", "affected by CVE-2021-44228 and cve-2022-12345")
	if code != 0 {
		t.Fatalf("cve extract exit code = %d, want 0", code)
	}
	want := "CVE-2021-44228\nCVE-2022-12345\n"
	if stdout != want {
		t.Fatalf("cve extract stdout = %q, want %q", stdout, want)
	}
}

// TestCLIExtractSubcommands 覆盖 first/last/year/seq/split 子命令的 Run 闭包。
func TestCLIExtractSubcommands(t *testing.T) {
	cases := []struct {
		args []string
		want string
	}{
		{[]string{"extract", "first", "CVE-2021-44228 and CVE-2022-12345"}, "CVE-2021-44228\n"},
		{[]string{"extract", "last", "CVE-2021-44228 and CVE-2022-12345"}, "CVE-2022-12345\n"},
		{[]string{"extract", "year", "CVE-2022-12345"}, "2022\n"},
		{[]string{"extract", "seq", "CVE-2022-12345"}, "12345\n"},
		{[]string{"extract", "split", "CVE-2022-12345"}, "2022\t12345\n"},
	}
	for _, tc := range cases {
		t.Run(strings.Join(tc.args, " "), func(t *testing.T) {
			stdout, _, code := runCve(t, tc.args...)
			if code != 0 {
				t.Fatalf("%v exit code = %d, want 0", tc.args, code)
			}
			if stdout != tc.want {
				t.Fatalf("%v stdout = %q, want %q", tc.args, stdout, tc.want)
			}
		})
	}
}

// TestCLIExtractNoInput 覆盖 extract 空输入 → os.Exit(1)。
func TestCLIExtractNoInput(t *testing.T) {
	_, _, code := runCve(t, "extract")
	if code != 1 {
		t.Fatalf("cve extract (no input) exit code = %d, want 1", code)
	}
}

// TestCLICompare 覆盖 compare 命令（ExactArgs(2)）：CompareCves 返回值打印。
func TestCLICompare(t *testing.T) {
	stdout, _, code := runCve(t, "compare", "CVE-2021-44228", "CVE-2022-12345")
	if code != 0 {
		t.Fatalf("cve compare exit code = %d, want 0", code)
	}
	if stdout != "-1\n" {
		t.Fatalf("cve compare stdout = %q, want %q", stdout, "-1\n")
	}
}

// TestCLICompareByYear 覆盖 compare by-year 子命令。
func TestCLICompareByYear(t *testing.T) {
	stdout, _, code := runCve(t, "compare", "by-year", "CVE-2021-44228", "CVE-2022-12345")
	if code != 0 {
		t.Fatalf("cve compare by-year exit code = %d, want 0", code)
	}
	if stdout != "-1\n" {
		t.Fatalf("cve compare by-year stdout = %q, want %q", stdout, "-1\n")
	}
}

// TestCLISort 覆盖 compare sort 子命令：readInputs + SortCves 逐行打印。
func TestCLISort(t *testing.T) {
	stdout, _, code := runCve(t, "compare", "sort", "CVE-2022-2222", "CVE-2020-1111", "CVE-2022-1111")
	if code != 0 {
		t.Fatalf("cve compare sort exit code = %d, want 0", code)
	}
	want := "CVE-2020-1111\nCVE-2022-1111\nCVE-2022-2222\n"
	if stdout != want {
		t.Fatalf("cve compare sort stdout = %q, want %q", stdout, want)
	}
}

// TestCLIFilterByYear 覆盖 by-year --year 有效 → FilterCvesByYear 打印。
func TestCLIFilterByYear(t *testing.T) {
	stdout, _, code := runCve(t, "filter", "by-year", "--year", "2022", "CVE-2021-1111", "CVE-2022-2222")
	if code != 0 {
		t.Fatalf("cve filter by-year exit code = %d, want 0", code)
	}
	if stdout != "CVE-2022-2222\n" {
		t.Fatalf("cve filter by-year stdout = %q, want %q", stdout, "CVE-2022-2222\n")
	}
}

// TestCLIFilterByYearMissingFlag 覆盖 --year==0 → 报错 os.Exit(1)。
func TestCLIFilterByYearMissingFlag(t *testing.T) {
	_, _, code := runCve(t, "filter", "by-year", "CVE-2022-1")
	if code != 1 {
		t.Fatalf("cve filter by-year (no --year) exit code = %d, want 1", code)
	}
}

// TestCLIFilterByYearRange 覆盖 by-year-range --start/--end 有效。
func TestCLIFilterByYearRange(t *testing.T) {
	stdout, _, code := runCve(t, "filter", "by-year-range", "--start", "2021", "--end", "2022",
		"CVE-2020-1111", "CVE-2021-2222", "CVE-2022-3333")
	if code != 0 {
		t.Fatalf("cve filter by-year-range exit code = %d, want 0", code)
	}
	want := "CVE-2021-2222\nCVE-2022-3333\n"
	if stdout != want {
		t.Fatalf("cve filter by-year-range stdout = %q, want %q", stdout, want)
	}
}

// TestCLIFilterByYearRangeMissingFlag 覆盖 --start==0 或 --end==0 → os.Exit(1)。
func TestCLIFilterByYearRangeMissingFlag(t *testing.T) {
	_, _, code := runCve(t, "filter", "by-year-range", "--end", "2022", "CVE-2022-1")
	if code != 1 {
		t.Fatalf("cve filter by-year-range (no --start) exit code = %d, want 1", code)
	}
}

// TestCLIFilterRecent 覆盖 recent --years 有效（输出依赖当前年，断言子集）。
func TestCLIFilterRecent(t *testing.T) {
	currentYear := time.Now().Year()
	// 用当前年份的 CVE 确保落在"最近 N 年"范围内，使 GetRecentCves 返回非空，
	// 从而覆盖 `for _, c := range filtered { fmt.Println(c) }` 循环体（filter.go:93-95）。
	recentCve := fmt.Sprintf("CVE-%d-2222", currentYear)
	stdout, _, code := runCve(t, "filter", "recent", "--years", "2",
		"CVE-2000-1111", recentCve)
	if code != 0 {
		t.Fatalf("cve filter recent exit code = %d, want 0", code)
	}
	// 当前年份 CVE 应被保留；2000 一定被排除。
	if !strings.Contains(stdout, recentCve) {
		t.Fatalf("cve filter recent should keep %s, got %q", recentCve, stdout)
	}
	if strings.Contains(stdout, "CVE-2000-1111") {
		t.Fatalf("cve filter recent should exclude CVE-2000-1111, got %q", stdout)
	}
}

// TestCLIFilterRecentMissingFlag 覆盖 --years==0 → os.Exit(1)。
func TestCLIFilterRecentMissingFlag(t *testing.T) {
	_, _, code := runCve(t, "filter", "recent", "CVE-2022-1")
	if code != 1 {
		t.Fatalf("cve filter recent (no --years) exit code = %d, want 1", code)
	}
}

// TestCLIGroupByYear 覆盖 group-by-year：按年分组打印（map 遍历顺序不定，断言内容存在）。
func TestCLIGroupByYear(t *testing.T) {
	stdout, _, code := runCve(t, "filter", "group-by-year", "CVE-2021-1111", "CVE-2022-2222")
	if code != 0 {
		t.Fatalf("cve filter group-by-year exit code = %d, want 0", code)
	}
	if !strings.Contains(stdout, "2021:") || !strings.Contains(stdout, "2022:") {
		t.Fatalf("cve filter group-by-year stdout missing year groups: %q", stdout)
	}
}

// TestCLIDedup 覆盖 dedup：去重后逐行打印。
func TestCLIDedup(t *testing.T) {
	stdout, _, code := runCve(t, "filter", "dedup", "CVE-2022-1111", "cve-2022-1111", "CVE-2022-2222")
	if code != 0 {
		t.Fatalf("cve filter dedup exit code = %d, want 0", code)
	}
	want := "CVE-2022-1111\nCVE-2022-2222\n"
	if stdout != want {
		t.Fatalf("cve filter dedup stdout = %q, want %q", stdout, want)
	}
}

// TestCLIValidate 覆盖 validate 命令：Format + ValidateCve 打印（有效与无效各一）。
func TestCLIValidate(t *testing.T) {
	stdout, _, code := runCve(t, "validate", "CVE-2022-12345", "CVE-1998-1")
	if code != 0 {
		t.Fatalf("cve validate exit code = %d, want 0", code)
	}
	if !strings.Contains(stdout, "CVE-2022-12345\ttrue") {
		t.Fatalf("cve validate stdout missing valid entry: %q", stdout)
	}
	if !strings.Contains(stdout, "CVE-1998-1\tfalse") {
		t.Fatalf("cve validate stdout missing invalid entry: %q", stdout)
	}
}

// TestCLIValidateNoInput 覆盖 validate 空输入 → os.Exit(1)。
func TestCLIValidateNoInput(t *testing.T) {
	_, _, code := runCve(t, "validate")
	if code != 1 {
		t.Fatalf("cve validate (no input) exit code = %d, want 1", code)
	}
}

// TestCLIIsCve 覆盖 is-cve：打印 input\ttrue/false。
func TestCLIIsCve(t *testing.T) {
	stdout, _, code := runCve(t, "validate", "is-cve", "CVE-2022-12345", "not-a-cve")
	if code != 0 {
		t.Fatalf("cve validate is-cve exit code = %d, want 0", code)
	}
	if !strings.Contains(stdout, "CVE-2022-12345\ttrue") {
		t.Fatalf("cve is-cve stdout missing true entry: %q", stdout)
	}
	if !strings.Contains(stdout, "not-a-cve\tfalse") {
		t.Fatalf("cve is-cve stdout missing false entry: %q", stdout)
	}
}

// TestCLIContainsCve 覆盖 contains-cve：打印 true/false。
func TestCLIContainsCve(t *testing.T) {
	stdout, _, code := runCve(t, "validate", "contains-cve", "has CVE-2022-12345", "no cve here")
	if code != 0 {
		t.Fatalf("cve validate contains-cve exit code = %d, want 0", code)
	}
	if !strings.Contains(stdout, "true") || !strings.Contains(stdout, "false") {
		t.Fatalf("cve contains-cve stdout missing true/false: %q", stdout)
	}
}

// TestCLIYearOkNoCutoff 覆盖 year-ok cutoff==0 分支（走 IsCveYearOk）：true 与 false 两路径。
func TestCLIYearOkNoCutoff(t *testing.T) {
	// true 路径：1999..当前年
	stdout, _, code := runCve(t, "validate", "year-ok", "CVE-2022-12345")
	if code != 0 {
		t.Fatalf("cve validate year-ok exit code = %d, want 0", code)
	}
	if !strings.Contains(stdout, "CVE-2022-12345\ttrue") {
		t.Fatalf("cve year-ok stdout missing true: %q", stdout)
	}

	// false 路径：年份 < 1999
	stdout, _, code = runCve(t, "validate", "year-ok", "CVE-1998-12345")
	if code != 0 {
		t.Fatalf("cve validate year-ok (old year) exit code = %d, want 0", code)
	}
	if !strings.Contains(stdout, "CVE-1998-12345\tfalse") {
		t.Fatalf("cve year-ok stdout missing false for old year: %q", stdout)
	}
}

// TestCLIYearOkWithCutoff 覆盖 year-ok cutoff>0 分支（走 IsCveYearOkWithCutoff）：true 与 false 两路径。
func TestCLIYearOkWithCutoff(t *testing.T) {
	currentYear := time.Now().Year()

	// true 路径：2030 在当前年+5 容忍范围内（当前年+5 >= 2030 时为 true）
	stdout, _, code := runCve(t, "validate", "year-ok", "--cutoff", "5", "CVE-2030-12345")
	if code != 0 {
		t.Fatalf("cve validate year-ok --cutoff exit code = %d, want 0", code)
	}
	wantSuffix := "false"
	if 2030 <= currentYear+5 {
		wantSuffix = "true"
	}
	if !strings.Contains(stdout, "CVE-2030-12345\t"+wantSuffix) {
		t.Fatalf("cve year-ok --cutoff 5 (2030) stdout = %q, want suffix %q", stdout, wantSuffix)
	}

	// false 路径：极远未来年份，超出当前年+5 容忍范围，必为 false
	farFuture := fmt.Sprintf("CVE-%d-12345", currentYear+100)
	stdout, _, code = runCve(t, "validate", "year-ok", "--cutoff", "5", farFuture)
	if code != 0 {
		t.Fatalf("cve validate year-ok --cutoff (far future) exit code = %d, want 0", code)
	}
	if !strings.Contains(stdout, farFuture+"\tfalse") {
		t.Fatalf("cve year-ok --cutoff 5 (far future) stdout = %q, want false", stdout)
	}
}

// TestCLIGenerateCve 覆盖 generate cve --year/--seq 有效。
func TestCLIGenerateCve(t *testing.T) {
	stdout, _, code := runCve(t, "generate", "cve", "--year", "2022", "--seq", "12345")
	if code != 0 {
		t.Fatalf("cve generate cve exit code = %d, want 0", code)
	}
	if stdout != "CVE-2022-12345\n" {
		t.Fatalf("cve generate cve stdout = %q, want %q", stdout, "CVE-2022-12345\n")
	}
}

// TestCLIGenerateCveMissingFlag 覆盖 year==0 或 seq==0 → 打印 error（return，不 exit 1，但 cobra 不视为 error，exit 0）。
func TestCLIGenerateCveMissingFlag(t *testing.T) {
	stdout, _, code := runCve(t, "generate", "cve", "--year", "2022")
	if code != 0 {
		t.Fatalf("cve generate cve (missing --seq) exit code = %d, want 0 (Run uses return not os.Exit)", code)
	}
	if !strings.Contains(stdout, "error: --year and --seq are required") {
		t.Fatalf("cve generate cve (missing --seq) stdout = %q, want error message", stdout)
	}
}

// TestCLIGenerateFake 覆盖 generate fake：输出随机但匹配 CVE 格式。
func TestCLIGenerateFake(t *testing.T) {
	stdout, _, code := runCve(t, "generate", "fake")
	if code != 0 {
		t.Fatalf("cve generate fake exit code = %d, want 0", code)
	}
	// 断言输出形如 CVE-YYYY-NNNNN
	if !regexp.MustCompile(`(?i)^CVE-\d{4}-\d+\s*$`).MatchString(stdout) {
		t.Fatalf("cve generate fake stdout = %q, want a CVE-shaped string", stdout)
	}
}

// TestCLIFilterPattern 覆盖 filter-pattern RunE：正常匹配 + 参数不足 error。
func TestCLIFilterPattern(t *testing.T) {
	stdout, _, code := runCve(t, "filter-pattern", "CVE-2022-*", "CVE-2022-1111,CVE-2023-2222")
	if code != 0 {
		t.Fatalf("cve filter-pattern exit code = %d, want 0", code)
	}
	if stdout != "CVE-2022-1111\n" {
		t.Fatalf("cve filter-pattern stdout = %q, want %q", stdout, "CVE-2022-1111\n")
	}
}

// TestCLIFilterPatternTooFewArgs 覆盖 len(inputs)<2 → RunE 返回 error → exit 1。
func TestCLIFilterPatternTooFewArgs(t *testing.T) {
	_, _, code := runCve(t, "filter-pattern", "CVE-2022-*")
	if code != 1 {
		t.Fatalf("cve filter-pattern (too few args) exit code = %d, want 1", code)
	}
}

// TestCLIFormatSeq 覆盖 format-seq RunE：正常 + 宽度非数字 error + 参数不足 error。
func TestCLIFormatSeq(t *testing.T) {
	stdout, _, code := runCve(t, "format-seq", "6", "CVE-2022-123")
	if code != 0 {
		t.Fatalf("cve format-seq exit code = %d, want 0", code)
	}
	if stdout != "CVE-2022-000123\n" {
		t.Fatalf("cve format-seq stdout = %q, want %q", stdout, "CVE-2022-000123\n")
	}
}

// TestCLIFormatSeqInvalidWidth 覆盖 strconv.Atoi 失败 → RunE 返回 error → exit 1。
func TestCLIFormatSeqInvalidWidth(t *testing.T) {
	_, _, code := runCve(t, "format-seq", "abc", "CVE-2022-123")
	if code != 1 {
		t.Fatalf("cve format-seq (invalid width) exit code = %d, want 1", code)
	}
}

// TestCLIFormatSeqTooFewArgs 覆盖 len(inputs)<2 → exit 1。
func TestCLIFormatSeqTooFewArgs(t *testing.T) {
	_, _, code := runCve(t, "format-seq", "6")
	if code != 1 {
		t.Fatalf("cve format-seq (too few args) exit code = %d, want 1", code)
	}
}

// TestCLIParseRange 覆盖 parse-range RunE：正常展开 + nil 结果 error。
func TestCLIParseRange(t *testing.T) {
	stdout, _, code := runCve(t, "parse-range", "CVE-2022-1..3")
	if code != 0 {
		t.Fatalf("cve parse-range exit code = %d, want 0", code)
	}
	want := "CVE-2022-1\nCVE-2022-2\nCVE-2022-3\n"
	if stdout != want {
		t.Fatalf("cve parse-range stdout = %q, want %q", stdout, want)
	}
}

// TestCLIParseRangeInvalid 覆盖 result==nil → RunE 返回 error → exit 1。
func TestCLIParseRangeInvalid(t *testing.T) {
	_, _, code := runCve(t, "parse-range", "not-a-range")
	if code != 1 {
		t.Fatalf("cve parse-range (invalid) exit code = %d, want 1", code)
	}
}

// TestCLIIsConsecutive 覆盖 is-consecutive RunE：连续 true 与非连续 false 两分支。
func TestCLIIsConsecutive(t *testing.T) {
	cases := []struct {
		args []string
		want string
	}{
		{[]string{"is-consecutive", "CVE-2022-1", "CVE-2022-2"}, "CVE-2022-1 and CVE-2022-2 are consecutive\n"},
		{[]string{"is-consecutive", "CVE-2022-1", "CVE-2022-3"}, "CVE-2022-1 and CVE-2022-3 are NOT consecutive\n"},
	}
	for _, tc := range cases {
		t.Run(strings.Join(tc.args, " "), func(t *testing.T) {
			stdout, _, code := runCve(t, tc.args...)
			if code != 0 {
				t.Fatalf("%v exit code = %d, want 0", tc.args, code)
			}
			if stdout != tc.want {
				t.Fatalf("%v stdout = %q, want %q", tc.args, stdout, tc.want)
			}
		})
	}
}

// TestCLISetOps 覆盖 intersect/union/diff RunE：正常 + 参数不足 error。
func TestCLISetOps(t *testing.T) {
	cases := []struct {
		args []string
		want string
	}{
		{[]string{"intersect", "CVE-2022-1,CVE-2022-2", "CVE-2022-2,CVE-2022-3"}, "CVE-2022-2\n"},
		{[]string{"union", "CVE-2022-1", "CVE-2022-2"}, "CVE-2022-1\nCVE-2022-2\n"},
		{[]string{"diff", "CVE-2022-1,CVE-2022-2", "CVE-2022-2"}, "CVE-2022-1\n"},
	}
	for _, tc := range cases {
		t.Run(tc.args[0], func(t *testing.T) {
			stdout, _, code := runCve(t, tc.args...)
			if code != 0 {
				t.Fatalf("%v exit code = %d, want 0", tc.args, code)
			}
			if stdout != tc.want {
				t.Fatalf("%v stdout = %q, want %q", tc.args, stdout, tc.want)
			}
		})
	}
}

// TestCLISetOpsTooFewArgs 覆盖 intersect/union/diff len(inputs)<2 → exit 1。
func TestCLISetOpsTooFewArgs(t *testing.T) {
	for _, op := range []string{"intersect", "union", "diff"} {
		t.Run(op, func(t *testing.T) {
			_, _, code := runCve(t, op, "CVE-2022-1")
			if code != 1 {
				t.Fatalf("cve %s (too few args) exit code = %d, want 1", op, code)
			}
		})
	}
}

// TestCLICountByYear 覆盖 count-by-year RunE：正常打印（map 顺序不定，断言内容）。
func TestCLICountByYear(t *testing.T) {
	stdout, _, code := runCve(t, "count-by-year", "CVE-2022-1,CVE-2022-2,CVE-2021-3")
	if code != 0 {
		t.Fatalf("cve count-by-year exit code = %d, want 0", code)
	}
	if !strings.Contains(stdout, "2022: 2") || !strings.Contains(stdout, "2021: 1") {
		t.Fatalf("cve count-by-year stdout = %q, want counts for 2022:2 and 2021:1", stdout)
	}
}

// TestCLIYearRange 覆盖 year-range RunE。
func TestCLIYearRange(t *testing.T) {
	stdout, _, code := runCve(t, "year-range", "CVE-2020-1,CVE-2022-2,CVE-2021-3")
	if code != 0 {
		t.Fatalf("cve year-range exit code = %d, want 0", code)
	}
	want := "Year range: 2020 - 2022 (span: 2 years)\n"
	if stdout != want {
		t.Fatalf("cve year-range stdout = %q, want %q", stdout, want)
	}
}

// TestCLISeqRange 覆盖 seq-range RunE：正常 + year 非数字 error + 参数不足 error。
func TestCLISeqRange(t *testing.T) {
	stdout, _, code := runCve(t, "seq-range", "2022", "CVE-2022-1,CVE-2022-3,CVE-2022-2")
	if code != 0 {
		t.Fatalf("cve seq-range exit code = %d, want 0", code)
	}
	want := "Year 2022 sequence range: 1 - 3\n"
	if stdout != want {
		t.Fatalf("cve seq-range stdout = %q, want %q", stdout, want)
	}
}

// TestCLISeqRangeInvalidYear 覆盖 strconv.Atoi(year) 失败 → exit 1。
func TestCLISeqRangeInvalidYear(t *testing.T) {
	_, _, code := runCve(t, "seq-range", "abc", "CVE-2022-1")
	if code != 1 {
		t.Fatalf("cve seq-range (invalid year) exit code = %d, want 1", code)
	}
}

// TestCLISeqRangeTooFewArgs 覆盖 len(inputs)<2 → exit 1。
func TestCLISeqRangeTooFewArgs(t *testing.T) {
	_, _, code := runCve(t, "seq-range", "2022")
	if code != 1 {
		t.Fatalf("cve seq-range (too few args) exit code = %d, want 1", code)
	}
}

// TestCLIValidateBatch 覆盖 validate-batch RunE：r.Valid 真假两分支（✓ 与 ✗）。
func TestCLIValidateBatch(t *testing.T) {
	stdout, _, code := runCve(t, "validate-batch", "CVE-2022-12345,CVE-1998-1")
	if code != 0 {
		t.Fatalf("cve validate-batch exit code = %d, want 0", code)
	}
	if !strings.Contains(stdout, "✓ CVE-2022-12345") {
		t.Fatalf("cve validate-batch stdout missing valid mark: %q", stdout)
	}
	if !strings.Contains(stdout, "✗ CVE-1998-1") || !strings.Contains(stdout, "before 1999") {
		t.Fatalf("cve validate-batch stdout missing invalid mark/reason: %q", stdout)
	}
}

// TestCLIFilterValid 覆盖 filter-valid RunE。
func TestCLIFilterValid(t *testing.T) {
	stdout, _, code := runCve(t, "filter-valid", "CVE-2022-12345,not-a-cve,CVE-1998-1")
	if code != 0 {
		t.Fatalf("cve filter-valid exit code = %d, want 0", code)
	}
	if stdout != "CVE-2022-12345\n" {
		t.Fatalf("cve filter-valid stdout = %q, want %q", stdout, "CVE-2022-12345\n")
	}
}

// TestCLISortNoInput 覆盖 compare sort 空输入 → os.Exit(1)。
func TestCLISortNoInput(t *testing.T) {
	_, _, code := runCve(t, "compare", "sort")
	if code != 1 {
		t.Fatalf("cve compare sort (no input) exit code = %d, want 1", code)
	}
}

// TestCLIExtractSubcommandsNoInput 覆盖 extract first/last/year/seq/split 空输入 → os.Exit(1)。
func TestCLIExtractSubcommandsNoInput(t *testing.T) {
	for _, sub := range []string{"first", "last", "year", "seq", "split"} {
		t.Run(sub, func(t *testing.T) {
			_, _, code := runCve(t, "extract", sub)
			if code != 1 {
				t.Fatalf("cve extract %s (no input) exit code = %d, want 1", sub, code)
			}
		})
	}
}

// TestCLIValidateSubcommandsNoInput 覆盖 validate is-cve/contains-cve/year-ok 空输入 → os.Exit(1)。
func TestCLIValidateSubcommandsNoInput(t *testing.T) {
	for _, sub := range []string{"is-cve", "contains-cve", "year-ok"} {
		t.Run(sub, func(t *testing.T) {
			_, _, code := runCve(t, "validate", sub)
			if code != 1 {
				t.Fatalf("cve validate %s (no input) exit code = %d, want 1", sub, code)
			}
		})
	}
}

// TestCLIFilterSubcommandsNoInput 覆盖 filter by-year/by-year-range/recent/group-by-year/dedup 的"flag 有效 + 输入空"路径 → os.Exit(1)。
func TestCLIFilterSubcommandsNoInput(t *testing.T) {
	cases := []struct {
		args []string
	}{
		{[]string{"filter", "by-year", "--year", "2022"}},
		{[]string{"filter", "by-year-range", "--start", "2021", "--end", "2022"}},
		{[]string{"filter", "recent", "--years", "2"}},
		{[]string{"filter", "group-by-year"}},
		{[]string{"filter", "dedup"}},
	}
	for _, tc := range cases {
		t.Run(strings.Join(tc.args, " "), func(t *testing.T) {
			_, _, code := runCve(t, tc.args...)
			if code != 1 {
				t.Fatalf("%v (no input) exit code = %d, want 1", tc.args, code)
			}
		})
	}
}

// TestCLIFilterParentHelp 覆盖 filter 父命令无子命令 → cmd.Help()，exit 0。
func TestCLIFilterParentHelp(t *testing.T) {
	stdout, _, code := runCve(t, "filter")
	if code != 0 {
		t.Fatalf("cve filter (no subcommand) exit code = %d, want 0 (Help)", code)
	}
	if !strings.Contains(stdout, "Filter") && !strings.Contains(stdout, "filter") {
		t.Fatalf("cve filter (no subcommand) stdout missing help: %q", stdout)
	}
}

// TestCLIGenerateParentHelp 覆盖 generate 父命令无子命令 → cmd.Help()，exit 0。
func TestCLIGenerateParentHelp(t *testing.T) {
	stdout, _, code := runCve(t, "generate")
	if code != 0 {
		t.Fatalf("cve generate (no subcommand) exit code = %d, want 0 (Help)", code)
	}
	if !strings.Contains(stdout, "Generate") && !strings.Contains(stdout, "generate") {
		t.Fatalf("cve generate (no subcommand) stdout missing help: %q", stdout)
	}
}

// TestCLINoArgEmptyInput 覆盖各命令"零位置参数 + 空输入"路径，触发 len(inputs)==0/<2 的 error/os.Exit(1) 分支。
// 现有的 TooFewArgs 测试传了 1 个参数（len==1），未触达 len==0；此测试不传任何位置参数，
// readInputs 收到空 args 且 stdin 为 /dev/null（字符设备）→ 返回 nil → 命中 len(inputs)==0 分支。
func TestCLINoArgEmptyInput(t *testing.T) {
	cases := []struct {
		name string
		args []string
	}{
		{"count-by-year no args", []string{"count-by-year"}},
		{"year-range no args", []string{"year-range"}},
		{"validate-batch no args", []string{"validate-batch"}},
		{"filter-valid no args", []string{"filter-valid"}},
		{"filter recent flag no input", []string{"filter", "recent", "--years", "2"}},
		{"parse-range no args", []string{"parse-range"}},
		{"is-consecutive no args", []string{"is-consecutive"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, _, code := runCve(t, tc.args...)
			if code != 1 {
				t.Fatalf("%v exit code = %d, want 1 (empty input → error/exit)", tc.args, code)
			}
		})
	}
}
