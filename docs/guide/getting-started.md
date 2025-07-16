# 快速开始

欢迎使用 CVE Utils！这个指南将帮助您快速上手使用这个强大的 CVE 处理工具库。

## 安装

### 使用 go get 安装

```bash
go get github.com/scagogogo/cve
```

### 验证安装

创建一个简单的测试文件来验证安装是否成功：

```go
// test.go
package main

import (
    "fmt"
    "github.com/scagogogo/cve"
)

func main() {
    // 测试基本功能
    result := cve.Format("cve-2022-12345")
    fmt.Println("格式化结果:", result)
    
    if result == "CVE-2022-12345" {
        fmt.Println("✅ CVE Utils 安装成功！")
    } else {
        fmt.Println("❌ 安装可能有问题")
    }
}
```

运行测试：

```bash
go run test.go
```

## 基本概念

### CVE 格式

CVE (Common Vulnerabilities and Exposures) 编号遵循特定的格式：

```
CVE-YYYY-NNNN
```

- `CVE`: 固定前缀
- `YYYY`: 4位年份
- `NNNN`: 序列号（至少4位数字）

例如：`CVE-2022-12345`、`CVE-2021-44228`

### 主要功能分类

CVE Utils 提供的功能可以分为以下几类：

1. **格式化与验证**: 标准化和验证 CVE 格式
2. **提取方法**: 从文本中提取 CVE 信息
3. **比较与排序**: 对 CVE 进行比较和排序
4. **过滤与分组**: 按条件过滤和分组 CVE
5. **生成与构造**: 生成新的 CVE 编号

## 第一个示例

让我们从一个简单的示例开始：

```go
package main

import (
    "fmt"
    "github.com/scagogogo/cve"
)

func main() {
    // 1. 格式化 CVE
    input := " cve-2022-12345 "
    formatted := cve.Format(input)
    fmt.Printf("原始输入: '%s'\n", input)
    fmt.Printf("格式化后: '%s'\n", formatted)
    
    // 2. 验证 CVE
    isValid := cve.ValidateCve(formatted)
    fmt.Printf("是否有效: %t\n", isValid)
    
    // 3. 从文本中提取 CVE
    text := "系统受到多个漏洞影响：CVE-2021-44228、CVE-2022-12345 和 cve-2023-1234"
    cves := cve.ExtractCve(text)
    fmt.Printf("提取的 CVE: %v\n", cves)
    
    // 4. 排序 CVE
    sorted := cve.SortCves(cves)
    fmt.Printf("排序后: %v\n", sorted)
}
```

运行结果：

```
原始输入: ' cve-2022-12345 '
格式化后: 'CVE-2022-12345'
是否有效: true
提取的 CVE: [CVE-2021-44228 CVE-2022-12345 CVE-2023-1234]
排序后: [CVE-2021-44228 CVE-2022-12345 CVE-2023-1234]
```

## 常用操作示例

### 处理用户输入

```go
func processUserInput(input string) {
    // 检查输入是否包含 CVE
    if !cve.IsContainsCve(input) {
        fmt.Println("输入中没有找到 CVE")
        return
    }
    
    // 提取第一个 CVE
    firstCve := cve.ExtractFirstCve(input)
    fmt.Printf("第一个 CVE: %s\n", firstCve)
    
    // 验证有效性
    if cve.ValidateCve(firstCve) {
        fmt.Println("✅ CVE 格式有效")
        
        // 提取年份和序列号
        year, seq := cve.Split(firstCve)
        fmt.Printf("年份: %s, 序列号: %s\n", year, seq)
    } else {
        fmt.Println("❌ CVE 格式无效")
    }
}

// 使用示例
processUserInput("漏洞编号：CVE-2022-12345")
```

### 批量处理 CVE

```go
func processCveList(cveList []string) {
    fmt.Printf("原始列表 (%d 个): %v\n", len(cveList), cveList)
    
    // 去重
    unique := cve.RemoveDuplicateCves(cveList)
    fmt.Printf("去重后 (%d 个): %v\n", len(unique), unique)
    
    // 排序
    sorted := cve.SortCves(unique)
    fmt.Printf("排序后: %v\n", sorted)
    
    // 按年份分组
    grouped := cve.GroupByYear(sorted)
    fmt.Println("按年份分组:")
    for year, cves := range grouped {
        fmt.Printf("  %s: %v\n", year, cves)
    }
    
    // 获取最近2年的 CVE
    recent := cve.GetRecentCves(sorted, 2)
    fmt.Printf("最近2年: %v\n", recent)
}

// 使用示例
cveList := []string{
    "CVE-2022-1111",
    "cve-2022-1111", // 重复项（大小写不同）
    "CVE-2021-2222",
    "CVE-2023-3333",
    "CVE-2022-4444",
}
processCveList(cveList)
```

## 错误处理

CVE Utils 的大部分函数都有良好的错误处理机制：

```go
func safeProcessing() {
    // 对于无效输入，函数会返回安全的默认值
    
    // 无效 CVE 返回空字符串
    seq := cve.ExtractCveSeq("invalid-input")
    fmt.Printf("无效输入的序列号: '%s'\n", seq) // 输出: ''
    
    // 无效 CVE 返回 0
    year := cve.ExtractCveYearAsInt("invalid-input")
    fmt.Printf("无效输入的年份: %d\n", year) // 输出: 0
    
    // 空文本返回空切片
    cves := cve.ExtractCve("")
    fmt.Printf("空文本提取结果: %v\n", cves) // 输出: []
}
```

## 性能考虑

CVE Utils 针对性能进行了优化：

```go
func performanceExample() {
    // 对于大量数据，建议批量处理
    largeCveList := make([]string, 10000)
    for i := 0; i < 10000; i++ {
        largeCveList[i] = fmt.Sprintf("CVE-2022-%d", i+1)
    }
    
    start := time.Now()
    
    // 批量去重和排序
    unique := cve.RemoveDuplicateCves(largeCveList)
    sorted := cve.SortCves(unique)
    
    duration := time.Since(start)
    fmt.Printf("处理 %d 个 CVE 耗时: %v\n", len(largeCveList), duration)
    fmt.Printf("结果数量: %d\n", len(sorted))
}
```

## 下一步

现在您已经了解了 CVE Utils 的基本用法，可以：

1. 查看 [API 文档](/api/) 了解所有可用函数
2. 浏览 [使用示例](/examples/) 学习更多实际应用场景
3. 查看 [基本使用指南](/guide/basic-usage) 了解更多细节

如果遇到问题，请查看 [GitHub Issues](https://github.com/scagogogo/cve/issues) 或提交新的问题。
