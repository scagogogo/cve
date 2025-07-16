package main

import (
	"fmt"

	"github.com/scagogogo/cve"
)

func main() {
	text := `系统安全报告：
首先发现的漏洞是CVE-2021-44228，这是最严重的。
随后还发现了CVE-2022-22965和CVE-2022-33891。
最新发现的漏洞是CVE-2023-12345，正在评估中。`

	fmt.Println("原始文本:")
	fmt.Println(text)

	// 提取第一个CVE
	firstCve := cve.ExtractFirstCve(text)
	fmt.Printf("\n第一个CVE: %s\n", firstCve)

	// 提取最后一个CVE
	lastCve := cve.ExtractLastCve(text)
	fmt.Printf("最后一个CVE: %s\n", lastCve)

	// 提取所有CVE作为对比
	allCves := cve.ExtractCve(text)
	fmt.Println("\n所有CVE:")
	for i, c := range allCves {
		fmt.Printf("[%d] %s\n", i+1, c)
	}

	// 处理没有CVE的文本
	emptyText := "这个文本中没有任何CVE编号信息。"
	fmt.Printf("\n没有CVE的文本中的第一个CVE: %q\n", cve.ExtractFirstCve(emptyText))
	fmt.Printf("没有CVE的文本中的最后一个CVE: %q\n", cve.ExtractLastCve(emptyText))
}
