package huawei

// 老师给两个同学分糖果，每袋糖果中的数量不完全一样。一袋糖果只能分给一个人，并且一次性全分完必须。两个人分到的糖果数必须相同。返回两个人分到的糖果数，无法平均分配，返回-1。

// 输入
// 第一行输入糖果袋数，范围是[1,100]
// 第二行是一个数组，每袋糖果数量，范围[1,100]

// 输出
// 第一行为每人分到的糖果总数，不能平均分配，输出-1
// 第二行，第三行为两个同学分配到的每袋糖果具体的糖果数，顺序不限
// 存在多种分配方式时，输出任意一种可行的方法即可
// example
// 5
// 7 4 5 3 3
// out
// 11
// 7 4
// 5 3 3
func HandOutCandy(nums []int) (ava int, r1, r2 []int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		ava = -1
		return
	}
	target := sum / 2
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, target+1)
	}
	for j := nums[0]; j < target; j++ {
		dp[0][j] = nums[0]
	}
	for i := 1; i < n; i++ {
		v := nums[i]
		for j := 1; j <= target; j++ {
			if j < v {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-v]+v)
			}
		}
	}
	ava = target
	for target > 0 {
		for i := 0; i < n; i++ {
			if dp[i][target] == target {
				r1 = append(r1, nums[i])
				target -= nums[i]
				n = i
				break
			}
		}
	}
	r2 = sub(nums, r1)
	return
}

// sub return a - b
func sub(a, b []int) (c []int) {
	bm := map[int]bool{}
	for _, num := range b {
		bm[num] = true
	}
	for _, num := range a {
		if _, ok := bm[num]; !ok {
			c = append(c, num)
		}
	}
	return
}
