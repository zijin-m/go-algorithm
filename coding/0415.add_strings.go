package coding

import "strconv"

func AddStrings(num1 string, num2 string) string {
	carry := 0
	ans := ""
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		var m, n, sum int
		if i >= 0 {
			m = int(num1[i] - '0')
		}
		if j >= 0 {
			n = int(num2[j] - '0')
		}
		sum, carry = (m+n+carry)%10, (m+n+carry)/10
		ans += strconv.Itoa(sum)
	}
	if carry == 1 {
		ans += "1"
	}
	return reverseStr(ans)
}

func reverseStr(str string) string {
	i, j := 0, len(str)-1
	bytes := []byte(str)
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	return string(bytes)
}
