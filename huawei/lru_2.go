package huawei

type ILinkList interface {
	Get(index int) ListElement
	AddAtHead(val ListElement)
	AddAtTail(val ListElement)
	AddAtIndex(index int, val ListElement)
	DeleteAtIndex(index int)
}

type LRU struct {
	capacity int
	size     int
	m        map[int]*ListElement
	list     ILinkList
}

func NewLRU(capacity int) *LRU {
	return &LRU{
		m:        map[int]*ListElement{},
		list:     NewLinkList(),
		capacity: capacity,
	}
}

func (c *LRU) Get(key int) int {

}

func (c *LRU) Set(key int, val int) {

}
