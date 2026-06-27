package main

import (
	"fmt"

	"github.com/scagogogo/cve-skills"
)

func main() {
	fmt.Println("=== 过滤有效CVE ===")

	mixedData := []string{
		"CVE-2022-1234",
		"invalid-data",
		"cve-2023-5678",
		"CVE-1998-0001",
		"CVE-2024-9999",
		"random-text",
		"CVE-2099-1234",
	}

	fmt.Println("混合数据:", mixedData)

	validCves := cve.FilterValidCves(mixedData)
	fmt.Printf("\n有效CVE: %v\n", validCves)
	fmt.Printf("有效数量: %d / %d\n", len(validCves), len(mixedData))

	fmt.Println("\n--- 与 ValidateCve 对比 ---")
	for _, item := range mixedData {
		status := "✗"
		if cve.ValidateCve(item) {
			status = "✓"
		}
		fmt.Printf("  %s %s\n", status, item)
	}
}
