package main

import (
	"fmt"

	"github.com/scagogogo/cve-skills"
)

func main() {
	fmt.Println("=== CVE 年份范围 ===")

	cveList := []string{
		"CVE-2015-1001",
		"CVE-2018-2001",
		"CVE-2020-3001",
		"CVE-2022-4001",
		"CVE-2024-5001",
		"CVE-2025-6001",
	}

	minYear, maxYear := cve.YearRange(cveList)
	fmt.Println("CVE列表:", cveList)
	fmt.Printf("\n年份范围: %d - %d\n", minYear, maxYear)
	fmt.Printf("时间跨度: %d 年\n", maxYear-minYear)

	fmt.Println("\n--- 边界情况 ---")
	minE, maxE := cve.YearRange([]string{})
	fmt.Printf("空列表: min=%d, max=%d\n", minE, maxE)
}
