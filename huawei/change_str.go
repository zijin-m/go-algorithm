package huawei

func ChangeStr(str string) string {
	if len(str) < 2 {
		return str
	}
	min := 'z' + 1
	minIdx := -1
	for i, ch := range str {
		if ch < min {
			min = ch
			minIdx = i
		}
		if min == 'a' {
			break
		}
	}
	if minIdx == 0 {
		return str
	}
	return str[minIdx:minIdx+1] + str[1:minIdx] + str[0:1] + str[minIdx+1:]
}
