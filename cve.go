// Package cve 提供一系列CVE（Common Vulnerabilities and Exposures）相关的工具方法。
//
// 本包中提供的功能包括：
// - CVE格式验证和标准化：检查CVE编号格式合法性，转换为标准大写格式
// - 从文本中提取CVE标识符：单个或批量提取文本中包含的CVE编号
// - CVE的年份和序列号提取与比较：分解CVE编号，获取和比较组成部分
// - CVE的排序、过滤和分组：按年份或其他条件处理CVE列表
// - 生成标准格式的CVE标识符：创建符合规范的CVE编号
// - 去重和验证工具：处理重复的CVE编号，验证CVE有效性
//
// 使用示例:
//
//	// 从文本中提取CVE编号
//	report := "系统受到CVE-2021-44228和CVE-2022-12345的影响"
//	cveList := cve.ExtractCve(report)
//
//	// 验证CVE格式
//	isValid := cve.ValidateCve("CVE-2022-12345")
//
//	// 按年份筛选CVE
//	cves2022 := cve.FilterCvesByYear(cveList, 2022)
//
// 更多信息请参考README文档。
package cve

// Version 表示当前包的版本号
//
// 记录当前包的版本信息，遵循语义化版本规范
//
// 格式: vX.Y.Z
// - X: 主版本号，不兼容的API修改
// - Y: 次版本号，向后兼容的功能性新增
// - Z: 修订号，向后兼容的问题修正
const Version = "v0.0.1"
