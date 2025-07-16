package cve

import (
	"sort"
)

// CompareByYear 根据CVE的年份比较大小
//
// 比较两个CVE编号的年份大小
//
// 参数:
//   - cveA: 第一个CVE编号
//   - cveB: 第二个CVE编号
//
// 返回值:
//   - int: 比较结果，具体规则如下
//   - 负数: cveA年份 < cveB年份（具体为 cveA年份 - cveB年份 的差值）
//   - 零: cveA年份 = cveB年份
//   - 正数: cveA年份 > cveB年份（具体为 cveA年份 - cveB年份 的差值）
//
// 示例:
//
//	输入: "CVE-2020-1111", "CVE-2022-2222" → 返回 -2
//	输入: "CVE-2022-1111", "CVE-2022-2222" → 返回 0
//	输入: "CVE-2023-1111", "CVE-2021-2222" → 返回 2
//
// 异常处理:
//   - 如果输入非有效CVE格式，会通过ExtractCveYearAsInt提取年份，无效CVE将被视为年份0
//
// 使用场景:
//   - CVE按年份排序时使用
//   - 比较两个CVE的发布时间（基于年份）
//
// 代码示例:
//
//	result := cve.CompareByYear("CVE-2020-1111", "CVE-2022-2222")
//	if result < 0 {
//	    fmt.Println("第一个CVE发布更早")
//	}
func CompareByYear(cveA, cveB string) int {
	return ExtractCveYearAsInt(cveA) - ExtractCveYearAsInt(cveB)
}

// SubByYear 把两个CVE根据年份相减
//
// 计算两个CVE编号的年份差值
//
// 参数:
//   - cveA: 第一个CVE编号
//   - cveB: 第二个CVE编号
//
// 返回值:
//   - int: 年份差值（cveA年份 - cveB年份）
//
// 示例:
//
//	输入: "CVE-2020-1111", "CVE-2022-2222" → 返回 -2
//	输入: "CVE-2022-1111", "CVE-2020-2222" → 返回 2
//	输入: "CVE-2022-1111", "CVE-2022-2222" → 返回 0
//
// 异常处理:
//   - 如果输入非有效CVE格式，会通过ExtractCveYearAsInt提取年份，无效CVE将被视为年份0
//
// 使用场景:
//   - 计算两个CVE之间的年份间隔
//   - 评估安全漏洞的时间分布
//
// 代码示例:
//
//	yearDiff := cve.SubByYear("CVE-2022-1111", "CVE-2020-2222")
//	// yearDiff 为 2，表示第一个CVE比第二个晚发布2年
func SubByYear(cveA, cveB string) int {
	return CompareByYear(cveA, cveB)
}

// CompareCves 根据CVE的年份和序列号比较大小
//
// 全面比较两个CVE编号的大小，首先比较年份，年份相同时比较序列号
//
// 参数:
//   - cveA: 第一个CVE编号
//   - cveB: 第二个CVE编号
//
// 返回值:
//   - int: 比较结果
//   - -1: cveA < cveB（cveA年份小于cveB，或年份相同但序列号较小）
//   - 0: cveA = cveB（年份和序列号完全相同）
//   - 1: cveA > cveB（cveA年份大于cveB，或年份相同但序列号较大）
//
// 示例:
//
//	输入: "CVE-2020-1111", "CVE-2022-2222" → 返回 -1 (不同年份)
//	输入: "CVE-2022-1111", "CVE-2022-2222" → 返回 -1 (相同年份，不同序列号)
//	输入: "CVE-2022-3333", "CVE-2022-2222" → 返回 1 (相同年份，序列号较大)
//	输入: "CVE-2022-2222", "CVE-2022-2222" → 返回 0 (完全相同)
//
// 异常处理:
//   - 如果输入的CVE格式不正确，将基于提取出的年份和序列号（可能为0）进行比较
//
// 使用场景:
//   - 需要完整排序CVE编号或比较两个CVE哪个更新时使用
//   - 对CVE列表按照发布顺序进行排序
//
// 代码示例:
//
//	result := cve.CompareCves("CVE-2022-1111", "CVE-2022-2222")
//	if result < 0 {
//	    fmt.Println("第一个CVE编号更小")
//	}
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

// SortCves 对CVE切片进行排序（按年份和序列号）并返回新的切片
//
// 将CVE列表按照年份和序列号排序，并统一格式，返回新的切片
//
// 参数:
//   - cveSlice: 需要排序的CVE编号数组
//
// 返回值:
//   - []string: 排序后的CVE编号数组，所有CVE均已标准化格式（大写）
//
// 排序规则:
//  1. 首先按年份升序排列（较早年份排在前面）
//  2. 年份相同时按序列号升序排列（较小序列号排在前面）
//
// 示例:
//
//	输入: ["CVE-2022-2222", "CVE-2020-1111", "CVE-2022-1111"]
//	输出: ["CVE-2020-1111", "CVE-2022-1111", "CVE-2022-2222"]
//
//	输入: ["cve-2022-2222", "CVE-2022-1111"]
//	输出: ["CVE-2022-1111", "CVE-2022-2222"] (注意格式已统一为大写)
//
// 性能特性:
//   - 时间复杂度: O(n log n)，其中n为数组长度
//   - 空间复杂度: O(n)，创建了新的数组存储结果
//
// 使用场景:
//   - 需要按时间顺序展示或处理一组CVE时使用
//   - 在生成漏洞报告时对CVE进行时间排序
//
// 代码示例:
//
//	cveList := []string{"CVE-2022-2222", "cve-2020-1111", "CVE-2022-1111"}
//	sortedList := cve.SortCves(cveList)
//	// sortedList 为 ["CVE-2020-1111", "CVE-2022-1111", "CVE-2022-2222"]
func SortCves(cveSlice []string) []string {
	result := make([]string, len(cveSlice))
	for i, cve := range cveSlice {
		result[i] = Format(cve)
	}

	sort.Slice(result, func(i, j int) bool {
		return CompareCves(result[i], result[j]) < 0
	})

	return result
}
