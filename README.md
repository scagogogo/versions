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

<details open>
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

<details open>
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

## 🧩 完整API文档

<div align="center">
<h3>核心类型与功能详解</h3>
</div>

### Version 类型

<details open>
<summary><b>结构定义</b></summary>

```go
type Version struct {
    // 原始版本号字符串
    Raw string
    
    // 版本发布时间
    PublicTime time.Time
    
    // 版本号数字部分，例如 1.2.3 中的 [1,2,3]
    VersionNumbers VersionNumbers
    
    // 版本号前缀，例如 v1.2.3 中的 "v"
    Prefix VersionPrefix
    
    // 版本号后缀，例如 1.2.3-beta 中的 "-beta"
    Suffix VersionSuffix
}
```
</details>

<details open>
<summary><b>NewVersion</b> - 创建版本号对象</summary>

```go
func NewVersion(versionString string) *Version
```

**参数:**
- `versionString string`: 版本号字符串，如 "1.2.3", "v1.0.0-beta" 等

**返回值:**
- `*Version`: 解析后的版本对象

**示例:**
```go
version := versions.NewVersion("v1.2.3-rc1")
fmt.Printf("前缀: %s, 版本号: %v, 后缀: %s\n", 
    version.Prefix, version.VersionNumbers, version.Suffix)
// 输出: 前缀: v, 版本号: [1 2 3], 后缀: -rc1
```
</details>

<details open>
<summary><b>NewVersionE</b> - 创建版本号对象（带错误返回）</summary>

```go
func NewVersionE(versionString string) (*Version, error)
```

**参数:**
- `versionString string`: 版本号字符串

**返回值:**
- `*Version`: 解析后的版本对象
- `error`: 解析过程中可能发生的错误

**示例:**
```go
version, err := versions.NewVersionE("v1.2.3-rc1")
if err != nil {
    log.Fatalf("版本号解析失败: %v", err)
}
```
</details>

<details open>
<summary><b>IsValid</b> - 检查版本号是否有效</summary>

```go
func (v *Version) IsValid() bool
```

**返回值:**
- `bool`: 版本号是否有效，必须至少包含一个版本数字

**示例:**
```go
version := versions.NewVersion("v1.2.3")
if version.IsValid() {
    fmt.Println("版本号有效")
}
```
</details>

<details open>
<summary><b>CompareTo</b> - 比较两个版本号</summary>

```go
func (v *Version) CompareTo(other *Version) int
```

**参数:**
- `other *Version`: 要比较的另一个版本对象

**返回值:**
- `int`: 小于0表示当前版本小于other，等于0表示相等，大于0表示当前版本大于other

**示例:**
```go
v1 := versions.NewVersion("1.2.3")
v2 := versions.NewVersion("1.3.0")
result := v1.CompareTo(v2)
if result < 0 {
    fmt.Printf("%s 小于 %s\n", v1.Raw, v2.Raw)
}
```
</details>

<details open>
<summary><b>String</b> - 获取版本号字符串表示</summary>

```go
func (v *Version) String() string
```

**返回值:**
- `string`: 版本号的字符串表示，通常等同于原始版本号

**示例:**
```go
version := versions.NewVersion("v1.2.3")
fmt.Println(version.String()) // 输出: v1.2.3
```
</details>

### VersionNumbers 类型

<details open>
<summary><b>结构定义与方法</b></summary>

```go
// VersionNumbers 是整数切片，表示版本号的数字部分
type VersionNumbers []int

// 获取主版本号
func (v VersionNumbers) MajorVersion() int

// 获取次版本号
func (v VersionNumbers) MinorVersion() int

// 获取修订版本号
func (v VersionNumbers) PatchVersion() int

// 比较两个版本号数字部分
func (v VersionNumbers) CompareTo(other VersionNumbers) int
```

**示例:**
```go
version := versions.NewVersion("1.2.3")
major := version.VersionNumbers.MajorVersion() // 返回 1
minor := version.VersionNumbers.MinorVersion() // 返回 2
patch := version.VersionNumbers.PatchVersion() // 返回 3
```
</details>

### VersionPrefix 类型

<details open>
<summary><b>结构定义与方法</b></summary>

```go
// VersionPrefix 是字符串，表示版本号前缀
type VersionPrefix string

// 检查前缀是否为空
func (v VersionPrefix) IsEmpty() bool

// 比较两个前缀
func (v VersionPrefix) CompareTo(other VersionPrefix) int
```

