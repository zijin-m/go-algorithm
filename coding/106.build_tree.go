package coding

func BuildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	// 先找到根节点（后续遍历的最后一个就是根节点
	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootVal}
	if len(postorder) == 1 {
		return root
	}

	// 从中序遍历中找到一分为二的点，左边为左子树，右边为右子树
	idx := findIndex(inorder, rootVal)

	// 将后续遍历依据中序大小一分为二，左边为左子树，右边为右子树
	root.Left = BuildTree(inorder[:idx], postorder[:idx])
	root.Right = BuildTree(inorder[idx+1:], postorder[idx:len(postorder)-1])

	// 返回根节点
	return root
}

func findIndex(inorder []int, node int) int {
	for i, val := range inorder {
		if val == node {
			return i
		}
	}
	return -1
}
