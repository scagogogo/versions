package versions

import "github.com/golang-infrastructure/go-tuple"

type SortedVersionGroups struct {

	// 用于根据组的ID快速定位到这个组在排好序的切片中的位置
	groupIdToIndexMap map[string]int

	// 排好序的组
	groupSlice []*VersionGroup
}

// NewSortedVersionGroups 为版本号创建有序的分组
func NewSortedVersionGroups(versions []*Version) *SortedVersionGroups {

	// 先对所有的版本进行分组
	groupMap := Group(versions)

	// 对所有的分组排序
	groupSlice := SortVersionGroupMap(groupMap)

	// 然后构造有序Group
	groups := &SortedVersionGroups{
		groupIdToIndexMap: make(map[string]int),
		groupSlice:        groupSlice,
	}
	for i, g := range groupSlice {
		groups.groupIdToIndexMap[g.ID()] = i
	}
	return groups
}

func (x *SortedVersionGroups) QueryRange(start, end *tuple.Tuple2[*Version, ContainsPolicy]) []*Version {

	// 如果要查询的版本组都不存在的话则直接返回空即可
	var i int
	if start.V1.Raw != "0" {
		// 仅当开始的版本不为0的时候才进行跳，否则认为是从最开始遍历
		var exists bool
		i, exists = x.groupIdToIndexMap[start.V1.BuildGroupID()]
		if !exists {
			return nil
		}
	}

	versions := make([]*Version, 0)
	for i < len(x.groupSlice) {
		g := x.groupSlice[i]
		i++

		versions = append(versions, g.QueryRangeVersions(start, end)...)
	}
	return versions
}
