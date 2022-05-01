package coding

// 给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，
// 找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
//
// 输入：candidates = [2,3,6,7], target = 7
//
// 输出：[[2,2,3],[7]]
//
// 解释：
//
// 2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
//
// 7 也是一个候选， 7 = 7 。
//
// 仅有这两种组合
func CombinationSum(candidates []int, target int) (res [][]int) {
	if len(candidates) < 1 {
		return
	}
	path := []int{}
	combinationSum(candidates, target, 0, path, &res)
	return
}

func combinationSum(candidates []int, target int, startIndex int, path []int, res *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}
	for i := startIndex; i < len(candidates); i++ {
		path = append(path, candidates[i])
		combinationSum(candidates, target-candidates[i], i, path, res)
		path = path[:len(path)-1]
	}
}
