# 基本使用

本指南将详细介绍 CVE Utils 的基本使用方法和最佳实践。

## 导入包

首先在您的 Go 代码中导入 CVE Utils：

```go
import "github.com/scagogogo/cve"
```

## 核心概念

### CVE 格式规范

CVE 编号遵循以下格式：
- 格式：`CVE-YYYY-NNNN`
- `CVE`：固定前缀（不区分大小写）
- `YYYY`：4位年份（1970年至今）
- `NNNN`：序列号（至少4位数字）

有效的 CVE 示例：
- `CVE-2022-12345`
- `CVE-2021-44228`
- `CVE-2023-1234`

### 函数分类

CVE Utils 的函数按功能分为五大类：

1. **格式化与验证**：处理 CVE 格式标准化和有效性验证
2. **提取方法**：从文本中提取 CVE 信息
3. **比较与排序**：对 CVE 进行比较和排序操作
4. **过滤与分组**：按条件筛选和分组 CVE
5. **生成与构造**：创建新的 CVE 编号

## 基本操作

### 1. 格式化 CVE

`Format` 函数将 CVE 转换为标准大写格式：

```go
package main

import (
    "fmt"
    "github.com/scagogogo/cve"
)

func main() {
    // 各种输入格式
    inputs := []string{
        " cve-2022-12345 ",  // 有空格
        "CVE-2022-12345",    // 已经是标准格式
        "cVe-2022-12345",    // 混合大小写
    }
    
    for _, input := range inputs {
        formatted := cve.Format(input)
        fmt.Printf("'%s' -> '%s'\n", input, formatted)
    }
}
```

输出：
```
' cve-2022-12345 ' -> 'CVE-2022-12345'
'CVE-2022-12345' -> 'CVE-2022-12345'
'cVe-2022-12345' -> 'CVE-2022-12345'
```

### 2. 验证 CVE

#### 基本格式验证

```go
func validateExamples() {
    testCases := []string{
        "CVE-2022-12345",           // 有效
        " CVE-2022-12345 ",         // 有效（有空格）
        "包含CVE-2022-12345的文本",   // 无效（包含其他文本）
        "2022-12345",               // 无效（缺少前缀）
        "CVE-2022-ABC",             // 无效（序列号不是数字）
    }
    
    for _, testCase := range testCases {
        isValid := cve.IsCve(testCase)
        fmt.Printf("'%s' -> %t\n", testCase, isValid)
    }
}
```

#### 检查文本是否包含 CVE

```go
func containsExamples() {
    texts := []string{
        "系统受到CVE-2022-12345影响",
        "这个文本不包含任何CVE",
        "多个漏洞：CVE-2021-1111 和 CVE-2022-2222",
    }
    
    for _, text := range texts {
        contains := cve.IsContainsCve(text)
        fmt.Printf("'%s' -> %t\n", text, contains)
    }
}
```

#### 全面验证

```go
func comprehensiveValidation() {
    testCases := []string{
        "CVE-2022-12345",  // 有效
        "CVE-1969-12345",  // 无效（年份太早）
        "CVE-2099-12345",  // 无效（年份太晚）
        "CVE-2022-0",      // 无效（序列号为0）
    }
    
    for _, testCase := range testCases {
        isValid := cve.ValidateCve(testCase)
        fmt.Printf("'%s' -> %t\n", testCase, isValid)
    }
}
```

### 3. 从文本中提取 CVE

#### 提取所有 CVE

```go
func extractAllExample() {
    text := `
    安全公告：系统受到多个漏洞影响
    - CVE-2021-44228 (Log4j)
    - cve-2022-12345 (自定义组件)
    - CVE-2023-1234 (第三方库)
    请尽快更新补丁。
    `
    
    cves := cve.ExtractCve(text)
    fmt.Printf("提取的 CVE (%d 个): %v\n", len(cves), cves)
    // 输出: 提取的 CVE (3 个): [CVE-2021-44228 CVE-2022-12345 CVE-2023-1234]
}
```

#### 提取第一个和最后一个 CVE

```go
func extractFirstLastExample() {
    text := "漏洞包括 CVE-2021-1111、CVE-2022-2222 和 CVE-2023-3333"
    
    first := cve.ExtractFirstCve(text)
    last := cve.ExtractLastCve(text)
    
    fmt.Printf("第一个 CVE: %s\n", first)  // CVE-2021-1111
    fmt.Printf("最后一个 CVE: %s\n", last) // CVE-2023-3333
}
```

