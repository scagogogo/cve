# 提取方法

这一类函数用于从文本中提取 CVE 信息，或从 CVE 编号中提取特定部分（如年份、序列号）。

## ExtractCve

从字符串中提取所有 CVE 编号。

### 函数签名

```go
func ExtractCve(text string) []string
```

### 参数

- `text` (string): 要从中提取 CVE 的文本

### 返回值

- `[]string`: 提取的 CVE 编号列表，已格式化为标准格式

### 功能描述

`ExtractCve` 函数从任意文本中提取所有符合 CVE 格式的编号：
- 使用正则表达式匹配 CVE 模式
- 自动格式化为标准大写格式
- 保持在文本中出现的顺序
- 不去重（如果需要去重，请使用 `RemoveDuplicateCves`）

### 使用示例

```go
func main() {
    // 基本使用
    text := "系统受到CVE-2021-44228和cve-2022-12345的影响"
    cves := cve.ExtractCve(text)
    fmt.Printf("提取结果: %v\n", cves)
    // 输出: [CVE-2021-44228 CVE-2022-12345]
    
    // 复杂文本示例
    complexText := `
    安全公告 2024-001
    
    本次更新修复了以下漏洞：
    1. CVE-2021-44228 - Log4Shell 漏洞
    2. cve-2022-12345 - 自定义组件漏洞
    3. CVE-2023-1234 - 第三方库漏洞
    
    另外还包括：CVE-2023-5678 和 CVE-2024-9999
    
    请立即更新到最新版本。
    `
    
    extracted := cve.ExtractCve(complexText)
    fmt.Printf("从复杂文本中提取 (%d 个): %v\n", len(extracted), extracted)
    
    // 处理空文本
    empty := cve.ExtractCve("")
    fmt.Printf("空文本提取结果: %v (长度: %d)\n", empty, len(empty))
    
    // 处理无 CVE 的文本
    noCve := cve.ExtractCve("这个文本不包含任何CVE编号")
    fmt.Printf("无CVE文本提取结果: %v (长度: %d)\n", noCve, len(noCve))
}
```

### 使用场景

- 从安全公告中提取所有相关 CVE
- 分析漏洞报告
- 处理邮件或文档中的 CVE 信息
- 日志分析和数据挖掘

---

## ExtractFirstCve

从字符串中提取第一个 CVE 编号。

### 函数签名

```go
func ExtractFirstCve(text string) string
```

### 参数

- `text` (string): 要从中提取 CVE 的文本

### 返回值

- `string`: 第一个 CVE 编号，如果没有找到则返回空字符串

### 功能描述

`ExtractFirstCve` 函数提取文本中出现的第一个 CVE 编号：
- 只返回第一个匹配的 CVE
- 自动格式化为标准格式
- 性能优于 `ExtractCve` 后取第一个元素

### 使用示例

```go
func main() {
    testCases := []string{
        "主要漏洞是CVE-2021-44228，还有CVE-2022-12345",
        "cve-2022-12345在前面",
        "没有CVE的文本",
        "",
        "CVE-2023-1111, CVE-2023-2222, CVE-2023-3333",
    }
    
    for _, text := range testCases {
        first := cve.ExtractFirstCve(text)
        fmt.Printf("文本: '%s'\n", text)
        fmt.Printf("第一个CVE: '%s'\n\n", first)
    }
}
```

### 使用场景

- 获取主要或最重要的 CVE
- 快速检查文本中的第一个 CVE
- 性能敏感的场景（只需要第一个结果）

---

## ExtractLastCve

从字符串中提取最后一个 CVE 编号。

### 函数签名

```go
func ExtractLastCve(text string) string
```

### 参数

- `text` (string): 要从中提取 CVE 的文本

### 返回值

- `string`: 最后一个 CVE 编号，如果没有找到则返回空字符串

### 功能描述

`ExtractLastCve` 函数提取文本中出现的最后一个 CVE 编号：
- 返回最后一个匹配的 CVE
- 自动格式化为标准格式
- 内部使用 `ExtractCve` 然后取最后一个元素

### 使用示例

```go
func main() {
    text := "漏洞包括CVE-2021-1111、CVE-2022-2222和CVE-2023-3333"
    
    first := cve.ExtractFirstCve(text)
    last := cve.ExtractLastCve(text)
    all := cve.ExtractCve(text)
    
    fmt.Printf("文本: %s\n", text)
    fmt.Printf("第一个: %s\n", first)
    fmt.Printf("最后一个: %s\n", last)
    fmt.Printf("全部: %v\n", all)
}
```

