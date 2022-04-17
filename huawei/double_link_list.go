package huawei

type DoubleListNode struct {
	Key, Val int
	prev     *DoubleListNode
	next     *DoubleListNode
}

type MyLinkedList struct {
	size int
	head *DoubleListNode
	tail *DoubleListNode
}

func NewLinkList() *MyLinkedList {
	head := &DoubleListNode{}
	tail := &DoubleListNode{}
	head.next = tail
	tail.prev = head
	return &MyLinkedList{
		head: head,
		tail: tail,
	}
}

func (l *MyLinkedList) Get(index int) *DoubleListNode {
	if index < 0 || index >= l.size {
		return nil
	}
	cur := l.head
	if index+1 < l.size-index {
		for i := 0; i < index+1; i++ {
			cur = cur.next
		}
	} else {
		cur = l.tail
		for i := 0; i < l.size-index; i++ {
			cur = cur.prev
		}
	}
	return cur
}

func (l *MyLinkedList) AddAtHead(val *DoubleListNode) {
	l.AddAtIndex(0, val)
}

func (l *MyLinkedList) AddAtTail(val *DoubleListNode) {
	l.AddAtIndex(l.size, val)
}

func (l *MyLinkedList) AddAtIndex(index int, val *DoubleListNode) {
	if index > l.size {
		return
	}
	if index < 0 {
		index = 0
	}
	cur := l.head
	if index < l.size-index {
		for i := 0; i < index; i++ {
			cur = cur.next
		}
	} else {
		cur = l.tail
		for i := 0; i < l.size-index; i++ {
			cur = cur.prev
		}
	}
	prev := cur.prev
	val.prev = prev
	val.next = cur
	prev.next = val
	cur.prev = val
	l.size++
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	node := l.Get(index)
	if node == nil {
		return
	}
	l.Delete(node)
}

func (l *MyLinkedList) Delete(node *DoubleListNode) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
	l.size--
}

func (l *MyLinkedList) MoveToBack(node *DoubleListNode) {
	l.Delete(node)
	l.AddAtTail(node)
}
