# 过滤与分组

这一类函数用于 CVE 的过滤、分组和去重操作，帮助您按各种条件组织和筛选 CVE 数据。

## FilterCvesByYear

筛选特定年份的 CVE。

### 函数签名

```go
func FilterCvesByYear(cveSlice []string, year int) []string
```

### 参数

- `cveSlice` ([]string): CVE 编号列表
- `year` (int): 要筛选的年份

### 返回值

- `[]string`: 指定年份的 CVE 编号列表

### 功能描述

`FilterCvesByYear` 函数从 CVE 列表中筛选出指定年份的所有 CVE：
- 自动格式化所有 CVE 为标准格式
- 只返回匹配指定年份的 CVE
- 保持原有顺序

### 使用示例

```go
func main() {
    cveList := []string{
        "CVE-2020-1111",
        "CVE-2021-2222",
        "cve-2021-3333",  // 小写
        "CVE-2022-4444",
        "CVE-2021-5555",
    }
    
    fmt.Printf("原始列表: %v\n", cveList)
    
    // 筛选2021年的 CVE
    cves2021 := cve.FilterCvesByYear(cveList, 2021)
    fmt.Printf("2021年的 CVE: %v\n", cves2021)
    // 输出: [CVE-2021-2222 CVE-2021-3333 CVE-2021-5555]
    
    // 筛选不存在的年份
    cves2025 := cve.FilterCvesByYear(cveList, 2025)
    fmt.Printf("2025年的 CVE: %v (长度: %d)\n", cves2025, len(cves2025))
    
    // 统计各年份的数量
    years := []int{2020, 2021, 2022, 2023}
    for _, year := range years {
        filtered := cve.FilterCvesByYear(cveList, year)
        fmt.Printf("%d年: %d个\n", year, len(filtered))
    }
}
```

### 使用场景

- 生成年度漏洞报告
- 按年份分析漏洞趋势
- 筛选特定时期的 CVE
- 数据分析和统计

---

## FilterCvesByYearRange

筛选指定年份范围内的 CVE。

### 函数签名

```go
func FilterCvesByYearRange(cveSlice []string, startYear, endYear int) []string
```

### 参数

- `cveSlice` ([]string): CVE 编号列表
- `startYear` (int): 开始年份（包含）
- `endYear` (int): 结束年份（包含）

### 返回值

- `[]string`: 指定年份范围内的 CVE 编号列表

### 功能描述

`FilterCvesByYearRange` 函数筛选年份范围内的 CVE：
- 包含开始年份和结束年份
- 自动格式化 CVE 为标准格式
- 如果 startYear > endYear，返回空列表

### 使用示例

```go
func main() {
    cveList := []string{
        "CVE-2019-1111",
        "CVE-2020-2222",
        "CVE-2021-3333",
        "CVE-2022-4444",
        "CVE-2023-5555",
        "CVE-2024-6666",
    }
    
    fmt.Printf("原始列表: %v\n", cveList)
    
    // 筛选2020-2022年的 CVE
    range2020to2022 := cve.FilterCvesByYearRange(cveList, 2020, 2022)
    fmt.Printf("2020-2022年: %v\n", range2020to2022)
    
    // 筛选单年（等同于 FilterCvesByYear）
    single2021 := cve.FilterCvesByYearRange(cveList, 2021, 2021)
    fmt.Printf("2021年: %v\n", single2021)
    
    // 筛选最近几年
    recent := cve.FilterCvesByYearRange(cveList, 2022, 2024)
    fmt.Printf("2022-2024年: %v\n", recent)
    
    // 无效范围
    invalid := cve.FilterCvesByYearRange(cveList, 2022, 2020)
    fmt.Printf("无效范围(2022-2020): %v (长度: %d)\n", invalid, len(invalid))
}
```

### 使用场景

- 分析特定时期的漏洞
- 生成时间段报告
- 趋势分析
- 数据切片和分段处理

---

## GetRecentCves

获取最近几年的 CVE。

### 函数签名

```go
func GetRecentCves(cveSlice []string, years int) []string
```

### 参数

- `cveSlice` ([]string): CVE 编号列表
- `years` (int): 最近几年（从当前年份往前计算）

### 返回值

- `[]string`: 最近几年的 CVE 编号列表

### 功能描述

`GetRecentCves` 函数获取最近指定年数的 CVE：
- 基于当前系统时间计算
- 包含当前年份
- 内部调用 `FilterCvesByYearRange`

### 计算逻辑

```go
currentYear := time.Now().Year()
startYear := currentYear - years + 1
endYear := currentYear
```

### 使用示例

