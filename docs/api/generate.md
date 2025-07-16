# 生成与构造

这一类函数用于生成新的 CVE 编号，主要用于测试、模拟和 CVE 编号的动态创建。

## GenerateCve

根据年份和序列号生成标准格式的 CVE 编号。

### 函数签名

```go
func GenerateCve(year int, seq int) string
```

### 参数

- `year` (int): CVE 年份
- `seq` (int): CVE 序列号

### 返回值

- `string`: 生成的标准格式 CVE 编号

### 功能描述

`GenerateCve` 函数通过给定的年份和序列号创建标准的 CVE 编号：
- 自动格式化为标准的 `CVE-YYYY-NNNN` 格式
- 不验证年份和序列号的有效性
- 内部使用 `Format` 函数确保标准格式

### 使用示例

```go
func main() {
    // 基本使用
    cveId := cve.GenerateCve(2022, 12345)
    fmt.Printf("生成的 CVE: %s\n", cveId)
    // 输出: CVE-2022-12345
    
    // 批量生成
    fmt.Println("\n批量生成示例:")
    for i := 1; i <= 5; i++ {
        cveId := cve.GenerateCve(2024, i)
        fmt.Printf("CVE #%d: %s\n", i, cveId)
    }
    
    // 不同年份的生成
    years := []int{2020, 2021, 2022, 2023, 2024}
    fmt.Println("\n不同年份:")
    for _, year := range years {
        cveId := cve.GenerateCve(year, 1000)
        fmt.Printf("%d年: %s\n", year, cveId)
    }
    
    // 不同序列号长度
    sequences := []int{1, 10, 100, 1000, 10000, 100000}
    fmt.Println("\n不同序列号长度:")
    for _, seq := range sequences {
        cveId := cve.GenerateCve(2024, seq)
        fmt.Printf("序列号 %6d: %s\n", seq, cveId)
    }
}
```

### 使用场景

- **测试数据生成**: 为单元测试创建 CVE 数据
- **模拟和演示**: 生成示例 CVE 用于演示
- **数据填充**: 为开发环境创建测试数据
- **CVE 编号预分配**: 在正式分配前生成占位符
- **批量数据生成**: 创建大量测试用的 CVE

### 注意事项

- 该函数不验证年份和序列号的合理性
- 可以生成无效的 CVE（如年份为负数）
- 建议与验证函数配合使用
- 生成的 CVE 可能与真实 CVE 冲突

## 实际应用示例

### 1. 测试数据生成器

```go
// CVE 测试数据生成器
type CveGenerator struct {
    BaseYear     int
    YearRange    int
    MaxSequence  int
    MinSequence  int
}

func NewCveGenerator() *CveGenerator {
    return &CveGenerator{
        BaseYear:    2020,
        YearRange:   5,     // 2020-2024
        MaxSequence: 99999,
        MinSequence: 1,
    }
}

func (g *CveGenerator) GenerateRandom(count int) []string {
    var cves []string
    
    for i := 0; i < count; i++ {
        year := g.BaseYear + rand.Intn(g.YearRange)
        seq := g.MinSequence + rand.Intn(g.MaxSequence-g.MinSequence+1)
        
        cveId := cve.GenerateCve(year, seq)
        cves = append(cves, cveId)
    }
    
    return cves
}

func (g *CveGenerator) GenerateSequential(year, startSeq, count int) []string {
    var cves []string
    
    for i := 0; i < count; i++ {
        seq := startSeq + i
        cveId := cve.GenerateCve(year, seq)
        cves = append(cves, cveId)
    }
    
    return cves
}

func (g *CveGenerator) GenerateByYear(yearCounts map[int]int) []string {
    var cves []string
    
    for year, count := range yearCounts {
        for i := 1; i <= count; i++ {
            cveId := cve.GenerateCve(year, i)
            cves = append(cves, cveId)
        }
    }
    
    return cves
}

func main() {
    generator := NewCveGenerator()
    
    // 生成随机 CVE
    randomCves := generator.GenerateRandom(10)
    fmt.Printf("随机生成 (%d个): %v\n", len(randomCves), randomCves)
    
    // 生成连续序列
    sequential := generator.GenerateSequential(2024, 1000, 5)
    fmt.Printf("连续序列: %v\n", sequential)
    
    // 按年份生成
    yearCounts := map[int]int{
        2022: 3,
        2023: 5,
        2024: 2,
    }
    byYear := generator.GenerateByYear(yearCounts)
    fmt.Printf("按年份生成: %v\n", byYear)
}
```

### 2. 单元测试辅助函数

