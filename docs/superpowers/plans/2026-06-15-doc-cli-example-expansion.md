# CVE Utils — Doc Site + CLI + Example Expansion Plan

> **For agentic workers:** REQUIRED SUB-SKILL: `superpowers:subagent-driven-development`
> Steps use checkbox (`- [ ]`) syntax.

**Goal:** 补全文档网站、CLI 和示例代码，使新增的 12 个 API 函数获得完整的文档覆盖、CLI 子命令支持、可运行示例，和此前 19 个已有函数保持一致。

**Architecture:**
- **文档层**：在 `docs/api/` 和 `docs/zh/api/` 下新增 4 个 API 文档页（集合运算、批量验证、范围模式匹配、统计分析），更新 api/index.md 和首页 features，注册侧边栏导航
- **示例层**：在 `examples/` 下创建 12 个编号目录（20-31），每个目录有 `main.go` 可独立运行，跟随现有 01-19 的命名和编码模式
- **CLI 层**：在 `cmd/` 下为新增函数创建 CLI 子命令文件，复用现有 cmd/compare.go、cmd/generate.go 的模式

**Tech Stack:** Go 1.18, Cobra v1.8.1, Vitepress (docs), Markdown (docs)

**Risks:**
- Vitepress 侧边栏配置需要手动注册新页面 → 缓解：修改 config.js 的 sidebar 数组
- CLI 子命令文件间命名可能冲突 → 缓解：每个新增功能独立文件，遵循命名约定

---

### Task 1: Update Documentation Site

**Depends on:** None
**Files:**
- Create: `docs/api/set-operations.md`
- Create: `docs/api/batch-validation.md`
- Create: `docs/api/range-pattern.md`
- Create: `docs/api/statistics.md`
- Create: `docs/zh/api/set-operations.md`
- Create: `docs/zh/api/batch-validation.md`
- Create: `docs/zh/api/range-pattern.md`
- Create: `docs/zh/api/statistics.md`
- Modify: `docs/api/index.md`
- Modify: `docs/zh/api/index.md`
- Modify: `docs/.vitepress/config.js`

- [ ] **Step 1: Create `docs/api/set-operations.md` — Document IntersectCves, UnionCves, DiffCves**

```markdown
# Set Operations

This category of functions provides set operations (intersection, union, difference) for CVE lists, useful when merging or comparing CVE data from multiple sources.

## IntersectCves

Compute the intersection of two CVE lists — returns CVEs that appear in both.

### Function Signature

```go
func IntersectCves(a, b []string) []string
```

### Parameters

- `a` ([]string): First CVE list
- `b` ([]string): Second CVE list

### Return Value

- `[]string`: CVEs present in both lists, deduplicated, formatted to uppercase, and sorted

### Description

The `IntersectCves` function returns CVE identifiers that exist in both input lists. Comparison is case-insensitive, and the result is sorted by year and sequence number.

### Example

```go
package main

import (
    "fmt"
    "github.com/scagogogo/cve"
)

func main() {
    list1 := []string{"CVE-2022-1111", "CVE-2022-2222", "CVE-2022-3333"}
    list2 := []string{"CVE-2022-2222", "CVE-2022-3333", "CVE-2022-4444"}

    common := cve.IntersectCves(list1, list2)
    fmt.Println(common)
    // Output: [CVE-2022-2222 CVE-2022-3333]
}
```

## UnionCves

Compute the union of two CVE lists — returns all CVEs from both lists, deduplicated.

### Function Signature

```go
func UnionCves(a, b []string) []string
```

### Parameters

- `a` ([]string): First CVE list
- `b` ([]string): Second CVE list

### Return Value

- `[]string`: All CVEs from both lists, deduplicated, formatted to uppercase, and sorted

### Description

The `UnionCves` function merges two CVE lists, removing duplicates. Comparison is case-insensitive, and the result is sorted by year and sequence number.

### Example

```go
list1 := []string{"CVE-2022-1111", "CVE-2022-2222"}
list2 := []string{"CVE-2022-2222", "CVE-2022-3333"}

all := cve.UnionCves(list1, list2)
fmt.Println(all)
// Output: [CVE-2022-1111 CVE-2022-2222 CVE-2022-3333]
```

## DiffCves

Compute the difference (a - b) of two CVE lists — returns CVEs in list a that are not in list b.

### Function Signature

```go
func DiffCves(a, b []string) []string
```

### Parameters

- `a` ([]string): The source CVE list (minuend)
- `b` ([]string): The CVE list to exclude (subtrahend)

### Return Value

- `[]string`: CVEs in `a` but not in `b`, deduplicated, formatted to uppercase, and sorted

### Description

The `DiffCves` function finds CVE identifiers that exist in list `a` but not in list `b`. This is useful for detecting new CVEs by comparing current data against historical data. Comparison is case-insensitive.

### Example

```go
current := []string{"CVE-2022-1111", "CVE-2022-2222", "CVE-2022-3333"}
historical := []string{"CVE-2022-2222", "CVE-2022-4444"}

newCves := cve.DiffCves(current, historical)
fmt.Println(newCves)
// Output: [CVE-2022-1111 CVE-2022-3333]
```
```

- [ ] **Step 2: Create `docs/api/batch-validation.md` — Document ValidateCves and FilterValidCves**

```markdown
# Batch Validation

This category of functions provides batch validation of CVE identifiers with detailed error reporting, useful for data quality checks when importing or processing large CVE datasets.

## ValidateCves

Batch validate CVE identifiers, returning detailed results for each.

### Function Signature

```go
func ValidateCves(cveSlice []string) []CveValidationResult
```

### CveValidationResult Type

```go
type CveValidationResult struct {
    Cve    string // The original CVE string
    Valid  bool   // Whether the CVE is valid
    Reason string // Reason for invalidity (empty if valid)
}
```

### Parameters

- `cveSlice` ([]string): Slice of CVE identifiers to validate

### Return Value

- `[]CveValidationResult`: Validation results for each input CVE

### Description

The `ValidateCves` function validates each CVE in the list and returns a structured result containing whether it's valid and, if not, a human-readable reason. Validation checks: format (must be CVE-YYYY-NNNNN), year (1999 to current year), and sequence number (positive integer).

### Example

```go
package main