```go
func main() {
    cveList := []string{
        "CVE-2019-1111",
        "CVE-2020-2222", 
        "CVE-2021-3333",
        "CVE-2022-4444",
        "CVE-2023-5555",
        "CVE-2024-6666",
    }
    
    currentYear := time.Now().Year()
    fmt.Printf("当前年份: %d\n", currentYear)
    fmt.Printf("原始列表: %v\n", cveList)
    
    // 获取最近2年的 CVE
    recent2 := cve.GetRecentCves(cveList, 2)
    fmt.Printf("最近2年: %v\n", recent2)
    
    // 获取最近3年的 CVE
    recent3 := cve.GetRecentCves(cveList, 3)
    fmt.Printf("最近3年: %v\n", recent3)
    
    // 获取今年的 CVE
    thisYear := cve.GetRecentCves(cveList, 1)
    fmt.Printf("今年: %v\n", thisYear)
    
    // 获取所有年份（使用大数值）
    allRecent := cve.GetRecentCves(cveList, 100)
    fmt.Printf("所有CVE: %v\n", allRecent)
}
```

### 使用场景

- 关注最新的安全漏洞
- 生成近期漏洞报告
- 实时监控和告警
- 优先级排序（新漏洞优先）

---

## GroupByYear

按年份对 CVE 进行分组。

### 函数签名

```go
func GroupByYear(cveSlice []string) map[string][]string
```

### 参数

- `cveSlice` ([]string): 要分组的 CVE 编号列表

### 返回值

- `map[string][]string`: 按年份分组的 CVE 编号映射表

### 功能描述

`GroupByYear` 函数将 CVE 列表按年份分组：
- 键为年份字符串（如 "2022"）
- 值为该年份的 CVE 列表
- 自动格式化所有 CVE 为标准格式
- 保持每组内的原有顺序

### 使用示例

```go
func main() {
    cveList := []string{
        "CVE-2021-1111",
        "cve-2022-2222",  // 小写
        "CVE-2021-3333",
        "CVE-2022-4444",
        "CVE-2023-5555",
        "CVE-2021-6666",
    }
    
    fmt.Printf("原始列表: %v\n", cveList)
    
    grouped := cve.GroupByYear(cveList)
    
    fmt.Println("\n按年份分组:")
    for year, cves := range grouped {
        fmt.Printf("  %s年 (%d个): %v\n", year, len(cves), cves)
    }
    
    // 按年份顺序遍历
    fmt.Println("\n按年份顺序:")
    years := make([]string, 0, len(grouped))
    for year := range grouped {
        years = append(years, year)
    }
    sort.Strings(years)  // 年份排序
    
    for _, year := range years {
        cves := grouped[year]
        fmt.Printf("%s年: %v\n", year, cves)
    }
    
    // 统计信息
    fmt.Println("\n统计信息:")
    totalCves := 0
    for year, cves := range grouped {
        count := len(cves)
        totalCves += count
        fmt.Printf("%s年: %d个 CVE\n", year, count)
    }
    fmt.Printf("总计: %d个 CVE，分布在 %d 个年份\n", totalCves, len(grouped))
}
```

### 使用场景

- 生成年度统计报告
- 可视化年份分布
- 按年份组织数据
- 趋势分析和对比

### 高级用法

```go
func analyzeYearlyDistribution(cveList []string) {
    grouped := cve.GroupByYear(cveList)
    
    // 找出CVE最多的年份
    maxYear := ""
    maxCount := 0
    for year, cves := range grouped {
        if len(cves) > maxCount {
            maxCount = len(cves)
            maxYear = year
        }
    }
    
    fmt.Printf("CVE最多的年份: %s年 (%d个)\n", maxYear, maxCount)
    
    // 计算平均每年的CVE数量
    if len(grouped) > 0 {
        totalCves := 0
        for _, cves := range grouped {
            totalCves += len(cves)
        }
        avgPerYear := float64(totalCves) / float64(len(grouped))
        fmt.Printf("平均每年: %.1f个 CVE\n", avgPerYear)
    }
}
```

---

## RemoveDuplicateCves

移除重复的 CVE 编号。

### 函数签名

```go
func RemoveDuplicateCves(cveSlice []string) []string
```

### 参数

- `cveSlice` ([]string): 可能包含重复项的 CVE 编号列表

### 返回值

- `[]string`: 去重后的 CVE 编号列表

### 功能描述

`RemoveDuplicateCves` 函数移除 CVE 列表中的重复项：
- 不区分大小写（"cve-2022-1" 和 "CVE-2022-1" 视为重复）
- 保持第一次出现的顺序
- 自动格式化为标准格式
- 使用 map 实现，时间复杂度 O(n)

### 使用示例

