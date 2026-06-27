package main

import (
	"fmt"

	"github.com/scagogogo/cve-skills"
)

func main() {
	fmt.Println("=== CVE 范围解析 ===")

	fmt.Println("--- 使用 'to' 关键字 ---")
	range1 := "CVE-2022-1000 to CVE-2022-1005"
	result1 := cve.ParseCveRange(range1)
	fmt.Printf("输入: %s\n", range1)
	fmt.Printf("输出: %v\n", result1)

	fmt.Println("\n--- 使用 '..' 双点号 ---")
	range2 := "CVE-2022-2000..2003"
	result2 := cve.ParseCveRange(range2)
	fmt.Printf("输入: %s\n", range2)
	fmt.Printf("输出: %v\n", result2)

	fmt.Println("\n--- 使用 '-' 连字符 ---")
	range3 := "CVE-2022-3000-3002"
	result3 := cve.ParseCveRange(range3)
	fmt.Printf("输入: %s\n", range3)
	fmt.Printf("输出: %v\n", result3)

	fmt.Println("\n--- 无效输入处理 ---")
	fmt.Printf("空字符串: %v\n", cve.ParseCveRange(""))
	fmt.Printf("单CVE格式: %v\n", cve.ParseCveRange("CVE-2022-12345"))
	fmt.Printf("反向范围: %v\n", cve.ParseCveRange("CVE-2022-1005..1000"))

	fmt.Println("\n--- 实际应用: 统计范围内CVE总数 ---")
	securityBulletin := "CVE-2023-5000 to CVE-2023-5999"
	affectedCves := cve.ParseCveRange(securityBulletin)
	fmt.Printf("安全公告: %s\n", securityBulletin)
	fmt.Printf("受影响CVE总数: %d\n", len(affectedCves))
	if len(affectedCves) > 0 {
		fmt.Printf("范围: %s 到 %s\n", affectedCves[0], affectedCves[len(affectedCves)-1])
	}
}
