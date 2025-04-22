package main

import (
	"fmt"
	"time"

	"github.com/scagogogo/cve"
)

func main() {
	fmt.Println("生成随机CVE编号示例")

	// 获取当前年份
	currentYear := time.Now().Year()
	fmt.Printf("当前年份: %d\n\n", currentYear)

	// 生成一个随机CVE
	fakeCve := cve.GenerateFakeCve()
	fmt.Printf("生成的随机CVE: %s\n", fakeCve)

	// 验证生成的CVE
	fmt.Printf("验证生成的CVE:\n")
	fmt.Printf("- 是否符合CVE格式: %v\n", cve.IsCve(fakeCve))
	fmt.Printf("- 是否有效的CVE: %v\n", cve.ValidateCve(fakeCve))

	// 提取并检查年份和序列号
	year := cve.ExtractCveYear(fakeCve)
	seq := cve.ExtractCveSeq(fakeCve)

	fmt.Printf("- 年份: %s (应该是当前年份 %d)\n", year, currentYear)
	fmt.Printf("- 序列号: %s (应该是一个5位以上的随机数)\n\n", seq)

	// 生成多个随机CVE以展示随机性
	fmt.Println("生成多个随机CVE:")

	count := 5
	for i := 0; i < count; i++ {
		id := cve.GenerateFakeCve()
		fmt.Printf("[%d] %s\n", i+1, id)
	}

	// 应用场景示例
	fmt.Println("\n应用场景示例 - 使用随机CVE进行测试:")
	fmt.Println("1. 创建测试数据集:")

	testDataset := make([]string, 10)
	for i := range testDataset {
		testDataset[i] = cve.GenerateFakeCve()
	}

	for i, id := range testDataset {
		fmt.Printf("  [%d] %s\n", i+1, id)
	}

	fmt.Println("\n2. 对测试数据集执行排序操作:")
	sortedData := cve.SortCves(testDataset)

	for i, id := range sortedData {
		fmt.Printf("  [%d] %s\n", i+1, id)
	}

	fmt.Println("\n3. 按年份分组 (所有CVE应该在同一组):")
	groupedData := cve.GroupByYear(testDataset)

	for year, ids := range groupedData {
		fmt.Printf("  %s年的CVE (%d个): %v\n", year, len(ids), ids)
	}
}
