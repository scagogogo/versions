package versions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestReadVersionsFromFile 测试从文件读取版本号功能
//
// 这个测试函数验证 ReadVersionsFromFile 函数能否正确地从文件中读取版本号列表。
// 测试使用 test_data 目录下的 fast_json_versions.txt 文件作为测试数据源。
//
// 测试步骤:
// 1. 调用 ReadVersionsFromFile 读取测试文件中的版本号
// 2. 验证读取过程中没有发生错误（error 为 nil）
// 3. 验证成功读取到版本号（返回的数组长度大于0）
//
// 预期结果:
// - 函数应当成功读取文件并解析其中的版本号
// - 返回的版本号数组不应为空
func TestReadVersionsFromFile(t *testing.T) {
	versions, err := ReadVersionsFromFile("test_data/fast_json_versions.txt")
	assert.Nil(t, err)
	assert.True(t, len(versions) > 0)
}
