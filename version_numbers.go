package versions

import (
	compare_anything "github.com/golang-infrastructure/go-compare-anything"
	"strconv"
	"strings"
)

const (
	// DefaultVersionDelimiter 默认的版本数字分隔符，即使版本的数字分隔符是其它的，解析完毕之后也会被统一为这个
	DefaultVersionDelimiter = "."
)

// VersionNumbers 表示版本号中的数字部分
type VersionNumbers []int

var _ compare_anything.Comparable[[]int] = VersionNumbers{}

func NewVersionNumbers(versionNumbers []int) VersionNumbers {
	return VersionNumbers(versionNumbers)
}

// CompareTo 比较两个版本号数字的大小
func (x VersionNumbers) CompareTo(target []int) int {
	for i := 0; i < len(x) || i < len(target); i++ {
		if i >= len(x) || i >= len(target) {
			return len(x) - len(target)
		} else if x[i] != target[i] {
			return x[i] - target[i]
		} else {
			// 版本当前部分的数字数相等，继续比较下一位
		}
	}
	return 0
}

// BuildGroupID 根据版本字符串构造版本组的ID，把版本中的数字部分以.分隔拼接为字符串
func (x VersionNumbers) BuildGroupID() string {
	s := strings.Builder{}
	for i, v := range x {
		s.WriteString(strconv.Itoa(v))
		if i+1 != len(x) {
			s.WriteString(DefaultVersionDelimiter)
		}
	}
	return s.String()
}
