package cve

import (
	"strconv"
	"time"
)

// GroupByYear 把一组CVE按照年份分组
//
// 将CVE列表按照年份进行分组
//
// 参数:
//   - cveSlice: 需要分组的CVE编号数组
//
// 返回值:
//   - map[string][]string: 分组结果，键为年份字符串（如"2021"），值为对应年份的CVE编号数组
//
// 示例:
//
//	输入: ["CVE-2021-1111", "CVE-2022-2222", "CVE-2021-3333"]
//	输出: {
//	  "2021": ["CVE-2021-1111", "CVE-2021-3333"],
//	  "2022": ["CVE-2022-2222"]
//	}
//
//	输入: ["CVE-2021-1111", "cve-2021-3333"] (不同大小写)
//	输出: {
//	  "2021": ["CVE-2021-1111", "CVE-2021-3333"]
//	} (注意返回结果已格式化为大写)
//
// 性能特性:
//   - 时间复杂度: O(n)，其中n为数组长度
//   - 空间复杂度: O(n)
//
// 使用场景:
//   - 按年份组织和展示多个CVE，例如生成年度漏洞报告
//   - 分析CVE随时间分布的趋势
//
// 代码示例:
//
//	cveList := []string{"CVE-2021-1111", "CVE-2022-2222", "CVE-2021-3333"}
//	yearGroups := cve.GroupByYear(cveList)
//	for year, cves := range yearGroups {
//	    fmt.Printf("%s年的CVE共有%d个\n", year, len(cves))
//	}
func GroupByYear(cveSlice []string) map[string][]string {
	groupMap := make(map[string][]string, 0)
	for _, cve := range cveSlice {
		year := ExtractCveYear(cve)
		groupMap[year] = append(groupMap[year], Format(cve))
	}
	return groupMap
}

// FilterCvesByYear 筛选特定年份的CVE
//
// 从CVE列表中筛选出指定年份的CVE编号
//
// 参数:
//   - cveSlice: 需要筛选的CVE编号数组
//   - year: 目标年份，整数格式，如2021
//
// 返回值:
//   - []string: 符合目标年份的CVE编号数组，已经过标准化格式处理
//     如果没有找到匹配项，则返回空数组
//
// 示例:
//
//	输入: ["CVE-2021-1111", "CVE-2022-2222", "CVE-2021-3333"], 2021
//	输出: ["CVE-2021-1111", "CVE-2021-3333"]
//
//	输入: ["CVE-2021-1111", "CVE-2022-2222"], 2023
//	输出: [] (空数组)
//
// 性能特性:
//   - 时间复杂度: O(n)，其中n为数组长度
//   - 空间复杂度: O(k)，其中k为结果数组长度（最坏情况为O(n)）
//
// 使用场景:
//   - 需要获取特定年份的CVE时使用
//   - 生成特定年度的安全报告
//
// 代码示例:
//
//	cveList := []string{"CVE-2021-1111", "CVE-2022-2222", "CVE-2021-3333"}
//	cves2021 := cve.FilterCvesByYear(cveList, 2021)
//	// cves2021 为 ["CVE-2021-1111", "CVE-2021-3333"]
func FilterCvesByYear(cveSlice []string, year int) []string {
	var result []string
	yearStr := strconv.Itoa(year)

	for _, cve := range cveSlice {
		formattedCve := Format(cve)
		if ExtractCveYear(formattedCve) == yearStr {
			result = append(result, formattedCve)
		}
	}

	return result
}

// FilterCvesByYearRange 筛选指定年份范围内的CVE
//
// 从CVE列表中筛选出在指定年份范围内的CVE编号
//
// 参数:
//   - cveSlice: 需要筛选的CVE编号数组
//   - startYear: 起始年份（含），整数格式
//   - endYear: 结束年份（含），整数格式
//
// 返回值:
//   - []string: 符合年份范围的CVE编号数组，已经过标准化格式处理
//     如果没有找到匹配项，则返回空数组
//
// 示例:
//
//	输入: ["CVE-2020-1111", "CVE-2021-2222", "CVE-2022-3333"], 2021, 2022
//	输出: ["CVE-2021-2222", "CVE-2022-3333"]
//
//	输入: ["CVE-2020-1111", "CVE-2021-2222"], 2022, 2023
//	输出: [] (空数组)
//
//	输入: ["CVE-2020-1111", "CVE-2021-2222", "CVE-2022-3333"], 2020, 2022
//	输出: ["CVE-2020-1111", "CVE-2021-2222", "CVE-2022-3333"]
//
// 性能特性:
//   - 时间复杂度: O(n)，其中n为数组长度
//   - 空间复杂度: O(k)，其中k为结果数组长度（最坏情况为O(n)）
//
// 使用场景:
//   - 需要获取一段时间内的CVE时使用
//   - 分析特定时间段内的安全漏洞
//
// 代码示例:
//
//	cveList := []string{"CVE-2020-1111", "CVE-2021-2222", "CVE-2022-3333"}
//	recentCves := cve.FilterCvesByYearRange(cveList, 2021, 2022)
//	// recentCves 为 ["CVE-2021-2222", "CVE-2022-3333"]
func FilterCvesByYearRange(cveSlice []string, startYear, endYear int) []string {
	var result []string

	for _, cve := range cveSlice {
		formattedCve := Format(cve)
		yearInt := ExtractCveYearAsInt(formattedCve)
		if yearInt >= startYear && yearInt <= endYear {
			result = append(result, formattedCve)
		}
	}

	return result
}

