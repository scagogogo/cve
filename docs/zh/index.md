---
# CVE Utils 文档网站首页
layout: home

hero:
  name: "CVE Utils"
  text: "CVE 工具方法集合"
  tagline: "强大、易用的 CVE (Common Vulnerabilities and Exposures) 处理工具库"
  actions:
    - theme: brand
      text: 快速开始
      link: /guide/getting-started
    - theme: alt
      text: API 文档
      link: /api/
    - theme: alt
      text: 查看 GitHub
      link: https://github.com/scagogogo/cve

features:
  - icon: 🔍
    title: CVE 格式验证
    details: 提供完整的 CVE 格式验证和标准化功能，确保 CVE 编号的正确性和一致性。
  - icon: 📝
    title: 智能提取
    details: 从任意文本中智能提取 CVE 编号，支持多种格式和大小写变化。
  - icon: 🔄
    title: 排序与比较
    details: 按年份和序列号对 CVE 进行排序和比较，便于管理和分析。
  - icon: 🎯
    title: 过滤与分组
    details: 按年份、年份范围等条件过滤 CVE，支持分组和去重操作。
  - icon: ⚡
    title: 高性能
    details: 使用 Go 语言编写，性能优异，适合处理大量 CVE 数据。
  - icon: 🛠️
    title: 易于使用
    details: 简洁的 API 设计，丰富的文档和示例，快速上手。
---

## 快速开始

### 安装

```bash
go get github.com/scagogogo/cve
```

### 基本使用

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

## 主要功能

### 🔍 格式化与验证
- **Format**: 标准化 CVE 格式
- **IsCve**: 验证是否为有效 CVE 格式
- **IsContainsCve**: 检查文本是否包含 CVE
- **ValidateCve**: 全面验证 CVE 合法性

### 📝 提取方法
- **ExtractCve**: 提取所有 CVE 编号
- **ExtractFirstCve**: 提取第一个 CVE
- **ExtractLastCve**: 提取最后一个 CVE
- **Split**: 分割年份和序列号

### 🔄 比较与排序
- **CompareCves**: 比较两个 CVE
- **SortCves**: 对 CVE 列表排序
- **CompareByYear**: 按年份比较

### 🎯 过滤与分组
- **FilterCvesByYear**: 按年份过滤
- **GroupByYear**: 按年份分组
- **RemoveDuplicateCves**: 去除重复项

## 使用场景

- **安全漏洞管理**: 整理和管理企业内部的漏洞清单
- **漏洞报告分析**: 从安全公告中提取和分析 CVE 信息
- **合规性检查**: 验证和标准化 CVE 编号格式
- **数据清洗**: 去重和排序 CVE 数据
- **漏洞趋势分析**: 按时间维度分析漏洞趋势

## 为什么选择 CVE Utils？

- ✅ **完整功能**: 涵盖 CVE 处理的各个方面
- ✅ **高质量代码**: 完整的测试覆盖和文档
- ✅ **性能优异**: Go 语言实现，处理速度快
- ✅ **易于集成**: 简单的 API，无外部依赖
- ✅ **持续维护**: 活跃的开发和社区支持
