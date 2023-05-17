package versions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortedVersionGroups_QueryRange(t *testing.T) {

	versions, err := ReadVersionsFromFile("./test_data/fast_json_versions.txt")
	assert.Nil(t, err)
	groups := NewSortedVersionGroups(versions)

	queryRange := groups.QueryRange(NewVersion("1.2.59"), NewVersion("1.2.70"))
	for _, v := range queryRange {
		fmt.Println(v.Raw)
	}
}
