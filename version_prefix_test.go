package versions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestVersionPrefix_IsEmpty 测试版本前缀的空值检查方法
func TestVersionPrefix_IsEmpty(t *testing.T) {
	// 空前缀
	emptyPrefix := VersionPrefix("")
	assert.True(t, emptyPrefix.IsEmpty())
	assert.True(t, EmptyVersionPrefix.IsEmpty())

	// 非空前缀
	nonEmptyPrefix1 := VersionPrefix("v")
	assert.False(t, nonEmptyPrefix1.IsEmpty())

	nonEmptyPrefix2 := VersionPrefix("version-")
	assert.False(t, nonEmptyPrefix2.IsEmpty())

	// 特殊字符前缀
	specialPrefix := VersionPrefix(" ")
	assert.False(t, specialPrefix.IsEmpty())
}
