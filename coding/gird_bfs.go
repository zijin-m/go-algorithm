package coding

func GirdBFS(grid [][]int, i, j int) {
	moves := [][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	queue := [][2]int{{i, j}}
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			cur := queue[0]
			queue = queue[1:]
			x, y := cur[0], cur[1]
			for _, m := range moves {
				x2, y2 := x+m[0], y+m[1]
				if inArea(grid, x2, y2) && grid[x2][y2] == 0 {
					grid[x2][y2] = 2
					queue = append(queue, [2]int{x2, y2})
				}
			}
		}
	}
}

func inArea(grid [][]int, i, j int) bool {
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
}
