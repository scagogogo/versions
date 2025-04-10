package versions

import (
	"os"
	"strings"
)

// ReadVersionsFromFile 从文件中读取版本号并解析为Version对象
//
// 该函数从指定的文件中读取版本号列表，每行一个版本号，并将其解析为Version对象数组。
// 函数会自动忽略空行和进行行尾空白字符的清理。
//
// 参数:
//   - filepath: 包含版本号列表的文件路径
//
// 返回:
//   - []*Version: 解析后的Version对象数组
//   - error: 如果文件读取失败则返回相应错误
//
// 示例文件内容:
//
//	1.1.28
//	1.1.29
//	1.1.30
//	1.1.31
//	1.1.31.sec01
//	1.1.31.sec04
//	1.1.31.sec06
//
// 使用示例:
//
//	// 从版本列表文件中读取版本
//	versions, err := versions.ReadVersionsFromFile("./versions.txt")
//	if err != nil {
//	    log.Fatalf("读取版本文件失败: %v", err)
//	}
//
//	// 打印解析的版本数
//	fmt.Printf("共读取 %d 个版本\n", len(versions))
//
//	// 对版本进行排序
//	sortedVersions := versions.SortVersionSlice(versions)
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

// ReadVersionsStringFromFile 从文件中读取版本号字符串
//
// 该函数从指定的文件中读取版本号列表，每行一个版本号，并返回字符串数组。
// 与 ReadVersionsFromFile 不同，此函数不会将版本号解析为Version对象，
// 而是保留原始字符串形式，适用于不需要解析和比较的场景。
//
// 参数:
//   - filepath: 包含版本号列表的文件路径
//
// 返回:
//   - []string: 读取的版本号字符串数组
//   - error: 如果文件读取失败则返回相应错误
//
// 使用示例:
//
//	// 从版本列表文件中读取版本字符串
//	versionStrings, err := versions.ReadVersionsStringFromFile("./versions.txt")
//	if err != nil {
//	    log.Fatalf("读取版本文件失败: %v", err)
//	}
//
//	// 使用版本字符串
//	for _, vStr := range versionStrings {
//	    fmt.Printf("发现版本: %s\n", vStr)
//	}
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
