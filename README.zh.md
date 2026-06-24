# CVE Utils

[![Go Tests](https://github.com/scagogogo/cve/actions/workflows/go-test.yml/badge.svg)](https://github.com/scagogogo/cve/actions/workflows/go-test.yml)
[![Documentation](https://github.com/scagogogo/cve/actions/workflows/docs.yml/badge.svg)](https://github.com/scagogogo/cve/actions/workflows/docs.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/scagogogo/cve.svg)](https://pkg.go.dev/github.com/scagogogo/cve)
[![Go Report Card](https://goreportcard.com/badge/github.com/scagogogo/cve)](https://goreportcard.com/report/github.com/scagogogo/cve)
[![License](https://img.shields.io/github/license/scagogogo/cve)](https://github.com/scagogogo/cve/blob/main/LICENSE)
[![Version](https://img.shields.io/badge/version-v0.0.1-blue)](https://github.com/scagogogo/cve/releases)

**🌐 Languages: [English](README.md) | [简体中文](README.zh.md)**

CVE (Common Vulnerabilities and Exposures) 相关的工具方法集合。这个包提供了一系列用于处理、验证、提取和操作 CVE 标识符的实用函数。

## 📖 文档

**完整的 API 文档和使用指南请访问：[https://scagogogo.github.io/cve/zh/](https://scagogogo.github.io/cve/zh/)**

文档包含：
- 🚀 [快速开始指南](https://scagogogo.github.io/cve/zh/guide/getting-started)
- 📚 [完整 API 参考](https://scagogogo.github.io/cve/zh/api/)
- 💡 [实际使用示例](https://scagogogo.github.io/cve/zh/examples/)
- 🔧 [安装和配置](https://scagogogo.github.io/cve/zh/guide/installation)

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
  - [集合运算](#集合运算)
  - [批量验证](#批量验证)
  - [范围与模式匹配](#范围与模式匹配)
  - [统计分析](#统计分析)
- [使用场景示例](#-使用场景示例)
- [项目结构](#-项目结构)
- [参考资料](#-参考资料)
- [许可证](#-许可证)

## ✨ 功能特性

- ✅ CVE 格式验证和标准化
- ✅ 从文本中提取 CVE 标识符
- ✅ CVE 的年份和序列号提取与比较
- ✅ CVE 的排序、过滤和分组
- ✅ 生成标准格式的 CVE 标识符
- ✅ 去重和验证工具
- ✅ 集合运算（交集、并集、差集）
- ✅ 批量验证及详细错误报告
- ✅ CVE 范围解析（支持 `to`、`..`、`-` 语法）
- ✅ 统计分析（按年计数、年份范围、序列号范围）
- ✅ 通配符模式匹配实现灵活筛选
- ✅ 序列号格式化（前补零）

## 📦 安装

```bash
go get github.com/scagogogo/cve
```

## 🚦 快速开始

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

## 📚 API 参考文档

### 格式化与验证

| 函数 | 描述 |
|------|------|
| `Format(cve string) string` | 将 CVE 转换为标准大写格式 |
| `IsCve(text string) bool` | 判断字符串是否为有效的 CVE 格式 |
| `IsContainsCve(text string) bool` | 判断字符串是否包含 CVE |
| `ValidateCve(cve string) bool` | 全面验证 CVE 编号的合法性 |

### 提取方法

| 函数 | 描述 |
|------|------|
| `ExtractCve(text string) []string` | 从文本中提取所有 CVE 编号 |
| `ExtractFirstCve(text string) string` | 提取第一个 CVE 编号 |
| `ExtractLastCve(text string) string` | 提取最后一个 CVE 编号 |
| `Split(cve string) (year string, seq string)` | 分割 CVE 为年份和序列号 |

### 比较与排序

| 函数 | 描述 |
|------|------|
| `CompareCves(cveA, cveB string) int` | 全面比较两个 CVE |
| `SortCves(cveSlice []string) []string` | 对 CVE 切片进行排序 |
| `CompareByYear(cveA, cveB string) int` | 根据年份比较两个 CVE |

### 过滤与分组

| 函数 | 描述 |
|------|------|
| `FilterCvesByYear(cveSlice []string, year int) []string` | 筛选特定年份的 CVE |
| `GroupByYear(cveSlice []string) map[string][]string` | 按年份分组 CVE |
| `RemoveDuplicateCves(cveSlice []string) []string` | 移除重复的 CVE |

### 生成与构造

| 函数 | 描述 |
|------|------|
| `GenerateCve(year int, seq int) string` | 根据年份和序列号生成 CVE |
| `GenerateFakeCve() string` | 生成一个随机的假 CVE 测试用 |
| `FormatSeq(cve string, width int) string` | 将 CVE 序列号格式化为指定位数（前补零） |

### 集合运算

| 函数 | 描述 |
|------|------|
| `IntersectCves(a, b []string) []string` | 求两个 CVE 列表的交集 |
| `UnionCves(a, b []string) []string` | 求两个 CVE 列表的并集 |
| `DiffCves(a, b []string) []string` | 求两个 CVE 列表的差集（a 有 b 无） |

### 批量验证

| 函数 | 描述 |
|------|------|
| `ValidateCves(cveSlice []string) []CveValidationResult` | 批量验证 CVE 并返回详细错误原因 |
| `FilterValidCves(cveSlice []string) []string` | 从列表中过滤出有效的 CVE |

### 范围与模式匹配

| 函数 | 描述 |
|------|------|
| `ParseCveRange(rangeExpr string) []string` | 解析 CVE 范围表达式（支持 `to`、`..`、`-`） |
| `IsCvesConsecutive(a, b string) bool` | 判断两个 CVE 是否连续 |
| `FilterCvesByPattern(cveSlice []string, pattern string) []string` | 通过通配符模式筛选 CVE（如 `CVE-2022-*`） |

### 统计分析

| 函数 | 描述 |
|------|------|
| `CountByYear(cveSlice []string) map[int]int` | 按年份统计 CVE 数量 |
| `YearRange(cveSlice []string) (min, max int)` | 获取 CVE 列表中最早和最晚的年份 |
| `SeqRange(cveSlice []string, year int) (min, max int)` | 获取指定年份下 CVE 序列号的范围 |

## 💡 使用场景示例

### 基本验证

```go
// 验证用户输入
func validateUserInput(input string) bool {
    return cve.ValidateCve(input)
}
```

### 文本处理

```go
// 从安全公告中提取 CVE
func extractFromBulletin(bulletin string) []string {
    return cve.ExtractCve(bulletin)
}
```

### 数据清洗

```go
// 清洗和排序 CVE 列表
func cleanCveList(rawList []string) []string {
    unique := cve.RemoveDuplicateCves(rawList)
    return cve.SortCves(unique)
}
```

## 🏗️ 项目结构

```
cve/
├── base.go              # 格式化、验证、批量验证
├── extract.go           # 提取方法、模式匹配
├── compare.go           # 比较与排序
├── filter.go            # 过滤分组、集合运算、统计分析
├── generate.go          # 生成构造、范围解析
├── *_test.go            # 单元测试（覆盖率 95%+）
├── README.md            # 英文文档
├── README.zh.md         # 中文文档
├── LICENSE              # 许可证
└── docs/                # 文档网站
    ├── index.md        # 英文首页
    ├── zh/             # 中文文档
    ├── api/            # API 文档
    ├── guide/          # 使用指南
    └── examples/       # 使用示例
```

## 📖 参考资料

- [CVE 官方网站](https://cve.mitre.org/)
- [CVE 编号规范](https://cve.mitre.org/cve/identifiers/)
- [Go 语言官方文档](https://golang.org/doc/)

## 📄 许可证

本项目采用 MIT 协议开源，详情请参阅 [LICENSE](LICENSE) 文件。
