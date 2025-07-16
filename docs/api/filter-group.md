# Filtering & Grouping

This category of functions is used for CVE filtering, grouping, and deduplication operations, helping you organize and filter CVE data by various conditions.

## FilterCvesByYear

Filter CVEs by a specific year.

### Function Signature

```go
func FilterCvesByYear(cveSlice []string, year int) []string
```

### Parameters

- `cveSlice` ([]string): Slice of CVE identifiers to filter
- `year` (int): Target year to filter by

### Return Value

- `[]string`: New slice containing only CVEs from the specified year

### Description

The `FilterCvesByYear` function filters a slice of CVE identifiers to return only those from a specific year.

### Example

```go
package main

import (
    "fmt"
    "github.com/scagogogo/cve"
)

func main() {
    cves := []string{
        "CVE-2021-1234",
        "CVE-2022-5678",
        "CVE-2021-9999",
        "CVE-2023-1111",
    }
    
    cves2021 := cve.FilterCvesByYear(cves, 2021)
    fmt.Println(cves2021)
    // Output: [CVE-2021-1234 CVE-2021-9999]
}
```

## FilterCvesByYearRange

Filter CVEs by a year range.

### Function Signature

```go
func FilterCvesByYearRange(cveSlice []string, startYear, endYear int) []string
```

### Parameters

- `cveSlice` ([]string): Slice of CVE identifiers to filter
- `startYear` (int): Start year (inclusive)
- `endYear` (int): End year (inclusive)

### Return Value

- `[]string`: New slice containing CVEs within the specified year range

### Description

The `FilterCvesByYearRange` function filters CVEs to return only those within a specified year range.

### Example

```go
package main

import (
    "fmt"
    "github.com/scagogogo/cve"
)

func main() {
    cves := []string{
        "CVE-2020-1234",
        "CVE-2021-5678",
        "CVE-2022-9999",
        "CVE-2023-1111",
    }
    
    filtered := cve.FilterCvesByYearRange(cves, 2021, 2022)
    fmt.Println(filtered)
    // Output: [CVE-2021-5678 CVE-2022-9999]
}
```

## GroupByYear

Group CVEs by year.

### Function Signature

```go
func GroupByYear(cveSlice []string) map[string][]string
```

### Parameters

- `cveSlice` ([]string): Slice of CVE identifiers to group

### Return Value

- `map[string][]string`: Map where keys are years and values are slices of CVEs

### Description

The `GroupByYear` function groups CVE identifiers by their year.

### Example

```go
package main

import (
    "fmt"
    "github.com/scagogogo/cve"
)

func main() {
    cves := []string{
        "CVE-2021-1234",
        "CVE-2022-5678",
        "CVE-2021-9999",
        "CVE-2022-1111",
    }
    
    grouped := cve.GroupByYear(cves)
    fmt.Println(grouped)
    // Output: map[2021:[CVE-2021-1234 CVE-2021-9999] 2022:[CVE-2022-5678 CVE-2022-1111]]
}
```

## RemoveDuplicateCves

Remove duplicate CVE identifiers.

### Function Signature

```go
func RemoveDuplicateCves(cveSlice []string) []string
```

### Parameters

- `cveSlice` ([]string): Slice of CVE identifiers that may contain duplicates

### Return Value

- `[]string`: New slice with duplicate CVEs removed

### Description

The `RemoveDuplicateCves` function removes duplicate CVE identifiers from a slice.

### Example

```go
package main

import (
    "fmt"
    "github.com/scagogogo/cve"
)

func main() {
    cves := []string{
        "CVE-2021-1234",
        "CVE-2022-5678",
        "CVE-2021-1234", // duplicate
        "CVE-2023-9999",
    }
    
    unique := cve.RemoveDuplicateCves(cves)
    fmt.Println(unique)
    // Output: [CVE-2021-1234 CVE-2022-5678 CVE-2023-9999]
}
```

## Best Practices

1. **Use FilterCvesByYear()** for single-year analysis
2. **Use FilterCvesByYearRange()** for multi-year analysis
3. **Use GroupByYear()** for statistical analysis and reporting
4. **Use RemoveDuplicateCves()** before processing to ensure data quality

## Performance Notes

- All filtering functions create new slices and don't modify the original
- GroupByYear() has O(n) time complexity
- RemoveDuplicateCves() has O(n) time complexity
