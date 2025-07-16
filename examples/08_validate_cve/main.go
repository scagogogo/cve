package main

import (
	"fmt"
	"time"

	"github.com/scagogogo/cve"
)

func main() {
	currentYear := time.Now().Year()

	fmt.Println("CVE全面验证示例:")

	// 有效的CVE示例
	validCVEs := []string{
		fmt.Sprintf("CVE-%d-12345", currentYear),   // 当前年份
		fmt.Sprintf("CVE-%d-12345", currentYear-1), // 去年
		"CVE-2020-1234", // 2020年
		"CVE-1999-0001", // 较早的CVE
	}

	fmt.Println("\n有效的CVE示例:")
	for _, id := range validCVEs {
		fmt.Printf("%s: %v\n", id, cve.ValidateCve(id))
	}

	// 无效的CVE示例
	invalidCVEs := []string{
		fmt.Sprintf("CVE-%d-12345", currentYear+1), // 未来年份
		"CVE-1998-1234",      // 早于1999
		"CVE-2022-ABC",       // 序列号不是数字
		"CVE2022-1234",       // 格式错误，缺少连字符
		"包含CVE-2022-1234的文本", // 非独立CVE
		"cve-2022--1234",     // 双连字符
		"CVE-2022-0",         // 序列号太短
	}

	fmt.Println("\n无效的CVE示例:")
	for _, id := range invalidCVEs {
		fmt.Printf("%s: %v\n", id, cve.ValidateCve(id))
	}

	// 解释验证规则
	fmt.Println("\nValidateCve函数验证规则说明:")
	fmt.Println("1. 必须是完整的CVE格式 (如 'CVE-YYYY-NNNNN')")
	fmt.Println("2. 年份必须在1999年至当前年份之间")
	fmt.Println("3. 序列号必须是正整数")
}
