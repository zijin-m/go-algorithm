package sort

import (
	"math/rand"
	"time"
)

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, l, r int) {
	if l < r {
		rand.Seed(time.Now().Unix())
		swap(arr, l+rand.Intn(r-l), r)
		p := partition(arr, l, r)
		quickSort(arr, l, p[0]-1)
		quickSort(arr, p[1]+1, r)
	}
}

func partition(arr []int, l, r int) [2]int {
	less, more := l-1, r
	for l < more {
		if arr[l] < arr[r] {
			swap(arr, l, less+1)
			less++
			l++
		} else if arr[l] > arr[r] {
			swap(arr, l, more-1)
			more--
		} else {
			l++
		}
	}
	swap(arr, more, r)
	return [2]int{less + 1, more}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
