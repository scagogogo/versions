package version

import (
	"testing"
)

func TestVersionGroup_QueryRangeVersions(t *testing.T) {

	// TODO 2023-5-16 19:12:11 单元测试，范围查询有个bug
	// https://repo1.maven.org/maven2/com/alibaba/fastjson/
	// 1.2.69/
	versions := NewVersions("1.2.6", "")
	group := NewVersionGroupFromVersions(versions)
	group.QueryRangeVersions(NewVersion("1.2.69/ "), NewVersion(""))

}
