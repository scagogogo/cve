package cve

import (
	"strconv"
	"time"
)

// GroupByYear 把一组CVE按照年份分组
//
// 将CVE列表按照年份进行分组
//
// 参数:
//   - cveSlice: 需要分组的CVE编号数组
//
// 返回值:
//   - map[string][]string: 分组结果，键为年份字符串（如"2021"），值为对应年份的CVE编号数组
//
// 示例:
//
//	输入: ["CVE-2021-1111", "CVE-2022-2222", "CVE-2021-3333"]
//	输出: {
//	  "2021": ["CVE-2021-1111", "CVE-2021-3333"],
//	  "2022": ["CVE-2022-2222"]
//	}
//
//	输入: ["CVE-2021-1111", "cve-2021-3333"] (不同大小写)
//	输出: {
//	  "2021": ["CVE-2021-1111", "CVE-2021-3333"]
//	} (注意返回结果已格式化为大写)
//
// 性能特性:
//   - 时间复杂度: O(n)，其中n为数组长度
//   - 空间复杂度: O(n)
//
// 使用场景:
//   - 按年份组织和展示多个CVE，例如生成年度漏洞报告
//   - 分析CVE随时间分布的趋势
//
// 代码示例:
//
//	cveList := []string{"CVE-2021-1111", "CVE-2022-2222", "CVE-2021-3333"}
//	yearGroups := cve.GroupByYear(cveList)
//	for year, cves := range yearGroups {
//	    fmt.Printf("%s年的CVE共有%d个\n", year, len(cves))
//	}
func GroupByYear(cveSlice []string) map[string][]string {
	groupMap := make(map[string][]string, 0)
	for _, cve := range cveSlice {
		year := ExtractCveYear(cve)
		groupMap[year] = append(groupMap[year], Format(cve))
	}
	return groupMap
}

// FilterCvesByYear 筛选特定年份的CVE
//
// 从CVE列表中筛选出指定年份的CVE编号
//
// 参数:
//   - cveSlice: 需要筛选的CVE编号数组
//   - year: 目标年份，整数格式，如2021
//
// 返回值:
//   - []string: 符合目标年份的CVE编号数组，已经过标准化格式处理
//     如果没有找到匹配项，则返回空数组
//
// 示例:
//
//	输入: ["CVE-2021-1111", "CVE-2022-2222", "CVE-2021-3333"], 2021
//	输出: ["CVE-2021-1111", "CVE-2021-3333"]
//
//	输入: ["CVE-2021-1111", "CVE-2022-2222"], 2023
//	输出: [] (空数组)
//
// 性能特性:
//   - 时间复杂度: O(n)，其中n为数组长度
//   - 空间复杂度: O(k)，其中k为结果数组长度（最坏情况为O(n)）
//
// 使用场景:
//   - 需要获取特定年份的CVE时使用
//   - 生成特定年度的安全报告
//
// 代码示例:
//
//	cveList := []string{"CVE-2021-1111", "CVE-2022-2222", "CVE-2021-3333"}
//	cves2021 := cve.FilterCvesByYear(cveList, 2021)
//	// cves2021 为 ["CVE-2021-1111", "CVE-2021-3333"]
func FilterCvesByYear(cveSlice []string, year int) []string {
	var result []string
	yearStr := strconv.Itoa(year)

	for _, cve := range cveSlice {
		formattedCve := Format(cve)
		if ExtractCveYear(formattedCve) == yearStr {
			result = append(result, formattedCve)
		}
	}

	return result
}

// FilterCvesByYearRange 筛选指定年份范围内的CVE
//
// 从CVE列表中筛选出在指定年份范围内的CVE编号
//
// 参数:
//   - cveSlice: 需要筛选的CVE编号数组
//   - startYear: 起始年份（含），整数格式
//   - endYear: 结束年份（含），整数格式
//
// 返回值:
//   - []string: 符合年份范围的CVE编号数组，已经过标准化格式处理
//     如果没有找到匹配项，则返回空数组
//
// 示例:
//
//	输入: ["CVE-2020-1111", "CVE-2021-2222", "CVE-2022-3333"], 2021, 2022
//	输出: ["CVE-2021-2222", "CVE-2022-3333"]
//
//	输入: ["CVE-2020-1111", "CVE-2021-2222"], 2022, 2023
//	输出: [] (空数组)
//
//	输入: ["CVE-2020-1111", "CVE-2021-2222", "CVE-2022-3333"], 2020, 2022
//	输出: ["CVE-2020-1111", "CVE-2021-2222", "CVE-2022-3333"]
//
// 性能特性:
//   - 时间复杂度: O(n)，其中n为数组长度
//   - 空间复杂度: O(k)，其中k为结果数组长度（最坏情况为O(n)）
//
// 使用场景:
//   - 需要获取一段时间内的CVE时使用
//   - 分析特定时间段内的安全漏洞
//
// 代码示例:
//
//	cveList := []string{"CVE-2020-1111", "CVE-2021-2222", "CVE-2022-3333"}
//	recentCves := cve.FilterCvesByYearRange(cveList, 2021, 2022)
//	// recentCves 为 ["CVE-2021-2222", "CVE-2022-3333"]
func FilterCvesByYearRange(cveSlice []string, startYear, endYear int) []string {
	var result []string

	for _, cve := range cveSlice {
		formattedCve := Format(cve)
		yearInt := ExtractCveYearAsInt(formattedCve)
		if yearInt >= startYear && yearInt <= endYear {
			result = append(result, formattedCve)
		}
	}

	return result
}

