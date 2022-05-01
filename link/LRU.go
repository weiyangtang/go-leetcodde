package link

type LRUCache struct {
	cacheMap       map[int]*DLinkedNode
	head, tail     *DLinkedNode
	capacity, size int
}

type DLinkedNode struct {
	key, val   int
	prev, next *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	cacheMap := make(map[int]*DLinkedNode)
	head := &DLinkedNode{}
	tail := &DLinkedNode{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		cacheMap: cacheMap,
		head:     head,
		tail:     tail,
		capacity: capacity,
		size:     0,
	}
}

func (this *LRUCache) Get(key int) int {
	val := -1
	if node, ok := this.cacheMap[key]; ok {
		val = node.val
		this.remove(node)
		this.addFirst(node)
	}
	return val
}

func (this *LRUCache) Put(key int, value int) {
	node := &DLinkedNode{key: key, val: value}
	if oldNode, ok := this.cacheMap[key]; ok {
		this.remove(oldNode)
		this.size--
	}
	this.cacheMap[key] = node
	this.addFirst(node)
	this.size++

	if this.size > this.capacity {
		delete(this.cacheMap, this.tail.prev.key)
		this.remove(this.tail.prev)
		this.size--
	}

}

func (this *LRUCache) remove(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) addFirst(node *DLinkedNode) {
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
	node.prev = this.head
}
