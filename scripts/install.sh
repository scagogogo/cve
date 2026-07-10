#!/usr/bin/env bash
# install.sh — 一键安装 cve CLI
#
# 用法:
#   curl -fsSL https://raw.githubusercontent.com/scagogogo/cve-skills/main/scripts/install.sh | bash
#
# 自动识别 OS/架构，从 GitHub Releases 下载对应预编译二进制并安装到 /usr/local/bin（或 ~/.local/bin）。
# AI Agent 也可直接调用本脚本完成无工具链安装。

set -euo pipefail

OWNER="scagogogo"
REPO="cve-skills"
BINARY_NAME="cve"
INSTALL_DIR="${CVE_INSTALL_DIR:-/usr/local/bin}"
FALLBACK_DIR="${HOME}/.local/bin"

# 可选覆盖（便于测试 / 内网镜像）：
#   CVE_INSTALL_VERSION    跳过 GitHub API，直接使用指定 tag（如 v0.1.0）
#   CVE_DOWNLOAD_BASE       替换下载源基址（默认 GitHub Releases，可用 file:// 本地目录）
#   CVE_INSTALL_DIR         覆盖安装目录（默认 /usr/local/bin）
DOWNLOAD_BASE="${CVE_DOWNLOAD_BASE:-https://github.com/${OWNER}/${REPO}/releases/download}"

# 临时目录（全局，供 EXIT trap 清理；trap 在 main 返回后才触发，故不能用 local）
tmp_dir=""
cleanup() { [ -n "${tmp_dir:-}" ] && rm -rf "$tmp_dir"; }
trap cleanup EXIT

# 颜色输出
info() { printf '\033[1;34m==>\033[0m %s\n' "$*"; }
err()  { printf '\033[1;31m错误:\033[0m %s\n' "$*" >&2; }

# 识别操作系统
detect_os() {
  case "$(uname -s)" in
    Linux*)  echo "linux" ;;
    Darwin*) echo "darwin" ;;
    FreeBSD*) echo "freebsd" ;;
    MINGW*|MSYS*|CYGWIN*) echo "windows" ;;
    *) err "不支持的操作系统: $(uname -s)"; exit 1 ;;
  esac
}

# 识别架构
detect_arch() {
  case "$(uname -m)" in
    x86_64|amd64) echo "amd64" ;;
    arm64|aarch64) echo "arm64" ;;
    armv7l|armv6l|arm) echo "arm" ;;
    i386|i686) echo "386" ;;
    *) err "不支持的架构: $(uname -m)"; exit 1 ;;
  esac
}

# 下载并解压
main() {
  local os arch version download_url archive_name ext
  os="$(detect_os)"
  arch="$(detect_arch)"

  # 获取版本 tag：优先用环境变量覆盖，否则查询 GitHub API 最新 release
  if [ -n "${CVE_INSTALL_VERSION:-}" ]; then
    version="$CVE_INSTALL_VERSION"
    info "使用指定版本: ${version}"
  else
    info "正在获取最新版本..."
    version="$(curl -fsSL "https://api.github.com/repos/${OWNER}/${REPO}/releases/latest" \
      | grep '"tag_name"' | sed -E 's/.*"([^"]+)".*/\1/')"
    if [ -z "${version:-}" ]; then
      err "无法获取最新版本号。请检查网络或前往 https://github.com/${OWNER}/${REPO}/releases 手动下载。"
      exit 1
    fi
    info "最新版本: ${version}"
  fi

  # 构造归档名（与 .goreleaser.yaml 的 name_template 对应）
  # 用 case 而非链式 [ ] && ...，避免 set -e 下未命中分支导致脚本提前退出。
  local os_label arch_label
  case "$os" in
    darwin) os_label="macos" ;;
    *)      os_label="$os" ;;
  esac
  case "$arch" in
    amd64) arch_label="x86_64" ;;
    386)   arch_label="i386" ;;
    arm64) arch_label="arm64" ;;
    arm)   arch_label="arm" ;;
    *)     err "不支持的架构: $arch"; exit 1 ;;
  esac

  if [ "$os" = "windows" ]; then
    ext="zip"
  else
    ext="tar.gz"
  fi
  archive_name="${BINARY_NAME}_${version#v}_${os_label}_${arch_label}.${ext}"
  download_url="${DOWNLOAD_BASE}/${version}/${archive_name}"

  tmp_dir="$(mktemp -d)"

  info "下载 ${download_url}"
  curl -fsSL -o "${tmp_dir}/${archive_name}" "$download_url"

  info "解压..."
  if [ "$ext" = "zip" ]; then
    unzip -o "${tmp_dir}/${archive_name}" -d "$tmp_dir" >/dev/null
  else
    tar -xzf "${tmp_dir}/${archive_name}" -C "$tmp_dir"
  fi

  # 选择可写安装目录
  local target_dir="$INSTALL_DIR"
  if ! [ -w "$INSTALL_DIR" ] && ! sudo -n true 2>/dev/null; then
    target_dir="$FALLBACK_DIR"
    mkdir -p "$target_dir"
    info "无 ${INSTALL_DIR} 写权限，安装到 ${target_dir}"
  fi

  if [ "$target_dir" = "$INSTALL_DIR" ]; then
    sudo install -m 0755 "${tmp_dir}/${BINARY_NAME}" "${target_dir}/${BINARY_NAME}"
  else
    install -m 0755 "${tmp_dir}/${BINARY_NAME}" "${target_dir}/${BINARY_NAME}"
  fi

  info "安装完成: ${target_dir}/${BINARY_NAME}"
  info "版本: $(${target_dir}/${BINARY_NAME} version)"

  if [ "$target_dir" = "$FALLBACK_DIR" ]; then
    printf '\n请确保 %s 在你的 PATH 中:\n  export PATH="%s:$PATH"\n' "$FALLBACK_DIR" "$FALLBACK_DIR"
  fi
}

# 仅在直接执行时运行（被 source 时不自动执行，便于测试）
if [ "${BASH_SOURCE[0]:-$0}" = "${0}" ]; then
  main "$@"
fi
