package main

import (
	"fmt"
	"sort"

	"github.com/scagogogo/cve-skills"
)

func main() {
	fmt.Println("=== 按年份统计CVE ===")

	cveList := []string{
		"CVE-2019-1001", "CVE-2019-1002",
		"CVE-2020-1001", "CVE-2020-1002", "CVE-2020-1003",
		"CVE-2021-1001", "CVE-2021-1002",
		"CVE-2022-1001", "CVE-2022-1002", "CVE-2022-1003", "CVE-2022-1004",
		"CVE-2023-1001",
		"CVE-2024-1001", "CVE-2024-1002", "CVE-2024-1003",
	}

	counts := cve.CountByYear(cveList)

	var years []int
	for y := range counts {
		years = append(years, y)
	}
	sort.Ints(years)

	fmt.Println("年份分布:")
	fmt.Println("年份    | 数量 | 柱状图")
	fmt.Println("--------|------|------")
	for _, year := range years {
		count := counts[year]
		bar := ""
		for i := 0; i < count; i++ {
			bar += "█"
		}
		fmt.Printf("%d    | %4d | %s\n", year, count, bar)
	}

	fmt.Printf("\n总年份跨度: %d 年\n", len(counts))
	fmt.Printf("总计CVE: %d\n", len(cveList))
}
