package huawei

import (
	"sort"
)

// 未点亮路段长度
// 在一段路径上，一共有N(N ∈ [1, 100000]N∈[1,100000])座路灯，已知路灯间距均为100m100m。现给出每个路灯的照亮范围light[i],且照亮范围在区间[1, 100 * N][1,100∗N]内。求出该段路中，总共未照亮的路段长度。
//
// 示例1
//
// 输入： N = 2, light[] = [50,50]
// 输出：0
//
// 示例2：
//
// 输入： N = 4, light[] = [50,70,20,70]
//
// 输出：20
//
// 可以借助 https://leetcode-cn.com/problems/merge-intervals/ 合并区间的思路来解决
func StreetLight(light []int, n int) int {
	sum := 0
	// to [][]int
	var intervals [][]int
	for i := 0; i < len(light); i++ {
		l := i*100 - light[i]
		r := i*100 + light[i]
		intervals = append(intervals, []int{l, r})
	}
	merged := Merge(intervals)
	for i := 0; i < len(merged)-1; i++ {
		sum += merged[i+1][0] - merged[i][1]
	}
	return sum
}

func Merge(intervals [][]int) (res [][]int) {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	for _, interval := range intervals {
		if len(res) == 0 || res[len(res)-1][1] < interval[0] {
			res = append(res, []int{interval[0], interval[1]})
		} else {
			res[len(res)-1][1] = max(res[len(res)-1][1], interval[1])
		}
	}
	return
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
