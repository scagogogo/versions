package version

// VersionPrefix 表示版本中数字部分之前的部分
type VersionPrefix string

// EmptyVersionPrefix 表示一个空的前缀
const EmptyVersionPrefix VersionPrefix = ""

// IsEmpty 返回前缀是否为空
func (x VersionPrefix) IsEmpty() bool {
	return x == EmptyVersionPrefix
}

// 去除了分隔符的纯净的前缀，暂时用不到先不实现了
//func (x VersionPrefix) PurePrefix() string {
//
//}
