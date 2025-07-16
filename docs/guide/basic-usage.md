# Basic Usage

This guide provides detailed instructions on the basic usage methods and best practices of CVE Utils.

## Import Package

First, import CVE Utils in your Go code:

```go
import "github.com/scagogogo/cve"
```

## Core Concepts

### CVE Format Specification

CVE identifiers follow this format:
- Format: `CVE-YYYY-NNNN`
- `CVE`: Fixed prefix (case insensitive)
- `YYYY`: 4-digit year (from 1999 to present)
- `NNNN`: Sequence number (at least 4 digits, zero-padded)

### Examples of Valid CVEs

```go
"CVE-2022-1234"    // Standard format
"CVE-2021-0001"    // Zero-padded sequence
"CVE-2023-123456"  // Longer sequence number
```

## Basic Operations

### 1. Format CVE

The `Format` function standardizes CVE identifiers:

```go
package main

import (
    "fmt"
    "github.com/scagogogo/cve"
)

func main() {
    // Various input formats
    inputs := []string{
        " cve-2022-12345 ",  // with spaces
        "CVE-2022-12345",    // already standard format
        "cVe-2022-12345",    // mixed case
    }
    
    for _, input := range inputs {
        formatted := cve.Format(input)
        fmt.Printf("'%s' -> '%s'\n", input, formatted)
    }
}
```

Output:
```
' cve-2022-12345 ' -> 'CVE-2022-12345'
'CVE-2022-12345' -> 'CVE-2022-12345'
'cVe-2022-12345' -> 'CVE-2022-12345'
```

### 2. Validate CVE

#### Basic Format Validation

```go
func validateExamples() {
    testCases := []string{
        "CVE-2022-1234",     // Valid
        "cve-2021-5678",     // Valid (case insensitive)
        "CVE-22-1234",       // Invalid (year too short)
        "CVE-2022-123",      // Invalid (sequence too short)
        "not-a-cve",         // Invalid (wrong format)
    }
    
    for _, test := range testCases {
        isValid := cve.IsCve(test)
        fmt.Printf("%-20s -> %t\n", test, isValid)
    }
}
```

#### Comprehensive Validation

```go
func comprehensiveValidation() {
    testCases := []string{
        "CVE-2022-1234",     // Valid
        "CVE-2030-1234",     // Invalid (future year)
        "CVE-1998-1234",     // Invalid (year too old)
    }
    
    for _, test := range testCases {
        isValid := cve.ValidateCve(test)
        fmt.Printf("%-20s -> %t\n", test, isValid)
    }
}
```

### 3. Extract CVE from Text

```go
func extractExamples() {
    text := "System affected by CVE-2021-44228 and CVE-2022-12345 vulnerabilities"
    
    // Extract all CVEs
    cves := cve.ExtractCve(text)
    fmt.Printf("Found CVEs: %v\n", cves)
    
    // Extract first CVE
    first := cve.ExtractFirstCve(text)
    fmt.Printf("First CVE: %s\n", first)
    
    // Extract last CVE
    last := cve.ExtractLastCve(text)
    fmt.Printf("Last CVE: %s\n", last)
}
```

### 4. Extract CVE Components

```go
func componentExamples() {
    cveId := "CVE-2022-12345"
    
    // Extract year and sequence
    year, seq := cve.Split(cveId)
    fmt.Printf("CVE: %s -> Year: %s, Sequence: %s\n", cveId, year, seq)
    
    // Extract year as integer
    yearInt := cve.ExtractCveYearAsInt(cveId)
    fmt.Printf("Year as int: %d\n", yearInt)
    
    // Extract sequence as integer
    seqInt := cve.ExtractCveSeqAsInt(cveId)
    fmt.Printf("Sequence as int: %d\n", seqInt)
}
```

## Working with CVE Collections

### Sorting CVEs

```go
func sortingExample() {
    cves := []string{
        "CVE-2022-9999",
        "CVE-2021-1234",
        "CVE-2022-1111",
        "CVE-2023-5678",
    }
    
    sorted := cve.SortCves(cves)
    fmt.Printf("Sorted: %v\n", sorted)
    // Output: [CVE-2021-1234 CVE-2022-1111 CVE-2022-9999 CVE-2023-5678]
}
```

### Filtering CVEs

```go
func filteringExample() {
    cves := []string{
        "CVE-2021-1234",
        "CVE-2022-5678",
        "CVE-2021-9999",
        "CVE-2023-1111",
    }
    
    // Filter by year
    cves2021 := cve.FilterCvesByYear(cves, 2021)
    fmt.Printf("2021 CVEs: %v\n", cves2021)
    
    // Filter by year range
    recent := cve.FilterCvesByYearRange(cves, 2022, 2023)
    fmt.Printf("2022-2023 CVEs: %v\n", recent)
    
    // Group by year
    grouped := cve.GroupByYear(cves)
    fmt.Printf("Grouped: %v\n", grouped)
}
```

## Best Practices

### 1. Always Format Before Processing

```go
func processUserInput(input string) {
    // Always format user input first
    formatted := cve.Format(input)
    
    // Then validate
    if cve.ValidateCve(formatted) {
        // Process the valid CVE
        fmt.Printf("Processing: %s\n", formatted)
    } else {
        fmt.Printf("Invalid CVE: %s\n", input)
    }
}
```

### 2. Use Appropriate Validation

```go
func chooseValidation(cveId string) {
    // For quick format checking
    if cve.IsCve(cveId) {
        fmt.Println("Valid format")
    }
    
    // For comprehensive validation (including year checks)
    if cve.ValidateCve(cveId) {
        fmt.Println("Valid CVE")
    }
}
```

### 3. Handle Collections Efficiently

```go
func processCollection(rawCves []string) []string {
    // 1. Remove duplicates first
    unique := cve.RemoveDuplicateCves(rawCves)
    
    // 2. Filter valid CVEs
    var valid []string
    for _, cveId := range unique {
        if cve.ValidateCve(cveId) {
            valid = append(valid, cveId)
        }
    }
    
    // 3. Sort for consistent output
    return cve.SortCves(valid)
}
```

## Error Handling

CVE Utils functions are designed to be safe and never panic:

```go
func safeUsage() {
    // These will not panic, even with invalid input
    fmt.Println(cve.Format(""))           // Returns empty string
    fmt.Println(cve.IsCve("invalid"))     // Returns false
    fmt.Println(cve.ExtractCve("no cve")) // Returns empty slice
    
    // Always check results
    cves := cve.ExtractCve("some text")
    if len(cves) > 0 {
        fmt.Printf("Found %d CVEs\n", len(cves))
    } else {
        fmt.Println("No CVEs found")
    }
}
```

## Performance Considerations

- All functions are optimized for performance
- Use `RemoveDuplicateCves()` before processing large collections
- `GroupByYear()` is efficient for statistical analysis
- Validation functions are fast and suitable for real-time use

## Next Steps

- Explore the [API Reference](/api/) for detailed function documentation
- Check out [Examples](/examples/) for real-world usage scenarios
