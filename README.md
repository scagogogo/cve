# CVE Utils

[![Go Tests](https://github.com/scagogogo/cve/actions/workflows/go-test.yml/badge.svg)](https://github.com/scagogogo/cve/actions/workflows/go-test.yml)
[![Documentation](https://github.com/scagogogo/cve/actions/workflows/docs.yml/badge.svg)](https://github.com/scagogogo/cve/actions/workflows/docs.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/scagogogo/cve.svg)](https://pkg.go.dev/github.com/scagogogo/cve)
[![Go Report Card](https://goreportcard.com/badge/github.com/scagogogo/cve)](https://goreportcard.com/report/github.com/scagogogo/cve)
[![License](https://img.shields.io/github/license/scagogogo/cve)](https://github.com/scagogogo/cve/blob/main/LICENSE)
[![Version](https://img.shields.io/badge/version-v0.0.1-blue)](https://github.com/scagogogo/cve/releases)

**🌐 Languages: [English](#english) | [简体中文](#简体中文)**

---

## English

A comprehensive collection of utility functions for handling CVE (Common Vulnerabilities and Exposures) identifiers. This package provides a series of practical functions for processing, validating, extracting, and manipulating CVE identifiers.

### 📖 Documentation

**Complete API documentation and usage guides: [https://scagogogo.github.io/cve/](https://scagogogo.github.io/cve/)**

Documentation includes:
- 🚀 [Quick Start Guide](https://scagogogo.github.io/cve/guide/getting-started)
- 📚 [Complete API Reference](https://scagogogo.github.io/cve/api/)
- 💡 [Practical Examples](https://scagogogo.github.io/cve/examples/)
- 🔧 [Installation & Configuration](https://scagogogo.github.io/cve/guide/installation)

### 📑 Table of Contents

- [Features](#-features)
- [Installation](#-installation)
- [Quick Start](#-quick-start)
- [API Reference](#-api-reference)
  - [Format & Validation](#format--validation)
  - [Extraction Methods](#extraction-methods)
  - [Comparison & Sorting](#comparison--sorting)
  - [Filtering & Grouping](#filtering--grouping)
  - [Generation & Construction](#generation--construction)
- [Usage Examples](#-usage-examples)
- [Project Structure](#-project-structure)
- [References](#-references)
- [License](#-license)

### ✨ Features

- ✅ CVE format validation and standardization
- ✅ Extract CVE identifiers from text
- ✅ Extract and compare CVE years and sequence numbers
- ✅ Sort, filter, and group CVEs
- ✅ Generate standard format CVE identifiers
- ✅ Deduplication and validation tools

### 📦 Installation

```bash
go get github.com/scagogogo/cve
```

### 🚦 Quick Start

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

<details open>
<summary><b>Format</b> - 将 CVE 编号转换为标准大写格式并移除前后空格</summary>

```go
func Format(cve string) string
```

**参数**：
- `cve` - 要格式化的 CVE 编号（字符串类型）

**返回值**：
- 标准化格式的 CVE 编号（大写，无前后空白）

**详细说明**：
- 此函数将传入的字符串转为大写并移除前后空白字符
- 即使输入的不是有效的 CVE 格式，也会返回转换后的结果
- 不检查 CVE 编号的有效性，仅进行格式标准化

**示例**：
```go
formattedCve := cve.Format(" cve-2022-12345 ")  // 返回 "CVE-2022-12345"
formattedCve = cve.Format("CVE-2021-44228")     // 返回 "CVE-2021-44228"
formattedCve = cve.Format("  cve-invalid ")     // 返回 "CVE-INVALID" (即使无效也会格式化)
```

**常见用途**：
- 在比较或存储 CVE 编号前进行标准化
- 用户输入处理前的清洗
- 确保与 CVE 数据库进行匹配时的一致性
</details>

<details open>
<summary><b>IsCve</b> - 判断字符串是否是有效的 CVE 格式</summary>

```go
func IsCve(text string) bool
```

**参数**：
- `text` - 要检查的字符串

**返回值**：
- 如果字符串是有效的 CVE 格式则返回 `true`，否则返回 `false`

**详细说明**：
- 验证字符串是否完全符合 CVE 格式（允许两侧有空白字符）
- 使用正则表达式 `(?i)^\s*CVE-\d+-\d+\s*$` 进行匹配
- 不验证 CVE 的年份或序列号是否在合理范围，仅检查格式

**示例**：
```go
isCve := cve.IsCve("CVE-2022-12345")           // 返回 true
isCve = cve.IsCve(" CVE-2022-12345 ")          // 返回 true (忽略前后空格)
isCve = cve.IsCve("包含CVE-2022-12345的文本")    // 返回 false (含有额外文本)
isCve = cve.IsCve("cve2022-12345")             // 返回 false (格式不正确，缺少连字符)
isCve = cve.IsCve("CVE-2022-ABCDE")            // 返回 false (序列号不是数字)
```

**常见用途**：
- 验证用户输入的字符串是否为有效的 CVE 编号
- 在进行更严格的 CVE 验证前进行初步格式检查
- 用于表单验证或数据导入前的检查
</details>

<details open>
<summary><b>IsContainsCve</b> - 判断字符串是否包含 CVE</summary>

```go
func IsContainsCve(text string) bool
```

**参数**：
- `text` - 要检查的字符串

**返回值**：
- 如果字符串包含 CVE 则返回 `true`，否则返回 `false`

**详细说明**：
- 检查字符串中是否包含 CVE 格式的内容，不要求字符串仅包含 CVE
- 使用正则表达式 `(?i)CVE-\d+-\d+` 进行匹配
- 不区分大小写，可匹配 "CVE" 或 "cve"
- 只要找到一个符合格式的 CVE，就返回 true

**示例**：
```go
contains := cve.IsContainsCve("这个漏洞的编号是CVE-2022-12345")  // 返回 true
contains = cve.IsContainsCve("修复了cve-2021-44228漏洞")        // 返回 true (不区分大小写)
contains = cve.IsContainsCve("这个文本不包含任何CVE标识符")        // 返回 false
contains = cve.IsContainsCve("错误格式：CVE2022-12345")         // 返回 false (格式不正确)
contains = cve.IsContainsCve("多个CVE-2021-44228和CVE-2022-12345") // 返回 true
```

**常见用途**：
- 快速检查文本是否提及任何 CVE
- 在日志或报告中查找漏洞信息
- 文档扫描中预筛选可能包含漏洞信息的内容
</details>

<details open>
<summary><b>IsCveYearOk</b> - 判断 CVE 的年份是否在合理的时间范围内</summary>

```go
func IsCveYearOk(cve string) bool
```

**参数**：
- `cve` - CVE 编号

**返回值**：
- 如果年份在1999年之后且不超过当前年份则返回 `true`，否则返回 `false`

**详细说明**：
- CVE 编号系统始于 1999 年，因此有效年份必须 >= 1999
- 年份不应超过当前年份，防止未来日期
- 此函数会先格式化输入的 CVE 编号
- 内部调用 `IsCveYearOkWithCutoff` 函数，偏移量设为 0

**示例**：
```go
// 假设当前年份是2023年
isYearOk := cve.IsCveYearOk("CVE-2022-12345")  // 返回 true
isYearOk = cve.IsCveYearOk("CVE-2023-12345")   // 返回 true (当前年份)
isYearOk = cve.IsCveYearOk("CVE-2030-12345")   // 返回 false (2030 > 2023)
isYearOk = cve.IsCveYearOk("CVE-1998-12345")   // 返回 false (1998 < 1999)
isYearOk = cve.IsCveYearOk("不是有效的CVE")       // 返回 false (无效格式)
```

**常见用途**：
- 验证 CVE 年份是否在合理范围内
- 筛选出可能为伪造或错误的 CVE 编号
- 作为 `ValidateCve` 函数的一部分使用
</details>

<details open>
<summary><b>IsCveYearOkWithCutoff</b> - 判断 CVE 的年份是否在合理的时间范围内（可设置偏移量）</summary>

```go
func IsCveYearOkWithCutoff(cve string, cutoff int) bool
```

**参数**：
- `cve` - CVE 编号
- `cutoff` - 允许的年份偏移量（正整数）

**返回值**：
- 如果年份在合理范围内则返回 `true`，否则返回 `false`

**详细说明**：
- 允许的年份范围为：1999 <= 年份 <= 当前年份 + cutoff
- 通过设置 cutoff 可以允许未来一定年份范围内的 CVE
- 适用于预留或预分配的 CVE 编号场景
- 内部先格式化输入的 CVE 编号

**示例**：
```go
// 假设当前年份是2023年
isYearOk := cve.IsCveYearOkWithCutoff("CVE-2022-12345", 0)  // 返回 true
isYearOk = cve.IsCveYearOkWithCutoff("CVE-2025-12345", 2)   // 返回 true (2025 <= 2023+2)
isYearOk = cve.IsCveYearOkWithCutoff("CVE-2030-12345", 5)   // 返回 false (2030 > 2023+5)
isYearOk = cve.IsCveYearOkWithCutoff("CVE-1998-12345", 0)   // 返回 false (1998 < 1999)
isYearOk = cve.IsCveYearOkWithCutoff("CVE-2024-12345", 1)   // 返回 true (2024 <= 2023+1)
```

**常见用途**：
- 设置灵活的 CVE 年份验证规则
- 处理预发布或预分配的 CVE 编号
- 在特定场景下验证带有未来日期的 CVE
</details>

<details open>
<summary><b>Split</b> - 将 CVE 分割成年份和编号两部分</summary>

```go
func Split(cve string) (year string, seq string)
```

**参数**：
- `cve` - 要分割的 CVE 编号

**返回值**：
- `year` - CVE 的年份部分（字符串类型）
- `seq` - CVE 的序列号部分（字符串类型）

**详细说明**：
- 将标准格式的 CVE（如 CVE-2022-12345）分割成两部分
- 函数内部会先格式化输入的 CVE 编号
- 如果输入不是有效的 CVE 格式，则返回空字符串
- 返回的是字符串类型，如需数值请使用 `strconv.Atoi` 转换

**示例**：
```go
year, seq := cve.Split("CVE-2022-12345")  // 返回 year="2022", seq="12345"
year, seq = cve.Split("cve-2021-44228")   // 返回 year="2021", seq="44228"
year, seq = cve.Split("不是CVE格式")        // 返回 year="", seq=""
year, seq = cve.Split("CVE-2022")         // 返回 year="", seq="" (不完整的CVE)
```

**常见用途**：
- 需要单独处理 CVE 的年份或序列号部分时使用
- 作为其他提取函数的基础
- 用于自定义 CVE 处理逻辑
</details>

<details open>
<summary><b>ValidateCve</b> - 全面验证 CVE 编号的合法性</summary>

```go
func ValidateCve(cve string) bool
```

**参数**：
- `cve` - 要验证的 CVE 编号

**返回值**：
- 如果 CVE 编号合法则返回 `true`，否则返回 `false`

**详细说明**：
- 执行最全面的 CVE 验证，包括格式、年份和序列号
- 验证规则如下：
  1. 必须符合 CVE 格式（CVE-YYYY-NNNNN）
  2. 年份必须在 1999 到当前年份之间
  3. 序列号必须为正整数
- 相比 `IsCve`，此函数增加了年份和序列号的合理性检查

**示例**：
```go
isValid := cve.ValidateCve("CVE-2022-12345")  // 正常情况返回 true
isValid = cve.ValidateCve("CVE-1998-12345")   // 返回 false (年份 < 1999)
isValid = cve.ValidateCve("CVE-2099-12345")   // 返回 false (假设当前为2023年，年份超前太多)
isValid = cve.ValidateCve("CVE-2022-0")       // 返回 false (序列号必须为正整数)
isValid = cve.ValidateCve("CVE2022-12345")    // 返回 false (缺少连字符)
isValid = cve.ValidateCve("CVE-2022-ABC")     // 返回 false (序列号必须为数字)
```

**完整代码示例**：
```go
func main() {
    userInput := "CVE-2022-12345"
    
    if cve.ValidateCve(userInput) {
        fmt.Println("这是有效的CVE编号")
        
        // 可以安全地使用该CVE进行后续操作
        year := cve.ExtractCveYear(userInput)
        fmt.Printf("CVE年份: %s\n", year)
    } else {
        fmt.Println("无效的CVE编号")
    }
}
```

**常见用途**：
- 用户输入的 CVE 编号的完整验证
- 导入 CVE 数据之前的验证
- 确保处理的 CVE 编号完全符合标准
</details>

### 提取方法

<details open>
<summary><b>ExtractCve</b> - 从字符串中提取所有 CVE 编号</summary>

```go
func ExtractCve(text string) []string
```

**参数**：
- `text` - 要从中提取 CVE 的文本

**返回值**：
- 提取的 CVE 编号列表，按标准格式返回（大写）

**详细说明**：
- 使用正则表达式 `(?i)(CVE-\d+-\d+)` 从文本中匹配所有 CVE
- 提取的每个 CVE 都会自动格式化为标准格式（大写）
- 如果文本中没有 CVE，返回空切片 `[]`
- 返回的结果可能包含重复的 CVE（如果文本中多次出现）

**示例**：
```go
text := "系统受到CVE-2021-44228和cve-2022-12345的影响"
cveList := cve.ExtractCve(text)  // 返回 ["CVE-2021-44228", "CVE-2022-12345"]

text = "没有包含任何CVE的文本"
cveList = cve.ExtractCve(text)   // 返回 [] (空切片)

text = "重复出现的CVE-2021-44228和CVE-2021-44228"
cveList = cve.ExtractCve(text)   // 返回 ["CVE-2021-44228", "CVE-2021-44228"] (有重复)
```

**完整代码示例**：
```go
func main() {
    report := `安全公告：系统受到多个漏洞影响，包括：
    - CVE-2021-44228 (Log4Shell)
    - CVE-2022-22965 (Spring4Shell)
    - 还有一些未公开的漏洞`
    
    // 提取所有CVE
    cves := cve.ExtractCve(report)
    
    // 打印结果
    fmt.Printf("发现 %d 个CVE:\n", len(cves))
    for i, id := range cves {
        fmt.Printf("%d. %s\n", i+1, id)
    }
    
    // 去重（可选）
    uniqueCves := cve.RemoveDuplicateCves(cves)
    if len(uniqueCves) < len(cves) {
        fmt.Println("去重后:", uniqueCves)
    }
}
```

**常见用途**：
- 从安全公告或漏洞报告中提取所有相关的 CVE 编号
- 自动化处理大量文本中的漏洞信息
- 构建 CVE 数据库或索引
</details>

<details open>
<summary><b>ExtractFirstCve</b> - 从字符串中提取第一个 CVE 编号</summary>

```go
func ExtractFirstCve(text string) string
```

**参数**：
- `text` - 要从中提取 CVE 的文本

**返回值**：
- 第一个 CVE 编号（标准格式），如果没有找到则返回空字符串

**详细说明**：
- 只返回文本中出现的第一个 CVE 编号
- 从左到右扫描文本，返回第一个匹配的结果
- 返回的 CVE 编号会格式化为标准格式（大写）
- 比完整提取更高效，当只需要第一个 CVE 时使用

**示例**：
```go
text := "系统受到CVE-2021-44228和CVE-2022-12345的影响"
firstCve := cve.ExtractFirstCve(text)  // 返回 "CVE-2021-44228"

text = "没有包含任何CVE的文本"
firstCve = cve.ExtractFirstCve(text)   // 返回 "" (空字符串)

text = "提到了 cve-2022-12345 这个漏洞"
firstCve = cve.ExtractFirstCve(text)   // 返回 "CVE-2022-12345" (转为大写)
```

**常见用途**：
- 当只需要获取文本中第一个提到的 CVE 时使用
- 快速检测最主要的漏洞标识符
- 提取标题或摘要中的主要 CVE
</details>

<details open>
<summary><b>ExtractLastCve</b> - 从字符串中提取最后一个 CVE 编号</summary>

```go
func ExtractLastCve(text string) string
```

**参数**：
- `text` - 要从中提取 CVE 的文本

**返回值**：
- 最后一个 CVE 编号（标准格式），如果没有找到则返回空字符串

**详细说明**：
- 提取文本中出现的最后一个 CVE 编号
- 内部首先调用 `ExtractCve` 提取所有 CVE，然后返回最后一个
- 如果文本中没有 CVE，返回空字符串
- 返回的 CVE 编号会格式化为标准格式（大写）

**示例**：
```go
text := "系统受到CVE-2021-44228和CVE-2022-12345的影响"
lastCve := cve.ExtractLastCve(text)  // 返回 "CVE-2022-12345"

text = "没有包含任何CVE的文本"
lastCve = cve.ExtractLastCve(text)   // 返回 "" (空字符串)

text = "最后提到了 cve-2021-44228"
lastCve = cve.ExtractLastCve(text)   // 返回 "CVE-2021-44228" (转为大写)
```

**常见用途**：
- 当需要获取文本中最后提到的 CVE 时使用
- 在有多个 CVE 的场景下，获取最新或最后出现的漏洞
- 处理按时间顺序排列的漏洞报告
</details>

<details open>
<summary><b>ExtractCveYear</b> - 从 CVE 中提取年份</summary>

```go
func ExtractCveYear(cve string) string
```

**参数**：
- `cve` - 要提取年份的 CVE 编号

**返回值**：
- CVE 的年份部分（字符串类型），如果不是有效 CVE 则返回空字符串

**详细说明**：
- 从 CVE 编号中提取年份部分，作为字符串返回
- 内部调用 `Split` 函数，返回年份部分
- 会先格式化输入的 CVE 编号
- 如果输入不是有效的 CVE 格式，返回空字符串

**示例**：
```go
year := cve.ExtractCveYear("CVE-2022-12345")  // 返回 "2022"
year = cve.ExtractCveYear("cve-2021-44228")   // 返回 "2021"
year = cve.ExtractCveYear("不是CVE格式")        // 返回 ""
year = cve.ExtractCveYear("CVE-ABCD-12345")   // 返回 "" (年份部分无效)
```

**常见用途**：
- 需要对 CVE 按年份进行分类或过滤时使用
- 作为字符串键值使用（如映射的键）
- 在不需要数值计算的场景使用
</details>

<details open>
<summary><b>ExtractCveYearAsInt</b> - 从 CVE 中提取年份并转换为整数</summary>

```go
func ExtractCveYearAsInt(cve string) int
```

**参数**：
- `cve` - 要提取年份的 CVE 编号

**返回值**：
- CVE 的年份（整数类型），如果不是有效 CVE 则返回 0

**详细说明**：
- 从 CVE 编号中提取年份并转换为整数类型
- 首先验证输入是否为有效的 CVE 格式
- 如果输入不是有效的 CVE 格式或年份部分无法转换为整数，返回 0
- 适用于需要对年份进行数值计算的场景

**示例**：
```go
year := cve.ExtractCveYearAsInt("CVE-2022-12345")  // 返回 2022
year = cve.ExtractCveYearAsInt("cve-2021-44228")   // 返回 2021
year = cve.ExtractCveYearAsInt("不是CVE格式")        // 返回 0
year = cve.ExtractCveYearAsInt("CVE-ABCD-12345")   // 返回 0 (年份部分无法转换为整数)
```

**完整代码示例**：
```go
func main() {
    cveId := "CVE-2022-12345"
    
    // 获取年份作为整数
    year := cve.ExtractCveYearAsInt(cveId)
    if year > 0 {
        currentYear := time.Now().Year()
        age := currentYear - year
        
        fmt.Printf("CVE %s 发布于 %d 年，距今已有 %d 年\n", 
                  cveId, year, age)
                  
        if age > 5 {
            fmt.Println("这是一个较老的漏洞")
        } else {
            fmt.Println("这是一个较新的漏洞")
        }
    }
}
```

**常见用途**：
- 需要对 CVE 年份进行数值计算或比较时使用
- 计算漏洞年龄或时间跨度
- 基于年份的统计和分析
</details>

<details open>
<summary><b>ExtractCveSeq</b> - 从 CVE 中提取序列号</summary>

```go
func ExtractCveSeq(cve string) string
```

**参数**：
- `cve` - 要提取序列号的 CVE 编号

**返回值**：
- CVE 的序列号部分（字符串类型），如果不是有效 CVE 则返回空字符串

**详细说明**：
- 从 CVE 编号中提取序列号部分，作为字符串返回
- 首先验证输入是否为有效的 CVE 格式
- 如果输入不是有效的 CVE 格式，返回空字符串
- 保留序列号前导零（如果有）

**示例**：
```go
seq := cve.ExtractCveSeq("CVE-2022-12345")  // 返回 "12345"
seq = cve.ExtractCveSeq("cve-2021-44228")   // 返回 "44228"
seq = cve.ExtractCveSeq("不是CVE格式")        // 返回 ""
seq = cve.ExtractCveSeq("CVE-2022-00123")   // 返回 "00123" (保留前导零)
```

**常见用途**：
- 需要单独处理 CVE 序列号时使用
- 作为字符串键值使用
- 在需要保留序列号原始格式（如前导零）的场景
</details>

<details open>
<summary><b>ExtractCveSeqAsInt</b> - 从 CVE 中提取序列号并转换为整数</summary>

```go
func ExtractCveSeqAsInt(cve string) int
```

**参数**：
- `cve` - 要提取序列号的 CVE 编号

**返回值**：
- CVE 的序列号（整数类型），如果不是有效 CVE 则返回 0

**详细说明**：
- 从 CVE 编号中提取序列号并转换为整数类型
- 首先通过 `ExtractCveSeq` 提取序列号字符串
- 然后将字符串转换为整数
- 如果输入不是有效的 CVE 格式或序列号部分无法转换为整数，返回 0
- 转换为整数会去除前导零

**示例**：
```go
seq := cve.ExtractCveSeqAsInt("CVE-2022-12345")  // 返回 12345
seq = cve.ExtractCveSeqAsInt("cve-2021-44228")   // 返回 44228
seq = cve.ExtractCveSeqAsInt("不是CVE格式")        // 返回 0
seq = cve.ExtractCveSeqAsInt("CVE-2022-00123")   // 返回 123 (去除前导零)
```

**常见用途**：
- 需要对 CVE 序列号进行数值计算或比较时使用
- 确定序列号的大小或范围
- 在序列号需要参与数值运算的场景
</details>

### 比较与排序

<details open>
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
- 正数：cveA 年份 > cveB 年份（具体值为年份差）

**详细说明**：
- 只比较 CVE 编号的年份部分，忽略序列号
- 返回的具体数值是两个 CVE 年份的差值（cveA 年份 - cveB 年份）
- 如果输入的 CVE 格式无效，将提取出 0 作为年份值进行比较
- 内部使用 `ExtractCveYearAsInt` 提取年份进行比较

**示例**：
```go
result := cve.CompareByYear("CVE-2020-1111", "CVE-2022-2222")  // 返回 -2
result = cve.CompareByYear("CVE-2022-1111", "CVE-2022-2222")   // 返回 0 (相同年份)
result = cve.CompareByYear("CVE-2022-1111", "CVE-2020-2222")   // 返回 2
result = cve.CompareByYear("cve-2022-1111", "CVE-2022-2222")   // 返回 0 (不区分大小写)
result = cve.CompareByYear("无效格式", "CVE-2022-2222")         // 返回 -2022 (无效格式视为年份0)
```

**常见用途**：
- 按年份对 CVE 进行排序
- 确定两个 CVE 的时间先后关系
- 用作排序函数的比较器
</details>

<details open>
<summary><b>SubByYear</b> - 计算两个 CVE 的年份差值</summary>

```go
func SubByYear(cveA, cveB string) int
```

**参数**：
- `cveA` - 第一个 CVE 编号
- `cveB` - 第二个 CVE 编号

**返回值**：
- cveA 年份 - cveB 年份的差值

**详细说明**：
- 计算两个 CVE 年份之间的差值
- 功能与 `CompareByYear` 相同，都返回年份差
- 将来可能会被 `CompareByYear` 替代
- 如果输入的 CVE 格式无效，将提取出 0 作为年份值进行计算

**示例**：
```go
diff := cve.SubByYear("CVE-2020-1111", "CVE-2022-2222")  // 返回 -2
diff = cve.SubByYear("CVE-2022-1111", "CVE-2020-2222")   // 返回 2
diff = cve.SubByYear("CVE-2022-1111", "CVE-2022-2222")   // 返回 0 (相同年份)
diff = cve.SubByYear("无效格式", "CVE-2020-2222")         // 返回 -2020
```

**常见用途**：
- 计算两个 CVE 的发布时间间隔
- 分析漏洞发现的时间趋势
- 基于年份的漏洞比较
</details>

<details open>
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

**详细说明**：
- 先比较 CVE 的年份，如果年份不同，返回年份比较结果
- 如果年份相同，则比较序列号
- 与标准的比较函数一致，返回 -1、0 或 1
- 无效的 CVE 格式视为年份和序列号都为 0
- 内部使用 `ExtractCveYearAsInt` 和 `ExtractCveSeqAsInt` 提取数据

**示例**：
```go
result := cve.CompareCves("CVE-2020-1111", "CVE-2022-2222")  // 返回 -1 (不同年份)
result = cve.CompareCves("CVE-2022-1111", "CVE-2022-2222")   // 返回 -1 (相同年份，不同序列号)
result = cve.CompareCves("CVE-2022-2222", "CVE-2022-2222")   // 返回 0 (完全相同)
result = cve.CompareCves("cve-2022-2222", "CVE-2022-2222")   // 返回 0 (不区分大小写)
result = cve.CompareCves("CVE-2022-3333", "CVE-2022-2222")   // 返回 1 (相同年份，序列号更大)
result = cve.CompareCves("CVE-2023-1", "CVE-2022-99999")     // 返回 1 (年份更大)
```

**完整代码示例**：
```go
func main() {
    // 对CVE列表进行排序
    cveList := []string{
        "CVE-2022-33891", // Apache Spark 
        "CVE-2021-44228", // Log4Shell
        "CVE-2022-22965", // Spring4Shell
        "CVE-2014-0160",  // Heartbleed
    }
    
    fmt.Println("原始CVE列表:")
    for _, id := range cveList {
        fmt.Println(id)
    }
    
    // 排序
    sortedList := cve.SortCves(cveList)
    
    fmt.Println("\n按时间顺序排序后:")
    for i, id := range sortedList {
        fmt.Printf("%d. %s\n", i+1, id)
    }
}
// 输出:
// 原始CVE列表:
// CVE-2022-33891
// CVE-2021-44228
// CVE-2022-22965
// CVE-2014-0160
//
// 按时间顺序排序后:
// 1. CVE-2014-0160
// 2. CVE-2021-44228
// 3. CVE-2022-22965
// 4. CVE-2022-33891
```

**常见用途**：
- 作为排序函数的比较器
- 按时间顺序（先年份后序列号）排列 CVE
- 在查找和排序算法中使用
</details>

<details open>
<summary><b>SortCves</b> - 对 CVE 切片进行排序（按年份和序列号）</summary>

```go
func SortCves(cveSlice []string) []string
```

**参数**：
- `cveSlice` - 要排序的 CVE 编号列表

**返回值**：
- 排序后的 CVE 编号列表（返回新的切片，不修改原切片）

**详细说明**：
- 按年份和序列号对 CVE 列表进行排序（从早到晚，从小到大）
- 内部使用 `CompareCves` 函数进行比较
- 返回新的排序后的切片，不修改原始输入
- 无效的 CVE 格式会排在有效 CVE 之前
- 排序结果是稳定的

**示例**：
```go
cveList := []string{"CVE-2022-2222", "cve-2020-1111", "CVE-2022-1111"}
sortedList := cve.SortCves(cveList)  // 返回 ["CVE-2020-1111", "CVE-2022-1111", "CVE-2022-2222"]

// 排序空切片或单元素切片
emptyList := cve.SortCves([]string{})           // 返回 []
singleList := cve.SortCves([]string{"CVE-2022-1111"})  // 返回 ["CVE-2022-1111"]

// 带有无效格式的排序
mixedList := cve.SortCves([]string{"无效格式", "CVE-2022-1111"})  
// 返回 ["无效格式", "CVE-2022-1111"] (无效格式排在前面)
```

**完整代码示例**：
```go
func main() {
    // 未排序的CVE列表
    cveList := []string{
        "CVE-2022-33891", // Apache Spark 
        "CVE-2021-44228", // Log4Shell
        "CVE-2022-22965", // Spring4Shell
        "CVE-2014-0160",  // Heartbleed
    }
    
    fmt.Println("原始CVE列表:")
    for _, id := range cveList {
        fmt.Println(id)
    }
    
    // 排序
    sortedList := cve.SortCves(cveList)
    
    fmt.Println("\n按时间顺序排序后:")
    for i, id := range sortedList {
        fmt.Printf("%d. %s\n", i+1, id)
    }
}
// 输出:
// 原始CVE列表:
// CVE-2022-33891
// CVE-2021-44228
// CVE-2022-22965
// CVE-2014-0160
//
// 按时间顺序排序后:
// 1. CVE-2014-0160
// 2. CVE-2021-44228
// 3. CVE-2022-22965
// 4. CVE-2022-33891
```

**常见用途**：
- 对 CVE 列表进行时间顺序排序
- 在漏洞报告或展示中按顺序显示 CVE
- 在分析和比较多个 CVE 时使用
</details>

### 过滤与分组

<details open>
<summary><b>FilterCvesByYear</b> - 筛选特定年份的 CVE</summary>

```go
func FilterCvesByYear(cveSlice []string, year int) []string
```

**参数**：
- `cveSlice` - CVE 编号列表
- `year` - 要筛选的年份（整数）

**返回值**：
- 指定年份的 CVE 编号列表

**详细说明**：
- 从给定的 CVE 列表中筛选出特定年份的 CVE
- 过滤时会先格式化 CVE 编号
- 如果列表中没有指定年份的 CVE，返回空切片 `[]`
- 不会修改原始输入的切片
- 保持原列表中 CVE 的顺序

**示例**：
```go
cveList := []string{"CVE-2021-1111", "CVE-2022-2222", "CVE-2021-3333"}
cves2021 := cve.FilterCvesByYear(cveList, 2021)  // 返回 ["CVE-2021-1111", "CVE-2021-3333"]
cves2022 := cve.FilterCvesByYear(cveList, 2022)  // 返回 ["CVE-2022-2222"]
cves2020 := cve.FilterCvesByYear(cveList, 2020)  // 返回 [] (没有2020年的CVE)

// 处理混合格式的列表
mixedList := []string{"cve-2021-1111", "CVE-2021-2222", "不是有效格式"}
cves2021 = cve.FilterCvesByYear(mixedList, 2021)  // 返回 ["CVE-2021-1111", "CVE-2021-2222"]
```

**完整代码示例**：
```go
func main() {
    // CVE列表，包含多个年份
    cveList := []string{
        "CVE-2022-33891", 
        "CVE-2021-44228", 
        "CVE-2022-22965", 
        "CVE-2021-3449",
        "CVE-2020-1234",
    }
    
    // 过滤2021年的CVE
    cves2021 := cve.FilterCvesByYear(cveList, 2021)
    
    // 过滤2022年的CVE
    cves2022 := cve.FilterCvesByYear(cveList, 2022)
    
    // 显示结果
    fmt.Println("2021年的CVE:")
    for _, id := range cves2021 {
        fmt.Println(" -", id)
    }
    
    fmt.Println("\n2022年的CVE:")
    for _, id := range cves2022 {
        fmt.Println(" -", id)
    }
}
```

**常见用途**：
- 按年份分类 CVE 记录
- 分析特定年份的漏洞数据
- 创建特定时间范围的漏洞报告
</details>

<details open>
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

**详细说明**：
- 筛选年份在 [startYear, endYear] 范围内的 CVE
- 包含起始年份和结束年份
- 如果 startYear > endYear，则返回空切片 `[]`
- 过滤时会先格式化 CVE 编号
- 保持原列表中 CVE 的顺序

**示例**：
```go
cveList := []string{"CVE-2020-1111", "CVE-2021-2222", "CVE-2022-3333"}
rangeCves := cve.FilterCvesByYearRange(cveList, 2021, 2022)  // 返回 ["CVE-2021-2222", "CVE-2022-3333"]
rangeCves = cve.FilterCvesByYearRange(cveList, 2020, 2020)   // 返回 ["CVE-2020-1111"]
rangeCves = cve.FilterCvesByYearRange(cveList, 2023, 2025)   // 返回 [] (没有该范围内的CVE)
// 注意：如果startYear > endYear，则返回空切片
rangeCves = cve.FilterCvesByYearRange(cveList, 2022, 2020)   // 返回 [] (无效范围)

// 处理混合格式的列表
mixedList := []string{"cve-2020-1111", "CVE-2021-2222", "不是有效格式", "CVE-2022-3333"}
rangeCves = cve.FilterCvesByYearRange(mixedList, 2020, 2021)  // 返回 ["CVE-2020-1111", "CVE-2021-2222"]
```

**完整代码示例**：
```go
func main() {
    // 导入的CVE列表
    cveList := []string{
        "CVE-2022-33891", 
        "CVE-2021-44228", 
        "CVE-2022-22965", 
        "CVE-2021-3449",
        "CVE-2020-1234",
        "CVE-2020-5902",
    }
    
    // 按年份分组
    groupedCves := cve.GroupByYear(cveList)
    
    // 按年份排序显示结果
    years := make([]string, 0, len(groupedCves))
    for year := range groupedCves {
        years = append(years, year)
    }
    sort.Strings(years)
    
    for _, year := range years {
        cvesInYear := groupedCves[year]
        fmt.Printf("%s年的CVE (%d个):\n", year, len(cvesInYear))
        for _, id := range cvesInYear {
            fmt.Printf("  - %s\n", id)
        }
        fmt.Println()
    }
}
```

**常见用途**：
- 按时间段筛选 CVE
- 分析特定时间范围内的漏洞趋势
- 创建时间跨度的安全报告
</details>

<details open>
<summary><b>GetRecentCves</b> - 获取最近几年的 CVE</summary>

```go
func GetRecentCves(cveSlice []string, years int) []string
```

**参数**：
- `cveSlice` - CVE 编号列表
- `years` - 最近几年（从当前年份往前计算）

**返回值**：
- 最近几年的 CVE 编号列表

**详细说明**：
- 筛选最近 `years` 年内的 CVE 编号
- 范围为 [当前年份-years+1, 当前年份]
- 如果 years <= 0，返回空切片 `[]`
- 内部使用 `FilterCvesByYearRange` 实现过滤
- 保持原列表中 CVE 的顺序

**示例**：
```go
// 假设当前年份是2023年
cveList := []string{"CVE-2020-1111", "CVE-2021-2222", "CVE-2022-3333", "CVE-2023-4444"}
recentCves := cve.GetRecentCves(cveList, 2)  // 返回 ["CVE-2022-3333", "CVE-2023-4444"]
recentCves = cve.GetRecentCves(cveList, 3)   // 返回 ["CVE-2021-2222", "CVE-2022-3333", "CVE-2023-4444"]
recentCves = cve.GetRecentCves(cveList, 1)   // 返回 ["CVE-2023-4444"]
recentCves = cve.GetRecentCves(cveList, 0)   // 返回 [] (无效参数)
```

**完整代码示例**：
```go
func main() {
    // 导入的CVE列表（跨多年）
    cveList := []string{
        "CVE-2023-23397", // 2023年
        "CVE-2022-26134", // 2022年
        "CVE-2021-44228", // 2021年 (Log4Shell)
        "CVE-2020-1472",  // 2020年 (Zerologon)
        "CVE-2019-11581", // 2019年
        "CVE-2018-13379", // 2018年
    }
    
    // 获取最近1年、2年和3年的CVE
    lastYear := cve.GetRecentCves(cveList, 1)
    last2Years := cve.GetRecentCves(cveList, 2)
    last3Years := cve.GetRecentCves(cveList, 3)
    
    // 计算每个时间段的CVE数量
    fmt.Printf("最近1年的CVE: %d 个\n", len(lastYear))
    fmt.Printf("最近2年的CVE: %d 个\n", len(last2Years))
    fmt.Printf("最近3年的CVE: %d 个\n", len(last3Years))
    
    // 输出最近2年的CVE详情
    fmt.Println("\n最近2年的CVE详情:")
    for _, id := range last2Years {
        fmt.Println(" -", id)
    }
}
```

**常见用途**：
- 获取近期的漏洞信息
- 分析最近几年的安全趋势
- 优先处理最新的漏洞
</details>

<details open>
<summary><b>GroupByYear</b> - 按年份对 CVE 进行分组</summary>

```go
func GroupByYear(cveSlice []string) map[string][]string
```

**参数**：
- `cveSlice` - 要分组的 CVE 编号列表

**返回值**：
- 按年份分组的 CVE 编号映射表，键为年份（字符串），值为该年份的 CVE 列表

**详细说明**：
- 将 CVE 列表按年份进行分组
- 返回的映射表中，键为年份字符串（如 "2022"），值为该年份的 CVE 列表
- 分组前会对每个 CVE 进行格式化
- 无效格式的 CVE 将被忽略
- 保持每个年份组中 CVE 的原始顺序

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

// 处理包含无效格式的列表
mixedList := []string{"CVE-2021-1111", "不是有效格式", "CVE-2022-2222"}
groupedMixed := cve.GroupByYear(mixedList)
// 返回:
// {
//   "2021": ["CVE-2021-1111"],
//   "2022": ["CVE-2022-2222"]
// }
// (无效格式被忽略)
```

**完整代码示例**：
```go
func main() {
    // 导入的CVE列表
    cveList := []string{
        "CVE-2022-33891", 
        "CVE-2021-44228", 
        "CVE-2022-22965", 
        "CVE-2021-3449",
        "CVE-2020-1234",
        "CVE-2020-5902",
    }
    
    // 按年份分组
    groupedCves := cve.GroupByYear(cveList)
    
    // 按年份排序显示结果
    years := make([]string, 0, len(groupedCves))
    for year := range groupedCves {
        years = append(years, year)
    }
    sort.Strings(years)
    
    for _, year := range years {
        cvesInYear := groupedCves[year]
        fmt.Printf("%s年的CVE (%d个):\n", year, len(cvesInYear))
        for _, id := range cvesInYear {
            fmt.Printf("  - %s\n", id)
        }
        fmt.Println()
    }
}
```

**常见用途**：
- 按年份组织和分类 CVE
- 生成按年份划分的安全报告
- 分析不同年份的漏洞分布
</details>

<details open>
<summary><b>RemoveDuplicateCves</b> - 移除重复的 CVE 编号</summary>

```go
func RemoveDuplicateCves(cveSlice []string) []string
```

**参数**：
- `cveSlice` - 可能包含重复项的 CVE 编号列表

**返回值**：
- 去重后的 CVE 编号列表

**详细说明**：
- 移除列表中重复的 CVE 编号
- 先对每个 CVE 进行格式化，确保不区分大小写的去重
- 保持第一次出现的 CVE 的顺序
- 如果输入为空切片，返回空切片
- 无效格式的 CVE 也会被当作独立项处理

**示例**：
```go
cveList := []string{"CVE-2022-1111", "cve-2022-1111", "CVE-2022-2222", "CVE-2022-1111"}
uniqueCves := cve.RemoveDuplicateCves(cveList)  // 返回 ["CVE-2022-1111", "CVE-2022-2222"]

// 处理空切片
emptyCves := cve.RemoveDuplicateCves([]string{})  // 返回 []

// 处理包含无效格式的列表
mixedList := []string{"CVE-2022-1111", "不是有效格式", "CVE-2022-1111", "另一个无效格式"}
uniqueMixed := cve.RemoveDuplicateCves(mixedList)  // 返回 ["CVE-2022-1111", "不是有效格式", "另一个无效格式"]
```

**完整代码示例**：
```go
func main() {
    // 包含重复项的CVE列表（因爬虫爬取或用户输入等原因）
    cveList := []string{
        "CVE-2022-33891", 
        "cve-2022-33891", // 重复，大小写不同
        "CVE-2021-44228", 
        "CVE-2022-22965", 
        "CVE-2021-44228",  // 重复
        "CVE-2020-1234",
    }
    
    // 计算去重前数量
    fmt.Printf("去重前: %d 条, 去重后: %d 条\n", len(cveList), len(cve.RemoveDuplicateCves(cveList)))
    
    // 去重
    uniqueCves := cve.RemoveDuplicateCves(cveList)
    
    // 显示去重结果
    fmt.Println("\n去重后的CVE列表:")
    for i, id := range uniqueCves {
        fmt.Printf("%d. %s\n", i+1, id)
    }
}
```

**常见用途**：
- 清理从不同来源收集的 CVE 数据
- 确保数据库或报告中不出现重复
- 合并多个数据源时去除重复项
</details>

### 生成与构造

<details open>
<summary><b>GenerateCve</b> - 根据年份和序列号生成标准格式的 CVE 编号</summary>

```go
func GenerateCve(year int, seq int) string
```

**参数**：
- `year` - CVE 年份（整数）
- `seq` - CVE 序列号（整数）

**返回值**：
- 生成的标准格式 CVE 编号

**详细说明**：
- 根据提供的年份和序列号，生成标准格式的 CVE 编号
- 生成格式为 `CVE-YYYY-NNNN`
- 不检查年份或序列号的有效性，仅进行格式化
- 序列号不会自动补零
- 如果输入负数，仍会转换为字符串（但不符合标准）

**示例**：
```go
cveId := cve.GenerateCve(2022, 12345)   // 返回 "CVE-2022-12345"
cveId = cve.GenerateCve(2021, 44228)    // 返回 "CVE-2021-44228" (Log4Shell)
cveId = cve.GenerateCve(2020, 0)        // 返回 "CVE-2020-0"
cveId = cve.GenerateCve(2023, 123456)   // 返回 "CVE-2023-123456"
cveId = cve.GenerateCve(2020, 42)       // 返回 "CVE-2020-42" (不会自动补零)
```

**完整代码示例**：
```go
func main() {
    // 从外部数据中获取的年份和序列号
    year := 2023
    seqNumbers := []int{1001, 1002, 1003, 1004}
    
    // 为每个序列号生成标准CVE标识符
    cveIds := make([]string, len(seqNumbers))
    for i, seq := range seqNumbers {
        cveIds[i] = cve.GenerateCve(year, seq)
    }
    
    // 显示生成的CVE编号
    fmt.Println("生成的CVE编号:")
    for i, id := range cveIds {
        fmt.Printf("%d. %s\n", i+1, id)
    }
    
    // 验证生成的CVE格式是否有效
    for _, id := range cveIds {
        if cve.ValidateCve(id) {
            fmt.Printf("%s: 有效\n", id)
        } else {
            fmt.Printf("%s: 无效\n", id)
        }
    }
}
```

**常见用途**：
- 根据数据库中存储的年份和序列号重建CVE标识符
- 在报告生成过程中创建标准格式的CVE
- 将拆分的CVE组件重新组合
</details>

<details open>
<summary><b>GenerateFakeCve</b> - 生成一个基于当前年份的随机 CVE 编号</summary>

```go
func GenerateFakeCve() string
```

**参数**：
- 无

**返回值**：
- 随机生成的 CVE 编号（当前年份+随机序列号）

**详细说明**：
- 生成基于当前年份的随机 CVE 编号
- 随机序列号在 1 到 99999 之间
- 用于测试、示例或占位符
- 返回的是标准格式的 CVE 编号
- 不保证生成的 CVE 编号不存在

**示例**：
```go
// 假设当前年份是2023年
fakeCve := cve.GenerateFakeCve()  // 可能返回 "CVE-2023-12345" (随机序列号)
anotherFake := cve.GenerateFakeCve()  // 可能返回 "CVE-2023-54321" (不同的随机序列号)
```

**完整代码示例**：
```go
func main() {
    // 生成10个随机CVE编号用于测试
    fmt.Println("生成的随机CVE编号:")
    for i := 1; i <= 10; i++ {
        fakeCve := cve.GenerateFakeCve()
        fmt.Printf("%d. %s\n", i, fakeCve)
    }
    
    // 使用随机生成的CVE创建模拟数据集
    fakeCves := make([]string, 5)
    for i := range fakeCves {
        fakeCves[i] = cve.GenerateFakeCve()
    }
    
    // 应用库中的函数进行测试
    groupedFakes := cve.GroupByYear(fakeCves)
    fmt.Println("\n按年份分组的随机CVE:")
    for year, cves := range groupedFakes {
        fmt.Printf("%s年: %v\n", year, cves)
    }
}
```

**常见用途**：
- 生成测试数据
- 在示例代码或文档中使用
- 作为开发或测试环境中的占位符
- 在模拟漏洞场景时使用
</details>

## 🔍 使用场景示例

### 漏洞报告分析

```go
package main

import (
	"fmt"
	"github.com/scagogogo/cve"
	"sort"
	"time"
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

	// 按顺序输出分组结果
	years := make([]string, 0, len(groupedCves))
	for year := range groupedCves {
		years = append(years, year)
	}
	sort.Strings(years)

	fmt.Println("\n按年份分组结果:")
	for _, year := range years {
		fmt.Printf("%s年的CVE：%v\n", year, groupedCves[year])
	}
	
	// 计算CVE的年龄（相对于当前年份）
	currentYear := time.Now().Year()
	fmt.Println("\nCVE年龄分析:")
	for _, id := range cves {
		year := cve.ExtractCveYearAsInt(id)
		age := currentYear - year
		ageDesc := "新"
		if age > 1 {
			ageDesc = "旧"
		}
		fmt.Printf("%s: %d年前发布 (%s漏洞)\n", id, age, ageDesc)
	}
}
```

### 漏洞库管理

```go
package main

import (
	"fmt"
	"github.com/scagogogo/cve"
	"time"
)

func main() {
	// 导入 CVE 并进行去重和排序
	importedCves := []string{
		"CVE-2022-1111", "cve-2022-1111", "CVE-2021-2222",
		"CVE-2020-3333", "CVE-2022-4444", "CVE-2022-1111",
	}

	// 去重
	uniqueCves := cve.RemoveDuplicateCves(importedCves)
	fmt.Printf("去重前: %d 条, 去重后: %d 条\n", len(importedCves), len(uniqueCves))

	// 排序
	sortedCves := cve.SortCves(uniqueCves)
	fmt.Println("排序后:", sortedCves)

	// 获取最近两年的 CVE
	currentYear := time.Now().Year()
	recentCves := cve.FilterCvesByYearRange(sortedCves, currentYear-1, currentYear)
	fmt.Printf("\n最近两年(%d-%d)的 CVE: %v\n", currentYear-1, currentYear, recentCves)
	
	// 按年份统计数量
	groupedCves := cve.GroupByYear(sortedCves)
	fmt.Println("\n各年份CVE数量统计:")
	for year, yearCves := range groupedCves {
		fmt.Printf("%s年: %d条\n", year, len(yearCves))
	}
	
	// 查找特定CVE并验证
	searchCve := "CVE-2021-2222"
	for _, id := range sortedCves {
		if id == searchCve {
			year := cve.ExtractCveYear(id)
			seq := cve.ExtractCveSeq(id)
			fmt.Printf("\n找到CVE: %s (年份: %s, 序列号: %s)\n", id, year, seq)
			
			if cve.ValidateCve(id) {
				fmt.Println("验证结果: 有效")
			} else {
				fmt.Println("验证结果: 无效")
			}
			break
		}
	}
}
```

### CVE 验证和处理

```go
package main

import (
	"fmt"
	"github.com/scagogogo/cve"
	"strings"
	"time"
)

func main() {
	// 验证用户输入的 CVE 格式
	userInputs := []string{
		" cve-2022-12345 ",
		"CVE-2021-44228",
		"CVE-1998-1234",   // 年份过早
		"CVE-2099-5678",   // 年份过晚
		"CVE2022-1234",    // 格式错误
		"不是CVE格式",
	}

	currentYear := time.Now().Year()
	
	fmt.Println("CVE验证结果:")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Printf("%-20s %-10s %-10s %-10s\n", "输入", "格式正确", "年份有效", "完全有效")
	fmt.Println(strings.Repeat("-", 50))
	
	for _, input := range userInputs {
		// 格式验证
		isFormatValid := cve.IsCve(input)
		
		// 格式化
		formatted := cve.Format(input)
		
		// 年份验证
		isYearValid := cve.IsCveYearOk(formatted)
		
		// 完整验证
		isFullyValid := cve.ValidateCve(formatted)
		
		// 打印结果
		fmt.Printf("%-20s %-10t %-10t %-10t", input, isFormatValid, isYearValid, isFullyValid)
		
		// 如果格式有效，提取更多信息
		if isFormatValid {
			year := cve.ExtractCveYearAsInt(formatted)
			yearDiff := year - currentYear
			
			if yearDiff > 0 {
				fmt.Printf(" (年份超前%d年)", yearDiff)
			} else if yearDiff < -20 {
				fmt.Printf(" (漏洞较老，%d年前)", -yearDiff)
			}
		}
		
		fmt.Println()
	}
	
	// 从文本中提取CVE
	text := "本系统受到CVE-2021-44228和CVE-2022-22965漏洞的影响。"
	extractedCves := cve.ExtractCve(text)
	
	fmt.Println("\n从文本中提取的CVE:")
	for i, id := range extractedCves {
		fmt.Printf("%d. %s\n", i+1, id)
	}
}
```

### 数据分析与处理管道

```go
package main

import (
	"fmt"
	"github.com/scagogogo/cve"
	"sort"
	"time"
)

func main() {
	// 模拟从多个来源收集的CVE数据
	source1 := []string{"CVE-2022-1111", "CVE-2021-2222", "CVE-2020-3333"}
	source2 := []string{"cve-2022-1111", "CVE-2022-4444", "CVE-2019-5555"}
	source3 := []string{"CVE-2021-2222", "CVE-2023-6666"}
	
	// 1. 合并所有来源
	allCves := append(source1, source2...)
	allCves = append(allCves, source3...)
	fmt.Printf("合并前总数: %d\n", len(allCves))
	
	// 2. 去重
	uniqueCves := cve.RemoveDuplicateCves(allCves)
	fmt.Printf("去重后总数: %d\n", len(uniqueCves))
	
	// 3. 排序
	sortedCves := cve.SortCves(uniqueCves)
	
	// 4. 按年份分组
	groupedCves := cve.GroupByYear(sortedCves)
	
	// 获取年份列表并排序
	years := make([]string, 0, len(groupedCves))
	for year := range groupedCves {
		years = append(years, year)
	}
	sort.Strings(years)
	
	// 5. 按年份显示统计信息
	fmt.Println("\n按年份统计:")
	for _, year := range years {
		cvesInYear := groupedCves[year]
		fmt.Printf("%s年: %d个CVE\n", year, len(cvesInYear))
	}
	
	// 6. 仅显示最近两年的详细CVE列表
	currentYear := time.Now().Year()
	recentCves := cve.GetRecentCves(sortedCves, 2)
	
	fmt.Printf("\n最近两年CVE详情 (%d-%d):\n", currentYear-1, currentYear)
	for i, id := range recentCves {
		year := cve.ExtractCveYear(id)
		seq := cve.ExtractCveSeq(id)
		fmt.Printf("%d. %s (年份: %s, 序列号: %s)\n", i+1, id, year, seq)
	}
	
	// 7. 验证所有CVE的有效性
	invalidCount := 0
	for _, id := range sortedCves {
		if !cve.ValidateCve(id) {
			invalidCount++
			fmt.Printf("无效的CVE: %s\n", id)
		}
	}
	
	if invalidCount == 0 {
		fmt.Println("\n所有CVE格式验证通过")
	} else {
		fmt.Printf("\n发现 %d 个无效格式的CVE\n", invalidCount)
	}
}
```

## ⭐ Star History

[![Star History Chart](https://api.star-history.com/svg?repos=scagogogo/cve&type=Date)](https://star-history.com/#scagogogo/cve&Date)

## 🤝 贡献指南

我们非常欢迎您的贡献！如果您有兴趣改进这个项目，请参考以下步骤：

1. Fork 这个仓库
2. 创建您的特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交您的更改 (`git commit -m 'feat: 添加了一些很棒的功能'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建一个 Pull Request

任何形式的贡献都将被感激，无论是新功能、文档改进、bug修复还是性能优化。

### 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 简体中文

CVE (Common Vulnerabilities and Exposures) 相关的工具方法集合。这个包提供了一系列用于处理、验证、提取和操作 CVE 标识符的实用函数。

### 📖 文档

**完整的 API 文档和使用指南请访问：[https://scagogogo.github.io/cve/zh/](https://scagogogo.github.io/cve/zh/)**

文档包含：
- 🚀 [快速开始指南](https://scagogogo.github.io/cve/zh/guide/getting-started)
- 📚 [完整 API 参考](https://scagogogo.github.io/cve/zh/api/)
- 💡 [实际使用示例](https://scagogogo.github.io/cve/zh/examples/)
- 🔧 [安装和配置](https://scagogogo.github.io/cve/zh/guide/installation)

### 📑 目录

- [功能特性](#功能特性-1)
- [安装](#安装-1)
- [快速开始](#快速开始-1)
- [API 参考文档](#api-参考文档)
  - [格式化与验证](#格式化与验证)
  - [提取方法](#提取方法)
  - [比较与排序](#比较与排序)
  - [过滤与分组](#过滤与分组)
  - [生成与构造](#生成与构造)
- [使用场景示例](#使用场景示例)
- [项目结构](#项目结构)
- [参考资料](#参考资料)
- [许可证](#许可证-1)

### ✨ 功能特性

- ✅ CVE 格式验证和标准化
- ✅ 从文本中提取 CVE 标识符
- ✅ CVE 的年份和序列号提取与比较
- ✅ CVE 的排序、过滤和分组
- ✅ 生成标准格式的 CVE 标识符
- ✅ 去重和验证工具

### 📦 安装

```bash
go get github.com/scagogogo/cve
```

### 🚦 快速开始

```go
package main

import (
    "fmt"
    "github.com/scagogogo/cve"
)

func main() {
    // 格式化 CVE
    formatted := cve.Format("cve-2022-12345")
    fmt.Println(formatted) // 输出: CVE-2022-12345

    // 验证 CVE
    isValid := cve.ValidateCve("CVE-2022-12345")
    fmt.Println(isValid) // 输出: true

    // 从文本中提取 CVE
    text := "系统受到 CVE-2021-44228 和 CVE-2022-12345 的影响"
    cves := cve.ExtractCve(text)
    fmt.Println(cves) // 输出: [CVE-2021-44228 CVE-2022-12345]
}
```

### 📄 许可证

本项目采用 MIT 协议开源，详情请参阅 [LICENSE](LICENSE) 文件。