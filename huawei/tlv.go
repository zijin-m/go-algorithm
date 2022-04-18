package huawei

import (
	"strconv"
	"strings"
)

func TLV(data []string, tag string) string {
	for i := 0; i < len(data); {
		length, _ := strconv.ParseInt(data[i+2]+data[i+1], 16, 16)
		if data[i] == tag {
			return strings.Join(data[i+3:i+3+int(length)], " ")
		} else {
			i += int(length) + 3
		}
	}
	return ""
}