import (
    "fmt"
    "github.com/scagogogo/cve"
)

func main() {
    cves := []string{"CVE-2022-1234", "not-a-cve", "CVE-1998-1234", "CVE-2023-5678"}

    results := cve.ValidateCves(cves)
    for _, r := range results {
        if r.Valid {
            fmt.Printf("✓ %s is valid\n", r.Cve)
        } else {
            fmt.Printf("✗ %s is invalid: %s\n", r.Cve, r.Reason)
        }
    }
    // Output:
    // ✓ CVE-2022-1234 is valid
    // ✗ not-a-cve is invalid: invalid CVE format
    // ✗ CVE-1998-1234 is invalid: year 1998 is before 1999
    // ✓ CVE-2023-5678 is valid
}
```

## FilterValidCves

Filter a CVE list to keep only valid CVE identifiers.

### Function Signature

```go
func FilterValidCves(cveSlice []string) []string
```

### Parameters

- `cveSlice` ([]string): Slice of CVE identifiers to filter

### Return Value

- `[]string`: Only the valid CVEs, formatted to uppercase

### Description

The `FilterValidCves` function is a convenience wrapper that filters out invalid CVEs from a list, returning only valid ones in standard uppercase format.

### Example

```go
raw := []string{"CVE-2022-1234", "invalid", "cve-2023-5678", "CVE-1998-1234"}
valid := cve.FilterValidCves(raw)
fmt.Println(valid)
// Output: [CVE-2022-1234 CVE-2023-5678]
```
```

- [ ] **Step 3: Create `docs/api/range-pattern.md` — Document ParseCveRange, IsCvesConsecutive, FilterCvesByPattern, FormatSeq**

