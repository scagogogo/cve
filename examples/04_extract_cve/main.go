package main

import (
	"fmt"

	"github.com/scagogogo/cve"
)

func main() {
	// 示例1：从文本中提取所有CVE编号
	text := `安全公告：系统受到多个漏洞影响，包括：
- cve-2021-44228（Log4Shell）
- CVE-2022-22965（Spring4Shell）
- CVE-2022-1234
建议尽快更新到最新版本。`

	fmt.Println("原始文本:")
	fmt.Println(text)

	fmt.Println("\n提取的CVE编号:")
	cveList := cve.ExtractCve(text)
	for i, c := range cveList {
		fmt.Printf("[%d] %s\n", i+1, c)
	}

	// 示例2：从不包含CVE的文本中提取
	text2 := "这个文本中不包含任何CVE编号。"
	fmt.Println("\n另一个示例文本:")
	fmt.Println(text2)

	cveList2 := cve.ExtractCve(text2)
	fmt.Println("\n提取的CVE编号:")
	if len(cveList2) == 0 {
		fmt.Println("未找到任何CVE编号")
	} else {
		for i, c := range cveList2 {
			fmt.Printf("[%d] %s\n", i+1, c)
		}
	}
}
