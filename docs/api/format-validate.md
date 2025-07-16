# Format & Validation

This category of functions is used for CVE format standardization and validity verification.

## Format

Convert CVE identifier to standard uppercase format.

### Function Signature

```go
func Format(cve string) string
```

### Example

```go
fmt.Println(cve.Format("cve-2022-12345"))     // Output: CVE-2022-12345
```

## IsCve

Check if a string is a valid CVE format.

### Function Signature

```go
func IsCve(text string) bool
```

### Example

```go
fmt.Println(cve.IsCve("CVE-2022-12345"))  // Output: true
```

## ValidateCve

Comprehensive validation of CVE identifier.

### Function Signature

```go
func ValidateCve(cve string) bool
```

### Example

```go
fmt.Println(cve.ValidateCve("CVE-2022-12345"))  // Output: true
```

## Best Practices

1. **Always use Format()** before storing CVE identifiers
2. **Use ValidateCve()** for user input validation
3. **Use IsCve()** for quick format checking
