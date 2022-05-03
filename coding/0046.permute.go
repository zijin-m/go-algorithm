package coding

func Permute(nums []int) (res [][]int) {
	if len(nums) < 1 {
		return
	}
	permute(nums, []bool{}, []int{}, &res)
	return
}

func permute(nums []int, used []bool, path []int, res *[][]int) {
	if len(nums) == len(path) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		used[i] = true
		path = append(path, nums[i])
		permute(nums, used, path, res)
		used[i] = false
		path = path[:len(path)-1]
	}
}