// GetRecentCves 获取最近n年的CVE
//
// 从CVE列表中获取最近几年内的CVE编号
//
// 参数:
//   - cveSlice: 需要筛选的CVE编号数组
//   - years: 最近几年的范围，整数，例如2表示最近两年
//
// 返回值:
//   - []string: 最近几年的CVE编号数组，已经过标准化格式处理
//     如果没有找到匹配项，则返回空数组
//
// 计算规则:
//
//	以当前年份为基准，获取(当前年份-years+1)到当前年份之间的所有CVE
//
// 示例:
//
//	假设当前年份为2023
//	输入: ["CVE-2020-1111", "CVE-2021-2222", "CVE-2022-3333", "CVE-2023-4444"], 2
//	输出: ["CVE-2022-3333", "CVE-2023-4444"] (2022和2023年的CVE)
//
//	输入: ["CVE-2020-1111", "CVE-2021-2222"], 1
//	输出: [] (空数组，因为没有2023年的CVE)
//
// 使用场景:
//   - 需要关注最近几年发布的CVE时使用
//   - 生成最新安全威胁报告
//
// 代码示例:
//
//	cveList := []string{"CVE-2020-1111", "CVE-2021-2222", "CVE-2022-3333", "CVE-2023-4444"}
//	recentTwoYears := cve.GetRecentCves(cveList, 2)
//	// 如果当前是2023年，recentTwoYears 为 ["CVE-2022-3333", "CVE-2023-4444"]
func GetRecentCves(cveSlice []string, years int) []string {
	currentYear := time.Now().Year()
	return FilterCvesByYearRange(cveSlice, currentYear-years+1, currentYear)
}

// IntersectCves 求两个CVE列表的交集
//
// 返回同时出现在两个列表中的CVE编号，结果已去重和格式化为大写
//
// 参数:
//   - a: 第一个CVE列表
//   - b: 第二个CVE列表
//
// 返回值:
//   - []string: 两个列表的共有CVE编号，已排序
//
// 比较规则:
//   - CVE比较时不区分大小写
//   - 结果按年份和序列号升序排列
//
// 示例:
//
//	输入: ["CVE-2022-1111", "CVE-2022-2222"], ["CVE-2022-2222", "CVE-2022-3333"]
//	输出: ["CVE-2022-2222"]
//
//	输入: ["CVE-2022-1111"], ["CVE-2023-2222"]
//	输出: [] (空数组)
//
// 性能特性:
//   - 时间复杂度: O(n+m)，其中n和m分别为两个列表的长度
//   - 空间复杂度: O(min(n,m))
//
// 使用场景:
//   - 多源CVE数据交叉对比
//   - 找出多个安全报告中共同提及的漏洞
//
// 代码示例:
//
//	list1 := []string{"CVE-2022-1111", "CVE-2022-2222"}
//	list2 := []string{"CVE-2022-2222", "CVE-2022-3333"}
//	common := cve.IntersectCves(list1, list2)
//	// common 为 ["CVE-2022-2222"]
func IntersectCves(a, b []string) []string {
	set := make(map[string]struct{}, len(a))
	for _, cve := range a {
		set[Format(cve)] = struct{}{}
	}

	var result []string
	seen := make(map[string]struct{}, len(b))
	for _, cve := range b {
		formatted := Format(cve)
		if _, inA := set[formatted]; inA {
			if _, exists := seen[formatted]; !exists {
				seen[formatted] = struct{}{}
				result = append(result, formatted)
			}
		}
	}

	return SortCves(result)
}

