package main

import (
	"fmt"

	"github.com/scagogogo/cve"
)

func main() {
	fmt.Println("去除重复CVE示例")

	// 创建一个包含重复CVE的列表
	cveList := []string{
		"CVE-2022-1111",
		"cve-2022-1111", // 与第一个相同，但大小写不同
		"CVE-2022-2222",
		"CVE-2021-3333",
		"CVE-2022-2222",   // 与第三个完全相同
		" CVE-2021-3333 ", // 与第四个相同，但有空格
		"CVE-2020-4444",
	}

	fmt.Println("原始CVE列表:")
	for i, id := range cveList {
		fmt.Printf("[%d] %q\n", i+1, id)
	}

	// 去除重复的CVE
	uniqueCves := cve.RemoveDuplicateCves(cveList)

	fmt.Println("\n去重后的CVE列表:")
	for i, id := range uniqueCves {
		fmt.Printf("[%d] %s\n", i+1, id)
	}

	// 演示时显示每个去重后CVE的来源
	fmt.Println("\n去重效果分析:")

	// 创建映射表示原始索引
	originalIndices := make(map[string][]int)
	for i, id := range cveList {
		formattedID := cve.Format(id)
		originalIndices[formattedID] = append(originalIndices[formattedID], i+1)
	}

	// 显示每个去重后的CVE来自哪些原始项
	for i, id := range uniqueCves {
		indices := originalIndices[id]
		indicesStr := ""
		for j, idx := range indices {
			if j > 0 {
				indicesStr += ", "
			}
			indicesStr += fmt.Sprintf("%d", idx)
		}
		fmt.Printf("[%d] %s - 来自原始列表中的第 %s 项\n", i+1, id, indicesStr)
	}

	// 应用场景示例
	fmt.Println("\n应用场景示例 - 合并多个来源的CVE:")

	// 模拟来自不同来源的CVE列表
	source1 := []string{"CVE-2022-1111", "CVE-2022-2222"}
	source2 := []string{"cve-2022-1111", "CVE-2022-3333"}
	source3 := []string{"CVE-2022-4444", "CVE-2022-2222"}

	fmt.Println("来源1的CVE:", source1)
	fmt.Println("来源2的CVE:", source2)
	fmt.Println("来源3的CVE:", source3)

	// 合并所有来源
	merged := make([]string, 0)
	merged = append(merged, source1...)
	merged = append(merged, source2...)
	merged = append(merged, source3...)

	fmt.Println("\n合并后的CVE列表:")
	for i, id := range merged {
		fmt.Printf("[%d] %s\n", i+1, id)
	}

	// 去重
	uniqueMerged := cve.RemoveDuplicateCves(merged)

	fmt.Println("\n合并并去重后的CVE列表:")
	for i, id := range uniqueMerged {
		fmt.Printf("[%d] %s\n", i+1, id)
	}
	fmt.Printf("\n总计: 从%d个条目中提取出%d个唯一的CVE\n",
		len(merged), len(uniqueMerged))
}
