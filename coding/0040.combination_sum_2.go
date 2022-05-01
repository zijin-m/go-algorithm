package coding

import "sort"

func CombinationSum2(candidates []int, target int) (res [][]int) {
	if len(candidates) < 1 {
		return
	}
	path := []int{}
	sort.Ints(candidates)
	combinationSum2(candidates, target, 0, path, &res)
	return
}

func combinationSum2(candidates []int, target int, startIndex int, path []int, res *[][]int) {
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
		if i > startIndex && candidates[i] == candidates[i-1] {
			continue
		}
		path = append(path, candidates[i])
		combinationSum2(candidates, target-candidates[i], i+1, path, res)
		path = path[:len(path)-1]
	}
}
