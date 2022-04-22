package coding

import (
	"sort"
	"strconv"
	"strings"
)

// 输入：nums = [10,2]
// 输出："210
func LargestNumber(nums []int) string {
	var str []string
	for _, num := range nums {
		str = append(str, strconv.Itoa(num))
	}
	sort.Slice(str, func(i, j int) bool { return str[i]+str[j] > str[j]+str[i] })
	if str[0] == "0" {
		return "0"
	}
	return strings.Join(str, "")
}
