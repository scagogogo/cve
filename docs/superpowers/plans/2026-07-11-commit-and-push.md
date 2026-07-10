# Commit All Local Changes & Push to Origin Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: `superpowers:subagent-driven-development`
> Steps use checkbox (`- [ ]`) syntax.

**Goal:** 将当前工作区 129 项未提交变更（14 修改 + 63 删除 + 52 未跟踪）按主题整理为 3 个提交，连同已领先的 13 个提交一并推送到 origin（git@github.com:scagogogo/cve-skills.git）的 main 分支。

**Architecture:** 当前 origin/main 落后本地 13 个提交（测试覆盖 + 配色改造）。工作区另有大规模未提交变更：website 从旧 VitePress(`docs/`)+React(`website/src/`) 双结构迁移到新 VitePress(`website/` 根)、新增 CI/release/goreleaser 工具链、库 Version 改为 ldflags 注入、计划文档更新。按主题分 3 个提交：(1) 库代码 + 计划文档，(2) CI/release/goreleaser 工具链，(3) website 全站迁移。关键排除项：`plans/` 孤儿目录（其内容 `2026-05-28-cobra-cli.md`/`2026-06-15-doc-cli-example-expansion.md` 与 HEAD 中已删除的 `docs/superpowers/plans/` 同名文件字节相同，是迁移残留，用户已主动删除，不可换路径塞回）。推送用 `git push origin main`，推送后验证 origin/main 与本地 HEAD 一致。

**Tech Stack:** Git 2.x, SSH (github.com)

**Risks:**
- `git add -A` / `git add .` 会把 `plans/` 孤儿目录一并暂存 → 缓解：每步用显式路径 add，或在 add 后用 `git reset plans/` 排除
- 63 个删除 + 50 个新增混在 website 迁移里，提交体积大 → 缓解：website 主题单独一个提交，提交信息说明迁移性质
- 推送可能因远程有新提交而拒绝（non-fast-forward）→ 缓解：推送前 `git fetch` + 检查 origin/main 是否落后，若远程有新提交则先 rebase
- `plans/` 目录若被误提交，等于恢复用户主动删除的旧文件 → 缓解：Step 1 显式验证 `plans/` 未被任何提交包含

---

### Task 1: 提交库代码改动与计划文档

**Depends on:** None
**Files:**
- Modify: `cve.go`（Version const→var，ldflags 注入）
- Modify: `base_test.go`, `extract.go`, `extract_test.go`, `filter_test.go`, `generate.go`, `generate_test.go`（库代码微调）
- Modify: `README.md`, `README.zh.md`（文档同步）
- Add: `docs/superpowers/plans/2026-07-08-cli-100pct-coverage.md`
- Add: `docs/superpowers/plans/2026-07-10-website-light-color-scheme.md`

- [ ] **Step 1: 暂存库代码、README 与计划文档 — 用显式路径避免误加 plans/ 孤儿目录**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills && git add cve.go base_test.go extract.go extract_test.go filter_test.go generate.go generate_test.go README.md README.zh.md docs/superpowers/plans/2026-07-08-cli-100pct-coverage.md docs/superpowers/plans/2026-07-10-website-light-color-scheme.md`
Expected:
  - Exit code: 0
  - 无报错

- [ ] **Step 2: 验证暂存区不含 plans/ 孤儿目录**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills && git diff --cached --name-only | grep '^plans/' && echo "错误：plans/ 被误暂存" || echo "✓ plans/ 未被暂存"`
Expected:
  - Exit code: 0
  - Output contains: "✓ plans/ 未被暂存"
  - 不得出现 "错误：plans/ 被误暂存"

- [ ] **Step 3: 确认暂存内容符合预期（应为 10 个文件）**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills && git diff --cached --name-only | wc -l`
Expected:
  - Exit code: 0
  - Output: `10`（cve.go + 6 个测试/库文件 + 2 个 README + 2 个计划文档）

- [ ] **Step 4: 提交库代码与计划文档**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills && git commit -m "$(cat <<'EOF'
chore(lib): switch Version to ldflags-injected var and sync docs/plans

- cve.go: Version from const "v0.0.1" to var "dev" for goreleaser ldflags injection
- sync minor library/test adjustments and README content
- add planning docs: CLI 100% coverage + website light color scheme

Co-Authored-By: Claude Opus 4.8 <noreply@anthropic.com>
EOF
)"`
Expected:
  - Exit code: 0
  - Output contains: "files changed" / "master" / "main"

