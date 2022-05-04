package coding

func LengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	length := 1
	d := make([]int, len(nums)+1)
	d[length] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > d[length] {
			length++
			d[length] = nums[i]
		} else {
			idx := binarySearch(d, nums[i], 1, length)
			d[idx] = nums[i]
		}
	}
	return length
}

func binarySearch(nums []int, target int, l, r int) int {
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
