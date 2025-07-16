# CVE工具包使用示例

本目录包含了 `github.com/scagogogo/cve` 包的所有API函数的使用示例。每个示例都是一个独立的Go程序，展示了相应函数的基本用法和应用场景。

## 示例目录

1. **01_format**: 演示格式化CVE编号的方法
2. **02_is_cve**: 演示如何验证字符串是否为标准CVE编号
3. **03_is_contains_cve**: 演示如何检查文本是否包含CVE编号
4. **04_extract_cve**: 演示从文本中提取所有CVE编号的方法
5. **05_extract_first_last_cve**: 演示提取第一个和最后一个CVE编号的方法
6. **06_extract_year_seq**: 演示从CVE中提取年份和序列号的方法
7. **07_validate_cve_year**: 演示验证CVE年份有效性的方法
8. **08_validate_cve**: 演示全面验证CVE有效性的方法
9. **09_compare_by_year**: 演示比较CVE年份的方法
10. **10_compare_cves**: 演示完整比较两个CVE的方法
11. **11_sort_cves**: 演示排序CVE列表的方法
12. **12_group_by_year**: 演示按年份分组CVE的方法
13. **13_filter_by_year**: 演示按特定年份过滤CVE的方法
14. **14_filter_by_year_range**: 演示按年份范围过滤CVE的方法
15. **15_get_recent_cves**: 演示获取最近几年CVE的方法
16. **16_remove_duplicate_cves**: 演示去除重复CVE的方法
17. **17_generate_cve**: 演示生成标准CVE编号的方法
18. **18_generate_fake_cve**: 演示生成随机测试用CVE的方法

## 运行示例

每个示例都可以独立运行。进入对应目录后执行以下命令：

```bash
go run main.go
```

## 示例功能

每个示例程序都包含以下几个部分：

1. 基本用法演示：展示函数的基本调用方式和参数传递
2. 多种情况测试：展示函数在不同输入下的行为
3. 实际应用场景：展示函数在真实场景中的使用方法

## 安装依赖

运行这些示例前，请确保已安装 `github.com/scagogogo/cve` 包：

```bash
go get github.com/scagogogo/cve
```

## 版本要求

- Go 1.18 或更高版本 