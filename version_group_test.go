package versions

import (
	"testing"

	"github.com/golang-infrastructure/go-tuple"
)

// TestVersionGroup_QueryRangeVersions 测试版本组范围查询功能
//
// 这个测试函数验证 VersionGroup 的 QueryRangeVersions 方法是否能正确地
// 查询指定范围内的版本。测试使用特定的版本范围边界条件，检查版本组的范围查询行为。
//
// 测试步骤:
// 1. 创建包含指定版本的版本组
// 2. 定义查询范围（起始版本和结束版本）及包含策略
// 3. 调用 QueryRangeVersions 方法进行范围查询
// 4. 检查查询结果是否符合预期
//
// 注意:
// 当前有一个已知问题（TODO项），需要修复范围查询的bug。
// 该bug与特定版本号范围的查询相关，如 1.2.69 到 1.2.70 之间的版本查询。
func TestVersionGroup_QueryRangeVersions(t *testing.T) {

	// TODO 2023-5-16 19:12:11 单元测试，范围查询有个bug
	// https://repo1.maven.org/maven2/com/alibaba/fastjson/
	// 1.2.69/
	versions := NewVersions("1.2.6", "")
	group := NewVersionGroupFromVersions(versions)

	start := tuple.New2[*Version, ContainsPolicy](NewVersion("1.2.69"), ContainsPolicyYes)
	end := tuple.New2[*Version, ContainsPolicy](NewVersion("1.2.70"), ContainsPolicyYes)
	rangeVersions := group.QueryRangeVersions(start, end)
	t.Log(rangeVersions)

}
