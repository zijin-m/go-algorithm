package huawei

func LevelOrder(root *TreeNode) [][]int {
	queue := []*TreeNode{root}
	res := [][]int{}
	for len(queue) > 0 {
		n := len(queue)
		level := []int{}
		for i := 0; i < n; i++ {
			cur := queue[0]
			queue = queue[1:]
			level = append(level, cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		res = append(res, level)
	}
	return res
}
