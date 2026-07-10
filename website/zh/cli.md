# CLI 参考

`cve` 命令行工具将整个库能力暴露为可组合、确定性的子命令。它专为 **AI agent 与
Shell 管道** 而设计:每个命令都可从参数 *或* stdin 读取输入,向 stdout 输出纯文本
行,并返回稳定的退出码——没有交互式提问、没有颜色、没有意外。

## 安装

预编译二进制与包管理器安装方式详见 [下载与安装](/zh/download)。最快捷的两种方式:

```bash
# 预编译二进制,无需 Go 工具链
curl -fsSL https://raw.githubusercontent.com/scagogogo/cve-skills/main/scripts/install.sh | bash

# 或使用 Go
go install github.com/scagogogo/cve-skills/cmd/cve@latest
```

验证:

```bash
cve version
```

## 输入输出约定

以下规则对每个子命令都成立,这正是该工具可安全用于脚本的原因:

| 方面 | 行为 |
|------|------|
| **输入** | 位置参数,**或** 在未提供参数时从 stdin 读取(每行一项)。 |
| **列表输入** | 接收 `<cve-list>` 的命令同时支持逗号分隔值 *与* 多个参数/多行。 |
| **输出** | stdout 每行一个结果。多字段行以 `\t` 分隔。 |
| **布尔值** | 直接打印为 `true` / `false`。 |
| **空结果** | 无 stdout 输出(退出码仍为 `0`)。 |
| **错误** | 信息输出到 stderr,退出码为 `1`。用法/错误不会污染 stdout。 |
| **`-q, --quiet`** | 全局标志,抑制非必要输出。 |

由于输入会回退到 stdin,命令天然可以串联:

```bash
cat advisory.txt | cve extract | cve compare sort | cve filter dedup
```

每一段都读取上一段的 stdout,每行一个 CVE:

```mermaid
flowchart LR
    F["advisory.txt"] -->|stdin| A["cve extract"]
    A -->|CVE 列表| B["cve compare sort"]
    B -->|已排序| C["cve filter dedup"]
    C -->|去重结果| O["stdout"]
```

## 命令树

```text
cve
├── format <cve...>                     格式化为标准大写
├── format-seq <width> <cve>            将序列号补零到固定宽度
├── validate <cve...>                   全面验证(格式 + 年份 + 序列号)
│   ├── is-cve <text...>                文本是否恰好是一个 CVE?
│   ├── contains-cve <text...>          文本是否包含 CVE?
│   └── year-ok <cve...> [--cutoff N]   年份是否在 1999..当前(+N)?
├── validate-batch <cve-list>          逐项验证并给出原因
├── filter-valid <cve-list>            仅保留有效 CVE
├── extract <text...>                   从文本提取所有 CVE
│   ├── first <text...>                 仅第一个 CVE
│   ├── last <text...>                  仅最后一个 CVE
│   ├── year <cve...>                   年份部分
│   ├── seq <cve...>                    序列号部分
│   └── split <cve...>                  年份<TAB>序列号
├── compare <a> <b>                     -1 / 0 / 1
│   ├── by-year <a> <b>                 仅按年份比较
│   └── sort <cve...>                   升序排序
├── filter                              (命令组;见子命令)
│   ├── by-year --year Y <cve...>       保留单一年份
│   ├── by-year-range --start --end     保留闭区间年份范围
│   ├── recent --years N <cve...>       保留最近 N 年
│   ├── group-by-year <cve...>          按年份分组打印
│   └── dedup <cve...>                  去重(大小写不敏感)
├── filter-pattern <pattern> <list>     通配符过滤,如 "CVE-2022-*"
├── intersect <list1> <list2>           集合交集
├── union <list1> <list2>               集合并集
├── diff <list1> <list2>                集合差集(list1 - list2)
├── parse-range <range-expr>            将范围表达式展开为逐个 CVE
├── is-consecutive <a> <b>              两个 CVE 是否相邻?
├── count-by-year <cve-list>            按年份计数
├── year-range <cve-list>              最早/最晚年份及跨度
├── seq-range <year> <cve-list>        指定年份的序列号最小/最大值
├── generate                            (命令组;见子命令)
│   ├── cve --year Y --seq S            构造一个 CVE
│   └── fake                            当前年份的随机 CVE(测试用)
└── version                             打印版本号
```

