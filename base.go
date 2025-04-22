package cve

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 内部变量：用于验证和匹配CVE格式的正则表达式
var (
	// 精确匹配CVE格式（允许两侧有空白字符）
	exactCveRegex = regexp.MustCompile(`(?i)^\s*CVE-\d+-\d+\s*$`)
	// 在文本中匹配CVE格式
	containsCveRegex = regexp.MustCompile(`(?i)CVE-\d+-\d+`)
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
	return exactCveRegex.MatchString(text)
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
	return containsCveRegex.MatchString(text)
}

// 内部函数：从CVE中提取年份并转为整数
func extractYear(cve string) int {
	cve = Format(cve)
	split := strings.Split(cve, "-")
	if len(split) != 3 {
		return 0
	}
	year, _ := strconv.Atoi(split[1])
	return year
}

// IsCveYearOk 判断CVE的年份是否在合理的时间范围内
//
// 验证CVE年份是否在1999年之后且不超过当前年份
//
// 示例:
//
//	输入: "CVE-2022-12345" → 当前是2023年时返回 true
//	输入: "CVE-1998-12345" → 返回 false (1998 < 1999)
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
// 验证CVE年份是否在1999年之后且不超过当前年份加上cutoff偏移值
//
// 参数:
//
//	cve: CVE编号
//	cutoff: 允许的年份偏移量，用于处理未来年份
//
// 示例:
//
//	输入: "CVE-2022-12345", 5 → 当前是2023年时返回 true
//	输入: "CVE-1998-12345", 5 → 返回 false (1998 < 1999)
//	输入: "CVE-2030-12345", 5 → 当前是2023年时返回 true (2030 <= 2023+5)
//
// 使用场景:
//
//	验证CVE年份的有效性，允许一定的未来年份偏移
func IsCveYearOkWithCutoff(cve string, cutoff int) bool {
	year := extractYear(cve)
	return year >= 1999 && year <= time.Now().Year()+cutoff
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

// ValidateCve 全面验证CVE编号的合法性
//
// 检查CVE编号是否符合格式要求并具有合理的年份和序列号
//
// 示例:
//
//	输入: "CVE-2022-12345" → 当前年份为2023时返回 true
//	输入: "CVE-1998-12345" → 返回 false (年份 < 1999)
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

	// 基础验证规则：年份在1999至今，序列号为正整数
	return yearInt >= 1999 && yearInt <= time.Now().Year() && seqInt > 0
}
