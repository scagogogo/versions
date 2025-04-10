package versions

import (
	"sort"
)

// SortVersionStringSlice 对字符串形式的版本数组进行排序
//
// 该函数接收一个字符串形式的版本号数组，将其解析为 Version 对象进行排序，
// 然后将排序结果转换回字符串数组返回。排序遵循版本号的自然排序规则。
//
// 参数:
//   - versionStringSlice: 待排序的版本号字符串数组
//
// 返回:
//   - []string: 排序后的版本号字符串数组
//
// 使用示例:
//
//	versions := []string{"1.2.0", "1.0.0", "1.10.0", "2.0.0"}
//	sorted := versions.SortVersionStringSlice(versions)
//	// 结果: ["1.0.0", "1.2.0", "1.10.0", "2.0.0"]
func SortVersionStringSlice(versionStringSlice []string) []string {
	versions := NewVersions(versionStringSlice...)
	slice := SortVersionSlice(versions)
	sortedVersionStringSlice := make([]string, 0)
	for _, v := range slice {
		sortedVersionStringSlice = append(sortedVersionStringSlice, v.Raw)
	}
	return sortedVersionStringSlice
}

// SortVersionSlice 对版本号对象数组进行排序
//
// 该函数实现了版本号的分组排序算法：
// 1. 首先将版本号按照主版本号分组
// 2. 对版本组进行排序
// 3. 在每个组内对版本号进行排序
// 4. 最后将所有组中的版本号按顺序合并
//
// 参数:
//   - versions: 待排序的版本号对象数组
//
// 返回:
//   - []*Version: 排序后的版本号对象数组
//
// 使用示例:
//
//	versions := versions.NewVersions("1.2.0", "1.0.0", "1.10.0", "2.0.0")
//	sorted := versions.SortVersionSlice(versions)
//	for _, v := range sorted {
//	    fmt.Println(v.Raw)
//	}
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

// SortVersionGroupMap 对版本组映射进行排序
//
// 该函数将版本组映射转换为切片，并对切片中的版本组进行排序。
//
// 参数:
//   - versionGroupMap: 版本组映射，键为版本组ID，值为版本组对象
//
// 返回:
//   - []*VersionGroup: 排序后的版本组切片
//
// 使用示例:
//
//	groupMap := Group(versions)
//	sortedGroups := SortVersionGroupMap(groupMap)
//	for _, group := range sortedGroups {
//	    fmt.Printf("组: %s, 版本数: %d\n", group.GroupID, len(group.Versions))
//	}
func SortVersionGroupMap(versionGroupMap map[string]*VersionGroup) []*VersionGroup {
	groupSlice := make([]*VersionGroup, 0)
	for _, g := range versionGroupMap {
		groupSlice = append(groupSlice, g)
	}
	SortVersionGroupSlice(groupSlice)
	return groupSlice
}

// SortVersionGroupSlice 对版本组切片进行排序
//
// 该函数直接对传入的版本组切片进行原地排序，排序依据是版本组之间的比较规则。
//
// 参数:
//   - groupSlice: 待排序的版本组切片，排序操作直接修改该切片
//
// 使用示例:
//
//	groups := []*VersionGroup{group1, group2, group3}
//	SortVersionGroupSlice(groups)
//	// 现在 groups 已按版本组规则排序
func SortVersionGroupSlice(groupSlice []*VersionGroup) {
	sort.Slice(groupSlice, func(i, j int) bool {
		return groupSlice[i].CompareTo(groupSlice[j]) < 0
	})
}
