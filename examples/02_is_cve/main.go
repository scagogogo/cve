package main

import (
	"fmt"

	"github.com/scagogogo/cve"
)

func main() {
	// 示例1：检查标准格式的CVE编号
	// IsCve函数用于检查字符串是否是标准格式的CVE编号（形如CVE-YYYY-NNNNN）
	input1 := "CVE-2022-12345"
	fmt.Printf("输入: %q, 是否为CVE: %v\n", input1, cve.IsCve(input1))
	// 预期输出:
	// 输入: "CVE-2022-12345", 是否为CVE: true

	// 示例2：检查包含空白字符的CVE编号
	// IsCve函数允许CVE编号两侧有空白字符
	input2 := " CVE-2021-44228 "
	fmt.Printf("输入: %q, 是否为CVE: %v\n", input2, cve.IsCve(input2))
	// 预期输出:
	// 输入: " CVE-2021-44228 ", 是否为CVE: true

	// 示例3：检查非标准格式
	// IsCve函数要求整个字符串都是CVE编号，而不只是包含CVE编号
	input3 := "包含CVE-2023-9999的文本"
	fmt.Printf("输入: %q, 是否为CVE: %v\n", input3, cve.IsCve(input3))
	// 预期输出:
	// 输入: "包含CVE-2023-9999的文本", 是否为CVE: false

	// 示例4：检查错误格式
	// IsCve函数检查格式是否严格符合CVE-YYYY-NNNNN
	input4 := "CVE2022-12345" // 缺少连字符
	fmt.Printf("输入: %q, 是否为CVE: %v\n", input4, cve.IsCve(input4))
	// 预期输出:
	// 输入: "CVE2022-12345", 是否为CVE: false

	// 总结: IsCve函数用于严格验证字符串是否为独立的CVE编号，
	// 常用于验证用户输入的字符串是否为有效的CVE编号，
	// 与IsContainsCve不同，它要求整个字符串就是一个CVE编号
}