// GetRecentCves 获取最近n年的CVE
//
// 从CVE列表中获取最近几年内的CVE编号
//
// 参数:
//   - cveSlice: 需要筛选的CVE编号数组
//   - years: 最近几年的范围，整数，例如2表示最近两年
//
// 返回值:
//   - []string: 最近几年的CVE编号数组，已经过标准化格式处理
//     如果没有找到匹配项，则返回空数组
//
// 计算规则:
//
//	以当前年份为基准，获取(当前年份-years+1)到当前年份之间的所有CVE
//
// 示例:
//
//	假设当前年份为2023
//	输入: ["CVE-2020-1111", "CVE-2021-2222", "CVE-2022-3333", "CVE-2023-4444"], 2
//	输出: ["CVE-2022-3333", "CVE-2023-4444"] (2022和2023年的CVE)
//
//	输入: ["CVE-2020-1111", "CVE-2021-2222"], 1
//	输出: [] (空数组，因为没有2023年的CVE)
//
// 使用场景:
//   - 需要关注最近几年发布的CVE时使用
//   - 生成最新安全威胁报告
//
// 代码示例:
//
//	cveList := []string{"CVE-2020-1111", "CVE-2021-2222", "CVE-2022-3333", "CVE-2023-4444"}
//	recentTwoYears := cve.GetRecentCves(cveList, 2)
//	// 如果当前是2023年，recentTwoYears 为 ["CVE-2022-3333", "CVE-2023-4444"]
func GetRecentCves(cveSlice []string, years int) []string {
	currentYear := time.Now().Year()
	return FilterCvesByYearRange(cveSlice, currentYear-years+1, currentYear)
}

// RemoveDuplicateCves 移除重复的CVE编号
//
// 去除CVE列表中的重复项，保留唯一的CVE编号（不区分大小写）
//
// 参数:
//   - cveSlice: 包含可能重复CVE的数组
//
// 返回值:
//   - []string: 去重后的CVE编号数组，所有CVE均已标准化格式（大写）
//
// 处理规则:
//   - CVE比较时不区分大小写，例如"CVE-2022-1111"和"cve-2022-1111"被视为重复
//   - 只保留每个CVE的第一次出现
//   - 所有返回的CVE均为标准化格式（大写）
//
// 示例:
//
//	输入: ["CVE-2022-1111", "cve-2022-1111", "CVE-2022-2222"]
//	输出: ["CVE-2022-1111", "CVE-2022-2222"]
//
//	输入: ["CVE-2022-1111", "CVE-2022-1111", "CVE-2022-1111"]
//	输出: ["CVE-2022-1111"]
//
// 性能特性:
//   - 时间复杂度: O(n)，其中n为数组长度
//   - 空间复杂度: O(n)
//
// 使用场景:
//   - 合并多个来源的CVE列表并去重时使用
//   - 在处理大量CVE数据前进行预处理
//
// 代码示例:
//
//	cveList := []string{"CVE-2022-1111", "cve-2022-1111", "CVE-2022-2222"}
//	uniqueCves := cve.RemoveDuplicateCves(cveList)
//	// uniqueCves 为 ["CVE-2022-1111", "CVE-2022-2222"]
func RemoveDuplicateCves(cveSlice []string) []string {
	cveMap := make(map[string]struct{})
	var result []string

	for _, cve := range cveSlice {
		formattedCve := Format(cve)
		if _, exists := cveMap[formattedCve]; !exists {
			cveMap[formattedCve] = struct{}{}
			result = append(result, formattedCve)
		}
	}

	return result
}
