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
