# Generation & Construction

This category of functions is used for generating and constructing new CVE identifiers from component parts.

## GenerateCve

Generate a CVE identifier from year and sequence number.

### Function Signature

```go
func GenerateCve(year int, seq int) string
```

### Parameters

- `year` (int): Year component of the CVE
- `seq` (int): Sequence number component of the CVE

### Return Value

- `string`: Generated CVE identifier in standard format

### Description

The `GenerateCve` function creates a properly formatted CVE identifier from a year and sequence number:
1. Validates that the year is reasonable (typically 1999 or later)
2. Ensures the sequence number is positive
3. Formats the result as "CVE-YYYY-NNNN" where NNNN is zero-padded to at least 4 digits

### Example

```go
package main

import (
    "fmt"
    "github.com/scagogogo/cve"
)

func main() {
    // Generate CVE with 4-digit sequence
    cve1 := cve.GenerateCve(2022, 1234)
    fmt.Println(cve1) // Output: CVE-2022-1234
    
    // Generate CVE with sequence less than 4 digits (zero-padded)
    cve2 := cve.GenerateCve(2023, 42)
    fmt.Println(cve2) // Output: CVE-2023-0042
    
    // Generate CVE with sequence more than 4 digits
    cve3 := cve.GenerateCve(2024, 123456)
    fmt.Println(cve3) // Output: CVE-2024-123456
    
    // Generate multiple CVEs in sequence
    for i := 1; i <= 5; i++ {
        cve := cve.GenerateCve(2023, i)
        fmt.Printf("CVE #%d: %s\n", i, cve)
    }
    // Output:
    // CVE #1: CVE-2023-0001
    // CVE #2: CVE-2023-0002
    // CVE #3: CVE-2023-0003
    // CVE #4: CVE-2023-0004
    // CVE #5: CVE-2023-0005
}
```

### Use Cases

1. **CVE Database Management**: Generate new CVE identifiers for vulnerability databases
2. **Testing and Simulation**: Create test CVE data for development and testing
3. **Data Migration**: Convert legacy vulnerability identifiers to CVE format
4. **Automated Systems**: Generate CVEs in automated vulnerability discovery systems

### Validation Rules

The function applies the following validation rules:
- Year must be 1999 or later (when CVE system was established)
- Sequence number must be positive (greater than 0)
- Generated CVE follows the standard CVE format specification

### Advanced Usage

```go
package main

import (
    "fmt"
    "time"
    "github.com/scagogogo/cve"
)

func main() {
    // Generate CVE for current year
    currentYear := time.Now().Year()
    newCve := cve.GenerateCve(currentYear, 1001)
    fmt.Printf("New CVE for %d: %s\n", currentYear, newCve)
    
    // Generate a batch of CVEs
    year := 2023
    startSeq := 1000
    count := 10
    
    fmt.Printf("Generating %d CVEs for year %d:\n", count, year)
    for i := 0; i < count; i++ {
        cveId := cve.GenerateCve(year, startSeq+i)
        fmt.Printf("  %s\n", cveId)
    }
    
    // Validate generated CVE
    generated := cve.GenerateCve(2022, 5678)
    isValid := cve.ValidateCve(generated)
    fmt.Printf("Generated CVE %s is valid: %t\n", generated, isValid)
}
```

### Integration with Other Functions

The `GenerateCve` function works seamlessly with other CVE Utils functions:

```go
package main

import (
    "fmt"
    "github.com/scagogogo/cve"
)

func main() {
    // Generate and validate
    generated := cve.GenerateCve(2023, 1234)
    isValid := cve.ValidateCve(generated)
    fmt.Printf("Generated: %s, Valid: %t\n", generated, isValid)
    
    // Generate and extract components
    cveId := cve.GenerateCve(2022, 5678)
    year := cve.ExtractCveYear(cveId)
    seq := cve.ExtractCveSeq(cveId)
    fmt.Printf("CVE: %s, Year: %s, Sequence: %s\n", cveId, year, seq)
    
    // Generate multiple and sort
    cves := []string{
        cve.GenerateCve(2023, 9999),
        cve.GenerateCve(2022, 1111),
        cve.GenerateCve(2023, 1234),
        cve.GenerateCve(2021, 5678),
    }
    
    sorted := cve.SortCves(cves)
    fmt.Printf("Sorted CVEs: %v\n", sorted)
    // Output: [CVE-2021-5678 CVE-2022-1111 CVE-2023-1234 CVE-2023-9999]
}
```

### Error Handling

The function handles edge cases gracefully:
- Invalid years return empty string or error (implementation dependent)
- Invalid sequence numbers return empty string or error (implementation dependent)
- Always check the result for validity using `ValidateCve()`

### Best Practices

1. **Always validate generated CVEs** using `ValidateCve()` before use
2. **Use reasonable years** (1999 or later for real CVEs)
3. **Maintain sequence number uniqueness** within the same year
4. **Consider zero-padding** for consistent formatting
5. **Document the source** of generated CVE identifiers

### Performance Notes

- `GenerateCve()` has O(1) time complexity
- String formatting is the primary operation
- Very fast for generating large batches of CVEs
- Memory usage is minimal

## Related Functions

- `ValidateCve()` - Validate generated CVE identifiers
- `ExtractCveYear()` - Extract year from generated CVE
- `ExtractCveSeq()` - Extract sequence from generated CVE
- `Format()` - Ensure consistent formatting
- `SortCves()` - Sort collections of generated CVEs
