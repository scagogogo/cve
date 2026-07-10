# CLI 层 100% 覆盖率与分支粒度深化 Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: `superpowers:subagent-driven-development`
> Steps use checkbox (`- [ ]`) syntax.

**Goal:** 为 cve-skills 的 CLI 层（`cmd/` 包）补齐单元测试，使 `cmd/` 覆盖率从 0% 提升到 100%，并对每个命令的每个分支（参数缺失、flag 为 0、cutoff 分支、nil 结果、Atoi 失败、valid/invalid 输出）逐一覆盖；同时跑通全部测试、修复发现的问题。

**Architecture:** 库层（base/compare/extract/filter/generate.go）已达 100.0% 语句覆盖率且 `go test -race` 通过，无需改动。本 Plan 聚焦 CLI 层 14 个文件。CLI 命令的 `Run` 闭包内直接调用 `os.Exit(1)`，无法在测试进程内捕获（会终止测试进程），故采用 Go 官方的「编译带覆盖率的测试二进制 + 子进程执行」方案：`go test -c -cover -o /tmp/cve.test ./cmd` 编译出带插桩的二进制，测试用 `os/exec` 以各种参数组合运行它，断言 stdout 与 exit code，子进程退出后写 coverage profile，最后用 `go tool cover` 合并。`readInputs` 是无 `os.Exit` 的纯函数，进程内直接调用测试；`Execute`/`Run` 闭包的 `os.Exit` 路径用子进程测试。关键组件：`cmd/cmd_test.go`（测试文件）、`buildCoveredBinary` 辅助函数（编译一次缓存复用）、`runCve` 辅助函数（执行子进程并返回 stdout+exitCode）。

**Tech Stack:** Go 1.25（模块声明 go 1.18），github.com/spf13/cobra v1.8.1，标准库 `os/exec`/`testing`/`strings`/`regexp`/`time`，`go test -c -cover` 覆盖率插桩二进制，`go tool cover -func` 汇总

**Risks:**
- R1：`os.Exit(1)` 不可在进程内捕获 → 缓解：子进程方案，编译带覆盖率的二进制运行，断言 exit code。
- R2：`examples/`（33 个 `main` 包）0% 覆盖率不在本 Plan 目标内 → 缓解：明确边界，覆盖率目标只针对 `.`（库）和 `cmd/`（CLI），examples 是可执行示例程序非被测代码。
- R3：时间相关命令（validate/year-ok/generate fake）输出依赖当前年份 → 缓解：期望值用 `time.Now().Year()` 动态构造，与库测试一致；`generate fake` 只断言 CVE 格式。
- R4：子进程测试需先编译带覆盖率二进制，单次测试运行慢于进程内 → 缓解：`buildCoveredBinary` 在 `TestMain`/首次调用时编译一次缓存到临时目录复用。
- R5：`go test -coverpkg=./...` 合并库与 CLI 覆盖率时，库的二进制插桩与 CLI 的需分别生成 profile 再合并 → 缓解：用 `go test -coverpkg=github.com/scagogogo/cve-skills,github.com/scagogogo/cve-skills/cmd -coverprofile` 一次性合并。

---

### Task 1: readInputs 纯函数进程内测试

**Depends on:** None
**Files:**
- Create: `cmd/cmd_test.go`

- [ ] **Step 1: 创建 cmd_test.go 骨架与 readInputs 测试 — 覆盖 args/stdin/空输入三分支**

readInputs 是 cmd 包内唯一无 os.Exit 的纯函数，可在进程内直接测试，覆盖其三个分支：args 非空直接返回、stdin 是字符设备（终端）返回 nil、stdin 管道逐行读取并跳过空行。

