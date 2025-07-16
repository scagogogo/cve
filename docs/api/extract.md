# Extraction Methods

This category of functions is used to extract CVE information from text.

## ExtractCve

Extract all CVE identifiers from a string.

### Function Signature

```go
func ExtractCve(text string) []string
```

### Example

```go
package main

import (
    "fmt"
    "github.com/scagogogo/cve"
)

func main() {
    text := "System affected by CVE-2021-44228 and cve-2022-12345"
    cves := cve.ExtractCve(text)
    fmt.Printf("Found CVEs: %v\n", cves)
    // Output: [CVE-2021-44228 CVE-2022-12345]
}
```

## ExtractFirstCve

Extract the first CVE identifier from text.

### Function Signature

```go
func ExtractFirstCve(text string) string
```

## Split

Split a CVE identifier into year and sequence components.

### Function Signature

```go
func Split(cve string) (year string, seq string)
```

## ExtractCveYear

Extract the year component from a CVE identifier.

### Function Signature

```go
func ExtractCveYear(cve string) string
```

## Best Practices

1. **Use ExtractCve()** for comprehensive text scanning
2. **Use Split()** when you need both year and sequence
3. **Always validate** extracted CVEs using ValidateCve()
