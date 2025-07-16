package main

import (
	"fmt"

	"github.com/scagogogo/cve"
)

func main() {
	// 示例CVE编号
	cveID := "CVE-2022-12345"
	fmt.Printf("CVE编号: %s\n\n", cveID)

	// 提取年份（字符串形式）
	year := cve.ExtractCveYear(cveID)
	fmt.Printf("年份(字符串): %s\n", year)

	// 提取年份（整数形式）
	yearInt := cve.ExtractCveYearAsInt(cveID)
	fmt.Printf("年份(整数): %d\n\n", yearInt)

	// 提取序列号（字符串形式）
	seq := cve.ExtractCveSeq(cveID)
	fmt.Printf("序列号(字符串): %s\n", seq)

	// 提取序列号（整数形式）
	seqInt := cve.ExtractCveSeqAsInt(cveID)
	fmt.Printf("序列号(整数): %d\n\n", seqInt)

	// 演示处理无效输入
	invalidCve := "这不是CVE格式"
	fmt.Printf("无效输入: %s\n", invalidCve)
	fmt.Printf("无效输入的年份(字符串): %q\n", cve.ExtractCveYear(invalidCve))
	fmt.Printf("无效输入的年份(整数): %d\n", cve.ExtractCveYearAsInt(invalidCve))
	fmt.Printf("无效输入的序列号(字符串): %q\n", cve.ExtractCveSeq(invalidCve))
	fmt.Printf("无效输入的序列号(整数): %d\n", cve.ExtractCveSeqAsInt(invalidCve))

	// 使用Split函数作为替代方法
	fmt.Println("\n使用Split函数:")
	splitYear, splitSeq := cve.Split(cveID)
	fmt.Printf("Split解析的年份: %s\n", splitYear)
	fmt.Printf("Split解析的序列号: %s\n", splitSeq)
}
