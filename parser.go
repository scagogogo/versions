package versions

// TODO 2023-5-31 12:13:27 一次解析多条版本，让它们之间互相印证

// VersionStringParser 把版本从字符串形式解析为struct
//
// VersionStringParser 负责将版本号字符串解析为结构化的 Version 对象。
// 它实现了版本号字符串的词法分析，将字符串划分为前缀、数字部分和后缀三个组成部分。
//
// 解析过程：
// 1. 首先识别和提取版本号的前缀部分（如 "v"）
// 2. 然后识别和提取版本号的数字部分（如 "1.2.3"）
// 3. 最后识别和提取版本号的后缀部分（如 "-beta1"）
//
// 使用示例:
//
//	// 创建一个版本解析器
//	parser := versions.NewVersionStringParser("v1.2.3-beta1")
//
//	// 解析版本字符串
//	version := parser.Parse()
//
//	// 使用解析结果
//	fmt.Printf("前缀: %s\n", version.Prefix)
//	fmt.Printf("数字部分: %v\n", version.VersionNumbers)
//	fmt.Printf("后缀: %s\n", version.Suffix)
type VersionStringParser struct {

	// versionStr 被解析的字符串原始样子
	versionStr string
	// versionRunes 转为字符序列方便处理
	versionRunes []rune
	// i 上面的字符序列，当前解析到哪个下标了
	i int

	// v 解析结果
	v *Version
}

// NewVersionStringParser 创建一个版本号Parser
//
// 该方法创建一个新的版本号解析器实例。每次解析都需要重新创建新的Parser，
// 因为解析状态保存在Parser对象中。
//
// 参数:
//   - versionStr: 要解析的版本号字符串
//
// 返回:
//   - *VersionStringParser: 新创建的版本号解析器
//
// 使用示例:
//
//	parser := versions.NewVersionStringParser("v1.2.3-rc1")
//	version := parser.Parse()
func NewVersionStringParser(versionStr string) *VersionStringParser {
	return &VersionStringParser{
		versionStr:   versionStr,
		versionRunes: []rune(versionStr),
		i:            0,
	}
}

// Parse 解析版本号字符串
//
// 该方法按照预定义的规则解析版本号字符串，提取前缀、数字部分和后缀，
// 并构造一个完整的 Version 对象返回。
//
// 返回:
//   - *Version: 解析后的版本对象
//
// 使用示例:
//
//	parser := versions.NewVersionStringParser("v1.2.3-beta1")
//	version := parser.Parse()
//	fmt.Printf("版本号: %s\n", version.Raw)
func (x *VersionStringParser) Parse() *Version {

	x.v = &Version{
		Raw: x.versionStr,
	}

	// 读取前缀
	x.readVersionPrefix()

	// 读取版本号的数字部分
	x.readVersionNumbers()

	// 读取版本号中的后缀
	x.readVersionSuffix()

	return x.v
}

// readVersionPrefix 读取版本中的前缀部分
//
// 该方法解析版本号字符串中的前缀部分。例如对于版本号 "v0.0.1"，前缀为 "v"。
// 方法通过向前搜索数字部分的起始位置，然后回溯确定前缀边界。
//
// 注意:
//   - TODO 2023-5-31 12:14:00 使用正则来定位版本号数字的位置，如果版本号数字有多个的话则选择最长的一个，如果一样长则选择靠前的那个
func (x *VersionStringParser) readVersionPrefix() {

	// 一直读取，直到读取到版本号中数字部分的分隔符
	for x.i < len(x.versionRunes) {
		if x.IsVersionNumberDelimiter(x.versionRunes[x.i]) {
			x.i--
			break
		} else {
			x.i++
		}
	}
	// 控制右边界
	if x.i >= len(x.versionRunes) && x.i > 0 {
		x.i--
	}

	// 然后再回退一个版本号数字
	for x.i > 0 {
		if x.IsDigit(x.versionRunes[x.i]) {
			x.i--
		} else {
			x.i++
			break
		}
	}
	// 控制左边界
	if x.i < 0 {
		x.i = 0
	}

	if x.i > 0 {
		x.v.Prefix = VersionPrefix(x.versionRunes[0:x.i])
	}
}

