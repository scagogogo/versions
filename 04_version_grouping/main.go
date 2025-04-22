package main

import (
	"fmt"
	"sort"

	"github.com/scagogogo/versions"
)

func main() {
	// 示例1：按主版本号分组
	fmt.Println("示例1 - 按主版本号分组:")
	versionStrings := []string{
		"1.0.0", "1.1.0", "1.2.0", "1.10.0",
		"2.0.0", "2.1.0", "2.1.1",
		"3.0.0-alpha", "3.0.0-beta", "3.0.0", "3.1.0",
	}

	// 将字符串转换为版本对象
	var versionObjs []*versions.Version
	for _, v := range versionStrings {
		versionObjs = append(versionObjs, versions.NewVersion(v))
	}

	// 按主版本号分组
	groupedVersions := versions.GroupVersionsByMajor(versionObjs)

	// 打印分组结果
	fmt.Println("版本分组结果:")
	for majorVersion, versions := range groupedVersions {
		fmt.Printf("主版本 %d:\n", majorVersion)
		for _, v := range versions {
			fmt.Printf("  - %s\n", v.Raw)
		}
	}
	/*
		输出示例:
		示例1 - 按主版本号分组:
		版本分组结果:
		主版本 1:
		  - 1.0.0
		  - 1.1.0
		  - 1.2.0
		  - 1.10.0
		主版本 2:
		  - 2.0.0
		  - 2.1.0
		  - 2.1.1
		主版本 3:
		  - 3.0.0-alpha
		  - 3.0.0-beta
		  - 3.0.0
		  - 3.1.0
	*/

	// 示例2：自定义分组函数
	fmt.Println("\n示例2 - 自定义分组函数:")
	// 按次版本号分组的函数
	groupByMinor := func(version *versions.Version) string {
		return fmt.Sprintf("%d.%d", version.Major, version.Minor)
	}

	// 使用自定义函数分组版本
	customGrouped := versions.GroupVersions(versionObjs, groupByMinor)

	// 为了有序显示，先对键进行排序
	var keys []string
	for k := range customGrouped {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 打印自定义分组结果
	fmt.Println("按次版本号分组结果:")
	for _, key := range keys {
		fmt.Printf("版本组 %s:\n", key)
		for _, v := range customGrouped[key] {
			fmt.Printf("  - %s\n", v.Raw)
		}
	}
	/*
		输出示例:
		示例2 - 自定义分组函数:
		按次版本号分组结果:
		版本组 1.0:
		  - 1.0.0
		版本组 1.1:
		  - 1.1.0
		版本组 1.2:
		  - 1.2.0
		版本组 1.10:
		  - 1.10.0
		版本组 2.0:
		  - 2.0.0
		版本组 2.1:
		  - 2.1.0
		  - 2.1.1
		版本组 3.0:
		  - 3.0.0-alpha
		  - 3.0.0-beta
		  - 3.0.0
		版本组 3.1:
		  - 3.1.0
	*/

	// 示例3：按发布阶段分组
	fmt.Println("\n示例3 - 按发布阶段分组:")
	releaseVersions := []string{
		"1.0.0", "1.1.0-alpha", "1.1.0-beta", "1.1.0-rc1", "1.1.0",
		"2.0.0-alpha", "2.0.0-beta.1", "2.0.0-beta.2", "2.0.0-rc1",
	}

	// 将字符串转换为版本对象
	var releaseObjs []*versions.Version
	for _, v := range releaseVersions {
		releaseObjs = append(releaseObjs, versions.NewVersion(v))
	}

	// 按发布阶段分组的函数
	groupByStage := func(version *versions.Version) string {
		if version.IsPreRelease() {
			return "预发布版本"
		}
		return "正式版本"
	}

	// 使用自定义函数按发布阶段分组
	stageGrouped := versions.GroupVersions(releaseObjs, groupByStage)

	// 打印按发布阶段分组的结果
	fmt.Println("按发布阶段分组结果:")
	if preReleases, exists := stageGrouped["预发布版本"]; exists {
		fmt.Println("预发布版本:")
		for _, v := range preReleases {
			fmt.Printf("  - %s\n", v.Raw)
		}
	}

	if releases, exists := stageGrouped["正式版本"]; exists {
		fmt.Println("正式版本:")
		for _, v := range releases {
			fmt.Printf("  - %s\n", v.Raw)
		}
	}
	/*
		输出示例:
		示例3 - 按发布阶段分组:
		按发布阶段分组结果:
		预发布版本:
		  - 1.1.0-alpha
		  - 1.1.0-beta
		  - 1.1.0-rc1
		  - 2.0.0-alpha
		  - 2.0.0-beta.1
		  - 2.0.0-beta.2
		  - 2.0.0-rc1
		正式版本:
		  - 1.0.0
		  - 1.1.0
	*/
}
