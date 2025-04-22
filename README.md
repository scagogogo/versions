# Versions - Go语言版本管理库

[![Go Tests](https://github.com/scagogogo/versions/actions/workflows/go-test.yml/badge.svg)](https://github.com/scagogogo/versions/actions/workflows/go-test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/scagogogo/versions)](https://goreportcard.com/report/github.com/scagogogo/versions)
[![GoDoc](https://godoc.org/github.com/scagogogo/versions?status.svg)](https://godoc.org/github.com/scagogogo/versions)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

`versions` 是一个功能强大的 Go 语言库，专门用于解析、比较、排序和管理软件版本号。支持语义化版本、前缀、后缀和版本分组，让您的版本管理更加简单高效。

## 📋 目录

- [特性](#-特性)
- [安装](#-安装)
- [快速开始](#-快速开始)
- [详细文档](#-详细文档)
  - [数据类型和常量](#数据类型和常量)
  - [主要函数](#主要函数)
- [使用示例](#-使用示例)
- [注意事项](#-注意事项)
- [许可证](#-许可证)

## ✨ 特性

- **全面的版本支持**: 支持标准语义化版本格式（如 `1.2.3`）和多种变体（如 `v1.2.3`、`1.2.3-beta` 等）
- **灵活的版本解析**: 自动识别前缀、版本号和后缀
- **智能版本比较**: 基于标准语义化版本规则进行版本号比较
- **版本分组和排序**: 按主版本号、次版本号分组，并提供多种排序方式
- **范围查询**: 支持查询指定版本范围内的所有版本
- **版本可视化**: 提供文本方式展示版本之间的层次关系
- **文件支持**: 直接从文件中读取和处理版本号
- **无外部依赖**: 核心功能无需额外依赖

## 📦 安装

使用 `go get` 命令安装:

```bash
go get -u github.com/scagogogo/versions
```

## 🚀 快速开始

以下是一个简单的示例，展示如何使用 `versions` 库解析和比较版本号:

```go
package main

import (
    "fmt"
    "github.com/scagogogo/versions"
)

func main() {
    // 创建版本对象
    v1 := versions.NewVersion("1.2.3")
    v2 := versions.NewVersion("v1.3.0")
    
    // 比较版本大小
    if v1.CompareTo(v2) < 0 {
        fmt.Printf("%s 小于 %s\n", v1.Raw, v2.Raw)
    }
    
    // 查看版本组成部分
    fmt.Printf("版本号数字: %v\n", v1.VersionNumbers)
    fmt.Printf("前缀: %s\n", v2.Prefix)  // 输出: "v"
    
    // 排序版本号
    versionList := []*versions.Version{
        versions.NewVersion("2.0.0"),
        versions.NewVersion("1.0.0"),
        versions.NewVersion("1.10.0"),
    }
    sortedVersions := versions.SortVersionSlice(versionList)
    for _, v := range sortedVersions {
        fmt.Println(v.Raw)  // 输出: 1.0.0, 1.10.0, 2.0.0
    }
}
```

## 📚 详细文档

### 数据类型和常量

| 类型 | 描述 |
|------|-----|
| `Version` | 表示一个版本号，包含原始字符串、版本号数字、前缀、后缀和发布时间 |
| `VersionNumbers` | 整数切片，表示版本号中的数字部分 |
| `VersionPrefix` | 字符串，表示版本号数字部分之前的前缀 |
| `VersionSuffix` | 字符串，表示版本号数字部分之后的后缀 |
| `ContainsPolicy` | 用于控制版本查询时的包含策略（包含、不包含） |
| `VersionGroup` | 版本组，包含相同主版本号的一组版本 |
| `SortedVersionGroups` | 有序的版本组集合，便于范围查询 |

### 主要函数

#### 版本解析与创建

```go
// 创建版本对象
version := versions.NewVersion("1.2.3")

// 带错误检查的版本创建
version, err := versions.NewVersionE("1.2.3")
if err != nil {
    log.Fatal(err)
}
```

#### 从文件读取版本

```go
// 读取版本号对象
versions, err := versions.ReadVersionsFromFile("path/to/versions.txt")
if err != nil {
    log.Fatal(err)
}

// 读取版本号字符串
versionStrings, err := versions.ReadVersionsStringFromFile("path/to/versions.txt")
if err != nil {
    log.Fatal(err)
}
```

#### 版本分组与排序

```go
// 版本分组
groupedVersions := versions.Group(versionList)

// 字符串版本排序
sortedStrings := versions.SortVersionStringSlice(versionStrings)

// 版本对象排序
sortedVersions := versions.SortVersionSlice(versionList)
```

#### 版本范围查询

```go
// 创建有序版本组
sortedGroups := versions.NewSortedVersionGroups(versionList)

// 定义查询范围和包含策略
startVersion := versions.NewVersion("1.0.0")
endVersion := versions.NewVersion("2.0.0")
startTuple := tuple.New2[*versions.Version, versions.ContainsPolicy](
    startVersion, versions.ContainsPolicyYes) // 包含起始版本
endTuple := tuple.New2[*versions.Version, versions.ContainsPolicy](
    endVersion, versions.ContainsPolicyNo)   // 不包含结束版本

// 执行范围查询
rangeResult := sortedGroups.QueryRange(startTuple, endTuple)
```

#### 版本可视化

```go
// 可视化所有版本（每组显示最多5个版本）
versions.VisualizeVersions(versionList, os.Stdout, 5)

// 可视化版本组层次结构
versions.VisualizeVersionGroups(versionList, os.Stdout)
```

## 🔍 使用示例

查看 [`examples`](./examples) 目录获取更多详细示例:

- [基本版本解析](./examples/01_basic_version_parsing/main.go) - 如何解析和比较不同格式的版本号
- [从文件读取版本](./examples/02_reading_versions_from_file/main.go) - 如何从文件中读取版本信息
- [版本排序](./examples/03_version_sorting/main.go) - 如何对版本号进行排序
- [版本分组](./examples/04_version_grouping/main.go) - 如何对版本进行分组管理
- [版本范围查询](./examples/05_version_range_queries/main.go) - 如何查询特定版本范围
- [版本可视化](./examples/06_version_visualization/main.go) - 如何可视化版本结构

## ⚠️ 注意事项

- 确保传入的文件路径正确，且文件格式符合要求
- 版本号的解析和比较依赖于正确的格式，非标准格式可能会导致解析错误
- 时间比较依赖于 `PublicTime` 被正确设置
- 对于非标准格式的版本号可能需要额外的处理

## 📄 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](./LICENSE) 文件。