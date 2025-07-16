package main

import (
	"fmt"

	"github.com/scagogogo/cve"
)

func main() {
	// 比较不同年份的CVE
	cve1 := "CVE-2020-1234"
	cve2 := "CVE-2022-5678"

	fmt.Printf("比较 %s 和 %s:\n", cve1, cve2)

	// 使用CompareByYear比较
	result := cve.CompareByYear(cve1, cve2)
	if result < 0 {
		fmt.Printf("CompareByYear结果: %d (第一个CVE的年份较早)\n", result)
	} else if result > 0 {
		fmt.Printf("CompareByYear结果: %d (第一个CVE的年份较晚)\n", result)
	} else {
		fmt.Printf("CompareByYear结果: %d (两个CVE的年份相同)\n", result)
	}

	// 使用SubByYear计算年份差
	diff := cve.SubByYear(cve1, cve2)
	fmt.Printf("SubByYear结果: %d (两个CVE的年份相差%d年)\n\n", diff, abs(diff))

	// 比较相同年份的CVE
	cve3 := "CVE-2022-1111"
	cve4 := "CVE-2022-9999"

	fmt.Printf("比较 %s 和 %s:\n", cve3, cve4)

	// 使用CompareByYear比较
	result2 := cve.CompareByYear(cve3, cve4)
	fmt.Printf("CompareByYear结果: %d (年份相同)\n", result2)

	// 使用SubByYear计算年份差
	diff2 := cve.SubByYear(cve3, cve4)
	fmt.Printf("SubByYear结果: %d (年份相同，无差异)\n\n", diff2)

	// 反向比较
	fmt.Printf("反向比较 %s 和 %s:\n", cve2, cve1)
	fmt.Printf("CompareByYear结果: %d\n", cve.CompareByYear(cve2, cve1))
	fmt.Printf("SubByYear结果: %d\n", cve.SubByYear(cve2, cve1))
}

// 辅助函数：计算绝对值
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
