package coding

func CanPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	n := len(nums)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, target+1)
	}
	if nums[0] < target {
		dp[0][nums[0]] = true
	}
	for i := 1; i < n; i++ {
		v := nums[i]
		for j := 0; j <= target; j++ {
			if j < v {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
			}
		}
	}
	return dp[n-1][target]
}
