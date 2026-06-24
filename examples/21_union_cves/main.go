package main

import (
	"fmt"

	"github.com/scagogogo/cve"
)

func main() {
	fmt.Println("=== CVE 并集运算 (Union) ===")

	teamA := []string{"CVE-2023-1001", "CVE-2023-1002", "CVE-2023-1003"}
	teamB := []string{"CVE-2023-1003", "CVE-2023-1004", "CVE-2023-1005"}
	teamC := []string{"CVE-2023-1004", "CVE-2023-1005", "CVE-2023-1006"}

	fmt.Println("团队A的CVE:", teamA)
	fmt.Println("团队B的CVE:", teamB)
	fmt.Println("团队C的CVE:", teamC)

	merged := cve.UnionCves(teamA, teamB)
	merged = cve.UnionCves(merged, teamC)
	fmt.Printf("\n全部团队的CVE (并集): %v\n", merged)
	fmt.Printf("总唯一CVE数量: %d\n", len(merged))

	fmt.Println("\n--- 去重效果 ---")
	withDups := []string{"CVE-2022-1111", "cve-2022-1111", "CVE-2022-1111", "CVE-2022-2222"}
	unique := cve.UnionCves(withDups, []string{})
	fmt.Printf("原始 (含重复): %v\n", withDups)
	fmt.Printf("并集 (去重后): %v\n", unique)
}
