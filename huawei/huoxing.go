package huawei

// 已知火星人使用的运算符为#、$，其与地球人的等价公式如下：

// x#y = 2x+3y+4

// x$y = 3*x+y+2

// 其中x、y是无符号整数
// 地球人公式按C语言规则计算
// 火星人公式中，$的优先级高于#，相同的运算符，按从左到右的顺序计算 现有一段火星人的字符串报文，请你来翻译并计算结果。
// 7#6$5#12
// 226
func Huoxing(str string) int {
	stack := []int{}
	num := 0
	preSign := '#'
	for i, ch := range str {
		if isDigit(ch) {
			num = num*10 + int(ch) - '0'
		}
		if !isDigit(ch) || i == len(str)-1 {
			if preSign == '#' {
				stack = append(stack, num)
			} else {
				stack[len(stack)-1] = 3*stack[len(stack)-1] + num + 2
			}
			num = 0
			preSign = ch
		}
	}
	ans := stack[0]
	for i := 1; i < len(stack); i++ {
		ans = 2*ans + 3*stack[i] + 4
	}
	return ans
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}
