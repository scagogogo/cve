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
// 参数:
//   - text: 需要提取CVE的文本内容，可以是任意字符串
//
// 返回值:
//   - []string: 所有提取到的CVE编号数组，标准格式为大写，如果没有找到则返回空数组
//
// 示例:
//
//	输入: "系统受到CVE-2021-44228和cve-2022-12345的影响"
//	输出: ["CVE-2021-44228", "CVE-2022-12345"]
//
//	输入: "本文档不包含任何CVE编号"
//	输出: [] (空数组)
//
// 性能特性:
//   - 空间复杂度：O(n)，其中n为提取的CVE数量
//   - 时间复杂度：O(m)，其中m为文本长度
//
// 使用场景:
//   - 从安全公告或漏洞报告中提取所有相关的CVE编号
//   - 批量处理安全通告以提取受影响的CVE列表
//
// 代码示例:
//
//	report := "系统受到CVE-2021-44228和CVE-2022-12345的影响"
//	cveList := cve.ExtractCve(report)
//	// cveList 包含 ["CVE-2021-44228", "CVE-2022-12345"]
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
// 参数:
//   - text: 需要提取CVE的文本内容
//
// 返回值:
//   - string: 第一个找到的CVE编号，格式化为标准大写形式
//     如果未找到任何CVE，则返回空字符串""
//
// 示例:
//
//	输入: "系统受到CVE-2021-44228和CVE-2022-12345的影响"
//	输出: "CVE-2021-44228"
//
//	输入: "本文档不包含任何CVE编号"
//	输出: "" (空字符串)
//
// 使用场景:
//   - 当只需要获取文本中第一个提到的CVE时使用
//   - 快速识别安全通告的主要CVE编号
//
// 代码示例:
//
//	report := "Log4j漏洞(CVE-2021-44228)非常严重"
//	firstCve := cve.ExtractFirstCve(report)
//	// firstCve 为 "CVE-2021-44228"
func ExtractFirstCve(text string) string {
	s := cveRegex.FindString(text)
	return Format(s)
}

// ExtractLastCve 从字符串中抽取出最后一个CVE编号
//
// 提取文本中出现的最后一个CVE编号
//
// 参数:
//   - text: 需要提取CVE的文本内容
//
// 返回值:
//   - string: 最后一个找到的CVE编号，格式化为标准大写形式
//     如果未找到任何CVE，则返回空字符串""
//
// 示例:
//
//	输入: "系统受到CVE-2021-44228和CVE-2022-12345的影响"
//	输出: "CVE-2022-12345"
//
//	输入: "本文档不包含任何CVE编号"
//	输出: "" (空字符串)
//
// 使用场景:
//   - 当需要获取文本中最后提到的CVE时使用
//   - 提取更新日志中最新的CVE编号
//
// 代码示例:
//
//	changelog := "修复了CVE-2021-44228和CVE-2022-12345漏洞"
//	lastCve := cve.ExtractLastCve(changelog)
//	// lastCve 为 "CVE-2022-12345"
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
// 参数:
//   - cve: CVE编号字符串
//
// 返回值:
//   - string: 提取的年份字符串，如"2022"
//     如果输入的不是有效的CVE格式，则返回空字符串""
//
// 示例:
//
//	输入: "CVE-2022-12345"
//	输出: "2022"
//
//	输入: "不是CVE格式"
//	输出: "" (空字符串)
//
// 使用场景:
//   - 需要对CVE按年份进行分类时使用
//   - 生成按年份组织的安全报告
//
// 代码示例:
//
//	year := cve.ExtractCveYear("CVE-2022-12345")
//	// year 为 "2022"
func ExtractCveYear(cve string) string {
	year, _ := Split(cve)
	return year
}