---

### Task 2: 提交 CI/release/goreleaser 工具链

**Depends on:** Task 1
**Files:**
- Add: `.github/workflows/ci.yml`（CI 流水线）
- Add: `.github/workflows/release.yml`（发布流水线）
- Add: `.goreleaser.yaml`（goreleaser 配置）
- Add: `scripts/install.sh`（安装脚本）
- Modify: `.github/workflows/website.yml`（站点部署工作流调整）
- Modify: `.gitignore`（新增 dist/ 等忽略项）

- [ ] **Step 1: 暂存工具链文件 — CI、release、goreleaser、install 脚本、gitignore、website 工作流**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills && git add .github/workflows/ci.yml .github/workflows/release.yml .github/workflows/website.yml .goreleaser.yaml scripts/install.sh .gitignore`
Expected:
  - Exit code: 0

- [ ] **Step 2: 验证暂存区不含 plans/ 孤儿目录**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills && git diff --cached --name-only | grep '^plans/' && echo "错误" || echo "✓ plans/ 未暂存"`
Expected:
  - Output contains: "✓ plans/ 未暂存"

- [ ] **Step 3: 确认暂存内容（应为 6 个文件）**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills && git diff --cached --name-only`
Expected:
  - 列出 `.github/workflows/ci.yml`、`.github/workflows/release.yml`、`.github/workflows/website.yml`、`.goreleaser.yaml`、`scripts/install.sh`、`.gitignore` 共 6 项

- [ ] **Step 4: 提交工具链**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills && git commit -m "$(cat <<'EOF'
ci: add goreleaser pipeline, CI workflow, and install script

- .goreleaser.yaml: cross-platform release config with ldflags Version injection
- .github/workflows/ci.yml: test pipeline
- .github/workflows/release.yml: release pipeline on tag
- .github/workflows/website.yml: site deploy adjustments
- scripts/install.sh: standalone installer
- .gitignore: ignore goreleaser dist/ output

Co-Authored-By: Claude Opus 4.8 <noreply@anthropic.com>
EOF
)"`
Expected:
  - Exit code: 0

---

### Task 3: 提交 website 全站迁移（删除旧站 + 新增新站）

**Depends on:** Task 2
**Files:**
- Delete: `docs/`（旧 VitePress 站，~40 文件）
- Delete: `website/src/`、`website/index.html`、`website/tsconfig*.json`、`website/vite.config.ts`、`website/.oxlintrc.json`、`website/README.md`、`website/public/icons.svg`、`website/public/images/*.png`（旧 React 站）
- Delete: `scripts/gen_architecture.py`、`scripts/gen_cli_tree.py`、`scripts/gen_feature_map.py`（旧图片生成脚本）
- Add: `website/api/`、`website/cli/`、`website/examples/`、`website/reference/`、`website/guide/*.md`、`website/cli.md`、`website/download.md`、`website/public/llms.txt` 及 `website/zh/` 对应目录（新 VitePress 站，~252 文件）
- Modify: `website/.gitignore`、`website/package.json`、`website/package-lock.json`

- [ ] **Step 1: 暂存剩余全部变更（website 迁移）— 此时 plans/ 仍未跟踪，git add -A 不会动它**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills && git add -A && git reset plans/ 2>/dev/null; git status --short | grep '^??' | grep '^?? plans/' && echo "plans/ 仍未跟踪（预期）" || echo "✓ 无 plans/ 残留未跟踪"`
Expected:
  - Exit code: 0
  - `git reset plans/` 确保即使被 -A 暂存也被移出
  - Output contains: "✓ 无 plans/ 残留未跟踪" 或 "plans/ 仍未跟踪（预期）"

  说明：`git add -A` 会暂存所有删除+修改+未跟踪。`plans/` 是未跟踪目录，会被 -A 暂存，故紧跟 `git reset plans/` 把它移回未跟踪态。

- [ ] **Step 2: 验证 plans/ 未被暂存（关键安全检查）**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills && (git diff --cached --name-only | grep '^plans/' && echo "❌ 错误：plans/ 被暂存，停止") || echo "✓ plans/ 未暂存"`
Expected:
  - Exit code: 0
  - Output contains: "✓ plans/ 未暂存"
  - 不得出现 "❌ 错误"

