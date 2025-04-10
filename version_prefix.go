package versions

// VersionPrefix 表示版本中数字部分之前的部分
//
// VersionPrefix 是一个字符串类型，用于表示和操作版本号的前缀部分。
// 在版本号格式中，前缀是位于数字部分之前的字符串，如 "v1.2.3" 中的 "v"。
// 前缀可以为空，表示版本号直接以数字部分开始，如 "1.2.3"。
//
// 使用示例:
//
//	// 检查前缀是否为空
//	prefix := versions.VersionPrefix("v")
//	if !prefix.IsEmpty() {
//	    fmt.Printf("版本前缀: %s\n", prefix)
//	}
//
//	// 在解析版本号时处理前缀
//	version := versions.NewVersion("v1.2.3")
//	fmt.Printf("版本 %s 的前缀是: %s\n", version.Raw, version.Prefix)
type VersionPrefix string

// EmptyVersionPrefix 表示一个空的前缀
//
// 该常量用于表示版本号没有前缀的情况，如纯数字版本号 "1.2.3"。
// 它也用于检查版本前缀是否为空。
const EmptyVersionPrefix VersionPrefix = ""

// IsEmpty 返回前缀是否为空
//
// 该方法检查版本前缀是否为空，即是否等于 EmptyVersionPrefix。
//
// 返回:
//   - bool: 如果前缀为空则返回 true，否则返回 false
//
// 使用示例:
//
//	version := versions.NewVersion("1.2.3") // 没有前缀
//	if version.Prefix.IsEmpty() {
//	    fmt.Println("版本没有前缀")
//	}
//
//	version2 := versions.NewVersion("v1.2.3") // 有前缀
//	if !version2.Prefix.IsEmpty() {
//	    fmt.Printf("版本前缀是: %s\n", version2.Prefix)
//	}
func (x VersionPrefix) IsEmpty() bool {
	return x == EmptyVersionPrefix
}

// 去除了分隔符的纯净的前缀，暂时用不到先不实现了
//func (x VersionPrefix) PurePrefix() string {
//
//}
