package main

import (
	"fmt"

	"github.com/scagogogo/cve"
)

func main() {
	// 示例1：格式化包含小写字母的CVE编号
	// Format函数会将CVE编号转换为标准大写格式
	input1 := "cve-2022-12345"
	fmt.Printf("原始输入: %s\n", input1)
	fmt.Printf("格式化后: %s\n\n", cve.Format(input1))
	// 预期输出:
	// 原始输入: cve-2022-12345
	// 格式化后: CVE-2022-12345

	// 示例2：格式化包含空白字符的CVE编号
	// Format函数会移除CVE编号两侧的空白字符
	input2 := " CVE-2021-44228 "
	fmt.Printf("原始输入: %q\n", input2)
	fmt.Printf("格式化后: %q\n\n", cve.Format(input2))
	// 预期输出:
	// 原始输入: " CVE-2021-44228 "
	// 格式化后: "CVE-2021-44228"

	// 示例3：格式化混合大小写的CVE编号
	// Format函数会统一将CVE部分转为大写
	input3 := "Cve-2023-9999"
	fmt.Printf("原始输入: %s\n", input3)
	fmt.Printf("格式化后: %s\n", cve.Format(input3))
	// 预期输出:
	// 原始输入: Cve-2023-9999
	// 格式化后: CVE-2023-9999

	// 总结: Format函数可以用于在比较或存储CVE编号前进行标准化，
	// 确保所有CVE编号遵循相同的格式规范，便于后续处理
}
