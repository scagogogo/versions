package versions

import (
	"fmt"
	"testing"

	"github.com/golang-infrastructure/go-shuffle"
	"github.com/golang-infrastructure/go-tuple"
	"github.com/stretchr/testify/assert"
)

// TestSortedVersionGroups_QueryRange 测试有序版本组的范围查询功能
//
// 这个测试函数验证 SortedVersionGroups 的 QueryRange 方法是否能正确地
// 查询指定范围内的版本。测试使用实际的版本数据集进行范围查询测试。
//
// 测试步骤:
// 1. 从测试数据文件中读取版本号列表（使用 fastjson 的版本列表）
// 2. 随机打乱版本号顺序，以确保排序逻辑正常工作
// 3. 创建 SortedVersionGroups 实例
// 4. 定义查询范围（从 "0" 到 "1.2.83"）和包含策略
// 5. 调用 QueryRange 方法进行范围查询
// 6. 输出查询结果的版本号
//
// 注意:
// 测试文件中还包含了其他版本数据集的查询代码，这些部分已被注释掉，
// 但可以取消注释以测试其他数据集上的范围查询功能。可用的测试数据集包括:
// - org.jboss_jboss-ejb-client.txt
// - de.tum.in.ase_artemis-java-test-sandbox.txt
func TestSortedVersionGroups_QueryRange(t *testing.T) {

	// fastjson
	versions, err := ReadVersionsFromFile("./test_data/fast_json_versions.txt")
	assert.Nil(t, err)

	// 故意打乱顺序，以防止影响到测试结果
	shuffle.Shuffle(versions)

	groups := NewSortedVersionGroups(versions)

	start := tuple.New2[*Version, ContainsPolicy](NewVersion("0"), ContainsPolicyYes)
	end := tuple.New2[*Version, ContainsPolicy](NewVersion("1.2.83"), ContainsPolicyNo)
	queryRange := groups.QueryRange(start, end)
	for _, v := range queryRange {
		fmt.Println(v.Raw)
	}

	// 以下是其他数据集的测试代码，目前已注释掉
	// 可以取消注释以测试不同数据集上的范围查询功能

	//versions, err := ReadVersionsFromFile("./test_data/org.jboss_jboss-ejb-client.txt")
	//assert.Nil(t, err)
	//groups := NewSortedVersionGroups(versions)
	//
	//start := tuple.New2[*Version, ContainsPolicy](NewVersion("0"), ContainsPolicyYes)
	//end := tuple.New2[*Version, ContainsPolicy](NewVersion("4.0.39"), ContainsPolicyNo)
	//queryRange := groups.QueryRange(start, end)
	//for _, v := range queryRange {
	//	fmt.Println(v.Raw)
	//}

	//versions, err := ReadVersionsFromFile("./test_data/de.tum.in.ase_artemis-java-test-sandbox.txt")
	//assert.Nil(t, err)
	//groups := NewSortedVersionGroups(versions)
	//
	//start := tuple.New2[*Version, ContainsPolicy](NewVersion("0"), ContainsPolicyYes)
	//end := tuple.New2[*Version, ContainsPolicy](NewVersion("1.8.0"), ContainsPolicyNo)
	//queryRange := groups.QueryRange(start, end)
	//for _, v := range queryRange {
	//	fmt.Println(v.Raw)
	//}
}
