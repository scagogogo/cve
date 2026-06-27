# 批量验证

这一类函数提供 CVE 编号的批量验证功能，并返回详细的错误报告，适用于导入或处理大型 CVE 数据集时的数据质量检查。

## ValidateCves

批量验证 CVE 编号，返回每个 CVE 的详细验证结果。

### 函数签名

```go
func ValidateCves(cveSlice []string) []CveValidationResult
```

### CveValidationResult 类型

```go
type CveValidationResult struct {
    Cve    string // 原始 CVE 字符串
    Valid  bool   // CVE 是否有效
    Reason string // 无效原因（有效时为空）
}
```

### 参数

- `cveSlice` ([]string): 要验证的 CVE 编号切片

### 返回值

- `[]CveValidationResult`: 每个输入 CVE 的验证结果

### 功能描述

`ValidateCves` 函数验证列表中的每个 CVE，并返回结构化结果，包含是否有效以及无效时的人类可读原因。验证检查包括：格式（必须为 CVE-YYYY-NNNNN）、年份（1999 至当前年份）和序列号（正整数）。

### 使用示例

```go
package main

import (
    "fmt"
    "github.com/scagogogo/cve-skills"
)

func main() {
    cves := []string{"CVE-2022-1234", "not-a-cve", "CVE-1998-1234", "CVE-2023-5678"}

    results := cve.ValidateCves(cves)
    for _, r := range results {
        if r.Valid {
            fmt.Printf("✓ %s 有效\n", r.Cve)
        } else {
            fmt.Printf("✗ %s 无效: %s\n", r.Cve, r.Reason)
        }
    }
    // 输出:
    // ✓ CVE-2022-1234 有效
    // ✗ not-a-cve 无效: CVE 格式无效
    // ✗ CVE-1998-1234 无效: 年份 1998 早于 1999
    // ✓ CVE-2023-5678 有效
}
```

## FilterValidCves

过滤 CVE 列表，仅保留有效的 CVE 编号。

### 函数签名

```go
func FilterValidCves(cveSlice []string) []string
```

### 参数

- `cveSlice` ([]string): 要过滤的 CVE 编号切片

### 返回值

- `[]string`: 仅包含有效的 CVE，格式化为大写

### 功能描述

`FilterValidCves` 函数是一个便捷封装，从列表中过滤掉无效的 CVE，仅返回标准大写格式的有效 CVE。

### 使用示例

```go
raw := []string{"CVE-2022-1234", "invalid", "cve-2023-5678", "CVE-1998-1234"}
valid := cve.FilterValidCves(raw)
fmt.Println(valid)
// 输出: [CVE-2022-1234 CVE-2023-5678]
```
