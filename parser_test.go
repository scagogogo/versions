package versions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersionStringParser_Parser(t *testing.T) {

	v := NewVersionStringParser("17-0.15.0-alpha1").Parse()
	assert.Equal(t, VersionPrefix("17-"), v.Prefix)
	assert.Equal(t, NewVersionNumbers([]int{0, 15, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-alpha1"), v.Suffix)

	v = NewVersionStringParser("2.10.0-M1-virtualized.rdev-4217-2012-01-24-g9118644").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, NewVersionNumbers([]int{2, 10, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-M1-virtualized.rdev-4217-2012-01-24-g9118644"), v.Suffix)

	v = NewVersionStringParser("0.0").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, NewVersionNumbers([]int{0, 0}), v.VersionNumbers)
	assert.Equal(t, EmptyVersionSuffix, v.Suffix)

	v = NewVersionStringParser("100").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{100}), v.VersionNumbers)
	assert.Equal(t, EmptyVersionSuffix, v.Suffix)

	v = NewVersionStringParser("1.").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{1}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("."), v.Suffix)

	v = NewVersionStringParser(".1").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{1}), v.VersionNumbers)
	assert.Equal(t, EmptyVersionSuffix, v.Suffix)

	v = NewVersionStringParser("1.1.alpha1").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{1, 1}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix(".alpha1"), v.Suffix)

	v = NewVersionStringParser("1.....1.alpha1").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{1, 1}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix(".alpha1"), v.Suffix)

	v = NewVersionStringParser("1.....1........alpha1").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{1, 1}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("........alpha1"), v.Suffix)

	v = NewVersionStringParser("15.05.01").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{15, 5, 1}), v.VersionNumbers)
	assert.Equal(t, EmptyVersionSuffix, v.Suffix)

	v = NewVersionStringParser("").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{}), v.VersionNumbers)
	assert.Equal(t, EmptyVersionSuffix, v.Suffix)

	v = NewVersionStringParser("2.0.0.Final").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{2, 0, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix(".Final"), v.Suffix)

	v = NewVersionStringParser("1.0.0-RC1").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{1, 0, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-RC1"), v.Suffix)

	v = NewVersionStringParser("3.0.0-M1").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{3, 0, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-M1"), v.Suffix)

	v = NewVersionStringParser("3.0.0-m4").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{3, 0, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-m4"), v.Suffix)

	v = NewVersionStringParser("1.0.0.CR2").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{1, 0, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix(".CR2"), v.Suffix)

	v = NewVersionStringParser("2.0.0-RC-1").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{2, 0, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-RC-1"), v.Suffix)

	v = NewVersionStringParser("7.45.0.t20201009").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{7, 45, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix(".t20201009"), v.Suffix)

	v = NewVersionStringParser("5.17.295.16").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{5, 17, 295, 16}), v.VersionNumbers)
	assert.Equal(t, EmptyVersionSuffix, v.Suffix)

	v = NewVersionStringParser("RELEASE126").Parse()
	assert.Equal(t, VersionPrefix("RELEASE"), v.Prefix)
	assert.Equal(t, VersionNumbers([]int{126}), v.VersionNumbers)
	assert.Equal(t, EmptyVersionSuffix, v.Suffix)

	// TODO 2023-5-31 11:25:49 此case未通过
	//v = NewVersionStringParser("RELEASE120-1").Parse()
	//assert.Equal(t, VersionPrefix("RELEASE"), v.Prefix)
	//assert.Equal(t, VersionNumbers([]int{120}), v.VersionNumbers)
	//assert.Equal(t, VersionSuffix("-1"), v.Suffix)

	// TODO 2023-5-31 11:25:49 此case未通过
	//v = NewVersionStringParser("RELEASE120-u1").Parse()
	//assert.Equal(t, VersionPrefix("RELEASE"), v.Prefix)
	//assert.Equal(t, VersionNumbers([]int{120}), v.VersionNumbers)
	//assert.Equal(t, VersionSuffix("-u1"), v.Suffix)

	v = NewVersionStringParser("0.26.1-v2-524.0").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{0, 26, 1}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-v2-524.0"), v.Suffix)

	v = NewVersionStringParser("1.0-alpha-1").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{1, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-alpha-1"), v.Suffix)

	v = NewVersionStringParser("7.45.0.t20201014").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{7, 45, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix(".t20201014"), v.Suffix)

	v = NewVersionStringParser("0.33.0-v2.892.0").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{0, 33, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-v2.892.0"), v.Suffix)

	v = NewVersionStringParser("0.31.0-v2.731.0").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{0, 31, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-v2.731.0"), v.Suffix)

	v = NewVersionStringParser("2.5.6-1-f8bff243").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{2, 5, 6}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-1-f8bff243"), v.Suffix)

	v = NewVersionStringParser("7.22.0.t042").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{7, 22, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix(".t042"), v.Suffix)

	v = NewVersionStringParser("7.22.0.t042").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{7, 22, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix(".t042"), v.Suffix)

	v = NewVersionStringParser("3.0.1-b18").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{3, 0, 1}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-b18"), v.Suffix)

	v = NewVersionStringParser("2.5.0-snapshot.20221107.10906.0.346bb489").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{2, 5, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-snapshot.20221107.10906.0.346bb489"), v.Suffix)

	v = NewVersionStringParser("0.9.0+121-bcc5decc").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{0, 9, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("+121-bcc5decc"), v.Suffix)

	v = NewVersionStringParser("394").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{394}), v.VersionNumbers)
	assert.Equal(t, EmptyVersionSuffix, v.Suffix)

	v = NewVersionStringParser("v1-rev4-1.18.0-rc").Parse()
	assert.Equal(t, VersionPrefix("v1-rev4-"), v.Prefix)
	assert.Equal(t, VersionNumbers([]int{1, 18, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-rc"), v.Suffix)

	v = NewVersionStringParser("1.7.0-snapshot.20201012.5405.0.af92198d").Parse()
	assert.Equal(t, EmptyVersionPrefix, v.Prefix)
	assert.Equal(t, VersionNumbers([]int{1, 7, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-snapshot.20201012.5405.0.af92198d"), v.Suffix)

	v = NewVersionStringParser("erpya-3.9.4-rc-1.0.2").Parse()
	assert.Equal(t, VersionPrefix("erpya-"), v.Prefix)
	assert.Equal(t, VersionNumbers([]int{3, 9, 4}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-rc-1.0.2"), v.Suffix)

	v = NewVersionStringParser("curl-7_85_0").Parse()
	assert.Equal(t, VersionPrefix("erpya-"), v.Prefix)
	assert.Equal(t, VersionNumbers([]int{7, 85, 0}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix("-rc-1.0.2"), v.Suffix)

	v = NewVersionStringParser("zoneinfo-tz").Parse()
	assert.Equal(t, VersionPrefix("zoneinfo-tz"), v.Prefix)
	assert.Equal(t, VersionNumbers([]int{}), v.VersionNumbers)
	assert.Equal(t, VersionSuffix(""), v.Suffix)

	// TODO
	//v = NewVersionStringParser("v1beta1-rev20191118-1.29.2").Parse()
	//assert.Equal(t, VersionPrefix("COMMIT-2fef389"), v.Prefix)
	//assert.Equal(t, VersionNumbers([]int{1, 18, 0}), v.VersionNumbers)
	//assert.Equal(t, VersionSuffix("-rc"), v.Suffix)

	// TODO sha1检查
	//v = NewVersionStringParser("18699aad7ce6e60980f876a27145b5b29e9fd55d").Parse()
	//assert.Equal(t, VersionPrefix("COMMIT-2fef389"), v.Prefix)
	//assert.Equal(t, VersionNumbers([]int{1, 18, 0}), v.VersionNumbers)
	//assert.Equal(t, VersionSuffix("-rc"), v.Suffix)

	// TODO
	//v = NewVersionStringParser("erpya-3.9.4-rc-1.0.4").Parse()
	//assert.Equal(t, VersionPrefix("erpya-"), v.Prefix)
	//assert.Equal(t, VersionNumbers([]int{7, 85, 0}), v.VersionNumbers)
	//assert.Equal(t, VersionSuffix("-rc-1.0.2"), v.Suffix)

	// TODO
	//v = NewVersionStringParser("COMMIT-2fef389").Parse()
	//assert.Equal(t, VersionPrefix("COMMIT-2fef389"), v.Prefix)
	//assert.Equal(t, VersionNumbers([]int{1, 18, 0}), v.VersionNumbers)
	//assert.Equal(t, VersionSuffix("-rc"), v.Suffix)

	// TODO 2023-5-16 17:03:44
	//v = NewVersionStringParser("  ").Parse()
	//assert.Equal(t, VersionPrefix("  "), v.Prefix)
	//assert.Equal(t, VersionNumbers([]int{}), v.VersionNumbers)
	//assert.Equal(t, VersionSuffix("  "), v.Suffix)

	// TODO 2023-5-16 17:03:40
	//v = NewVersionStringParser("a.a.a.b.c.123").Parse()
	//assert.Equal(t, VersionPrefix("  "), v.Prefix)
	//assert.Equal(t, VersionNumbers([]int{}), v.VersionNumbers)
	//assert.Equal(t, EmptyVersionSuffix, v.Suffix)

}
