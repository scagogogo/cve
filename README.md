# CVE Utils

[![Go Tests](https://github.com/scagogogo/cve-skills/actions/workflows/go-test.yml/badge.svg)](https://github.com/scagogogo/cve-skills/actions/workflows/go-test.yml)
[![Documentation](https://github.com/scagogogo/cve-skills/actions/workflows/docs.yml/badge.svg)](https://github.com/scagogogo/cve-skills/actions/workflows/docs.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/scagogogo/cve-skills.svg)](https://pkg.go.dev/github.com/scagogogo/cve-skills)
[![Go Report Card](https://goreportcard.com/badge/github.com/scagogogo/cve-skills)](https://goreportcard.com/report/github.com/scagogogo/cve-skills)
[![License](https://img.shields.io/github/license/scagogogo/cve)](https://github.com/scagogogo/cve-skills/blob/main/LICENSE)
[![Version](https://img.shields.io/badge/version-v0.0.1-blue)](https://github.com/scagogogo/cve-skills/releases)

**🌐 Languages: [English](README.md) | [简体中文](README.zh.md)**

---

## What is CVE Utils?

**CVE Utils** is a comprehensive Go library and CLI tool for processing CVE (Common Vulnerabilities and Exposures) identifiers. It provides 30+ utility functions covering everything from basic format validation to advanced set operations and statistical analysis.

## The Problem It Solves

When working with CVE identifiers in security tools, vulnerability scanners, and compliance systems, developers repeatedly face these challenges:

- **Format inconsistency** — CVE IDs appear as `cve-2022-12345`, `CVE-2022-12345`, `CVE-2022-012345`, or even mixed in text. Standardization is tedious but essential.
- **Manual extraction** — Parsing CVE IDs from security advisories, NVD feeds, and vulnerability reports requires custom regex logic every time.
- **No native comparison** — Go has no built-in way to compare, sort, or filter CVE identifiers by year or sequence number.
- **Duplicate handling** — Merging CVE lists from multiple sources creates duplicates and format mismatches.
- **Range parsing** — Security bulletins often describe CVE ranges (`CVE-2022-1000 to CVE-2022-1050`), which must be expanded manually.
- **Repetitive validation** — Every project re-implements CVE validation with slightly different rules, leading to inconsistencies.

**CVE Utils eliminates all of these problems** with a single, well-tested dependency.

## Feature Map

![Feature Map](docs/images/feature-map.png)

## Architecture

![Architecture](docs/images/architecture.png)

## CLI Command Tree

![CLI Command Tree](docs/images/cli-tree.png)

## Features at a Glance

| Category | Functions | Highlights |
|----------|-----------|------------|
| Format & Validation | 7 | Standardize, validate, year-check with cutoff |
| Extraction | 8 | Parse from text, split year/seq, int variants |
| Compare & Sorting | 4 | Full comparison, sort, year diff |
| Filter & Grouping | 5 | By year, year range, recent, dedup |
| Generation | 3 | Generate, fake, zero-pad seq |
| Set Operations | 3 | Intersection, union, difference |
| Batch Validation | 2 | Batch validate with reasons, filter valid |
| Range & Pattern | 3 | Parse ranges, check consecutive, wildcard |
| Statistics | 3 | Count by year, year range, seq range |

## Installation

### As a Go Library

```bash
go get github.com/scagogogo/cve-skills
```

### As a CLI Tool

```bash
go install github.com/scagogogo/cve-skills/cmd/cve@latest
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/scagogogo/cve-skills"
)

func main() {
    // 1. Format & Validate
    formatted := cve.Format("cve-2022-12345")
    fmt.Println(formatted) // CVE-2022-12345

    isValid := cve.ValidateCve("CVE-2022-12345")
    fmt.Println(isValid) // true

    // 2. Extract from text
    text := "Affected by CVE-2021-44228 and CVE-2022-12345"
    cves := cve.ExtractCve(text)
    fmt.Println(cves) // [CVE-2021-44228 CVE-2022-12345]

    // 3. Sort & Filter
    list := []string{"CVE-2022-3333", "CVE-2020-1111", "CVE-2022-1111"}
    sorted := cve.SortCves(list)
    fmt.Println(sorted) // [CVE-2020-1111 CVE-2022-1111 CVE-2022-3333]

    // 4. Set operations
    common := cve.IntersectCves(
        []string{"CVE-2022-1111", "CVE-2022-2222"},
        []string{"CVE-2022-2222", "CVE-2022-3333"},
    )
    fmt.Println(common) // [CVE-2022-2222]
}
```

## CLI Usage

```bash
# Format CVE identifiers
cve format CVE-2022-12345 cve-2023-54321

# Validate
cve validate CVE-2022-12345 CVE-1998-12345

# Extract from text
cve extract "System affected by CVE-2021-44228 and CVE-2022-12345"

# Compare
cve compare CVE-2021-44228 CVE-2022-12345

# Sort
cve sort CVE-2022-3333 CVE-2020-1111 CVE-2022-1111

# Filter by year
cve filter by-year --year 2022 CVE-2021-1111 CVE-2022-2222 CVE-2023-3333

# Generate
cve generate cve --year 2024 --seq 56789
cve generate fake

# Set operations
cve intersect "CVE-2022-1111,CVE-2022-2222" "CVE-2022-2222,CVE-2022-3333"

# Parse CVE range
cve parse-range "CVE-2022-1000 to CVE-2022-1005"

# Statistics
cve count-by-year "CVE-2022-1111,CVE-2022-2222,CVE-2021-3333"
cve year-range "CVE-2020-1111,CVE-2023-9999"
```

## API Reference

### Format & Validation

| Function | Description |
|----------|-------------|
| `Format(cve string) string` | Convert to standard uppercase format |
| `IsCve(text string) bool` | Check if string is a valid CVE format |
| `IsContainsCve(text string) bool` | Check if text contains a CVE |
| `ValidateCve(cve string) bool` | Comprehensive validation (format + year + seq) |
| `IsCveYearOk(cve string) bool` | Check if year is in 1999–current year |
| `IsCveYearOkWithCutoff(cve string, cutoff int) bool` | Year check with future-year offset |
| `FormatSeq(cve string, width int) string` | Zero-pad sequence number to fixed width |

### Extraction

| Function | Description |
|----------|-------------|
| `ExtractCve(text string) []string` | Extract all CVEs from text |
| `ExtractFirstCve(text string) string` | Extract the first CVE |
| `ExtractLastCve(text string) string` | Extract the last CVE |
| `Split(cve string) (year, seq string)` | Split CVE into year and sequence |
| `ExtractCveYear(cve string) string` | Extract year as string |
| `ExtractCveYearAsInt(cve string) int` | Extract year as integer |
| `ExtractCveSeq(cve string) string` | Extract sequence as string |
| `ExtractCveSeqAsInt(cve string) int` | Extract sequence as integer |

### Comparison & Sorting

| Function | Description |
|----------|-------------|
| `CompareCves(cveA, cveB string) int` | Full comparison (year, then sequence) |
| `CompareByYear(cveA, cveB string) int` | Compare by year only |
| `SubByYear(cveA, cveB string) int` | Year difference between two CVEs |
| `SortCves(cveSlice []string) []string` | Sort by year and sequence |

### Filtering & Grouping

| Function | Description |
|----------|-------------|
| `FilterCvesByYear(cveSlice []string, year int) []string` | Filter by specific year |
| `FilterCvesByYearRange(cveSlice []string, start, end int) []string` | Filter by year range |
| `GetRecentCves(cveSlice []string, years int) []string` | Get CVEs from last N years |
| `GroupByYear(cveSlice []string) map[string][]string` | Group CVEs by year |
| `RemoveDuplicateCves(cveSlice []string) []string` | Remove duplicates (case-insensitive) |

### Generation & Construction

| Function | Description |
|----------|-------------|
| `GenerateCve(year, seq int) string` | Generate CVE from year and sequence |
| `GenerateFakeCve() string` | Generate random CVE for testing |
| `FormatSeq(cve string, width int) string` | Format sequence number with zero-padding |

### Set Operations

| Function | Description |
|----------|-------------|
| `IntersectCves(a, b []string) []string` | Intersection of two CVE lists |
| `UnionCves(a, b []string) []string` | Union of two CVE lists |
| `DiffCves(a, b []string) []string` | Difference (a - b) of two CVE lists |

### Batch Validation

| Function | Description |
|----------|-------------|
| `ValidateCves(cveSlice []string) []CveValidationResult` | Batch validate with error reasons |
| `FilterValidCves(cveSlice []string) []string` | Filter out only valid CVEs |

### Range & Pattern Matching

| Function | Description |
|----------|-------------|
| `ParseCveRange(rangeExpr string) []string` | Parse range expression (`to`, `..`, `-`) |
| `IsCvesConsecutive(a, b string) bool` | Check if two CVEs are consecutive |
| `FilterCvesByPattern(cveSlice []string, pattern string) []string` | Wildcard pattern filter |

### Statistical Analysis

| Function | Description |
|----------|-------------|
| `CountByYear(cveSlice []string) map[int]int` | Count CVEs per year |
| `YearRange(cveSlice []string) (min, max int)` | Earliest and latest year |
| `SeqRange(cveSlice []string, year int) (min, max int)` | Sequence range for a given year |

## Real-World Use Cases

### Security Advisory Parser

```go
// Extract and normalize CVEs from a security advisory
func parseAdvisory(advisory string) []string {
    raw := cve.ExtractCve(advisory)
    unique := cve.RemoveDuplicateCves(raw)
    return cve.SortCves(unique)
}
```

### Vulnerability Dashboard Data

```go
// Generate year-by-year statistics for a dashboard
func dashboardStats(cveList []string) {
    counts := cve.CountByYear(cveList)
    minYear, maxYear := cve.YearRange(cveList)
    fmt.Printf("CVEs span %d to %d\n", minYear, maxYear)
    for year, count := range counts {
        fmt.Printf("  %d: %d vulnerabilities\n", year, count)
    }
}
```

### Compliance Report Generator

```go
// Find new CVEs not in last year's report
func findNewCves(current, historical []string) []string {
    return cve.DiffCves(current, historical)
}
```

### CVE Range Expansion

```go
// Expand "CVE-2022-1000 to CVE-2022-1050" into individual CVEs
func expandRange(rangeExpr string) []string {
    return cve.ParseCveRange(rangeExpr)
}
```

## Documentation

**Full API docs and guides: [https://scagogogo.github.io/cve-skills/](https://scagogogo.github.io/cve-skills/)**

- [Quick Start Guide](https://scagogogo.github.io/cve-skills/guide/getting-started)
- [Complete API Reference](https://scagogogo.github.io/cve-skills/api/)
- [Practical Examples](https://scagogogo.github.io/cve-skills/examples/)
- [Installation & Configuration](https://scagogogo.github.io/cve-skills/guide/installation)

## Project Structure

```
cve/
├── cve.go              # Package info & version
├── base.go             # Format, validation, batch validation
├── extract.go          # Extraction, pattern matching
├── compare.go          # Comparison and sorting
├── filter.go           # Filtering, grouping, set ops, statistics
├── generate.go         # Generation, range parsing
├── *_test.go           # Unit tests (95%+ coverage)
├── cmd/                # CLI implementation
│   ├── root.go         # Root command
│   ├── format.go       # Format subcommand
│   ├── validate.go     # Validate subcommands
│   ├── extract.go      # Extract subcommands
│   ├── compare.go      # Compare & sort subcommands
│   ├── filter.go       # Filter & group subcommands
│   ├── generate.go     # Generate subcommands
│   ├── set.go          # Set operation subcommands
│   ├── range.go        # Range & pattern subcommands
│   ├── stats.go        # Statistics subcommands
│   └── ...
├── examples/           # 30+ runnable examples
├── docs/               # VitePress documentation site
└── scripts/            # Image generation scripts
```

## References

- [CVE Official Website](https://cve.mitre.org/)
- [CVE Identifier Specification](https://cve.mitre.org/cve/identifiers/)
- [Go Documentation](https://golang.org/doc/)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
