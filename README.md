# Versions - Go语言版本号解析计算SDK

<div align="center">

[![Go Tests](https://github.com/scagogogo/versions/actions/workflows/go-test.yml/badge.svg)](https://github.com/scagogogo/versions/actions/workflows/go-test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/scagogogo/versions)](https://goreportcard.com/report/github.com/scagogogo/versions)
[![GoDoc](https://godoc.org/github.com/scagogogo/versions?status.svg)](https://godoc.org/github.com/scagogogo/versions)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

<img src="https://user-images.githubusercontent.com/5877/236610549-d20056f0-db64-4ba4-aabd-4f3cf78fb8d5.png" alt="Versions Logo" width="180"/>

</div>

<p align="center">
<b>专业的 Go 语言版本号解析计算SDK - 轻松解析、比较、排序软件版本号</b>
</p>

`versions` 是一个专为 Go 开发者设计的版本号解析计算SDK，专注于处理语义化版本号的解析、比较、排序和查询。它不是一个版本管理系统，而是一个帮助开发者处理版本号字符串的工具库。无论是依赖管理、API兼容性检查还是软件更新逻辑，都能高效完成版本号的计算需求。

---

## 📋 目录

- [✨ 特性](#-特性)
- [📦 安装](#-安装)
- [🚀 快速开始](#-快速开始)
- [📚 详细文档](#-详细文档)
  - [数据类型和常量](#数据类型和常量)
  - [主要函数](#主要函数)
- [🔍 使用示例](#-使用示例)
- [⚠️ 注意事项](#-注意事项)
- [📈 性能](#-性能)
- [📄 许可证](#-许可证)

---

## ✨ 特性

<table>
  <tr>
    <td><b>🔄 全面的版本号支持</b></td>
    <td>支持标准语义化版本格式（如 <code>1.2.3</code>）和多种变体（如 <code>v1.2.3</code>、<code>1.2.3-beta</code> 等）</td>
  </tr>
  <tr>
    <td><b>🧩 灵活的版本号解析</b></td>
    <td>自动识别前缀、版本号和后缀，处理各种版本号格式</td>
  </tr>
  <tr>
    <td><b>📊 版本号比较</b></td>
    <td>基于标准语义化版本规则进行版本号比较，支持前缀和后缀处理</td>
  </tr>
  <tr>
    <td><b>📦 版本号分组和排序</b></td>
    <td>按主版本号、次版本号分组，并提供多种排序方式</td>
  </tr>
  <tr>
    <td><b>🔍 版本号范围查询</b></td>
    <td>支持查询指定版本号范围内的所有版本，带有灵活的包含/排除边界选项</td>
  </tr>
  <tr>
    <td><b>📋 版本号可视化</b></td>
    <td>提供文本方式展示版本号之间的层次关系，直观查看版本号组织结构</td>
  </tr>
  <tr>
    <td><b>📁 文件支持</b></td>
    <td>直接从文件中读取和处理版本号</td>
  </tr>
  <tr>
    <td><b>🚀 无外部依赖</b></td>
    <td>核心功能无需额外依赖，轻量快速</td>
  </tr>
</table>

---

## 📦 安装

使用 `go get` 命令安装:

```bash
go get -u github.com/scagogogo/versions
```

<details>
<summary>需要特定版本？</summary>

```bash
go get -u github.com/scagogogo/versions@v1.0.0  # 指定版本
```
</details>

---

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

<details>
<summary>查看输出结果</summary>

```
1.2.3 小于 v1.3.0
版本号数字: [1 2 3]
前缀: v
1.0.0
1.10.0
2.0.0
```
</details>

---

## 📚 详细文档

### 数据类型和常量

<div align="center">

| 类型 | 描述 |
|:------:|:-----|
| <kbd>Version</kbd> | 表示一个版本号，包含原始字符串、版本号数字、前缀、后缀和发布时间 |
| <kbd>VersionNumbers</kbd> | 整数切片，表示版本号中的数字部分 |
| <kbd>VersionPrefix</kbd> | 字符串，表示版本号数字部分之前的前缀 |
| <kbd>VersionSuffix</kbd> | 字符串，表示版本号数字部分之后的后缀 |
| <kbd>ContainsPolicy</kbd> | 用于控制版本查询时的包含策略（包含、不包含） |
| <kbd>VersionGroup</kbd> | 版本组，包含相同主版本号的一组版本 |
| <kbd>SortedVersionGroups</kbd> | 有序的版本组集合，便于范围查询 |

</div>

### 主要函数

<details open>
<summary><b>版本解析与创建</b></summary>

```go
// 创建版本对象
version := versions.NewVersion("1.2.3")

// 带错误检查的版本创建
version, err := versions.NewVersionE("1.2.3")
if err != nil {
    log.Fatal(err)
}
```
</details>

<details open>
<summary><b>从文件读取版本</b></summary>

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
</details>

<details open>
<summary><b>版本分组与排序</b></summary>

```go
// 版本分组
groupedVersions := versions.Group(versionList)

// 字符串版本排序
sortedStrings := versions.SortVersionStringSlice(versionStrings)

// 版本对象排序
sortedVersions := versions.SortVersionSlice(versionList)
```
</details>

<details open>
<summary><b>版本范围查询</b></summary>

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
</details>

<details open>
<summary><b>版本可视化</b></summary>

```go
// 可视化所有版本（每组显示最多5个版本）
versions.VisualizeVersions(versionList, os.Stdout, 5)

// 可视化版本组层次结构
versions.VisualizeVersionGroups(versionList, os.Stdout)
```

**可视化输出示例:**

```
版本总数: 15
版本组数: 3

┌─ 版本组: 1.0 (3个版本)
├── 1.0.0 (发布时间: 2020-01-01)
├── 1.0.1 (发布时间: 2020-02-01)
└── 1.0.2 (发布时间: 2020-03-01)

┌─ 版本组: 2.0 (4个版本)
├── 2.0.0 (发布时间: 2021-01-01)
├── 2.0.1 (发布时间: 2021-02-01)
├── 2.0.2 (发布时间: 2021-03-01)
└── ...还有1个版本未显示
```
</details>

---

## 🔍 使用示例

<div align="center">

| 示例 | 描述 |
|:------:|:-----|
| [📚 基本版本解析](./examples/01_basic_version_parsing/main.go) | 如何解析和比较不同格式的版本号 |
| [📂 从文件读取版本](./examples/02_reading_versions_from_file/main.go) | 如何从文件中读取版本信息 |
| [🔢 版本排序](./examples/03_version_sorting/main.go) | 如何对版本号进行排序 |
| [📦 版本分组](./examples/04_version_grouping/main.go) | 如何对版本进行分组管理 |
| [🔍 版本范围查询](./examples/05_version_range_queries/main.go) | 如何查询特定版本范围 |
| [📊 版本可视化](./examples/06_version_visualization/main.go) | 如何可视化版本结构 |

</div>

<div align="center">
<img src="https://user-images.githubusercontent.com/5877/236610730-b2f3c58f-b70b-4621-9f1a-ae99928dae99.png" alt="版本树示例" width="600"/>
<br>
<i>版本树可视化示例</i>
</div>

---

## ⚠️ 注意事项

- ⚠️ 确保传入的文件路径正确，且文件格式符合要求
- ⚠️ 版本号的解析和比较依赖于正确的格式，非标准格式可能会导致解析错误
- ⚠️ 时间比较依赖于 `PublicTime` 被正确设置
- ⚠️ 对于非标准格式的版本号可能需要额外的处理

---

## 📈 性能

- 版本解析: `O(n)`，其中 n 是版本号字符串的长度
- 版本比较: `O(m)`，其中 m 是版本号中数字部分的最大长度
- 版本排序: `O(n log n)`，其中 n 是版本列表的长度
- 范围查询: `O(log n)`，基于有序版本组的二分查找

---

## 📄 许可证

<div align="center">
  
本项目采用 [MIT 许可证](./LICENSE) - 详见 LICENSE 文件

<b>Copyright © 2023-2025 scagogogo</b>

</div>