# Set Operations

This category of functions provides set operations (intersection, union, difference) for CVE lists, useful when merging or comparing CVE data from multiple sources.

## IntersectCves

Compute the intersection of two CVE lists — returns CVEs that appear in both.

### Function Signature

```go
func IntersectCves(a, b []string) []string
```

### Parameters

- `a` ([]string): First CVE list
- `b` ([]string): Second CVE list

### Return Value

- `[]string`: CVEs present in both lists, deduplicated, formatted to uppercase, and sorted

### Description

The `IntersectCves` function returns CVE identifiers that exist in both input lists. Comparison is case-insensitive, and the result is sorted by year and sequence number.

### Example

```go
package main

import (
    "fmt"
    "github.com/scagogogo/cve"
)

func main() {
    list1 := []string{"CVE-2022-1111", "CVE-2022-2222", "CVE-2022-3333"}
    list2 := []string{"CVE-2022-2222", "CVE-2022-3333", "CVE-2022-4444"}

    common := cve.IntersectCves(list1, list2)
    fmt.Println(common)
    // Output: [CVE-2022-2222 CVE-2022-3333]
}
```

## UnionCves

Compute the union of two CVE lists — returns all CVEs from both lists, deduplicated.

### Function Signature

```go
func UnionCves(a, b []string) []string
```

### Parameters

- `a` ([]string): First CVE list
- `b` ([]string): Second CVE list

### Return Value

- `[]string`: All CVEs from both lists, deduplicated, formatted to uppercase, and sorted

### Description

The `UnionCves` function merges two CVE lists, removing duplicates. Comparison is case-insensitive, and the result is sorted by year and sequence number.

### Example

```go
list1 := []string{"CVE-2022-1111", "CVE-2022-2222"}
list2 := []string{"CVE-2022-2222", "CVE-2022-3333"}

all := cve.UnionCves(list1, list2)
fmt.Println(all)
// Output: [CVE-2022-1111 CVE-2022-2222 CVE-2022-3333]
```

## DiffCves

Compute the difference (a - b) of two CVE lists — returns CVEs in list a that are not in list b.

### Function Signature

```go
func DiffCves(a, b []string) []string
```

### Parameters

- `a` ([]string): The source CVE list (minuend)
- `b` ([]string): The CVE list to exclude (subtrahend)

### Return Value

- `[]string`: CVEs in `a` but not in `b`, deduplicated, formatted to uppercase, and sorted

### Description

The `DiffCves` function finds CVE identifiers that exist in list `a` but not in list `b`. This is useful for detecting new CVEs by comparing current data against historical data. Comparison is case-insensitive.

### Example

```go
current := []string{"CVE-2022-1111", "CVE-2022-2222", "CVE-2022-3333"}
historical := []string{"CVE-2022-2222", "CVE-2022-4444"}

newCves := cve.DiffCves(current, historical)
fmt.Println(newCves)
// Output: [CVE-2022-1111 CVE-2022-3333]
```