```markdown
# Range, Pattern & Formatting

This category of functions provides CVE range parsing, wildcard pattern matching, consecutive CVE checking, and sequence number formatting.

## ParseCveRange

Parse a CVE range expression and expand it into a list of individual CVE identifiers.

### Function Signature

```go
func ParseCveRange(rangeExpr string) []string
```

### Parameters

- `rangeExpr` (string): Range expression in one of the supported formats

### Return Value

- `[]string`: All CVE identifiers within the range; nil if the expression is invalid

### Supported Formats

- `CVE-2022-12345 to CVE-2022-12350` (keyword "to")
- `CVE-2022-12345..12350` (double dots)
- `CVE-2022-12345-12350` (dash separator)

### Description

The `ParseCveRange` function parses CVE range expressions commonly found in security bulletins (e.g., "CVEs CVE-2022-1000 through CVE-2022-1050 are affected") and expands them into individual CVE identifiers. Both the start and end are within the same year.

### Example

```go
cves := cve.ParseCveRange("CVE-2022-12345 to CVE-2022-12348")
fmt.Println(cves)
// Output: [CVE-2022-12345 CVE-2022-12346 CVE-2022-12347 CVE-2022-12348]
```

## IsCvesConsecutive

Check if two CVE identifiers are consecutive (same year, adjacent sequence numbers).

### Function Signature

```go
func IsCvesConsecutive(a, b string) bool
```

### Parameters

- `a` (string): First CVE identifier
- `b` (string): Second CVE identifier

### Return Value

- `bool`: true if the two CVEs have the same year and their sequence numbers differ by exactly 1

### Example

```go
fmt.Println(cve.IsCvesConsecutive("CVE-2022-12345", "CVE-2022-12346"))
// Output: true
fmt.Println(cve.IsCvesConsecutive("CVE-2022-12345", "CVE-2022-12347"))
// Output: false
```

## FilterCvesByPattern

Filter CVE identifiers by wildcard pattern.

### Function Signature

```go
func FilterCvesByPattern(cveSlice []string, pattern string) []string
```

### Parameters

- `cveSlice` ([]string): CVE identifiers to filter
- `pattern` (string): Wildcard pattern (e.g., `CVE-2022-*`, `CVE-*-1111`, `CVE-2022-1*`)

### Return Value

- `[]string`: CVEs matching the pattern, formatted and sorted; nil if no matches

### Description

The `FilterCvesByPattern` function supports simple wildcard matching with `*` matching any sequence of characters. Pattern is case-insensitive. Special characters like `.`, `+`, `[` etc. are auto-escaped.

### Example

```go
cves := []string{"CVE-2022-1111", "CVE-2022-2222", "CVE-2023-1111", "CVE-2023-2222"}

cves2022 := cve.FilterCvesByPattern(cves, "CVE-2022-*")
fmt.Println(cves2022)
// Output: [CVE-2022-1111 CVE-2022-2222]

cve1111 := cve.FilterCvesByPattern(cves, "CVE-*-1111")
fmt.Println(cve1111)
// Output: [CVE-2021-1111 CVE-2022-1111]  (if present)
```

## FormatSeq

Format a CVE's sequence number to a fixed width with zero-padding.

### Function Signature

```go
func FormatSeq(cve string, width int) string
```

### Parameters

- `cve` (string): CVE identifier
- `width` (int): Target width for the sequence number (e.g., 6)

### Return Value

- `string`: CVE with zero-padded sequence number; returns original string if invalid

### Example

```go
fmt.Println(cve.FormatSeq("CVE-2022-123", 6))
// Output: CVE-2022-000123
fmt.Println(cve.FormatSeq("CVE-2022-12345", 6))
// Output: CVE-2022-012345
```
```

- [ ] **Step 4: Create `docs/api/statistics.md` — Document CountByYear, YearRange, SeqRange**

```markdown
# Statistical Analysis

This category of functions provides statistical analysis of CVE data, including year-based counting, year range detection, and sequence number range analysis.

## CountByYear

Count the number of CVEs per year.

### Function Signature

```go
func CountByYear(cveSlice []string) map[int]int
```

### Parameters

- `cveSlice` ([]string): Slice of CVE identifiers

### Return Value

- `map[int]int`: Map from year to CVE count

### Description

The `CountByYear` function groups CVE identifiers by year and returns the count for each year. Invalid CVEs (unparseable year) are excluded from the result.

### Example

```go
cves := []string{"CVE-2022-1111", "CVE-2022-2222", "CVE-2021-3333", "CVE-2022-4444"}
counts := cve.CountByYear(cves)
fmt.Println(counts)
// Output: map[2021:1 2022:3]
```

## YearRange

Get the earliest and latest year from a CVE list.

### Function Signature

```go
func YearRange(cveSlice []string) (min, max int)
```

### Parameters

- `cveSlice` ([]string): Slice of CVE identifiers

### Return Value

- `min` (int): Earliest year; 0 if no valid CVEs found
- `max` (int): Latest year; 0 if no valid CVEs found

### Example

```go
cves := []string{"CVE-2020-1111", "CVE-2022-2222", "CVE-2021-3333"}
min, max := cve.YearRange(cves)
fmt.Printf("Range: %d - %d\n", min, max)
// Output: Range: 2020 - 2022
```

## SeqRange

Get the sequence number range for CVEs in a specific year.

### Function Signature

```go
func SeqRange(cveSlice []string, year int) (min, max int)
```

### Parameters

- `cveSlice` ([]string): Slice of CVE identifiers
- `year` (int): Target year to analyze

### Return Value

- `min` (int): Smallest sequence number in the given year; 0 if no CVEs found
- `max` (int): Largest sequence number in the given year; 0 if no CVEs found

### Example

```go
cves := []string{"CVE-2022-1111", "CVE-2022-5555", "CVE-2022-3333", "CVE-2021-9999"}
min, max := cve.SeqRange(cves, 2022)
fmt.Printf("2022 sequence range: %d - %d\n", min, max)
// Output: 2022 sequence range: 1111 - 5555
```
```

- [ ] **Step 5: Create Chinese translations of all 4 new API docs**

Create the following files with equivalent Chinese content:

- `docs/zh/api/set-operations.md` — 集合运算
- `docs/zh/api/batch-validation.md` — 批量验证
- `docs/zh/api/range-pattern.md` — 范围、模式匹配与格式化
- `docs/zh/api/statistics.md` — 统计分析

Each file follows the same structure as the English counterpart but with Chinese function descriptions, parameter explanations, and example comments. Use the existing `docs/zh/api/filter-group.md` as the translation style reference.

- [ ] **Step 6: Update `docs/api/index.md` — Add links to new API pages**

In the function categories list, add after the existing `Generation & Construction` section:

```markdown
- **Set Operations**: Compute intersection, union, and difference of CVE lists — [View Docs](/api/set-operations)
- **Batch Validation**: Validate CVEs in batch with detailed error reporting — [View Docs](/api/batch-validation)
- **Range & Pattern**: Parse CVE range expressions, filter by wildcard patterns — [View Docs](/api/range-pattern)
- **Statistical Analysis**: Count by year, get year/sequence ranges — [View Docs](/api/statistics)
```

- [ ] **Step 7: Update `docs/zh/api/index.md` — Add Chinese links to new API pages**

Follow the same pattern as Step 6 but in Chinese:

```markdown
- **集合运算**: 计算 CVE 列表的交集、并集和差集 — [查看文档](/zh/api/set-operations)
- **批量验证**: 批量验证 CVE 并返回详细的错误报告 — [查看文档](/zh/api/batch-validation)
- **范围与模式**: 解析 CVE 范围表达式，通配符模式筛选 — [查看文档](/zh/api/range-pattern)
- **统计分析**: 按年份计数，获取年份和序列号范围 — [查看文档](/zh/api/statistics)
```

- [ ] **Step 8: Update `docs/.vitepress/config.js` — Register new API pages in sidebar**

Add the following sidebar entries after the existing `Generation & Construction` item in BOTH the `root` (English) and `zh` (Chinese) locale sections:

English section (around line 42):
```javascript
{ text: 'Set Operations', link: '/api/set-operations' },
{ text: 'Batch Validation', link: '/api/batch-validation' },
{ text: 'Range & Pattern', link: '/api/range-pattern' },
{ text: 'Statistical Analysis', link: '/api/statistics' },
```

Chinese section (around line 94):
```javascript
{ text: '集合运算', link: '/zh/api/set-operations' },
{ text: '批量验证', link: '/zh/api/batch-validation' },
{ text: '范围与模式', link: '/zh/api/range-pattern' },
{ text: '统计分析', link: '/zh/api/statistics' },
```

---

### Task 2: Create Runnable Examples for All New Functions

**Depends on:** None (can parallel with Task 1)
**Files:**
- Create: `examples/20_intersect_cves/main.go`
- Create: `examples/21_union_cves/main.go`
- Create: `examples/22_diff_cves/main.go`
- Create: `examples/23_validate_cves/main.go`
- Create: `examples/24_filter_valid_cves/main.go`
- Create: `examples/25_parse_cve_range/main.go`
- Create: `examples/26_is_cves_consecutive/main.go`
- Create: `examples/27_count_by_year/main.go`
- Create: `examples/28_year_range/main.go`
- Create: `examples/29_seq_range/main.go`
- Create: `examples/30_filter_by_pattern/main.go`
- Create: `examples/31_format_seq/main.go`
- Modify: `examples/README.md`

**Pattern to follow:** Each example is a standalone `main.go` in a numbered directory, imports `github.com/scagogogo/cve`, prints a descriptive header, demonstrates the function with realistic data, and shows use case scenarios. See `examples/16_remove_duplicate_cves/main.go` for the reference pattern.

- [ ] **Step 1: Create `examples/20_intersect_cves/main.go` — Demonstrate IntersectCves**

File: `examples/20_intersect_cves/main.go`

```go
package main

import (
	"fmt"

	"github.com/scagogogo/cve"
)

func main() {
	fmt.Println("=== CVE 交集运算 (Intersection) ===\n")

	// 模拟来自两个不同安全扫描工具的CVE报告
	scannerA := []string{"CVE-2022-1111", "CVE-2022-2222", "CVE-2022-3333", "CVE-2022-4444"}
	scannerB := []string{"CVE-2022-2222", "CVE-2022-3333", "CVE-2022-5555", "CVE-2022-6666"}

	fmt.Println("扫描器A发现的CVE:", scannerA)
	fmt.Println("扫描器B发现的CVE:", scannerB)

	// 求交集：两个扫描器都发现的CVE（更可信）
	common := cve.IntersectCves(scannerA, scannerB)
	fmt.Printf("\n共同发现的CVE (交集): %v\n", common)
	fmt.Printf("共同发现数量: %d\n", len(common))

	// 演示大小写不敏感
	fmt.Println("\n--- 大小写不敏感示例 ---")
	list1 := []string{"cve-2022-1111", "CVE-2022-2222", "Cve-2022-3333"}
	list2 := []string{"CVE-2022-1111", "cve-2022-3333", "CVE-2022-4444"}
	fmt.Println("列表1:", list1)
	fmt.Println("列表2:", list2)
	fmt.Printf("交集: %v\n", cve.IntersectCves(list1, list2))

	// 空列表场景
	fmt.Println("\n--- 空列表场景 ---")
	fmt.Printf("空列表交集: %v\n", cve.IntersectCves([]string{}, []string{"CVE-2022-1111"}))
}
```

- [ ] **Step 2: Create `examples/21_union_cves/main.go` — Demonstrate UnionCves**

```go
package main

import (
	"fmt"

	"github.com/scagogogo/cve"
)

func main() {
	fmt.Println("=== CVE 并集运算 (Union) ===\n")

	// 模拟多个团队的CVE数据
	teamA := []string{"CVE-2023-1001", "CVE-2023-1002", "CVE-2023-1003"}
	teamB := []string{"CVE-2023-1003", "CVE-2023-1004", "CVE-2023-1005"}
	teamC := []string{"CVE-2023-1004", "CVE-2023-1005", "CVE-2023-1006"}

	fmt.Println("团队A的CVE:", teamA)
	fmt.Println("团队B的CVE:", teamB)
	fmt.Println("团队C的CVE:", teamC)

	// 合并：先合并AB，再与C合并
	merged := cve.UnionCves(teamA, teamB)
	merged = cve.UnionCves(merged, teamC)
	fmt.Printf("\n全部团队的CVE (并集): %v\n", merged)
	fmt.Printf("总唯一CVE数量: %d\n", len(merged))

	// 演示重复去除
	fmt.Println("\n--- 去重效果 ---")
	withDups := []string{"CVE-2022-1111", "cve-2022-1111", "CVE-2022-1111", "CVE-2022-2222"}
	unique := cve.UnionCves(withDups, []string{})
	fmt.Printf("原始 (含重复): %v\n", withDups)
	fmt.Printf("并集 (去重后): %v\n", unique)
}
```

- [ ] **Step 3: Create `examples/22_diff_cves/main.go` — Demonstrate DiffCves**

```go
package main

import (
	"fmt"

	"github.com/scagogogo/cve"
)

func main() {
	fmt.Println("=== CVE 差集运算 (Difference) ===\n")

	// 模拟当前扫描结果与历史数据的对比
	currentScan := []string{"CVE-2024-1001", "CVE-2024-1002", "CVE-2024-1003", "CVE-2024-1004", "CVE-2024-1005"}
	previousScan := []string{"CVE-2024-1001", "CVE-2024-1003", "CVE-2024-1005"}

	fmt.Println("当前扫描结果:", currentScan)
	fmt.Println("前一次扫描结果:", previousScan)

	// 差集：找出新出现的CVE
	newCves := cve.DiffCves(currentScan, previousScan)
	fmt.Printf("\n新出现的CVE (差集): %v\n", newCves)
	fmt.Printf("新增数量: %d\n", len(newCves))

	// 反向差集：找出已修复的CVE
	fixedCves := cve.DiffCves(previousScan, currentScan)
	fmt.Printf("\n已修复的CVE (反向差集): %v\n", fixedCves)
	fmt.Printf("修复数量: %d\n", len(fixedCves))

	// 演示：差集为空表示完全覆盖
	fmt.Println("\n--- 完全覆盖场景 ---")
	subset := []string{"CVE-2022-1111", "CVE-2022-2222"}
	superset := []string{"CVE-2022-1111", "CVE-2022-2222", "CVE-2022-3333"}
	fmt.Printf("%v \\ %v = %v\n", subset, superset, cve.DiffCves(subset, superset))
}
```

- [ ] **Step 4: Create `examples/23_validate_cves/main.go` — Demonstrate ValidateCves**

```go
package main

import (
	"fmt"

	"github.com/scagogogo/cve"
)

func main() {
	fmt.Println("=== 批量CVE验证 ===\n")

	// 模拟从数据源导入的CVE列表
	rawCves := []string{
		"CVE-2022-1234",
		"cve-2023-5678",
		"CVE-1998-1234",  // 年份太早
		"not-a-cve",       // 格式错误
		"CVE-2099-9999",   // 未来年份
		"CVE-2022-ABCD",   // 非数字序列号
		"CVE-2022-0",      // 序列号为0
		" CVE-2024-8888 ", // 有效但含空格
	}

	fmt.Println("验证以下CVE:\n")
	results := cve.ValidateCves(rawCves)

	validCount := 0
	for _, r := range results {
		if r.Valid {
			fmt.Printf("  ✓ %-25s 有效\n", r.Cve)
			validCount++
		} else {
			fmt.Printf("  ✗ %-25s 无效 — %s\n", r.Cve, r.Reason)
		}
	}

	fmt.Printf("\n统计: %d/%d 有效\n", validCount, len(rawCves))
}
```

- [ ] **Step 5: Create `examples/24_filter_valid_cves/main.go` — Demonstrate FilterValidCves**

```go
package main

import (
	"fmt"

	"github.com/scagogogo/cve"
)

func main() {
	fmt.Println("=== 过滤有效CVE ===\n")

	// 从混合数据中提取有效CVE
	mixedData := []string{
		"CVE-2022-1234",
		"invalid-data",
		"cve-2023-5678",
		"CVE-1998-0001",
		"CVE-2024-9999",
		"random-text",
		"CVE-2099-1234",
	}

	fmt.Println("混合数据:", mixedData)

	validCves := cve.FilterValidCves(mixedData)
	fmt.Printf("\n有效CVE: %v\n", validCves)
	fmt.Printf("有效数量: %d / %d\n", len(validCves), len(mixedData))

	// 对比原始函数
	fmt.Println("\n--- 与 ValidateCve 对比 ---")
	for _, item := range mixedData {
		status := "✗"
		if cve.ValidateCve(item) {
			status = "✓"
		}
		fmt.Printf("  %s %s\n", status, item)
	}
}
```

- [ ] **Step 6: Create `examples/25_parse_cve_range/main.go` — Demonstrate ParseCveRange**

```go
package main

import (
	"fmt"

	"github.com/scagogogo/cve"
)

func main() {
	fmt.Println("=== CVE 范围解析 ===\n")

	// 安全公告常见的范围表达方式
	fmt.Println("--- 使用 'to' 关键字 ---")
	range1 := "CVE-2022-1000 to CVE-2022-1005"
	result1 := cve.ParseCveRange(range1)
	fmt.Printf("输入: %s\n", range1)
	fmt.Printf("输出: %v\n", result1)

	fmt.Println("\n--- 使用 '..' 双点号 ---")
	range2 := "CVE-2022-2000..2003"
	result2 := cve.ParseCveRange(range2)
	fmt.Printf("输入: %s\n", range2)
	fmt.Printf("输出: %v\n", result2)

	fmt.Println("\n--- 使用 '-' 连字符 ---")
	range3 := "CVE-2022-3000-3002"
	result3 := cve.ParseCveRange(range3)
	fmt.Printf("输入: %s\n", range3)
	fmt.Printf("输出: %v\n", result3)

	// 无效表达式
	fmt.Println("\n--- 无效输入处理 ---")
	fmt.Printf("空字符串: %v\n", cve.ParseCveRange(""))
	fmt.Printf("单CVE格式: %v\n", cve.ParseCveRange("CVE-2022-12345"))
	fmt.Printf("反向范围: %v\n", cve.ParseCveRange("CVE-2022-1005..1000"))

	// 实际应用场景
	fmt.Println("\n--- 实际应用: 统计范围内CVE总数 ---")
	securityBulletin := "CVE-2023-5000 to CVE-2023-5999"
	affectedCves := cve.ParseCveRange(securityBulletin)
	fmt.Printf("安全公告: %s\n", securityBulletin)
	fmt.Printf("受影响CVE总数: %d\n", len(affectedCves))
	fmt.Printf("范围: %s 到 %s\n", affectedCves[0], affectedCves[len(affectedCves)-1])
}
```

- [ ] **Step 7: Create `examples/26_is_cves_consecutive/main.go` — Demonstrate IsCvesConsecutive**

```go
package main

import (
	"fmt"

	"github.com/scagogogo/cve"
)

func main() {
	fmt.Println("=== CVE 连续性判断 ===\n")

	// 检查连续的CVE对
	pairs := []struct {
		a, b string
	}{
		{"CVE-2022-12345", "CVE-2022-12346"}, // 连续
		{"CVE-2022-12345", "CVE-2022-12347"}, // 不连续
		{"CVE-2022-12345", "CVE-2023-12345"}, // 不同年份
		{"CVE-2022-12346", "CVE-2022-12345"}, // 顺序反转但连续
		{"CVE-2022-12345", "CVE-2022-12345"}, // 相同CVE
	}

	for _, p := range pairs {
		consecutive := cve.IsCvesConsecutive(p.a, p.b)
		mark := "✗"
		if consecutive {
			mark = "✓"
		}
		fmt.Printf("  %s %s <-> %s: 连续=%v\n", mark, p.a, p.b, consecutive)
	}

	// 检测列表中可合并为范围的连续CVE
	fmt.Println("\n--- 检测可合并列表 ---")
	cveList := []string{
		"CVE-2022-1001", "CVE-2022-1002", "CVE-2022-1003", // 连续
		"CVE-2022-2001", "CVE-2022-2003",                   // 不连续
	}
	fmt.Println("CVE列表:", cveList)

	for i := 0; i < len(cveList)-1; i++ {
		if cve.IsCvesConsecutive(cveList[i], cveList[i+1]) {
			fmt.Printf("  %s 和 %s 连续\n", cveList[i], cveList[i+1])
		} else {
			fmt.Printf("  %s 和 %s 不连续\n", cveList[i], cveList[i+1])
		}
	}
}
```

- [ ] **Step 8: Create `examples/27_count_by_year/main.go` — Demonstrate CountByYear**

```go
package main

import (
	"fmt"
	"sort"

	"github.com/scagogogo/cve"
)

func main() {
	fmt.Println("=== 按年份统计CVE ===\n")

	cveList := []string{
		"CVE-2019-1001", "CVE-2019-1002",
		"CVE-2020-1001", "CVE-2020-1002", "CVE-2020-1003",
		"CVE-2021-1001", "CVE-2021-1002",
		"CVE-2022-1001", "CVE-2022-1002", "CVE-2022-1003", "CVE-2022-1004",
		"CVE-2023-1001",
		"CVE-2024-1001", "CVE-2024-1002", "CVE-2024-1003",
	}

	counts := cve.CountByYear(cveList)

	// 按年份排序显示
	var years []int
	for y := range counts {
		years = append(years, y)
	}
	sort.Ints(years)

	fmt.Println("年份分布:")
	fmt.Println("年份    | 数量 | 柱状图")
	fmt.Println("--------|------|------")
	for _, year := range years {
		count := counts[year]
		bar := ""
		for i := 0; i < count; i++ {
			bar += "█"
		}
		fmt.Printf("%d    | %4d | %s\n", year, count, bar)
	}

	fmt.Printf("\n总年份跨度: %d 年\n", len(counts))
	fmt.Printf("总计CVE: %d\n", len(cveList))
}
```

- [ ] **Step 9: Create `examples/28_year_range/main.go` — Demonstrate YearRange**

```go
package main

import (
	"fmt"

	"github.com/scagogogo/cve"
)

func main() {
	fmt.Println("=== CVE 年份范围 ===\n")

	// 模拟一个跨年份的CVE数据集
	cveList := []string{
		"CVE-2015-1001",
		"CVE-2018-2001",
		"CVE-2020-3001",
		"CVE-2022-4001",
		"CVE-2024-5001",
		"CVE-2025-6001",
	}

	minYear, maxYear := cve.YearRange(cveList)
	fmt.Println("CVE列表:", cveList)
	fmt.Printf("\n年份范围: %d - %d\n", minYear, maxYear)
	fmt.Printf("时间跨度: %d 年\n", maxYear-minYear)

	// 空列表处理
	fmt.Println("\n--- 边界情况 ---")
	minE, maxE := cve.YearRange([]string{})
	fmt.Printf("空列表: min=%d, max=%d\n", minE, maxE)
}
```

- [ ] **Step 10: Create `examples/29_seq_range/main.go` — Demonstrate SeqRange**

```go
package main

import (
	"fmt"

	"github.com/scagogogo/cve"
)

func main() {
	fmt.Println("=== CVE 序列号范围 ===\n")

	// 模拟某一年份的CVE分配情况
	cveList := []string{
		"CVE-2022-1001", "CVE-2022-5050", "CVE-2022-3025",
		"CVE-2022-8888", "CVE-2022-1500", "CVE-2021-9999",
		"CVE-2023-1234", "CVE-2022-7777",
	}

	targetYears := []int{2022, 2021, 2023, 2020}

	for _, year := range targetYears {
		minSeq, maxSeq := cve.SeqRange(cveList, year)
		if minSeq == 0 && maxSeq == 0 {
			fmt.Printf("%d 年: 无CVE数据\n", year)
		} else {
			fmt.Printf("%d 年: 序列号范围 %d - %d (共 %d 个可能位置)\n",
				year, minSeq, maxSeq, maxSeq-minSeq+1)
		}
	}

	// 统计某个年份所有CVE
	fmt.Println("\n--- 列出2022年所有CVE ---")
	cves2022 := cve.FilterCvesByYear(cveList, 2022)
	sorted := cve.SortCves(cves2022)
	fmt.Println(sorted)
}
```

- [ ] **Step 11: Create `examples/30_filter_by_pattern/main.go` — Demonstrate FilterCvesByPattern**

```go
package main

import (
	"fmt"

	"github.com/scagogogo/cve"
)

func main() {
	fmt.Println("=== 通配符模式匹配 CVE ===\n")

	cveList := []string{
		"CVE-2021-1111", "CVE-2021-2222",
		"CVE-2022-1111", "CVE-2022-1122", "CVE-2022-2222", "CVE-2022-3333",
		"CVE-2023-1111", "CVE-2023-2222", "CVE-2023-3333",
	}

	fmt.Printf("CVE列表 (共 %d 个):\n", len(cveList))
	fmt.Println("  ", cveList)

	// 按年份筛选
	fmt.Println("\n--- 按年份筛选: CVE-2022-* ---")
	fmt.Println("  ", cve.FilterCvesByPattern(cveList, "CVE-2022-*"))

	// 按序列号筛选
	fmt.Println("\n--- 按序列号筛选: CVE-*-1111 ---")
	fmt.Println("  ", cve.FilterCvesByPattern(cveList, "CVE-*-1111"))

	// 前缀匹配
	fmt.Println("\n--- 前缀匹配: CVE-2022-11* ---")
	fmt.Println("  ", cve.FilterCvesByPattern(cveList, "CVE-2022-11*"))

	// 精确匹配（无通配符）
	fmt.Println("\n--- 精确匹配: CVE-2022-2222 ---")
	fmt.Println("  ", cve.FilterCvesByPattern(cveList, "CVE-2022-2222"))

	// 无匹配
	fmt.Println("\n--- 无匹配: CVE-2020-* ---")
	fmt.Println("  ", cve.FilterCvesByPattern(cveList, "CVE-2020-*"))
}
```

- [ ] **Step 12: Create `examples/31_format_seq/main.go` — Demonstrate FormatSeq**

```go
package main

import (
	"fmt"

	"github.com/scagogogo/cve"
)

func main() {
	fmt.Println("=== CVE 序列号格式化 ===\n")

	// 不同序列号长度的CVE格式化
	cves := []string{
		"CVE-2022-1",
		"CVE-2022-12",
		"CVE-2022-123",
		"CVE-2022-1234",
		"CVE-2022-12345",
		"CVE-2022-123456",
	}

	fmt.Println("宽度为 6 的格式化效果:")
	fmt.Println("原始            | 格式化后")
	fmt.Println("----------------|---------")
	for _, cve := range cves {
		formatted := cve.FormatSeq(cve, 6)
		fmt.Printf("%-16s| %s\n", cve, formatted)
	}

	// 不同宽度
	fmt.Println("\n--- 不同宽度效果 (CVE-2022-123) ---")
	for _, width := range []int{4, 5, 6, 8} {
		fmt.Printf("  宽度 %d: %s\n", width, cve.FormatSeq("CVE-2022-123", width))
	}

	// 无效输入
	fmt.Println("\n--- 无效输入 ---")
	fmt.Printf("  'not-a-cve' -> %s\n", cve.FormatSeq("not-a-cve", 6))
}
```

- [ ] **Step 13: Update `examples/README.md` — Add descriptions for examples 20-31**

Add the following entries after the existing list (after entry 19):

```markdown
| 20 | [Intersect CVEs](./20_intersect_cves/) | 求 CVE 列表交集 |
| 21 | [Union CVEs](./21_union_cves/) | 求 CVE 列表并集 |
| 22 | [Diff CVEs](./22_diff_cves/) | 求 CVE 列表差集 |
| 23 | [Validate CVEs](./23_validate_cves/) | 批量验证 CVE 及错误原因 |
| 24 | [Filter Valid CVEs](./24_filter_valid_cves/) | 从列表中过滤有效 CVE |
| 25 | [Parse CVE Range](./25_parse_cve_range/) | 解析 CVE 范围表达式 |
| 26 | [Is CVEs Consecutive](./26_is_cves_consecutive/) | 判断 CVE 是否连续 |
| 27 | [Count by Year](./27_count_by_year/) | 按年份统计 CVE 数量 |
| 28 | [Year Range](./28_year_range/) | 获取 CVE 年份范围 |
| 29 | [Seq Range](./29_seq_range/) | 获取序列号范围 |
| 30 | [Filter by Pattern](./30_filter_by_pattern/) | 通配符模式筛选 CVE |
| 31 | [Format Seq](./31_format_seq/) | CVE 序列号补零格式化 |
```

---

### Task 3: Add CLI Subcommands for New Functions

**Depends on:** None (can parallel with Task 1, Task 2)
**Files:**
- Create: `cmd/set.go`
- Create: `cmd/validate_batch.go`
- Create: `cmd/range.go`
- Create: `cmd/stats.go`
- Create: `cmd/pattern.go`

**Pattern to follow:** Existing `cmd/compare.go`, `cmd/generate.go` using Cobra `*cobra.Command` with `RunE`, reading from positional args or stdin via `cmd/helpers.go:readInputs`.

- [ ] **Step 1: Create `cmd/set.go` — Add intersect, union, diff subcommands**

Create a new file `cmd/set.go` with three Cobra subcommands:
- `cve intersect <list1> <list2>` — computes intersection
- `cve union <list1> <list2>` — computes union
- `cve diff <list1> <list2>` — computes difference

Each reads two CVE lists from args (comma-separated or from stdin), calls the corresponding `cve.*` function, and prints results. Register all three in `init()`.

```go
package cmd

import (
	"fmt"
	"strings"

	"github.com/scagogogo/cve"
	"github.com/spf13/cobra"
)

var intersectCmd = &cobra.Command{
	Use:   "intersect <list1> <list2>",
	Short: "Compute intersection of two CVE lists",
	Long:  `Compute the intersection (CVEs in both lists) of two CVE lists. Input lists may be comma-separated or read from stdin.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		inputs := readInputs(args)
		if len(inputs) < 2 {
			return fmt.Errorf("requires exactly 2 arguments (two CVE lists)")
		}
		list1 := strings.Split(inputs[0], ",")
		list2 := strings.Split(inputs[1], ",")
		result := cve.IntersectCves(list1, list2)
		for _, v := range result {
			fmt.Println(v)
		}
		return nil
	},
}

