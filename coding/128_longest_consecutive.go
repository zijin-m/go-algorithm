package coding

func LongestConsecutive(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}
	max := 1
	for num := range set {
		// 前一个必然不存在，才开始尝试遍历
		if !set[num-1] {
			start := num
			curMax := 1
			for set[start+1] {
				start++
				curMax++
			}
			if curMax > max {
				max = curMax
			}
		}
	}
	return max
}
