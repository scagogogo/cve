package main

import (
	"fmt"

	"github.com/scagogogo/cve"
)

func main() {
	fmt.Println("=== CVE 连续性判断 ===")

	pairs := []struct {
		a, b string
	}{
		{"CVE-2022-12345", "CVE-2022-12346"},
		{"CVE-2022-12345", "CVE-2022-12347"},
		{"CVE-2022-12345", "CVE-2023-12345"},
		{"CVE-2022-12346", "CVE-2022-12345"},
		{"CVE-2022-12345", "CVE-2022-12345"},
	}

	for _, p := range pairs {
		consecutive := cve.IsCvesConsecutive(p.a, p.b)
		mark := "✗"
		if consecutive {
			mark = "✓"
		}
		fmt.Printf("  %s %s <-> %s: 连续=%v\n", mark, p.a, p.b, consecutive)
	}

	fmt.Println("\n--- 检测可合并列表 ---")
	cveList := []string{
		"CVE-2022-1001", "CVE-2022-1002", "CVE-2022-1003",
		"CVE-2022-2001", "CVE-2022-2003",
	}
	fmt.Println("CVE列表:", cveList)

	for i := 0; i < len(cveList)-1; i++ {
		if cve.IsCvesConsecutive(cveList[i], cveList[i+1]) {
			fmt.Printf("  %s 和 %s 连续\n", cveList[i], cveList[i+1])
		} else {
			fmt.Printf("  %s 和 %s 不连续\n", cveList[i], cveList[i+1])
		}
	}
}
