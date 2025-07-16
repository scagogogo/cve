# API 参考

CVE Utils 提供了一套完整的 CVE 处理函数，涵盖了从基本的格式化验证到复杂的分析处理的各个方面。

## 函数分类

### 🔍 [格式化与验证](/api/format-validate)

用于 CVE 格式标准化和有效性验证的函数：

| 函数 | 描述 |
|------|------|
| `Format(cve string) string` | 将 CVE 转换为标准大写格式 |
| `IsCve(text string) bool` | 判断字符串是否为有效的 CVE 格式 |
| `IsContainsCve(text string) bool` | 判断字符串是否包含 CVE |
| `IsCveYearOk(cve string, cutoff int) bool` | 判断 CVE 年份是否合理 |
| `ValidateCve(cve string) bool` | 全面验证 CVE 编号的合法性 |

### 📝 [提取方法](/api/extract)

从文本或 CVE 中提取信息的函数：

| 函数 | 描述 |
|------|------|
| `ExtractCve(text string) []string` | 从文本中提取所有 CVE 编号 |
| `ExtractFirstCve(text string) string` | 提取第一个 CVE 编号 |
| `ExtractLastCve(text string) string` | 提取最后一个 CVE 编号 |
| `Split(cve string) (year string, seq string)` | 分割 CVE 为年份和序列号 |
| `ExtractCveYear(cve string) string` | 提取 CVE 年份（字符串） |
| `ExtractCveYearAsInt(cve string) int` | 提取 CVE 年份（整数） |
| `ExtractCveSeq(cve string) string` | 提取 CVE 序列号（字符串） |
| `ExtractCveSeqAsInt(cve string) int` | 提取 CVE 序列号（整数） |

### 🔄 [比较与排序](/api/compare-sort)

用于 CVE 比较和排序的函数：

| 函数 | 描述 |
|------|------|
| `CompareByYear(cveA, cveB string) int` | 根据年份比较两个 CVE |
| `SubByYear(cveA, cveB string) int` | 计算两个 CVE 的年份差值 |
| `CompareCves(cveA, cveB string) int` | 全面比较两个 CVE |
| `SortCves(cveSlice []string) []string` | 对 CVE 切片进行排序 |

### 🎯 [过滤与分组](/api/filter-group)

用于 CVE 过滤、分组和去重的函数：

| 函数 | 描述 |
|------|------|
| `FilterCvesByYear(cveSlice []string, year int) []string` | 筛选特定年份的 CVE |
| `FilterCvesByYearRange(cveSlice []string, startYear, endYear int) []string` | 筛选年份范围内的 CVE |
| `GetRecentCves(cveSlice []string, years int) []string` | 获取最近几年的 CVE |
| `GroupByYear(cveSlice []string) map[string][]string` | 按年份分组 CVE |
| `RemoveDuplicateCves(cveSlice []string) []string` | 移除重复的 CVE |

### ⚡ [生成与构造](/api/generate)

用于生成新 CVE 编号的函数：

| 函数 | 描述 |
|------|------|
| `GenerateCve(year int, seq int) string` | 根据年份和序列号生成 CVE |

## 快速参考

### 常用操作

```go
import "github.com/scagogogo/cve"

// 格式化
formatted := cve.Format(" cve-2022-12345 ")  // "CVE-2022-12345"

// 验证
isValid := cve.ValidateCve("CVE-2022-12345")  // true

// 提取
cves := cve.ExtractCve("文本包含 CVE-2022-12345")  // ["CVE-2022-12345"]

// 排序
sorted := cve.SortCves([]string{"CVE-2022-2", "CVE-2021-1"})

// 分组
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

## 版本兼容性

当前 API 版本：`v1.0.0`

- ✅ **稳定 API**：所有公开函数的签名保持稳定
- ✅ **向后兼容**：新版本保持向后兼容
- ✅ **语义版本**：遵循语义版本规范

## 下一步

选择您感兴趣的 API 分类深入了解：

- [格式化与验证](/api/format-validate) - 学习如何验证和标准化 CVE
- [提取方法](/api/extract) - 了解如何从文本中提取 CVE 信息
- [比较与排序](/api/compare-sort) - 掌握 CVE 比较和排序技巧
- [过滤与分组](/api/filter-group) - 学习高级的过滤和分组操作
- [生成与构造](/api/generate) - 了解如何生成新的 CVE 编号

或者查看 [使用示例](/examples/) 了解实际应用场景。
