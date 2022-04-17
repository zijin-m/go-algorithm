package huawei

func ThreeOrders(root *TreeNode) [][]int {
	// write code here
	return [][]int{
		mirrosPre(root),
		mirrosIn(root),
		mirrosPost(root),
	}
}

func mirrosPre(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	cur := root
	var mr *TreeNode
	res := []int{}
	for cur != nil {
		if cur.Left != nil {
			mr = cur.Left
			for mr.Right != nil && mr.Right != cur {
				mr = mr.Right
			}
			if mr.Right == nil {
				res = append(res, cur.Val)
				mr.Right = cur
				cur = cur.Left
			} else {
				mr.Right = nil
				cur = cur.Right
			}
		} else {
			res = append(res, cur.Val)
			cur = cur.Right
		}
	}
	return res
}

func mirrosIn(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	cur := root
	var mr *TreeNode
	res := []int{}
	for cur != nil {
		if cur.Left != nil {
			mr = cur.Left
			for mr.Right != nil && mr.Right != cur {
				mr = mr.Right
			}
			if mr.Right == nil {
				mr.Right = cur
				cur = cur.Left
			} else {
				res = append(res, cur.Val)
				mr.Right = nil
				cur = cur.Right
			}
		} else {
			res = append(res, cur.Val)
			cur = cur.Right
		}
	}
	return res
}

func mirrosPost(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	cur := root
	var mr *TreeNode
	res := []int{}
	for cur != nil {
		if cur.Left != nil {
			mr = cur.Left
			for mr.Right != nil && mr.Right != cur {
				mr = mr.Right
			}
			if mr.Right == nil {
				mr.Right = cur
				cur = cur.Left
			} else {
				mr.Right = nil
				// 遍历左子树右边界
				res = append(res, addEdge(cur.Left)...)
				cur = cur.Right
			}
		} else {
			cur = cur.Right
		}
	}
	// 遍历整树右边界
	res = append(res, addEdge(root)...)
	return res
}

func addEdge(root *TreeNode) []int {
	tail := reverseEdge(root)
	cur := tail
	res := []int{}
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Right
	}
	reverseEdge(tail)
	return res
}

func reverseEdge(root *TreeNode) *TreeNode {
	if root == nil || root.Right == nil {
		return root
	}
	newHead := reverseEdge(root.Right)
	root.Right.Right = root
	root.Right = nil
	return newHead
}
