# cve filter group-by-year 按年分组

:::tip 📂 查看源码
[`cmd/filter.go:99`](https://github.com/scagogogo/cve-skills/blob/main/cmd/filter.go#L99-L126) — 在 GitHub 上查看 cobra 命令定义（第 99–126 行）。
:::

将一组 CVE 编号按年份分组，每行输出一个年份，其后紧跟该年份的所有 CVE，每条缩进一行。

:::tip 🖥️ 适用场景
- 在报告或趋势可视化前，将原始 CVE 列表按发布年份分桶。
- 快速查看某个漏洞源或公告包中哪些年份占主导。
- 将混合年份的输入拆分为按年分桶，便于后续按年份的专项处理。
:::

## 命令语法

```bash
cve filter group-by-year [cve-id...]
```

该命令以位置参数接收 CVE 编号。当未提供参数且 stdin 有管道输入时，改为从 stdin 逐行读取编号。命令自身不定义任何 flag。

## 参数与选项

- `[cve-id...]`（位置参数，可选）：零个或多个 CVE 编号，例如 `CVE-2021-44228`。省略时回退到 stdin（每行一个，空行跳过）。
- 本命令自身**不定义任何 flag**。全局 `-q, --quiet` 标志继承自根命令。
- 若未提供参数且 stdin 为终端（非管道），`readInputs` 返回 `nil`，命令以退出码 `1` 退出且不输出任何内容。

## 使用示例

将三个跨两年的 CVE 分组；年份按升序输出，每条 CVE 缩进两个空格：

```bash
$ cve filter group-by-year CVE-2021-1111 CVE-2022-2222 CVE-2021-3333
2021:
  CVE-2021-1111
  CVE-2021-3333
2022:
  CVE-2022-2222
```

当输入来自其他工具时，通过 stdin 管道传入列表：

```bash
$ printf 'CVE-2020-5\nCVE-2023-7\nCVE-2020-9\n' | cve filter group-by-year
2020:
  CVE-2020-5
  CVE-2020-9
2023:
  CVE-2023-7
```

可任意混合年份 —— 输入中每个不同年份都会成为一个分桶：

```bash
$ cve filter group-by-year CVE-1999-1 CVE-2024-1 CVE-1999-2 CVE-2024-2
1999:
  CVE-1999-1
  CVE-1999-2
2024:
  CVE-2024-1
  CVE-2024-2
```

与 `cve filter dedup` 组合成管道，在分组前先去重：

```bash
$ cve filter dedup CVE-2022-1111 cve-2022-1111 CVE-2022-2222 | cve filter group-by-year
2022:
  CVE-2022-1111
  CVE-2022-2222
```

## 工作流程

```mermaid
flowchart LR
    A["输入: 参数或 stdin"] --> B["readInputs"]
    B --> C["GroupByYear(cves)"]
    C --> D["经 ExtractCveYear 提取每个 CVE 的年份"]
    D --> E["经 Format 规范化每个 CVE"]
    E --> F["map[年份] -> []cve"]
    F --> G["按年份升序排序"]
    G --> H["stdout: '年份:' + 缩进的 CVE"]
    H --> I["退出码 0"]
```

## 对应 Go API

本命令是 [`GroupByYear`](/api/functions/group-by-year) 的轻量封装，后者返回以提取出的年份为键的 `map[string][]string`。每个 CVE 在加入分桶前会经 `Format` 规范化。CLI 随后对键排序并输出；直接调用该 Go 函数时你拿到的是原始 map，需自行处理排序与渲染。

## 退出码与输出

- 退出码 `0`：命令正常结束，向 stdout 输出分组结果。
- 退出码 `1`：未提供任何输入（参数为空且 stdin 无管道输入）。不输出任何内容。
- stdout：每个年份（升序）输出一行 `年份:` 表头，其后每行一个缩进的 CVE。
- stderr：正常运行时不写入任何内容。

## 注意事项

- 年份在输出前通过 `sort.Strings` **按字符串升序**排序。由于 CVE 年份为四位零填充值，这与数值顺序一致，但比较本身是文本比较而非数值比较。
- 每个 CVE 在分组前会经 `Format` 重新格式化，因此 `cve-2022-1111` 与 `CVE-2022-1111` 会落入同一分桶，并以规范的大小写形式输出。
- 年份提取纯属文本解析；年份段非数字或缺失的非法编号会被归入 `ExtractCveYear` 所返回的值（通常为空字符串）对应的分桶。若对此敏感，请先用 `cve validate` 校验输入。
- 不会校验年份是否落在 `1999..当前年份` 范围内 —— 假设的 `CVE-1800-1` 会被归入 `1800` 分桶而不报错。
- 本命令不去重；重复的 CVE 会在其年份分桶内多次出现。若需要唯一条目，请先通过 `cve filter dedup` 处理。
- 同一年份分桶内，CVE 按**输入中出现顺序**输出，而非按序列号排序。如需分桶内排序，请随后使用 `cve compare sort`。

## 内部实现

`groupByYearCmd` cobra 命令（定义于 `cmd/filter.go:99-126`）以一条简明的流水线运行，自身不定义任何 flag：

- **输入收集**：`Run` 接收 `args []string`，直接传给 `readInputs(args)` —— 该共享辅助函数在有位置参数时返回参数，否则逐行读取 stdin（跳过空行）。本子命令内不做任何 flag 解析。
- **空输入守卫**：`if len(inputs) == 0 { os.Exit(1) }` 在没有任何输入时于调用库函数前立即中止，因此分组逻辑永不会在空切片上运行。
- **库函数调用**：`groups := cvepkg.GroupByYear(inputs)` 返回以提取出的年份为键的 `map[string][]string`。每个 CVE 在该函数内部经 `Format` 规范化、由 `ExtractCveYear` 分桶 —— CLI 自身不处理格式化。
- **输出格式化**：CLI 将 map 的键收集到 `[]string`，用 `sort.Strings` 排序，随后对每个年份输出 `fmt.Printf("%s:\n", y)` 作为表头，再对该分桶内每个 CVE 输出 `fmt.Printf("  %s\n", c)`（缩进两空格，分桶内保持输入顺序）。

## 参数流

```text
+----------------------+   +----------------------+   +---------------------------+
| CLI 参数 / stdin     |-->| readInputs(args)     |-->| []string inputs           |
| (每行一个 cve-id)    |   | (位置参数或管道)     |   | (为空? -> os.Exit(1))     |
+----------------------+   +----------------------+   +---------------------------+
                                                                |
                                                                v
                          +---------------------------------------+
                          | cvepkg.GroupByYear(inputs)            |
                          |   对每个 CVE 调用 ExtractCveYear      |
                          |   对每个 CVE 调用 Format (规范大小写) |
                          |   -> map[string][]string (年份->cve)  |
                          +---------------------------------------+
                                                                |
                                                                v
                          +---------------------------------------+
                          | 将 map 键收集到 []string years        |
                          | sort.Strings(years)  (升序)           |
                          +---------------------------------------+
                                                                |
                                                                v
                          +---------------------------------------+
                          | 对每个年份 y:                         |
                          |   fmt.Printf("%s:\n", y)              |
                          |   对 groups[y] 中每个 c:              |
                          |     fmt.Printf("  %s\n", c)           |
                          +---------------------------------------+
                                                                |
                                                                v
                          stdout (分组、缩进) -> 退出码 0
```

## 边界情形

| 输入 | 行为 | 退出码/输出 |
|---|---|---|
| 无位置参数且 stdin 为终端 | `readInputs` 返回 `nil`；空输入守卫立即触发 | 退出 `1`，stdout 与 stderr 均无输出 |
| 无位置参数且 stdin 有管道但为空 | `readInputs` 返回空切片；守卫触发 | 退出 `1`，无输出 |
| stdin 含空行 | `readInputs` 跳过空行；仅非空 token 进入 `GroupByYear` | 退出 `0`，输出非空输入的分组 |
| 非法 CVE（年份非数字或缺失） | `ExtractCveYear` 返回其文本结果（常为空字符串）；该 CVE 归入该键 | 退出 `0`，分桶表头可能为空行 |
| 超范围年份（如 `CVE-1800-1`） | 不做范围校验；归入字面年份 `1800` | 退出 `0`，输出 `1800:` 分桶 |
| 输入含重复 CVE | 不去重；重复项在其年份分桶内再次出现 | 退出 `0`，重复项被打印 |
| 大小写混合输入（`cve-2022-1`、`CVE-2022-1`） | 分桶前经 `Format` 规范化为标准形式 | 退出 `0`，单一分桶，标准大小写 |
| 分组后结果为空（不应发生） | 不可达：任何非空输入至少产生一个分桶 | 不适用 |

## 退出码

- **成功（退出 `0`）**：`readInputs` 返回非空切片且 `GroupByYear` 生成了 map；循环向 stdout 输出分组结果。成功路径上命令不调用 `os.Exit`，因此采用 Go 默认退出码 `0`。
- **失败（退出 `1`）**：唯一的显式失败是空输入守卫 `if len(inputs) == 0 { os.Exit(1) }`。它立即以 `1` 退出，**不**向 stderr 输出任何内容 —— 对比 `by-year` 等同级子命令会先输出 `error: --year is required`。
- **stderr**：本子命令从不写入 stderr；所有诊断信息只可能来自 cobra 自身的 flag/usage 处理，但 `group-by-year` 不定义 flag，故该路径在此不会被触发。

## 相关命令

- [cve filter by-year](/cli/commands/filter-by-year) —— 仅保留指定单一年份的 CVE。
- [cve filter by-year-range](/cli/commands/filter-by-year-range) —— 保留年份范围内的 CVE（含边界）。
- [cve filter recent](/cli/commands/filter-recent) —— 保留最近 N 年的 CVE。
- [cve filter dedup](/cli/commands/filter-dedup) —— 在分组前去除重复的 CVE 编号。
- [CLI 参考](/cli) —— 完整命令树与 I/O 约定。