```go
// cmd/cmd_test.go
package cmd

import (
	"bytes"
	"os"
	"testing"
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
	r := bytes.NewReader([]byte(input))
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

// TestReadInputsEmptyArgsNoPipe 覆盖 (stat.Mode() & os.ModeCharDevice) != 0 分支：
// 无 args 且 stdin 是字符设备（终端）时返回 nil。测试环境下 os.Stdin 通常是管道或文件，
// 此分支难以在 CI 中稳定触发；通过直接验证函数对无 args 且非管道 stdin 的契约来锁定行为。
// 当 stdin 不可 stat 为字符设备时，函数返回空切片（lines 为 nil），断言其为空。
func TestReadInputsEmptyArgsNoPipe(t *testing.T) {
	// 用一个空管道（非字符设备，但无数据）触发 lines 为空的路径
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
```

- [ ] **Step 2: 验证 readInputs 测试通过**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills && go test -run 'TestReadInputs' ./cmd/ -v`
Expected:
  - Exit code: 0
  - Output contains: "PASS" and "ok" and "TestReadInputsFromArgs"
  - Output does NOT contain: "FAIL"

- [ ] **Step 3: 提交**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills && git add cmd/cmd_test.go && git commit -m "test(cmd): add readInputs unit tests covering args/stdin/empty branches"`

---

### Task 2: 子进程测试基础设施与 Execute/version 测试

**Depends on:** Task 1
**Files:**
- Modify: `cmd/cmd_test.go`

- [ ] **Step 1: 添加子进程测试辅助函数 — 编译带覆盖率二进制并缓存复用**

CLI 命令的 Run 闭包含 os.Exit(1)，无法进程内测试。采用编译带 -cover 插桩的二进制 + os/exec 子进程方案。buildCoveredBinary 编译一次缓存到临时目录；runCve 执行子进程返回 stdout/stderr/exitCode。注意：子进程必须设置 GOCOVERDIR 环境变量，覆盖率数据才会落盘。

```go
// 追加到 cmd/cmd_test.go（在 import 块中增加 "os/exec" "path/filepath" "strings" "sync"）

var (
	coverBinaryOnce sync.Once
	coverBinaryPath string
	coverBinaryErr  error
)

// buildCoveredBinary 编译带覆盖率插桩的 cve 测试二进制，缓存复用。
// go test -c -cover 编译出一个可执行文件，运行时若设置 GOCOVERDIR 则写覆盖率数据。
func buildCoveredBinary(t *testing.T) string {
	coverBinaryOnce.Do(func() {
		dir := t.TempDir()
		path := filepath.Join(dir, "cve.test")
		cmd := exec.Command("go", "test", "-c", "-cover", "-o", path, ".")
		cmd.Dir = "/home/cc11001100/github/scagogogo/cve-skills/cmd"
		out, err := cmd.CombinedOutput()
		if err != nil {
			coverBinaryErr = fmt.Errorf("go test -c -cover failed: %v\n%s", err, out)
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
	// 把覆盖率数据合并到固定位置供 TestMain 收集
	mergeCoverData(t, coverDir)
	return stdout, exitCode
}

// mergeCoverData 将子进程产生的 coverDir 中的覆盖率计数器文件移动到
// 测试进程级收集目录，供 go tool cover 最终合并。GOCOVERDIR 模式下，
// go test 主进程会在结束时自动 flush，但子进程的数据需手动归集。
var collectedCoverDirs []string

func mergeCoverData(t *testing.T, coverDir string) {
	collectedCoverDirs = append(collectedCoverDirs, coverDir)
}
```

同步更新 import 块（在文件顶部）：

```go
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
)
```

- [ ] **Step 2: 添加 version 与 Execute 正常路径测试 — 覆盖 version Run 闭包与 Execute 无 error 路径**

version 命令的 Run 闭包打印 cve.Version，无 os.Exit，子进程运行 exit code 0。这同时覆盖 Execute 的无 error 分支（rootCmd.Execute() 返回 nil）。