### 4. 分解 CVE

#### 分割年份和序列号

```go
func splitExample() {
    cveId := "CVE-2022-12345"
    
    year, seq := cve.Split(cveId)
    fmt.Printf("CVE: %s\n", cveId)
    fmt.Printf("年份: %s\n", year)      // 2022
    fmt.Printf("序列号: %s\n", seq)     // 12345
    
    // 获取整数类型
    yearInt := cve.ExtractCveYearAsInt(cveId)
    seqInt := cve.ExtractCveSeqAsInt(cveId)
    fmt.Printf("年份(int): %d\n", yearInt)  // 2022
    fmt.Printf("序列号(int): %d\n", seqInt) // 12345
}
```

### 5. 比较和排序

#### 比较两个 CVE

```go
func compareExample() {
    cveA := "CVE-2020-1111"
    cveB := "CVE-2022-2222"
    
    // 比较年份
    yearComp := cve.CompareByYear(cveA, cveB)
    fmt.Printf("年份比较 %s vs %s: %d\n", cveA, cveB, yearComp) // -2
    
    // 全面比较
    fullComp := cve.CompareCves(cveA, cveB)
    fmt.Printf("全面比较 %s vs %s: %d\n", cveA, cveB, fullComp) // -1
    
    // 年份差值
    yearDiff := cve.SubByYear(cveB, cveA)
    fmt.Printf("年份差值: %d\n", yearDiff) // 2
}
```

#### 排序 CVE 列表

```go
func sortExample() {
    cveList := []string{
        "CVE-2022-2222",
        "cve-2020-1111",  // 小写
        "CVE-2022-1111",
        "CVE-2021-3333",
    }
    
    fmt.Printf("原始列表: %v\n", cveList)
    
    sorted := cve.SortCves(cveList)
    fmt.Printf("排序后: %v\n", sorted)
    // 输出: [CVE-2020-1111 CVE-2021-3333 CVE-2022-1111 CVE-2022-2222]
}
```

### 6. 过滤和分组

#### 按年份过滤

```go
func filterExample() {
    cveList := []string{
        "CVE-2020-1111",
        "CVE-2021-2222",
        "CVE-2021-3333",
        "CVE-2022-4444",
    }
    
    // 筛选2021年的 CVE
    cves2021 := cve.FilterCvesByYear(cveList, 2021)
    fmt.Printf("2021年的 CVE: %v\n", cves2021)
    // 输出: [CVE-2021-2222 CVE-2021-3333]
    
    // 筛选年份范围
    recentCves := cve.FilterCvesByYearRange(cveList, 2021, 2022)
    fmt.Printf("2021-2022年的 CVE: %v\n", recentCves)
    // 输出: [CVE-2021-2222 CVE-2021-3333 CVE-2022-4444]
    
    // 获取最近2年的 CVE
    recent := cve.GetRecentCves(cveList, 2)
    fmt.Printf("最近2年的 CVE: %v\n", recent)
}
```

#### 按年份分组

```go
func groupExample() {
    cveList := []string{
        "CVE-2021-1111",
        "CVE-2022-2222",
        "CVE-2021-3333",
        "CVE-2022-4444",
    }
    
    grouped := cve.GroupByYear(cveList)
    fmt.Println("按年份分组:")
    for year, cves := range grouped {
        fmt.Printf("  %s: %v\n", year, cves)
    }
    // 输出:
    // 2021: [CVE-2021-1111 CVE-2021-3333]
    // 2022: [CVE-2022-2222 CVE-2022-4444]
}
```

#### 去除重复

```go
func deduplicateExample() {
    cveList := []string{
        "CVE-2022-1111",
        "cve-2022-1111",  // 重复（大小写不同）
        "CVE-2022-2222",
        "CVE-2022-1111",  // 重复
    }
    
    fmt.Printf("原始列表 (%d 个): %v\n", len(cveList), cveList)
    
    unique := cve.RemoveDuplicateCves(cveList)
    fmt.Printf("去重后 (%d 个): %v\n", len(unique), unique)
    // 输出: 去重后 (2 个): [CVE-2022-1111 CVE-2022-2222]
}
```