- [ ] **Step 3: 验证 plans/ 仍保持未跟踪态（未被删除也未暂存）**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills && git status --short plans/`
Expected:
  - Exit code: 0
  - Output: `?? plans/`（仍为未跟踪，原样保留在工作区）

- [ ] **Step 4: 确认暂存内容均为 website/docs/scripts 迁移相关**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills && git diff --cached --name-status | awk '{print $1}' | sort | uniq -c`
Expected:
  - Exit code: 0
  - 输出含 `A`（新增）、`D`（删除）、`M`（修改）三类计数
  - 总文件数应 ~245 项（52 未跟踪 - 1 plans/ + 63 删除 + 部分 website 修改）

- [ ] **Step 5: 提交 website 全站迁移**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills && git commit -m "$(cat <<'EOF'
feat(website): migrate site to VitePress root structure

- remove legacy VitePress site under docs/ and React app under website/src/
- remove obsolete image-generation scripts (gen_architecture/cli_tree/feature_map)
- consolidate to single VitePress site at website/ root with EN (root) + ZH (zh/)
- add full content: api/, cli/, examples/, reference/, guide/* (EN+ZH parity)
- add llms.txt for AI crawlers
- update website package.json/.gitignore for VitePress toolchain

Co-Authored-By: Claude Opus 4.8 <noreply@anthropic.com>
EOF
)"`
Expected:
  - Exit code: 0

---

### Task 4: 推送全部提交到 origin/main 并验证

**Depends on:** Task 3
**Files:**
- Verify only (no file changes)

- [ ] **Step 1: 推送前确认本地领先 origin 的提交数**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills && git fetch origin && git rev-list --count origin/main..HEAD`
Expected:
  - Exit code: 0
  - 输出一个数字（13 已有 + 3 新提交 = 16，或 fetch 后若远程无新提交则为 16）
  - 若输出为 0，说明远程已与本地一致（异常，需排查）

- [ ] **Step 2: 检查远程是否有本地没有的提交（避免 non-fast-forward）**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills && git rev-list --count HEAD..origin/main`
Expected:
  - Exit code: 0
  - 输出 `0`（远程无本地缺失的提交，可直接 fast-forward push）
  - 若输出 >0，说明远程有新提交 → 需先 `git pull --rebase origin main`，再推送

- [ ] **Step 3: 推送到 origin/main**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills && git push origin main`
Expected:
  - Exit code: 0
  - Output contains: "To github.com:scagogogo/cve-skills.git" and "main -> main"

- [ ] **Step 4: 验证推送成功 — origin/main 与本地 HEAD 一致**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills && git fetch origin && echo "本地 HEAD: $(git rev-parse HEAD)" && echo "远程 main: $(git rev-parse origin/main)" && [ "$(git rev-parse HEAD)" = "$(git rev-parse origin/main)" ] && echo "✓ 推送成功，本地与远程一致" || echo "❌ 不一致，需排查"`
Expected:
  - Exit code: 0
  - 本地 HEAD 与远程 main 的 commit hash 相同
  - Output contains: "✓ 推送成功，本地与远程一致"

- [ ] **Step 5: 确认工作区只剩 plans/ 孤儿目录（未提交，符合预期）**
Run: `cd /home/cc11001100/github/scagogogo/cve-skills && git status --short`
Expected:
  - Exit code: 0
  - 仅剩 `?? plans/`（未跟踪孤儿目录，按设计保留在工作区不提交）
  - 无其他未提交/未推送项

---

## 完成标准

- [ ] 3 个主题提交完成（库+计划 / 工具链 / website 迁移）
- [ ] `plans/` 孤儿目录未被任何提交包含，保留在工作区未跟踪态
- [ ] `git push origin main` 成功
- [ ] origin/main 与本地 HEAD commit hash 一致
- [ ] 工作区仅剩 `?? plans/`
