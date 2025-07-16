# Comparison & Sorting

This category of functions is used for CVE identifier comparison and sorting operations, supporting comparison by year, sequence number, or comprehensive comparison.

## CompareByYear

Compare CVEs based on their year.

### Function Signature

```go
func CompareByYear(cveA, cveB string) int
```

### Parameters

- `cveA` (string): First CVE identifier
- `cveB` (string): Second CVE identifier

### Return Value

- `int`: Comparison result
  - Negative: cveA year < cveB year
  - Zero: cveA year = cveB year  
  - Positive: cveA year > cveB year

### Description

The `CompareByYear` function only compares the year portion of two CVEs. This is useful for chronological analysis and time-based sorting.

### Example

```go
package main

import (
    "fmt"
    "github.com/scagogogo/cve"
)

func main() {
    result1 := cve.CompareByYear("CVE-2021-1234", "CVE-2022-5678")
    fmt.Println(result1) // Output: negative number (2021 < 2022)
    
    result2 := cve.CompareByYear("CVE-2022-1111", "CVE-2022-9999")
    fmt.Println(result2) // Output: 0 (same year)
}
```

## CompareCves

Comprehensive comparison of two CVE identifiers.

### Function Signature

```go
func CompareCves(cveA, cveB string) int
```

### Parameters

- `cveA` (string): First CVE identifier
- `cveB` (string): Second CVE identifier

### Return Value

- `int`: Comparison result
  - Negative: cveA < cveB
  - Zero: cveA = cveB
  - Positive: cveA > cveB

### Description

The `CompareCves` function performs a comprehensive comparison of two CVE identifiers by comparing both year and sequence number.

### Example

```go
package main

import (
    "fmt"
    "github.com/scagogogo/cve"
)

func main() {
    result1 := cve.CompareCves("CVE-2021-1234", "CVE-2022-5678")
    fmt.Println(result1) // Output: negative (2021 < 2022)
    
    result2 := cve.CompareCves("CVE-2022-1111", "CVE-2022-9999")
    fmt.Println(result2) // Output: negative (1111 < 9999)
}
```

## SortCves

Sort a slice of CVE identifiers.

### Function Signature

```go
func SortCves(cveSlice []string) []string
```

### Parameters

- `cveSlice` ([]string): Slice of CVE identifiers to sort

### Return Value

- `[]string`: New sorted slice of CVE identifiers

### Description

The `SortCves` function sorts CVE identifiers in ascending order by year first, then by sequence number.

### Example

```go
package main

import (
    "fmt"
    "github.com/scagogogo/cve"
)

func main() {
    cves := []string{
        "CVE-2022-9999",
        "CVE-2021-1234",
        "CVE-2022-1111",
        "CVE-2023-5678",
    }
    
    sorted := cve.SortCves(cves)
    fmt.Println(sorted)
    // Output: [CVE-2021-1234 CVE-2022-1111 CVE-2022-9999 CVE-2023-5678]
}
```

## Best Practices

1. **Use CompareCves()** for general sorting and comparison needs
2. **Use CompareByYear()** when you only care about chronological order
3. **Use SortCves()** for sorting collections of CVE identifiers

## Performance Notes

- All comparison functions have O(1) time complexity
- SortCves() has O(n log n) time complexity where n is the number of CVEs
