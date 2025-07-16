# Installation Guide

This page provides detailed instructions on how to install and configure CVE Utils in different environments.

## System Requirements

- Go 1.18 or higher
- Supported operating systems: Linux, macOS, Windows

## Installation Methods

### Method 1: Using go get (Recommended)

This is the simplest and recommended installation method:

```bash
go get github.com/scagogogo/cve
```

### Method 2: Using go mod

If you are using Go modules in your project:

1. Initialize your module (if not already done):
```bash
go mod init your-project-name
```

2. Add CVE Utils as a dependency:
```bash
go get github.com/scagogogo/cve
```

3. Import in your Go code:
```go
import "github.com/scagogogo/cve"
```

## Verification

After installation, verify that CVE Utils is working correctly:

```go
package main

import (
    "fmt"
    "github.com/scagogogo/cve"
)

func main() {
    // Test basic functionality
    result := cve.Format("cve-2022-12345")
    if result == "CVE-2022-12345" {
        fmt.Println("✅ CVE Utils installed successfully!")
    } else {
        fmt.Println("❌ Installation failed")
    }
}
```

## Troubleshooting

### Common Issues

1. **Go version too old**: Ensure you are using Go 1.18 or higher
2. **Module not found**: Make sure you have internet access and can reach GitHub
3. **Import errors**: Verify the import path is correct: `github.com/scagogogo/cve`

### Getting Help

If you encounter issues:
1. Check the [GitHub Issues](https://github.com/scagogogo/cve/issues)
2. Review the [documentation](https://scagogogo.github.io/cve/)
3. Create a new issue if needed

## Next Steps

- Continue to [Getting Started](/guide/getting-started)
- Read the [Basic Usage Guide](/guide/basic-usage)
- Explore the [API Reference](/api/)