```go
// 测试辅助函数
func createTestCves(years []int, seqsPerYear int) []string {
    var testCves []string
    
    for _, year := range years {
        for seq := 1; seq <= seqsPerYear; seq++ {
            cveId := cve.GenerateCve(year, seq)
            testCves = append(testCves, cveId)
        }
    }
    
    return testCves
}

func createMixedFormatCves(year, count int) []string {
    var cves []string
    
    for i := 1; i <= count; i++ {
        cveId := cve.GenerateCve(year, i)
        
        // 随机改变格式用于测试
        switch i % 3 {
        case 0:
            cves = append(cves, strings.ToLower(cveId)) // 小写
        case 1:
            cves = append(cves, " "+cveId+" ")          // 添加空格
        default:
            cves = append(cves, cveId)                  // 标准格式
        }
    }
    
    return cves
}

// 使用示例
func TestCveProcessing(t *testing.T) {
    // 创建测试数据
    testYears := []int{2020, 2021, 2022}
    testCves := createTestCves(testYears, 3)
    
    // 测试排序功能
    sorted := cve.SortCves(testCves)
    if len(sorted) != len(testCves) {
        t.Errorf("排序后长度不匹配")
    }
    
    // 测试分组功能
    grouped := cve.GroupByYear(testCves)
    if len(grouped) != len(testYears) {
        t.Errorf("分组数量不匹配")
    }
    
    // 测试格式处理
    mixedCves := createMixedFormatCves(2024, 6)
    unique := cve.RemoveDuplicateCves(mixedCves)
    if len(unique) != 6 {
        t.Errorf("去重结果不正确")
    }
}
```

### 3. 性能测试数据生成

```go
func generatePerformanceTestData(size int) []string {
    fmt.Printf("生成 %d 个 CVE 用于性能测试...\n", size)
    
    var cves []string
    startTime := time.Now()
    
    // 生成大量 CVE 数据
    for i := 0; i < size; i++ {
        year := 2020 + (i % 5)        // 2020-2024 循环
        seq := i + 1
        
        cveId := cve.GenerateCve(year, seq)
        cves = append(cves, cveId)
    }
    
    duration := time.Since(startTime)
    fmt.Printf("生成完成，耗时: %v\n", duration)
    
    return cves
}

func benchmarkCveOperations(cves []string) {
    fmt.Println("=== CVE 操作性能测试 ===")
    
    // 测试排序性能
    start := time.Now()
    sorted := cve.SortCves(cves)
    sortDuration := time.Since(start)
    fmt.Printf("排序 %d 个 CVE 耗时: %v\n", len(cves), sortDuration)
    
    // 测试去重性能
    start = time.Now()
    unique := cve.RemoveDuplicateCves(sorted)
    dedupDuration := time.Since(start)
    fmt.Printf("去重 %d 个 CVE 耗时: %v\n", len(sorted), dedupDuration)
    
    // 测试分组性能
    start = time.Now()
    grouped := cve.GroupByYear(unique)
    groupDuration := time.Since(start)
    fmt.Printf("分组 %d 个 CVE 耗时: %v\n", len(unique), groupDuration)
    
    fmt.Printf("分组结果: %d 个年份\n", len(grouped))
}

func main() {
    // 生成不同规模的测试数据
    sizes := []int{1000, 10000, 100000}
    
    for _, size := range sizes {
        fmt.Printf("\n=== 测试规模: %d ===\n", size)
        testData := generatePerformanceTestData(size)
        benchmarkCveOperations(testData)
    }
}
```

### 4. 模拟真实场景

```go
// 模拟漏洞发现场景
type VulnerabilitySimulator struct {
    CurrentYear int
    NextSeq     map[int]int // 每年的下一个序列号
}

func NewVulnerabilitySimulator() *VulnerabilitySimulator {
    return &VulnerabilitySimulator{
        CurrentYear: time.Now().Year(),
        NextSeq:     make(map[int]int),
    }
}

func (vs *VulnerabilitySimulator) DiscoverVulnerability(year int) string {
    if vs.NextSeq[year] == 0 {
        vs.NextSeq[year] = 1
    }
    
    cveId := cve.GenerateCve(year, vs.NextSeq[year])
    vs.NextSeq[year]++
    
    return cveId
}

func (vs *VulnerabilitySimulator) SimulateYear(year int, count int) []string {
    var discoveries []string
    
    for i := 0; i < count; i++ {
        cveId := vs.DiscoverVulnerability(year)
        discoveries = append(discoveries, cveId)
    }
    
    return discoveries
}

func (vs *VulnerabilitySimulator) GetStatistics() map[int]int {
    stats := make(map[int]int)
    for year, nextSeq := range vs.NextSeq {
        stats[year] = nextSeq - 1 // 减1因为nextSeq是下一个要分配的序列号
    }
    return stats
}

func main() {
    simulator := NewVulnerabilitySimulator()
    
    // 模拟多年的漏洞发现
    fmt.Println("=== 漏洞发现模拟 ===")
    
    allDiscoveries := []string{}
    
    // 模拟2022-2024年的漏洞发现
    yearlyCount := map[int]int{
        2022: 15,
        2023: 25,
        2024: 10,
    }
    
    for year, count := range yearlyCount {
        discoveries := simulator.SimulateYear(year, count)
        allDiscoveries = append(allDiscoveries, discoveries...)
        fmt.Printf("%d年发现 %d 个漏洞: %v\n", year, count, discoveries[:3]) // 只显示前3个
    }
    
    // 统计信息
    stats := simulator.GetStatistics()
    fmt.Println("\n统计信息:")
    for year, count := range stats {
        fmt.Printf("%d年: %d个漏洞\n", year, count)
    }
    
    // 分析生成的数据
    fmt.Printf("\n总计生成: %d个 CVE\n", len(allDiscoveries))
    grouped := cve.GroupByYear(allDiscoveries)
    fmt.Printf("分布在 %d 个年份\n", len(grouped))
}
```

