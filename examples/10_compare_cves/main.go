package main

import (
	"fmt"

	"github.com/scagogogo/cve"
)

func main() {
	fmt.Println("CVE完整比较示例")
	// 预期输出:
	// CVE完整比较示例

	// 比较不同年份的CVE
	cve1 := "CVE-2020-1234"
	cve2 := "CVE-2022-5678"

	fmt.Printf("1. 比较不同年份的CVE: %s 和 %s\n", cve1, cve2)
	result1 := cve.CompareCves(cve1, cve2)
	printCompareResult(result1)
	// 预期输出:
	// 1. 比较不同年份的CVE: CVE-2020-1234 和 CVE-2022-5678
	// CompareCves结果: -1 (第一个CVE更早)

	// 比较相同年份但不同序列号的CVE
	cve3 := "CVE-2022-1111"
	cve4 := "CVE-2022-9999"

	fmt.Printf("\n2. 比较相同年份不同序列号的CVE: %s 和 %s\n", cve3, cve4)
	result2 := cve.CompareCves(cve3, cve4)
	printCompareResult(result2)
	// 预期输出:
	// 2. 比较相同年份不同序列号的CVE: CVE-2022-1111 和 CVE-2022-9999
	// CompareCves结果: -1 (第一个CVE更早)

	// 比较完全相同的CVE
	cve5 := "CVE-2022-1111"
	cve6 := "cve-2022-1111" // 大小写不同，但格式化后相同

	fmt.Printf("\n3. 比较完全相同的CVE (大小写不同): %s 和 %s\n", cve5, cve6)
	result3 := cve.CompareCves(cve5, cve6)
	printCompareResult(result3)
	// 预期输出:
	// 3. 比较完全相同的CVE (大小写不同): CVE-2022-1111 和 cve-2022-1111
	// CompareCves结果: 0 (两个CVE完全相同)

	// 反向比较
	fmt.Printf("\n4. 反向比较: %s 和 %s\n", cve2, cve1)
	result4 := cve.CompareCves(cve2, cve1)
	printCompareResult(result4)
	// 预期输出:
	// 4. 反向比较: CVE-2022-5678 和 CVE-2020-1234
	// CompareCves结果: 1 (第一个CVE更晚)

	// 演示使用场景
	fmt.Println("\n5. 使用场景示例 - 确定两个CVE的时间顺序:")
	cveA := "CVE-2021-44228" // Log4Shell漏洞
	cveB := "CVE-2022-22965" // Spring4Shell漏洞

	fmt.Printf("比较 %s 和 %s:\n", cveA, cveB)

	result := cve.CompareCves(cveA, cveB)
	if result < 0 {
		fmt.Printf("%s 出现在 %s 之前\n", cveA, cveB)
	} else if result > 0 {
		fmt.Printf("%s 出现在 %s 之后\n", cveA, cveB)
	} else {
		fmt.Printf("%s 和 %s 在同一时间点发布\n", cveA, cveB)
	}
	// 预期输出:
	// 5. 使用场景示例 - 确定两个CVE的时间顺序:
	// 比较 CVE-2021-44228 和 CVE-2022-22965:
	// CVE-2021-44228 出现在 CVE-2022-22965 之前
}

// 辅助函数：打印比较结果
func printCompareResult(result int) {
	fmt.Printf("CompareCves结果: %d ", result)
	switch result {
	case -1:
		fmt.Println("(第一个CVE更早)")
	case 0:
		fmt.Println("(两个CVE完全相同)")
	case 1:
		fmt.Println("(第一个CVE更晚)")
	}
}
