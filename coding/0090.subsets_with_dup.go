package coding

import "sort"

func SubsetsWithDup(nums []int) (res [][]int) {
	if len(nums) < 1 {
		return
	}
	sort.Ints(nums)
	subsetsWithDup(nums, 0, []int{}, &res)
	return
}

func subsetsWithDup(nums []int, start int, help []int, res *[][]int) {
	tmp := make([]int, len(help))
	copy(tmp, help)
	*res = append(*res, tmp)

	if start >= len(nums) {
		return
	}

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		help = append(help, nums[i])
		subsetsWithDup(nums, i+1, help, res)
		help = help[:len(help)-1]
	}
}