var unionCmd = &cobra.Command{
	Use:   "union <list1> <list2>",
	Short: "Compute union of two CVE lists",
	Long:  `Compute the union (all CVEs from both lists) of two CVE lists.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		inputs := readInputs(args)
		if len(inputs) < 2 {
			return fmt.Errorf("requires exactly 2 arguments (two CVE lists)")
		}
		list1 := strings.Split(inputs[0], ",")
		list2 := strings.Split(inputs[1], ",")
		result := cve.UnionCves(list1, list2)
		for _, v := range result {
			fmt.Println(v)
		}
		return nil
	},
}

var diffCmd = &cobra.Command{
	Use:   "diff <list1> <list2>",
	Short: "Compute difference (a - b) of two CVE lists",
	Long:  `Compute the difference (CVEs in list1 but not in list2) of two CVE lists.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		inputs := readInputs(args)
		if len(inputs) < 2 {
			return fmt.Errorf("requires exactly 2 arguments (two CVE lists)")
		}
		list1 := strings.Split(inputs[0], ",")
		list2 := strings.Split(inputs[1], ",")
		result := cve.DiffCves(list1, list2)
		for _, v := range result {
			fmt.Println(v)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(intersectCmd)
	rootCmd.AddCommand(unionCmd)
	rootCmd.AddCommand(diffCmd)
}
```

- [ ] **Step 2: Create `cmd/validate_batch.go` — Add validate-batch and filter-valid subcommands**

```go
package cmd

import (
	"fmt"
	"strings"

	"github.com/scagogogo/cve"
	"github.com/spf13/cobra"
)

var validateBatchCmd = &cobra.Command{
	Use:   "validate-batch <cve-list>",
	Short: "Batch validate CVE identifiers",
	Long:  `Validate a batch of CVE identifiers and report detailed results for each.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		inputs := readInputs(args)
		if len(inputs) == 0 {
			return fmt.Errorf("requires at least 1 argument (CVE list)")
		}
		var cveList []string
		for _, input := range inputs {
			cveList = append(cveList, strings.Split(input, ",")...)
		}
		results := cve.ValidateCves(cveList)
		for _, r := range results {
			if r.Valid {
				fmt.Printf("✓ %s\n", r.Cve)
			} else {
				fmt.Printf("✗ %s — %s\n", r.Cve, r.Reason)
			}
		}
		return nil
	},
}

var filterValidCmd = &cobra.Command{
	Use:   "filter-valid <cve-list>",
	Short: "Filter out only valid CVE identifiers",
	Long:  `Filter a list of CVE identifiers, keeping only valid ones.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		inputs := readInputs(args)
		if len(inputs) == 0 {
			return fmt.Errorf("requires at least 1 argument (CVE list)")
		}
		var cveList []string
		for _, input := range inputs {
			cveList = append(cveList, strings.Split(input, ",")...)
		}
		result := cve.FilterValidCves(cveList)
		for _, v := range result {
			fmt.Println(v)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(validateBatchCmd)
	rootCmd.AddCommand(filterValidCmd)
}
```

- [ ] **Step 3: Create `cmd/range.go` — Add parse-range and is-consecutive subcommands**

```go
package cmd

import (
	"fmt"
	"strings"

	"github.com/scagogogo/cve"
	"github.com/spf13/cobra"
)

var parseRangeCmd = &cobra.Command{
	Use:   "parse-range <range-expr>",
	Short: "Parse a CVE range expression",
	Long:  `Parse a CVE range expression (supports "to", "..", "-" syntax) and expand into individual CVEs.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		inputs := readInputs(args)
		if len(inputs) == 0 {
			return fmt.Errorf("requires at least 1 argument (range expression)")
		}
		rangeExpr := strings.TrimSpace(inputs[0])
		result := cve.ParseCveRange(rangeExpr)
		if result == nil {
			return fmt.Errorf("invalid range expression: %s", rangeExpr)
		}
		for _, v := range result {
			fmt.Println(v)
		}
		return nil
	},
}

var isConsecutiveCmd = &cobra.Command{
	Use:   "is-consecutive <cve-a> <cve-b>",
	Short: "Check if two CVEs are consecutive",
	Long:  `Check whether two CVE identifiers are consecutive (same year, adjacent sequence numbers).`,
	RunE: func(cmd *cobra.Command, args []string) error {
		inputs := readInputs(args)
		if len(inputs) < 2 {
			return fmt.Errorf("requires exactly 2 CVE identifiers")
		}
		result := cve.IsCvesConsecutive(inputs[0], inputs[1])
		if result {
			fmt.Printf("%s and %s are consecutive\n", inputs[0], inputs[1])
		} else {
			fmt.Printf("%s and %s are NOT consecutive\n", inputs[0], inputs[1])
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(parseRangeCmd)
	rootCmd.AddCommand(isConsecutiveCmd)
}
```

- [ ] **Step 4: Create `cmd/stats.go` — Add count-by-year, year-range, seq-range subcommands**

```go
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/scagogogo/cve"
	"github.com/spf13/cobra"
)

var countByYearCmd = &cobra.Command{
	Use:   "count-by-year <cve-list>",
	Short: "Count CVEs by year",
	Long:  `Group and count CVE identifiers by year.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		inputs := readInputs(args)
		if len(inputs) == 0 {
			return fmt.Errorf("requires at least 1 argument (CVE list)")
		}
		var cveList []string
		for _, input := range inputs {
			cveList = append(cveList, strings.Split(input, ",")...)
		}
		counts := cve.CountByYear(cveList)
		for year, count := range counts {
			fmt.Printf("%d: %d\n", year, count)
		}
		return nil
	},
}

var yearRangeCmd = &cobra.Command{
	Use:   "year-range <cve-list>",
	Short: "Get the earliest and latest year of CVEs",
	Long:  `Get the earliest and latest year from a CVE list.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		inputs := readInputs(args)
		if len(inputs) == 0 {
			return fmt.Errorf("requires at least 1 argument (CVE list)")
		}
		var cveList []string
		for _, input := range inputs {
			cveList = append(cveList, strings.Split(input, ",")...)
		}
		min, max := cve.YearRange(cveList)
		fmt.Printf("Year range: %d - %d (span: %d years)\n", min, max, max-min)
		return nil
	},
}

var seqRangeCmd = &cobra.Command{
	Use:   "seq-range <year> <cve-list>",
	Short: "Get sequence number range for a given year",
	Long:  `Get the smallest and largest sequence number for CVEs in a specific year.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		inputs := readInputs(args)
		if len(inputs) < 2 {
			return fmt.Errorf("requires year and CVE list")
		}
		year, err := strconv.Atoi(strings.TrimSpace(inputs[0]))
		if err != nil {
			return fmt.Errorf("invalid year: %s", inputs[0])
		}
		var cveList []string
		for _, input := range inputs[1:] {
			cveList = append(cveList, strings.Split(input, ",")...)
		}
		min, max := cve.SeqRange(cveList, year)
		fmt.Printf("Year %d sequence range: %d - %d\n", year, min, max)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(countByYearCmd)
	rootCmd.AddCommand(yearRangeCmd)
	rootCmd.AddCommand(seqRangeCmd)
}
```

- [ ] **Step 5: Create `cmd/pattern.go` — Add filter-pattern and format-seq subcommands**

```go
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/scagogogo/cve"
	"github.com/spf13/cobra"
)

