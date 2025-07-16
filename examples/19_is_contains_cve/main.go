package main

import (
	"fmt"

	"github.com/scagogogo/cve"
)

func main() {
	fmt.Println("检测文本中是否包含CVE示例")
	// 预期输出:
	// 检测文本中是否包含CVE示例

	// 包含CVE的文本示例
	text1 := "这是一个包含CVE-2021-44228漏洞的文本。"
	fmt.Printf("文本1: %s\n", text1)
	containsCve1 := cve.IsContainsCve(text1)
	fmt.Printf("检测结果: %v\n", containsCve1)
	// 预期输出:
	// 文本1: 这是一个包含CVE-2021-44228漏洞的文本。
	// 检测结果: true

	// 不包含CVE的文本示例
	text2 := "这是一个不包含任何CVE编号的普通文本。"
	fmt.Printf("\n文本2: %s\n", text2)
	containsCve2 := cve.IsContainsCve(text2)
	fmt.Printf("检测结果: %v\n", containsCve2)
	// 预期输出:
	// 文本2: 这是一个不包含任何CVE编号的普通文本。
	// 检测结果: false

	// 包含多个CVE的文本示例
	text3 := "这个文本包含多个CVE：CVE-2022-22965和CVE-2021-45046。"
	fmt.Printf("\n文本3: %s\n", text3)
	containsCve3 := cve.IsContainsCve(text3)
	fmt.Printf("检测结果: %v\n", containsCve3)
	// 预期输出:
	// 文本3: 这个文本包含多个CVE：CVE-2022-22965和CVE-2021-45046。
	// 检测结果: true

	// 包含不规范CVE格式的文本示例
	text4 := "这个文本包含不规范格式的cve-2022-1234和CVE2023-5678。"
	fmt.Printf("\n文本4: %s\n", text4)
	containsCve4 := cve.IsContainsCve(text4)
	fmt.Printf("检测结果: %v\n", containsCve4)
	// 预期输出:
	// 文本4: 这个文本包含不规范格式的cve-2022-1234和CVE2023-5678。
	// 检测结果: true

	// 提取文本中的所有CVE
	fmt.Printf("\n提取文本3中的所有CVE:\n")
	cves := cve.ExtractCve(text3)
	for i, id := range cves {
		fmt.Printf("[%d] %s\n", i+1, id)
	}
	// 预期输出:
	// 提取文本3中的所有CVE:
	// [1] CVE-2022-22965
	// [2] CVE-2021-45046

	// 应用场景示例
	fmt.Println("\n应用场景示例:")
	fmt.Println("1. 安全公告分析：自动扫描安全公告中提到的CVE")
	fmt.Println("2. 漏洞跟踪：从各种文档中提取CVE进行追踪管理")
	fmt.Println("3. 威胁情报分析：检测威胁情报报告中的CVE编号")
	// 预期输出:
	// 应用场景示例:
	// 1. 安全公告分析：自动扫描安全公告中提到的CVE
	// 2. 漏洞跟踪：从各种文档中提取CVE进行追踪管理
	// 3. 威胁情报分析：检测威胁情报报告中的CVE编号

	// 与ExtractCve的区别
	fmt.Println("\n与ExtractCve的区别:")
	fmt.Println("1. IsContainsCve - 仅检测是否存在，返回布尔值")
	fmt.Println("2. ExtractCve - 提取所有CVE并返回标准格式的CVE切片")
	// 预期输出:
	// 与ExtractCve的区别:
	// 1. IsContainsCve - 仅检测是否存在，返回布尔值
	// 2. ExtractCve - 提取所有CVE并返回标准格式的CVE切片
}
