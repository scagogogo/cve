# API Reference

CVE Utils provides a complete set of CVE processing functions, covering everything from basic format validation to complex analysis and processing.

## Function Categories

### 🔍 [Format & Validation](/api/format-validate)

Functions for CVE format standardization and validity verification:

| Function | Description |
|----------|-------------|
| `Format(cve string) string` | Convert CVE to standard uppercase format |
| `IsCve(text string) bool` | Check if string is a valid CVE format |
| `IsContainsCve(text string) bool` | Check if string contains CVE |
| `IsCveYearOk(cve string, cutoff int) bool` | Check if CVE year is reasonable |
| `ValidateCve(cve string) bool` | Comprehensive validation of CVE identifier |

### 📝 [Extraction Methods](/api/extract)

Functions for extracting information from text or CVE:

| Function | Description |
|----------|-------------|
| `ExtractCve(text string) []string` | Extract all CVE identifiers from text |
| `ExtractFirstCve(text string) string` | Extract the first CVE identifier |
| `ExtractLastCve(text string) string` | Extract the last CVE identifier |
| `Split(cve string) (year string, seq string)` | Split CVE into year and sequence |
| `ExtractCveYear(cve string) string` | Extract CVE year (string) |
| `ExtractCveYearAsInt(cve string) int` | Extract CVE year (integer) |
| `ExtractCveSeq(cve string) string` | Extract CVE sequence (string) |
| `ExtractCveSeqAsInt(cve string) int` | Extract CVE sequence (integer) |

### 🔄 [Comparison & Sorting](/api/compare-sort)

Functions for CVE comparison and sorting:

| Function | Description |
|----------|-------------|
| `CompareByYear(cveA, cveB string) int` | Compare two CVEs by year |
| `SubByYear(cveA, cveB string) int` | Calculate year difference between two CVEs |
| `CompareCves(cveA, cveB string) int` | Comprehensive comparison of two CVEs |
| `SortCves(cveSlice []string) []string` | Sort CVE slice |

### 🎯 [Filtering & Grouping](/api/filter-group)

Functions for CVE filtering, grouping, and deduplication:

| Function | Description |
|----------|-------------|
| `FilterCvesByYear(cveSlice []string, year int) []string` | Filter CVEs by specific year |
| `FilterCvesByYearRange(cveSlice []string, startYear, endYear int) []string` | Filter CVEs by year range |
| `GetRecentCves(cveSlice []string, years int) []string` | Get CVEs from recent years |
| `GroupByYear(cveSlice []string) map[string][]string` | Group CVEs by year |
| `RemoveDuplicateCves(cveSlice []string) []string` | Remove duplicate CVEs |

### ⚡ [Generation & Construction](/api/generate)

Functions for generating new CVE identifiers:

| Function | Description |
|----------|-------------|
| `GenerateCve(year int, seq int) string` | Generate CVE from year and sequence |

## Quick Reference

### Common Operations

```go
import "github.com/scagogogo/cve"

// Format
formatted := cve.Format(" cve-2022-12345 ")  // "CVE-2022-12345"

// Validate
isValid := cve.ValidateCve("CVE-2022-12345")  // true

// Extract
cves := cve.ExtractCve("Text contains CVE-2022-12345")  // ["CVE-2022-12345"]

// Sort
sorted := cve.SortCves([]string{"CVE-2022-2", "CVE-2021-1"})

// Group
grouped := cve.GroupByYear([]string{"CVE-2021-1", "CVE-2022-1"})
```

### 函数返回值说明

| 返回类型 | 说明 | 示例 |
|----------|------|------|
| `string` | 单个字符串结果，无效输入返回空字符串 | `"CVE-2022-12345"` |
| `[]string` | 字符串切片，无结果返回空切片 | `["CVE-2022-1", "CVE-2022-2"]` |
| `bool` | 布尔值，表示是/否或真/假 | `true` |
| `int` | 整数，无效输入通常返回 0 | `2022` |
| `map[string][]string` | 字符串到字符串切片的映射 | `{"2022": ["CVE-2022-1"]}` |

### 错误处理

CVE Utils 的函数设计为对无效输入有良好的容错性：

- 字符串函数对无效输入返回空字符串 `""`
- 整数函数对无效输入返回 `0`
- 切片函数对无效输入返回空切片 `[]string{}`
- 布尔函数返回 `false` 表示无效或不匹配

### 性能特性

- **内存效率**：函数避免不必要的内存分配
- **并发安全**：所有函数都是并发安全的（无状态）
- **正则表达式优化**：内部使用编译后的正则表达式
- **批量处理**：支持高效的批量操作

## 使用模式

### 1. 数据清洗流水线

```go
func cleanCveData(rawData []string) []string {
    // 1. 提取所有可能的 CVE
    var allCves []string
    for _, text := range rawData {
        cves := cve.ExtractCve(text)
        allCves = append(allCves, cves...)
    }
    
    // 2. 去重
    unique := cve.RemoveDuplicateCves(allCves)
    
    // 3. 验证
    var valid []string
    for _, cveId := range unique {
        if cve.ValidateCve(cveId) {
            valid = append(valid, cveId)
        }
    }
    
    // 4. 排序
    return cve.SortCves(valid)
}
```

### 2. 条件过滤

```go
func filterRecentHighPriority(cveList []string) []string {
    // 获取最近2年的 CVE
    recent := cve.GetRecentCves(cveList, 2)
    
    // 进一步过滤（示例：序列号大于10000的）
    var highPriority []string
    for _, cveId := range recent {
        if cve.ExtractCveSeqAsInt(cveId) > 10000 {
            highPriority = append(highPriority, cveId)
        }
    }
    
    return highPriority
}
```

### 3. 统计分析

```go
func analyzeCveDistribution(cveList []string) {
    // 按年份分组
    grouped := cve.GroupByYear(cveList)
    
    // 统计每年的数量
    for year, cves := range grouped {
        fmt.Printf("%s年: %d个 CVE\n", year, len(cves))
        
        // 分析序列号范围
        if len(cves) > 0 {
            sorted := cve.SortCves(cves)
            firstSeq := cve.ExtractCveSeqAsInt(sorted[0])
            lastSeq := cve.ExtractCveSeqAsInt(sorted[len(sorted)-1])
            fmt.Printf("  序列号范围: %d - %d\n", firstSeq, lastSeq)
        }
    }
}
```

## Version Compatibility

Current API version: `v1.0.0`

- ✅ **Stable API**: All public function signatures remain stable
- ✅ **Backward Compatible**: New versions maintain backward compatibility
- ✅ **Semantic Versioning**: Follows semantic versioning specification

## Next Steps

Choose the API category you're interested in to learn more:

- [Format & Validation](/api/format-validate) - Learn how to validate and standardize CVEs
- [Extraction Methods](/api/extract) - Understand how to extract CVE information from text
- [Comparison & Sorting](/api/compare-sort) - Master CVE comparison and sorting techniques
- [Filtering & Grouping](/api/filter-group) - Learn advanced filtering and grouping operations
- [Generation & Construction](/api/generate) - Understand how to generate new CVE identifiers

Or check out [Examples](/examples/) to see real-world usage scenarios.
