# Website Light Color Scheme & Clarity Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: `superpowers:subagent-driven-development`
> Steps use checkbox (`- [ ]`) syntax.

**Goal:** 将 cve-skills 文档站（VitePress）从"偏重且割裂的玫红+靛蓝配色"改造为"统一、偏亮系的天青蓝配色"，并提升首页清晰度，使整体视觉清爽明亮。

**Architecture:** 当前配色割裂——VitePress 默认品牌色为靛蓝 `--vp-c-brand-1: var(--vp-c-indigo-1)` (#3451b2，偏深冷)，而 favicon/hero.svg/mermaid/theme-color 用玫红 #e94560（偏重暖），两套色系打架。改造路径：建立 `.vitepress/theme/` 自定义主题入口，用 CSS 变量覆盖 `--vp-c-brand-*` 链为天青蓝（亮系，明度高、安全领域常用、与白底对比清爽），同时把 favicon.svg/hero.svg 的玫红渐变替换为天青渐变，mermaid primaryColor 与 theme-color meta 同步换色，实现"一套色系贯穿 brand 按钮/链接/图标/图表"。首页清晰度通过精简 hero actions（5 个→3 个主次分明）+ 调整 hero 副标题实现。

**Tech Stack:** VitePress 1.6.4, vitepress-plugin-mermaid 2.0.17, CSS custom properties, SVG

**Risks:**
- Task 1 的 CSS 变量覆盖需同时覆盖亮色 `:root` 与暗色 `.dark`，否则暗色模式品牌色过暗或错乱 → 缓解：style.css 两套均显式声明，亮色用 #0ea5e9 系、暗色用提亮的 #38bdf8 系
- Task 2 的 hero.svg 重绘必须保持"盾牌路径 + CVE 字 + 校验勾 + YYYY-NNNNN 掩码"语义不变，仅换渐变 stop-color → 缓解：只替换两个 stop 的 stop-color 值，不触碰任何 path d 属性
- mermaid primaryColor 改色后，节点文字/边框由插件自动派生，可能对比度不足 → 缓解：Task 4 构建后人工核查首页 mermaid 图表可读性
- VitePress 1.6.4 的 theme/index.ts 必须默认导出 `Theme` 并 `extend: DefaultTheme`，否则主题失效回退默认 → 缓解：按官方签名编写

---

### Task 1: 创建自定义主题入口与品牌色 CSS 覆盖

**Depends on:** None
**Files:**
- Create: `website/.vitepress/theme/index.ts`
- Create: `website/.vitepress/theme/style.css`

- [ ] **Step 1: 创建 theme/index.ts — 注册自定义主题并引入样式覆盖**

VitePress 通过 `.vitepress/theme/index.ts` 的默认导出识别自定义主题。该文件导出 `Theme` 接口、`extend: DefaultTheme`（继承默认主题全部功能），并 `import './style.css'` 注入品牌色覆盖。不创建此文件则 CSS 变量覆盖不会生效。

```typescript
// website/.vitepress/theme/index.ts
import type { Theme } from 'vitepress'
import DefaultTheme from 'vitepress/theme'
import './style.css'

export default {
  extends: DefaultTheme,
} satisfies Theme
```

- [ ] **Step 2: 创建 theme/style.css — 覆盖品牌色为天青蓝亮系（亮色 + 暗色双模式）**

VitePress 默认品牌色链为 `--vp-c-brand-1/2/3/soft`，默认指向 indigo。此文件将其重链为天青蓝（cyan/sky 系）。亮色用沉稳但明亮的 `#0ea5e9`（sky-500）主色，暗色用提亮的 `#38bdf8`（sky-400）保证暗底下不沉闷。同时覆盖 `--vp-c-brand` 旧别名（部分内部组件仍引用）。`--vp-c-brand-soft` 为浅底按钮/标签背景，亮色用淡天青、暗色用半透明天青。

```css
/* website/.vitepress/theme/style.css */

/**
 * 品牌色：天青蓝（亮系）。
 * 覆盖 VitePress 默认的 indigo，统一全站按钮/链接/强调色。
 * 亮色用 sky-500（#0ea5e9）为主，暗色用 sky-400（#38bdf8）提亮。
 */
:root {
  --vp-c-brand-1: #0ea5e9;
  --vp-c-brand-2: #0284c7;
  --vp-c-brand-3: #38bdf8;
  --vp-c-brand-soft: #e0f2fe;
  --vp-c-brand: var(--vp-c-brand-1);
}

.dark {
  --vp-c-brand-1: #38bdf8;
  --vp-c-brand-2: #0ea5e9;
  --vp-c-brand-3: #7dd3fc;
  --vp-c-brand-soft: rgba(56, 189, 248, 0.16);
  --vp-c-brand: var(--vp-c-brand-1);
}

/* 首页 hero 主按钮（brand）与次按钮（alt）的视觉强化 */
:root {
  --vp-button-brand-bg: #0ea5e9;
  --vp-button-brand-hover-bg: #0284c7;
  --vp-button-alt-bg: #f1f5f9;
  --vp-button-alt-hover-bg: #e2e8f0;
  --vp-button-alt-text: #0f172a;
}

.dark {
  --vp-button-brand-bg: #0ea5e9;
  --vp-button-brand-hover-bg: #38bdf8;
  --vp-button-alt-bg: rgba(255, 255, 255, 0.08);
  --vp-button-alt-hover-bg: rgba(255, 255, 255, 0.14);
  --vp-button-alt-text: #e2e8f0;
}
```

- [ ] **Step 3: 验证主题色覆盖生效**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills/website && npm run build 2>&1 | tail -8`
Expected:
  - Exit code: 0
  - Output contains: "building client + server bundles" and "build complete"
  - Output does NOT contain: "error" or "Error"

- [ ] **Step 4: 提交**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills && git add website/.vitepress/theme/index.ts website/.vitepress/theme/style.css && git commit -m "feat(website): add custom theme with light cyan brand color"`

---

### Task 2: 替换 favicon 与 hero 的玫红配色为天青渐变

**Depends on:** Task 1
**Files:**
- Modify: `website/public/favicon.svg:4`（rect fill 玫红→天青）
- Modify: `website/public/hero.svg:7-8`（渐变 stop-color 玫红→天青）

- [ ] **Step 1: 修改 favicon.svg 的背景方块色 — 玫红 #e94560 → 天青 #0ea5e9**
文件: `website/public/favicon.svg`（替换第 4 行 rect 的 fill 属性）

```svg
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 64 64" width="64" height="64">
  <rect width="64" height="64" rx="12" fill="#0ea5e9"/>
  <text x="32" y="42" font-family="-apple-system,Segoe UI,Roboto,sans-serif" font-size="30" font-weight="700" fill="#fff" text-anchor="middle">C</text>
</svg>
```

- [ ] **Step 2: 修改 hero.svg 的渐变色 — 玫红渐变 → 天青渐变**
文件: `website/public/hero.svg`（仅替换 `<linearGradient>` 内两个 `<stop>` 的 stop-color，盾牌 path/CVE 字/校验勾全部保持原样）

```svg
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 240 240" width="240" height="240" role="img" aria-label="CVE Utils logo">
  <defs>
    <linearGradient id="g" x1="0" y1="0" x2="1" y2="1">
      <stop offset="0" stop-color="#38bdf8"/>
      <stop offset="1" stop-color="#0369a1"/>
    </linearGradient>
  </defs>
  <!-- 盾牌：安全 / 漏洞领域意象 -->
  <path d="M120 18 L206 52 V120 C206 172 168 208 120 224 C72 208 34 172 34 120 V52 Z"
        fill="url(#g)"/>
  <path d="M120 40 L186 66 V120 C186 160 158 189 120 202 C82 189 54 160 54 120 V66 Z"
        fill="#ffffff" opacity="0.12"/>
  <!-- CVE 字形 -->
  <text x="120" y="118" font-family="ui-monospace,SFMono-Regular,Menlo,Consolas,monospace"
        font-size="52" font-weight="700" fill="#ffffff" text-anchor="middle">CVE</text>
  <!-- 标识符掩码：CVE-YYYY-NNNNN 的确定性格式 -->
  <text x="120" y="150" font-family="ui-monospace,SFMono-Regular,Menlo,Consolas,monospace"
        font-size="16" font-weight="500" fill="#ffffff" opacity="0.85" text-anchor="middle" letter-spacing="1">YYYY-NNNNN</text>
  <!-- 校验勾：validate 语义 -->
  <path d="M97 170 l14 14 l30 -32" fill="none" stroke="#ffffff" stroke-width="7"
        stroke-linecap="round" stroke-linejoin="round"/>
</svg>
```

渐变色选择：起点 `#38bdf8`（sky-400，亮）到终点 `#0369a1`（sky-700，深），形成有层次的天青渐变，亮而不失立体感，比原玫红渐变更清爽。

- [ ] **Step 3: 同步更新 config.js 的 theme-color meta — 玫红 → 天青**
文件: `website/.vitepress/config.js:85`（theme-color meta content 值）

```javascript
      ['meta', { name: 'theme-color', content: '#0ea5e9' }],
```

- [ ] **Step 4: 同步更新 config.js 的 mermaid primaryColor — 玫红 → 天青**
文件: `website/.vitepress/config.js:115`（mermaid themeVariables primaryColor 值）

```javascript
      themeVariables: { primaryColor: '#0ea5e9' },
```

- [ ] **Step 5: 验证 SVG 与 mermaid 配色统一为天青**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills/website && grep -h 'fill\|stop-color\|primaryColor\|theme-color' public/favicon.svg public/hero.svg .vitepress/config.js | grep -iE '#0ea5e9|#38bdf8|#0369a1'`
Expected:
  - Exit code: 0
  - Output contains: `#0ea5e9` at least 3 times（favicon + theme-color + mermaid）
  - Output does NOT contain: `#e94560` or `#ff5c7a` or `#c9304b`

- [ ] **Step 6: 验证构建**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills/website && npm run build 2>&1 | tail -5`
Expected:
  - Exit code: 0
  - Output contains: "build complete"

- [ ] **Step 7: 提交**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills && git add website/public/favicon.svg website/public/hero.svg website/.vitepress/config.js && git commit -m "feat(website): replace rose palette with cyan in favicon/hero/mermaid"`

---

### Task 3: 提升首页清晰度 — 精简 hero actions 与副标题

**Depends on:** Task 2
**Files:**
- Modify: `website/index.md:1-26`（EN 首页 frontmatter hero 区）
- Modify: `website/zh/index.md:1-26`（ZH 首页 frontmatter hero 区，对齐）

- [ ] **Step 1: 精简 EN 首页 hero actions — 5 个按钮 → 3 个主次分明**

当前首页 hero 有 5 个 action（Quick Start + Download + API + CLI + GitHub），按钮过多导致视觉拥挤、不清晰。保留 1 个 brand 主按钮（Quick Start）+ 2 个 alt 次按钮（API Reference、GitHub），Download 与 CLI 降级到 features 区或导航栏已有。同时 tagline 微调更聚焦。

文件: `website/index.md:1-26`（替换整个 frontmatter 的 hero 块）

```yaml
---
layout: home

hero:
  name: "CVE Utils"
  text: "AI First CVE Toolkit"
  tagline: 30+ Go functions + cross-platform CLI for CVE identifier processing — built to be read, installed, and driven by AI agents.
  image:
    src: /hero.svg
    alt: CVE Utils
  actions:
    - theme: brand
      text: Quick Start
      link: /guide/getting-started
    - theme: alt
      text: API Reference
      link: /api/
    - theme: alt
      text: GitHub
      link: https://github.com/scagogogo/cve-skills

features:
```

注意：`features:` 行之后的内容完全不变，本次只替换 hero 块（第 1 行到第 26 行的 `features:` 之前）。原 tagline 仅微调（"designed to be read" → "built to be read"，更主动；删去冗余的 "— designed to be read, installed, and driven by AI agents." 改为更紧凑的 "— built to be read, installed, and driven by AI agents." 保持语义）。实际改动核心是 actions 从 5 减到 3。

- [ ] **Step 2: 精简 ZH 首页 hero actions — 与 EN 对齐**
文件: `website/zh/index.md:1-26`（替换整个 frontmatter 的 hero 块）

```yaml
---
layout: home

hero:
  name: "CVE Utils"
  text: "AI First 的 CVE 工具集"
  tagline: 30+ Go 函数 + 跨平台 CLI，专为 AI Agent 读取、安装与驱动而设计。
  image:
    src: /hero.svg
    alt: CVE Utils
  actions:
    - theme: brand
      text: 快速开始
      link: /zh/guide/getting-started
    - theme: alt
      text: API 文档
      link: /zh/api/
    - theme: alt
      text: GitHub
      link: https://github.com/scagogogo/cve-skills

features:
```

原 ZH 有 5 个按钮（快速开始 + 下载 CLI + API 文档 + CLI 参考 + GitHub），精简为 3 个（快速开始 + API 文档 + GitHub），与 EN 对齐。Download 与 CLI 参考在导航栏 nav 中已有入口，不在 hero 重复。

- [ ] **Step 3: 验证首页构建**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills/website && npm run build 2>&1 | tail -5`
Expected:
  - Exit code: 0
  - Output contains: "build complete"
  - Output does NOT contain: "error" or "dead link"

- [ ] **Step 4: 提交**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills && git add website/index.md website/zh/index.md && git commit -m "feat(website): streamline homepage hero actions for clarity"`

---

### Task 4: 整体验证 — 构建、死链、配色统一性自查

**Depends on:** Task 3
**Files:**
- Verify only (no file changes unless issues found)

- [ ] **Step 1: 验证全站构建无死链**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills/website && npm run build 2>&1 | tail -10`
Expected:
  - Exit code: 0
  - Output contains: "build complete"
  - Output does NOT contain: "dead link" or "error"

- [ ] **Step 2: 自查配色统一性 — 全站无遗留玫红色**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills/website && grep -rn 'e94560\|ff5c7a\|c9304b' .vitepress/ public/ index.md zh/index.md 2>/dev/null | grep -v node_modules | grep -v dist || echo "无遗留玫红色"`
Expected:
  - Exit code: 0
  - Output contains: "无遗留玫红色"
  - 不得出现任何 `#e94560` / `#ff5c7a` / `#c9304b`

- [ ] **Step 3: 自查品牌色覆盖已写入编译产物**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills/website && grep -rohE '\-\-vp-c-brand-1:\s*#[0-9a-f]{6}' .vitepress/dist/assets/*.css 2>/dev/null | sort -u`
Expected:
  - Exit code: 0
  - Output contains: `#0ea5e9`（亮色）或 `#38bdf8`（暗色）—— 证明 style.css 覆盖已编译进产物
  - Output does NOT contain: `var(--vp-c-indigo`（证明 indigo 链已被覆盖）

- [ ] **Step 4: 启动预览服务确认实际渲染（可选，人工核查）**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills/website && npm run preview &`
Expected:
  - Exit code: 0
  - 输出含本地预览地址（如 `http://localhost:4173/`）
  - 人工核查：hero 按钮为天青蓝、盾牌图标渐变为天青、mermaid 图表节点为天青、链接 hover 为天青
  - 核查完毕后 `kill %1` 关闭预览

- [ ] **Step 5: 提交（如有预览发现的修复）**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills && git add -A && git commit -m "fix(website): color unification fixes from preview review" || echo "无需额外修复"`

---

## 完成标准

- [ ] 全站配色统一为天青蓝亮系（brand 按钮/链接 + favicon + hero + mermaid + theme-color）
- [ ] 亮色与暗色模式品牌色均明亮清爽，不沉闷
- [ ] 首页 hero actions 从 5 个精简为 3 个，主次分明
- [ ] EN/ZH 首页对齐
- [ ] 构建无错、无死链
- [ ] 无任何遗留玫红色（#e94560/#ff5c7a/#c9304b）