### 使用场景

- 获取最新提及的 CVE
- 处理按时间顺序排列的 CVE 列表
- 获取补充或更新的 CVE 信息

---

## Split

将 CVE 分割成年份和序列号两部分。

### 函数签名

```go
func Split(cve string) (year string, seq string)
```

### 参数

- `cve` (string): 要分割的 CVE 编号

### 返回值

- `year` (string): CVE 的年份部分
- `seq` (string): CVE 的序列号部分

### 功能描述

`Split` 函数将 CVE 编号拆分为年份和序列号：
- 自动格式化输入的 CVE
- 按 `-` 分割字符串
- 对于无效格式，返回空字符串

### 使用示例

```go
func main() {
    testCases := []string{
        "CVE-2022-12345",
        " cve-2021-44228 ",
        "CVE-2023-1",
        "invalid-format",
        "",
    }
    
    for _, cveId := range testCases {
        year, seq := cve.Split(cveId)
        fmt.Printf("CVE: %-20s -> 年份: %-6s 序列号: %-8s\n", 
            cveId, year, seq)
    }
}
```

### 使用场景

- 需要单独处理年份或序列号
- 数据分析和统计
- 自定义排序逻辑
- 数据库存储（分别存储年份和序列号）

---

## ExtractCveYear

从 CVE 中提取年份（字符串格式）。

### 函数签名

```go
func ExtractCveYear(cve string) string
```

### 参数

- `cve` (string): 要提取年份的 CVE 编号

### 返回值

- `string`: CVE 的年份部分

### 功能描述

`ExtractCveYear` 函数提取 CVE 编号中的年份部分：
- 内部调用 `Split` 函数
- 返回字符串格式的年份
- 对于无效 CVE，返回空字符串

### 使用示例

```go
func main() {
    cveList := []string{
        "CVE-2020-1111",
        "CVE-2021-2222", 
        "CVE-2022-3333",
        "CVE-2023-4444",
    }
    
    fmt.Println("CVE 年份提取:")
    for _, cveId := range cveList {
        year := cve.ExtractCveYear(cveId)
        fmt.Printf("%-15s -> %s\n", cveId, year)
    }
}
```

---

## ExtractCveYearAsInt

从 CVE 中提取年份（整数格式）。

### 函数签名

```go
func ExtractCveYearAsInt(cve string) int
```

### 参数

- `cve` (string): 要提取年份的 CVE 编号

### 返回值

- `int`: CVE 的年份（整数类型），无效输入返回 0

### 功能描述

`ExtractCveYearAsInt` 函数提取年份并转换为整数：
- 便于进行数值计算和比较
- 对于无效输入或转换失败，返回 0
- 常用于年份比较和统计

### 使用示例

```go
func main() {
    cveList := []string{
        "CVE-2020-1111",
        "CVE-2021-2222", 
        "CVE-2022-3333",
        "invalid-cve",
    }
    
    currentYear := time.Now().Year()
    
    for _, cveId := range cveList {
        year := cve.ExtractCveYearAsInt(cveId)
        if year > 0 {
            age := currentYear - year
            fmt.Printf("%-15s -> %d年 (距今%d年)\n", cveId, year, age)
        } else {
            fmt.Printf("%-15s -> 无效CVE\n", cveId)
        }
    }
}
```

---

## ExtractCveSeq

从 CVE 中提取序列号（字符串格式）。

### 函数签名

```go
func ExtractCveSeq(cve string) string
```

### 参数

- `cve` (string): 要提取序列号的 CVE 编号

### 返回值

- `string`: CVE 的序列号部分，如果不是有效 CVE 则返回空字符串

### 功能描述

`ExtractCveSeq` 函数提取 CVE 编号中的序列号部分：
- 先验证是否为有效 CVE 格式
- 返回字符串格式的序列号
- 保持原始的序列号格式（包括前导零）

### 使用示例

```go
func main() {
    testCases := []string{
        "CVE-2022-12345",
        "CVE-2022-0001",    // 有前导零
        "CVE-2022-1",       // 短序列号
        "invalid-format",   // 无效格式
    }
    
    for _, cveId := range testCases {
        seq := cve.ExtractCveSeq(cveId)
        fmt.Printf("%-20s -> 序列号: '%s'\n", cveId, seq)
    }
}
```

