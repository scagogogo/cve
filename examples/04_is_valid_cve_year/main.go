package main

import (
	"fmt"

	"github.com/scagogogo/cve"
)

func main() {
	// 示例1：检查有效的CVE年份
	// CVE编号的年份部分从1999年开始，到当前年份有效
	year1 := 2021
	fmt.Printf("年份: %d\n", year1)
	fmt.Printf("是否为有效的CVE年份: %v\n\n", cve.IsCveYearOk(fmt.Sprintf("CVE-%d-1234", year1)))
	// 预期输出:
	// 年份: 2021
	// 是否为有效的CVE年份: true

	// 示例2：检查1999年（CVE编号的第一个有效年份）
	// 1999年是CVE编号系统开始使用的年份，是有效的
	year2 := 1999
	fmt.Printf("年份: %d\n", year2)
	fmt.Printf("是否为有效的CVE年份: %v\n\n", cve.IsCveYearOk(fmt.Sprintf("CVE-%d-1234", year2)))
	// 预期输出:
	// 年份: 1999
	// 是否为有效的CVE年份: true

	// 示例3：检查1998年（早于CVE编号系统开始的年份）
	// 1999年之前的年份对CVE编号无效
	year3 := 1998
	fmt.Printf("年份: %d\n", year3)
	fmt.Printf("是否为有效的CVE年份: %v\n\n", cve.IsCveYearOk(fmt.Sprintf("CVE-%d-1234", year3)))
	// 预期输出:
	// 年份: 1998
	// 是否为有效的CVE年份: false

	// 示例4：检查未来年份
	// 超过当前年份的未来年份不是有效的CVE年份
	year4 := 2100
	fmt.Printf("年份: %d\n", year4)
	fmt.Printf("是否为有效的CVE年份: %v\n", cve.IsCveYearOk(fmt.Sprintf("CVE-%d-1234", year4)))
	// 预期输出(假设当前年份小于2100):
	// 年份: 2100
	// 是否为有效的CVE年份: false

	// 总结: IsCveYearOk函数用于验证一个CVE编号的年份部分是否有效
	// 有效范围: 1999年至当前年份(含)
}
