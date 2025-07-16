package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/scagogogo/cve"
)

func main() {
	fmt.Println("按年份分组CVE示例")
	// 预期输出:
	// 按年份分组CVE示例

	// 创建一个CVE列表
	cveList := []string{
		"CVE-2022-22965", // Spring4Shell
		"CVE-2021-44228", // Log4Shell
		"CVE-2020-1337",  // 2020年的CVE
		"CVE-2021-3156",  // Sudo漏洞
		"CVE-2022-0847",  // Dirty Pipe
		"CVE-2020-0601",  // Windows CryptoAPI
		"CVE-2019-0708",  // BlueKeep
		"CVE-2017-0144",  // EternalBlue
		"CVE-2022-42889", // Text4Shell
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
	// [4] CVE-2021-3156
	// [5] CVE-2022-0847
	// [6] CVE-2020-0601
	// [7] CVE-2019-0708
	// [8] CVE-2017-0144
	// [9] CVE-2022-42889

	// 按年份分组
	groupedCves := cve.GroupByYear(cveList)

	// 获取所有年份并排序，以便按年份顺序打印
	years := make([]string, 0, len(groupedCves))
	for year := range groupedCves {
		years = append(years, year)
	}

	// 转换为整数进行排序
	yearInts := make([]int, len(years))
	for i, year := range years {
		yearInt, _ := strconv.Atoi(year)
		yearInts[i] = yearInt
	}
	sort.Ints(yearInts)

	// 将排序后的整数年份转回字符串
	sortedYears := make([]string, len(yearInts))
	for i, yearInt := range yearInts {
		sortedYears[i] = strconv.Itoa(yearInt)
	}

	// 打印分组结果
	fmt.Println("\n按年份分组结果:")
	for _, year := range sortedYears {
		yearInt, _ := strconv.Atoi(year)
		fmt.Printf("%d年的CVE (%d个):\n", yearInt, len(groupedCves[year]))
		for i, id := range groupedCves[year] {
			fmt.Printf("  [%d] %s\n", i+1, id)
		}
	}
	// 预期输出:
	// 按年份分组结果:
	// 2017年的CVE (1个):
	//   [1] CVE-2017-0144
	// 2019年的CVE (1个):
	//   [1] CVE-2019-0708
	// 2020年的CVE (2个):
	//   [1] CVE-2020-0601
	//   [2] CVE-2020-1337
	// 2021年的CVE (2个):
	//   [1] CVE-2021-3156
	//   [2] CVE-2021-44228
	// 2022年的CVE (3个):
	//   [1] CVE-2022-0847
	//   [2] CVE-2022-22965
	//   [3] CVE-2022-42889

	// 示例：统计每年的CVE数量
	fmt.Println("\n每年CVE数量统计:")
	for _, year := range sortedYears {
		yearInt, _ := strconv.Atoi(year)
		fmt.Printf("%d年: %d个CVE\n", yearInt, len(groupedCves[year]))
	}
	// 预期输出:
	// 每年CVE数量统计:
	// 2017年: 1个CVE
	// 2019年: 1个CVE
	// 2020年: 2个CVE
	// 2021年: 2个CVE
	// 2022年: 3个CVE

	// 应用场景
	fmt.Println("\n应用场景示例:")
	fmt.Println("1. 漏洞趋势分析：对比不同年份的CVE数量和类型")
	fmt.Println("2. 漏洞响应优先级：优先处理最近年份的CVE")
	fmt.Println("3. 报告生成：按年份组织CVE列表，生成漏洞报告")
	// 预期输出:
	// 应用场景示例:
	// 1. 漏洞趋势分析：对比不同年份的CVE数量和类型
	// 2. 漏洞响应优先级：优先处理最近年份的CVE
	// 3. 报告生成：按年份组织CVE列表，生成漏洞报告
}
