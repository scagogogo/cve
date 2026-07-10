# 示例：生成 CVE

:::tip 📂 查看源码
[`examples/17_generate_cve/main.go`](https://github.com/scagogogo/cve-skills/blob/main/examples/17_generate_cve/main.go) — 在 GitHub 上查看完整可运行示例。
:::

使用 `cve.GenerateCve` 根据年份和序列号组装标准的 `CVE-YYYY-NNNNN` 标识符。这是将数据库、工单或表格中零散的 `(year, seq)` 数据规范化为标准 CVE 字符串的最简方式。

:::tip 🎯 学习目标
- 了解 `cve.GenerateCve` 的签名与行为
- 明确序列号保持原样、**不会**自动补前导零
- 学会用 `GenerateCve` 进行批量生成与跨数据源格式归一化
:::

## 场景

某漏洞响应团队需要从多个情报源接入漏洞通告。每个源都把 CVE 的年份和序列号存为独立的整数字段，且格式不统一。在将记录写入统一追踪系统之前，团队需要把每条记录渲染为规范的 `CVE-YYYY-NNNNN` 字符串。`GenerateCve` 接收年份和序列号并返回标准标识符，是连接原始数据源与下游 CVE 工具链的理想桥梁。

## 完整代码

```go
package main

import (
	"fmt"
	"time"

	"github.com/scagogogo/cve-skills"
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
```

## 运行方式

```bash
cd examples/17_generate_cve && go run main.go
```

## 预期输出

使用 `currentYear` 的那一行反映运行时的系统年份（请将 `<当前年份>` 替换为实际年份，如 `2026`）。

```text
生成CVE编号示例
使用年份 2022 和序列号 12345 生成的CVE: CVE-2022-12345

使用当前年份 <当前年份> 生成的CVE: CVE-<当前年份>-99999

使用小序列号 123 生成的CVE: CVE-2022-123

应用场景示例 - 批量生成一组特定年份的CVE:
生成2023年的5个连续CVE，起始序列号为10001:
[1] CVE-2023-10001
[2] CVE-2023-10002
[3] CVE-2023-10003
[4] CVE-2023-10004
[5] CVE-2023-10005

应用场景示例 - 从不同来源整理CVE格式:
标准化格式后的CVE:
[1] 源数据(2022, 44228) -> CVE-2022-44228
[2] 源数据(2021, 45046) -> CVE-2021-45046
[3] 源数据(2022, 22965) -> CVE-2022-22965
```

## 代码讲解

示例先打印标题，并通过 `time.Now().Year()` 读取当前年份，用于演示动态年份的用法。

- 📋 **基础生成** — `cve.GenerateCve(2022, 12345)` 生成规范的 `CVE-2022-12345`。函数内部先拼出字符串 `CVE-2022-12345`，再交给 `Format` 处理，因此结果始终为大写。
- ▶️ **当前年份生成** — `cve.GenerateCve(currentYear, 99999)` 展示了如何把运行时年份与序列号配对，适合为当前年份生成占位或测试标识符。
- 💡 **不补前导零** — `cve.GenerateCve(2022, 123)` 返回的是 `CVE-2022-123`，而不是 `CVE-2022-00123`。源码注释明确指出：序列号保持原样，不会自动补前导零。若需要补零，请改用 `FormatSeq`。
- 🔗 **批量生成** — 从 `startSeq` 循环到 `startSeq + count`，为 2023 年生成 5 个连续 CVE，展示了如何把 `(year, start, count)` 元组展开为标识符列表。
- 🔗 **格式归一化** — 一组来自异构数据源的 `(Year, Seq)` 结构体通过 `GenerateCve` 映射为统一的 `CVE-YYYY-NNNNN` 字符串，包括 Log4Shell（`CVE-2022-44228`）、Spring4Shell（`CVE-2022-22965`）等知名条目。

```mermaid
flowchart TD
    A["(year, seq) 数据对"] --> B["GenerateCve(year, seq)"]
    B --> C["Format(\"CVE-YYYY-NNNNN\")"]
    C --> D["标准 CVE 字符串"]
    D --> E1["CVE-2022-12345"]
    D --> E2["CVE-2022-123（不补零）"]
    D --> E3["批量：CVE-2023-10001..10005"]
    D --> E4["归一化：CVE-2022-44228 等"]
```

## 涉及函数

- [GenerateCve](/zh/api/functions/generate-cve) — 本示例使用的函数
- [GenerateFakeCve](/zh/api/functions/generate-fake-cve) — 无参数生成随机假 CVE
- [Format](/zh/api/functions/format) — `GenerateCve` 内部调用的归一化辅助函数
- [FormatSeq](/zh/api/functions/format-seq) — 在需要补零时对序列号进行补零
- [ParseCveRange](/zh/api/functions/parse-cve-range) — 将 CVE 范围表达式展开回标识符列表

## 扩展练习

- 🎯 调用 `GenerateCve(2022, 1)` 和 `GenerateCve(2022, 99999)`，验证序列号既不会被补零，也不会被截断。
- 🎯 用一次 `ParseCveRange("CVE-2023-10001 to CVE-2023-10005")` 替换批量循环，比较输出结果。
- 🎯 将归一化后的 `sourceData` 输出送入 `ValidateCves`，确认每个生成的字符串都能通过格式校验。