var filterPatternCmd = &cobra.Command{
	Use:   "filter-pattern <pattern> <cve-list>",
	Short: "Filter CVEs by wildcard pattern",
	Long:  `Filter CVE identifiers using wildcard pattern matching (e.g., "CVE-2022-*").`,
	RunE: func(cmd *cobra.Command, args []string) error {
		inputs := readInputs(args)
		if len(inputs) < 2 {
			return fmt.Errorf("requires pattern and CVE list")
		}
		pattern := strings.TrimSpace(inputs[0])
		var cveList []string
		for _, input := range inputs[1:] {
			cveList = append(cveList, strings.Split(input, ",")...)
		}
		result := cve.FilterCvesByPattern(cveList, pattern)
		for _, v := range result {
			fmt.Println(v)
		}
		return nil
	},
}

var formatSeqCmd = &cobra.Command{
	Use:   "format-seq <width> <cve>",
	Short: "Format CVE sequence number with zero-padding",
	Long:  `Format a CVE's sequence number to a fixed width with leading zeros.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		inputs := readInputs(args)
		if len(inputs) < 2 {
			return fmt.Errorf("requires width and CVE identifier")
		}
		width, err := strconv.Atoi(strings.TrimSpace(inputs[0]))
		if err != nil {
			return fmt.Errorf("invalid width: %s", inputs[0])
		}
		result := cve.FormatSeq(inputs[1], width)
		fmt.Println(result)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(filterPatternCmd)
	rootCmd.AddCommand(formatSeqCmd)
}
```

---

### Task 4: Verify Everything Builds and Tests Pass

**Depends on:** Task 1, Task 2, Task 3
**Files:** None (verification only)

- [ ] **Step 1: Run full test suite**
Run: `go test ./...`
Expected:
  - Exit code: 0
  - Output contains: "ok  github.com/scagogogo/cve"
  - Output does NOT contain: "FAIL"

- [ ] **Step 2: Verify all new Go files compile**
Run: `go build ./...`
Expected:
  - Exit code: 0
  - Output is empty (no errors)

- [ ] **Step 3: Verify test coverage on core library**
Run: `go test -cover github.com/scagogogo/cve`
Expected:
  - Exit code: 0
  - Output contains: "coverage:" with >= 90%

- [ ] **Step 4: Verify all examples compile**
Run: `for d in examples/2*/; do (cd "$d" && go build -o /dev/null .) || echo "FAIL: $d"; done`
Expected:
  - Exit code: 0
  - No "FAIL:" output

- [ ] **Step 5: Commit all changes**
Run: `git add docs/api/ docs/zh/api/ docs/.vitepress/config.js examples/ cmd/ && git commit -m "docs: add API docs, CLI commands, and examples for new functions (set ops, batch validation, range, stats, pattern, format)"`
