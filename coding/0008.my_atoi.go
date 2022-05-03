package coding

import (
	"math"
	"unicode"
)

func MyAtoi(s string) int {
	if len(s) < 1 {
		return 0
	}
	var n int32 = 0
	numStart := false
	// 1 正数 -1 负数
	var sign int8 = 1
	for _, letter := range s {
		if !numStart && letter == ' ' {
			continue
		}
		if !numStart && letter == '+' {
			sign = 1
			numStart = true
			continue
		}
		if !numStart && letter == '-' {
			sign = -1
			numStart = true
			continue
		}
		numStart = true
		// 后续再碰到非数字
		if !unicode.IsDigit(letter) {
			break
		}
		digit := int32(letter - '0')
		// 越界
		if n > (math.MaxInt32-digit)/10 {
			if sign == 1 {
				return 1<<31 - 1
			} else {
				return -1 << 31
			}
		}
		n = n*10 + digit
	}
	return int(sign) * int(n)
}
