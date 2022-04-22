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
	dp := make([]int, target+1)
	for i := 1; i < n; i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	return dp[target] == target
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
