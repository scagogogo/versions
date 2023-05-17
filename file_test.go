package versions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadVersionsFromFile(t *testing.T) {
	versions, err := ReadVersionsFromFile("test_data/fast_json_versions.txt")
	assert.Nil(t, err)
	assert.True(t, len(versions) > 0)
}
