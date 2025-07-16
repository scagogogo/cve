# Installation Guide

This page provides detailed instructions on how to install and configure CVE Utils in different environments.

## System Requirements

- Go 1.18 or higher
- Supported operating systems: Linux, macOS, Windows

## Installation Methods

### Method 1: Using go get (Recommended)

This is the simplest and recommended installation method:

```bash
go get github.com/scagogogo/cve
```

### Method 2: Using go mod

如果您使用 Go modules，可以在项目中直接导入：

```go
import "github.com/scagogogo/cve"
```

然后运行：

```bash
go mod tidy
```

### 方法三：手动下载

您也可以手动克隆仓库：

```bash
git clone https://github.com/scagogogo/cve.git
cd cve
go build
```

## 验证安装

### 基本验证

创建一个测试文件验证安装：

```go
// verify.go
package main

import (
    "fmt"
    "github.com/scagogogo/cve"
)

func main() {
    // 测试基本功能
    testCve := "CVE-2022-12345"
    
    // 测试格式化
    formatted := cve.Format(testCve)
    fmt.Printf("格式化测试: %s -> %s\n", testCve, formatted)
    
    // 测试验证
    isValid := cve.ValidateCve(testCve)
    fmt.Printf("验证测试: %s -> %t\n", testCve, isValid)
    
    // 测试提取
    text := "系统受到 CVE-2021-44228 影响"
    extracted := cve.ExtractCve(text)
    fmt.Printf("提取测试: %s -> %v\n", text, extracted)
    
    if len(extracted) > 0 && isValid {
        fmt.Println("✅ CVE Utils 安装成功！")
    } else {
        fmt.Println("❌ 安装验证失败")
    }
}
```

运行验证：

```bash
go run verify.go
```

预期输出：

```
格式化测试: CVE-2022-12345 -> CVE-2022-12345
验证测试: CVE-2022-12345 -> true
提取测试: 系统受到 CVE-2021-44228 影响 -> [CVE-2021-44228]
✅ CVE Utils 安装成功！
```

### 完整功能测试

运行项目自带的测试套件：

```bash
# 克隆仓库（如果还没有）
git clone https://github.com/scagogogo/cve.git
cd cve

# 运行所有测试
go test -v

# 运行测试并显示覆盖率
go test -v -cover
```

## 在项目中使用

### 新项目

如果您正在创建新的 Go 项目：

```bash
# 创建新项目
mkdir my-cve-project
cd my-cve-project

# 初始化 Go module
go mod init my-cve-project

# 添加 CVE Utils 依赖
go get github.com/scagogogo/cve

# 创建主文件
cat > main.go << 'EOF'
package main

import (
    "fmt"
    "github.com/scagogogo/cve"
)

func main() {
    cves := cve.ExtractCve("发现漏洞 CVE-2022-12345")
    fmt.Println("提取的 CVE:", cves)
}
EOF

# 运行项目
go run main.go
```

### 现有项目

如果您要在现有项目中添加 CVE Utils：

```bash
# 在项目根目录下
go get github.com/scagogogo/cve

# 更新依赖
go mod tidy
```

然后在代码中导入：

```go
import "github.com/scagogogo/cve"
```

## 版本管理

### 使用特定版本

如果您需要使用特定版本：

```bash
# 使用特定标签版本
go get github.com/scagogogo/cve@v1.0.0

# 使用特定提交
go get github.com/scagogogo/cve@commit-hash
```

### 查看当前版本

```bash
go list -m github.com/scagogogo/cve
```

### 更新到最新版本

```bash
go get -u github.com/scagogogo/cve
go mod tidy
```

## 常见问题

### 问题 1：Go 版本过低

**错误信息**：
```
go: module github.com/scagogogo/cve requires go >= 1.18
```

**解决方案**：
升级 Go 到 1.18 或更高版本。

### 问题 2：网络连接问题

**错误信息**：
```
go: github.com/scagogogo/cve: dial tcp: lookup github.com: no such host
```

**解决方案**：
1. 检查网络连接
2. 配置 Go 代理（如果在中国）：

```bash
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GOSUMDB=sum.golang.google.cn
```

### 问题 3：模块缓存问题

**错误信息**：
```
go: github.com/scagogogo/cve@v1.0.0: invalid version: unknown revision
```

**解决方案**：
清理模块缓存：

```bash
go clean -modcache
go get github.com/scagogogo/cve
```

### 问题 4：导入路径错误

**错误信息**：
```
package github.com/scagogogo/cve is not in GOROOT
```

**解决方案**：
确保使用 Go modules：

```bash
# 检查是否启用了 Go modules
go env GO111MODULE

# 如果输出不是 "on"，则启用它
go env -w GO111MODULE=on
```

## 开发环境设置

如果您想参与 CVE Utils 的开发：

### 1. 克隆仓库

```bash
git clone https://github.com/scagogogo/cve.git
cd cve
```

### 2. 安装开发依赖

```bash
# 安装测试工具
go install golang.org/x/tools/cmd/cover@latest

# 安装代码格式化工具
go install golang.org/x/tools/cmd/goimports@latest
```

### 3. 运行开发测试

```bash
# 运行所有测试
go test ./...

# 运行测试并生成覆盖率报告
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

# 格式化代码
gofmt -w .
goimports -w .
```

### 4. 构建

```bash
# 构建项目
go build

# 交叉编译（例如为 Linux 构建）
GOOS=linux GOARCH=amd64 go build
```

## 下一步

安装完成后，您可以：

1. 阅读 [快速开始](/guide/getting-started) 学习基本用法
2. 查看 [基本使用指南](/guide/basic-usage) 了解更多功能
3. 浏览 [API 文档](/api/) 了解所有可用函数
4. 查看 [使用示例](/examples/) 学习实际应用场景
