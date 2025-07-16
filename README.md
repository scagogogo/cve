# CVE Utils

[![Go Tests](https://github.com/scagogogo/cve/actions/workflows/go-test.yml/badge.svg)](https://github.com/scagogogo/cve/actions/workflows/go-test.yml)
[![Documentation](https://github.com/scagogogo/cve/actions/workflows/docs.yml/badge.svg)](https://github.com/scagogogo/cve/actions/workflows/docs.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/scagogogo/cve.svg)](https://pkg.go.dev/github.com/scagogogo/cve)
[![Go Report Card](https://goreportcard.com/badge/github.com/scagogogo/cve)](https://goreportcard.com/report/github.com/scagogogo/cve)
[![License](https://img.shields.io/github/license/scagogogo/cve)](https://github.com/scagogogo/cve/blob/main/LICENSE)
[![Version](https://img.shields.io/badge/version-v0.0.1-blue)](https://github.com/scagogogo/cve/releases)

**🌐 Languages: [English](README.md) | [简体中文](README.zh.md)**

A comprehensive collection of utility functions for handling CVE (Common Vulnerabilities and Exposures) identifiers. This package provides a series of practical functions for processing, validating, extracting, and manipulating CVE identifiers.

## 📖 Documentation

**Complete API documentation and usage guides: [https://scagogogo.github.io/cve/](https://scagogogo.github.io/cve/)**

Documentation includes:
- 🚀 [Quick Start Guide](https://scagogogo.github.io/cve/guide/getting-started)
- 📚 [Complete API Reference](https://scagogogo.github.io/cve/api/)
- 💡 [Practical Examples](https://scagogogo.github.io/cve/examples/)
- 🔧 [Installation & Configuration](https://scagogogo.github.io/cve/guide/installation)

## 📑 Table of Contents

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

## ✨ Features

- ✅ CVE format validation and standardization
- ✅ Extract CVE identifiers from text
- ✅ Extract and compare CVE years and sequence numbers
- ✅ Sort, filter, and group CVEs
- ✅ Generate standard format CVE identifiers
- ✅ Deduplication and validation tools

## 📦 Installation

```bash
go get github.com/scagogogo/cve
```

## 🚦 Quick Start

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

## 📚 API Reference

### Format & Validation

| Function | Description |
|----------|-------------|
| `Format(cve string) string` | Convert CVE to standard uppercase format |
| `IsCve(text string) bool` | Check if string is a valid CVE format |
| `IsContainsCve(text string) bool` | Check if string contains CVE |
| `ValidateCve(cve string) bool` | Comprehensive validation of CVE identifier |

### Extraction Methods

| Function | Description |
|----------|-------------|
| `ExtractCve(text string) []string` | Extract all CVE identifiers from text |
| `ExtractFirstCve(text string) string` | Extract the first CVE identifier |
| `ExtractLastCve(text string) string` | Extract the last CVE identifier |
| `Split(cve string) (year string, seq string)` | Split CVE into year and sequence |

### Comparison & Sorting

| Function | Description |
|----------|-------------|
| `CompareCves(cveA, cveB string) int` | Comprehensive comparison of two CVEs |
| `SortCves(cveSlice []string) []string` | Sort CVE slice |
| `CompareByYear(cveA, cveB string) int` | Compare two CVEs by year |

### Filtering & Grouping

| Function | Description |
|----------|-------------|
| `FilterCvesByYear(cveSlice []string, year int) []string` | Filter CVEs by specific year |
| `GroupByYear(cveSlice []string) map[string][]string` | Group CVEs by year |
| `RemoveDuplicateCves(cveSlice []string) []string` | Remove duplicate CVEs |

### Generation & Construction

| Function | Description |
|----------|-------------|
| `GenerateCve(year int, seq int) string` | Generate CVE from year and sequence |

## 💡 Usage Examples

### Basic Validation

```go
// Validate user input
func validateUserInput(input string) bool {
    return cve.ValidateCve(input)
}
```

### Text Processing

```go
// Extract CVEs from security bulletin
func extractFromBulletin(bulletin string) []string {
    return cve.ExtractCve(bulletin)
}
```

### Data Cleaning

```go
// Clean and sort CVE list
func cleanCveList(rawList []string) []string {
    unique := cve.RemoveDuplicateCves(rawList)
    return cve.SortCves(unique)
}
```

## 🏗️ Project Structure

```
cve/
├── cve.go              # Main functionality
├── cve_test.go         # Unit tests
├── README.md           # English documentation
├── README.zh.md        # Chinese documentation
├── LICENSE             # License file
└── docs/               # Documentation website
    ├── index.md        # English homepage
    ├── zh/             # Chinese documentation
    ├── api/            # API documentation
    ├── guide/          # Usage guides
    └── examples/       # Usage examples
```

## 📖 References

- [CVE Official Website](https://cve.mitre.org/)
- [CVE Identifier Specification](https://cve.mitre.org/cve/identifiers/)
- [Go Language Documentation](https://golang.org/doc/)

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
