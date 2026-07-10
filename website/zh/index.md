---
layout: home

hero:
  name: "CVE Utils"
  text: "AI First 的 CVE 工具集"
  tagline: 30+ Go 函数 + 跨平台 CLI，专为 AI Agent 读取、安装与驱动而设计。
  image:
    src: /hero.svg
    alt: CVE Utils
  actions:
    - theme: brand
      text: 快速开始
      link: /zh/guide/getting-started
    - theme: alt
      text: API 文档
      link: /zh/api/
    - theme: alt
      text: GitHub
      link: https://github.com/scagogogo/cve-skills

features:
  - icon: 🤖
    title: AI First
    details: 机器可读的 API 表面、确定性的 CLI 输出、单一依赖。Agent 无需猜测即可安装、调用并解析结果。
    link: /zh/guide/getting-started
  - icon: 🔍
    title: 格式化与验证
    details: 标准化、验证、含截断年份校验。7 个函数覆盖完整格式生命周期。
    link: /zh/api/format-validate
  - icon: 📝
    title: 智能提取
    details: 从任意文本解析 CVE 编号 —— 安全公告、NVD 数据源、报告。支持首/尾/批量提取。
    link: /zh/api/extract
  - icon: 🔄
    title: 比较与排序
    details: 按年份和序列号原生比较与排序，告别自定义正则。
    link: /zh/api/compare-sort
  - icon: 🎯
    title: 过滤与分组
    details: 按年份、年份范围、最近 N 年过滤，自动去重。5 个函数打造干净漏洞清单。
    link: /zh/api/filter-group
  - icon: 🛠️
    title: 集合运算
    details: 交集、并集、差集 —— 快速发现新增、共同与已移除的漏洞。
    link: /zh/api/set-operations
  - icon: ⚡
    title: 范围与模式
    details: 解析 CVE 范围（CVE-2022-1000 ~ CVE-2022-1050）、连续性检查、通配符匹配。
    link: /zh/api/range-pattern
  - icon: 📊
    title: 统计分析
    details: 按年份计数、年份范围、序列号范围。把原始 CVE 数据转化为趋势洞察。
    link: /zh/api/statistics
---

## AI Agent 如何使用它

```mermaid
flowchart LR
    A["Agent 收到<br/>安全公告文本"] --> B["ExtractCve()"]
    B --> C["ValidateCves()"]
    C --> D["SortCves() +<br/>RemoveDuplicateCves()"]
    D --> E["DiffCves()<br/>对比已知清单"]
    E --> F["待分诊的<br/>新漏洞"]
```

::: code-group

```bash [安装 CLI]
# 覆盖所有主流平台的预编译二进制
go install github.com/scagogogo/cve-skills/cmd/cve@latest
# 或从 Releases 下载:
# https://github.com/scagogogo/cve-skills/releases
```

```bash [使用 CLI]
# 从任意文本提取并验证 CVE —— 确定性 stdout
echo "Affected by CVE-2021-44228 and CVE-2022-12345" | cve extract
# 34 个子命令,全部支持参数或 stdin —— 详见 CLI 参考: /zh/cli
```

```go [作为 Go 库使用]
package main

import (
    "fmt"
    "github.com/scagogogo/cve-skills" // go get github.com/scagogogo/cve-skills
)

func main() {
    text := "System affected by CVE-2021-44228 and CVE-2022-12345"

    // 提取、验证、去重、排序 —— 一条流水线
    cves := cve.SortCves(cve.RemoveDuplicateCves(cve.ExtractCve(text)))
    fmt.Println(cves) // [CVE-2021-44228 CVE-2022-12345]

    fmt.Println(cve.ValidateCve("CVE-2022-12345")) // true
}
```

:::

## 函数地图

9 大类共 30+ 函数。单一依赖，零 CVE 格式重复造轮子。

```mermaid
graph TD
    ROOT["cve package"] --> F1["格式化与验证<br/>7 函数"]
    ROOT --> F2["提取<br/>8 函数"]
    ROOT --> F3["比较与排序<br/>4 函数"]
    ROOT --> F4["过滤与分组<br/>5 函数"]
    ROOT --> F5["生成<br/>2 函数"]
    ROOT --> F6["集合运算<br/>3 函数"]
    ROOT --> F7["批量验证<br/>2 函数"]
    ROOT --> F8["范围与模式<br/>3 函数"]
    ROOT --> F9["统计<br/>3 函数"]
```

## 为什么选 CVE Utils？

| 问题 | CVE Utils 的解法 |
|------|------------------|
| 格式不一致（`cve-...`、`CVE-...`、大小写混用） | `Format()` → 标准 `CVE-YYYY-NNNNN` |
| 手写正则提取 | `ExtractCve()` 从任意文本提取 |
| Go 无原生比较/排序 | `CompareCves()` / `SortCves()` |
| 多来源合并产生重复 | `RemoveDuplicateCves()` |
| 公告里的范围描述 | `ParseCveRange()` → 展开为列表 |

🚀 **[30 秒上手 →](/zh/guide/getting-started)**
