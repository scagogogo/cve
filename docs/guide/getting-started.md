# Getting Started

Welcome to CVE Utils! This guide will help you quickly get started with this powerful CVE processing library.

## Installation

Install using go get:

```bash
go get github.com/scagogogo/cve
```

## Verify Installation

Create a simple test file:

```go
package main

import (
    "fmt"
    "github.com/scagogogo/cve"
)

func main() {
    result := cve.Format("cve-2022-12345")
    fmt.Println("Format result:", result)
    
    if result == "CVE-2022-12345" {
        fmt.Println("✅ CVE Utils installed successfully!")
    } else {
        fmt.Println("❌ Installation may have issues")
    }
}
```

## Basic Concepts

### CVE Format

CVE identifiers follow this format:
- Format: `CVE-YYYY-NNNN`
- `CVE`: Fixed prefix (case insensitive)
- `YYYY`: 4-digit year (from 1999 to present)
- `NNNN`: Sequence number (at least 4 digits)

## Quick Examples

### Format CVE

```go
formatted := cve.Format("cve-2022-12345")
fmt.Println(formatted) // Output: CVE-2022-12345
```

### Validate CVE

```go
isValid := cve.ValidateCve("CVE-2022-12345")
fmt.Println(isValid) // Output: true
```

### Extract CVEs from Text

```go
text := "System affected by CVE-2021-44228 and CVE-2022-12345"
cves := cve.ExtractCve(text)
fmt.Println(cves) // Output: [CVE-2021-44228 CVE-2022-12345]
```

## Next Steps

- Read the [Basic Usage Guide](/guide/basic-usage)
- Explore the [API Reference](/api/)
- Check out [Examples](/examples/)
