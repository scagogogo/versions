package version

import (
	compare_anything "github.com/golang-infrastructure/go-compare-anything"
	"sort"
)

// VersionGroup 表示一个版本组，一个版本组中可能有一个或多个版本
type VersionGroup struct {

	// 组的版本号中的数字部分
	GroupVersionNumbers VersionNumbers

	// 组中的版本号都有哪些
	VersionMap map[string]*Version
}

var _ compare_anything.Comparable[*VersionGroup] = &VersionGroup{}

// NewVersionGroup 创建组的时候必须传递能够在包范围下唯一区分组的组id，这个id选择的是版本中的数字部分
func NewVersionGroup(groupVersionNumbers VersionNumbers) *VersionGroup {
	return &VersionGroup{
		GroupVersionNumbers: groupVersionNumbers,
		VersionMap:          make(map[string]*Version, 0),
	}
}

func NewVersionGroupFromVersions(versions []*Version) *VersionGroup {
	if len(versions) == 0 {
		return nil
	}
	group := NewVersionGroup(versions[0].VersionNumbers)
	for _, v := range versions {
		group.Add(v)
	}
	return group
}

// Add 把给定的版本添加到本版本组中
func (x *VersionGroup) Add(v *Version) bool {
	_, exists := x.VersionMap[v.Raw]
	x.VersionMap[v.Raw] = v
	return exists
}

// Contains 判断本版本组中是否包含给定的版本
func (x *VersionGroup) Contains(v *Version) bool {
	_, exists := x.VersionMap[v.Raw]
	return exists
}

// ID 返回组的ID
func (x *VersionGroup) ID() string {
	return x.GroupVersionNumbers.BuildGroupID()
}

// CompareTo 比较两个版本组的大小
func (x *VersionGroup) CompareTo(target *VersionGroup) int {
	return x.GroupVersionNumbers.CompareTo(target.GroupVersionNumbers)
}

// Versions 返回组下的所有版本
func (x *VersionGroup) Versions() []*Version {
	slice := make([]*Version, 0)
	for _, v := range x.VersionMap {
		slice = append(slice, v)
	}
	return slice
}

// SortVersions 对组下的所有版本进行排序返回
func (x *VersionGroup) SortVersions() []*Version {
	versions := x.Versions()
	sort.Slice(versions, func(i, j int) bool {
		return versions[i].CompareTo(versions[j]) < 0
	})
	return versions
}


// QueryRangeVersions 获取组内要求的区间内的版本
func (x *VersionGroup) QueryRangeVersions(start, end *Version, containsEndPolicy ...ContainsPolicy) []*Version {

	// 默认情况下是包含
	if len(containsEndPolicy) == 0 {
		containsEndPolicy = append(containsEndPolicy, ContainsEndPolicyIfSuffixEmptyAllElseLe)
	}

	// 因为这里认为同一个版本组中不会有特别多的版本，所以就不再做索引表直接跳了，如果后面有发现特殊情况再来做优化
	sortedVersions := x.SortVersions()
	versions := make([]*Version, 0)
	for _, v := range sortedVersions {

		// 获取完了则结束
		if v.CompareTo(end) > 0 {
			break
		}

		// 如果是主版本大于等于开始版本
		versionNumberIsContains := v.VersionNumbers.CompareTo(start.VersionNumbers) >= 0 && v.VersionNumbers.CompareTo(end.VersionNumbers) == 0
		if versionNumberIsContains {
			versions = append(versions, v)
		} else if v.VersionNumbers.CompareTo(end.VersionNumbers) == 0 && x.isRangeContains(v, end, containsEndPolicy[0]) {
			versions = append(versions, v)
		}
	}
	return versions
}

func (x *VersionGroup) isRangeContains(needJudgeVersion, rangeEndVersion *Version, policy ContainsPolicy) bool {
	// 根据不同的包含策略进行
	switch policy {
	case ContainsPolicyNoneSingle:

	}
}