## 格式化与验证

### `format`

将 CVE 编号标准化为大写形式(去除空白、`CVE` 前缀转大写)。

```bash
$ cve format cve-2022-12345 " CVE-2021-44228 "
CVE-2022-12345
CVE-2021-44228
```

### `format-seq`

将序列号补零到固定宽度。参数:`<width> <cve>`。

```bash
$ cve format-seq 7 CVE-2022-123
CVE-2022-0000123
```

### `validate`

全面验证——检查格式、年份范围(1999 至当前年份)以及序列号为正。
输出为 `标准化 CVE<TAB>布尔值`。

```bash
$ cve validate CVE-2022-12345 CVE-1998-12345
CVE-2022-12345	true
CVE-1998-12345	false
```

#### `validate is-cve`

整个输入字符串是否 *恰好* 是一个 CVE 编号?输出:`输入<TAB>布尔值`。

```bash
$ cve validate is-cve CVE-2022-12345 "text CVE-2022-1"
CVE-2022-12345	true
text CVE-2022-1	false
```

#### `validate contains-cve`

输入中是否包含至少一个 CVE?输出:`true` / `false`。

```bash
$ cve validate contains-cve "affected by CVE-2021-44228"
true
```

#### `validate year-ok`

CVE 年份是否落在 `1999..当前年份`?通过 `--cutoff N`(`-c`)允许向未来放宽
至多 `N` 年。输出:`标准化 CVE<TAB>布尔值`。

```bash
$ cve validate year-ok CVE-2022-1 CVE-1998-1
CVE-2022-1	true
CVE-1998-1	false
```

### `validate-batch`

验证一个列表并逐项打印结论,失败时附带原因。

```bash
$ cve validate-batch "CVE-2022-12345,CVE-1998-1,not-a-cve"
✓ CVE-2022-12345
✗ CVE-1998-1 — year 1998 is before 1999
✗ not-a-cve — invalid CVE format
```

### `filter-valid`

将列表过滤为仅保留有效的 CVE。

```bash
$ cve filter-valid "CVE-2022-12345,bad,CVE-2021-44228"
CVE-2022-12345
CVE-2021-44228
```

## 提取

### `extract`

从自由文本中提取所有 CVE。

```bash
$ cve extract "System affected by CVE-2021-44228 and CVE-2022-12345"
CVE-2021-44228
CVE-2022-12345
```

子命令用于缩小结果:

```bash
$ cve extract first "CVE-2021-44228 and CVE-2022-12345"   # → CVE-2021-44228
$ cve extract last  "CVE-2021-44228 and CVE-2022-12345"   # → CVE-2022-12345
$ cve extract year  CVE-2022-12345                        # → 2022
$ cve extract seq   CVE-2022-12345                        # → 12345
```

### `extract split`

拆分为年份与序列号,以制表符分隔(`年份<TAB>序列号`)。

```bash
$ cve extract split CVE-2022-12345
2022	12345
```

## 比较与排序

### `compare`

按年份 *与* 序列号比较两个 CVE。输出:`-1`(a < b)、`0`(相等)、`1`(a > b)。

```bash
$ cve compare CVE-2021-44228 CVE-2022-12345
-1
```

### `compare by-year`

仅按年份比较(符号表示先后,数值为年份差)。

```bash
$ cve compare by-year CVE-2022-1 CVE-2021-9
1
```

### `compare sort`

按年份、再按序列号升序排序。

```bash
$ cve compare sort CVE-2022-2222 CVE-2020-1111 CVE-2022-1111
CVE-2020-1111
CVE-2022-1111
CVE-2022-2222
```

## 过滤与分组

### `filter by-year`

仅保留指定年份的 CVE(`--year`、`-y`,必填)。

```bash
$ cve filter by-year --year 2022 CVE-2021-1111 CVE-2022-2222 CVE-2022-3333
CVE-2022-2222
CVE-2022-3333
```

### `filter by-year-range`

保留闭区间年份范围内的 CVE(`--start`/`-s`,`--end`/`-e`)。

