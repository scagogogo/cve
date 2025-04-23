package cve

import (
	"fmt"
	"time"
)

// GenerateCve 根据年份和序列号生成标准格式的CVE编号
//
// 通过给定的年份和序列号创建标准的CVE编号
//
// 参数:
//   - year: CVE年份，整数格式，如2022
//   - seq: CVE序列号，整数格式，如12345
//
// 返回值:
//   - string: 生成的标准格式CVE编号，如"CVE-2022-12345"
//
// 格式规则:
//   - 生成的CVE编号格式为"CVE-YYYY-NNNNN"
//   - 返回结果始终为大写形式
//
// 示例:
//
//	输入: 2022, 12345
//	输出: "CVE-2022-12345"
//
//	输入: 2021, 44228
//	输出: "CVE-2021-44228"
//
// 注意事项:
//   - 此函数不会验证年份是否合理（如是否在1999年之后）
//   - 序列号可以是任意整数，不限制位数
//
// 使用场景:
//   - 需要动态生成CVE编号时使用
//   - 创建新的CVE标识符
//
// 代码示例:
//
//	cveId := cve.GenerateCve(2022, 12345)
//	// cveId 为 "CVE-2022-12345"
//
//	// 可以组合使用提取和生成功能
//	year := 2023
//	seq := cve.ExtractCveSeqAsInt("CVE-2022-67890")
//	newCve := cve.GenerateCve(year, seq) // 生成"CVE-2023-67890"
func GenerateCve(year int, seq int) string {
	return Format(fmt.Sprintf("CVE-%d-%d", year, seq))
}

// GenerateFakeCve 生成一个假的CVE编号
//
// 无需提供参数，自动使用当前年份和随机序列号生成假的CVE编号
//
// 参数:
//
//	无
//
// 返回值:
//   - string: 生成的标准格式CVE编号，如"CVE-2023-54321"
//
// 生成规则:
//   - 使用当前系统年份作为CVE年份
//   - 序列号随机生成，范围在10000到99999之间
//   - 返回结果始终为大写形式
//
// 示例:
//
//	输出结果类似: "CVE-2023-54321"（假设当前年份为2023）
//
// 随机性:
//   - 序列号基于当前时间的纳秒部分生成，具有一定随机性
//   - 不保证全局唯一性，仅用于测试或示例目的
//
// 使用场景:
//   - 用于测试、示例或者占位符
//   - 快速创建模拟CVE数据
//
// 代码示例:
//
//	fakeCve := cve.GenerateFakeCve()
//	// fakeCve 可能为 "CVE-2023-12345"
//
//	// 在测试中创建多个随机CVE
//	var testCves []string
//	for i := 0; i < 5; i++ {
//	    testCves = append(testCves, cve.GenerateFakeCve())
//	}
func GenerateFakeCve() string {
	currentYear := time.Now().Year()
	randomSeq := 10000 + time.Now().Nanosecond()%90000
	return GenerateCve(currentYear, randomSeq)
}
