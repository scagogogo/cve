import { defineConfig } from 'vitepress'
import { withMermaid } from 'vitepress-plugin-mermaid'

// 站点常量：base 与仓库名一致，用于 GitHub Pages 部署到 https://scagogogo.github.io/cve-skills/
const HOSTNAME = 'https://scagogogo.github.io'
const BASE = '/cve-skills/'
const SITE_URL = `${HOSTNAME}${BASE}` // 末尾带斜杠，用于绝对 URL 拼接

// AI First 结构化数据：让抓取器 / AI 无需解析 HTML 即可读到软件元信息
const JSON_LD = {
  '@context': 'https://schema.org',
  '@type': 'SoftwareSourceCode',
  name: 'CVE Utils',
  description:
    'AI First CVE identifier toolkit — 30+ Go functions + cross-platform CLI for format/validate, extract, compare/sort, filter/group, generate, set operations, range parsing and statistics.',
  url: SITE_URL,
  codeRepository: 'https://github.com/scagogogo/cve-skills',
  programmingLanguage: 'Go',
  runtimePlatform: 'Go 1.18+',
  license: 'https://opensource.org/licenses/MIT',
  author: { '@type': 'Organization', name: 'scagogogo' },
  keywords: [
    'CVE',
    'vulnerability',
    'security',
    'Go library',
    'CLI',
    'AI agent',
    'CVE parser',
  ],
}

// 本地搜索 UI 文案（中文）——VitePress local search 默认英文，需显式翻译
const zhSearchTranslations = {
  button: { buttonText: '搜索文档', buttonAriaLabel: '搜索文档' },
  modal: {
    displayDetails: '显示详细列表',
    resetButtonTitle: '清除查询条件',
    backButtonTitle: '关闭搜索',
    noResultsText: '无法找到相关结果',
    footer: {
      selectText: '选择',
      selectKeyAriaLabel: '回车',
      navigateText: '切换',
      navigateUpKeyAriaLabel: '上箭头',
      navigateDownKeyAriaLabel: '下箭头',
      closeText: '关闭',
      closeKeyAriaLabel: 'esc',
    },
  },
}

