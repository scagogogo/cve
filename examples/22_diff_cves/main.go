package main

import (
	"fmt"

	"github.com/scagogogo/cve-skills"
)

func main() {
	fmt.Println("=== CVE 差集运算 (Difference) ===")

	currentScan := []string{"CVE-2024-1001", "CVE-2024-1002", "CVE-2024-1003", "CVE-2024-1004", "CVE-2024-1005"}
	previousScan := []string{"CVE-2024-1001", "CVE-2024-1003", "CVE-2024-1005"}

	fmt.Println("当前扫描结果:", currentScan)
	fmt.Println("前一次扫描结果:", previousScan)

	newCves := cve.DiffCves(currentScan, previousScan)
	fmt.Printf("\n新出现的CVE (差集): %v\n", newCves)
	fmt.Printf("新增数量: %d\n", len(newCves))

	fixedCves := cve.DiffCves(previousScan, currentScan)
	fmt.Printf("\n已修复的CVE (反向差集): %v\n", fixedCves)

	fmt.Println("\n--- 完全覆盖场景 ---")
	subset := []string{"CVE-2022-1111", "CVE-2022-2222"}
	superset := []string{"CVE-2022-1111", "CVE-2022-2222", "CVE-2022-3333"}
	fmt.Printf("差集结果: %v\n", cve.DiffCves(subset, superset))
}