```bash
$ cve filter by-year-range --start 2021 --end 2022 CVE-2020-1 CVE-2021-2 CVE-2022-3 CVE-2023-4
CVE-2021-2
CVE-2022-3
```

### `filter recent`

保留最近 `N` 年的 CVE(`--years`、`-n`,必填)。窗口为
`[当前年份 - N + 1, 当前年份]`。

```bash
$ cve filter recent --years 2 CVE-2020-1 CVE-2025-2 CVE-2026-3
CVE-2025-2
CVE-2026-3
```

### `filter group-by-year`

按年份分组,以 `年份:` 为标题、成员缩进打印(年份已排序)。

```bash
$ cve filter group-by-year CVE-2021-1111 CVE-2022-2222 CVE-2021-3333
2021:
  CVE-2021-1111
  CVE-2021-3333
2022:
  CVE-2022-2222
```

### `filter dedup`

去重(大小写不敏感),保留首次出现的顺序。

```bash
$ cve filter dedup CVE-2022-1111 cve-2022-1111 CVE-2022-2222
CVE-2022-1111
CVE-2022-2222
```

### `filter-pattern`

使用 `*` 的通配符过滤。参数:`<pattern> <cve-list>`。

```bash
$ cve filter-pattern "CVE-2022-*" "CVE-2021-1,CVE-2022-2,CVE-2022-3"
CVE-2022-2
CVE-2022-3
```

## 集合运算

每个命令接收两个逗号分隔的列表。

```bash
$ cve intersect "CVE-2021-1,CVE-2022-2" "CVE-2022-2,CVE-2023-3"
CVE-2022-2

$ cve union "CVE-2021-1,CVE-2022-2" "CVE-2022-2,CVE-2023-3"
CVE-2021-1
CVE-2022-2
CVE-2023-3

$ cve diff "CVE-2021-1,CVE-2022-2" "CVE-2022-2,CVE-2023-3"
CVE-2021-1
```

## 范围与模式

### `parse-range`

将范围表达式展开为逐个 CVE。支持三种分隔符语法:

| 语法 | 右侧内容 | 示例 |
|------|----------|------|
| `to` | 完整 CVE | `CVE-2022-1 to CVE-2022-3` |
| `..` | 仅序列号 | `CVE-2022-1..4` |
| `-` | 仅序列号 | `CVE-2022-1 - 3` |

```bash
$ cve parse-range "CVE-2022-1..4"
CVE-2022-1
CVE-2022-2
CVE-2022-3
CVE-2022-4
```

### `is-consecutive`

检查两个 CVE 是否同年且序列号相邻。

```bash
$ cve is-consecutive CVE-2022-1 CVE-2022-2
CVE-2022-1 and CVE-2022-2 are consecutive
```

## 统计

### `count-by-year`

按年份统计 CVE 数量。

```bash
$ cve count-by-year "CVE-2021-1,CVE-2022-2,CVE-2021-3"
2021: 2
2022: 1
```

### `year-range`

报告最早与最晚年份及其跨度。

```bash
$ cve year-range "CVE-2019-1,CVE-2022-2,CVE-2025-3"
Year range: 2019 - 2025 (span: 6 years)
```

### `seq-range`

报告指定年份的序列号最小/最大值。参数:`<year> <cve-list>`。

```bash
$ cve seq-range 2022 "CVE-2022-100,CVE-2022-5000,CVE-2022-42"
Year 2022 sequence range: 42 - 5000
```

## 生成

### `generate cve`

根据 `--year`(`-y`)与 `--seq`(`-s`)构造一个 CVE。

```bash
$ cve generate cve --year 2022 --seq 12345
CVE-2022-12345
```

### `generate fake`

生成一个当前年份的随机 CVE(仅供测试/演示)。

```bash
$ cve generate fake
CVE-2026-56291
```

## 版本

```bash
$ cve version
```

打印构建版本(发布时由 goreleaser 注入;源码构建报告为 `dev`)。

## 另请参阅

- [API 参考](/zh/api/) —— 以类型化 Go 函数形式提供的相同能力。
- [下载与安装](/zh/download) —— 全部平台与包格式。