### 5. CVE 编号验证器

```go
func validateGeneratedCves(cves []string) (valid, invalid []string) {
    for _, cveId := range cves {
        if cve.ValidateCve(cveId) {
            valid = append(valid, cveId)
        } else {
            invalid = append(invalid, cveId)
        }
    }
    return
}

func generateAndValidate(year, count int) {
    fmt.Printf("=== 生成并验证 %d年的 %d 个 CVE ===\n", year, count)
    
    // 生成 CVE
    var generated []string
    for i := 1; i <= count; i++ {
        cveId := cve.GenerateCve(year, i)
        generated = append(generated, cveId)
    }
    
    // 验证生成的 CVE
    valid, invalid := validateGeneratedCves(generated)
    
    fmt.Printf("生成: %d个\n", len(generated))
    fmt.Printf("有效: %d个\n", len(valid))
    fmt.Printf("无效: %d个\n", len(invalid))
    
    if len(invalid) > 0 {
        fmt.Printf("无效的 CVE: %v\n", invalid)
    }
    
    // 测试边界情况
    fmt.Println("\n边界情况测试:")
    
    // 测试无效年份
    invalidYear := cve.GenerateCve(-1, 1)
    fmt.Printf("负年份: %s (有效: %t)\n", invalidYear, cve.ValidateCve(invalidYear))
    
    // 测试零序列号
    zeroSeq := cve.GenerateCve(2024, 0)
    fmt.Printf("零序列号: %s (有效: %t)\n", zeroSeq, cve.ValidateCve(zeroSeq))
    
    // 测试大序列号
    largeSeq := cve.GenerateCve(2024, 999999)
    fmt.Printf("大序列号: %s (有效: %t)\n", largeSeq, cve.ValidateCve(largeSeq))
}

func main() {
    generateAndValidate(2024, 5)
}
```

## 最佳实践

### 1. 与验证函数配合使用

```go
func generateValidCve(year, seq int) (string, error) {
    cveId := cve.GenerateCve(year, seq)
    
    if !cve.ValidateCve(cveId) {
        return "", fmt.Errorf("生成的 CVE 无效: %s", cveId)
    }
    
    return cveId, nil
}
```

### 2. 批量生成优化

```go
func generateCveBatch(year, startSeq, count int) []string {
    cves := make([]string, count)
    
    for i := 0; i < count; i++ {
        cves[i] = cve.GenerateCve(year, startSeq+i)
    }
    
    return cves
}
```

### 3. 避免冲突

```go
func generateUniqueCves(existingCves []string, year, count int) []string {
    existing := make(map[string]bool)
    for _, cveId := range existingCves {
        existing[cveId] = true
    }
    
    var newCves []string
    seq := 1
    
    for len(newCves) < count {
        candidate := cve.GenerateCve(year, seq)
        if !existing[candidate] {
            newCves = append(newCves, candidate)
            existing[candidate] = true
        }
        seq++
    }
    
    return newCves
}
```

## 性能说明

- `GenerateCve` 函数性能很高，主要开销在字符串格式化
- 批量生成时建议预分配切片容量
- 对于大量数据生成，考虑使用 goroutine 并行处理
- 内存使用量与生成的 CVE 数量成正比

## 注意事项

1. **不验证有效性**: `GenerateCve` 不检查年份和序列号的合理性
2. **可能冲突**: 生成的 CVE 可能与真实 CVE 冲突
3. **仅用于测试**: 不应用于生产环境的真实 CVE 分配
4. **格式保证**: 始终生成标准格式的 CVE 编号
