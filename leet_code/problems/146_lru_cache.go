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

func (cache *LRUCache) Get(key int) int {
	node, ok := cache.LinkedNodeMap[key]
	if !ok {
		return -1
	}
	cache.LinkedList.MoveToTop(node)
	return node.Value.(CacheNode).Value
}

func (cache *LRUCache) Put(key int, value int) {
	newCacheNode := CacheNode{key, value}
	node, ok := cache.LinkedNodeMap[key]

	if ok {
		node.Value = newCacheNode
		cache.LinkedList.MoveToTop(node)
		return
	}

	if len(cache.LinkedNodeMap) == cache.capacity {
		delete(cache.LinkedNodeMap, cache.LinkedList.Tail.Pre.Value.(CacheNode).Key)
		cache.LinkedList.Remove(cache.LinkedList.Tail.Pre)
	}

	newLinkedNode := &DoubleLinkedNode{Value: newCacheNode}
	cache.LinkedNodeMap[key] = newLinkedNode
	cache.LinkedList.AddToTop(newLinkedNode)
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