// UnionCves 求两个CVE列表的并集
//
// 合并两个列表中的所有CVE编号，结果已去重和格式化为大写
//
// 参数:
//   - a: 第一个CVE列表
//   - b: 第二个CVE列表
//
// 返回值:
//   - []string: 两个列表的所有CVE编号（去重），已排序
//
// 比较规则:
//   - CVE比较时不区分大小写
//   - 结果按年份和序列号升序排列
//
// 示例:
//
//	输入: ["CVE-2022-1111", "CVE-2022-2222"], ["CVE-2022-2222", "CVE-2022-3333"]
//	输出: ["CVE-2022-1111", "CVE-2022-2222", "CVE-2022-3333"]
//
// 性能特性:
//   - 时间复杂度: O(n+m)，其中n和m分别为两个列表的长度
//   - 空间复杂度: O(n+m)
//
// 使用场景:
//   - 合并多个来源的CVE列表
//   - 整合多方安全通告中的漏洞信息
//
// 代码示例:
//
//	list1 := []string{"CVE-2022-1111", "CVE-2022-2222"}
//	list2 := []string{"CVE-2022-2222", "CVE-2022-3333"}
//	all := cve.UnionCves(list1, list2)
//	// all 为 ["CVE-2022-1111", "CVE-2022-2222", "CVE-2022-3333"]
func UnionCves(a, b []string) []string {
	set := make(map[string]struct{}, len(a)+len(b))
	var result []string

	for _, cve := range a {
		formatted := Format(cve)
		if _, exists := set[formatted]; !exists {
			set[formatted] = struct{}{}
			result = append(result, formatted)
		}
	}

	for _, cve := range b {
		formatted := Format(cve)
		if _, exists := set[formatted]; !exists {
			set[formatted] = struct{}{}
			result = append(result, formatted)
		}
	}

	return SortCves(result)
}

// DiffCves 求两个CVE列表的差集（a有b无）
//
// 返回在列表a中出现但不在列表b中出现的CVE编号
//
// 参数:
//   - a: 被减CVE列表
//   - b: 需要排除的CVE列表
//
// 返回值:
//   - []string: 只在列表a中出现的CVE编号，已去重和格式化，已排序
//
// 比较规则:
//   - CVE比较时不区分大小写
//   - 结果按年份和序列号升序排列
//
// 示例:
//
//	输入: a=["CVE-2022-1111", "CVE-2022-2222"], b=["CVE-2022-2222", "CVE-2022-3333"]
//	输出: ["CVE-2022-1111"]
//
//	输入: a=["CVE-2022-1111","CVE-2022-1111"], b=["CVE-2022-3333"]
//	输出: ["CVE-2022-1111"] (注意输入a中的重复已被去除)
//
// 性能特性:
//   - 时间复杂度: O(n+m)，其中n和m分别为两个列表的长度
//   - 空间复杂度: O(n+m)
//
// 使用场景:
//   - 找出新增的CVE（与历史数据对比）
//   - 检测某个列表中独有的漏洞
//
// 代码示例:
//
//	current := []string{"CVE-2022-1111", "CVE-2022-2222"}
//	historical := []string{"CVE-2022-2222", "CVE-2022-3333"}
//	newCves := cve.DiffCves(current, historical)
//	// newCves 为 ["CVE-2022-1111"] — 历史列表中未出现的新CVE
func DiffCves(a, b []string) []string {
	bSet := make(map[string]struct{}, len(b))
	for _, cve := range b {
		bSet[Format(cve)] = struct{}{}
	}

	aSeen := make(map[string]struct{}, len(a))
	var result []string
	for _, cve := range a {
		formatted := Format(cve)
		if _, inB := bSet[formatted]; !inB {
			if _, exists := aSeen[formatted]; !exists {
				aSeen[formatted] = struct{}{}
				result = append(result, formatted)
			}
		}
	}

	return SortCves(result)
}

// RemoveDuplicateCves 移除重复的CVE编号
//
// 去除CVE列表中的重复项，保留唯一的CVE编号（不区分大小写）
//
// 参数:
//   - cveSlice: 包含可能重复CVE的数组
//
// 返回值:
//   - []string: 去重后的CVE编号数组，所有CVE均已标准化格式（大写）
//
// 处理规则:
//   - CVE比较时不区分大小写，例如"CVE-2022-1111"和"cve-2022-1111"被视为重复
//   - 只保留每个CVE的第一次出现
//   - 所有返回的CVE均为标准化格式（大写）
//
// 示例:
//
//	输入: ["CVE-2022-1111", "cve-2022-1111", "CVE-2022-2222"]
//	输出: ["CVE-2022-1111", "CVE-2022-2222"]
//
//	输入: ["CVE-2022-1111", "CVE-2022-1111", "CVE-2022-1111"]
//	输出: ["CVE-2022-1111"]
//
// 性能特性:
//   - 时间复杂度: O(n)，其中n为数组长度
//   - 空间复杂度: O(n)
//
// 使用场景:
//   - 合并多个来源的CVE列表并去重时使用
//   - 在处理大量CVE数据前进行预处理
//
// 代码示例:
//
//	cveList := []string{"CVE-2022-1111", "cve-2022-1111", "CVE-2022-2222"}
//	uniqueCves := cve.RemoveDuplicateCves(cveList)
//	// uniqueCves 为 ["CVE-2022-1111", "CVE-2022-2222"]
func RemoveDuplicateCves(cveSlice []string) []string {
	cveMap := make(map[string]struct{})
	var result []string

	for _, cve := range cveSlice {
		formattedCve := Format(cve)
		if _, exists := cveMap[formattedCve]; !exists {
			cveMap[formattedCve] = struct{}{}
			result = append(result, formattedCve)
		}
	}

	return result
}

