# CVE Utils

[![Go Tests](https://github.com/scagogogo/cve/actions/workflows/go-test.yml/badge.svg)](https://github.com/scagogogo/cve/actions/workflows/go-test.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/scagogogo/cve.svg)](https://pkg.go.dev/github.com/scagogogo/cve)
[![Go Report Card](https://goreportcard.com/badge/github.com/scagogogo/cve)](https://goreportcard.com/report/github.com/scagogogo/cve)
[![License](https://img.shields.io/github/license/scagogogo/cve)](https://github.com/scagogogo/cve/blob/main/LICENSE)
[![Version](https://img.shields.io/badge/version-v0.0.1-blue)](https://github.com/scagogogo/cve/releases)

CVE (Common Vulnerabilities and Exposures) 相关的工具方法集合。这个包提供了一系列用于处理、验证、提取和操作 CVE 标识符的实用函数。

## 📑 目录

- [功能特性](#-功能特性)
- [安装](#-安装)
- [快速开始](#-快速开始)
- [API 参考文档](#-api-参考文档)
  - [格式化与验证](#格式化与验证)
  - [提取方法](#提取方法)
  - [比较与排序](#比较与排序)
  - [过滤与分组](#过滤与分组)
  - [生成与构造](#生成与构造)
- [使用场景示例](#-使用场景示例)
- [项目结构](#-项目结构)
- [参考资料](#-参考资料)
- [许可证](#-许可证)

## 🚀 功能特性

- ✅ CVE 格式验证和标准化
- ✅ 从文本中提取 CVE 标识符
- ✅ CVE 的年份和序列号提取与比较
- ✅ CVE 的排序、过滤和分组
- ✅ 生成标准格式的 CVE 标识符
- ✅ 去重和验证工具

## 📦 安装

```bash
go get github.com/scagogogo/cve
```

## 🚦 快速开始

### 基本使用

```go
package main

import (
	"fmt"
	"github.com/scagogogo/cve"
)

func main() {
	// 格式化 CVE
	formattedCve := cve.Format("cve-2022-12345")
	fmt.Println("格式化后:", formattedCve) // 输出: "CVE-2022-12345"

	// 验证是否为合法 CVE
	isValid := cve.ValidateCve("CVE-2022-12345")
	fmt.Println("是否有效:", isValid) // 输出: true

	// 从文本中提取 CVE
	text := "系统中发现了多个漏洞：CVE-2021-44228 和 CVE-2022-12345"
	cveList := cve.ExtractCve(text)
	fmt.Println("提取的CVE:", cveList) // 输出: ["CVE-2021-44228", "CVE-2022-12345"]
}
```

### 高级功能

```go
package main

import (
	"fmt"
	"github.com/scagogogo/cve"
)

func main() {
	// 排序 CVE 列表
	cveList := []string{"CVE-2022-12345", "CVE-2021-44228", "CVE-2022-10000"}
	sortedList := cve.SortCves(cveList)
	fmt.Println("排序后:", sortedList)
	
	// 按年份过滤 CVE
	recentCves := cve.GetRecentCves(cveList, 2)
	fmt.Println("最近两年的CVE:", recentCves)
	
	// 去除重复的 CVE
	duplicatedList := []string{"CVE-2022-12345", "cve-2022-12345", "CVE-2021-44228"}
	uniqueCves := cve.RemoveDuplicateCves(duplicatedList)
	fmt.Println("去重后:", uniqueCves)
}
```

## 📚 API 参考文档

### 格式化与验证

<details>
<summary><b>Format</b> - 将 CVE 编号转换为标准大写格式并移除前后空格</summary>

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
formattedCve = cve.Format("CVE-2021-44228")     // 返回 "CVE-2021-44228"
```
</details>

<details>
<summary><b>IsCve</b> - 判断字符串是否是有效的 CVE 格式</summary>

```go
func IsCve(text string) bool
```

**参数**：
- `text` - 要检查的字符串

**返回值**：
- 如果字符串是有效的 CVE 格式则返回 `true`，否则返回 `false`

**示例**：
```go
isCve := cve.IsCve("CVE-2022-12345")           // 返回 true
isCve = cve.IsCve(" CVE-2022-12345 ")          // 返回 true (忽略前后空格)
isCve = cve.IsCve("包含CVE-2022-12345的文本")    // 返回 false
isCve = cve.IsCve("cve2022-12345")             // 返回 false (格式不正确)
```
</details>

<details>
<summary><b>IsContainsCve</b> - 判断字符串是否包含 CVE</summary>

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
contains = cve.IsContainsCve("修复了cve-2021-44228漏洞")        // 返回 true (不区分大小写)
contains = cve.IsContainsCve("这个文本不包含任何CVE标识符")        // 返回 false
```
</details>

<details>
<summary><b>IsCveYearOk</b> - 判断 CVE 的年份是否在合理的时间范围内</summary>

```go
func IsCveYearOk(cve string) bool
```

**参数**：
- `cve` - CVE 编号

**返回值**：
- 如果年份在1999年之后且不超过当前年份则返回 `true`，否则返回 `false`

**示例**：
```go
// 假设当前年份是2023年
isYearOk := cve.IsCveYearOk("CVE-2022-12345")  // 返回 true
isYearOk = cve.IsCveYearOk("CVE-2023-12345")   // 返回 true (当前年份)
isYearOk = cve.IsCveYearOk("CVE-2030-12345")   // 返回 false (2030 > 2023)
isYearOk = cve.IsCveYearOk("CVE-1998-12345")   // 返回 false (1998 < 1999)
```
</details>

<details>
<summary><b>IsCveYearOkWithCutoff</b> - 判断 CVE 的年份是否在合理的时间范围内（可设置偏移量）</summary>

```go
func IsCveYearOkWithCutoff(cve string, cutoff int) bool
```

**参数**：
- `cve` - CVE 编号
- `cutoff` - 允许的年份偏移量

**返回值**：
- 如果年份在合理范围内则返回 `true`，否则返回 `false`

**示例**：
```go
// 假设当前年份是2023年
isYearOk := cve.IsCveYearOkWithCutoff("CVE-2022-12345", 0)  // 返回 true
isYearOk = cve.IsCveYearOkWithCutoff("CVE-2025-12345", 2)   // 返回 true (2025 <= 2023+2)
isYearOk = cve.IsCveYearOkWithCutoff("CVE-2030-12345", 5)   // 返回 false (2030 > 2023+5)
isYearOk = cve.IsCveYearOkWithCutoff("CVE-1998-12345", 0)   // 返回 false (1998 < 1999)
```
</details>

<details>
<summary><b>ValidateCve</b> - 全面验证 CVE 编号的合法性</summary>

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
isValid = cve.ValidateCve("CVE-1998-12345")   // 返回 false (年份 < 1999)
isValid = cve.ValidateCve("CVE-2099-12345")   // 返回 false (假设当前为2023年，年份超前太多)
isValid = cve.ValidateCve("CVE-2022-0")       // 返回 false (序列号格式不正确)
isValid = cve.ValidateCve("CVE2022-12345")    // 返回 false (缺少连字符)
```
</details>

### 提取方法

<details>
<summary><b>ExtractCve</b> - 从字符串中提取所有 CVE 编号</summary>

```go
func ExtractCve(text string) []string
```

**参数**：
- `text` - 要从中提取 CVE 的文本

**返回值**：
- 提取的 CVE 编号列表，按标准格式返回

**示例**：
```go
text := "系统受到CVE-2021-44228和cve-2022-12345的影响"
cveList := cve.ExtractCve(text)  // 返回 ["CVE-2021-44228", "CVE-2022-12345"]

text = "没有包含任何CVE的文本"
cveList = cve.ExtractCve(text)   // 返回 [] (空切片)
```
</details>

<details>
<summary><b>ExtractFirstCve</b> - 从字符串中提取第一个 CVE 编号</summary>

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

text = "没有包含任何CVE的文本"
firstCve = cve.ExtractFirstCve(text)   // 返回 "" (空字符串)
```
</details>

<details>
<summary><b>ExtractLastCve</b> - 从字符串中提取最后一个 CVE 编号</summary>

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

text = "没有包含任何CVE的文本"
lastCve = cve.ExtractLastCve(text)   // 返回 "" (空字符串)
```
</details>

<details>
<summary><b>Split</b> - 将 CVE 分割成年份和编号两部分</summary>

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
year, seq = cve.Split("cve-2021-44228")   // 返回 year="2021", seq="44228"
year, seq = cve.Split("不是CVE格式")        // 返回 year="", seq=""
```
</details>

<details>
<summary><b>ExtractCveYear</b> - 从 CVE 中提取年份</summary>

```go
func ExtractCveYear(cve string) string
```

**参数**：
- `cve` - 要提取年份的 CVE 编号

**返回值**：
- CVE 的年份部分，如果不是有效 CVE 则返回空字符串

**示例**：
```go
year := cve.ExtractCveYear("CVE-2022-12345")  // 返回 "2022"
year = cve.ExtractCveYear("cve-2021-44228")   // 返回 "2021"
year = cve.ExtractCveYear("不是CVE格式")        // 返回 ""
```
</details>

<details>
<summary><b>ExtractCveYearAsInt</b> - 从 CVE 中提取年份并转换为整数</summary>

```go
func ExtractCveYearAsInt(cve string) int
```

**参数**：
- `cve` - 要提取年份的 CVE 编号

**返回值**：
- CVE 的年份（整数类型），如果不是有效 CVE 则返回 0

**示例**：
```go
year := cve.ExtractCveYearAsInt("CVE-2022-12345")  // 返回 2022
year = cve.ExtractCveYearAsInt("cve-2021-44228")   // 返回 2021
year = cve.ExtractCveYearAsInt("不是CVE格式")        // 返回 0
```
</details>

<details>
<summary><b>ExtractCveSeq</b> - 从 CVE 中提取序列号</summary>

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
seq = cve.ExtractCveSeq("cve-2021-44228")   // 返回 "44228"
seq = cve.ExtractCveSeq("不是CVE格式")        // 返回 ""
```
</details>

<details>
<summary><b>ExtractCveSeqAsInt</b> - 从 CVE 中提取序列号并转换为整数</summary>

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
seq = cve.ExtractCveSeqAsInt("cve-2021-44228")   // 返回 44228
seq = cve.ExtractCveSeqAsInt("不是CVE格式")        // 返回 0
```
</details>

### 比较与排序

<details>
<summary><b>CompareByYear</b> - 根据 CVE 的年份比较大小</summary>

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
result = cve.CompareByYear("CVE-2022-1111", "CVE-2022-2222")   // 返回 0 (相同年份)
result = cve.CompareByYear("CVE-2022-1111", "CVE-2020-2222")   // 返回 2
result = cve.CompareByYear("cve-2022-1111", "CVE-2022-2222")   // 返回 0 (不区分大小写)
```
</details>

<details>
<summary><b>SubByYear</b> - 计算两个 CVE 的年份差值</summary>

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
diff = cve.SubByYear("CVE-2022-1111", "CVE-2020-2222")   // 返回 2
diff = cve.SubByYear("CVE-2022-1111", "CVE-2022-2222")   // 返回 0 (相同年份)
```
</details>

<details>
<summary><b>CompareCves</b> - 全面比较两个 CVE 编号的大小</summary>

```go
func CompareCves(cveA, cveB string) int
```

**参数**：
- `cveA` - 第一个 CVE 编号
- `cveB` - 第二个 CVE 编号

**返回值**：
- -1：cveA < cveB (年份更早或年份相同但序列号更小)
- 0：cveA = cveB (年份和序列号都相同)
- 1：cveA > cveB (年份更晚或年份相同但序列号更大)

**示例**：
```go
result := cve.CompareCves("CVE-2020-1111", "CVE-2022-2222")  // 返回 -1 (不同年份)
result = cve.CompareCves("CVE-2022-1111", "CVE-2022-2222")   // 返回 -1 (相同年份，不同序列号)
result = cve.CompareCves("CVE-2022-2222", "CVE-2022-2222")   // 返回 0 (完全相同)
result = cve.CompareCves("cve-2022-2222", "CVE-2022-2222")   // 返回 0 (不区分大小写)
result = cve.CompareCves("CVE-2022-3333", "CVE-2022-2222")   // 返回 1 (相同年份，序列号更大)
```
</details>

<details>
<summary><b>SortCves</b> - 对 CVE 切片进行排序（按年份和序列号）</summary>

```go
func SortCves(cveSlice []string) []string
```

**参数**：
- `cveSlice` - 要排序的 CVE 编号列表

**返回值**：
- 排序后的 CVE 编号列表（返回新的切片，不修改原切片）

**示例**：
```go
cveList := []string{"CVE-2022-2222", "cve-2020-1111", "CVE-2022-1111"}
sortedList := cve.SortCves(cveList)  // 返回 ["CVE-2020-1111", "CVE-2022-1111", "CVE-2022-2222"]

// 排序空切片或单元素切片
emptyList := cve.SortCves([]string{})           // 返回 []
singleList := cve.SortCves([]string{"CVE-2022-1111"})  // 返回 ["CVE-2022-1111"]
```
</details>

### 过滤与分组

<details>
<summary><b>FilterCvesByYear</b> - 筛选特定年份的 CVE</summary>

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
cves2022 := cve.FilterCvesByYear(cveList, 2022)  // 返回 ["CVE-2022-2222"]
cves2020 := cve.FilterCvesByYear(cveList, 2020)  // 返回 [] (没有2020年的CVE)
```
</details>

<details>
<summary><b>FilterCvesByYearRange</b> - 筛选指定年份范围内的 CVE</summary>

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
rangeCves = cve.FilterCvesByYearRange(cveList, 2020, 2020)   // 返回 ["CVE-2020-1111"]
rangeCves = cve.FilterCvesByYearRange(cveList, 2023, 2025)   // 返回 [] (没有该范围内的CVE)
// 注意：如果startYear > endYear，则返回空切片
rangeCves = cve.FilterCvesByYearRange(cveList, 2022, 2020)   // 返回 [] (无效范围)
```
</details>

<details>
<summary><b>GetRecentCves</b> - 获取最近几年的 CVE</summary>

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
recentCves = cve.GetRecentCves(cveList, 3)   // 返回 ["CVE-2021-2222", "CVE-2022-3333", "CVE-2023-4444"]
recentCves = cve.GetRecentCves(cveList, 1)   // 返回 ["CVE-2023-4444"]
```
</details>

<details>
<summary><b>GroupByYear</b> - 按年份对 CVE 进行分组</summary>

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

// 处理空切片
emptyCves := cve.GroupByYear([]string{})  // 返回空映射 map[string][]string{}
```
</details>

<details>
<summary><b>RemoveDuplicateCves</b> - 移除重复的 CVE 编号</summary>

```go
func RemoveDuplicateCves(cveSlice []string) []string
```

**参数**：
- `cveSlice` - 可能包含重复项的 CVE 编号列表

**返回值**：
- 去重后的 CVE 编号列表

**示例**：
```go
cveList := []string{"CVE-2022-1111", "cve-2022-1111", "CVE-2022-2222", "CVE-2022-1111"}
uniqueCves := cve.RemoveDuplicateCves(cveList)  // 返回 ["CVE-2022-1111", "CVE-2022-2222"]

// 处理空切片
emptyCves := cve.RemoveDuplicateCves([]string{})  // 返回 []
```
</details>

### 生成与构造

<details>
<summary><b>GenerateCve</b> - 根据年份和序列号生成标准格式的 CVE 编号</summary>

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
cveId := cve.GenerateCve(2022, 12345)   // 返回 "CVE-2022-12345"
cveId = cve.GenerateCve(2021, 44228)    // 返回 "CVE-2021-44228" (Log4Shell)
cveId = cve.GenerateCve(2020, 0)        // 返回 "CVE-2020-0"
cveId = cve.GenerateCve(2023, 123456)   // 返回 "CVE-2023-123456"
```
</details>

<details>
<summary><b>GenerateFakeCve</b> - 生成一个基于当前年份的随机 CVE 编号</summary>

```go
func GenerateFakeCve() string
```

**参数**：
- 无

**返回值**：
- 随机生成的 CVE 编号（当前年份+随机序列号）

**示例**：
```go
// 假设当前年份是2023年
fakeCve := cve.GenerateFakeCve()  // 返回类似 "CVE-2023-12345" 的格式
```
</details>

## 🔍 使用场景示例

### 漏洞报告分析

```go
package main

import (
	"fmt"
	"github.com/scagogogo/cve"
)

func main() {
	// 从文本中提取所有 CVE 并按年份分组
	text := `安全公告：系统受到多个漏洞影响，包括 CVE-2021-44228、
CVE-2021-45046、CVE-2022-1234 和 CVE-2022-5678。
建议尽快更新补丁。`

	// 提取所有 CVE
	cves := cve.ExtractCve(text)
	fmt.Println("提取的CVE:", cves)

	// 按年份分组
	groupedCves := cve.GroupByYear(cves)

	// 输出分组结果
	for year, yearCves := range groupedCves {
		fmt.Printf("%s年的CVE：%v\n", year, yearCves)
	}
}
```

### 漏洞库管理

```go
package main

import (
	"fmt"
	"github.com/scagogogo/cve"
)

func main() {
	// 导入 CVE 并进行去重和排序
	importedCves := []string{
		"CVE-2022-1111", "cve-2022-1111", "CVE-2021-2222",
		"CVE-2020-3333", "CVE-2022-4444", "CVE-2022-1111",
	}

	// 去重
	uniqueCves := cve.RemoveDuplicateCves(importedCves)
	fmt.Println("去重后:", uniqueCves)

	// 排序
	sortedCves := cve.SortCves(uniqueCves)
	fmt.Println("排序后:", sortedCves)

	// 获取最近两年的 CVE
	recentCves := cve.GetRecentCves(sortedCves, 2)
	fmt.Println("最近两年的 CVE:", recentCves)
}
```

### CVE 验证和处理

```go
package main

import (
	"fmt"
	"github.com/scagogogo/cve"
)

func main() {
	// 验证用户输入的 CVE 格式
	userInput := " cve-2022-12345 "

	if cve.IsCve(userInput) {
		formattedCve := cve.Format(userInput)
		fmt.Println("有效的 CVE:", formattedCve)
		
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
}
```

## 📂 项目结构

本项目按功能模块拆分为多个文件:

| 文件名 | 描述 |
|-------|------|
| **cve.go** | 包的主入口，包含版本信息 |
| **base.go** | 基础功能，如格式化、验证CVE格式等 |
| **extract.go** | 提取功能，从文本中提取CVE编号及其组成部分 |
| **compare.go** | 比较功能，比较CVE的年份和序列号，排序等 |
| **filter.go** | 过滤功能，按照年份范围过滤CVE，去重等 |
| **generate.go** | 生成功能，创建标准格式的CVE编号 |

每个功能模块都有对应的测试文件，确保功能正确性。此外，`examples` 目录下包含了每个函数的使用示例。

## 📚 参考资料

- [CVE 官方网站](https://cve.mitre.org/)
- [NIST 国家漏洞数据库](https://nvd.nist.gov/)
- [CISA 已知漏洞目录](https://www.cisa.gov/known-exploited-vulnerabilities-catalog)
- [CVSS 评分系统](https://www.first.org/cvss/)

## 📄 许可证

本项目使用 MIT 许可证。详见 [LICENSE](LICENSE) 文件。 
