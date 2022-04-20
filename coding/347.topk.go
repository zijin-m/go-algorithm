package coding

import "container/heap"

type MyHeap [][2]int

func (h MyHeap) Len() int { return len(h) }

func (h MyHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }

func (h MyHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MyHeap) Push(v interface{}) {
	*h = append(*h, v.([2]int))
}

func (h *MyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func TopKFrequent(nums []int, k int) []int {
	h := &MyHeap{}
	heap.Init(h)
	cm := map[int]int{}
	for _, v := range nums {
		cm[v]++
	}
	for key, value := range cm {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return res
}
