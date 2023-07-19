package cve

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Format 把CVE的格式统一化，其实就是转大写、去除两侧空格啥的
func Format(cve string) string {
	return strings.ToUpper(strings.TrimSpace(cve))
}

// IsCve 判断字符串是否是CVE
func IsCve(text string) bool {
	// 允许两侧有空白字符，但是不允许有除空白字符以外的其他字符
	return regexp.MustCompile("(?i)^\\s*CVE-\\d+-\\d+\\s*$").MatchString(text)
}

// IsContainsCve 判断字符串是否包含CVE
func IsContainsCve(text string) bool {
	return regexp.MustCompile("(?i)CVE-\\d+-\\d+").MatchString(text)
}

// IsCveYearOk 判断CVE的年份是否在合理的时间范围内
// 允许年份适当的向后偏移
func IsCveYearOk(cve string, cutoff int) bool {
	year := ExtractCveYearAsInt(cve)
	return year >= 1970 && time.Now().Year()-year <= cutoff
}

// ExtractCve 从字符串中抽取出CVE编号
func ExtractCve(text string) []string {
	slice := regexp.MustCompile("(?i)(CVE-\\d+-\\d+)").FindAllString(text, -1)
	for i, cve := range slice {
		slice[i] = Format(cve)
	}
	return slice
}

// ExtractFirstCve 从字符串中抽取出第一个CVE编号
func ExtractFirstCve(text string) string {
	//slice := ExtractCve(text)
	//if len(slice) == 0 {
	//	return ""
	//}
	//return slice[0]

	s := regexp.MustCompile("(?i)(CVE-\\d+-\\d+)").FindString(text)
	return Format(s)
}

// ExtractLastCve 从字符串中抽取出最后一个CVE编号
func ExtractLastCve(text string) string {
	slice := ExtractCve(text)
	if len(slice) == 0 {
		return ""
	}
	return slice[len(slice)-1]
}

// Split 把CVE分割成年份和编号两部分，比如：
// CVE-2022-100087
// return:
// year is 2022
// seq is 100087
func Split(cve string) (year string, seq string) {
	cve = Format(cve)
	split := strings.Split(cve, "-")
	if len(split) != 3 {
		return
	}
	return split[1], split[2]
}

// ExtractCveYear 从CVE中抽取出年份
func ExtractCveYear(cve string) string {
	year, _ := Split(cve)
	return year
}

// ExtractCveYearAsInt 抽取CVE中的年份，解析为int返回
func ExtractCveYearAsInt(cve string) int {
	year := ExtractCveYear(cve)
	i, _ := strconv.Atoi(year)
	return i
}

// GroupByYear 把一组CVE按照年份分组
func GroupByYear(cveSlice []string) map[string][]string {
	groupMap := make(map[string][]string, 0)
	for _, cve := range cveSlice {
		year := ExtractCveYear(cve)
		groupMap[year] = append(groupMap[year], Format(cve))
	}
	return groupMap
}

// CompareByYear 根据CVE的年份比较大小，年份相等返回0
func CompareByYear(cveA, cveB string) int {
	return ExtractCveYearAsInt(cveA) - ExtractCveYearAsInt(cveB)
}

// SubByYear 把两个CVE根据年份相减
func SubByYear(cveA, cveB string) int {
	return ExtractCveYearAsInt(cveA) - ExtractCveYearAsInt(cveB)
}
