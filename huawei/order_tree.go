package huawei

func threeOrders(root *TreeNode) [][]int {
	// write code here
	return [][]int{
		preorder(root),
		midorder(root),
		postorder(root),
	}
}

func preorder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{root.Val}
	left := preorder(root.Left)
	right := preorder(root.Right)
	res = append(append(res, left...), right...)
	return res
}

func midorder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	left := midorder(root.Left)
	right := midorder(root.Right)
	res = append(append(append(res, left...), root.Val), right...)
	return res
}

func postorder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	left := postorder(root.Left)
	right := postorder(root.Right)
	res = append(append(append(res, left...), right...), root.Val)
	return res
}
