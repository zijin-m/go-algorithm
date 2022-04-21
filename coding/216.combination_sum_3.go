package coding

var result [][]int
var help []int

func CombinationSum3(k int, n int) [][]int {
	result = [][]int{}
	help = []int{}
	backtracking(k, n, 1)
	return result
}

func backtracking(k, n, start int) {
	if len(help)+9-start+1 < k {
		return
	}
	if len(help) == k {
		sum := 0
		for _, v := range help {
			sum += v
		}
		if sum == n {
			temp := make([]int, k)
			copy(temp, help)
			result = append(result, temp)
		}
		return
	}
	for i := start; i <= 9; i++ {
		help = append(help, i)
		backtracking(k, n, i+1)
		help = help[:len(help)-1]
	}
}
