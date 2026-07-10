# cve count-by-year 按年计数

:::tip 📂 查看源码
[`cmd/stats.go:12`](https://github.com/scagogogo/cve-skills/blob/main/cmd/stats.go#L12-L31) — 在 GitHub 上查看 cobra 命令定义（第 12–31 行）。
:::

将一组 CVE 编号按年份分组，并输出每个年份下的 CVE 数量 —— 一份快速的时间维度漏洞分布直方图。

:::tip 🖥️ 适用场景
- 对一份 CVE 语料按年统计，快速发现发布高峰或低谷。
- 从扁平的编号列表生成逐年分布报告。
- 在将另一个 `cve` 命令（如 `filter`、`extract`）的输出继续传入下游前先做汇总。
:::

## 命令语法

```bash
cve count-by-year <cve-list>...
```

该命令接收一个或多个位置参数。每个参数本身也可包含逗号分隔的 CVE，因此 `CVE-2021-1,CVE-2022-1` 与 `CVE-2021-1 CVE-2022-1` 等价。当未提供参数且 stdin 为管道输入时，它会逐行读取 CVE 编号（跳过空行）。

## 参数与选项

- `<cve-list>`（位置参数，至少一个）：一个或多个 CVE 编号，例如 `CVE-2021-44228`。多个 CVE 可作为独立参数传入，也可作为单个逗号分隔的参数传入。
- stdin 回退：若未提供参数且 stdin 不是终端，则每一非空行视为一个输入（随后同样按逗号拆分）。
- 本命令自身**不定义任何 flag**。全局 `-q, --quiet` 标志继承自根命令。
- 若既无参数、也无管道 stdin 输入，命令以退出码 `1` 退出，并输出错误 `requires at least 1 argument (CVE list)`。

## 使用示例

统计三年份下的三个 CVE：

```bash
$ cve count-by-year CVE-2021-44228 CVE-2022-12345 CVE-2021-7
2021: 2
2022: 1
```

单个逗号分隔的参数会被同样拆分：

```bash
$ cve count-by-year CVE-2024-1,CVE-2024-2,CVE-2023-9
2023: 1
2024: 2
```

通过管道从其他命令传入列表 —— 每一行成为一个输入：

```bash
$ printf 'CVE-2020-1\nCVE-2020-2\nCVE-2019-5\n' | cve count-by-year
2019: 1
2020: 2
```

混合使用参数与逗号，统计混合语料：

```bash
$ cve count-by-year CVE-2018-1000 CVE-2019-1,CVE-2019-2 CVE-2018-5
2018: 2
2019: 2
```

与 `cve filter valid` 组合，先剔除非法项再计数：

```bash
$ cve filter valid CVE-2021-1 not-a-cve CVE-2022-3 | cve count-by-year
2021: 1
2022: 1
```

## 工作流程

```mermaid
flowchart LR
    A["参数或 stdin"] --> B["按逗号拆分每个输入"]
    B --> C["cve.CountByYear(list)"]
    C --> D["map[年份]数量"]
    D --> E["逐项打印 '年份: 数量'"]
    E --> F["退出码 0"]
```

## 对应 Go API

本命令是 [`CountByYear`](/api/functions/count-by-year) 的轻量封装，后者遍历列表、提取每个 CVE 的年份，返回 `map[int]int`（年份到数量）。CLI 遍历该 map 并逐项打印 `年份: 数量`。当你在代码中需要结构化的 map 而非纯文本输出时，请直接使用该 Go 函数 —— 例如用于喂给绘图库，或计算排序后的明细。

## 退出码与输出

- 退出码 `0`：命令正常结束，每个不同年份输出一行 `年份: 数量`。
- 退出码 `1`：未提供任何输入（既无参数也无管道 stdin）。此时向 stderr 输出错误 `requires at least 1 argument (CVE list)`。
- stdout：每个年份一行，格式为 `<年份>: <数量>`。
- stderr：仅输出上述缺失参数错误。

## 注意事项

- 输出顺序**未经排序** —— 遵循 Go map 的迭代顺序，是随机的。若需要年份升序，请管道传入 `sort -n`。
- 年份提取纯属文本解析，非法 CVE 会被 `ExtractCveYearAsInt` 映射为年份 `0`。此类条目会以 `0: <数量>` 出现在输出中。若对此敏感，请先用 `cve validate` 或 `cve filter valid` 预过滤。
- 不会校验年份是否落在 `1999..当前年份` 范围内。假设的 `CVE-1800-1` 会被计入年份 `1800` 而不报错。
- 计数前不会规范化大小写与两侧空白 —— 请传入已格式化的编号，或先执行 `cve format`。
- 逗号拆分适用于**每一个**输入，包括从 stdin 读入的每一行。形如 `CVE-2021-1, CVE-2022-2` 的行会产出两个条目（注意第二个的前导空格不会被去除）。

## 内部实现

该命令是一个 cobra 命令，其 `RunE`（cmd/stats.go L16-L30）驱动整个流程：

- **参数接入**：`inputs := readInputs(args)` 先收集位置参数；当没有参数且 stdin 为管道输入时，回退为逐行读取 stdin 中的非空行。该函数负责 args 与 stdin 的合并，因此 `RunE` 本身并不直接读取 stdin。
- **无 flag**：本命令不定义任何本地 flag，仅继承根命令的 `-q, --quiet`。`RunE` 不调用 `cmd.Flags()`，完全基于 `args` 工作。
- **逗号拆分**：每个收集到的输入都经 `strings.Split(input, ",")` 拆分，拆出的片段追加进同一个 `cveList []string`。这就是 `CVE-2021-1,CVE-2022-1` 与 `CVE-2021-1 CVE-2022-1` 等价的原因。
- **库函数调用与输出**：`counts := cve.CountByYear(cveList)` 返回 `map[int]int`；`RunE` 遍历该 map，用 `fmt.Printf("%d: %d\n", year, count)` 逐项打印，然后返回 `nil`。空输入守卫（`len(inputs) == 0`）在调用库函数之前就返回 `fmt.Errorf("requires at least 1 argument (CVE list)")`。

## 参数流

```text
+----------------------+    +---------------------------+
| 命令行参数（位置）   |    | stdin（仅当无参数且为管道 |
|  CVE-2021-1 CVE-2022 |    |   时逐行读取）            |
+----------+-----------+    +-------------+-------------+
           |                              |
           +--------------+---------------+
                          v
              +---------------------------+
              | readInputs(args)          |
              |  - 合并参数与 stdin        |
              |  - 跳过 stdin 空行         |
              +-------------+-------------+
                            |
                            v
              +---------------------------+
              | 对每个 input：            |
              |  strings.Split(input, ",")|
              |  追加到 cveList            |
              +-------------+-------------+
                            |
                            v
              +---------------------------+
              | len(inputs) == 0 ?        |
              |   是 -> 返回错误           |
              +-------------+-------------+
                            | 否
                            v
              +---------------------------+
              | cve.CountByYear(cveList)  |
              |  -> map[int]int           |
              +-------------+-------------+
                            |
                            v
              +---------------------------+
              | for year, count := range  |
              |   fmt.Printf("%d: %d\n")  |
              +-------------+-------------+
                            |
                            v
                  +-----------------+
                  | return nil      |
                  | 退出码 0        |
                  +-----------------+
```

## 边界情形

| 输入 | 行为 | 退出码 / 输出 |
|---|---|---|
| 无参数且 stdin 为终端（无管道输入） | `readInputs` 返回空切片，触发守卫 | 退出 `1`；stderr：`requires at least 1 argument (CVE list)` |
| 无参数且管道 stdin 为空（如 `printf '' \| cve ...`） | stdin 行全为空，`inputs` 为空 | 退出 `1`；同上 stderr 错误 |
| 有参数但全部非法（如 `not-a-cve`） | 年份提取将每项映射为年份 `0`；`CountByYear` 返回 `{0: N}` | 退出 `0`；stdout：`0: N` |
| 单个逗号分隔参数（如 `CVE-2021-1,CVE-2022-2`） | 一个输入被拆成两个 CVE | 退出 `0`；按年统计 |
| 参数与逗号混合（如 `CVE-2018-1 CVE-2019-1,CVE-2019-2`） | 每个参数独立拆分后拼接 | 退出 `0`；按年统计 |
| stdin 行带尾随逗号（如 `CVE-2021-1,`） | `strings.Split` 产出 `["CVE-2021-1", ""]`；空片段计入年份 `0` | 退出 `0`；可能出现 `0:` 行 |
| stdin 行含空格（如 `CVE-2021-1, CVE-2022-2`） | 仅按 `,` 拆分，不去除空白，第二个片段保留前导空格 | 退出 `0`；非法片段映射为年份 `0` |
| `CountByYear` 返回空结果（本命令中不会发生） | 守卫已阻止空 `cveList`，map 至少含一项 | 不适用 |

## 退出码

- **成功 —— 退出 `0`**：`RunE` 在打印后返回 `nil`，cobra 以 `0` 退出。stdout 为每个不同年份输出一行 `<年份>: <数量>`（顺序由 Go map 迭代随机决定）。
- **失败 —— 退出 `1`**：当 `len(inputs) == 0` 时，`RunE` 返回 `fmt.Errorf("requires at least 1 argument (CVE list)")`。cobra 将该错误打印到 stderr 并以 `1` 退出。这是唯一显式返回的错误。
- **无其他显式错误路径**：`RunE` 不调用 `os.Exit`，也不包装库函数错误；`cve.CountByYear` 本身不返回错误。因此任何非 `0` 退出都仅来自缺失输入守卫（或意外 panic，cobra 会以非零码退出）。

## 相关命令

- [cve year-range](/cli/commands/year-range) —— 获取列表中最早与最晚的年份。
- [cve seq-range](/cli/commands/seq-range) —— 获取指定年份的序列号范围。
- [cve filter by-year](/cli/commands/filter-by-year) —— 仅保留指定年份的 CVE。
- [cve filter group-by-year](/cli/commands/filter-group-by-year) —— 将 CVE 按年份分桶。
- [CLI 参考](/cli) —— 完整命令树与 I/O 约定。