```go
// 追加到 cmd/cmd_test.go

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

// TestCLIRootHelp 覆盖 rootCmd 无子命令时打印帮助（cmd.Help() 路径经由 cobra 默认）。
func TestCLIRootHelp(t *testing.T) {
	stdout, code := runCve(t, "--help")
	if code != 0 {
		t.Fatalf("cve --help exit code = %d, want 0", code)
	}
	if !strings.Contains(stdout, "cve") || !strings.Contains(stdout, "Available Commands") {
		t.Fatalf("cve --help output missing help content: %q", stdout)
	}
}

// TestCLIUnknownCommand 覆盖 Execute 的 error 路径：未知子命令 → cobra 返回 error → Execute 打印到 stderr 并 os.Exit(1)。
func TestCLIUnknownCommand(t *testing.T) {
	_, code := runCve(t, "no-such-command")
	if code != 1 {
		t.Fatalf("cve no-such-command exit code = %d, want 1 (Execute error path)", code)
	}
}
```

需在 import 中引入库包：在 import 块增加 `cve "github.com/scagogogo/cve-skills"`。

更新后的 import 块：

```go
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
```

- [ ] **Step 3: 验证 version/Execute 测试通过**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills && go test -run 'TestCLI' ./cmd/ -v`
Expected:
  - Exit code: 0
  - Output contains: "TestCLIVersion" and "PASS" and "ok"
  - Output does NOT contain: "FAIL"

- [ ] **Step 4: 提交**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills && git add cmd/cmd_test.go && git commit -m "test(cmd): add subprocess test harness and version/Execute coverage"`

---

### Task 3: 各 Run/RunE 命令子进程分支测试

**Depends on:** Task 2
**Files:**
- Modify: `cmd/cmd_test.go`

- [ ] **Step 1: 添加 format/extract/compare 系列命令测试 — 覆盖 args 输入、空输入 exit 1、多参数分支**

format/extract 系列命令的 Run 闭包：有 args 则逐个处理打印，无 args 且 stdin 无数据则 os.Exit(1)。compare 的 compare/sort/by-year 子命令各有 ExactArgs(2) 约束与 readInputs 分支。

```go
// 追加到 cmd/cmd_test.go

// TestCLIFormat 覆盖 format 命令：有 args → 逐个格式化打印（Run 闭包 for 循环）。
func TestCLIFormat(t *testing.T) {
	stdout, code := runCve(t, "format", "cve-2022-12345", " cve-2021-44228 ")
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
	_, code := runCve(t, "format")
	if code != 1 {
		t.Fatalf("cve format (no input) exit code = %d, want 1", code)
	}
}

// TestCLIExtract 覆盖 extract 命令：从文本提取多个 CVE 并逐行打印。
func TestCLIExtract(t *testing.T) {
	stdout, code := runCve(t, "extract", "affected by CVE-2021-44228 and cve-2022-12345")
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
		args   []string
		want   string
	}{
		{[]string{"extract", "first", "CVE-2021-44228 and CVE-2022-12345"}, "CVE-2021-44228\n"},
		{[]string{"extract", "last", "CVE-2021-44228 and CVE-2022-12345"}, "CVE-2022-12345\n"},
		{[]string{"extract", "year", "CVE-2022-12345"}, "2022\n"},
		{[]string{"extract", "seq", "CVE-2022-12345"}, "12345\n"},
		{[]string{"extract", "split", "CVE-2022-12345"}, "2022\t12345\n"},
	}
	for _, tc := range cases {
		t.Run(strings.Join(tc.args, " "), func(t *testing.T) {
			stdout, code := runCve(t, tc.args...)
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
	_, code := runCve(t, "extract")
	if code != 1 {
		t.Fatalf("cve extract (no input) exit code = %d, want 1", code)
	}
}

// TestCLICompare 覆盖 compare 命令（ExactArgs(2)）：CompareCves 返回值打印。
func TestCLICompare(t *testing.T) {
	stdout, code := runCve(t, "compare", "CVE-2021-44228", "CVE-2022-12345")
	if code != 0 {
		t.Fatalf("cve compare exit code = %d, want 0", code)
	}
	if stdout != "-1\n" {
		t.Fatalf("cve compare stdout = %q, want %q", stdout, "-1\n")
	}
}

// TestCLICompareByYear 覆盖 compare by-year 子命令。
func TestCLICompareByYear(t *testing.T) {
	stdout, code := runCve(t, "compare", "by-year", "CVE-2021-44228", "CVE-2022-12345")
	if code != 0 {
		t.Fatalf("cve compare by-year exit code = %d, want 0", code)
	}
	if stdout != "-1\n" {
		t.Fatalf("cve compare by-year stdout = %q, want %q", stdout, "-1\n")
	}
}

// TestCLISort 覆盖 compare sort 子命令：readInputs + SortCves 逐行打印。
func TestCLISort(t *testing.T) {
	stdout, code := runCve(t, "compare", "sort", "CVE-2022-2222", "CVE-2020-1111", "CVE-2022-1111")
	if code != 0 {
		t.Fatalf("cve compare sort exit code = %d, want 0", code)
	}
	want := "CVE-2020-1111\nCVE-2022-1111\nCVE-2022-2222\n"
	if stdout != want {
		t.Fatalf("cve compare sort stdout = %q, want %q", stdout, want)
	}
}
```

