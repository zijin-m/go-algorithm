package huawei

type ILinkList interface {
	Get(index int) *DoubleListNode
	AddAtHead(val *DoubleListNode)
	AddAtTail(val *DoubleListNode)
	AddAtIndex(index int, val *DoubleListNode)
	DeleteAtIndex(index int)
	Delete(val *DoubleListNode)
	MoveToBack(val *DoubleListNode)
}

type LRU struct {
	capacity int
	size     int
	m        map[int]*DoubleListNode
	l        ILinkList
}

func NewLRU(capacity int) *LRU {
	return &LRU{
		capacity: capacity,
		m:        make(map[int]*DoubleListNode),
		l:        NewLinkList(),
	}
}

func (c *LRU) Get(key int) int {
	val, ok := c.m[key]
	if !ok {
		return -1
	}
	// move to tail
	c.l.MoveToBack(val)
	return val.Val
}

func (c *LRU) Set(key int, val int) {
	if v := c.Get(key); v != -1 {
		// update
		c.m[key].Val = val
		return
	}
	if c.size >= c.capacity {
		c.removeHead()
	}
	// insert
	node := &DoubleListNode{Key: key, Val: val}
	c.m[key] = node
	c.l.AddAtTail(node)
	c.size++
}

func (c *LRU) removeHead() {
	node := c.l.Get(0)
	if node == nil {
		return
	}
	c.l.Delete(node)
	delete(c.m, node.Key)
	c.size--
}
