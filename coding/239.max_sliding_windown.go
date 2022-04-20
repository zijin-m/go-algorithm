package coding

type Queue interface {
	Push(v int)
	Pop(v int)
	Front() int
}

type queue struct {
	arr []int
}

func NewQueue() *queue {
	return &queue{}
}

func (q *queue) Push(v int) {
	for !q.Empty() && v > q.Back() {
		q.arr = q.arr[:len(q.arr)-1]
	}
	q.arr = append(q.arr, v)
}

func (q *queue) Pop(v int) {
	if !q.Empty() && v == q.arr[0] {
		q.arr = q.arr[1:]
	}
}

func (q *queue) Front() int {
	return q.arr[0]
}

func (q *queue) Back() int {
	return q.arr[len(q.arr)-1]
}

func (q *queue) Empty() bool {
	return len(q.arr) == 0
}

func MaxSlidingWindow(nums []int, k int) []int {
	q := NewQueue()
	for i := 0; i < k; i++ {
		q.Push(nums[i])
	}
	res := []int{q.Front()}
	for i := 0; i+k < len(nums); i++ {
		q.Pop(nums[i])
		q.Push(nums[i+k])
		res = append(res, q.Front())
	}
	return res
}
