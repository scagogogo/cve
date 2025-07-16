# 比较与排序

这一类函数用于 CVE 编号的比较和排序操作，支持按年份、序列号或综合比较。

## CompareByYear

根据 CVE 的年份比较大小。

### 函数签名

```go
func CompareByYear(cveA, cveB string) int
```

### 参数

- `cveA` (string): 第一个 CVE 编号
- `cveB` (string): 第二个 CVE 编号

### 返回值

- `int`: 比较结果
  - 负数: cveA 年份 < cveB 年份
  - 零: cveA 年份 = cveB 年份  
  - 正数: cveA 年份 > cveB 年份

### 功能描述

`CompareByYear` 函数只比较两个 CVE 的年份部分：
- 提取两个 CVE 的年份并进行数值比较
- 返回年份差值（cveA年份 - cveB年份）
- 不考虑序列号部分

### 使用示例

```go
func main() {
    testCases := []struct {
        cveA, cveB string
        desc       string
    }{
        {"CVE-2020-1111", "CVE-2022-2222", "不同年份"},
        {"CVE-2022-1111", "CVE-2022-2222", "相同年份"},
        {"CVE-2023-1111", "CVE-2021-2222", "A年份更新"},
        {"cve-2022-1111", "CVE-2022-2222", "大小写混合"},
    }
    
    for _, tc := range testCases {
        result := cve.CompareByYear(tc.cveA, tc.cveB)
        var relation string
        if result < 0 {
            relation = "早于"
        } else if result > 0 {
            relation = "晚于"
        } else {
            relation = "同年"
        }
        
        fmt.Printf("%-15s %s %-15s (差值: %d) - %s\n", 
            tc.cveA, relation, tc.cveB, result, tc.desc)
    }
}
```

### 使用场景

- 按年份对 CVE 进行粗略排序
- 年份统计和分析
- 时间范围过滤的辅助函数

---

## SubByYear

计算两个 CVE 的年份差值。

### 函数签名

```go
func SubByYear(cveA, cveB string) int
```

### 参数

- `cveA` (string): 第一个 CVE 编号
- `cveB` (string): 第二个 CVE 编号

### 返回值

- `int`: cveA 年份 - cveB 年份的差值

### 功能描述

`SubByYear` 函数计算两个 CVE 编号的年份差值：
- 功能上等同于 `CompareByYear`
- 提供更直观的函数名
- 常用于计算时间间隔

### 使用示例

```go
func main() {
    pairs := []struct {
        cveA, cveB string
    }{
        {"CVE-2023-1111", "CVE-2020-2222"},
        {"CVE-2020-1111", "CVE-2023-2222"},
        {"CVE-2022-1111", "CVE-2022-2222"},
    }
    
    for _, pair := range pairs {
        diff := cve.SubByYear(pair.cveA, pair.cveB)
        fmt.Printf("%s - %s = %d年\n", pair.cveA, pair.cveB, diff)
    }
}
```

### 使用场景

- 计算 CVE 之间的时间间隔
- 漏洞趋势分析
- 时间序列数据处理

---

## CompareCves

全面比较两个 CVE 编号的大小。

### 函数签名

```go
func CompareCves(cveA, cveB string) int
```

### 参数

- `cveA` (string): 第一个 CVE 编号
- `cveB` (string): 第二个 CVE 编号

### 返回值

- `int`: 比较结果
  - `-1`: cveA < cveB
  - `0`: cveA = cveB
  - `1`: cveA > cveB

### 功能描述

`CompareCves` 函数进行完整的 CVE 比较：
1. 首先比较年份
2. 年份相同时比较序列号
3. 返回标准化的比较结果（-1, 0, 1）

### 比较逻辑

```go
// 伪代码
if yearA != yearB {
    return yearA < yearB ? -1 : 1
}
if seqA != seqB {
    return seqA < seqB ? -1 : 1  
}
return 0
```

### 使用示例

