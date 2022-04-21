package coding

var res [][]int
var path []int

func Combine(n int, k int) [][]int {
	res = [][]int{}
	path = []int{}
	backtracing(n, k, 1)
	return res
}

func backtracing(n, k, start int) {
	if len(path)+n-start+1 < k {
		return
	}
	if len(path) == k {
		// ! 重新创建切片，避免使用同样的地址，造成值被覆盖
		temp := make([]int, k)
		copy(temp, path)
		res = append(res, temp)
		return
	}
	for i := start; i <= n; i++ {
		path = append(path, i)
		backtracing(n, k, i+1)
		path = path[:len(path)-1]
	}
}
