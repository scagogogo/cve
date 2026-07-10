---
layout: home

hero:
  name: "CVE Utils"
  text: "AI First CVE Toolkit"
  tagline: 30+ Go functions + cross-platform CLI for CVE identifier processing — designed to be read, installed, and driven by AI agents.
  image:
    src: /hero.svg
    alt: CVE Utils
  actions:
    - theme: brand
      text: Quick Start
      link: /guide/getting-started
    - theme: alt
      text: API Reference
      link: /api/
    - theme: alt
      text: GitHub
      link: https://github.com/scagogogo/cve-skills

features:
  - icon: 🤖
    title: AI First
    details: Machine-readable API surface, deterministic CLI output, and a single dependency. Agents install, call, and parse results without guessing.
    link: /guide/getting-started
  - icon: 🔍
    title: Format & Validation
    details: Standardize, validate, year-check with cutoff. 7 functions covering the full format lifecycle.
    link: /api/format-validate
  - icon: 📝
    title: Smart Extraction
    details: Parse CVE IDs from any text — advisories, NVD feeds, reports. First/last/batch extraction.
    link: /api/extract
  - icon: 🔄
    title: Compare & Sort
    details: Native comparison and sorting by year and sequence number. No more custom regex logic.
    link: /api/compare-sort
  - icon: 🎯
    title: Filter & Group
    details: By year, year range, recent N years, dedup. 5 functions for clean vulnerability inventories.
    link: /api/filter-group
  - icon: 🛠️
    title: Set Operations
    details: Intersect, union, diff across CVE lists — find new, common, and removed vulnerabilities.
    link: /api/set-operations
  - icon: ⚡
    title: Range & Pattern
    details: Parse CVE ranges (CVE-2022-1000 ~ CVE-2022-1050), check continuity, wildcard matching.
    link: /api/range-pattern
  - icon: 📊
    title: Statistics
    details: Count by year, year range, sequence range. Turn raw CVE data into trend insights.
    link: /api/statistics
---

## How an AI Agent uses it

```mermaid
flowchart LR
    A["Agent receives<br/>advisory text"] --> B["ExtractCve()"]
    B --> C["ValidateCves()"]
    C --> D["SortCves() +<br/>RemoveDuplicateCves()"]
    D --> E["DiffCves()<br/>vs known list"]
    E --> F["New vulnerabilities<br/>for triage"]
```

::: code-group

```bash [Install CLI]
# Prebuilt binary for every major platform
go install github.com/scagogogo/cve-skills/cmd/cve@latest
# or download from Releases:
# https://github.com/scagogogo/cve-skills/releases
```

```bash [Use the CLI]
# Extract & validate CVEs from any text — deterministic stdout
echo "Affected by CVE-2021-44228 and CVE-2022-12345" | cve extract
# 34 subcommands, all args-or-stdin — see the CLI Reference: /cli
```

```go [Use as a Go library]
package main

import (
    "fmt"
    "github.com/scagogogo/cve-skills" // go get github.com/scagogogo/cve-skills
)

func main() {
    text := "System affected by CVE-2021-44228 and CVE-2022-12345"

    // Extract, validate, dedup, sort — in one pipeline
    cves := cve.SortCves(cve.RemoveDuplicateCves(cve.ExtractCve(text)))
    fmt.Println(cves) // [CVE-2021-44228 CVE-2022-12345]

    fmt.Println(cve.ValidateCve("CVE-2022-12345")) // true
}
```

:::

## Function Map

30+ functions across 9 categories. One dependency, zero CVE-format reinvention.

```mermaid
graph TD
    ROOT["cve package"] --> F1["Format & Validation<br/>7 funcs"]
    ROOT --> F2["Extraction<br/>8 funcs"]
    ROOT --> F3["Compare & Sort<br/>4 funcs"]
    ROOT --> F4["Filter & Group<br/>5 funcs"]
    ROOT --> F5["Generate<br/>2 funcs"]
    ROOT --> F6["Set Operations<br/>3 funcs"]
    ROOT --> F7["Batch Validation<br/>2 funcs"]
    ROOT --> F8["Range & Pattern<br/>3 funcs"]
    ROOT --> F9["Statistics<br/>3 funcs"]
```

## Why CVE Utils?

| Problem | CVE Utils solves it |
|---------|---------------------|
| Format inconsistency (`cve-...`, `CVE-...`, mixed case) | `Format()` → canonical `CVE-YYYY-NNNNN` |
| Manual extraction with custom regex | `ExtractCve()` from any text |
| No native Go comparison/sort | `CompareCves()` / `SortCves()` |
| Duplicates when merging sources | `RemoveDuplicateCves()` |
| Range parsing in bulletins | `ParseCveRange()` → expanded list |

🚀 **[Get started in 30 seconds →](/guide/getting-started)**
