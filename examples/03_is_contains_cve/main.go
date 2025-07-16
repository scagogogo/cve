package main

import (
	"fmt"

	"github.com/scagogogo/cve"
)

func main() {
	// 示例1：检查包含一个CVE编号的文本
	// IsContainsCve函数用于检查字符串中是否包含CVE编号，不要求整个字符串就是CVE编号
	text1 := "系统受到CVE-2021-44228漏洞的影响，需要立即修复。"
	fmt.Printf("文本: %q\n", text1)
	fmt.Printf("是否包含CVE: %v\n\n", cve.IsContainsCve(text1))
	// 预期输出:
	// 文本: "系统受到CVE-2021-44228漏洞的影响，需要立即修复。"
	// 是否包含CVE: true

	// 示例2：检查包含多个CVE编号的文本
	// IsContainsCve函数只检查是否包含，不会提取出具体哪些CVE
	text2 := "安全公告：发现多个漏洞，包括CVE-2022-12345和CVE-2023-67890。"
	fmt.Printf("文本: %q\n", text2)
	fmt.Printf("是否包含CVE: %v\n\n", cve.IsContainsCve(text2))
	// 预期输出:
	// 文本: "安全公告：发现多个漏洞，包括CVE-2022-12345和CVE-2023-67890。"
	// 是否包含CVE: true

	// 示例3：检查不包含CVE编号的文本
	// 当文本中没有CVE编号时，返回false
	text3 := "这份文档中没有任何安全漏洞信息。"
	fmt.Printf("文本: %q\n", text3)
	fmt.Printf("是否包含CVE: %v\n\n", cve.IsContainsCve(text3))
	// 预期输出:
	// 文本: "这份文档中没有任何安全漏洞信息。"
	// 是否包含CVE: false

	// 示例4：检查包含小写cve编号的文本
	// IsContainsCve函数不区分大小写，能识别小写的cve编号
	text4 := "注意检查cve-2022-98765漏洞。"
	fmt.Printf("文本: %q\n", text4)
	fmt.Printf("是否包含CVE: %v\n", cve.IsContainsCve(text4))
	// 预期输出:
	// 文本: "注意检查cve-2022-98765漏洞。"
	// 是否包含CVE: true

	// 总结: IsContainsCve函数适用于从文章或报告中检测是否有提及CVE，
	// 与IsCve函数的区别在于它只检查包含关系，不要求整个字符串是CVE编号
}
