package offer

func findContinuousSequence(target int) (res [][]int) {
	left := 1
	right := 2
	sum := left + right
	nums := make([]int, target)
	for i := range nums {
		nums[i] = i
	}
	for left <= target/2 {
		if sum < target {
			right++
			sum += right
		} else if sum > target {
			sum -= left
			left++
		} else {
			res = append(res, nums[left:right+1])
			right++
			sum += right
			sum -= left
			left++
		}
	}
	return
}
