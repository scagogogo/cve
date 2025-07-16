package main

import (
	"fmt"
	"time"

	"github.com/scagogogo/cve"
)

func main() {
	currentYear := time.Now().Year()
	fmt.Printf("当前年份: %d\n\n", currentYear)

	// 测试正常年份的CVE
	normalCve := fmt.Sprintf("CVE-%d-12345", currentYear-1)
	fmt.Printf("去年的CVE: %s\n", normalCve)
	fmt.Printf("使用IsCveYearOk验证: %v\n\n", cve.IsCveYearOk(normalCve))

	// 测试当前年份的CVE
	currentYearCve := fmt.Sprintf("CVE-%d-12345", currentYear)
	fmt.Printf("当前年份的CVE: %s\n", currentYearCve)
	fmt.Printf("使用IsCveYearOk验证: %v\n\n", cve.IsCveYearOk(currentYearCve))

	// 测试未来年份的CVE (超出范围)
	futureCve := fmt.Sprintf("CVE-%d-12345", currentYear+2)
	fmt.Printf("未来年份的CVE: %s\n", futureCve)
	fmt.Printf("使用IsCveYearOk验证: %v\n", cve.IsCveYearOk(futureCve))

	// 使用带偏移量的验证 (允许未来2年)
	fmt.Printf("使用IsCveYearOkWithCutoff验证(偏移量=2): %v\n\n",
		cve.IsCveYearOkWithCutoff(futureCve, 2))

	// 测试未来年份的CVE (超出允许的偏移量)
	farFutureCve := fmt.Sprintf("CVE-%d-12345", currentYear+5)
	fmt.Printf("远期未来的CVE: %s\n", farFutureCve)
	fmt.Printf("使用IsCveYearOkWithCutoff验证(偏移量=2): %v\n\n",
		cve.IsCveYearOkWithCutoff(farFutureCve, 2))

	// 测试过早的年份
	oldCve := "CVE-1998-12345" // 早于1999年
	fmt.Printf("1999年之前的CVE: %s\n", oldCve)
	fmt.Printf("使用IsCveYearOk验证: %v\n", cve.IsCveYearOk(oldCve))
}
