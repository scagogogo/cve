# Statistical Analysis

This category of functions provides statistical analysis of CVE data, including year-based counting, year range detection, and sequence number range analysis.

## CountByYear

Count the number of CVEs per year.

### Function Signature

```go
func CountByYear(cveSlice []string) map[int]int
```

### Parameters

- `cveSlice` ([]string): Slice of CVE identifiers

### Return Value

- `map[int]int`: Map from year to CVE count

### Description

The `CountByYear` function groups CVE identifiers by year and returns the count for each year. Invalid CVEs (unparseable year) are excluded from the result.

### Example

```go
cves := []string{"CVE-2022-1111", "CVE-2022-2222", "CVE-2021-3333", "CVE-2022-4444"}
counts := cve.CountByYear(cves)
fmt.Println(counts)
// Output: map[2021:1 2022:3]
```

## YearRange

Get the earliest and latest year from a CVE list.

### Function Signature

```go
func YearRange(cveSlice []string) (min, max int)
```

### Parameters

- `cveSlice` ([]string): Slice of CVE identifiers

### Return Value

- `min` (int): Earliest year; 0 if no valid CVEs found
- `max` (int): Latest year; 0 if no valid CVEs found

### Example

```go
cves := []string{"CVE-2020-1111", "CVE-2022-2222", "CVE-2021-3333"}
min, max := cve.YearRange(cves)
fmt.Printf("Range: %d - %d\n", min, max)
// Output: Range: 2020 - 2022
```

## SeqRange

Get the sequence number range for CVEs in a specific year.

### Function Signature

```go
func SeqRange(cveSlice []string, year int) (min, max int)
```

### Parameters

- `cveSlice` ([]string): Slice of CVE identifiers
- `year` (int): Target year to analyze

### Return Value

- `min` (int): Smallest sequence number in the given year; 0 if no CVEs found
- `max` (int): Largest sequence number in the given year; 0 if no CVEs found

### Example

```go
cves := []string{"CVE-2022-1111", "CVE-2022-5555", "CVE-2022-3333", "CVE-2021-9999"}
min, max := cve.SeqRange(cves, 2022)
fmt.Printf("2022 sequence range: %d - %d\n", min, max)
// Output: 2022 sequence range: 1111 - 5555
```
