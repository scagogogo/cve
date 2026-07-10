# Example: GenerateCve

:::tip 📂 View Source
[`examples/17_generate_cve/main.go`](https://github.com/scagogogo/cve-skills/blob/main/examples/17_generate_cve/main.go) — open the full runnable example on GitHub.
:::

Compose a standard `CVE-YYYY-NNNNN` identifier from a year and a sequence number with `cve.GenerateCve`. This is the simplest way to normalize loose `(year, seq)` pairs coming from databases, tickets, or spreadsheets into canonical CVE strings.

:::tip 🎯 Learning objectives
- Understand the signature and behavior of `cve.GenerateCve`
- Learn that the sequence number is kept as-is and is NOT zero-padded
- Use `GenerateCve` for batch generation and format normalization across data sources
:::

## Scenario

A vulnerability response team ingests advisories from several feeds. Each feed stores the CVE year and sequence number as separate integer fields, and the formats are inconsistent. Before pushing the records into a unified tracking system, the team needs every record rendered as a canonical `CVE-YYYY-NNNNN` string. `GenerateCve` takes a year and a sequence number and returns the standardized identifier, making it the ideal glue between raw data sources and downstream CVE tooling.

## Full code

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

## How to run

```bash
cd examples/17_generate_cve && go run main.go
```

## Expected output

The line that uses `currentYear` reflects the system year at runtime (replace `<当前年份>` with the actual year, e.g. `2026`).

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

## Code walkthrough

The example starts by printing a header and reading the current year from `time.Now().Year()` so the program can demonstrate dynamic year usage.

- 📋 **Basic generation** — `cve.GenerateCve(2022, 12345)` produces the canonical `CVE-2022-12345`. The function internally builds the string `CVE-2022-12345` and runs it through `Format`, so the result is always uppercase.
- ▶️ **Current-year generation** — `cve.GenerateCve(currentYear, 99999)` shows how to pair a runtime year with a sequence number, useful when minting placeholder or test identifiers for the current year.
- 💡 **No zero-padding** — `cve.GenerateCve(2022, 123)` returns `CVE-2022-123`, not `CVE-2022-00123`. The source comment calls this out explicitly: the sequence number is kept as-is, with no automatic leading zeros. If you need padded sequences, reach for `FormatSeq` instead.
- 🔗 **Batch generation** — a loop from `startSeq` to `startSeq + count` generates five consecutive CVEs for 2023, illustrating how to expand a `(year, start, count)` tuple into a list of identifiers.
- 🔗 **Format normalization** — a slice of `(Year, Seq)` structs taken from heterogeneous sources is mapped through `GenerateCve` into uniform `CVE-YYYY-NNNNN` strings, including well-known entries such as Log4Shell (`CVE-2022-44228`) and Spring4Shell (`CVE-2022-22965`).

```mermaid
flowchart TD
    A["(year, seq) pair"] --> B["GenerateCve(year, seq)"]
    B --> C["Format(\"CVE-YYYY-NNNNN\")"]
    C --> D["Canonical CVE string"]
    D --> E1["CVE-2022-12345"]
    D --> E2["CVE-2022-123 (no padding)"]
    D --> E3["Batch: CVE-2023-10001..10005"]
    D --> E4["Normalized: CVE-2022-44228 etc."]
```

## Related functions

- [GenerateCve](/api/functions/generate-cve) — the function used in this example
- [GenerateFakeCve](/api/functions/generate-fake-cve) — generate a random fake CVE with no arguments
- [Format](/api/functions/format) — the normalization helper invoked internally by `GenerateCve`
- [FormatSeq](/api/functions/format-seq) — zero-pad a sequence number when padding is required
- [ParseCveRange](/api/functions/parse-cve-range) — expand a CVE range expression back into a list of identifiers

## Extensions

- 🎯 Call `GenerateCve(2022, 1)` and `GenerateCve(2022, 99999)` to confirm the sequence number width is never padded or truncated.
- 🎯 Replace the batch loop with a single `ParseCveRange("CVE-2023-10001 to CVE-2023-10005")` call and compare the output.
- 🎯 Feed the normalized `sourceData` output into `ValidateCves` to verify every generated string passes format validation.
