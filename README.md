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
  - [Set Operations](#set-operations)
  - [Batch Validation](#batch-validation)
  - [Range & Pattern Matching](#range--pattern-matching)
  - [Statistical Analysis](#statistical-analysis)
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
- ✅ Set operations (intersection, union, difference)
- ✅ Batch validation with detailed error reporting
- ✅ CVE range parsing (supports `to`, `..`, `-` syntax)
- ✅ Statistical analysis (count by year, year range, sequence range)
- ✅ Wildcard pattern matching for flexible filtering
- ✅ Sequence number formatting (zero-padding)

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
| `GenerateFakeCve() string` | Generate a random fake CVE for testing |
| `FormatSeq(cve string, width int) string` | Format CVE sequence number to fixed width with zero-padding |

### Set Operations

| Function | Description |
|----------|-------------|
| `IntersectCves(a, b []string) []string` | Compute intersection of two CVE lists |
| `UnionCves(a, b []string) []string` | Compute union of two CVE lists |
| `DiffCves(a, b []string) []string` | Compute difference (a - b) of two CVE lists |

### Batch Validation

| Function | Description |
|----------|-------------|
| `ValidateCves(cveSlice []string) []CveValidationResult` | Batch validate CVEs with detailed error reasons |
| `FilterValidCves(cveSlice []string) []string` | Filter out only valid CVEs from a list |

### Range & Pattern Matching

| Function | Description |
|----------|-------------|
| `ParseCveRange(rangeExpr string) []string` | Parse CVE range expression (supports `to`, `..`, `-`) |
| `IsCvesConsecutive(a, b string) bool` | Check if two CVEs are consecutive |
| `FilterCvesByPattern(cveSlice []string, pattern string) []string` | Filter CVEs by wildcard pattern (e.g., `CVE-2022-*`) |

### Statistical Analysis

| Function | Description |
|----------|-------------|
| `CountByYear(cveSlice []string) map[int]int` | Count CVEs by year |
| `YearRange(cveSlice []string) (min, max int)` | Get the earliest and latest year of CVEs |
| `SeqRange(cveSlice []string, year int) (min, max int)` | Get sequence number range for a given year |

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
├── base.go              # Format, validation, batch validation
├── extract.go           # Extraction, pattern matching
├── compare.go           # Comparison and sorting
├── filter.go            # Filtering, grouping, set operations, statistics
├── generate.go          # Generation, range parsing
├── *_test.go            # Unit tests (95%+ coverage)
├── README.md            # English documentation
├── README.zh.md         # Chinese documentation
├── LICENSE              # License file
└── docs/                # Documentation website
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