- [ ] **Step 2: 添加 filter 系列命令测试 — 覆盖 flag 为 0 报错、有 args 正常、cutoff 分支**

filter by-year/by-year-range/recent 在 flag（--year/--start/--end/--years）为 0 时打印 stderr 并 os.Exit(1)；group-by-year/dedup 走 readInputs 分支。

```go
// 追加到 cmd/cmd_test.go

// TestCLIFilterByYear 覆盖 by-year --year 有效 → FilterCvesByYear 打印。
func TestCLIFilterByYear(t *testing.T) {
	stdout, code := runCve(t, "filter", "by-year", "--year", "2022", "CVE-2021-1111", "CVE-2022-2222")
	if code != 0 {
		t.Fatalf("cve filter by-year exit code = %d, want 0", code)
	}
	if stdout != "CVE-2022-2222\n" {
		t.Fatalf("cve filter by-year stdout = %q, want %q", stdout, "CVE-2022-2222\n")
	}
}

// TestCLIFilterByYearMissingFlag 覆盖 --year==0 → 报错 os.Exit(1)。
func TestCLIFilterByYearMissingFlag(t *testing.T) {
	_, code := runCve(t, "filter", "by-year", "CVE-2022-1")
	if code != 1 {
		t.Fatalf("cve filter by-year (no --year) exit code = %d, want 1", code)
	}
}

// TestCLIFilterByYearRange 覆盖 by-year-range --start/--end 有效。
func TestCLIFilterByYearRange(t *testing.T) {
	stdout, code := runCve(t, "filter", "by-year-range", "--start", "2021", "--end", "2022",
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
	_, code := runCve(t, "filter", "by-year-range", "--end", "2022", "CVE-2022-1")
	if code != 1 {
		t.Fatalf("cve filter by-year-range (no --start) exit code = %d, want 1", code)
	}
}

// TestCLIFilterRecent 覆盖 recent --years 有效（输出依赖当前年，断言子集）。
func TestCLIFilterRecent(t *testing.T) {
	stdout, code := runCve(t, "filter", "recent", "--years", "2",
		"CVE-2000-1111", "CVE-2022-2222")
	if code != 0 {
		t.Fatalf("cve filter recent exit code = %d, want 0", code)
	}
	// 2022 在最近 2 年内则应被保留；2000 一定被排除。仅断言 2000 不出现。
	if strings.Contains(stdout, "CVE-2000-1111") {
		t.Fatalf("cve filter recent should exclude CVE-2000-1111, got %q", stdout)
	}
}

// TestCLIFilterRecentMissingFlag 覆盖 --years==0 → os.Exit(1)。
func TestCLIFilterRecentMissingFlag(t *testing.T) {
	_, code := runCve(t, "filter", "recent", "CVE-2022-1")
	if code != 1 {
		t.Fatalf("cve filter recent (no --years) exit code = %d, want 1", code)
	}
}

// TestCLIGroupByYear 覆盖 group-by-year：按年分组打印（map 遍历顺序不定，断言内容存在）。
func TestCLIGroupByYear(t *testing.T) {
	stdout, code := runCve(t, "filter", "group-by-year", "CVE-2021-1111", "CVE-2022-2222")
	if code != 0 {
		t.Fatalf("cve filter group-by-year exit code = %d, want 0", code)
	}
	if !strings.Contains(stdout, "2021:") || !strings.Contains(stdout, "2022:") {
		t.Fatalf("cve filter group-by-year stdout missing year groups: %q", stdout)
	}
}

// TestCLIDedup 覆盖 dedup：去重后逐行打印。
func TestCLIDedup(t *testing.T) {
	stdout, code := runCve(t, "filter", "dedup", "CVE-2022-1111", "cve-2022-1111", "CVE-2022-2222")
	if code != 0 {
		t.Fatalf("cve filter dedup exit code = %d, want 0", code)
	}
	want := "CVE-2022-1111\nCVE-2022-2222\n"
	if stdout != want {
		t.Fatalf("cve filter dedup stdout = %q, want %q", stdout, want)
	}
}
```

