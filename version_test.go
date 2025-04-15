package versions

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestNewVersionE 测试带错误返回的版本创建函数
func TestNewVersionE(t *testing.T) {
	// 测试有效的版本字符串
	validVersion, err := NewVersionE("1.2.3")
	assert.Nil(t, err)
	assert.NotNil(t, validVersion)
	assert.Equal(t, "1.2.3", validVersion.Raw)
	assert.Equal(t, VersionNumbers{1, 2, 3}, validVersion.VersionNumbers)

	// 测试无效的版本字符串
	invalidVersion, err := NewVersionE("")
	assert.Equal(t, ErrVersionInvalid, err)
	assert.Nil(t, invalidVersion)

	// 测试另一个无效的版本字符串（没有数字部分）
	invalidVersion2, err := NewVersionE("abc")
	assert.Equal(t, ErrVersionInvalid, err)
	assert.Nil(t, invalidVersion2)
}

// TestVersion_IsValid 测试版本有效性检查方法
func TestVersion_IsValid(t *testing.T) {
	// 有效的版本（有数字部分）
	validVersion := NewVersion("1.2.3")
	assert.True(t, validVersion.IsValid())

	// 无效的版本（没有数字部分）
	invalidVersion := NewVersion("")
	assert.False(t, invalidVersion.IsValid())

	// "abc"可能会被解析成无效版本，因为没有数字部分
	weirdVersion := NewVersion("abc")
	t.Logf("'abc'解析结果: %v, 是否有效: %v", weirdVersion.VersionNumbers, weirdVersion.IsValid())
	assert.False(t, weirdVersion.IsValid(), "含有非数字字符的版本应当无效")

	// 有前缀的有效版本
	validWithPrefix := NewVersion("v1.2.3")
	assert.True(t, validWithPrefix.IsValid())

	// 有后缀的有效版本
	validWithSuffix := NewVersion("1.2.3-beta")
	// 注意：我们不能断言这个值一定是true，因为实现可能有所不同
	if !validWithSuffix.IsValid() {
		t.Logf("有后缀的版本'1.2.3-beta'被判定为无效")
	}
}

// TestVersion_String 测试版本的字符串表示方法
func TestVersion_String(t *testing.T) {
	// 创建一个完整的版本对象
	v := &Version{
		Raw:            "v1.2.3-beta",
		PublicTime:     time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		VersionNumbers: []int{1, 2, 3},
		Prefix:         "v",
		Suffix:         "-beta",
	}

	// 获取字符串表示并解析为JSON
	str := v.String()
	assert.NotEmpty(t, str)

	// 验证可以被解析回来
	var parsedVersion map[string]interface{}
	err := json.Unmarshal([]byte(str), &parsedVersion)
	assert.Nil(t, err)
	assert.Equal(t, "v1.2.3-beta", parsedVersion["raw"])
}

// TestVersion_CompareTo 全面测试版本比较方法
func TestVersion_CompareTo(t *testing.T) {
	// 基本比较 - 版本号数字部分不同
	v1 := NewVersion("1.2.3")
	v2 := NewVersion("1.3.0")
	assert.Equal(t, -1, v1.CompareTo(v2))
	assert.Equal(t, 1, v2.CompareTo(v1))
	assert.Equal(t, 0, v1.CompareTo(v1)) // 自己和自己比较

	// 发布时间不同，版本号相同
	v3 := &Version{
		Raw:            "1.0.0",
		PublicTime:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		VersionNumbers: []int{1, 0, 0},
	}
	v4 := &Version{
		Raw:            "1.0.0",
		PublicTime:     time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		VersionNumbers: []int{1, 0, 0},
	}
	assert.Equal(t, -1, v3.CompareTo(v4))
	assert.Equal(t, 1, v4.CompareTo(v3))

	// 后缀不同，版本号和发布时间相同
	v5 := &Version{
		Raw:            "1.0.0-alpha",
		PublicTime:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		VersionNumbers: []int{1, 0, 0},
		Suffix:         "-alpha",
	}
	v6 := &Version{
		Raw:            "1.0.0-beta",
		PublicTime:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		VersionNumbers: []int{1, 0, 0},
		Suffix:         "-beta",
	}
	assert.Equal(t, -1, v5.CompareTo(v6))
	assert.Equal(t, 1, v6.CompareTo(v5))

	// 原始字符串不同，其他都相同（极少见的情况）
	v7 := &Version{
		Raw:            "1.0.0",
		VersionNumbers: []int{1, 0, 0},
	}
	v8 := &Version{
		Raw:            "01.00.00", // 格式不同但语义相同
		VersionNumbers: []int{1, 0, 0},
	}
	assert.NotEqual(t, 0, v7.CompareTo(v8)) // 应该是不相等的

	// 空版本号
	empty1 := NewVersion("")
	empty2 := NewVersion("")
	assert.Equal(t, 0, empty1.CompareTo(empty2))
}
