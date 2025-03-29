# CVE Utils

[![Go Tests](https://github.com/scagogogo/cve/actions/workflows/go-test.yml/badge.svg)](https://github.com/scagogogo/cve/actions/workflows/go-test.yml)

CVE (Common Vulnerabilities and Exposures) 相关的工具方法集合。这个包提供了一系列用于处理、验证、提取和操作 CVE 标识符的实用函数。

## 功能特性

- CVE 格式验证和标准化
- 从文本中提取 CVE 标识符
- CVE 的年份和序列号提取与比较
- CVE 的排序、过滤和分组
- 生成标准格式的 CVE 标识符
- 去重和验证工具

## 安装

```go
go get github.com/scagogogo/cve
```

## 使用示例

### 基本使用

```go
import "github.com/scagogogo/cve"

// 格式化 CVE
formattedCve := cve.Format("cve-2022-12345")  // 返回 "CVE-2022-12345"

// 验证是否为合法 CVE
isValid := cve.ValidateCve("CVE-2022-12345")  // 验证 CVE 是否符合规范

// 从文本中提取 CVE
text := "系统中发现了多个漏洞：CVE-2021-44228 和 CVE-2022-12345"
cveList := cve.ExtractCve(text)  // 返回 ["CVE-2021-44228", "CVE-2022-12345"]
```

### 高级功能

```go
// 排序 CVE 列表
cveList := []string{"CVE-2022-12345", "CVE-2021-44228", "CVE-2022-10000"}
sortedList := cve.SortCves(cveList)  // 按年份和序列号排序

// 按年份过滤 CVE
recentCves := cve.GetRecentCves(cveList, 2)  // 获取最近两年的 CVE

// 去除重复的 CVE
uniqueCves := cve.RemoveDuplicateCves(cveList)
```

## API 详细文档

### 格式化与验证

#### Format

将 CVE 编号转换为标准大写格式并移除前后空格。

```go
func Format(cve string) string
```

**参数**：
- `cve` - 要格式化的 CVE 编号

**返回值**：
- 标准化格式的 CVE 编号

**示例**：
```go
formattedCve := cve.Format(" cve-2022-12345 ")  // 返回 "CVE-2022-12345"
```

#### IsCve

判断字符串是否是有效的 CVE 格式。

```go
func IsCve(text string) bool
```

**参数**：
- `text` - 要检查的字符串

**返回值**：
- 如果字符串是有效的 CVE 格式则返回 `true`，否则返回 `false`

**示例**：
```go
isCve := cve.IsCve("CVE-2022-12345")  // 返回 true
isCve = cve.IsCve("包含CVE-2022-12345的文本")  // 返回 false
```

#### IsContainsCve

判断字符串是否包含 CVE。

```go
func IsContainsCve(text string) bool
```

**参数**：
- `text` - 要检查的字符串

**返回值**：
- 如果字符串包含 CVE 则返回 `true`，否则返回 `false`

**示例**：
```go
contains := cve.IsContainsCve("这个漏洞的编号是CVE-2022-12345")  // 返回 true
contains = cve.IsContainsCve("这个文本不包含任何CVE标识符")  // 返回 false
```

#### IsCveYearOk

判断 CVE 的年份是否在合理的时间范围内。

```go
func IsCveYearOk(cve string, cutoff int) bool
```

**参数**：
- `cve` - CVE 编号
- `cutoff` - 允许的年份偏移量

**返回值**：
- 如果年份在合理范围内则返回 `true`，否则返回 `false`

**示例**：
```go
// 假设当前年份是2023年
isYearOk := cve.IsCveYearOk("CVE-2022-12345", 5)  // 返回 true
isYearOk = cve.IsCveYearOk("CVE-2030-12345", 5)  // 返回 false (2030 > 2023+5)
```

#### ValidateCve

全面验证 CVE 编号的合法性。

```go
func ValidateCve(cve string) bool
```

**参数**：
- `cve` - 要验证的 CVE 编号

**返回值**：
- 如果 CVE 编号合法则返回 `true`，否则返回 `false`

**示例**：
```go
isValid := cve.ValidateCve("CVE-2022-12345")  // 正常情况返回 true
isValid = cve.ValidateCve("CVE-1960-12345")  // 返回 false (年份 < 1970)
isValid = cve.ValidateCve("CVE-2099-12345")  // 返回 false (假设当前为2023年，年份超前太多)
```

### 提取方法

#### ExtractCve

从字符串中提取所有 CVE 编号。

```go
func ExtractCve(text string) []string
```

**参数**：
- `text` - 要从中提取 CVE 的文本

**返回值**：
- 提取的 CVE 编号列表

**示例**：
```go
text := "系统受到CVE-2021-44228和cve-2022-12345的影响"
cveList := cve.ExtractCve(text)  // 返回 ["CVE-2021-44228", "CVE-2022-12345"]
```

#### ExtractFirstCve

从字符串中提取第一个 CVE 编号。

```go
func ExtractFirstCve(text string) string
```

**参数**：
- `text` - 要从中提取 CVE 的文本

**返回值**：
- 第一个 CVE 编号，如果没有找到则返回空字符串

**示例**：
```go
text := "系统受到CVE-2021-44228和CVE-2022-12345的影响"
firstCve := cve.ExtractFirstCve(text)  // 返回 "CVE-2021-44228"
```

#### ExtractLastCve

从字符串中提取最后一个 CVE 编号。

```go
func ExtractLastCve(text string) string
```

**参数**：
- `text` - 要从中提取 CVE 的文本

**返回值**：
- 最后一个 CVE 编号，如果没有找到则返回空字符串

**示例**：
```go
text := "系统受到CVE-2021-44228和CVE-2022-12345的影响"
lastCve := cve.ExtractLastCve(text)  // 返回 "CVE-2022-12345"
```

#### Split

将 CVE 分割成年份和编号两部分。

```go
func Split(cve string) (year string, seq string)
```

**参数**：
- `cve` - 要分割的 CVE 编号

**返回值**：
- `year` - CVE 的年份部分
- `seq` - CVE 的序列号部分

**示例**：
```go
year, seq := cve.Split("CVE-2022-12345")  // 返回 year="2022", seq="12345"
```

#### ExtractCveYear

从 CVE 中提取年份。

```go
func ExtractCveYear(cve string) string
```

**参数**：
- `cve` - 要提取年份的 CVE 编号

**返回值**：
- CVE 的年份部分

**示例**：
```go
year := cve.ExtractCveYear("CVE-2022-12345")  // 返回 "2022"
```

#### ExtractCveYearAsInt

从 CVE 中提取年份，并解析为整数类型。

```go
func ExtractCveYearAsInt(cve string) int
```

**参数**：
- `cve` - 要提取年份的 CVE 编号

**返回值**：
- CVE 的年份（整数类型）

**示例**：
```go
year := cve.ExtractCveYearAsInt("CVE-2022-12345")  // 返回 2022
```

#### ExtractCveSeq

从 CVE 中提取序列号。

```go
func ExtractCveSeq(cve string) string
```

**参数**：
- `cve` - 要提取序列号的 CVE 编号

**返回值**：
- CVE 的序列号部分，如果不是有效 CVE 则返回空字符串

**示例**：
```go
seq := cve.ExtractCveSeq("CVE-2022-12345")  // 返回 "12345"
seq = cve.ExtractCveSeq("非CVE格式文本")  // 返回 ""
```

#### ExtractCveSeqAsInt

从 CVE 中提取序列号，并解析为整数类型。

```go
func ExtractCveSeqAsInt(cve string) int
```

**参数**：
- `cve` - 要提取序列号的 CVE 编号

**返回值**：
- CVE 的序列号（整数类型），如果不是有效 CVE 则返回 0

**示例**：
```go
seq := cve.ExtractCveSeqAsInt("CVE-2022-12345")  // 返回 12345
seq = cve.ExtractCveSeqAsInt("非CVE格式文本")  // 返回 0
```

### 比较与排序

#### CompareByYear

根据 CVE 的年份比较大小。

```go
func CompareByYear(cveA, cveB string) int
```

**参数**：
- `cveA` - 第一个 CVE 编号
- `cveB` - 第二个 CVE 编号

**返回值**：
- 负数：cveA 年份 < cveB 年份
- 零：cveA 年份 = cveB 年份
- 正数：cveA 年份 > cveB 年份

**示例**：
```go
result := cve.CompareByYear("CVE-2020-1111", "CVE-2022-2222")  // 返回 -2
result = cve.CompareByYear("CVE-2022-1111", "CVE-2022-2222")  // 返回 0
result = cve.CompareByYear("CVE-2022-1111", "CVE-2020-2222")  // 返回 2
```

#### SubByYear

计算两个 CVE 的年份差值。

```go
func SubByYear(cveA, cveB string) int
```

**参数**：
- `cveA` - 第一个 CVE 编号
- `cveB` - 第二个 CVE 编号

**返回值**：
- cveA 年份 - cveB 年份的差值

**示例**：
```go
diff := cve.SubByYear("CVE-2020-1111", "CVE-2022-2222")  // 返回 -2
diff = cve.SubByYear("CVE-2022-1111", "CVE-2020-2222")  // 返回 2
```

#### CompareCves

全面比较两个 CVE 编号的大小。

```go
func CompareCves(cveA, cveB string) int
```

**参数**：
- `cveA` - 第一个 CVE 编号
- `cveB` - 第二个 CVE 编号

**返回值**：
- -1：cveA < cveB
- 0：cveA = cveB
- 1：cveA > cveB

**示例**：
```go
result := cve.CompareCves("CVE-2020-1111", "CVE-2022-2222")  // 返回 -1 (不同年份)
result = cve.CompareCves("CVE-2022-1111", "CVE-2022-2222")  // 返回 -1 (相同年份，不同序列号)
result = cve.CompareCves("CVE-2022-2222", "CVE-2022-2222")  // 返回 0 (完全相同)
```

#### SortCves

对 CVE 切片进行排序（按年份和序列号）。

```go
func SortCves(cveSlice []string) []string
```

**参数**：
- `cveSlice` - 要排序的 CVE 编号列表

**返回值**：
- 排序后的 CVE 编号列表

**示例**：
```go
cveList := []string{"CVE-2022-2222", "cve-2020-1111", "CVE-2022-1111"}
sortedList := cve.SortCves(cveList)  // 返回 ["CVE-2020-1111", "CVE-2022-1111", "CVE-2022-2222"]
```

### 过滤与分组

#### FilterCvesByYear

筛选特定年份的 CVE。

```go
func FilterCvesByYear(cveSlice []string, year int) []string
```

**参数**：
- `cveSlice` - CVE 编号列表
- `year` - 要筛选的年份

**返回值**：
- 指定年份的 CVE 编号列表

**示例**：
```go
cveList := []string{"CVE-2021-1111", "CVE-2022-2222", "CVE-2021-3333"}
cves2021 := cve.FilterCvesByYear(cveList, 2021)  // 返回 ["CVE-2021-1111", "CVE-2021-3333"]
```

#### FilterCvesByYearRange

筛选指定年份范围内的 CVE。

```go
func FilterCvesByYearRange(cveSlice []string, startYear, endYear int) []string
```

**参数**：
- `cveSlice` - CVE 编号列表
- `startYear` - 开始年份（含）
- `endYear` - 结束年份（含）

**返回值**：
- 指定年份范围内的 CVE 编号列表

**示例**：
```go
cveList := []string{"CVE-2020-1111", "CVE-2021-2222", "CVE-2022-3333"}
rangeCves := cve.FilterCvesByYearRange(cveList, 2021, 2022)  // 返回 ["CVE-2021-2222", "CVE-2022-3333"]
```

#### GetRecentCves

获取最近几年的 CVE。

```go
func GetRecentCves(cveSlice []string, years int) []string
```

**参数**：
- `cveSlice` - CVE 编号列表
- `years` - 最近几年（从当前年份往前计算）

**返回值**：
- 最近几年的 CVE 编号列表

**示例**：
```go
// 假设当前年份是2023年
cveList := []string{"CVE-2020-1111", "CVE-2021-2222", "CVE-2022-3333", "CVE-2023-4444"}
recentCves := cve.GetRecentCves(cveList, 2)  // 返回 ["CVE-2022-3333", "CVE-2023-4444"]
```

#### GroupByYear

按年份对 CVE 进行分组。

```go
func GroupByYear(cveSlice []string) map[string][]string
```

**参数**：
- `cveSlice` - 要分组的 CVE 编号列表

**返回值**：
- 按年份分组的 CVE 编号映射表

**示例**：
```go
cveList := []string{"CVE-2021-1111", "CVE-2022-2222", "CVE-2021-3333"}
groupedCves := cve.GroupByYear(cveList)
// 返回:
// {
//   "2021": ["CVE-2021-1111", "CVE-2021-3333"],
//   "2022": ["CVE-2022-2222"]
// }

// 遍历示例
for year, cves := range groupedCves {
    fmt.Printf("Year %s: %v\n", year, cves)
}
```

#### RemoveDuplicateCves

移除重复的 CVE 编号。

```go
func RemoveDuplicateCves(cveSlice []string) []string
```

**参数**：
- `cveSlice` - 可能包含重复项的 CVE 编号列表

**返回值**：
- 去重后的 CVE 编号列表

**示例**：
```go
cveList := []string{"CVE-2022-1111", "cve-2022-1111", "CVE-2022-2222"}
uniqueCves := cve.RemoveDuplicateCves(cveList)  // 返回 ["CVE-2022-1111", "CVE-2022-2222"]
```

### 生成与构造

#### GenerateCve

根据年份和序列号生成标准格式的 CVE 编号。

```go
func GenerateCve(year int, seq int) string
```

**参数**：
- `year` - CVE 年份
- `seq` - CVE 序列号

**返回值**：
- 生成的标准格式 CVE 编号

**示例**：
```go
cveId := cve.GenerateCve(2022, 12345)  // 返回 "CVE-2022-12345"
```

## 常见使用场景

### 漏洞报告分析

```go
// 从文本中提取所有 CVE 并按年份分组
text := `安全公告：系统受到多个漏洞影响，包括 CVE-2021-44228、
CVE-2021-45046、CVE-2022-1234 和 CVE-2022-5678。
建议尽快更新补丁。`

// 提取所有 CVE
cves := cve.ExtractCve(text)

// 按年份分组
groupedCves := cve.GroupByYear(cves)

// 输出分组结果
for year, yearCves := range groupedCves {
    fmt.Printf("%s年的CVE：%v\n", year, yearCves)
}
```

### 漏洞库管理

```go
// 导入 CVE 并进行去重和排序
importedCves := []string{
    "CVE-2022-1111", "cve-2022-1111", "CVE-2021-2222",
    "CVE-2020-3333", "CVE-2022-4444", "CVE-2022-1111",
}

// 去重
uniqueCves := cve.RemoveDuplicateCves(importedCves)

// 排序
sortedCves := cve.SortCves(uniqueCves)

// 获取最近两年的 CVE
recentCves := cve.GetRecentCves(sortedCves, 2)

fmt.Println("最近两年的 CVE：", recentCves)
```

### CVE 验证和处理

```go
// 验证用户输入的 CVE 格式
userInput := " cve-2022-12345 "

if cve.IsCve(userInput) {
    formattedCve := cve.Format(userInput)
    fmt.Println("有效的 CVE：", formattedCve)
    
    year := cve.ExtractCveYear(formattedCve)
    seq := cve.ExtractCveSeq(formattedCve)
    fmt.Printf("年份: %s, 序列号: %s\n", year, seq)
    
    if cve.ValidateCve(formattedCve) {
        fmt.Println("CVE 验证通过")
    } else {
        fmt.Println("CVE 格式正确但验证失败（可能年份过早或过晚）")
    }
} else {
    fmt.Println("无效的 CVE 格式")
}
```

## 参考资料

- [CVE 官方网站](https://cve.mitre.org/)
- [NIST 国家漏洞数据库](https://nvd.nist.gov/)

## 许可证

[许可证信息] 