```go
func main() {
    testCases := []struct {
        cveA, cveB string
        desc       string
    }{
        {"CVE-2020-1111", "CVE-2022-2222", "不同年份"},
        {"CVE-2022-1111", "CVE-2022-2222", "相同年份，不同序列号"},
        {"CVE-2022-2222", "CVE-2022-2222", "完全相同"},
        {"CVE-2022-2222", "CVE-2022-1111", "A序列号更大"},
        {"CVE-2023-1111", "CVE-2022-9999", "A年份更新"},
    }
    
    for _, tc := range testCases {
        result := cve.CompareCves(tc.cveA, tc.cveB)
        var relation string
        switch result {
        case -1:
            relation = "<"
        case 0:
            relation = "="
        case 1:
            relation = ">"
        }
        
        fmt.Printf("%-15s %s %-15s (%d) - %s\n", 
            tc.cveA, relation, tc.cveB, result, tc.desc)
    }
}
```

### 使用场景

- 完整的 CVE 排序
- 查找最新或最旧的 CVE
- 实现自定义排序算法
- 数据结构中的比较函数

---

## SortCves

对 CVE 切片进行排序。

### 函数签名

```go
func SortCves(cveSlice []string) []string
```

### 参数

- `cveSlice` ([]string): 要排序的 CVE 编号列表

### 返回值

- `[]string`: 排序后的 CVE 编号列表（新切片）

### 功能描述

`SortCves` 函数对 CVE 列表进行完整排序：
- 创建新的切片，不修改原始数据
- 自动格式化所有 CVE 为标准格式
- 使用 `CompareCves` 进行比较
- 按年份优先，序列号次之的顺序排序

### 使用示例

```go
func main() {
    // 基本排序
    cveList := []string{
        "CVE-2022-2222",
        "cve-2020-1111",  // 小写
        "CVE-2022-1111",
        "CVE-2021-3333",
        "CVE-2020-9999",
    }
    
    fmt.Printf("原始列表: %v\n", cveList)
    
    sorted := cve.SortCves(cveList)
    fmt.Printf("排序后: %v\n", sorted)
    
    // 验证原始列表未被修改
    fmt.Printf("原始列表（未变）: %v\n", cveList)
    
    // 复杂排序示例
    complexList := []string{
        "CVE-2022-12345",
        "CVE-2022-1",      // 短序列号
        "CVE-2021-99999",  // 大序列号但年份早
        "cve-2022-12344",  // 小写且序列号相近
        "CVE-2023-1",      // 新年份
    }
    
    complexSorted := cve.SortCves(complexList)
    fmt.Printf("\n复杂排序:\n")
    fmt.Printf("原始: %v\n", complexList)
    fmt.Printf("排序: %v\n", complexSorted)
}
```

### 排序规则

1. **年份优先**: 年份小的排在前面
2. **序列号次之**: 年份相同时，序列号小的排在前面
3. **格式标准化**: 自动转换为标准大写格式

### 使用场景

- 生成有序的 CVE 报告
- 时间线分析
- 数据展示和可视化
- 查找最早或最新的 CVE

## 实际应用示例

### 1. CVE 时间线分析

```go
func analyzeCveTimeline(cveList []string) {
    // 排序 CVE
    sorted := cve.SortCves(cveList)
    
    fmt.Println("=== CVE 时间线分析 ===")
    fmt.Printf("总计 %d 个 CVE\n", len(sorted))
    
    if len(sorted) == 0 {
        return
    }
    
    // 最早和最新的 CVE
    earliest := sorted[0]
    latest := sorted[len(sorted)-1]
    
    fmt.Printf("最早: %s (%s年)\n", earliest, cve.ExtractCveYear(earliest))
    fmt.Printf("最新: %s (%s年)\n", latest, cve.ExtractCveYear(latest))
    
    // 时间跨度
    yearSpan := cve.SubByYear(latest, earliest)
    fmt.Printf("时间跨度: %d年\n", yearSpan)
    
    // 按年份统计
    yearCount := make(map[string]int)
    for _, cveId := range sorted {
        year := cve.ExtractCveYear(cveId)
        yearCount[year]++
    }
    
    fmt.Println("\n年份分布:")
    for year, count := range yearCount {
        fmt.Printf("  %s年: %d个\n", year, count)
    }
}
```

### 2. 自定义排序

