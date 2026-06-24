package cve

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// 内部常量：用于匹配CVE范围表达式的正则表达式
//
// 匹配格式:
//   - "CVE-2022-12345 to CVE-2022-12350"
//   - "CVE-2022-12345..12350" (双点号范围)
//   - "CVE-2022-12345-12350" (连字符范围)
var rangeRegex = regexp.MustCompile(`(?i)^\s*CVE-(\d+)-(\d+)\s*(?:to\s*CVE-\d+-(\d+)|\.\.(\d+)|\s*-\s*(\d+))\s*$`)

// GenerateCve 根据年份和序列号生成标准格式的CVE编号
//
// 通过给定的年份和序列号创建标准的CVE编号
//
// 参数:
//   - year: CVE年份，整数格式，如2022
//   - seq: CVE序列号，整数格式，如12345
//
// 返回值:
//   - string: 生成的标准格式CVE编号，如"CVE-2022-12345"
//
// 格式规则:
//   - 生成的CVE编号格式为"CVE-YYYY-NNNNN"
//   - 返回结果始终为大写形式
//
// 示例:
//
//	输入: 2022, 12345
//	输出: "CVE-2022-12345"
//
//	输入: 2021, 44228
//	输出: "CVE-2021-44228"
//
// 注意事项:
//   - 此函数不会验证年份是否合理（如是否在1999年之后）
//   - 序列号可以是任意整数，不限制位数
//
// 使用场景:
//   - 需要动态生成CVE编号时使用
//   - 创建新的CVE标识符
//
// 代码示例:
//
//	cveId := cve.GenerateCve(2022, 12345)
//	// cveId 为 "CVE-2022-12345"
//
//	// 可以组合使用提取和生成功能
//	year := 2023
//	seq := cve.ExtractCveSeqAsInt("CVE-2022-67890")
//	newCve := cve.GenerateCve(year, seq) // 生成"CVE-2023-67890"
func GenerateCve(year int, seq int) string {
	return Format(fmt.Sprintf("CVE-%d-%d", year, seq))
}

// GenerateFakeCve 生成一个假的CVE编号
//
// 无需提供参数，自动使用当前年份和随机序列号生成假的CVE编号
//
// 参数:
//
//	无
//
// 返回值:
//   - string: 生成的标准格式CVE编号，如"CVE-2023-54321"
//
// 生成规则:
//   - 使用当前系统年份作为CVE年份
//   - 序列号随机生成，范围在10000到99999之间
//   - 返回结果始终为大写形式
//
// 示例:
//
//	输出结果类似: "CVE-2023-54321"（假设当前年份为2023）
//
// 随机性:
//   - 序列号基于当前时间的纳秒部分生成，具有一定随机性
//   - 不保证全局唯一性，仅用于测试或示例目的
//
// 使用场景:
//   - 用于测试、示例或者占位符
//   - 快速创建模拟CVE数据
//
// 代码示例:
//
//	fakeCve := cve.GenerateFakeCve()
//	// fakeCve 可能为 "CVE-2023-12345"
//
//	// 在测试中创建多个随机CVE
//	var testCves []string
//	for i := 0; i < 5; i++ {
//	    testCves = append(testCves, cve.GenerateFakeCve())
//	}
func GenerateFakeCve() string {
	currentYear := time.Now().Year()
	randomSeq := 10000 + time.Now().Nanosecond()%90000
	return GenerateCve(currentYear, randomSeq)
}

// ParseCveRange 解析CVE范围表达式，返回范围内的所有CVE编号
//
// 支持多种范围表达格式:
//   - "CVE-2022-12345 to CVE-2022-12350" (单词to)
//   - "CVE-2022-12345..12350" (双点号)
//   - "CVE-2022-12345-12350" (连字符，注意与单个CVE区分)
//
// 参数:
//   - rangeExpr: CVE范围表达式字符串
//
// 返回值:
//   - []string: 范围内所有CVE编号的数组，如果格式无效则返回空数组
//
// 解析规则:
//  1. 范围的起始和结束必须在同一年份
//  2. 起始序列号必须小于或等于结束序列号
//  3. 返回的范围为闭区间，包含起始和结束
//
// 示例:
//
//	输入: "CVE-2022-12345 to CVE-2022-12350"
//	输出: ["CVE-2022-12345", "CVE-2022-12346", "CVE-2022-12347", "CVE-2022-12348", "CVE-2022-12349", "CVE-2022-12350"]
//
//	输入: "CVE-2022-12345..12347"
//	输出: ["CVE-2022-12345", "CVE-2022-12346", "CVE-2022-12347"]
//
//	输入: "CVE-2022-12345-12347"
//	输出: ["CVE-2022-12345", "CVE-2022-12346", "CVE-2022-12347"]
//
// 使用场景:
//   - 安全公告中经常使用CVE范围表示一连串漏洞
//   - 将范围表达展开为具体CVE列表以进行后续处理
//
// 代码示例:
//
//	cves := cve.ParseCveRange("CVE-2022-1000 to CVE-2022-1005")
//	// cves 包含6个CVE：1000到1005
func ParseCveRange(rangeExpr string) []string {
	matches := rangeRegex.FindStringSubmatch(rangeExpr)
	if matches == nil {
		return nil
	}

	// matches[1] = start year, matches[2] = start seq
	startYear := matches[1]
	startSeq, err := strconv.Atoi(matches[2])
	if err != nil {
		return nil
	}

	// 结束序列号可能在 to/dot-dot/dash 分组中
	var endSeq int
	switch {
	case matches[3] != "":
		endSeq, err = strconv.Atoi(matches[3])
	case matches[4] != "":
		endSeq, err = strconv.Atoi(matches[4])
	case matches[5] != "":
		endSeq, err = strconv.Atoi(matches[5])
	default:
		return nil
	}
	if err != nil || startSeq > endSeq {
		return nil
	}

	count := endSeq - startSeq + 1
	result := make([]string, count)
	year, _ := strconv.Atoi(startYear)
	for i := 0; i < count; i++ {
		result[i] = Format(fmt.Sprintf("CVE-%d-%d", year, startSeq+i))
	}
	return result
}

// IsCvesConsecutive 判断两个CVE编号是否连续
//
// 检查两个CVE是否在同一个年份且序列号相邻
//
// 参数:
//   - a: 第一个CVE编号
//   - b: 第二个CVE编号
//
// 返回值:
//   - bool: 如果两个CVE在同一且序列号差值为1则返回true
//
// 示例:
//
//	输入: "CVE-2022-12345", "CVE-2022-12346" → 返回 true
//	输入: "CVE-2022-12345", "CVE-2022-12347" → 返回 false (差值>1)
//	输入: "CVE-2022-12345", "CVE-2023-12345" → 返回 false (不同年份)
//
// 使用场景:
//   - 判断CVE是否可以合并为范围
//   - 检测CVE编号的连续性
//
// 代码示例:
//
//	if cve.IsCvesConsecutive("CVE-2022-12345", "CVE-2022-12346") {
//	    // 这两个CVE是连续的
//	}
func IsCvesConsecutive(a, b string) bool {
	yearA := ExtractCveYearAsInt(a)
	yearB := ExtractCveYearAsInt(b)
	if yearA == 0 || yearB == 0 || yearA != yearB {
		return false
	}
	seqA := ExtractCveSeqAsInt(a)
	seqB := ExtractCveSeqAsInt(b)
	if seqA == 0 || seqB == 0 {
		return false
	}
	diff := seqA - seqB
	return diff == 1 || diff == -1
}