// CountByYear 统计CVE列表中各年份的数量
//
// 对CVE列表按年份进行计数，返回年份到数量的映射
//
// 参数:
//   - cveSlice: 需要统计的CVE编号数组
//
// 返回值:
//   - map[int]int: 年份到CVE数量的映射，key为年份，value为该年份的CVE数量
//
// 示例:
//
//	输入: ["CVE-2022-1111", "CVE-2022-2222", "CVE-2021-3333", "cve-2022-4444"]
//	输出: {2021: 1, 2022: 3}
//
// 使用场景:
//   - CVE趋势分析：了解各年份的漏洞分布
//   - 安全报告：生成年度CVE统计
//
// 代码示例:
//
//	counts := cve.CountByYear(cveList)
//	for year, count := range counts {
//	    fmt.Printf("%d: %d CVEs\n", year, count)
//	}
func CountByYear(cveSlice []string) map[int]int {
	result := make(map[int]int)
	for _, cve := range cveSlice {
		year := ExtractCveYearAsInt(cve)
		if year > 0 {
			result[year]++
		}
	}
	return result
}

// YearRange 获取CVE列表中最早和最晚的年份
//
// 返回CVE列表中年份的最小值和最大值
//
// 参数:
//   - cveSlice: CVE编号数组
//
// 返回值:
//   - min: 最早的年份（最小值），如果列表为空或没有有效CVE则返回0
//   - max: 最晚的年份（最大值），如果列表为空或没有有效CVE则返回0
//
// 示例:
//
//	输入: ["CVE-2020-1111", "CVE-2022-2222", "CVE-2021-3333"]
//	输出: min=2020, max=2022
//
//	输入: [] (空数组)
//	输出: min=0, max=0
//
// 使用场景:
//   - 确定CVE数据的时间跨度
//   - 生成时间范围描述
//
// 代码示例:
//
//	minYear, maxYear := cve.YearRange(cveList)
//	fmt.Printf("CVEs span from %d to %d\n", minYear, maxYear)
func YearRange(cveSlice []string) (min, max int) {
	if len(cveSlice) == 0 {
		return 0, 0
	}

	min = -1
	for _, cve := range cveSlice {
		year := ExtractCveYearAsInt(cve)
		if year <= 0 {
			continue
		}
		if min == -1 || year < min {
			min = year
		}
		if year > max {
			max = year
		}
	}

	if min == -1 {
		return 0, 0
	}
	return min, max
}

// SeqRange 获取指定年份下CVE序列号的范围
//
// 从CVE列表中筛选出指定年份的CVE，然后返回其序列号的最小值和最大值
//
// 参数:
//   - cveSlice: CVE编号数组
//   - year: 目标年份
//
// 返回值:
//   - min: 该年份下的最小序列号，如果没有找到则返回0
//   - max: 该年份下的最大序列号，如果没有找到则返回0
//
// 示例:
//
//	输入: ["CVE-2022-1111", "CVE-2022-5555", "CVE-2022-3333", "CVE-2021-9999"], 2022
//	输出: min=1111, max=5555
//
//	输入: ["CVE-2022-1111"], 2023
//	输出: min=0, max=0
//
// 使用场景:
//   - 了解某个年份CVE编号的分配范围
//   - 估算CVE密度的辅助信息
//
// 代码示例:
//
//	minSeq, maxSeq := cve.SeqRange(cveList, 2022)
//	fmt.Printf("2022年CVE序列号范围: %d - %d\n", minSeq, maxSeq)
func SeqRange(cveSlice []string, year int) (min, max int) {
	min = -1
	for _, cve := range cveSlice {
		cveYear := ExtractCveYearAsInt(cve)
		if cveYear != year {
			continue
		}
		seq := ExtractCveSeqAsInt(cve)
		if seq <= 0 {
			continue
		}
		if min == -1 || seq < min {
			min = seq
		}
		if seq > max {
			max = seq
		}
	}

	if min == -1 {
		return 0, 0
	}
	return min, max
}