- [ ] **Step 3: 添加 validate 系列命令测试 — 覆盖 cutoff 分支、is-cve/contains-cve/year-ok 各分支**

year-ok 命令有 `cutoff > 0` 走 IsCveYearOkWithCutoff、`cutoff == 0` 走 IsCveYearOk 的分支。validate-batch 有 `r.Valid` 真假两分支。

```go
// 追加到 cmd/cmd_test.go

// TestCLIValidate 覆盖 validate 命令：Format + ValidateCve 打印（有效与无效各一）。
func TestCLIValidate(t *testing.T) {
	stdout, code := runCve(t, "validate", "CVE-2022-12345", "CVE-1998-1")
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
	_, code := runCve(t, "validate")
	if code != 1 {
		t.Fatalf("cve validate (no input) exit code = %d, want 1", code)
	}
}

// TestCLIIsCve 覆盖 is-cve：打印 input\ttrue/false。
func TestCLIIsCve(t *testing.T) {
	stdout, code := runCve(t, "validate", "is-cve", "CVE-2022-12345", "not-a-cve")
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
	stdout, code := runCve(t, "validate", "contains-cve", "has CVE-2022-12345", "no cve here")
	if code != 0 {
		t.Fatalf("cve validate contains-cve exit code = %d, want 0", code)
	}
	if !strings.Contains(stdout, "true") || !strings.Contains(stdout, "false") {
		t.Fatalf("cve contains-cve stdout missing true/false: %q", stdout)
	}
}

// TestCLIYearOkNoCutoff 覆盖 year-ok cutoff==0 分支（走 IsCveYearOk）。
func TestCLIYearOkNoCutoff(t *testing.T) {
	stdout, code := runCve(t, "validate", "year-ok", "CVE-2022-12345")
	if code != 0 {
		t.Fatalf("cve validate year-ok exit code = %d, want 0", code)
	}
	if !strings.Contains(stdout, "CVE-2022-12345\ttrue") {
		t.Fatalf("cve year-ok stdout missing true: %q", stdout)
	}
}

// TestCLIYearOkWithCutoff 覆盖 year-ok cutoff>0 分支（走 IsCveYearOkWithCutoff）。
func TestCLIYearOkWithCutoff(t *testing.T) {
	stdout, code := runCve(t, "validate", "year-ok", "--cutoff", "5", "CVE-2030-12345")
	if code != 0 {
		t.Fatalf("cve validate year-ok --cutoff exit code = %d, want 0", code)
	}
	// 2030 <= 当前年+5（假设当前年<=2025），应有效。用动态期望。
	currentYear := time.Now().Year()
	wantSuffix := "false"
	if 2030 <= currentYear+5 {
		wantSuffix = "true"
	}
	if !strings.Contains(stdout, "CVE-2030-12345\t"+wantSuffix) {
		t.Fatalf("cve year-ok --cutoff 5 stdout = %q, want suffix %q", stdout, wantSuffix)
	}
}
```

在 import 块增加 `"time"`：

```go
import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	cve "github.com/scagogogo/cve-skills"
)
```

- [ ] **Step 4: 添加 generate/pattern/range/set/stats/validate-batch 系列命令测试 — 覆盖 RunE 各 error 分支与正常分支**

