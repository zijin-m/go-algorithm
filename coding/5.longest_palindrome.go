package coding

func LongestPalindrome(s string) string {
	str := maracherStr(s)
	C, R := -1, -1
	help := make([]int, len(str))
	longest := 0
	start := -1
	end := -1
	for i := 0; i < len(str); i++ {
		if R > i {
			help[i] = min(help[(2*C)-i], R-i)
		} else {
			help[i] = 0
		}
		// 两边扩张
		for i+help[i] < len(str) && i-help[i] > -1 {
			if str[i+help[i]] != str[i-help[i]] {
				break
			}
			help[i]++
		}
		// 更新 C R 的位置
		if i+help[i] > R {
			R = i + help[i]
			C = i
		}
		// 记录最大的位置
		if help[i] > longest {
			longest = help[i]
			// 开始位置，因为 help[i] 算上了自身
			// 如第一个字符的最长为 1
			// 此处开始位置需要加上 1 （如 0 - 1 + 1 开始为 0）
			start = i - help[i] + 1
			// 结束位置
			end = i + help[i]
		}
	}
	ans := ""
	for i := start; i < end; i++ {
		if str[i] != '#' {
			ans += string(str[i])
		}
	}
	return ans
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func maracherStr(s string) string {
	buf := make([]byte, len(s)*2+1)
	index := 0
	for i := 0; i < len(buf); i++ {
		if i&1 == 0 {
			buf[i] = '#'
		} else {
			buf[i] = s[index]
			index++
		}
	}
	return string(buf)
}
