package versions

import (
	"fmt"
	"testing"

	"github.com/golang-infrastructure/go-shuffle"
	"github.com/stretchr/testify/assert"
)

// TestSortVersionSlice 测试版本对象数组排序功能
//
// 这个测试函数验证 SortVersionSlice 函数能否正确排序 Version 对象数组。
// 测试步骤:
// 1. 从测试数据文件中读取版本号列表
// 2. 随机打乱版本号顺序
// 3. 调用 SortVersionSlice 函数进行排序
// 4. 验证排序结果是否满足版本号的自然排序规则
//
// 预期结果:
// - 排序后的版本号数组应按照版本号的自然顺序排列
// - 对于任意相邻的两个版本，后一个版本应大于或等于前一个版本
func TestSortVersionSlice(t *testing.T) {

	// 读取测试文件中的版本号
	versions, err := ReadVersionsFromFile("./test_data/fast_json_versions.txt")
	assert.Nil(t, err)

	// 顺序打乱
	shuffle.Shuffle(versions)

	// 排序
	sortedVersions := SortVersionSlice(versions)

	// 检查是否有序
	for i := 1; i < len(sortedVersions); i++ {
		assert.True(t, sortedVersions[i].CompareTo(sortedVersions[i-1]) >= 0)
	}

	for _, v := range sortedVersions {
		fmt.Println(v.Raw)
	}

}

// TestSortVersionStringSlice 测试版本字符串数组排序功能
//
// 这个测试函数验证 SortVersionStringSlice 函数能否正确排序版本号字符串数组。
// 与 TestSortVersionSlice 不同，该测试直接处理字符串形式的版本号。
//
// 测试步骤:
// 1. 从测试数据文件中读取版本号字符串列表
// 2. 随机打乱版本号顺序
// 3. 调用 SortVersionStringSlice 函数进行排序
// 4. 验证排序结果是否满足版本号的自然排序规则
//
// 预期结果:
// - 排序后的版本号字符串数组应按照版本号的自然顺序排列
// - 对于任意相邻的两个版本字符串，后一个版本应大于或等于前一个版本
func TestSortVersionStringSlice(t *testing.T) {
	// 读取测试文件中的版本号
	//versions, err := ReadVersionsStringFromFile("./test_data/fast_json_versions.txt")
	versions, err := ReadVersionsStringFromFile("./test_data/org.apache.tomcat_tomcat-juli.txt")
	assert.Nil(t, err)

	shuffle.Shuffle(versions)
	versions = SortVersionStringSlice(versions)
	for i := 1; i < len(versions); i++ {
		assert.True(t, NewVersion(versions[i]).CompareTo(NewVersion(versions[i-1])) >= 0)
		t.Log(versions[i])
	}
}