```go
func main() {
    cveList := []string{
        "CVE-2022-1111",
        "cve-2022-1111",  // 重复（大小写不同）
        "CVE-2022-2222",
        "CVE-2022-1111",  // 重复
        "CVE-2023-3333",
        "cve-2023-3333",  // 重复（大小写不同）
        "CVE-2022-2222",  // 重复
    }
    
    fmt.Printf("原始列表 (%d个): %v\n", len(cveList), cveList)
    
    unique := cve.RemoveDuplicateCves(cveList)
    fmt.Printf("去重后 (%d个): %v\n", len(unique), unique)
    
    // 验证去重效果
    fmt.Printf("去重前后数量: %d -> %d (减少了 %d 个)\n", 
        len(cveList), len(unique), len(cveList)-len(unique))
    
    // 空列表和单元素列表
    empty := cve.RemoveDuplicateCves([]string{})
    single := cve.RemoveDuplicateCves([]string{"CVE-2022-1111"})
    fmt.Printf("空列表去重: %v\n", empty)
    fmt.Printf("单元素去重: %v\n", single)
}
```

### 使用场景

- 合并多个来源的 CVE 列表
- 数据清洗和预处理
- 避免重复处理相同的 CVE
- 统计唯一 CVE 数量

### 性能特性

- 时间复杂度：O(n)
- 空间复杂度：O(n)
- 保持插入顺序
- 内存效率高

## 实际应用示例

### 1. 综合数据分析

```go
func comprehensiveAnalysis(cveList []string) {
    fmt.Println("=== CVE 数据综合分析 ===")
    
    // 1. 基本统计
    fmt.Printf("原始数据: %d个 CVE\n", len(cveList))
    
    // 2. 去重
    unique := cve.RemoveDuplicateCves(cveList)
    fmt.Printf("去重后: %d个 CVE (去除了 %d 个重复)\n", 
        len(unique), len(cveList)-len(unique))
    
    // 3. 按年份分组
    grouped := cve.GroupByYear(unique)
    fmt.Printf("涉及年份: %d个\n", len(grouped))
    
    // 4. 年份分布
    fmt.Println("\n年份分布:")
    for year, cves := range grouped {
        fmt.Printf("  %s年: %d个\n", year, len(cves))
    }
    
    // 5. 最近几年的趋势
    recent3 := cve.GetRecentCves(unique, 3)
    recent2 := cve.GetRecentCves(unique, 2)
    recent1 := cve.GetRecentCves(unique, 1)
    
    fmt.Printf("\n时间趋势:\n")
    fmt.Printf("  最近1年: %d个\n", len(recent1))
    fmt.Printf("  最近2年: %d个\n", len(recent2))
    fmt.Printf("  最近3年: %d个\n", len(recent3))
    
    // 6. 特定年份分析
    currentYear := time.Now().Year()
    thisYear := cve.FilterCvesByYear(unique, currentYear)
    lastYear := cve.FilterCvesByYear(unique, currentYear-1)
    
    fmt.Printf("\n年度对比:\n")
    fmt.Printf("  %d年: %d个\n", currentYear, len(thisYear))
    fmt.Printf("  %d年: %d个\n", currentYear-1, len(lastYear))
    
    if len(lastYear) > 0 {
        change := len(thisYear) - len(lastYear)
        changePercent := float64(change) / float64(len(lastYear)) * 100
        fmt.Printf("  同比变化: %+d个 (%+.1f%%)\n", change, changePercent)
    }
}
```

### 2. 数据清洗流水线

```go
func cleanAndOrganizeCves(rawData []string) map[string][]string {
    fmt.Println("=== CVE 数据清洗流水线 ===")
    
    // 1. 提取所有可能的 CVE
    var allCves []string
    for _, text := range rawData {
        extracted := cve.ExtractCve(text)
        allCves = append(allCves, extracted...)
    }
    fmt.Printf("步骤1 - 提取: %d个 CVE\n", len(allCves))
    
    // 2. 去重
    unique := cve.RemoveDuplicateCves(allCves)
    fmt.Printf("步骤2 - 去重: %d个 CVE (去除 %d 个重复)\n", 
        len(unique), len(allCves)-len(unique))
    
    // 3. 验证有效性
    var valid []string
    for _, cveId := range unique {
        if cve.ValidateCve(cveId) {
            valid = append(valid, cveId)
        }
    }
    fmt.Printf("步骤3 - 验证: %d个有效 CVE (过滤 %d 个无效)\n", 
        len(valid), len(unique)-len(valid))
    
    // 4. 排序
    sorted := cve.SortCves(valid)
    fmt.Printf("步骤4 - 排序: 完成\n")
    
    // 5. 按年份分组
    grouped := cve.GroupByYear(sorted)
    fmt.Printf("步骤5 - 分组: %d个年份\n", len(grouped))
    
    return grouped
}
```

