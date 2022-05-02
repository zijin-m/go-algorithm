package coding

func Subsets(nums []int) (res [][]int) {
	if len(nums) < 1 {
		return
	}
	subsets(nums, 0, []int{}, &res)
	return
}

func subsets(nums []int, start int, path []int, res *[][]int) {
	tmp := make([]int, len(path))
	copy(tmp, path)
	*res = append(*res, tmp)
	if start >= len(nums) {
		return
	}
	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		subsets(nums, i+1, path, res)
		path = path[:len(path)-1]
	}
}