### 7. 生成 CVE

```go
func generateExample() {
    // 生成新的 CVE 编号
    newCve := cve.GenerateCve(2024, 12345)
    fmt.Printf("生成的 CVE: %s\n", newCve) // CVE-2024-12345
    
    // 批量生成
    for i := 1; i <= 5; i++ {
        cveId := cve.GenerateCve(2024, i)
        fmt.Printf("CVE #%d: %s\n", i, cveId)
    }
}
```

## 实际应用示例

### 处理安全报告

```go
func processSecurityReport(reportText string) {
    fmt.Println("=== 安全报告分析 ===")
    
    // 1. 检查是否包含 CVE
    if !cve.IsContainsCve(reportText) {
        fmt.Println("报告中未发现 CVE")
        return
    }
    
    // 2. 提取所有 CVE
    allCves := cve.ExtractCve(reportText)
    fmt.Printf("发现 %d 个 CVE: %v\n", len(allCves), allCves)
    
    // 3. 去重和排序
    uniqueCves := cve.RemoveDuplicateCves(allCves)
    sortedCves := cve.SortCves(uniqueCves)
    fmt.Printf("去重排序后: %v\n", sortedCves)
    
    // 4. 按年份分组
    grouped := cve.GroupByYear(sortedCves)
    fmt.Println("按年份分组:")
    for year, cves := range grouped {
        fmt.Printf("  %s年: %d个 - %v\n", year, len(cves), cves)
    }
    
    // 5. 分析最近的漏洞
    recentCves := cve.GetRecentCves(sortedCves, 2)
    fmt.Printf("最近2年的漏洞: %v\n", recentCves)
}

// 使用示例
reportText := `
安全公告 2024-001
系统受到以下漏洞影响：
- CVE-2021-44228 (Log4Shell)
- CVE-2022-12345 (自定义组件)
- cve-2021-44228 (重复)
- CVE-2023-1234 (第三方库)
建议立即更新。
`

processSecurityReport(reportText)
```

## 最佳实践

### 1. 错误处理

```go
func safeProcessing(input string) {
    // CVE Utils 的函数对无效输入有良好的处理
    
    // 无效输入返回空字符串
    seq := cve.ExtractCveSeq("invalid")
    if seq == "" {
        fmt.Println("无法提取序列号")
    }
    
    // 无效输入返回 0
    year := cve.ExtractCveYearAsInt("invalid")
    if year == 0 {
        fmt.Println("无法提取年份")
    }
    
    // 验证后再处理
    if cve.ValidateCve(input) {
        // 安全处理
        year, seq := cve.Split(input)
        fmt.Printf("年份: %s, 序列号: %s\n", year, seq)
    } else {
        fmt.Printf("无效的 CVE: %s\n", input)
    }
}
```

### 2. 性能优化

```go
func efficientProcessing(largeCveList []string) {
    // 对于大量数据，建议批量处理
    
    // 一次性去重和排序
    unique := cve.RemoveDuplicateCves(largeCveList)
    sorted := cve.SortCves(unique)
    
    // 避免重复调用格式化函数
    // 好的做法：
    formatted := make([]string, len(largeCveList))
    for i, cveId := range largeCveList {
        formatted[i] = cve.Format(cveId)
    }
    
    // 避免的做法：
    // for _, cveId := range largeCveList {
    //     cve.Format(cveId) // 每次都调用
    // }
}
```

### 3. 数据验证

```go
func validateInput(userInput string) error {
    // 先检查基本格式
    if !cve.IsCve(userInput) {
        return fmt.Errorf("无效的 CVE 格式: %s", userInput)
    }
    
    // 再进行全面验证
    if !cve.ValidateCve(userInput) {
        return fmt.Errorf("CVE 验证失败: %s", userInput)
    }
    
    return nil
}
```

## 下一步

现在您已经掌握了 CVE Utils 的基本使用方法，可以：

1. 查看 [API 文档](/api/) 了解所有函数的详细信息
2. 浏览 [使用示例](/examples/) 学习更多实际应用场景
3. 参考具体的 API 分类文档：
   - [格式化与验证](/api/format-validate)
   - [提取方法](/api/extract)
   - [比较与排序](/api/compare-sort)
   - [过滤与分组](/api/filter-group)
   - [生成与构造](/api/generate)
