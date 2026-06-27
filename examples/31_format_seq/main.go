package main

import (
	"fmt"

	"github.com/scagogogo/cve-skills"
)

func main() {
	fmt.Println("=== CVE 序列号格式化 ===")

	cves := []string{
		"CVE-2022-1",
		"CVE-2022-12",
		"CVE-2022-123",
		"CVE-2022-1234",
		"CVE-2022-12345",
		"CVE-2022-123456",
	}

	fmt.Println("宽度为 6 的格式化效果:")
	fmt.Println("原始            | 格式化后")
	fmt.Println("----------------|---------")
	for _, id := range cves {
		formatted := cve.FormatSeq(id, 6)
		fmt.Printf("%-16s| %s\n", id, formatted)
	}

	fmt.Println("\n--- 不同宽度效果 (CVE-2022-123) ---")
	for _, width := range []int{4, 5, 6, 8} {
		fmt.Printf("  宽度 %d: %s\n", width, cve.FormatSeq("CVE-2022-123", width))
	}

	fmt.Println("\n--- 无效输入 ---")
	fmt.Printf("  'not-a-cve' -> %s\n", cve.FormatSeq("not-a-cve", 6))
}