// IsDigit 判断是否是数字
//
// 该方法检查给定的字符是否为数字字符（0-9）。
//
// 参数:
//   - c: 要检查的字符
//
// 返回:
//   - bool: 如果是数字则返回 true，否则返回 false
func (x *VersionStringParser) IsDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

// readVersionNumbers 读取版本号中的数字部分
//
// 该方法解析版本号字符串中的数字部分，如 "1.2.3"。它识别由分隔符（通常是点）
// 分隔的数字序列，并将其转换为整数数组。
//
// 示例:
//   - 对于 "1.2.48.sec06"，从适当的位置开始读取，可能返回 [1,2,48]
//   - 解析在遇到非数字且非分隔符的字符时停止
func (x *VersionStringParser) readVersionNumbers() {
	numbers := make([]int, 0)
	nowNumberDigits := make([]rune, 0)
	for x.i < len(x.versionRunes) {
		c := x.versionRunes[x.i]

		if x.IsDigit(c) {
			nowNumberDigits = append(nowNumberDigits, c)
			x.i++
			continue
		}

		// 版本号读取已经完毕了一部分，则将其处理一下加入到版本数字数组中
		if len(nowNumberDigits) != 0 {
			number := x.parseDigitsToNumber(nowNumberDigits)
			numbers = append(numbers, number)
			nowNumberDigits = make([]rune, 0)
		}

		// 如果不是版本号的分隔符的话则不再继续读取了
		if !x.IsVersionNumberDelimiter(c) {
			break
		}
		x.i++
	}
	// 如果只有版本号结尾的话，则最后一部分要能够正确处理
	if len(nowNumberDigits) != 0 {
		number := x.parseDigitsToNumber(nowNumberDigits)
		numbers = append(numbers, number)
	}

	x.v.VersionNumbers = NewVersionNumbers(numbers)

	// 上次读取的字符必须是数字，否则就吐出来直到是个数字
	for x.i > 0 && !x.IsDigit(x.versionRunes[x.i-1]) {
		x.i--
	}

	return
}

// IsVersionNumberDelimiter 判断是否是版本数字的分隔符
//
// 该方法检查给定的字符是否为版本号数字部分的分隔符（目前仅支持点号）。
//
// 参数:
//   - c: 要检查的字符
//
// 返回:
//   - bool: 如果是分隔符则返回 true，否则返回 false
func (x *VersionStringParser) IsVersionNumberDelimiter(c rune) bool {
	return c == '.'
}

// parseDigitsToNumber 把数字字符数组解析为int
//
// 该方法将数字字符数组转换为对应的整数值。例如 ['1','2','3'] 将被转换为 123。
//
// 参数:
//   - digits: 数字字符数组
//
// 返回:
//   - int: 解析后的整数值
func (x *VersionStringParser) parseDigitsToNumber(digits []rune) int {
	r := 0
	weight := 0
	for i := len(digits) - 1; i >= 0; i-- {
		r += int(digits[i]-'0') * x.pow(10, weight)
		weight++
	}
	return r
}

// pow 求幂
//
// 该方法计算 p 的 q 次方（p^q）。
//
// 参数:
//   - p: 底数
//   - q: 指数
//
// 返回:
//   - int: 计算结果
func (x *VersionStringParser) pow(p, q int) int {
	r := 1
	for q > 0 {
		r *= p
		q--
	}
	return r
}

// readVersionSuffix 读取字符串中的后缀
//
// 该方法解析版本号字符串中的后缀部分。后缀是版本号数字部分之后的所有内容。
// 例如对于版本号 "1.2.3-beta1"，后缀为 "-beta1"。
func (x *VersionStringParser) readVersionSuffix() {
	// 如果还有剩余的话，则将剩余都认为是后缀部分
	if x.i < len(x.versionRunes) {
		x.v.Suffix = VersionSuffix(x.versionRunes[x.i:])
	}
}
