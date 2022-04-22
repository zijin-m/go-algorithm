package huawei

// 输入一组数字，数字之间用空格隔开，判断输入的数字是否可以满足A=B+2C，每个元素最多只可用一次。若有满足的数字组合，依次输出A、B、C三个数字，之间用空格隔开；若无满足条件的组合，输出0。
//
// 例如：输入 2 3 5 7 输出 7 3 2
var res []int

func HasNumCombine(nums []int) []int {
	res = []int{}
	path := []int{}
	nm := make(map[int]bool)
	for _, num := range nums {
		nm[num] = true
	}
	hasCombine(nums, nm, 0, path)
	return res
}

func hasCombine(nums []int, nm map[int]bool, start int, path []int) {
	if len(res) > 0 {
		return
	}
	if len(path) == 2 {
		a1 := path[0] + 2*path[1]
		if exist(nm, a1) {
			res = append(res, a1, path[0], path[1])
		}
		a2 := path[1] + 2*path[0]
		if exist(nm, a2) {
			res = append(res, a2, path[1], path[0])
		}
		return
	}
	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		hasCombine(nums, nm, i+1, path)
		path = path[:len(path)-1]
	}
}

func exist(nm map[int]bool, a int) bool {
	_, ok := nm[a]
	return ok
}
