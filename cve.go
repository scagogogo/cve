package cve

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Format 把CVE的格式统一化
//
// 将CVE编号转换为标准大写格式并移除前后空格
//
// 示例:
//
//	输入: " cve-2022-12345 "
//	输出: "CVE-2022-12345"
//
// 使用场景:
//
//	在比较或存储CVE编号前进行标准化
func Format(cve string) string {
	return strings.ToUpper(strings.TrimSpace(cve))
}

// IsCve 判断字符串是否是有效的CVE格式
//
// 验证字符串是否完全符合CVE格式（允许两侧有空白字符）
//
// 示例:
//
//	输入: "CVE-2022-12345" 或 " CVE-2022-12345 " → 返回 true
//	输入: "包含CVE-2022-12345的文本" → 返回 false
//
// 使用场景:
//
//	验证用户输入的字符串是否为有效的CVE编号
func IsCve(text string) bool {
	// 允许两侧有空白字符，但是不允许有除空白字符以外的其他字符
	return regexp.MustCompile(`(?i)^\s*CVE-\d+-\d+\s*$`).MatchString(text)
}

// IsContainsCve 判断字符串是否包含CVE
//
// 检查字符串中是否包含CVE格式的内容
//
// 示例:
//
//	输入: "这个漏洞的编号是CVE-2022-12345" → 返回 true
//	输入: "这个文本不包含任何CVE" → 返回 false
//
// 使用场景:
//
//	从文章或报告中检测是否有提及CVE
func IsContainsCve(text string) bool {
	return regexp.MustCompile(`(?i)CVE-\d+-\d+`).MatchString(text)
}

// IsCveYearOk 判断CVE的年份是否在合理的时间范围内
//
// 验证CVE年份是否在1970年之后且不超过当前年份
//
// 示例:
//
//	输入: "CVE-2022-12345" → 当前是2023年时返回 true
//	输入: "CVE-2030-12345" → 当前是2023年时返回 false (2030 > 2023)
//
// 使用场景:
//
//	验证CVE年份的有效性
func IsCveYearOk(cve string) bool {
	return IsCveYearOkWithCutoff(cve, 0)
}

// IsCveYearOkWithCutoff 判断CVE的年份是否在合理的时间范围内，可设置偏移量
//
// 验证CVE年份是否在1970年之后且不超过当前年份加上cutoff偏移值
//
// 参数:
//
//	cve: CVE编号
//	cutoff: 允许的年份偏移量，用于处理未来年份
//
// 示例:
//
//	输入: "CVE-2022-12345", 5 → 当前是2023年时返回 true
//	输入: "CVE-2030-12345", 5 → 当前是2023年时返回 false (2030 > 2023+5)
//
// 使用场景:
//
//	验证CVE年份的有效性，允许一定的未来年份偏移
func IsCveYearOkWithCutoff(cve string, cutoff int) bool {
	year := ExtractCveYearAsInt(cve)
	return year >= 1970 && time.Now().Year()-year <= cutoff
}

// ExtractCve 从字符串中抽取出所有CVE编号
//
// 提取文本中所有CVE编号并标准化格式
//
// 示例:
//
//	输入: "系统受到CVE-2021-44228和cve-2022-12345的影响"
//	输出: ["CVE-2021-44228", "CVE-2022-12345"]
//
// 使用场景:
//
//	从安全公告或漏洞报告中提取所有相关的CVE编号
func ExtractCve(text string) []string {
	slice := regexp.MustCompile(`(?i)(CVE-\d+-\d+)`).FindAllString(text, -1)
	for i, cve := range slice {
		slice[i] = Format(cve)
	}
	return slice
}

// ExtractFirstCve 从字符串中抽取出第一个CVE编号
//
// 提取文本中出现的第一个CVE编号
//
// 示例:
//
//	输入: "系统受到CVE-2021-44228和CVE-2022-12345的影响"
//	输出: "CVE-2021-44228"
//
// 使用场景:
//
//	当只需要获取文本中第一个提到的CVE时使用
func ExtractFirstCve(text string) string {
	//slice := ExtractCve(text)
	//if len(slice) == 0 {
	//	return ""
	//}
	//return slice[0]

	s := regexp.MustCompile(`(?i)(CVE-\d+-\d+)`).FindString(text)
	return Format(s)
}

// ExtractLastCve 从字符串中抽取出最后一个CVE编号
//
// 提取文本中出现的最后一个CVE编号
//
// 示例:
//
//	输入: "系统受到CVE-2021-44228和CVE-2022-12345的影响"
//	输出: "CVE-2022-12345"
//
// 使用场景:
//
//	当需要获取文本中最后提到的CVE时使用
func ExtractLastCve(text string) string {
	slice := ExtractCve(text)
	if len(slice) == 0 {
		return ""
	}
	return slice[len(slice)-1]
}

