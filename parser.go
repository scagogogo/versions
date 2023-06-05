package versions

// TODO 2023-5-31 12:13:27 一次解析多条版本，让它们之间互相印证

// VersionStringParser 把版本从字符串形式解析为struct
type VersionStringParser struct {

	// 被解析的字符串原始样子
	versionStr string
	// 转为字符序列方便处理
	versionRunes []rune
	// 上面的字符序列，当前解析到哪个下标了
	i int

	// 解析结果
	v *Version
}

// NewVersionStringParser 创建一个版本号Parser，每次解析都需要重新创建新的Parser
func NewVersionStringParser(versionStr string) *VersionStringParser {
	return &VersionStringParser{
		versionStr:   versionStr,
		versionRunes: []rune(versionStr),
		i:            0,
	}
}

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

// 读取版本中的前缀部分，比如对于版本号 v0.0.1，则prefix为 v
func (x *VersionStringParser) readVersionPrefix() {

	// TODO 2023-5-31 12:14:00 使用正则来定位版本号数字的位置，如果版本号数字有多个的话则选择最长的一个，如果一样长则选择靠前的那个

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
func (x *VersionStringParser) IsDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

// 读取版本号中的数字部分
//
// Example:
// versionStr 1.2.48.sec06
// i 从这个下标开始读取
// return 1.0.10
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
func (x *VersionStringParser) IsVersionNumberDelimiter(c rune) bool {
	return c == '.'
}

// 把数字字符数组解析为int
func (x *VersionStringParser) parseDigitsToNumber(digits []rune) int {
	r := 0
	weight := 0
	for i := len(digits) - 1; i >= 0; i-- {
		r += int(digits[i]-'0') * x.pow(10, weight)
		weight++
	}
	return r
}

// 求幂
func (x *VersionStringParser) pow(p, q int) int {
	r := 1
	for q > 0 {
		r *= p
		q--
	}
	return r
}

// 读取字符串中的后缀
func (x *VersionStringParser) readVersionSuffix() {
	// 如果还有剩余的话，则将剩余都认为是后缀部分
	if x.i < len(x.versionRunes) {
		x.v.Suffix = VersionSuffix(x.versionRunes[x.i:])
	}
}
