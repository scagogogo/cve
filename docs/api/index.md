# API Reference

CVE Utils provides a complete set of CVE processing functions, covering everything from basic format validation to complex analysis and processing.

## Function Categories

### 🔍 [Format & Validation](/api/format-validate)

Functions for CVE format standardization and validity verification:

| Function | Description |
|----------|-------------|
| `Format(cve string) string` | Convert CVE to standard uppercase format |
| `IsCve(text string) bool` | Check if string is a valid CVE format |
| `IsContainsCve(text string) bool` | Check if string contains CVE |
| `IsCveYearOk(cve string, cutoff int) bool` | Check if CVE year is reasonable |
| `ValidateCve(cve string) bool` | Comprehensive validation of CVE identifier |

### 📝 [Extraction Methods](/api/extract)

Functions for extracting information from text or CVE:

| Function | Description |
|----------|-------------|
| `ExtractCve(text string) []string` | Extract all CVE identifiers from text |
| `ExtractFirstCve(text string) string` | Extract the first CVE identifier |
| `ExtractLastCve(text string) string` | Extract the last CVE identifier |
| `Split(cve string) (year string, seq string)` | Split CVE into year and sequence |
| `ExtractCveYear(cve string) string` | Extract CVE year (string) |
| `ExtractCveYearAsInt(cve string) int` | Extract CVE year (integer) |
| `ExtractCveSeq(cve string) string` | Extract CVE sequence (string) |
| `ExtractCveSeqAsInt(cve string) int` | Extract CVE sequence (integer) |

### 🔄 [Comparison & Sorting](/api/compare-sort)

Functions for CVE comparison and sorting:

| Function | Description |
|----------|-------------|
| `CompareByYear(cveA, cveB string) int` | Compare two CVEs by year |
| `SubByYear(cveA, cveB string) int` | Calculate year difference between two CVEs |
| `CompareCves(cveA, cveB string) int` | Comprehensive comparison of two CVEs |
| `SortCves(cveSlice []string) []string` | Sort CVE slice |

### 🎯 [Filtering & Grouping](/api/filter-group)

Functions for CVE filtering, grouping, and deduplication:

| Function | Description |
|----------|-------------|
| `FilterCvesByYear(cveSlice []string, year int) []string` | Filter CVEs by specific year |
| `FilterCvesByYearRange(cveSlice []string, startYear, endYear int) []string` | Filter CVEs by year range |
| `GetRecentCves(cveSlice []string, years int) []string` | Get CVEs from recent years |
| `GroupByYear(cveSlice []string) map[string][]string` | Group CVEs by year |
| `RemoveDuplicateCves(cveSlice []string) []string` | Remove duplicate CVEs |

### ⚡ [Generation & Construction](/api/generate)

Functions for generating new CVE identifiers:

| Function | Description |
|----------|-------------|
| `GenerateCve(year int, seq int) string` | Generate CVE from year and sequence |

## Quick Reference

### Common Operations

```go
import "github.com/scagogogo/cve"

// Format
formatted := cve.Format(" cve-2022-12345 ")  // "CVE-2022-12345"

// Validate
isValid := cve.ValidateCve("CVE-2022-12345")  // true

// Extract
cves := cve.ExtractCve("Text contains CVE-2022-12345")  // ["CVE-2022-12345"]

// Sort
sorted := cve.SortCves([]string{"CVE-2022-2", "CVE-2021-1"})

// Group
grouped := cve.GroupByYear([]string{"CVE-2021-1", "CVE-2022-1"})
```

### Function Return Types

| Return Type | Description | Example |
|-------------|-------------|---------|
| `string` | Single string result, returns empty string for invalid input | `"CVE-2022-12345"` |
| `[]string` | String slice, returns empty slice when no results | `["CVE-2022-1", "CVE-2022-2"]` |
| `bool` | Boolean value, indicates yes/no or true/false | `true` |
| `int` | Integer, usually returns 0 for invalid input | `2022` |
| `map[string][]string` | Map from string to string slice | `{"2022": ["CVE-2022-1"]}` |

### Error Handling

CVE Utils functions are designed with good fault tolerance for invalid input:

- String functions return empty string `""` for invalid input
- Integer functions return `0` for invalid input
- Slice functions return empty slice `[]string{}` for invalid input
- Boolean functions return `false` for invalid or non-matching input

### Performance Characteristics

- **Memory Efficient**: Functions avoid unnecessary memory allocations
- **Concurrency Safe**: All functions are concurrency-safe (stateless)
- **Regex Optimization**: Uses compiled regular expressions internally
- **Batch Processing**: Supports efficient batch operations

## Usage Patterns

### 1. Data Cleaning Pipeline

```go
func cleanCveData(rawData []string) []string {
    // 1. Extract all possible CVEs
    var allCves []string
    for _, text := range rawData {
        cves := cve.ExtractCve(text)
        allCves = append(allCves, cves...)
    }

    // 2. Remove duplicates
    unique := cve.RemoveDuplicateCves(allCves)

    // 3. Validate
    var valid []string
    for _, cveId := range unique {
        if cve.ValidateCve(cveId) {
            valid = append(valid, cveId)
        }
    }

    // 4. Sort
    return cve.SortCves(valid)
}
```

### 2. Conditional Filtering

```go
func filterRecentHighPriority(cveList []string) []string {
    // Get CVEs from recent 2 years
    recent := cve.GetRecentCves(cveList, 2)

    // Further filtering (example: sequence number > 10000)
    var highPriority []string
    for _, cveId := range recent {
        if cve.ExtractCveSeqAsInt(cveId) > 10000 {
            highPriority = append(highPriority, cveId)
        }
    }
    
    return highPriority
}
```

### 3. Statistical Analysis

```go
func analyzeCveDistribution(cveList []string) {
    // Group by year
    grouped := cve.GroupByYear(cveList)

    // Count for each year
    for year, cves := range grouped {
        fmt.Printf("Year %s: %d CVEs\n", year, len(cves))

        // Analyze sequence number range
        if len(cves) > 0 {
            sorted := cve.SortCves(cves)
            firstSeq := cve.ExtractCveSeqAsInt(sorted[0])
            lastSeq := cve.ExtractCveSeqAsInt(sorted[len(sorted)-1])
            fmt.Printf("  Sequence range: %d - %d\n", firstSeq, lastSeq)
        }
    }
}
```

## Version Compatibility

Current API version: `v1.0.0`

- ✅ **Stable API**: All public function signatures remain stable
- ✅ **Backward Compatible**: New versions maintain backward compatibility
- ✅ **Semantic Versioning**: Follows semantic versioning specification

## Next Steps

Choose the API category you're interested in to learn more:

- [Format & Validation](/api/format-validate) - Learn how to validate and standardize CVEs
- [Extraction Methods](/api/extract) - Understand how to extract CVE information from text
- [Comparison & Sorting](/api/compare-sort) - Master CVE comparison and sorting techniques
- [Filtering & Grouping](/api/filter-group) - Learn advanced filtering and grouping operations
- [Generation & Construction](/api/generate) - Understand how to generate new CVE identifiers

Or check out [Examples](/examples/) to see real-world usage scenarios.