export default withMermaid(
  defineConfig({
    title: 'CVE Utils',
    description: 'AI First 的 CVE 标识符处理工具集 —— Go 库 + CLI',
    base: BASE,
    lang: 'en',
    lastUpdated: true,
    cleanUrls: true,
    metaChunk: true, // 将元数据抽到独立 chunk，避免每页重复内联，减小体积

    // 生成 sitemap.xml，利于搜索引擎与 AI 抓取器发现全部页面
    sitemap: { hostname: SITE_URL },

    // 逐页注入 canonical 规范链接 + 页面级 og:title/og:url
    // 避免带/不带斜杠、base 前缀导致的重复内容；让每页被分享/抓取时携带精确元信息
    transformHead({ pageData }) {
      const rel = pageData.relativePath || ''
      // cleanUrls 下：index.md → 目录，其余去掉 .md 扩展名
      const path = rel.replace(/(^|\/)index\.md$/, '$1').replace(/\.md$/, '')
      const url = `${SITE_URL}${path}`
      const pageTitle = pageData.title
        ? `${pageData.title} | CVE Utils`
        : 'CVE Utils — AI First CVE Toolkit'
      return [
        ['link', { rel: 'canonical', href: url }],
        ['meta', { property: 'og:url', content: url }],
        ['meta', { property: 'og:title', content: pageTitle }],
      ]
    },

    head: [
      ['link', { rel: 'icon', type: 'image/svg+xml', href: `${BASE}favicon.svg` }],
      ['meta', { name: 'theme-color', content: '#0ea5e9' }],
      // Open Graph（社交分享 / AI 预览卡片）；og:title 与 og:url 由 transformHead 逐页动态注入
      ['meta', { property: 'og:type', content: 'website' }],
      ['meta', { property: 'og:site_name', content: 'CVE Utils' }],
      [
        'meta',
        {
          property: 'og:description',
          content:
            'Go library + cross-platform CLI for CVE identifier processing. 30+ deterministic functions, built to be driven by AI agents.',
        },
      ],
      // Twitter card
      ['meta', { name: 'twitter:card', content: 'summary' }],
      ['meta', { name: 'twitter:title', content: 'CVE Utils — AI First CVE Toolkit' }],
      [
        'meta',
        {
          name: 'twitter:description',
          content:
            'Go library + cross-platform CLI for CVE identifier processing. Built to be driven by AI agents.',
        },
      ],
      // JSON-LD 结构化数据
      ['script', { type: 'application/ld+json' }, JSON.stringify(JSON_LD)],
    ],

    // mermaid 初始化选项；插件会跟随 VitePress 深浅色模式自动切换主题
    mermaid: {
      theme: 'default',
      themeVariables: { primaryColor: '#0ea5e9' },
    },

    locales: {
      root: {
        label: 'English',
        lang: 'en',
        title: 'CVE Utils',
        description: 'AI First CVE identifier toolkit — Go library + CLI',
        themeConfig: {
          nav: [
            { text: 'Guide', link: '/guide/getting-started' },
            { text: 'API', link: '/api/' },
            { text: 'CLI', link: '/cli' },
            { text: 'Examples', link: '/examples/' },
            { text: 'Download', link: '/download' },
            { text: 'GitHub', link: 'https://github.com/scagogogo/cve-skills' },
          ],
          sidebar: {
            '/guide/': [
              {
                text: 'Guide',
                items: [
                  { text: "Getting Started", link: '/guide/getting-started' },
                  { text: "Installation", link: '/guide/installation' },
                  { text: "Basic Usage", link: '/guide/basic-usage' },
                  { text: "CVE Identifier System", link: '/guide/cve-identifier-system' },
                  { text: "Year Validation Rules", link: '/guide/year-rules' },
                  { text: "Regex Matching Internals", link: '/guide/regex-internals' },
                  { text: "Formatting & Normalization", link: '/guide/formatting-normalization' },
                  { text: "Validation Strategy", link: '/guide/validation-strategy' },
                  { text: "Comparison & Ordering", link: '/guide/comparison-ordering' },
                  { text: "Set Operations Guide", link: '/guide/set-operations-guide' },
                  { text: "Range Parsing Guide", link: '/guide/range-parsing-guide' },
                  { text: "Statistics & Analysis", link: '/guide/statistics-analysis' },
                  { text: "Performance Characteristics", link: '/guide/performance' },
                  { text: "AI Agent Integration", link: '/guide/ai-integration' },
                  { text: "Error Handling & Edge Cases", link: '/guide/error-handling' },
                  { text: "Library Design Philosophy", link: '/guide/library-design' },
                  { text: "Testing Strategy", link: '/guide/testing' }
                ],
              },
              {
                text: 'Reference',
                items: [
                  { text: "FAQ", link: '/reference/faq' },
                  { text: "Glossary", link: '/reference/glossary' },
                  { text: "Changelog", link: '/reference/changelog' },
                  { text: "CLI Conventions", link: '/reference/cli-conventions' },
                  { text: "Migration Guide", link: '/reference/migration' },
                  { text: "CI/CD Pipeline", link: '/reference/ci-cd' },
                  { text: "Release & goreleaser", link: '/reference/release' }
                ],
              },
            ],
            '/api/': [
              {
                text: 'API Reference',
                items: [
                  { text: 'Overview', link: '/api/' },
                  { text: 'Format & Validation', link: '/api/format-validate' },
                  { text: 'Format & Validation Functions', collapsed: true, items: [
                  { text: "Format", link: '/api/functions/format' },
                  { text: "FormatSeq", link: '/api/functions/format-seq' },
                  { text: "IsCve", link: '/api/functions/is-cve' },
                  { text: "IsContainsCve", link: '/api/functions/is-contains-cve' },
                  { text: "IsCveYearOk", link: '/api/functions/is-cve-year-ok' },
                  { text: "IsCveYearOkWithCutoff", link: '/api/functions/is-cve-year-ok-with-cutoff' },
                  { text: "Split", link: '/api/functions/split' },
                  { text: "ValidateCve", link: '/api/functions/validate-cve' }
                  ] },
                  { text: 'Batch Validation', link: '/api/batch-validation' },
                  { text: 'Batch Validation Functions', collapsed: true, items: [
                  { text: "ValidateCves", link: '/api/functions/validate-cves' },
                  { text: "FilterValidCves", link: '/api/functions/filter-valid-cves' }
                  ] },
                  { text: 'Extraction Methods', link: '/api/extract' },
                  { text: 'Extraction Functions', collapsed: true, items: [
                  { text: "ExtractCve", link: '/api/functions/extract-cve' },
                  { text: "ExtractFirstCve", link: '/api/functions/extract-first-cve' },
                  { text: "ExtractLastCve", link: '/api/functions/extract-last-cve' },
                  { text: "ExtractCveYear", link: '/api/functions/extract-cve-year' },
                  { text: "ExtractCveYearAsInt", link: '/api/functions/extract-cve-year-as-int' },
                  { text: "ExtractCveSeq", link: '/api/functions/extract-cve-seq' },
                  { text: "ExtractCveSeqAsInt", link: '/api/functions/extract-cve-seq-as-int' }
                  ] },
                  { text: 'Comparison & Sorting', link: '/api/compare-sort' },
                  { text: 'Comparison & Sorting Functions', collapsed: true, items: [
                  { text: "CompareByYear", link: '/api/functions/compare-by-year' },
                  { text: "SubByYear", link: '/api/functions/sub-by-year' },
                  { text: "CompareCves", link: '/api/functions/compare-cves' },
                  { text: "SortCves", link: '/api/functions/sort-cves' }
                  ] },
                  { text: 'Filtering & Grouping', link: '/api/filter-group' },
                  { text: 'Filtering & Grouping Functions', collapsed: true, items: [
                  { text: "GroupByYear", link: '/api/functions/group-by-year' },
                  { text: "FilterCvesByYear", link: '/api/functions/filter-cves-by-year' },
                  { text: "FilterCvesByYearRange", link: '/api/functions/filter-cves-by-year-range' },
                  { text: "GetRecentCves", link: '/api/functions/get-recent-cves' }
                  ] },
                  { text: 'Set Operations', link: '/api/set-operations' },
                  { text: 'Set Operation Functions', collapsed: true, items: [
                  { text: "IntersectCves", link: '/api/functions/intersect-cves' },
                  { text: "UnionCves", link: '/api/functions/union-cves' },
                  { text: "DiffCves", link: '/api/functions/diff-cves' },
                  { text: "RemoveDuplicateCves", link: '/api/functions/remove-duplicate-cves' }
                  ] },
                  { text: 'Range & Pattern', link: '/api/range-pattern' },
                  { text: 'Range & Pattern Functions', collapsed: true, items: [
                  { text: "ParseCveRange", link: '/api/functions/parse-cve-range' },
                  { text: "IsCvesConsecutive", link: '/api/functions/is-cves-consecutive' },
                  { text: "FilterCvesByPattern", link: '/api/functions/filter-cves-by-pattern' }
                  ] },
                  { text: 'Statistical Analysis', link: '/api/statistics' },
                  { text: 'Statistics Functions', collapsed: true, items: [
                  { text: "CountByYear", link: '/api/functions/count-by-year' },
                  { text: "YearRange", link: '/api/functions/year-range' },
                  { text: "SeqRange", link: '/api/functions/seq-range' }
                  ] },
                  { text: 'Generation & Construction', link: '/api/generate' },
                  { text: 'Generation Functions', collapsed: true, items: [
                  { text: "GenerateCve", link: '/api/functions/generate-cve' },
                  { text: "GenerateFakeCve", link: '/api/functions/generate-fake-cve' }
                  ] },
                  { text: 'Package Metadata', collapsed: true, items: [
                  { text: "Version", link: '/api/functions/version' },
                  { text: "CveValidationResult", link: '/api/functions/cve-validation-result' }
                  ] },
                ],
              },
            ],
            '/examples/': [
              {
                text: 'Examples',
                items: [
                  { text: 'Overview', link: '/examples/' },
                  { text: 'Vulnerability Analysis', link: '/examples/vulnerability-analysis' },
                  { text: 'Vulnerability Management', link: '/examples/vulnerability-management' },
                  { text: 'CVE Validation', link: '/examples/cve-validation' },
                  { text: 'Runnable Examples', collapsed: true, items: [
                  { text: "Example: Format CVE", link: '/examples/01-format' },
                  { text: "Example: IsCve", link: '/examples/02-is-cve' },
                  { text: "Example: IsContainsCve", link: '/examples/03-is-contains-cve' },
                  { text: "Example: ExtractCve", link: '/examples/04-extract-cve' },
                  { text: "Example: IsCveYearOk", link: '/examples/04-is-valid-cve-year' },
                  { text: "Example: ExtractFirstCve / ExtractLastCve", link: '/examples/05-extract-first-last-cve' },
                  { text: "Example: Extract Year & Seq", link: '/examples/06-extract-year-seq' },
                  { text: "Example: Validate CVE Year", link: '/examples/07-validate-cve-year' },
                  { text: "Example: ValidateCve", link: '/examples/08-validate-cve' },
                  { text: "Example: CompareByYear", link: '/examples/09-compare-by-year' },
                  { text: "Example: CompareCves", link: '/examples/10-compare-cves' },
                  { text: "Example: SortCves", link: '/examples/11-sort-cves' },
                  { text: "Example: GroupByYear", link: '/examples/12-group-by-year' },
                  { text: "Example: FilterCvesByYear", link: '/examples/13-filter-by-year' },
                  { text: "Example: FilterCvesByYearRange", link: '/examples/14-filter-by-year-range' },
                  { text: "Example: GetRecentCves", link: '/examples/15-get-recent-cves' },
                  { text: "Example: RemoveDuplicateCves", link: '/examples/16-remove-duplicate-cves' },
                  { text: "Example: GenerateCve", link: '/examples/17-generate-cve' },
                  { text: "Example: GenerateFakeCve", link: '/examples/18-generate-fake-cve' },
                  { text: "Example: IsContainsCve (report)", link: '/examples/19-is-contains-cve' },
                  { text: "Example: IntersectCves", link: '/examples/20-intersect-cves' },
                  { text: "Example: UnionCves", link: '/examples/21-union-cves' },
                  { text: "Example: DiffCves", link: '/examples/22-diff-cves' },
                  { text: "Example: ValidateCves", link: '/examples/23-validate-cves' },
                  { text: "Example: FilterValidCves", link: '/examples/24-filter-valid-cves' },
                  { text: "Example: ParseCveRange", link: '/examples/25-parse-cve-range' },
                  { text: "Example: IsCvesConsecutive", link: '/examples/26-is-cves-consecutive' },
                  { text: "Example: CountByYear", link: '/examples/27-count-by-year' },
                  { text: "Example: YearRange", link: '/examples/28-year-range' },
                  { text: "Example: SeqRange", link: '/examples/29-seq-range' },
                  { text: "Example: FilterCvesByPattern", link: '/examples/30-filter-by-pattern' },
                  { text: "Example: FormatSeq", link: '/examples/31-format-seq' }
                  ] },
                ],
              },
            ],
            '/cli': [
              {
                text: 'CLI Reference',
                items: [
                  { text: 'Overview & Conventions', link: '/cli' },
                  { text: 'Commands', collapsed: true, items: [
                  { text: "cve format", link: '/cli/commands/format' },
                  { text: "cve validate", link: '/cli/commands/validate' },
                  { text: "cve validate is-cve", link: '/cli/commands/validate-is-cve' },
                  { text: "cve validate contains-cve", link: '/cli/commands/validate-contains-cve' },
                  { text: "cve validate year-ok", link: '/cli/commands/validate-year-ok' },
                  { text: "cve validate-batch", link: '/cli/commands/validate-batch' },
                  { text: "cve filter-valid", link: '/cli/commands/filter-valid' },
                  { text: "cve extract", link: '/cli/commands/extract' },
                  { text: "cve extract first", link: '/cli/commands/extract-first' },
                  { text: "cve extract last", link: '/cli/commands/extract-last' },
                  { text: "cve extract year", link: '/cli/commands/extract-year' },
                  { text: "cve extract seq", link: '/cli/commands/extract-seq' },
                  { text: "cve extract split", link: '/cli/commands/extract-split' },
                  { text: "cve compare", link: '/cli/commands/compare' },
                  { text: "cve compare sort", link: '/cli/commands/compare-sort' },
                  { text: "cve compare by-year", link: '/cli/commands/compare-by-year' },
                  { text: "cve filter by-year", link: '/cli/commands/filter-by-year' },
                  { text: "cve filter by-year-range", link: '/cli/commands/filter-by-year-range' },
                  { text: "cve filter recent", link: '/cli/commands/filter-recent' },
                  { text: "cve filter group-by-year", link: '/cli/commands/filter-group-by-year' },
                  { text: "cve filter dedup", link: '/cli/commands/filter-dedup' },
                  { text: "cve generate cve", link: '/cli/commands/generate-cve' },
                  { text: "cve generate fake", link: '/cli/commands/generate-fake' },
                  { text: "cve parse-range", link: '/cli/commands/parse-range' },
                  { text: "cve is-consecutive", link: '/cli/commands/is-consecutive' },
                  { text: "cve count-by-year", link: '/cli/commands/count-by-year' },
                  { text: "cve year-range", link: '/cli/commands/year-range' },
                  { text: "cve seq-range", link: '/cli/commands/seq-range' },
                  { text: "cve filter-pattern", link: '/cli/commands/filter-pattern' },
                  { text: "cve format-seq", link: '/cli/commands/format-seq' },
                  { text: "cve intersect", link: '/cli/commands/intersect' },
                  { text: "cve union", link: '/cli/commands/union' },
                  { text: "cve diff", link: '/cli/commands/diff' },
                  { text: "cve version", link: '/cli/commands/version' }
                  ] },
                ],
              },
            ],
          },
          editLink: {
            pattern: 'https://github.com/scagogogo/cve-skills/edit/main/website/:path',
            text: 'Edit this page on GitHub',
          },
          outline: { level: [2, 3], label: 'On this page' },
          lastUpdatedText: 'Last updated',
          footer: {
            message: 'Released under the MIT License.',
            copyright: 'Copyright © 2024-2026 scagogogo',
          },
          docFooter: { prev: 'Previous', next: 'Next' },
        },
      },
      zh: {
        label: '简体中文',
        lang: 'zh-CN',
        title: 'CVE Utils',
        description: 'AI First 的 CVE 标识符处理工具集 —— Go 库 + CLI',
        themeConfig: {
          nav: [
            { text: '指南', link: '/zh/guide/getting-started' },
            { text: 'API', link: '/zh/api/' },
            { text: 'CLI', link: '/zh/cli' },
            { text: '示例', link: '/zh/examples/' },
            { text: '下载', link: '/zh/download' },
            { text: 'GitHub', link: 'https://github.com/scagogogo/cve-skills' },
          ],
          sidebar: {
            '/zh/guide/': [
              {
                text: '指南',
                items: [
                  { text: "快速开始", link: '/zh/guide/getting-started' },
                  { text: "安装", link: '/zh/guide/installation' },
                  { text: "基本使用", link: '/zh/guide/basic-usage' },
                  { text: "CVE 编号体系", link: '/zh/guide/cve-identifier-system' },
                  { text: "年份校验规则", link: '/zh/guide/year-rules' },
                  { text: "正则匹配原理", link: '/zh/guide/regex-internals' },
                  { text: "格式化与标准化", link: '/zh/guide/formatting-normalization' },
                  { text: "验证策略", link: '/zh/guide/validation-strategy' },
                  { text: "比较与排序", link: '/zh/guide/comparison-ordering' },
                  { text: "集合运算指南", link: '/zh/guide/set-operations-guide' },
                  { text: "范围解析指南", link: '/zh/guide/range-parsing-guide' },
                  { text: "统计分析", link: '/zh/guide/statistics-analysis' },
                  { text: "性能特性", link: '/zh/guide/performance' },
                  { text: "AI 代理集成", link: '/zh/guide/ai-integration' },
                  { text: "错误处理与边界", link: '/zh/guide/error-handling' },
                  { text: "库设计哲学", link: '/zh/guide/library-design' },
                  { text: "测试策略", link: '/zh/guide/testing' }
                ],
              },
              {
                text: '参考',
                items: [
                  { text: "常见问题", link: '/zh/reference/faq' },
                  { text: "术语表", link: '/zh/reference/glossary' },
                  { text: "更新日志", link: '/zh/reference/changelog' },
                  { text: "CLI 约定", link: '/zh/reference/cli-conventions' },
                  { text: "迁移指南", link: '/zh/reference/migration' },
                  { text: "CI/CD 流水线", link: '/zh/reference/ci-cd' },
                  { text: "发布与 goreleaser", link: '/zh/reference/release' }
                ],
              },
            ],
            '/zh/api/': [
              {
                text: 'API 参考',
                items: [
                  { text: '概览', link: '/zh/api/' },
                  { text: '格式化与验证', link: '/zh/api/format-validate' },
                  { text: '格式化与验证函数', collapsed: true, items: [
                  { text: "Format 格式化", link: '/zh/api/functions/format' },
                  { text: "FormatSeq 序列号定宽", link: '/zh/api/functions/format-seq' },
                  { text: "IsCve 格式判断", link: '/zh/api/functions/is-cve' },
                  { text: "IsContainsCve 包含判断", link: '/zh/api/functions/is-contains-cve' },
                  { text: "IsCveYearOk 年份校验", link: '/zh/api/functions/is-cve-year-ok' },
                  { text: "IsCveYearOkWithCutoff 带偏移年份校验", link: '/zh/api/functions/is-cve-year-ok-with-cutoff' },
                  { text: "Split 拆分", link: '/zh/api/functions/split' },
                  { text: "ValidateCve 单个验证", link: '/zh/api/functions/validate-cve' }
                  ] },
                  { text: '批量验证', link: '/zh/api/batch-validation' },
                  { text: '批量验证函数', collapsed: true, items: [
                  { text: "ValidateCves 批量验证", link: '/zh/api/functions/validate-cves' },
                  { text: "FilterValidCves 过滤有效", link: '/zh/api/functions/filter-valid-cves' }
                  ] },
                  { text: '提取方法', link: '/zh/api/extract' },
                  { text: '提取函数', collapsed: true, items: [
                  { text: "ExtractCve 提取全部", link: '/zh/api/functions/extract-cve' },
                  { text: "ExtractFirstCve 提取首个", link: '/zh/api/functions/extract-first-cve' },
                  { text: "ExtractLastCve 提取末个", link: '/zh/api/functions/extract-last-cve' },
                  { text: "ExtractCveYear 提取年份", link: '/zh/api/functions/extract-cve-year' },
                  { text: "ExtractCveYearAsInt 提取年份整数", link: '/zh/api/functions/extract-cve-year-as-int' },
                  { text: "ExtractCveSeq 提取序列号", link: '/zh/api/functions/extract-cve-seq' },
                  { text: "ExtractCveSeqAsInt 提取序列号整数", link: '/zh/api/functions/extract-cve-seq-as-int' }
                  ] },
                  { text: '比较与排序', link: '/zh/api/compare-sort' },
                  { text: '比较与排序函数', collapsed: true, items: [
                  { text: "CompareByYear 按年份比较", link: '/zh/api/functions/compare-by-year' },
                  { text: "SubByYear 年份相减", link: '/zh/api/functions/sub-by-year' },
                  { text: "CompareCves 完整比较", link: '/zh/api/functions/compare-cves' },
                  { text: "SortCves 排序", link: '/zh/api/functions/sort-cves' }
                  ] },
                  { text: '过滤与分组', link: '/zh/api/filter-group' },
                  { text: '过滤与分组函数', collapsed: true, items: [
                  { text: "GroupByYear 按年分组", link: '/zh/api/functions/group-by-year' },
                  { text: "FilterCvesByYear 按年筛选", link: '/zh/api/functions/filter-cves-by-year' },
                  { text: "FilterCvesByYearRange 年份范围筛选", link: '/zh/api/functions/filter-cves-by-year-range' },
                  { text: "GetRecentCves 最近N年", link: '/zh/api/functions/get-recent-cves' }
                  ] },
                  { text: '集合运算', link: '/zh/api/set-operations' },
                  { text: '集合运算函数', collapsed: true, items: [
                  { text: "IntersectCves 交集", link: '/zh/api/functions/intersect-cves' },
                  { text: "UnionCves 并集", link: '/zh/api/functions/union-cves' },
                  { text: "DiffCves 差集", link: '/zh/api/functions/diff-cves' },
                  { text: "RemoveDuplicateCves 去重", link: '/zh/api/functions/remove-duplicate-cves' }
                  ] },
                  { text: '范围与模式', link: '/zh/api/range-pattern' },
                  { text: '范围与模式函数', collapsed: true, items: [
                  { text: "ParseCveRange 范围解析", link: '/zh/api/functions/parse-cve-range' },
                  { text: "IsCvesConsecutive 连续判断", link: '/zh/api/functions/is-cves-consecutive' },
                  { text: "FilterCvesByPattern 通配符筛选", link: '/zh/api/functions/filter-cves-by-pattern' }
                  ] },
                  { text: '统计分析', link: '/zh/api/statistics' },
                  { text: '统计函数', collapsed: true, items: [
                  { text: "CountByYear 按年计数", link: '/zh/api/functions/count-by-year' },
                  { text: "YearRange 年份范围", link: '/zh/api/functions/year-range' },
                  { text: "SeqRange 序列号范围", link: '/zh/api/functions/seq-range' }
                  ] },
                  { text: '生成与构造', link: '/zh/api/generate' },
                  { text: '生成函数', collapsed: true, items: [
                  { text: "GenerateCve 生成", link: '/zh/api/functions/generate-cve' },
                  { text: "GenerateFakeCve 生成假CVE", link: '/zh/api/functions/generate-fake-cve' }
                  ] },
                  { text: '包元数据', collapsed: true, items: [
                  { text: "Version 版本号", link: '/zh/api/functions/version' },
                  { text: "CveValidationResult 验证结果", link: '/zh/api/functions/cve-validation-result' }
                  ] },
                ],
              },
            ],
            '/zh/examples/': [
              {
                text: '使用示例',
                items: [
                  { text: '概览', link: '/zh/examples/' },
                  { text: '漏洞报告分析', link: '/zh/examples/vulnerability-analysis' },
                  { text: '漏洞库管理', link: '/zh/examples/vulnerability-management' },
                  { text: 'CVE 验证处理', link: '/zh/examples/cve-validation' },
                  { text: '可运行示例', collapsed: true, items: [
                  { text: "示例：格式化 CVE", link: '/zh/examples/01-format' },
                  { text: "示例：判断 CVE 格式", link: '/zh/examples/02-is-cve' },
                  { text: "示例：检测文本是否包含 CVE", link: '/zh/examples/03-is-contains-cve' },
                  { text: "示例：提取全部 CVE", link: '/zh/examples/04-extract-cve' },
                  { text: "示例：校验 CVE 年份", link: '/zh/examples/04-is-valid-cve-year' },
                  { text: "示例：提取首个/末个 CVE", link: '/zh/examples/05-extract-first-last-cve' },
                  { text: "示例：提取年份与序列号", link: '/zh/examples/06-extract-year-seq' },
                  { text: "示例：验证 CVE 年份范围", link: '/zh/examples/07-validate-cve-year' },
                  { text: "示例：全面验证 CVE", link: '/zh/examples/08-validate-cve' },
                  { text: "示例：按年份比较", link: '/zh/examples/09-compare-by-year' },
                  { text: "示例：完整比较 CVE", link: '/zh/examples/10-compare-cves' },
                  { text: "示例：排序 CVE 列表", link: '/zh/examples/11-sort-cves' },
                  { text: "示例：按年份分组", link: '/zh/examples/12-group-by-year' },
                  { text: "示例：按年份筛选", link: '/zh/examples/13-filter-by-year' },
                  { text: "示例：年份范围筛选", link: '/zh/examples/14-filter-by-year-range' },
                  { text: "示例：最近 N 年 CVE", link: '/zh/examples/15-get-recent-cves' },
                  { text: "示例：去重", link: '/zh/examples/16-remove-duplicate-cves' },
                  { text: "示例：生成 CVE", link: '/zh/examples/17-generate-cve' },
                  { text: "示例：生成假 CVE", link: '/zh/examples/18-generate-fake-cve' },
                  { text: "示例：检测报告中的 CVE", link: '/zh/examples/19-is-contains-cve' },
                  { text: "示例：交集", link: '/zh/examples/20-intersect-cves' },
                  { text: "示例：并集", link: '/zh/examples/21-union-cves' },
                  { text: "示例：差集", link: '/zh/examples/22-diff-cves' },
                  { text: "示例：批量验证", link: '/zh/examples/23-validate-cves' },
                  { text: "示例：过滤有效 CVE", link: '/zh/examples/24-filter-valid-cves' },
                  { text: "示例：范围解析", link: '/zh/examples/25-parse-cve-range' },
                  { text: "示例：连续判断", link: '/zh/examples/26-is-cves-consecutive' },
                  { text: "示例：按年计数", link: '/zh/examples/27-count-by-year' },
                  { text: "示例：年份范围", link: '/zh/examples/28-year-range' },
                  { text: "示例：序列号范围", link: '/zh/examples/29-seq-range' },
                  { text: "示例：通配符筛选", link: '/zh/examples/30-filter-by-pattern' },
                  { text: "示例：序列号定宽", link: '/zh/examples/31-format-seq' }
                  ] },
                ],
              },
            ],
            '/zh/cli': [
              {
                text: 'CLI 参考',
                items: [
                  { text: '概览与约定', link: '/zh/cli' },
                  { text: '命令', collapsed: true, items: [
                  { text: "cve format", link: '/zh/cli/commands/format' },
                  { text: "cve validate", link: '/zh/cli/commands/validate' },
                  { text: "cve validate is-cve", link: '/zh/cli/commands/validate-is-cve' },
                  { text: "cve validate contains-cve", link: '/zh/cli/commands/validate-contains-cve' },
                  { text: "cve validate year-ok", link: '/zh/cli/commands/validate-year-ok' },
                  { text: "cve validate-batch", link: '/zh/cli/commands/validate-batch' },
                  { text: "cve filter-valid", link: '/zh/cli/commands/filter-valid' },
                  { text: "cve extract", link: '/zh/cli/commands/extract' },
                  { text: "cve extract first", link: '/zh/cli/commands/extract-first' },
                  { text: "cve extract last", link: '/zh/cli/commands/extract-last' },
                  { text: "cve extract year", link: '/zh/cli/commands/extract-year' },
                  { text: "cve extract seq", link: '/zh/cli/commands/extract-seq' },
                  { text: "cve extract split", link: '/zh/cli/commands/extract-split' },
                  { text: "cve compare", link: '/zh/cli/commands/compare' },
                  { text: "cve compare sort", link: '/zh/cli/commands/compare-sort' },
                  { text: "cve compare by-year", link: '/zh/cli/commands/compare-by-year' },
                  { text: "cve filter by-year", link: '/zh/cli/commands/filter-by-year' },
                  { text: "cve filter by-year-range", link: '/zh/cli/commands/filter-by-year-range' },
                  { text: "cve filter recent", link: '/zh/cli/commands/filter-recent' },
                  { text: "cve filter group-by-year", link: '/zh/cli/commands/filter-group-by-year' },
                  { text: "cve filter dedup", link: '/zh/cli/commands/filter-dedup' },
                  { text: "cve generate cve", link: '/zh/cli/commands/generate-cve' },
                  { text: "cve generate fake", link: '/zh/cli/commands/generate-fake' },
                  { text: "cve parse-range", link: '/zh/cli/commands/parse-range' },
                  { text: "cve is-consecutive", link: '/zh/cli/commands/is-consecutive' },
                  { text: "cve count-by-year", link: '/zh/cli/commands/count-by-year' },
                  { text: "cve year-range", link: '/zh/cli/commands/year-range' },
                  { text: "cve seq-range", link: '/zh/cli/commands/seq-range' },
                  { text: "cve filter-pattern", link: '/zh/cli/commands/filter-pattern' },
                  { text: "cve format-seq", link: '/zh/cli/commands/format-seq' },
                  { text: "cve intersect", link: '/zh/cli/commands/intersect' },
                  { text: "cve union", link: '/zh/cli/commands/union' },
                  { text: "cve diff", link: '/zh/cli/commands/diff' },
                  { text: "cve version", link: '/zh/cli/commands/version' }
                  ] },
                ],
              },
            ],
          },
          editLink: {
            pattern: 'https://github.com/scagogogo/cve-skills/edit/main/website/:path',
            text: '在 GitHub 上编辑此页',
          },
          outline: { level: [2, 3], label: '本页目录' },
          lastUpdatedText: '最后更新于',
          returnToTopLabel: '回到顶部',
          sidebarMenuLabel: '菜单',
          darkModeSwitchLabel: '外观',
          lightModeSwitchTitle: '切换到浅色模式',
          darkModeSwitchTitle: '切换到深色模式',
          langMenuLabel: '切换语言',
          skipToContentLabel: '跳到主要内容',
          footer: {
            message: '基于 MIT 许可证发布。',
            copyright: 'Copyright © 2024-2026 scagogogo',
          },
          docFooter: { prev: '上一页', next: '下一页' },
        },
      },
    },

    themeConfig: {
      socialLinks: [
        { icon: 'github', link: 'https://github.com/scagogogo/cve-skills' },
      ],
      search: {
        provider: 'local',
        options: {
          locales: { zh: zhSearchTranslations },
        },
      },
    },
  }),
)
