package main

import (
	"fmt"
	"time"

	"github.com/scagogogo/cve"
)

func main() {
	fmt.Println("生成CVE编号示例")

	// 获取当前年份
	currentYear := time.Now().Year()

	// 使用指定年份和序列号生成CVE
	year := 2022
	seq := 12345
	generatedCve := cve.GenerateCve(year, seq)

	fmt.Printf("使用年份 %d 和序列号 %d 生成的CVE: %s\n\n", year, seq, generatedCve)

	// 使用当前年份生成CVE
	currentYearCve := cve.GenerateCve(currentYear, 99999)
	fmt.Printf("使用当前年份 %d 生成的CVE: %s\n\n", currentYear, currentYearCve)

	// 演示序列号格式化
	// 注意：序列号保持原样，不会自动添加前导零
	smallSeq := 123
	smallSeqCve := cve.GenerateCve(year, smallSeq)
	fmt.Printf("使用小序列号 %d 生成的CVE: %s\n\n", smallSeq, smallSeqCve)

	// 应用场景示例 - 批量生成CVE
	fmt.Println("应用场景示例 - 批量生成一组特定年份的CVE:")

	batchYear := 2023
	startSeq := 10001
	count := 5

	fmt.Printf("生成%d年的%d个连续CVE，起始序列号为%d:\n", batchYear, count, startSeq)
	for i := 0; i < count; i++ {
		seq := startSeq + i
		id := cve.GenerateCve(batchYear, seq)
		fmt.Printf("[%d] %s\n", i+1, id)
	}

	// 应用场景示例 - 格式化输入
	fmt.Println("\n应用场景示例 - 从不同来源整理CVE格式:")

	// 假设这些数据来自不同的数据源，格式不统一
	sourceData := []struct {
		Year int
		Seq  int
	}{
		{2022, 44228}, // Log4Shell
		{2021, 45046}, // Log4j漏洞
		{2022, 22965}, // Spring4Shell
	}

	fmt.Println("标准化格式后的CVE:")
	for i, data := range sourceData {
		standardCve := cve.GenerateCve(data.Year, data.Seq)
		fmt.Printf("[%d] 源数据(%d, %d) -> %s\n", i+1, data.Year, data.Seq, standardCve)
	}
}