// ExtractCveYearAsInt 抽取CVE中的年份，解析为int返回
//
// 提取CVE编号中的年份并转换为整数类型
//
// 参数:
//   - cve: CVE编号字符串
//
// 返回值:
//   - int: 提取的年份整数值，如2022
//     如果输入的不是有效的CVE格式或年份解析失败，则返回0
//
// 示例:
//
//	输入: "CVE-2022-12345"
//	输出: 2022 (整数类型)
//
//	输入: "不是CVE格式"
//	输出: 0
//
// 异常处理:
//   - 如果CVE格式不正确或年份部分不是有效数字，返回0
//
// 使用场景:
//   - 需要对CVE年份进行数值计算或比较时使用
//   - 计算CVE发布的年限
//
// 代码示例:
//
//	yearInt := cve.ExtractCveYearAsInt("CVE-2022-12345")
//	currentYear := time.Now().Year()
//	age := currentYear - yearInt
//	// 计算CVE发布至今的年数
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
// 参数:
//   - cve: CVE编号字符串
//
// 返回值:
//   - string: 提取的序列号字符串，如"12345"
//     如果输入的不是有效的CVE格式，则返回空字符串""
//
// 示例:
//
//	输入: "CVE-2022-12345"
//	输出: "12345"
//
//	输入: "不是CVE格式"
//	输出: "" (空字符串)
//
// 使用场景:
//   - 需要单独处理CVE序列号时使用
//   - 获取特定范围内的CVE序列号
//
// 代码示例:
//
//	seq := cve.ExtractCveSeq("CVE-2022-12345")
//	// seq 为 "12345"
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
// 参数:
//   - cve: CVE编号字符串
//
// 返回值:
//   - int: 提取的序列号整数值，如12345
//     如果输入的不是有效的CVE格式或序列号解析失败，则返回0
//
// 示例:
//
//	输入: "CVE-2022-12345"
//	输出: 12345 (整数类型)
//
//	输入: "不是CVE格式"
//	输出: 0
//
//	输入: "CVE-2022-ABC" (无效序列号)
//	输出: 0
//
// 异常处理:
//   - 如果CVE格式不正确或序列号部分不是有效数字，返回0
//
// 使用场景:
//   - 需要对CVE序列号进行数值计算或比较时使用
//   - 过滤指定序列号范围的CVE
//
// 代码示例:
//
//	seqInt := cve.ExtractCveSeqAsInt("CVE-2022-12345")
//	if seqInt > 10000 {
//	    // 处理序列号大于10000的CVE
//	}
func ExtractCveSeqAsInt(cve string) int {
	seq := ExtractCveSeq(cve)
	i, _ := strconv.Atoi(seq)
	return i
}

// FilterCvesByPattern 根据通配符模式筛选CVE列表
//
// 支持简单的通配符模式匹配：
//   - "*" 匹配任意内容
//   - "CVE-2022-*" 匹配2022年的所有CVE
//   - "CVE-*-1234" 匹配序列号为1234的所有年份的CVE
//   - "CVE-2022-1*" 匹配2022年以1开头的序列号
//
// 参数:
//   - cveSlice: 需要筛选的CVE编号数组
//   - pattern: 通配符模式，自动格式化为大写
//
// 返回值:
//   - []string: 匹配模式的所有CVE编号，已格式化和排序
//
// 示例:
//
//	输入: ["CVE-2022-1111", "CVE-2022-2222", "CVE-2023-1111"], "CVE-2022-*"
//	输出: ["CVE-2022-1111", "CVE-2022-2222"]
//
//	输入: ["CVE-2022-1111", "CVE-2021-1111", "CVE-2023-2222"], "CVE-*-1111"
//	输出: ["CVE-2021-1111", "CVE-2022-1111"]
//
// 使用场景:
//   - 通过通配符快速筛选CVE
//   - 构建灵活的CVE查询功能
//
// 代码示例:
//
//	cves2022 := cve.FilterCvesByPattern(cveList, "CVE-2022-*")
//	cve1111 := cve.FilterCvesByPattern(cveList, "CVE-*-1111")
func FilterCvesByPattern(cveSlice []string, pattern string) []string {
	pattern = Format(pattern)
	// 将通配符模式转换为正则表达式
	patternParts := []rune(pattern)
	var regexParts []rune
	for _, ch := range patternParts {
		switch ch {
		case '*':
			regexParts = append(regexParts, []rune(".*")...)
		case '.', '+', '(', ')', '[', ']', '{', '}', '\\', '^', '$', '|':
			// 转义正则特殊字符
			regexParts = append(regexParts, '\\', ch)
		default:
			regexParts = append(regexParts, ch)
		}
	}

	// 转换逻辑已将所有正则特殊字符转义为字面字符，* 转为 .*，
	// 产出的正则必然可编译，故用 MustCompile 表达此不变式。
	regex := regexp.MustCompile(string(regexParts))

	var result []string
	for _, cve := range cveSlice {
		formatted := Format(cve)
		if regex.MatchString(formatted) {
			result = append(result, formatted)
		}
	}

	return SortCves(result)
}
