import { Layout, Typography, Button, Row, Col, Card, Tag, Divider, Space, Tabs } from 'antd'
import {
  GithubOutlined,
  SafetyCertificateOutlined,
  ThunderboltOutlined,
  ApiOutlined,
  CodeOutlined,
  ToolOutlined,
  CheckCircleOutlined,
  ArrowRightOutlined,
  BugOutlined,
  FilterOutlined,
  SwapOutlined,
  BarChartOutlined,
  CopyOutlined,
  RocketOutlined,
} from '@ant-design/icons'
import './HomePage.css'

const { Header, Content, Footer } = Layout
const { Title, Paragraph, Text } = Typography

const features = [
  {
    icon: <SafetyCertificateOutlined style={{ fontSize: 32 }} />,
    title: '格式化与验证',
    desc: '标准化 CVE 编号格式，全面验证（格式+年份+序列号），支持未来年份偏移检测',
    tags: ['Format()', 'ValidateCve()', 'IsCveYearOk()'],
    color: '#e94560',
  },
  {
    icon: <CopyOutlined style={{ fontSize: 32 }} />,
    title: '智能提取',
    desc: '从安全公告、NVD 数据源和漏洞报告中一键提取 CVE 编号，支持首尾提取与拆分',
    tags: ['ExtractCve()', 'Split()', 'ExtractCveSeq()'],
    color: '#4ecdc4',
  },
  {
    icon: <SwapOutlined style={{ fontSize: 32 }} />,
    title: '比较与排序',
    desc: '原生支持 CVE 编号的比较与排序，按年份或序列号灵活比较',
    tags: ['CompareCves()', 'SortCves()', 'SubByYear()'],
    color: '#45b7d1',
  },
  {
    icon: <FilterOutlined style={{ fontSize: 32 }} />,
    title: '过滤与分组',
    desc: '按年份、年份范围过滤，自动去重，按年份分组，获取最近 N 年的 CVE',
    tags: ['FilterByYear()', 'GroupByYear()', 'RemoveDuplicate()'],
    color: '#96ceb4',
  },
  {
    icon: <BugOutlined style={{ fontSize: 32 }} />,
    title: '生成与构造',
    desc: '根据年份和序列号生成 CVE，一键生成测试用伪造 CVE，序列号补零格式化',
    tags: ['GenerateCve()', 'GenerateFakeCve()', 'FormatSeq()'],
    color: '#f9ca24',
  },
  {
    icon: <ToolOutlined style={{ fontSize: 32 }} />,
    title: '集合运算',
    desc: '对两个 CVE 列表执行交集、并集、差集运算，快速发现新增或共同漏洞',
    tags: ['IntersectCves()', 'UnionCves()', 'DiffCves()'],
    color: '#a29bfe',
  },
  {
    icon: <CodeOutlined style={{ fontSize: 32 }} />,
    title: '批量验证',
    desc: '批量验证 CVE 编号并返回详细错误原因，一键过滤出有效 CVE',
    tags: ['ValidateCves()', 'FilterValidCves()'],
    color: '#fd79a8',
  },
  {
    icon: <RocketOutlined style={{ fontSize: 32 }} />,
    title: '范围与模式',
    desc: '解析 CVE 范围表达式，检查连续性，通配符模式匹配',
    tags: ['ParseCveRange()', 'IsCvesConsecutive()', 'FilterByPattern()'],
    color: '#fab1a0',
  },
  {
    icon: <BarChartOutlined style={{ fontSize: 32 }} />,
    title: '统计分析',
    desc: '按年份统计 CVE 数量，获取年份跨度和序列号范围',
    tags: ['CountByYear()', 'YearRange()', 'SeqRange()'],
    color: '#00cec9',
  },
]