// Split 把CVE分割成年份和编号两部分
//
// 将CVE编号拆分为年份和序列号两个部分
//
// 示例:
//
//	输入: "CVE-2022-12345"
//	输出: year="2022", seq="12345"
//
// 使用场景:
//
//	需要单独处理CVE的年份或序列号部分时使用
func Split(cve string) (year string, seq string) {
	cve = Format(cve)
	split := strings.Split(cve, "-")
	if len(split) != 3 {
		return
	}
	return split[1], split[2]
}

// ExtractCveYear 从CVE中抽取出年份
//
// 提取CVE编号中的年份部分（字符串格式）
//
// 示例:
//
//	输入: "CVE-2022-12345"
//	输出: "2022"
//
// 使用场景:
//
//	需要对CVE按年份进行分类时使用
func ExtractCveYear(cve string) string {
	year, _ := Split(cve)
	return year
}

// ExtractCveYearAsInt 抽取CVE中的年份，解析为int返回
//
// 提取CVE编号中的年份并转换为整数类型
//
// 示例:
//
//	输入: "CVE-2022-12345"
//	输出: 2022 (整数类型)
//
// 使用场景:
//
//	需要对CVE年份进行数值计算或比较时使用
func ExtractCveYearAsInt(cve string) int {
	year := ExtractCveYear(cve)
	i, _ := strconv.Atoi(year)
	return i
}

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

// ExtractCveSeq 从CVE中抽取出序列号
//
// 提取CVE编号中的序列号部分（字符串格式）
//
// 示例:
//
//	输入: "CVE-2022-12345"
//	输出: "12345"
//	输入: "不是CVE格式"
//	输出: "" (空字符串)
//
// 使用场景:
//
//	需要单独处理CVE序列号时使用
func ExtractCveSeq(cve string) string {
	if !IsCve(cve) {
		return ""
	}
	_, seq := Split(cve)
	return seq
}

// ExtractCveSeqAsInt 抽取CVE中的序列号，解析为int返回
//
// 提取CVE编号中的序列号并转换为整数类型
//
// 示例:
//
//	输入: "CVE-2022-12345"
//	输出: 12345 (整数类型)
//	输入: "不是CVE格式"
//	输出: 0
//
// 使用场景:
//
//	需要对CVE序列号进行数值计算或比较时使用
func ExtractCveSeqAsInt(cve string) int {
	seq := ExtractCveSeq(cve)
	i, _ := strconv.Atoi(seq)
	return i
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

// GenerateCve 根据年份和序列号生成标准格式的CVE编号
//
// 通过给定的年份和序列号创建标准的CVE编号
//
// 示例:
//
//	输入: 2022, 12345
//	输出: "CVE-2022-12345"
//
// 使用场景:
//
//	需要动态生成CVE编号时使用
//
// 代码示例:
//
//	cveId := cve.GenerateCve(2022, 12345)
//	// cveId 为 "CVE-2022-12345"
func GenerateCve(year int, seq int) string {
	return Format(fmt.Sprintf("CVE-%d-%d", year, seq))
}

// ValidateCve 全面验证CVE编号的合法性
//
// 检查CVE编号是否符合格式要求并具有合理的年份和序列号
//
// 示例:
//
//	输入: "CVE-2022-12345" → 当前年份为2023时返回 true
//	输入: "CVE-1960-12345" → 返回 false (年份 < 1970)
//	输入: "CVE-2030-12345" → 当前年份为2023时返回 false (年份 > 当前年份)
//	输入: "CVE-2022-ABC" → 返回 false (序列号不是数字)
//
// 使用场景:
//
//	验证用户输入的CVE编号是否有效
//
// 代码示例:
//
//	isValid := cve.ValidateCve("CVE-2022-12345")
//	if isValid {
//	    // 进行处理...
//	}
func ValidateCve(cve string) bool {
	if !IsCve(cve) {
		return false
	}

	year, seq := Split(cve)
	yearInt, yearErr := strconv.Atoi(year)
	seqInt, seqErr := strconv.Atoi(seq)

	if yearErr != nil || seqErr != nil {
		return false
	}

	// 基础验证规则：年份在1970至今，序列号为正整数
	return yearInt >= 1970 && yearInt <= time.Now().Year() && seqInt > 0
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

// GenerateFakeCve 生成一个假的CVE编号
//
// 无需提供参数，自动使用当前年份和随机序列号生成假的CVE编号
//
// 示例:
//
//	生成结果类似: "CVE-2023-54321"（假设当前年份为2023）
//
// 使用场景:
//
//	用于测试、示例或者占位符
//
// 代码示例:
//
//	fakeCve := cve.GenerateFakeCve()
//	// fakeCve 可能为 "CVE-2023-12345"
func GenerateFakeCve() string {
	currentYear := time.Now().Year()
	randomSeq := 10000 + time.Now().Nanosecond()%90000
	return GenerateCve(currentYear, randomSeq)
}
