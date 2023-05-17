package version

// Group 对版本号进行分组
func Group(versions []*Version) map[string]*VersionGroup {
	groupMap := make(map[string]*VersionGroup, 0)
	for _, v := range versions {
		group := groupMap[v.BuildGroupID()]
		if group == nil {
			group = NewVersionGroup(v.VersionNumbers)
			groupMap[v.BuildGroupID()] = group
		}
		group.Add(v)
	}
	return groupMap
}