const cliExamples = [
  {
    key: 'format',
    label: '格式化',
    children: (
      <pre className="code-block">{`# 格式化 CVE 编号为标准大写格式
$ cve format CVE-2024-0001 cve-2023-54321
CVE-2024-0001
CVE-2023-54321`}</pre>
    ),
  },
  {
    key: 'validate',
    label: '验证',
    children: (
      <pre className="code-block">{`# 验证 CVE 编号是否合法
$ cve validate CVE-2024-0001 CVE-1998-12345
CVE-2024-0001    true
CVE-1998-12345  false  (年份不合法: 1998 < 1999)`}</pre>
    ),
  },
  {
    key: 'extract',
    label: '提取',
    children: (
      <pre className="code-block">{`# 从文本中提取 CVE 编号
$ cve extract "受 CVE-2021-44228 和 CVE-2022-12345 影响"
CVE-2021-44228
CVE-2022-12345`}</pre>
    ),
  },
  {
    key: 'filter',
    label: '过滤',
    children: (
      <pre className="code-block">{`# 按年份过滤 CVE
$ cve filter by-year --year 2022 CVE-2021-1111 CVE-2022-2222 CVE-2023-3333
CVE-2022-2222`}</pre>
    ),
  },
  {
    key: 'set',
    label: '集合运算',
    children: (
      <pre className="code-block">{`# 求两个 CVE 列表的交集
$ cve intersect "CVE-2022-1111,CVE-2022-2222" "CVE-2022-2222,CVE-2022-3333"
CVE-2022-2222`}</pre>
    ),
  },
  {
    key: 'stats',
    label: '统计',
    children: (
      <pre className="code-block">{`# 按年份统计 CVE 数量
$ cve count-by-year "CVE-2022-1111,CVE-2022-2222,CVE-2021-3333"
2021: 1
2022: 2`}</pre>
    ),
  },
]

const goCodeExample = `package main

import (
    "fmt"
    "github.com/scagogogo/cve-skills"
)

func main() {
    // 1. 格式化与验证
    formatted := cve.Format("cve-2022-12345")
    fmt.Println(formatted) // CVE-2022-12345

    isValid := cve.ValidateCve("CVE-2022-12345")
    fmt.Println(isValid) // true

    // 2. 从文本提取
    text := "受 CVE-2021-44228 和 CVE-2022-12345 影响"
    cves := cve.ExtractCve(text)
    fmt.Println(cves) // [CVE-2021-44228 CVE-2022-12345]

    // 3. 排序与过滤
    list := []string{"CVE-2022-3333", "CVE-2020-1111", "CVE-2022-1111"}
    sorted := cve.SortCves(list)
    fmt.Println(sorted) // [CVE-2020-1111 CVE-2022-1111 CVE-2022-3333]

    // 4. 集合运算
    common := cve.IntersectCves(
        []string{"CVE-2022-1111", "CVE-2022-2222"},
        []string{"CVE-2022-2222", "CVE-2022-3333"},
    )
    fmt.Println(common) // [CVE-2022-2222]
}`

