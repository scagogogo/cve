package main

import (
	"fmt"

	"github.com/scagogogo/cve"
)

func main() {
	fmt.Println("=== CVE 交集运算 (Intersection) ===")

	scannerA := []string{"CVE-2022-1111", "CVE-2022-2222", "CVE-2022-3333", "CVE-2022-4444"}
	scannerB := []string{"CVE-2022-2222", "CVE-2022-3333", "CVE-2022-5555", "CVE-2022-6666"}

	fmt.Println("扫描器A发现的CVE:", scannerA)
	fmt.Println("扫描器B发现的CVE:", scannerB)

	common := cve.IntersectCves(scannerA, scannerB)
	fmt.Printf("\n共同发现的CVE (交集): %v\n", common)
	fmt.Printf("共同发现数量: %d\n", len(common))

	fmt.Println("\n--- 大小写不敏感示例 ---")
	list1 := []string{"cve-2022-1111", "CVE-2022-2222", "Cve-2022-3333"}
	list2 := []string{"CVE-2022-1111", "cve-2022-3333", "CVE-2022-4444"}
	fmt.Println("列表1:", list1)
	fmt.Println("列表2:", list2)
	fmt.Printf("交集: %v\n", cve.IntersectCves(list1, list2))

	fmt.Println("\n--- 空列表场景 ---")
	fmt.Printf("空列表交集: %v\n", cve.IntersectCves([]string{}, []string{"CVE-2022-1111"}))
}
