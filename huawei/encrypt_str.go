package huawei

// pasword___a123"__"456"________"78__timeout__""_100
func EncryptStr(str string, k int) string {
	// 当前是第几个命令
	cur := 0
	// 是否有 " 等待匹配
	flag := false
	len := 0
	for i, ch := range str {
		if ch != '_' && ch != '#' {
			len++
			continue
		}
		if ch == '#' {
			len++
			if str[i-1] == '_' {
				flag = false
			} else {
				flag = true
			}
			continue
		}
		if ch == '_' && !flag {
			cur++
			flag = false
			if cur > k {
				end := i
				start := end - len
				return str[0:start] + "******" + str[end:]
			}
			len = 0
			continue
		}
		len++

	}
	return "ERROR"
}