---

## ExtractCveSeqAsInt

从 CVE 中提取序列号（整数格式）。

### 函数签名

```go
func ExtractCveSeqAsInt(cve string) int
```

### 参数

- `cve` (string): 要提取序列号的 CVE 编号

### 返回值

- `int`: CVE 的序列号（整数类型），无效输入返回 0

### 功能描述

`ExtractCveSeqAsInt` 函数提取序列号并转换为整数：
- 便于进行数值比较和排序
- 自动处理前导零
- 对于无效输入，返回 0

### 使用示例

```go
func main() {
    cveList := []string{
        "CVE-2022-00001",   // 有前导零
        "CVE-2022-12345",
        "CVE-2022-1",
        "CVE-2022-99999",
    }
    
    // 按序列号排序
    sort.Slice(cveList, func(i, j int) bool {
        seqA := cve.ExtractCveSeqAsInt(cveList[i])
        seqB := cve.ExtractCveSeqAsInt(cveList[j])
        return seqA < seqB
    })
    
    fmt.Println("按序列号排序:")
    for _, cveId := range cveList {
        seq := cve.ExtractCveSeqAsInt(cveId)
        fmt.Printf("%-15s -> %d\n", cveId, seq)
    }
}
```

## 实际应用示例

### 1. 文本分析流水线

```go
func analyzeSecurityReport(reportText string) {
    fmt.Println("=== 安全报告分析 ===")
    
    // 检查是否包含 CVE
    if !cve.IsContainsCve(reportText) {
        fmt.Println("报告中未发现 CVE")
        return
    }
    
    // 提取所有 CVE
    allCves := cve.ExtractCve(reportText)
    fmt.Printf("发现 %d 个 CVE: %v\n", len(allCves), allCves)
    
    // 分析第一个和最后一个 CVE
    if len(allCves) > 0 {
        first := cve.ExtractFirstCve(reportText)
        last := cve.ExtractLastCve(reportText)
        fmt.Printf("第一个 CVE: %s\n", first)
        fmt.Printf("最后一个 CVE: %s\n", last)
        
        // 分析年份分布
        yearMap := make(map[string]int)
        for _, cveId := range allCves {
            year := cve.ExtractCveYear(cveId)
            yearMap[year]++
        }
        
        fmt.Println("年份分布:")
        for year, count := range yearMap {
            fmt.Printf("  %s年: %d个\n", year, count)
        }
    }
}
```

### 2. CVE 信息提取器

```go
type CveInfo struct {
    FullId   string
    Year     int
    Sequence int
    YearStr  string
    SeqStr   string
}

func extractCveInfo(cveId string) *CveInfo {
    if !cve.IsCve(cveId) {
        return nil
    }
    
    year, seq := cve.Split(cveId)
    
    return &CveInfo{
        FullId:   cve.Format(cveId),
        Year:     cve.ExtractCveYearAsInt(cveId),
        Sequence: cve.ExtractCveSeqAsInt(cveId),
        YearStr:  year,
        SeqStr:   seq,
    }
}

func main() {
    cveId := "cve-2022-12345"
    info := extractCveInfo(cveId)
    
    if info != nil {
        fmt.Printf("CVE 信息:\n")
        fmt.Printf("  完整编号: %s\n", info.FullId)
        fmt.Printf("  年份: %d (%s)\n", info.Year, info.YearStr)
        fmt.Printf("  序列号: %d (%s)\n", info.Sequence, info.SeqStr)
    }
}
```

### 3. 批量文本处理

```go
func processBatchTexts(texts []string) map[string][]string {
    result := make(map[string][]string)
    
    for i, text := range texts {
        key := fmt.Sprintf("text_%d", i+1)
        
        // 提取所有 CVE
        cves := cve.ExtractCve(text)
        
        // 去重并排序
        unique := cve.RemoveDuplicateCves(cves)
        sorted := cve.SortCves(unique)
        
        result[key] = sorted
    }
    
    return result
}
```

## 性能说明

- `ExtractFirstCve` 比 `ExtractCve` 后取第一个元素更高效
- `ExtractLastCve` 内部使用 `ExtractCve`，性能相当
- 整数提取函数（`*AsInt`）包含类型转换，略慢于字符串版本
- 所有函数都使用预编译的正则表达式，性能优异
- 对于大量文本处理，建议批量操作而不是逐个调用
