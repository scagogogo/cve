package main

import (
	"fmt"

	"github.com/scagogogo/cve-skills"
)

func main() {
	fmt.Println("=== CVE 序列号范围 ===")

	cveList := []string{
		"CVE-2022-1001", "CVE-2022-5050", "CVE-2022-3025",
		"CVE-2022-8888", "CVE-2022-1500", "CVE-2021-9999",
		"CVE-2023-1234", "CVE-2022-7777",
	}

	targetYears := []int{2022, 2021, 2023, 2020}

	for _, year := range targetYears {
		minSeq, maxSeq := cve.SeqRange(cveList, year)
		if minSeq == 0 && maxSeq == 0 {
			fmt.Printf("%d 年: 无CVE数据\n", year)
		} else {
			fmt.Printf("%d 年: 序列号范围 %d - %d (共 %d 个可能位置)\n",
				year, minSeq, maxSeq, maxSeq-minSeq+1)
		}
	}

	fmt.Println("\n--- 列出2022年所有CVE ---")
	cves2022 := cve.FilterCvesByYear(cveList, 2022)
	sorted := cve.SortCves(cves2022)
	fmt.Println(sorted)
}
