package versions

import (
	"sort"
)

// SortVersionStringSlice 对字符串形式的版本数组排序，返回值也是字符串形式的
func SortVersionStringSlice(versionStringSlice []string) []string {
	versions := NewVersions(versionStringSlice...)
	slice := SortVersionSlice(versions)
	sortedVersionStringSlice := make([]string, 0)
	for _, v := range slice {
		sortedVersionStringSlice = append(sortedVersionStringSlice, v.Raw)
	}
	return sortedVersionStringSlice
}

// SortVersionSlice 对版本号排序
func SortVersionSlice(versions []*Version) []*Version {

	// 先对所有的版本进行分组
	groupMap := Group(versions)

	// 对所有的分组排序
	groupSlice := SortVersionGroupMap(groupMap)

	// 收集所有的组，同时对组中的版本号也排序，保证最终收集到的版本号都是有序的
	newVersions := make([]*Version, 0)
	for _, g := range groupSlice {
		newVersions = append(newVersions, g.SortVersions()...)
	}
	return newVersions
}

func SortVersionGroupMap(versionGroupMap map[string]*VersionGroup) []*VersionGroup {
	groupSlice := make([]*VersionGroup, 0)
	for _, g := range versionGroupMap {
		groupSlice = append(groupSlice, g)
	}
	SortVersionGroupSlice(groupSlice)
	return groupSlice
}

func SortVersionGroupSlice(groupSlice []*VersionGroup) {
	sort.Slice(groupSlice, func(i, j int) bool {
		return groupSlice[i].CompareTo(groupSlice[j]) < 0
	})
}