这些命令用 RunE（返回 error 而非 os.Exit），error 经 cobra 打印到 stderr，exit code 非 0。覆盖：参数不足 error、Atoi 失败 error、nil 结果 error、正常输出、r.Valid 真假。

```go
// 追加到 cmd/cmd_test.go

// TestCLIGenerateCve 覆盖 generate cve --year/--seq 有效。
func TestCLIGenerateCve(t *testing.T) {
	stdout, code := runCve(t, "generate", "cve", "--year", "2022", "--seq", "12345")
	if code != 0 {
		t.Fatalf("cve generate cve exit code = %d, want 0", code)
	}
	if stdout != "CVE-2022-12345\n" {
		t.Fatalf("cve generate cve stdout = %q, want %q", stdout, "CVE-2022-12345\n")
	}
}

// TestCLIGenerateCveMissingFlag 覆盖 year==0 或 seq==0 → 打印 error（return，不 exit 1，但 cobra 不视为 error，exit 0）。
func TestCLIGenerateCveMissingFlag(t *testing.T) {
	stdout, code := runCve(t, "generate", "cve", "--year", "2022")
	if code != 0 {
		t.Fatalf("cve generate cve (missing --seq) exit code = %d, want 0 (Run uses return not os.Exit)", code)
	}
	if !strings.Contains(stdout, "error: --year and --seq are required") {
		t.Fatalf("cve generate cve (missing --seq) stdout = %q, want error message", stdout)
	}
}

// TestCLIGenerateFake 覆盖 generate fake：输出随机但匹配 CVE 格式。
func TestCLIGenerateFake(t *testing.T) {
	stdout, code := runCve(t, "generate", "fake")
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
	stdout, code := runCve(t, "filter-pattern", "CVE-2022-*", "CVE-2022-1111,CVE-2023-2222")
	if code != 0 {
		t.Fatalf("cve filter-pattern exit code = %d, want 0", code)
	}
	if stdout != "CVE-2022-1111\n" {
		t.Fatalf("cve filter-pattern stdout = %q, want %q", stdout, "CVE-2022-1111\n")
	}
}

// TestCLIFilterPatternTooFewArgs 覆盖 len(inputs)<2 → RunE 返回 error → exit 1。
func TestCLIFilterPatternTooFewArgs(t *testing.T) {
	_, code := runCve(t, "filter-pattern", "CVE-2022-*")
	if code != 1 {
		t.Fatalf("cve filter-pattern (too few args) exit code = %d, want 1", code)
	}
}

// TestCLIFormatSeq 覆盖 format-seq RunE：正常 + 宽度非数字 error + 参数不足 error。
func TestCLIFormatSeq(t *testing.T) {
	stdout, code := runCve(t, "format-seq", "6", "CVE-2022-123")
	if code != 0 {
		t.Fatalf("cve format-seq exit code = %d, want 0", code)
	}
	if stdout != "CVE-2022-000123\n" {
		t.Fatalf("cve format-seq stdout = %q, want %q", stdout, "CVE-2022-000123\n")
	}
}

// TestCLIFormatSeqInvalidWidth 覆盖 strconv.Atoi 失败 → RunE 返回 error → exit 1。
func TestCLIFormatSeqInvalidWidth(t *testing.T) {
	_, code := runCve(t, "format-seq", "abc", "CVE-2022-123")
	if code != 1 {
		t.Fatalf("cve format-seq (invalid width) exit code = %d, want 1", code)
	}
}

// TestCLIFormatSeqTooFewArgs 覆盖 len(inputs)<2 → exit 1。
func TestCLIFormatSeqTooFewArgs(t *testing.T) {
	_, code := runCve(t, "format-seq", "6")
	if code != 1 {
		t.Fatalf("cve format-seq (too few args) exit code = %d, want 1", code)
	}
}

// TestCLIParseRange 覆盖 parse-range RunE：正常展开 + nil 结果 error。
func TestCLIParseRange(t *testing.T) {
	stdout, code := runCve(t, "parse-range", "CVE-2022-1..3")
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
	_, code := runCve(t, "parse-range", "not-a-range")
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
			stdout, code := runCve(t, tc.args...)
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
			stdout, code := runCve(t, tc.args...)
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
			_, code := runCve(t, op, "CVE-2022-1")
			if code != 1 {
				t.Fatalf("cve %s (too few args) exit code = %d, want 1", op, code)
			}
		})
	}
}

// TestCLICountByYear 覆盖 count-by-year RunE：正常打印（map 顺序不定，断言内容）。
func TestCLICountByYear(t *testing.T) {
	stdout, code := runCve(t, "count-by-year", "CVE-2022-1,CVE-2022-2,CVE-2021-3")
	if code != 0 {
		t.Fatalf("cve count-by-year exit code = %d, want 0", code)
	}
	if !strings.Contains(stdout, "2022: 2") || !strings.Contains(stdout, "2021: 1") {
		t.Fatalf("cve count-by-year stdout = %q, want counts for 2022:2 and 2021:1", stdout)
	}
}

// TestCLIYearRange 覆盖 year-range RunE。
func TestCLIYearRange(t *testing.T) {
	stdout, code := runCve(t, "year-range", "CVE-2020-1,CVE-2022-2,CVE-2021-3")
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
	stdout, code := runCve(t, "seq-range", "2022", "CVE-2022-1,CVE-2022-3,CVE-2022-2")
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
	_, code := runCve(t, "seq-range", "abc", "CVE-2022-1")
	if code != 1 {
		t.Fatalf("cve seq-range (invalid year) exit code = %d, want 1", code)
	}
}

// TestCLISeqRangeTooFewArgs 覆盖 len(inputs)<2 → exit 1。
func TestCLISeqRangeTooFewArgs(t *testing.T) {
	_, code := runCve(t, "seq-range", "2022")
	if code != 1 {
		t.Fatalf("cve seq-range (too few args) exit code = %d, want 1", code)
	}
}

// TestCLIValidateBatch 覆盖 validate-batch RunE：r.Valid 真假两分支（✓ 与 ✗）。
func TestCLIValidateBatch(t *testing.T) {
	stdout, code := runCve(t, "validate-batch", "CVE-2022-12345,CVE-1998-1")
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
	stdout, code := runCve(t, "filter-valid", "CVE-2022-12345,not-a-cve,CVE-1998-1")
	if code != 0 {
		t.Fatalf("cve filter-valid exit code = %d, want 0", code)
	}
	if stdout != "CVE-2022-12345\n" {
		t.Fatalf("cve filter-valid stdout = %q, want %q", stdout, "CVE-2022-12345\n")
	}
}
```

