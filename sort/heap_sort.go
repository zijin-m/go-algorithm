package sort

type Heap struct {
	arr  []int
	size int
}

func NewHeap(arr []int) *Heap {
	return &Heap{
		arr: arr,
	}
}

func (h *Heap) Sort() {
	if h.arr == nil || len(h.arr) < 2 {
		return
	}
	for i := range h.arr {
		h.insert(i)
		h.size++
	}

	for h.size > 0 {
		h.swap(0, h.size-1)
		h.size--
		h.heapify(0)
	}
}

func (h *Heap) insert(index int) {
	for h.arr[index] > h.arr[(index-1)/2] {
		h.swap(index, (index-1)/2)
		index = (index - 1) / 2
	}
}

func (h *Heap) heapify(index int) {
	left := 2*index + 1
	for left < h.size {
		right := left + 1
		var largest int
		if right < h.size && h.arr[right] > h.arr[left] {
			largest = right
		} else {
			largest = left
		}
		if h.arr[largest] < h.arr[index] {
			largest = index
		}
		if index == largest {
			break
		}
		h.swap(index, largest)
		index = largest
		left = 2*index + 1
	}
}

func (h Heap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}
