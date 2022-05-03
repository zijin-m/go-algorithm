package coding

func FindSubsequences(nums []int) (res [][]int) {
	if len(nums) < 1 {
		return
	}
	findSubsequences(nums, 0, []int{}, &res)
	return
}

func findSubsequences(nums []int, start int, path []int, res *[][]int) {
	if len(path) >= 2 && isIncr(path) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
	}
	if start >= len(nums) {
		return
	}
	used := make(map[int]bool)
	for i := start; i < len(nums); i++ {
		if used[nums[i]] {
			continue
		}
		used[nums[i]] = true
		path = append(path, nums[i])
		findSubsequences(nums, i+1, path, res)
		path = path[:len(path)-1]
	}
}

func isIncr(nums []int) bool {
	i := 1
	for i < len(nums) {
		if nums[i] < nums[i-1] {
			return false
		}
		i++
	}
	return true
}
