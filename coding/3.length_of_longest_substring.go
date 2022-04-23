package coding

// pwwkew -> 3
func LengthOfLongestSubstring(s string) int {
	if len(s) < 1 {
		return 0
	}
	left := 0
	right := 0
	// 元素 -> 索引
	m := map[string]int{}
	max := 1
	for right < len(s) {
		b := string(s[right])
		i, ok := m[b]
		if ok {
			if i+1 > left {
				left = i + 1
			}
		}
		m[b] = right
		curMax := right - left + 1
		if curMax > max {
			max = curMax
		}
		right++
	}
	return max
}
