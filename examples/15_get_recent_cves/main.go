package main

import (
	"fmt"
	"time"

	"github.com/scagogogo/cve"
)

func main() {
	fmt.Println("获取最近几年的CVE示例")
	// 预期输出:
	// 获取最近几年的CVE示例

	// 创建一个包含多个年份CVE的列表
	cveList := []string{
		"CVE-2023-22965", // 今年的CVE
		"CVE-2022-1111",  // 去年的CVE
		"CVE-2021-44228", // Log4Shell (2年前)
		"CVE-2020-1337",  // 3年前的CVE
		"CVE-2019-0708",  // BlueKeep (4年前)
		"CVE-2018-1000",  // 5年前的CVE
		"CVE-2017-0144",  // EternalBlue (6年前)
		"CVE-2016-0123",  // 7年前的CVE
	}

	fmt.Println("原始CVE列表:")
	for i, id := range cveList {
		fmt.Printf("[%d] %s\n", i+1, id)
	}
	// 预期输出:
	// 原始CVE列表:
	// [1] CVE-2023-22965
	// [2] CVE-2022-1111
	// [3] CVE-2021-44228
	// [4] CVE-2020-1337
	// [5] CVE-2019-0708
	// [6] CVE-2018-1000
	// [7] CVE-2017-0144
	// [8] CVE-2016-0123

	// 获取最近1年的CVE
	years := 1
	recentCves := cve.GetRecentCves(cveList, years)
	fmt.Printf("\n最近%d年的CVE (包含%d年):\n", years, time.Now().Year())
	for i, id := range recentCves {
		fmt.Printf("[%d] %s\n", i+1, id)
	}
	currentYear := time.Now().Year()
	// 预期输出:
	// 最近1年的CVE (包含2023年):
	// [1] CVE-2023-22965

	// 获取最近2年的CVE
	years = 2
	recentCves = cve.GetRecentCves(cveList, years)
	fmt.Printf("\n最近%d年的CVE (包含%d-%d年):\n", years, currentYear-years+1, currentYear)
	for i, id := range recentCves {
		fmt.Printf("[%d] %s\n", i+1, id)
	}
	// 预期输出:
	// 最近2年的CVE (包含2022-2023年):
	// [1] CVE-2022-1111
	// [2] CVE-2023-22965

	// 获取最近3年的CVE
	years = 3
	recentCves = cve.GetRecentCves(cveList, years)
	fmt.Printf("\n最近%d年的CVE (包含%d-%d年):\n", years, currentYear-years+1, currentYear)
	for i, id := range recentCves {
		fmt.Printf("[%d] %s\n", i+1, id)
	}
	// 预期输出:
	// 最近3年的CVE (包含2021-2023年):
	// [1] CVE-2021-44228
	// [2] CVE-2022-1111
	// [3] CVE-2023-22965

	// 应用场景示例
	fmt.Println("\n应用场景示例:")
	fmt.Println("1. 漏洞响应优先级 - 优先修复最近几年的漏洞")
	fmt.Println("2. 安全态势感知 - 分析最近一段时间内的CVE趋势")
	fmt.Println("3. 合规性检查 - 检查系统是否受到最近发布的高危CVE影响")
	// 预期输出:
	// 应用场景示例:
	// 1. 漏洞响应优先级 - 优先修复最近几年的漏洞
	// 2. 安全态势感知 - 分析最近一段时间内的CVE趋势
	// 3. 合规性检查 - 检查系统是否受到最近发布的高危CVE影响

	// 注意事项
	fmt.Println("\n注意事项:")
	fmt.Println("1. 函数会自动基于当前时间计算年份范围")
	fmt.Println("2. 包含参数指定的年数(包括当前年份)")
	fmt.Println("3. 结果会按照CVE的标准格式返回")
	// 预期输出:
	// 注意事项:
	// 1. 函数会自动基于当前时间计算年份范围
	// 2. 包含参数指定的年数(包括当前年份)
	// 3. 结果会按照CVE的标准格式返回
}

// 辅助函数：打印CVE列表
func printCveList(list []string) {
	for i, id := range list {
		fmt.Printf("[%d] %s\n", i+1, id)
	}
}
