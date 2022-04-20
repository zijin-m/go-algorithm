package coding

func IsValid(s string) bool {
	rule := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := []rune{}
	for _, c := range s {
		_, ok := rule[c]
		if ok {
			stack = append(stack, c)
		} else {
			n := len(stack)
			if n == 0 {
				return false
			}
			// (
			l := stack[n-1]
			stack = stack[:n-1]
			// )
			if c != rule[l] {
				return false
			}
		}
	}
	return true
}
