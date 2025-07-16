import { defineConfig } from 'vitepress'

export default defineConfig({
  title: 'CVE Utils',
  description: 'A comprehensive collection of utility functions for handling CVE identifiers',
  base: '/cve/',

  locales: {
    root: {
      label: 'English',
      lang: 'en',
      title: 'CVE Utils',
      description: 'A comprehensive collection of utility functions for handling CVE identifiers'
    },
    zh: {
      label: '简体中文',
      lang: 'zh-CN',
      title: 'CVE Utils',
      description: 'CVE (Common Vulnerabilities and Exposures) 相关的工具方法集合'
    }
  },
  
  themeConfig: {
    // English (root) configuration
    nav: [
      { text: 'Home', link: '/' },
      { text: 'API Docs', link: '/api/' },
      { text: 'Quick Start', link: '/guide/getting-started' },
      { text: 'Examples', link: '/examples/' },
      { text: 'GitHub', link: 'https://github.com/scagogogo/cve' }
    ],

    sidebar: {
      '/guide/': [
        {
          text: 'Guide',
          items: [
            { text: 'Getting Started', link: '/guide/getting-started' },
            { text: 'Installation', link: '/guide/installation' },
            { text: 'Basic Usage', link: '/guide/basic-usage' }
          ]
        }
      ],
      '/api/': [
        {
          text: 'API Reference',
          items: [
            { text: 'Overview', link: '/api/' },
            { text: 'Format & Validation', link: '/api/format-validate' },
            { text: 'Extraction Methods', link: '/api/extract' },
            { text: 'Comparison & Sorting', link: '/api/compare-sort' },
            { text: 'Filtering & Grouping', link: '/api/filter-group' },
            { text: 'Generation & Construction', link: '/api/generate' }
          ]
        }
      ],
      '/examples/': [
        {
          text: 'Examples',
          items: [
            { text: 'Overview', link: '/examples/' },
            { text: 'Vulnerability Analysis', link: '/examples/vulnerability-analysis' },
            { text: 'Vulnerability Management', link: '/examples/vulnerability-management' },
            { text: 'CVE Validation', link: '/examples/cve-validation' }
          ]
        }
      ]
    },

    socialLinks: [
      { icon: 'github', link: 'https://github.com/scagogogo/cve' }
    ],

    footer: {
      message: 'Released under the MIT License.',
      copyright: 'Copyright © 2024 scagogogo'
    },

    search: {
      provider: 'local'
    },

    // Localized theme config
    locales: {
      zh: {
        nav: [
          { text: '首页', link: '/zh/' },
          { text: 'API 文档', link: '/zh/api/' },
          { text: '快速开始', link: '/zh/guide/getting-started' },
          { text: '示例', link: '/zh/examples/' },
          { text: 'GitHub', link: 'https://github.com/scagogogo/cve' }
        ],

        sidebar: {
          '/zh/guide/': [
            {
              text: '指南',
              items: [
                { text: '快速开始', link: '/zh/guide/getting-started' },
                { text: '安装', link: '/zh/guide/installation' },
                { text: '基本使用', link: '/zh/guide/basic-usage' }
              ]
            }
          ],
          '/zh/api/': [
            {
              text: 'API 参考',
              items: [
                { text: '概览', link: '/zh/api/' },
                { text: '格式化与验证', link: '/zh/api/format-validate' },
                { text: '提取方法', link: '/zh/api/extract' },
                { text: '比较与排序', link: '/zh/api/compare-sort' },
                { text: '过滤与分组', link: '/zh/api/filter-group' },
                { text: '生成与构造', link: '/zh/api/generate' }
              ]
            }
          ],
          '/zh/examples/': [
            {
              text: '使用示例',
              items: [
                { text: '概览', link: '/zh/examples/' },
                { text: '漏洞报告分析', link: '/zh/examples/vulnerability-analysis' },
                { text: '漏洞库管理', link: '/zh/examples/vulnerability-management' },
                { text: 'CVE 验证处理', link: '/zh/examples/cve-validation' }
              ]
            }
          ]
        },

        footer: {
          message: '基于 MIT 许可证发布。',
          copyright: 'Copyright © 2024 scagogogo'
        }
      }
    }
  }
})
