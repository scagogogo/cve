package main

import (
	"fmt"

	"github.com/scagogogo/cve-skills"
)

func main() {
	fmt.Println("=== 通配符模式匹配 CVE ===")

	cveList := []string{
		"CVE-2021-1111", "CVE-2021-2222",
		"CVE-2022-1111", "CVE-2022-1122", "CVE-2022-2222", "CVE-2022-3333",
		"CVE-2023-1111", "CVE-2023-2222", "CVE-2023-3333",
	}

	fmt.Printf("CVE列表 (共 %d 个):\n", len(cveList))
	fmt.Println("  ", cveList)

	fmt.Println("\n--- 按年份筛选: CVE-2022-* ---")
	fmt.Println("  ", cve.FilterCvesByPattern(cveList, "CVE-2022-*"))

	fmt.Println("\n--- 按序列号筛选: CVE-*-1111 ---")
	fmt.Println("  ", cve.FilterCvesByPattern(cveList, "CVE-*-1111"))

	fmt.Println("\n--- 前缀匹配: CVE-2022-11* ---")
	fmt.Println("  ", cve.FilterCvesByPattern(cveList, "CVE-2022-11*"))

	fmt.Println("\n--- 精确匹配: CVE-2022-2222 ---")
	fmt.Println("  ", cve.FilterCvesByPattern(cveList, "CVE-2022-2222"))

	fmt.Println("\n--- 无匹配: CVE-2020-* ---")
	fmt.Println("  ", cve.FilterCvesByPattern(cveList, "CVE-2020-*"))
}
