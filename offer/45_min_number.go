package offer

import (
	"sort"
	"strconv"
	"strings"
)

func MinNumber(nums []int) string {
	var str []string
	for _, n := range nums {
		s := strconv.Itoa(n)
		str = append(str, s)
	}
	sort.Slice(str, func(i, j int) bool { return str[i]+str[j] < str[j]+str[i] })
	return strings.Join(str, "")
}
