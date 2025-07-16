---
layout: home

hero:
  name: "CVE Utils"
  text: "CVE Utility Functions"
  tagline: "A powerful and easy-to-use library for handling CVE (Common Vulnerabilities and Exposures) identifiers"
  actions:
    - theme: brand
      text: Get Started
      link: /guide/getting-started
    - theme: alt
      text: API Docs
      link: /api/
    - theme: alt
      text: View on GitHub
      link: https://github.com/scagogogo/cve

features:
  - icon: 🔍
    title: CVE Format Validation
    details: Complete CVE format validation and standardization to ensure correctness and consistency of CVE identifiers.
  - icon: 📝
    title: Smart Extraction
    details: Intelligently extract CVE identifiers from any text, supporting various formats and case variations.
  - icon: 🔄
    title: Sorting & Comparison
    details: Sort and compare CVEs by year and sequence number for easy management and analysis.
  - icon: 🎯
    title: Filtering & Grouping
    details: Filter CVEs by year, year range, and other conditions with support for grouping and deduplication.
  - icon: ⚡
    title: High Performance
    details: Written in Go for excellent performance, suitable for processing large amounts of CVE data.
  - icon: 🛠️
    title: Easy to Use
    details: Clean API design with rich documentation and examples for quick adoption.
---

## Quick Start

### Installation

```bash
go get github.com/scagogogo/cve
```

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/scagogogo/cve"
)

func main() {
    // Format CVE
    formatted := cve.Format("cve-2022-12345")
    fmt.Println(formatted) // Output: CVE-2022-12345

    // Validate CVE
    isValid := cve.ValidateCve("CVE-2022-12345")
    fmt.Println(isValid) // Output: true

    // Extract CVE from text
    text := "System affected by CVE-2021-44228 and CVE-2022-12345"
    cves := cve.ExtractCve(text)
    fmt.Println(cves) // Output: [CVE-2021-44228 CVE-2022-12345]
}
```

## Main Features

### 🔍 Format & Validation
- **Format**: Standardize CVE format
- **IsCve**: Validate if string is a valid CVE format
- **IsContainsCve**: Check if text contains CVE
- **ValidateCve**: Comprehensive CVE validation

### 📝 Extraction Methods
- **ExtractCve**: Extract all CVE identifiers
- **ExtractFirstCve**: Extract the first CVE
- **ExtractLastCve**: Extract the last CVE
- **Split**: Split year and sequence number

### 🔄 Comparison & Sorting
- **CompareCves**: Compare two CVEs
- **SortCves**: Sort CVE list
- **CompareByYear**: Compare by year

### 🎯 Filtering & Grouping
- **FilterCvesByYear**: Filter by year
- **GroupByYear**: Group by year
- **RemoveDuplicateCves**: Remove duplicates

## Use Cases

- **Security Vulnerability Management**: Organize and manage enterprise vulnerability inventories
- **Vulnerability Report Analysis**: Extract and analyze CVE information from security bulletins
- **Compliance Checking**: Validate and standardize CVE identifier formats
- **Data Cleaning**: Deduplicate and sort CVE data
- **Vulnerability Trend Analysis**: Analyze vulnerability trends by time dimension

## Why Choose CVE Utils?

- ✅ **Complete Features**: Covers all aspects of CVE processing
- ✅ **High Quality Code**: Complete test coverage and documentation
- ✅ **Excellent Performance**: Go implementation for fast processing
- ✅ **Easy Integration**: Simple API with no external dependencies
- ✅ **Continuous Maintenance**: Active development and community support
