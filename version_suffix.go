package versions

import compare_anything "github.com/golang-infrastructure/go-compare-anything"

// VersionSuffix 表示版本号的后缀，版本号中数字后面的部分
type VersionSuffix string

// EmptyVersionSuffix 空的后缀
const EmptyVersionSuffix VersionSuffix = ""

var _ compare_anything.Comparable[VersionSuffix] = VersionSuffix("")

// IsEmpty 判断版本后缀是否为空
func (x VersionSuffix) IsEmpty() bool {
	return x == EmptyVersionSuffix
}

// CompareTo TODO 2023-5-16 17:26:02 后缀比较算法引入权重，不同的后缀权重不同
func (x VersionSuffix) CompareTo(target VersionSuffix) int {
	if x > target {
		return 1
	} else if x < target {
		return -1
	} else {
		return 0
	}
}
