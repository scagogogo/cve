package main

import (
	"fmt"

	"github.com/scagogogo/cve"
)

func main() {
	fmt.Println("按年份筛选CVE示例")
	// 预期输出:
	// 按年份筛选CVE示例

	// 创建一个包含不同年份CVE的列表
	cveList := []string{
		"CVE-2022-22965", // Spring4Shell
		"CVE-2021-44228", // Log4Shell
		"CVE-2020-1337",  // 2020年的CVE
		"CVE-2019-0708",  // BlueKeep
		"CVE-2017-0144",  // EternalBlue
		"CVE-2014-0160",  // Heartbleed
		"CVE-2023-9999",  // 2023年的CVE
	}

	fmt.Println("原始CVE列表:")
	for i, id := range cveList {
		fmt.Printf("[%d] %s\n", i+1, id)
	}
	// 预期输出:
	// 原始CVE列表:
	// [1] CVE-2022-22965
	// [2] CVE-2021-44228
	// [3] CVE-2020-1337
	// [4] CVE-2019-0708
	// [5] CVE-2017-0144
	// [6] CVE-2014-0160
	// [7] CVE-2023-9999

	// 筛选2022年的CVE
	cves2022 := cve.FilterCvesByYear(cveList, 2022)
	fmt.Println("\n2022年的CVE:")
	for i, id := range cves2022 {
		fmt.Printf("[%d] %s\n", i+1, id)
	}
	// 预期输出:
	// 2022年的CVE:
	// [1] CVE-2022-22965

	// 筛选2021年的CVE
	cves2021 := cve.FilterCvesByYear(cveList, 2021)
	fmt.Println("\n2021年的CVE:")
	for i, id := range cves2021 {
		fmt.Printf("[%d] %s\n", i+1, id)
	}
	// 预期输出:
	// 2021年的CVE:
	// [1] CVE-2021-44228

	// 筛选2020年的CVE
	cves2020 := cve.FilterCvesByYear(cveList, 2020)
	fmt.Println("\n2020年的CVE:")
	for i, id := range cves2020 {
		fmt.Printf("[%d] %s\n", i+1, id)
	}
	// 预期输出:
	// 2020年的CVE:
	// [1] CVE-2020-1337

	// 应用场景示例
	fmt.Println("\n应用场景示例 - 安全分析:")
	fmt.Println("安全团队需要对特定年份的CVE进行单独分析")
	fmt.Println("例如分析2021年的Log4Shell漏洞及相关CVE")
	// 预期输出:
	// 应用场景示例 - 安全分析:
	// 安全团队需要对特定年份的CVE进行单独分析
	// 例如分析2021年的Log4Shell漏洞及相关CVE

	// 注意事项
	fmt.Println("\n注意事项:")
	fmt.Println("1. 年份必须是有效的CVE年份(1999年至今)")
	fmt.Println("2. 非CVE格式的字符串会被自动过滤")
	fmt.Println("3. 年份不匹配的CVE会被过滤掉")
	// 预期输出:
	// 注意事项:
	// 1. 年份必须是有效的CVE年份(1999年至今)
	// 2. 非CVE格式的字符串会被自动过滤
	// 3. 年份不匹配的CVE会被过滤掉
}
