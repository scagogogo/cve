package main

import (
	"fmt"

	"github.com/scagogogo/cve"
)

func main() {
	fmt.Println("CVE排序示例")
	// 预期输出:
	// CVE排序示例

	// 创建一个混乱顺序的CVE列表
	cveList := []string{
		"CVE-2022-22965", // Spring4Shell
		"cve-2021-44228", // Log4Shell (小写格式)
		"CVE-2022-1234",  // 随机示例
		"CVE-2020-1337",  // 较早的CVE
		"CVE-2022-0000",  // 相同年份，序列号较小
		" CVE-2023-9999", // 带有空格的CVE
	}

	fmt.Println("原始CVE列表:")
	printCveList(cveList)
	// 预期输出:
	// 原始CVE列表:
	// [1] CVE-2022-22965
	// [2] cve-2021-44228
	// [3] CVE-2022-1234
	// [4] CVE-2020-1337
	// [5] CVE-2022-0000
	// [6]  CVE-2023-9999

	// 使用SortCves函数对列表进行排序
	sortedList := cve.SortCves(cveList)

	fmt.Println("\n排序后的CVE列表:")
	printCveList(sortedList)
	// 预期输出:
	// 排序后的CVE列表:
	// [1] CVE-2020-1337
	// [2] CVE-2021-44228
	// [3] CVE-2022-0000
	// [4] CVE-2022-1234
	// [5] CVE-2022-22965
	// [6] CVE-2023-9999

	// 演示SortCves函数的格式化功能
	fmt.Println("\n注意事项:")
	fmt.Println("1. SortCves函数会自动对所有CVE进行格式化")
	fmt.Println("2. 排序首先按年份，然后按序列号进行")
	// 预期输出:
	// 注意事项:
	// 1. SortCves函数会自动对所有CVE进行格式化
	// 2. 排序首先按年份，然后按序列号进行

	// 实际应用场景
	fmt.Println("\n应用场景示例 - 按时间顺序显示CVE的安全公告:")
	for i, id := range sortedList {
		var description string
		switch id {
		case "CVE-2020-1337":
			description = "Windows内核权限提升漏洞"
		case "CVE-2021-44228":
			description = "Log4Shell远程代码执行漏洞"
		case "CVE-2022-0000":
			description = "示例低序列号漏洞"
		case "CVE-2022-1234":
			description = "示例中等序列号漏洞"
		case "CVE-2022-22965":
			description = "Spring4Shell远程代码执行漏洞"
		case "CVE-2023-9999":
			description = "示例未来漏洞"
		}
		fmt.Printf("%d. %s - %s\n", i+1, id, description)
	}
	// 预期输出:
	// 应用场景示例 - 按时间顺序显示CVE的安全公告:
	// 1. CVE-2020-1337 - Windows内核权限提升漏洞
	// 2. CVE-2021-44228 - Log4Shell远程代码执行漏洞
	// 3. CVE-2022-0000 - 示例低序列号漏洞
	// 4. CVE-2022-1234 - 示例中等序列号漏洞
	// 5. CVE-2022-22965 - Spring4Shell远程代码执行漏洞
	// 6. CVE-2023-9999 - 示例未来漏洞
}

// 辅助函数：打印CVE列表
func printCveList(list []string) {
	for i, id := range list {
		fmt.Printf("[%d] %s\n", i+1, id)
	}
}
