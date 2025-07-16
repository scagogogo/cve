package main

import (
	"fmt"

	"github.com/scagogogo/cve"
)

func main() {
	fmt.Println("按年份区间筛选CVE示例")
	// 预期输出:
	// 按年份区间筛选CVE示例

	// 创建一个包含不同年份CVE的列表
	cveList := []string{
		"CVE-2022-22965", // Spring4Shell
		"CVE-2021-44228", // Log4Shell
		"CVE-2020-1337",  // 2020年的CVE
		"CVE-2019-0708",  // BlueKeep
		"CVE-2017-0144",  // EternalBlue
		"CVE-2014-0160",  // Heartbleed
		"CVE-2010-3333",  // RTF Stack Buffer Overflow
		"CVE-2008-4250",  // Conficker
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
	// [7] CVE-2010-3333
	// [8] CVE-2008-4250

	// 筛选2020-2022年的CVE
	startYear := 2020
	endYear := 2022
	filteredCves := cve.FilterCvesByYearRange(cveList, startYear, endYear)

	fmt.Printf("\n%d-%d年的CVE:\n", startYear, endYear)
	for i, id := range filteredCves {
		fmt.Printf("[%d] %s\n", i+1, id)
	}
	// 预期输出:
	// 2020-2022年的CVE:
	// [1] CVE-2020-1337
	// [2] CVE-2021-44228
	// [3] CVE-2022-22965

	// 筛选2010-2018年的CVE
	startYear = 2010
	endYear = 2018
	filteredCves = cve.FilterCvesByYearRange(cveList, startYear, endYear)

	fmt.Printf("\n%d-%d年的CVE:\n", startYear, endYear)
	for i, id := range filteredCves {
		fmt.Printf("[%d] %s\n", i+1, id)
	}
	// 预期输出:
	// 2010-2018年的CVE:
	// [1] CVE-2010-3333
	// [2] CVE-2014-0160
	// [3] CVE-2017-0144

	// 应用场景示例
	fmt.Println("\n应用场景示例:")
	fmt.Println("1. 安全漏洞分析：分析特定时间范围内发布的CVE")
	fmt.Println("2. 合规性检查：检查系统是否受到特定年份范围内发布的CVE影响")
	fmt.Println("3. 趋势分析：分析不同年份区间CVE的数量和类型变化")
	// 预期输出:
	// 应用场景示例:
	// 1. 安全漏洞分析：分析特定时间范围内发布的CVE
	// 2. 合规性检查：检查系统是否受到特定年份范围内发布的CVE影响
	// 3. 趋势分析：分析不同年份区间CVE的数量和类型变化

	// 注意事项
	fmt.Println("\n注意事项:")
	fmt.Println("1. 年份范围包含起始年和结束年")
	fmt.Println("2. 如果开始年份大于结束年份，函数会返回空列表")
	fmt.Println("3. 筛选的年份必须在有效的CVE年份范围内(1999年至今)")
	// 预期输出:
	// 注意事项:
	// 1. 年份范围包含起始年和结束年
	// 2. 如果开始年份大于结束年份，函数会返回空列表
	// 3. 筛选的年份必须在有效的CVE年份范围内(1999年至今)
}
