package cve

import (
	"regexp"
	"strconv"
)

// 内部常量：用于匹配CVE格式的正则表达式
var cveRegex = regexp.MustCompile(`(?i)(CVE-\d+-\d+)`)

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
	slice := cveRegex.FindAllString(text, -1)
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
	s := cveRegex.FindString(text)
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
//	输入: "不是CVE格式"
//	输出: 0
//
// 使用场景:
//
//	需要对CVE年份进行数值计算或比较时使用
func ExtractCveYearAsInt(cve string) int {
	if !IsCve(cve) {
		return 0
	}
	year := ExtractCveYear(cve)
	i, _ := strconv.Atoi(year)
	return i
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
