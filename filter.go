package cve

import (
	"strconv"
	"time"
)

// GroupByYear 把一组CVE按照年份分组
//
// 将CVE列表按照年份进行分组
//
// 示例:
//
//	输入: ["CVE-2021-1111", "CVE-2022-2222", "CVE-2021-3333"]
//	输出: {
//	  "2021": ["CVE-2021-1111", "CVE-2021-3333"],
//	  "2022": ["CVE-2022-2222"]
//	}
//
// 使用场景:
//
//	按年份组织和展示多个CVE，例如生成年度漏洞报告
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
// 示例:
//
//	输入: ["CVE-2021-1111", "CVE-2022-2222", "CVE-2021-3333"], 2021
//	输出: ["CVE-2021-1111", "CVE-2021-3333"]
//
// 使用场景:
//
//	需要获取特定年份的CVE时使用
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
// 示例:
//
//	输入: ["CVE-2020-1111", "CVE-2021-2222", "CVE-2022-3333"], 2021, 2022
//	输出: ["CVE-2021-2222", "CVE-2022-3333"]
//
// 使用场景:
//
//	需要获取一段时间内的CVE时使用
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
// 示例:
//
//	假设当前年份为2023
//	输入: ["CVE-2020-1111", "CVE-2021-2222", "CVE-2022-3333", "CVE-2023-4444"], 2
//	输出: ["CVE-2022-3333", "CVE-2023-4444"] (2022和2023年的CVE)
//
// 使用场景:
//
//	需要关注最近几年发布的CVE时使用
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
// 示例:
//
//	输入: ["CVE-2022-1111", "cve-2022-1111", "CVE-2022-2222"]
//	输出: ["CVE-2022-1111", "CVE-2022-2222"]
//
// 使用场景:
//
//	合并多个来源的CVE列表并去重时使用
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
