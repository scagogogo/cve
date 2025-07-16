import { defineConfig } from 'vitepress'

export default defineConfig({
  title: 'CVE Utils',
  description: 'CVE (Common Vulnerabilities and Exposures) 相关的工具方法集合',
  base: '/cve/',
  
  themeConfig: {
    logo: '/logo.svg',
    
    nav: [
      { text: '首页', link: '/' },
      { text: 'API 文档', link: '/api/' },
      { text: '快速开始', link: '/guide/getting-started' },
      { text: '示例', link: '/examples/' },
      { text: 'GitHub', link: 'https://github.com/scagogogo/cve' }
    ],

    sidebar: {
      '/guide/': [
        {
          text: '指南',
          items: [
            { text: '快速开始', link: '/guide/getting-started' },
            { text: '安装', link: '/guide/installation' },
            { text: '基本使用', link: '/guide/basic-usage' }
          ]
        }
      ],
      '/api/': [
        {
          text: 'API 参考',
          items: [
            { text: '概览', link: '/api/' },
            { text: '格式化与验证', link: '/api/format-validate' },
            { text: '提取方法', link: '/api/extract' },
            { text: '比较与排序', link: '/api/compare-sort' },
            { text: '过滤与分组', link: '/api/filter-group' },
            { text: '生成与构造', link: '/api/generate' }
          ]
        }
      ],
      '/examples/': [
        {
          text: '使用示例',
          items: [
            { text: '概览', link: '/examples/' },
            { text: '漏洞报告分析', link: '/examples/vulnerability-analysis' },
            { text: '漏洞库管理', link: '/examples/vulnerability-management' },
            { text: 'CVE 验证处理', link: '/examples/cve-validation' }
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
    }
  }
})
