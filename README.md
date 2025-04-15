## 添加依赖

```bash
go get -u github.com/scagogogo/versions
```

## versions API文档

### 1. 概述
`versions` 包是一个Go语言库，旨在处理软件版本号。它提供了从文件读取版本号、解析、比较、排序以及查询特定版本范围内的功能。

### 2. 数据类型和常量

#### Version
- **属性**:
  - `Raw` string: 原始版本号字符串。
  - `PublicTime` time.Time: 版本发布的时间。
  - `VersionNumbers` VersionNumbers: 版本号中的数字部分。
  - `Prefix` VersionPrefix: 版本号数字部分之前的前缀。
  - `Suffix` VersionSuffix: 版本号数字部分之后的后缀。

#### VersionNumbers
- **属性**:
  - 一个整数切片，表示版本号中的数字部分。

#### VersionPrefix
- **属性**:
  - 一个字符串，表示版本号数字部分之前的前缀。

#### VersionSuffix
- **属性**:
  - 一个字符串，表示版本号数字部分之后的后缀。

#### ContainsPolicy
- **说明**: 用于控制版本查询时的参数。
  - `ContainsPolicyNone`: 未指定。
  - `ContainsPolicyYes`: 包含。
  - `ContainsPolicyNo`: 不包含。

### 3. 函数

#### ReadVersionsFromFile
- **描述**: 从文件中读取版本号，一行一个版本号，返回解析后的版本切片。
- **参数**:
  - `filepath string`: 文件路径。
- **返回值**:
  - `[]*Version`: 解析后的版本号切片。
  - `error`: 错误信息。

**示例代码**:
```go
versions, err := versions.ReadVersionsFromFile("path/to/versions.txt")
if err != nil {
    log.Fatal(err)
}
fmt.Println(versions)
```

#### ReadVersionsStringFromFile
- **描述**: 从文件中读取版本号，一行一个版本号，返回原始字符串切片。
- **参数**:
  - `filepath string`: 文件路径。
- **返回值**:
  - `[]string`: 原始版本号字符串切片。
  - `error`: 错误信息。

**示例代码**:
```go
versionStrings, err := versions.ReadVersionsStringFromFile("path/to/versions.txt")
if err != nil {
    log.Fatal(err)
}
fmt.Println(versionStrings)
```

#### Group
- **描述**: 对版本号进行分组。
- **参数**:
  - `versions []*Version`: 版本号切片。
- **返回值**:
  - `map[string]*VersionGroup`: 版本号分组映射。

**示例代码**:
```go
groupedVersions := versions.Group(versions)
for _, group := range groupedVersions {
    fmt.Println(group.Versions())
}
```

#### SortVersionStringSlice
- **描述**: 对字符串形式的版本数组排序，返回值也是字符串形式的。
- **参数**:
  - `versionStringSlice []string`: 版本号字符串切片。
- **返回值**:
  - `[]string`: 排序后的版本号字符串切片。

**示例代码**:
```go
sortedVersions := versions.SortVersionStringSlice(versionStrings)
fmt.Println(sortedVersions)
```

#### SortVersionSlice
- **描述**: 对版本号切片进行排序。
- **参数**:
  - `versions []*Version`: 版本号切片。
- **返回值**:
  - `[]*Version`: 排序后的版本号切片。

**示例代码**:
```go
sortedVersions := versions.SortVersionSlice(versions)
fmt.Println(sortedVersions)
```

#### NewSortedVersionGroups
- **描述**: 为版本号创建有序的分组。
- **参数**:
  - `versions []*Version`: 版本号切片。
- **返回值**:
  - `*SortedVersionGroups`: 有序的版本号分组。

**示例代码**:
```go
sortedGroups := versions.NewSortedVersionGroups(versions)
fmt.Println(sortedGroups.groupSlice)
```

#### QueryRange
- **描述**: 查询指定范围内的版本号。
- **参数**:
  - `start *tuple.Tuple2[*Version, ContainsPolicy]`: 开始版本号和包含策略。
  - `end *tuple.Tuple2[*Version, ContainsPolicy]`: 结束版本号和包含策略。
- **返回值**:
  - `[]*Version`: 查询范围内的版本号切片。

**示例代码**:
```go
startVersion := versions.NewVersion("1.0.0")
endVersion := versions.NewVersion("2.0.0")
queryRange := sortedGroups.QueryRange(&tuple.Tuple2[startVersion, versions.ContainsPolicyYes], &tuple.Tuple2[endVersion, versions.ContainsPolicyNo])
fmt.Println(queryRange)
```

#### VisualizeVersions
- **描述**: 可视化版本号之间的关系和结构，以文本形式展示版本层次结构。
- **参数**:
  - `versions []*Version`: 要可视化的版本集合。
  - `w io.Writer`: 输出写入的目标。
  - `maxItems int`: 每个版本组最多显示的版本数量，0表示不限制。
- **返回值**: 无直接返回值，结果写入指定的writer。

**示例代码**:
```go
versions, _ := versions.ReadVersionsFromFile("path/to/versions.txt")
versions.VisualizeVersions(versions, os.Stdout, 5)
```

**输出示例**:
```
版本总数: 10
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

#### VisualizeVersionGroups
- **描述**: 可视化版本组之间的层次关系，以树状结构展示版本组织架构。
- **参数**:
  - `versions []*Version`: 要可视化的版本集合。
  - `w io.Writer`: 输出写入的目标。
- **返回值**: 无直接返回值，结果写入指定的writer。

**示例代码**:
```go
versions, _ := versions.ReadVersionsFromFile("path/to/versions.txt")
versions.VisualizeVersionGroups(versions, os.Stdout)
```

**输出示例**:
```
版本总数: 15
版本组数: 6

├─ 1 (1个版本)
│ ├─ 1.0 (3个版本)
│ └─ 1.1 (2个版本)
└─ 2 (2个版本)
  ├─ 2.0 (4个版本)
  └─ 2.1 (3个版本)
```

### 4. 注意事项
- 确保传入的文件路径正确，且文件格式符合要求。
- 版本号的解析和比较依赖于正确的格式，非标准格式可能会导致解析错误。
- 时间比较依赖于 `PublicTime` 被正确设置。

### 5. 错误处理
每个函数都可能返回一个错误值，调用时应该检查并妥善处理这些错误。

### 6. 版本兼容性
该库设计为兼容Go语言的标准版本号格式，对于非标准格式的版本号可能需要额外的处理。

这个API文档提供了`versions`包的基本使用方法和函数描述，包括详细的代码示例，帮助开发者理解和使用这个版本号处理库。