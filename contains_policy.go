package versions

// ContainsPolicy 用于控制版本查询的时候的参数
type ContainsPolicy int

const (

	// ContainsPolicyNone 未指定
	ContainsPolicyNone ContainsPolicy = iota

	// ContainsPolicyYes 包含
	ContainsPolicyYes

	// ContainsPolicyNo 不包含
	ContainsPolicyNo
)