function HomePage() {
  return (
    <Layout className="site-layout">
      {/* ====== Header / Nav ====== */}
      <Header className="site-header">
        <div className="header-inner">
          <div className="logo">
            <SafetyCertificateOutlined style={{ fontSize: 24, marginRight: 10 }} />
            <span className="logo-text">CVE Skills</span>
          </div>
          <Space size="large">
            <a href="#features" className="nav-link">功能特性</a>
            <a href="#architecture" className="nav-link">架构</a>
            <a href="#cli" className="nav-link">CLI</a>
            <a href="#quickstart" className="nav-link">快速开始</a>
            <Button
              type="primary"
              icon={<GithubOutlined />}
              href="https://github.com/scagogogo/cve-skills"
              target="_blank"
            >
              GitHub
            </Button>
          </Space>
        </div>
      </Header>

      <Content>
        {/* ====== Hero Section ====== */}
        <section className="hero-section">
          <div className="hero-content">
            <div className="hero-badge">
              <Tag color="red">Go Library</Tag>
              <Tag color="blue">CLI Tool</Tag>
              <Tag color="green">30+ Functions</Tag>
            </div>
            <Title level={1} className="hero-title">
              CVE 标识符处理的<span className="highlight">全能工具集</span>
            </Title>
            <Paragraph className="hero-desc">
              一个全面的 Go 语言库和命令行工具，用于处理 CVE（通用漏洞披露）标识符。
              从格式验证到集合运算，从文本提取到统计分析，一个依赖解决所有问题。
            </Paragraph>
            <Space size="large" className="hero-actions">
              <Button type="primary" size="large" icon={<ArrowRightOutlined />} href="#quickstart">
                快速开始
              </Button>
              <Button size="large" icon={<ApiOutlined />} href="https://pkg.go.dev/github.com/scagogogo/cve-skills" target="_blank">
                API 文档
              </Button>
            </Space>
            <div className="hero-stats">
              <div className="stat-item">
                <Text className="stat-number">30+</Text>
                <Text className="stat-label">工具函数</Text>
              </div>
              <div className="stat-item">
                <Text className="stat-number">95%+</Text>
                <Text className="stat-label">测试覆盖率</Text>
              </div>
              <div className="stat-item">
                <Text className="stat-number">20+</Text>
                <Text className="stat-label">CLI 命令</Text>
              </div>
              <div className="stat-item">
                <Text className="stat-number">MIT</Text>
                <Text className="stat-label">开源协议</Text>
              </div>
            </div>
          </div>
        </section>

        {/* ====== Problem Section ====== */}
        <section className="problem-section" id="problem">
          <div className="section-inner">
            <Title level={2} className="section-title">
              <ThunderboltOutlined /> 它解决了什么问题？
            </Title>
            <Row gutter={[24, 24]}>
              <Col xs={24} md={12}>
                <Card className="problem-card" hoverable>
                  <Title level={4}>😰 没有它的时候</Title>
                  <ul className="problem-list">
                    <li>每个项目都自己写正则提取 CVE，规则不一致</li>
                    <li>手动处理格式不一致 — <code>cve-2022-12345</code> vs <code>CVE-2022-12345</code></li>
                    <li>无法原生比较、排序 CVE 编号</li>
                    <li>多来源合并列表产生大量重复</li>
                    <li>范围表达式 <code>CVE-2022-1000 to 1050</code> 需手动展开</li>
                    <li>反复重新实现验证逻辑，规则各不相同</li>
                  </ul>
                </Card>
              </Col>
              <Col xs={24} md={12}>
                <Card className="solution-card" hoverable>
                  <Title level={4}>🚀 有了它之后</Title>
                  <ul className="problem-list">
                    <li><CheckCircleOutlined style={{color:'#00b894'}} /> 一个 import 搞定所有 CVE 处理</li>
                    <li><CheckCircleOutlined style={{color:'#00b894'}} /> 自动标准化格式，零样板代码</li>
                    <li><CheckCircleOutlined style={{color:'#00b894'}} /> 原生比较、排序、过滤</li>
                    <li><CheckCircleOutlined style={{color:'#00b894'}} /> 一键去重、交集、并集、差集</li>
                    <li><CheckCircleOutlined style={{color:'#00b894'}} /> 自动展开 CVE 范围表达式</li>
                    <li><CheckCircleOutlined style={{color:'#00b894'}} /> 统一验证规则，经过充分测试</li>
                  </ul>
                </Card>
              </Col>
            </Row>
          </div>
        </section>

        {/* ====== Feature Section ====== */}
        <section className="feature-section" id="features">
          <div className="section-inner">
            <Title level={2} className="section-title">
              <ApiOutlined /> 功能特性
            </Title>
            <Paragraph className="section-subtitle">
              覆盖 CVE 处理全流程的 30+ 工具函数，9 大功能模块
            </Paragraph>
            <Row gutter={[24, 24]}>
              {features.map((f, i) => (
                <Col xs={24} sm={12} lg={8} key={i}>
                  <Card className="feature-card" hoverable>
                    <div className="feature-icon" style={{ color: f.color }}>
                      {f.icon}
                    </div>
                    <Title level={4}>{f.title}</Title>
                    <Paragraph className="feature-desc">{f.desc}</Paragraph>
                    <div className="feature-tags">
                      {f.tags.map((t, j) => (
                        <Tag key={j} color={f.color} className="feature-tag">{t}</Tag>
                      ))}
                    </div>
                  </Card>
                </Col>
              ))}
            </Row>
          </div>
        </section>

        {/* ====== Architecture Section ====== */}
        <section className="architecture-section" id="architecture">
          <div className="section-inner">
            <Title level={2} className="section-title">
              架构概览
            </Title>
            <Paragraph className="section-subtitle">
              三层架构设计：CLI 工具 → Go 包 API → 功能模块 → 底层函数
            </Paragraph>
            <div className="image-container">
              <img src="/cve-skills/images/architecture.png" alt="Architecture" className="diagram-img" />
            </div>

            <Title level={3} className="sub-title" style={{ marginTop: 48 }}>
              CLI 命令树
            </Title>
            <Paragraph className="section-subtitle">
              20+ 子命令覆盖所有功能，支持 Shell 自动补全
            </Paragraph>
            <div className="image-container">
              <img src="/cve-skills/images/cli-tree.png" alt="CLI Command Tree" className="diagram-img" />
            </div>

            <Title level={3} className="sub-title" style={{ marginTop: 48 }}>
              功能思维导图
            </Title>
            <div className="image-container">
              <img src="/cve-skills/images/feature-map.png" alt="Feature Map" className="diagram-img" />
            </div>
          </div>
        </section>

        {/* ====== CLI Section ====== */}
        <section className="cli-section" id="cli">
          <div className="section-inner">
            <Title level={2} className="section-title">
              <CodeOutlined /> CLI 命令行工具
            </Title>
            <Paragraph className="section-subtitle">
              安装即用，无需编写代码，在终端中快速处理 CVE 标识符
            </Paragraph>
            <Card className="cli-card">
              <pre className="install-block">{`# 安装 CLI 工具
go install github.com/scagogogo/cve-skills/cmd/cve@latest

# 查看帮助
cve --help`}</pre>
              <Tabs items={cliExamples} className="cli-tabs" />
            </Card>
          </div>
        </section>

        {/* ====== Quick Start Section ====== */}
        <section className="quickstart-section" id="quickstart">
          <div className="section-inner">
            <Title level={2} className="section-title">
              <RocketOutlined /> 快速开始
            </Title>
            <Row gutter={[32, 32]}>
              <Col xs={24} md={12}>
                <Card className="install-card">
                  <Title level={4}>📦 安装</Title>
                  <pre className="install-block">{`# 作为 Go 库
go get github.com/scagogogo/cve-skills

# 作为 CLI 工具
go install github.com/scagogogo/cve-skills/cmd/cve@latest`}</pre>
                </Card>
              </Col>
              <Col xs={24} md={12}>
                <Card className="code-card">
                  <Title level={4}>💻 代码示例</Title>
                  <pre className="code-block">{goCodeExample}</pre>
                </Card>
              </Col>
            </Row>
          </div>
        </section>

        {/* ====== Links Section ====== */}
        <section className="links-section">
          <div className="section-inner">
            <Row gutter={[24, 24]}>
              <Col xs={24} md={8}>
                <Card className="link-card" hoverable>
                  <ApiOutlined style={{ fontSize: 36, color: '#e94560' }} />
                  <Title level={4}>API 文档</Title>
                  <Paragraph>完整的 Go 包 API 参考文档</Paragraph>
                  <Button type="link" href="https://pkg.go.dev/github.com/scagogogo/cve-skills" target="_blank">
                    pkg.go.dev →
                  </Button>
                </Card>
              </Col>
              <Col xs={24} md={8}>
                <Card className="link-card" hoverable>
                  <CodeOutlined style={{ fontSize: 36, color: '#4ecdc4' }} />
                  <Title level={4}>使用指南</Title>
                  <Paragraph>快速开始、安装配置和最佳实践</Paragraph>
                  <Button type="link" href="/docs/">
                    查看文档 →
                  </Button>
                </Card>
              </Col>
              <Col xs={24} md={8}>
                <Card className="link-card" hoverable>
                  <GithubOutlined style={{ fontSize: 36, color: '#2d3436' }} />
                  <Title level={4}>GitHub 仓库</Title>
                  <Paragraph>源代码、Issue 和贡献指南</Paragraph>
                  <Button type="link" href="https://github.com/scagogogo/cve-skills" target="_blank">
                    访问仓库 →
                  </Button>
                </Card>
              </Col>
            </Row>
          </div>
        </section>
      </Content>

      {/* ====== Footer ====== */}
      <Footer className="site-footer">
        <Divider style={{ margin: '0 0 24px 0' }} />
        <div className="footer-content">
          <Paragraph type="secondary">
            CVE Skills — 一个全面的 Go 语言库和 CLI 工具，用于处理 CVE 标识符
          </Paragraph>
          <Space split={<Divider type="vertical" />}>
            <a href="https://github.com/scagogogo/cve-skills" target="_blank" className="footer-link">GitHub</a>
            <a href="https://pkg.go.dev/github.com/scagogogo/cve-skills" target="_blank" className="footer-link">API 文档</a>
            <a href="/docs/" className="footer-link">使用文档</a>
            <a href="https://github.com/scagogogo/cve-skills/issues" target="_blank" className="footer-link">Issue</a>
          </Space>
          <Paragraph type="secondary" style={{ marginTop: 12 }}>
            MIT License © {new Date().getFullYear()} scagogogo
          </Paragraph>
        </div>
      </Footer>
    </Layout>
  )
}

export default HomePage
