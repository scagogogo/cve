# Batch Validation

This category of functions provides batch validation of CVE identifiers with detailed error reporting, useful for data quality checks when importing or processing large CVE datasets.

## ValidateCves

Batch validate CVE identifiers, returning detailed results for each.

### Function Signature

```go
func ValidateCves(cveSlice []string) []CveValidationResult
```

### CveValidationResult Type

```go
type CveValidationResult struct {
    Cve    string // The original CVE string
    Valid  bool   // Whether the CVE is valid
    Reason string // Reason for invalidity (empty if valid)
}
```

### Parameters

- `cveSlice` ([]string): Slice of CVE identifiers to validate

### Return Value

- `[]CveValidationResult`: Validation results for each input CVE

### Description

The `ValidateCves` function validates each CVE in the list and returns a structured result containing whether it's valid and, if not, a human-readable reason. Validation checks: format (must be CVE-YYYY-NNNNN), year (1999 to current year), and sequence number (positive integer).

### Example

```go
package main

import (
    "fmt"
    "github.com/scagogogo/cve"
)

func main() {
    cves := []string{"CVE-2022-1234", "not-a-cve", "CVE-1998-1234", "CVE-2023-5678"}

    results := cve.ValidateCves(cves)
    for _, r := range results {
        if r.Valid {
            fmt.Printf("✓ %s is valid\n", r.Cve)
        } else {
            fmt.Printf("✗ %s is invalid: %s\n", r.Cve, r.Reason)
        }
    }
    // Output:
    // ✓ CVE-2022-1234 is valid
    // ✗ not-a-cve is invalid: invalid CVE format
    // ✗ CVE-1998-1234 is invalid: year 1998 is before 1999
    // ✓ CVE-2023-5678 is valid
}
```

## FilterValidCves

Filter a CVE list to keep only valid CVE identifiers.

### Function Signature

```go
func FilterValidCves(cveSlice []string) []string
```

### Parameters

- `cveSlice` ([]string): Slice of CVE identifiers to filter

### Return Value

- `[]string`: Only the valid CVEs, formatted to uppercase

### Description

The `FilterValidCves` function is a convenience wrapper that filters out invalid CVEs from a list, returning only valid ones in standard uppercase format.

### Example

```go
raw := []string{"CVE-2022-1234", "invalid", "cve-2023-5678", "CVE-1998-1234"}
valid := cve.FilterValidCves(raw)
fmt.Println(valid)
// Output: [CVE-2022-1234 CVE-2023-5678]
```