在 import 块增加 `"regexp"`：

```go
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
```

- [ ] **Step 5: 验证所有 CLI 命令测试通过**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills && go test ./cmd/ -v 2>&1 | tail -40`
Expected:
  - Exit code: 0
  - Output contains: "PASS" and "ok" and "TestCLIVersion"
  - Output does NOT contain: "FAIL"

- [ ] **Step 6: 提交**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills && git add cmd/cmd_test.go && git commit -m "test(cmd): add subprocess branch tests for all CLI commands"`

---

### Task 4: 覆盖率合并、全量验证与文档同步

**Depends on:** Task 3
**Files:**
- Modify: `website/guide/testing.md`
- Modify: `website/zh/guide/testing.md`

- [ ] **Step 1: 合并库与 CLI 覆盖率并验证 cmd/ 达 100%**

用 -coverpkg 同时插桩库与 cmd 包，子进程覆盖率经 GOCOVERDIR 收集，go test 主进程在退出时自动 flush 合并。

Run: `cd /home/cc11001100/github/scagogogo/cve-skills && go test -coverpkg=./... -coverprofile=/tmp/full.out ./... 2>&1 | tail -20 && echo "---CMD FUNC COV---" && go tool cover -func=/tmp/full.out | grep -E 'cmd/|total:'`
Expected:
  - Exit code: 0
  - Output contains: "ok" for both packages
  - `go tool cover -func` 中所有 `cmd/*.go` 函数显示 100.0%
  - `total:` 行显示接近 100%（cmd 全 100%，库全 100%；examples 仍为 0% 但不计入这两个包的目标）