```go
// 按序列号降序排序（年份仍然升序）
func sortBySeqDesc(cveList []string) []string {
    result := make([]string, len(cveList))
    copy(result, cveList)
    
    sort.Slice(result, func(i, j int) bool {
        yearComp := cve.CompareByYear(result[i], result[j])
        if yearComp != 0 {
            return yearComp < 0  // 年份升序
        }
        // 序列号降序
        seqA := cve.ExtractCveSeqAsInt(result[i])
        seqB := cve.ExtractCveSeqAsInt(result[j])
        return seqA > seqB
    })
    
    return result
}

// 只按年份排序，忽略序列号
func sortByYearOnly(cveList []string) []string {
    result := make([]string, len(cveList))
    copy(result, cveList)
    
    sort.Slice(result, func(i, j int) bool {
        return cve.CompareByYear(result[i], result[j]) < 0
    })
    
    return result
}
```

### 3. 查找操作

```go
func findCveOperations(cveList []string) {
    sorted := cve.SortCves(cveList)
    
    // 查找特定年份的第一个和最后一个 CVE
    targetYear := "2022"
    var firstInYear, lastInYear string
    
    for _, cveId := range sorted {
        year := cve.ExtractCveYear(cveId)
        if year == targetYear {
            if firstInYear == "" {
                firstInYear = cveId
            }
            lastInYear = cveId
        }
    }
    
    fmt.Printf("%s年第一个CVE: %s\n", targetYear, firstInYear)
    fmt.Printf("%s年最后一个CVE: %s\n", targetYear, lastInYear)
    
    // 查找中位数 CVE
    if len(sorted) > 0 {
        midIndex := len(sorted) / 2
        median := sorted[midIndex]
        fmt.Printf("中位数CVE: %s\n", median)
    }
}
```

### 4. 比较分析

```go
func compareCveGroups(groupA, groupB []string) {
    sortedA := cve.SortCves(groupA)
    sortedB := cve.SortCves(groupB)
    
    fmt.Println("=== CVE 组比较 ===")
    fmt.Printf("组A: %d个 CVE\n", len(sortedA))
    fmt.Printf("组B: %d个 CVE\n", len(sortedB))
    
    if len(sortedA) > 0 && len(sortedB) > 0 {
        // 比较最早的 CVE
        earliestComp := cve.CompareCves(sortedA[0], sortedB[0])
        if earliestComp < 0 {
            fmt.Printf("组A的最早CVE更早: %s vs %s\n", sortedA[0], sortedB[0])
        } else if earliestComp > 0 {
            fmt.Printf("组B的最早CVE更早: %s vs %s\n", sortedB[0], sortedA[0])
        } else {
            fmt.Printf("两组最早CVE相同: %s\n", sortedA[0])
        }
        
        // 比较最新的 CVE
        latestA := sortedA[len(sortedA)-1]
        latestB := sortedB[len(sortedB)-1]
        latestComp := cve.CompareCves(latestA, latestB)
        
        if latestComp < 0 {
            fmt.Printf("组B的最新CVE更新: %s vs %s\n", latestB, latestA)
        } else if latestComp > 0 {
            fmt.Printf("组A的最新CVE更新: %s vs %s\n", latestA, latestB)
        } else {
            fmt.Printf("两组最新CVE相同: %s\n", latestA)
        }
    }
}
```

## 性能说明

- `CompareByYear` 和 `SubByYear` 性能相当，都很快
- `CompareCves` 包含更多逻辑，但仍然高效
- `SortCves` 的时间复杂度为 O(n log n)，空间复杂度为 O(n)
- 对于大量数据，建议使用 `SortCves` 而不是多次调用比较函数
- 所有函数都是并发安全的

## 最佳实践

### 1. 选择合适的比较函数

```go
// 只需要年份比较时
if cve.CompareByYear(cveA, cveB) < 0 {
    // cveA 年份更早
}

// 需要完整比较时
if cve.CompareCves(cveA, cveB) < 0 {
    // cveA 完全小于 cveB
}
```

### 2. 批量排序

```go
// 好的做法：一次性排序
sorted := cve.SortCves(largeCveList)

// 避免的做法：多次比较
// for i := 0; i < len(largeCveList); i++ {
//     for j := i + 1; j < len(largeCveList); j++ {
//         if cve.CompareCves(largeCveList[i], largeCveList[j]) > 0 {
//             // 交换...
//         }
//     }
// }
```

### 3. 保持数据不变性

```go
// SortCves 返回新切片，原数据不变
original := []string{"CVE-2022-2", "CVE-2021-1"}
sorted := cve.SortCves(original)
// original 保持不变
```