**示例:**
```go
version := versions.NewVersion("v1.2.3")
if !version.Prefix.IsEmpty() {
    fmt.Printf("版本前缀: %s\n", version.Prefix) // 输出: 版本前缀: v
}
```
</details>

### VersionSuffix 类型

<details open>
<summary><b>结构定义与方法</b></summary>

```go
// VersionSuffix 是字符串，表示版本号后缀
type VersionSuffix string

// 检查后缀是否为空
func (v VersionSuffix) IsEmpty() bool

// 比较两个后缀
func (v VersionSuffix) CompareTo(other VersionSuffix) int
```

**示例:**
```go
version := versions.NewVersion("1.2.3-beta")
if !version.Suffix.IsEmpty() {
    fmt.Printf("版本后缀: %s\n", version.Suffix) // 输出: 版本后缀: -beta
}
```
</details>

### ContainsPolicy 类型

<details open>
<summary><b>定义与常量</b></summary>

```go
// ContainsPolicy 用于控制版本范围查询时是否包含边界版本
type ContainsPolicy int

const (
    // 未指定包含策略
    ContainsPolicyNone ContainsPolicy = iota
    
    // 包含边界版本
    ContainsPolicyYes
    
    // 不包含边界版本
    ContainsPolicyNo
)
```

**使用场景:**
在范围查询中指定是否包含起始版本和结束版本，例如：
- `[1.0.0, 2.0.0]` - 包含1.0.0和2.0.0
- `(1.0.0, 2.0.0)` - 不包含1.0.0和2.0.0
- `[1.0.0, 2.0.0)` - 包含1.0.0但不包含2.0.0
</details>

### VersionGroup 类型

<details open>
<summary><b>结构定义与方法</b></summary>

```go
// VersionGroup 表示具有相同主版本号的一组版本
type VersionGroup struct {
    // ...内部字段
}

// 创建新的版本组
func NewVersionGroup(id string) *VersionGroup

// 添加版本到组中
func (g *VersionGroup) Add(version *Version)

// 检查组是否包含某个版本
func (g *VersionGroup) Contains(version *Version) bool

// 获取组ID
func (g *VersionGroup) ID() string

// 获取组中的所有版本
func (g *VersionGroup) Versions() []*Version

// 获取组中版本的数量
func (g *VersionGroup) Count() int

// 按版本号排序组内的版本
func (g *VersionGroup) SortVersions() []*Version

// 查询范围内的版本
func (g *VersionGroup) QueryRangeVersions(start, end *Version) []*Version
```

**示例:**
```go
// 创建版本组
group := versions.NewVersionGroup("1")

// 添加版本
group.Add(versions.NewVersion("1.0.0"))
group.Add(versions.NewVersion("1.1.0"))
group.Add(versions.NewVersion("1.2.0"))

// 获取组内所有版本
allVersions := group.Versions()
fmt.Printf("版本组 %s 包含 %d 个版本\n", group.ID(), group.Count())

// 排序组内版本
sortedVersions := group.SortVersions()
```
</details>

### SortedVersionGroups 类型

<details open>
<summary><b>结构定义与方法</b></summary>

```go
// SortedVersionGroups 表示一组有序的版本组
type SortedVersionGroups struct {
    // ...内部字段
}

// 创建新的有序版本组
func NewSortedVersionGroups(versions []*Version) *SortedVersionGroups

// 获取所有版本组ID
func (s *SortedVersionGroups) GroupIDs() []string

// 查询指定范围内的版本
func (s *SortedVersionGroups) QueryRange(
    start *tuple.Tuple2[*Version, ContainsPolicy],
    end *tuple.Tuple2[*Version, ContainsPolicy],
) []*Version
```

**示例:**
```go
// 创建测试版本列表
versions := []*versions.Version{
    versions.NewVersion("1.0.0"),
    versions.NewVersion("1.1.0"),
    versions.NewVersion("2.0.0"),
    versions.NewVersion("2.1.0"),
    versions.NewVersion("3.0.0"),
}

// 创建有序版本组
sortedGroups := versions.NewSortedVersionGroups(versions)

// 获取所有组ID
groupIDs := sortedGroups.GroupIDs()
fmt.Printf("共有 %d 个版本组: %v\n", len(groupIDs), groupIDs)

// 执行范围查询：获取1.0.0（包含）到2.0.0（不包含）之间的所有版本
startVersion := versions.NewVersion("1.0.0")
endVersion := versions.NewVersion("2.0.0")
startTuple := tuple.New2[*versions.Version, versions.ContainsPolicy](
    startVersion, versions.ContainsPolicyYes)
endTuple := tuple.New2[*versions.Version, versions.ContainsPolicy](
    endVersion, versions.ContainsPolicyNo)
    
result := sortedGroups.QueryRange(startTuple, endTuple)
fmt.Printf("查询结果包含 %d 个版本\n", len(result))
```
</details>

