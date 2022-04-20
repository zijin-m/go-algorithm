package coding

func Calculate(s string) int {
	stack := []int{}
	preSign := '+'
	num := 0
	for i, c := range s {
		if isDigit(c) {
			num = num*10 + int(c) - '0'
		}
		// 非数字，非空格，或者如果是最后的数字，则进行计算
		if !isDigit(c) && c != ' ' || i == len(s)-1 {
			switch preSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			default:
				stack[len(stack)-1] /= num
			}
			num = 0
			preSign = c
		}
	}
	sum := 0
	for _, n := range stack {
		sum += n
	}
	return sum
}

func isDigit(c rune) bool {
	return '0' <= c && c <= '9'
}
