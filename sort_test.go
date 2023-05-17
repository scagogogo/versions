package versions

import (
	"fmt"
	"github.com/golang-infrastructure/go-shuffle"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