### 文件操作函数

<details open>
<summary><b>从文件读取版本号</b></summary>

```go
// 读取文件中的版本号字符串
func ReadVersionsStringFromFile(filepath string) ([]string, error)

// 读取并解析文件中的版本号
func ReadVersionsFromFile(filepath string) ([]*Version, error)
```

**参数:**
- `filepath string`: 文件路径，文件中每行应包含一个版本号

**返回值:**
- `[]string` 或 `[]*Version`: 版本号字符串或版本对象切片
- `error`: 读取过程中可能发生的错误

**示例:**
```go
// 从文件读取版本号字符串
versionStrings, err := versions.ReadVersionsStringFromFile("versions.txt")
if err != nil {
    log.Fatalf("读取版本号失败: %v", err)
}

// 从文件读取并解析版本号
versionObjects, err := versions.ReadVersionsFromFile("versions.txt")
if err != nil {
    log.Fatalf("解析版本号失败: %v", err)
}
```
</details>

### 排序函数

<details open>
<summary><b>版本排序函数</b></summary>

```go
// 对版本字符串切片进行排序
func SortVersionStringSlice(versionStringSlice []string) []string

// 对版本对象切片进行排序
func SortVersionSlice(versions []*Version) []*Version
```

**参数:**
- `versionStringSlice []string` 或 `versions []*Version`: 要排序的版本号切片

**返回值:**
- `[]string` 或 `[]*Version`: 排序后的版本号切片

**示例:**
```go
// 排序版本号字符串
unsortedStrings := []string{"2.0.0", "1.0.0", "1.10.0", "1.2.0"}
sortedStrings := versions.SortVersionStringSlice(unsortedStrings)
// 结果: ["1.0.0", "1.2.0", "1.10.0", "2.0.0"]

// 排序版本对象
unsortedVersions := []*versions.Version{
    versions.NewVersion("2.0.0"),
    versions.NewVersion("1.0.0"),
    versions.NewVersion("1.10.0"),
}
sortedVersions := versions.SortVersionSlice(unsortedVersions)
```
</details>

### 分组函数

<details open>
<summary><b>版本分组函数</b></summary>

```go
// 将版本列表按主版本号分组
func Group(versions []*Version) map[string]*VersionGroup
```

**参数:**
- `versions []*Version`: 要分组的版本对象列表

**返回值:**
- `map[string]*VersionGroup`: 以组ID为键的版本组映射

**示例:**
```go
// 创建版本列表
versionList := []*versions.Version{
    versions.NewVersion("1.0.0"),
    versions.NewVersion("1.1.0"),
    versions.NewVersion("2.0.0"),
    versions.NewVersion("2.1.0"),
}

// 按主版本号分组
groupMap := versions.Group(versionList)

// 遍历分组结果
for groupID, group := range groupMap {
    fmt.Printf("版本组 %s 包含 %d 个版本:\n", groupID, group.Count())
    for _, v := range group.Versions() {
        fmt.Printf("  - %s\n", v.Raw)
    }
}
```
</details>

### 可视化函数

<details open>
<summary><b>版本可视化函数</b></summary>

```go
// 以文本树形式可视化版本结构
func VisualizeVersions(versions []*Version, w io.Writer, maxItemsPerGroup int)

// 以文本树形式可视化版本组层次结构
func VisualizeVersionGroups(versions []*Version, w io.Writer)
```

**参数:**
- `versions []*Version`: 要可视化的版本对象列表
- `w io.Writer`: 输出写入目标，通常是 `os.Stdout`
- `maxItemsPerGroup int`: 每组最多显示的版本数量，0表示不限制

**示例:**
```go
// 可视化版本结构
versions.VisualizeVersions(versionList, os.Stdout, 0)

// 可视化版本组层次结构
versions.VisualizeVersionGroups(versionList, os.Stdout)
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