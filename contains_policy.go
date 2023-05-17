package version

// ContainsPolicy 用于控制版本查询的时候的参数
type ContainsPolicy int

const (

	// ContainsPolicyNoneSingle 不包含此版本号，但是如果此组下有其它版本号的大于或者小于的话则会被包含
	ContainsPolicyNoneSingle ContainsPolicy = iota

	// ContainsPolicyNoneGroup 不包含此版本号所在的组
	ContainsPolicyNoneGroup

	ContainsPolicySingle
	ContainsPolicyGroup
)
