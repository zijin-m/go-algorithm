package sort

func MergeSort(arr []int) {
	mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, l, r int) {
	if l == r {
		return
	}
	mid := l + (r-l)>>1
	mergeSort(arr, l, mid)
	mergeSort(arr, mid+1, r)
	merge(arr, l, mid, r)
}

func merge(arr []int, l, m, r int) {
	help := make([]int, r-l+1)
	i := 0
	p1 := l
	p2 := m + 1
	for p1 <= m && p2 <= r {
		if arr[p1] <= arr[p2] {
			help[i] = arr[p1]
			p1++
			i++
		} else {
			help[i] = arr[p2]
			p2++
			i++
		}
	}
	for p1 <= m {
		help[i] = arr[p1]
		p1++
		i++
	}
	for p2 <= r {
		help[i] = arr[p2]
		p2++
		i++
	}
	for i, v := range help {
		arr[l+i] = v
	}
}
