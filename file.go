package versions

import (
	"os"
	"strings"
)

// ReadVersionsFromFile 从文件中读取版本号，一行一个版本号，把解析好的版本切片返回
// 示例：
// 1.1.28
// 1.1.29
// 1.1.30
// 1.1.31
// 1.1.31.sec01
// 1.1.31.sec04
// 1.1.31.sec06
func ReadVersionsFromFile(filepath string) ([]*Version, error) {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	versions := make([]*Version, 0)
	for _, line := range strings.Split(string(bytes), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		v := NewVersionStringParser(line).Parse()
		versions = append(versions, v)
	}
	return versions, nil
}

func ReadVersionsStringFromFile(filepath string) ([]string, error) {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	versions := make([]string, 0)
	for _, line := range strings.Split(string(bytes), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		versions = append(versions, line)
	}
	return versions, nil
}