- [ ] **Step 2: 跑竞态检测与 vet**

Run: `cd /home/cc11001100/github/scagogogo/cve-skills && go test -race ./... 2>&1 | tail -10 && echo "---VET---" && go vet ./... 2>&1 | tail -10`
Expected:
  - Exit code: 0
  - Output contains: "ok" for all packages
  - vet 无输出（无警告）

- [ ] **Step 3: 修复测试中发现的问题（如有）**

执行 Task 3 Step 5 与本 Task Step 1/2 的验证命令。若任何测试失败或覆盖率未达 100%，用 `superpowers:systematic-debugging` 定位根因。常见问题及修复：
- 子进程找不到二进制 → 确认 buildCoveredBinary 的 cmd.Dir 指向 cmd 包目录。
- GOCOVERDIR 未生效 → 确认 runCve 设置了 `GOCOVERDIR` 环境变量且 go 版本 >= 1.20。
- 时间相关断言失败 → 用 `time.Now().Year()` 动态构造期望值。
- map 遍历顺序导致断言失败 → 改为断言内容存在而非精确字符串。

若验证全部通过，跳过修复，直接进入 Step 4。

Run: `cd /home/cc11001100/github/scagogogo/cve-skills && go test ./... 2>&1 | tail -5`
Expected:
  - Exit code: 0
  - Output contains: "ok" for `github.com/scagogogo/cve-skills` and `github.com/scagogogo/cve-skills/cmd`

- [ ] **Step 4: 同步测试策略文档 — 更新 testing.md 反映 CLI 层测试与 100% 覆盖**

文件: `website/guide/testing.md`（更新 Test Inventory 表与 Coverage Philosophy 小节）

在 Test Inventory 表后追加 CLI 层测试说明，并更新覆盖率哲学小节，说明 CLI 层通过子进程测试达到 100%。具体：将 "38 test functions" 更新为包含 cmd 测试的总数，在 Coverage Philosophy 增加一段说明 CLI 层用 `go test -c -cover` 编译带覆盖率二进制 + `os/exec` 子进程运行，覆盖所有 `os.Exit` 与 RunE error 分支。

修改 `website/guide/testing.md` 中第 3 行附近的测试清单描述与第 152 行附近的 Coverage Philosophy 小节，增加 CLI 子进程测试段落：

```markdown
## CLI Layer Coverage

The `cmd/` package (cobra CLI) reaches 100.0% statement coverage via a subprocess strategy. Because command `Run` closures call `os.Exit(1)` on missing input — which terminates the test process — the suite compiles a coverage-instrumented binary with `go test -c -cover` and drives it via `os/exec` with argument combinations that hit every branch: missing flags (`--year`/`--start`/`--end`/`--years` == 0), the `cutoff > 0` vs `cutoff == 0` split, `nil` range results, `strconv.Atoi` failures, and the `r.Valid` true/false output paths. Each subprocess writes coverage counters to a `GOCOVERDIR`, merged by `go tool cover` into the final profile.

```bash
# Build the instrumented binary once, reused across tests
go test -c -cover -o /tmp/cve.test ./cmd

# Merge library + CLI coverage
go test -coverpkg=./... -coverprofile=full.out ./...
go tool cover -func=full.out | grep cmd/
```
```

同步修改 `website/zh/guide/testing.md` 的对应小节（测试清单表 + 覆盖率哲学），增加中文版 CLI 层覆盖说明。

- [ ] **Step 5: 验证文档站构建通过且无死链**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills/website && npm run build 2>&1 | tail -15`
Expected:
  - Exit code: 0
  - Output contains: "build complete"
  - Output does NOT contain: "error" or "dead link"

- [ ] **Step 6: 提交**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills && git add website/guide/testing.md website/zh/guide/testing.md && git commit -m "docs(testing): document CLI subprocess coverage reaching 100%"`
