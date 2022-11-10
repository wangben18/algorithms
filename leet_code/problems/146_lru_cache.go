package problems

type LRUCache struct {
	LinkedList    *DoubleLinkedList
	LinkedNodeMap map[int]*DoubleLinkedNode
	capacity      int
}

func Constructor(capacity int) LRUCache {
	if capacity == 0 {
		panic("capacity should not be zero")
	}
	return LRUCache{
		LinkedList:    ConstructDoubleLinkedList(),
		LinkedNodeMap: map[int]*DoubleLinkedNode{},
		capacity:      capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	// get from map
	node, ok := this.LinkedNodeMap[key]
	// if not exist, return -1
	if !ok {
		return -1
	}
	// if exist, move to top
	this.LinkedList.MoveToTop(node)
	return node.Value.(CacheNode).Value
}

func (this *LRUCache) Put(key int, value int) {
	newCacheNode := CacheNode{key, value}
	// get from map
	node, ok := this.LinkedNodeMap[key]
	// if exist, update value, move to top
	if ok {
		node.Value = newCacheNode
		this.LinkedList.MoveToTop(node)
		return
	}
	// if not exist,
	// if len(map) == capacity, delete last node
	if len(this.LinkedNodeMap) == this.capacity {
		delete(this.LinkedNodeMap, this.LinkedList.Tail.Pre.Value.(CacheNode).Key)
		this.LinkedList.Remove(this.LinkedList.Tail.Pre)
	}
	// add new node
	newLinkedNode := &DoubleLinkedNode{Value: newCacheNode}
	this.LinkedNodeMap[key] = newLinkedNode
	this.LinkedList.AddToTop(newLinkedNode)
}

type DoubleLinkedList struct {
	Head *DoubleLinkedNode
	Tail *DoubleLinkedNode
}

func ConstructDoubleLinkedList() *DoubleLinkedList {
	tail := &DoubleLinkedNode{}
	head := &DoubleLinkedNode{}
	tail.Pre = head
	head.Next = tail
	return &DoubleLinkedList{
		Head: head,
		Tail: tail,
	}
}

func (list *DoubleLinkedList) MoveToTop(node *DoubleLinkedNode) {
	list.Remove(node)
	list.AddToTop(node)
}

func (list *DoubleLinkedList) Remove(node *DoubleLinkedNode) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (list *DoubleLinkedList) AddToTop(node *DoubleLinkedNode) {
	list.Head.Next.Pre = node
	node.Next = list.Head.Next
	node.Pre = list.Head
	list.Head.Next = node
}

type CacheNode struct {
	Key   int
	Value int
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