### 3. 时间范围过滤器

```go
type TimeRangeFilter struct {
    StartYear int
    EndYear   int
    RecentYears int
}

func (f *TimeRangeFilter) Apply(cveList []string) []string {
    if f.RecentYears > 0 {
        return cve.GetRecentCves(cveList, f.RecentYears)
    }
    
    if f.StartYear > 0 && f.EndYear > 0 {
        return cve.FilterCvesByYearRange(cveList, f.StartYear, f.EndYear)
    }
    
    if f.StartYear > 0 {
        return cve.FilterCvesByYear(cveList, f.StartYear)
    }
    
    return cveList
}

func main() {
    cveList := []string{
        "CVE-2020-1111", "CVE-2021-2222", "CVE-2022-3333", 
        "CVE-2023-4444", "CVE-2024-5555",
    }
    
    // 不同的过滤条件
    filters := []TimeRangeFilter{
        {RecentYears: 2},           // 最近2年
        {StartYear: 2021, EndYear: 2023}, // 2021-2023年
        {StartYear: 2022},          // 2022年
    }
    
    for i, filter := range filters {
        result := filter.Apply(cveList)
        fmt.Printf("过滤器%d结果: %v\n", i+1, result)
    }
}
```

### 4. 统计报告生成器

```go
func generateStatisticsReport(cveList []string) {
    fmt.Println("=== CVE 统计报告 ===")
    
    // 数据预处理
    unique := cve.RemoveDuplicateCves(cveList)
    grouped := cve.GroupByYear(unique)
    
    // 基本统计
    fmt.Printf("总计: %d个唯一 CVE\n", len(unique))
    fmt.Printf("年份跨度: %d个年份\n", len(grouped))
    
    // 年份统计
    fmt.Println("\n年份分布:")
    years := make([]string, 0, len(grouped))
    for year := range grouped {
        years = append(years, year)
    }
    sort.Strings(years)
    
    for _, year := range years {
        count := len(grouped[year])
        percentage := float64(count) / float64(len(unique)) * 100
        fmt.Printf("  %s年: %3d个 (%5.1f%%)\n", year, count, percentage)
    }
    
    // 趋势分析
    fmt.Println("\n趋势分析:")
    for i := 1; i <= 5; i++ {
        recent := cve.GetRecentCves(unique, i)
        fmt.Printf("  最近%d年: %d个\n", i, len(recent))
    }
    
    // 活跃度分析
    if len(years) >= 2 {
        fmt.Println("\n活跃度分析:")
        recentYear := years[len(years)-1]
        previousYear := years[len(years)-2]
        
        recentCount := len(grouped[recentYear])
        previousCount := len(grouped[previousYear])
        
        fmt.Printf("  %s年: %d个\n", recentYear, recentCount)
        fmt.Printf("  %s年: %d个\n", previousYear, previousCount)
        
        if previousCount > 0 {
            change := float64(recentCount-previousCount) / float64(previousCount) * 100
            fmt.Printf("  同比变化: %+.1f%%\n", change)
        }
    }
}
```

## 性能说明

- `FilterCvesByYear`: O(n) 时间复杂度，遍历一次列表
- `FilterCvesByYearRange`: O(n) 时间复杂度，遍历一次列表
- `GetRecentCves`: 等同于 `FilterCvesByYearRange` 的性能
- `GroupByYear`: O(n) 时间复杂度，O(n) 空间复杂度
- `RemoveDuplicateCves`: O(n) 时间复杂度，使用 map 去重

## 最佳实践

### 1. 组合使用过滤函数

```go
// 获取最近2年的唯一CVE并按年份分组
recent := cve.GetRecentCves(cveList, 2)
unique := cve.RemoveDuplicateCves(recent)
grouped := cve.GroupByYear(unique)
```

### 2. 数据预处理

```go
// 在分析前先清洗数据
func preprocessCves(rawCves []string) []string {
    unique := cve.RemoveDuplicateCves(rawCves)
    
    // 过滤有效的CVE
    var valid []string
    for _, cveId := range unique {
        if cve.ValidateCve(cveId) {
            valid = append(valid, cveId)
        }
    }
    
    return cve.SortCves(valid)
}
```

### 3. 避免重复计算

```go
// 好的做法：缓存分组结果
grouped := cve.GroupByYear(cveList)
for year, cves := range grouped {
    // 使用 cves 进行处理
}

// 避免的做法：重复过滤
// for year := 2020; year <= 2024; year++ {
//     cves := cve.FilterCvesByYear(cveList, year) // 每次都遍历整个列表
// }
```
