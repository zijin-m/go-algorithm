package huawei

import "container/list"

type Solution struct {
	capacity int
	size     int
	m        map[int]*list.Element
	list     *list.List
}

type element struct {
	key, val int
}

func Constructor(capacity int) *Solution {
	return &Solution{
		m:        map[int]*list.Element{},
		list:     list.New(),
		capacity: capacity,
	}
}

func (s *Solution) Get(key int) int {
	e, ok := s.m[key]
	if !ok {
		return -1
	}
	s.list.MoveToBack(e)
	return e.Value.(*element).val
}

func (s *Solution) Set(key int, val int) {
	// key exists
	if n := s.Get(key); n != -1 {
		ptr := s.m[key].Value.(*element)
		ptr.val = val
		return
	}
	// key not exists
	if s.size >= s.capacity {
		s.removeHead()
	}
	// insert
	s.insert(key, val)
}

func (s *Solution) insert(key, val int) {
	e := &element{key: key, val: val}
	s.m[key] = s.list.PushBack(e)
	s.size++
}

func (s *Solution) removeHead() {
	head := s.list.Front()
	s.list.Remove(head)
	key := head.Value.(*element).key
	delete(s.m, key)
	s.size--
}
