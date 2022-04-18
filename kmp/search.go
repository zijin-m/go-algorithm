package kmp

func IndexOf(s, substr string) int {
	if len(substr) > len(s) {
		return -1
	}
	next := getNext(substr)
	i, j := 0, 0
	for i < len(s) && j < len(substr) {
		if j == -1 || s[i] == substr[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == len(substr) {
		return i - j
	}
	return -1

}

func getNext(substr string) []int {
	next := make([]int, len(substr))
	next[0] = -1
	k, j := -1, 0
	for j < len(substr)-1 {
		if k == -1 || substr[k] == substr[j] {
			k++
			j++
			next[j] = k
		} else {
			k = next[k]
		}
	}
	return next
}
