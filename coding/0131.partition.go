package coding

func Partition(s string) (res [][]string) {
	if len(s) == 0 {
		return
	}
	partition(s, 0, []string{}, &res)
	return
}

func partition(s string, start int, path []string, res *[][]string) {
	if start >= len(s) {
		tmp := make([]string, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}
	for i := start; i < len(s); i++ {
		if !isPalindrome(s, start, i) {
			continue
		}
		path = append(path, s[start:i+1])
		partition(s, i+1, path, res)
		path = path[:len(path)-1]
	}
}

func isPalindrome(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
