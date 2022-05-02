package coding

func GetAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	arr1, arr2 := inorder(root1), inorder(root2)
	i, j := 0, 0
	res := make([]int, len(arr1)+len(arr2))
	index := 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			res[index] = arr1[i]
			i++
		} else {
			res[index] = arr2[j]
			j++
		}
		index++
	}
	for i < len(arr1) {
		res[index] = arr1[i]
		index++
		i++
	}
	for j < len(arr2) {
		res[index] = arr2[j]
		index++
		j++
	}
	return res
}

func inorder(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	left := inorder(root.Left)
	right := inorder(root.Right)
	res = append(append(append(res, left...), root.Val), right...)
	return
}
