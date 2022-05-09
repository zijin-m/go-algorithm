package coding

import "math"

func FindDuplicates(nums []int) (res []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		num := int(math.Abs(float64(nums[i])))
		if nums[num-1] > 0 {
			nums[num-1] *= -1
		} else {
			res = append(res, num)
		}
	}
	return
}
