package versions

import (
	"encoding/json"
	"errors"
	compare_anything "github.com/golang-infrastructure/go-compare-anything"
	"time"
)

var (
	ErrVersionInvalid = errors.New("version invalid")
)

// Version 用于表示一个版本号
type Version struct {

	// 原始的版本号
	Raw string `json:"raw"`

	// 此版本的发布时间
	PublicTime time.Time `json:"public_time"`

	// 版本号中的数字部分
	VersionNumbers VersionNumbers `json:"version_numbers"`

	// 版本号数字部分之前的前缀
	Prefix VersionPrefix `json:"prefix"`

	// 版本号数字部分之后的后缀
	Suffix VersionSuffix `json:"suffix"`
}

var _ compare_anything.Comparable[*Version] = &Version{}

func NewVersion(versionStr string) *Version {
	return NewVersionStringParser(versionStr).Parse()
}

func NewVersionE(versionStr string) (*Version, error) {
	v := NewVersionStringParser(versionStr).Parse()
	if v.IsValid() {
		return v, nil
	} else {
		return nil, ErrVersionInvalid
	}
}

func NewVersions(versionStringSlice ...string) []*Version {
	versions := make([]*Version, len(versionStringSlice))
	for i, versionStr := range versionStringSlice {
		versions[i] = NewVersion(versionStr)
	}
	return versions
}

func (x *Version) IsValid() bool {
	// 只有当解析到了版本号数字的时候才认为是有效的版本号，其它情况认为是无效的版本
	return len(x.VersionNumbers) > 0
}

// BuildGroupID 构造版本所属的组的ID
func (x *Version) BuildGroupID() string {
	return x.VersionNumbers.BuildGroupID()
}

func (x *Version) CompareTo(target *Version) int {

	// 1. 先按照主版本号排序
	r := x.VersionNumbers.CompareTo(target.VersionNumbers)
	if r != 0 {
		return r
	}

	// 2. 然后按照发布时间排序
	r2 := target.PublicTime.UnixMilli() - x.PublicTime.UnixMilli()
	if r2 != 0 {
		if r2 > 0 {
			return 1
		} else {
			return -1
		}
	}

	// 3. 然后按照后缀的字典序排序
	return x.Suffix.CompareTo(target.Suffix)
}

func (x *Version) String() string {
	marshal, _ := json.Marshal(x)
	return string(marshal)
}
