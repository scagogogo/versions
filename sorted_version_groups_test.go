package versions

import (
	"fmt"
	"github.com/golang-infrastructure/go-shuffle"
	"github.com/golang-infrastructure/go-tuple"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
