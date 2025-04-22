package cve

import (
	"fmt"
	"time"
)

// GenerateCve 根据年份和序列号生成标准格式的CVE编号
//
// 通过给定的年份和序列号创建标准的CVE编号
//
// 示例:
//
//	输入: 2022, 12345
//	输出: "CVE-2022-12345"
//
// 使用场景:
//
//	需要动态生成CVE编号时使用
//
// 代码示例:
//
//	cveId := cve.GenerateCve(2022, 12345)
//	// cveId 为 "CVE-2022-12345"
func GenerateCve(year int, seq int) string {
	return Format(fmt.Sprintf("CVE-%d-%d", year, seq))
}

// GenerateFakeCve 生成一个假的CVE编号
//
// 无需提供参数，自动使用当前年份和随机序列号生成假的CVE编号
//
// 示例:
//
//	生成结果类似: "CVE-2023-54321"（假设当前年份为2023）
//
// 使用场景:
//
//	用于测试、示例或者占位符
//
// 代码示例:
//
//	fakeCve := cve.GenerateFakeCve()
//	// fakeCve 可能为 "CVE-2023-12345"
func GenerateFakeCve() string {
	currentYear := time.Now().Year()
	randomSeq := 10000 + time.Now().Nanosecond()%90000
	return GenerateCve(currentYear, randomSeq)
}
