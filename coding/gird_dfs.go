package coding

func dfs(gird [][]int, r, c int) {
	if !inGird(gird, r, c) {
		return
	}
	if gird[r][c] != 1 {
		return
	}
	gird[r][c] = 2
	dfs(gird, r-1, c)
	dfs(gird, r+1, c)
	dfs(gird, r, c-1)
	dfs(gird, r, c+1)
}

func inGird(gird [][]int, r, c int) bool {
	rmax := len(gird[0])
	cmax := len(gird)
	return 0 < r && r < rmax && 0 < c && c < cmax
}
