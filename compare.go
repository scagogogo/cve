package cve

import (
	"sort"
)

// CompareByYear 根据CVE的年份比较大小
//
// 比较两个CVE编号的年份大小
//
// 返回值:
//
//	负数: cveA年份 < cveB年份
//	零: cveA年份 = cveB年份
//	正数: cveA年份 > cveB年份
//
// 示例:
//
//	输入: "CVE-2020-1111", "CVE-2022-2222" → 返回 -2
//	输入: "CVE-2022-1111", "CVE-2022-2222" → 返回 0
//
// 使用场景:
//
//	CVE按年份排序时使用
func CompareByYear(cveA, cveB string) int {
	return ExtractCveYearAsInt(cveA) - ExtractCveYearAsInt(cveB)
}

// SubByYear 把两个CVE根据年份相减
//
// 计算两个CVE编号的年份差值
//
// 示例:
//
//	输入: "CVE-2020-1111", "CVE-2022-2222" → 返回 -2
//	输入: "CVE-2022-1111", "CVE-2020-2222" → 返回 2
//
// 使用场景:
//
//	计算两个CVE之间的年份间隔
func SubByYear(cveA, cveB string) int {
	return ExtractCveYearAsInt(cveA) - ExtractCveYearAsInt(cveB)
}

// CompareCves 根据CVE的年份和序列号比较大小
//
// 全面比较两个CVE编号的大小，首先比较年份，年份相同时比较序列号
//
// 返回值:
//
//	-1: cveA < cveB
//	 0: cveA = cveB
//	 1: cveA > cveB
//
// 示例:
//
//	输入: "CVE-2020-1111", "CVE-2022-2222" → 返回 -1 (不同年份)
//	输入: "CVE-2022-1111", "CVE-2022-2222" → 返回 -1 (相同年份，不同序列号)
//	输入: "CVE-2022-2222", "CVE-2022-2222" → 返回 0 (完全相同)
//
// 使用场景:
//
//	需要完整排序CVE编号或比较两个CVE哪个更新时使用
func CompareCves(cveA, cveB string) int {
	yearComp := CompareByYear(cveA, cveB)
	if yearComp != 0 {
		if yearComp < 0 {
			return -1
		}
		return 1
	}

	seqA := ExtractCveSeqAsInt(cveA)
	seqB := ExtractCveSeqAsInt(cveB)

	if seqA < seqB {
		return -1
	} else if seqA > seqB {
		return 1
	}
	return 0
}

// SortedCves 对CVE切片进行排序（按年份和序列号）并返回新的切片
//
// 将CVE列表按照年份和序列号排序，并统一格式，返回新的切片
//
// 示例:
//
//	输入: ["CVE-2022-2222", "CVE-2020-1111", "CVE-2022-1111"]
//	输出: ["CVE-2020-1111", "CVE-2022-1111", "CVE-2022-2222"]
//
// 使用场景:
//
//	需要按时间顺序展示或处理一组CVE时使用
//
// 代码示例:
//
//	cveList := []string{"CVE-2022-2222", "cve-2020-1111", "CVE-2022-1111"}
//	sortedList := cve.SortedCves(cveList)
//	// sortedList 为 ["CVE-2020-1111", "CVE-2022-1111", "CVE-2022-2222"]
func SortedCves(cveSlice []string) []string {
	result := make([]string, len(cveSlice))
	for i, cve := range cveSlice {
		result[i] = Format(cve)
	}

	sort.Slice(result, func(i, j int) bool {
		return CompareCves(result[i], result[j]) < 0
	})

	return result
}

// SortCves 是 SortedCves 的别名，为保持向后兼容
//
// 注意: 此函数不会修改原始切片，而是返回一个新的排序后的切片
func SortCves(cveSlice []string) []string {
	return SortedCves(cveSlice)
}
