package hashmap

type LRUCache struct {
	capacity   int
	size       int
	Caches     map[int]*LinkedNode
	head, tail *LinkedNode
}

type LinkedNode struct {
	key, value int
	prev, next *LinkedNode
}

func Constructor(capacity int) LRUCache {
	res := LRUCache{
		capacity: capacity,
		size:     0,
		Caches:   make(map[int]*LinkedNode),
		head:     &LinkedNode{0, 0, nil, nil},
		tail:     &LinkedNode{0, 0, nil, nil},
	}
	res.head.next = res.tail
	res.tail.prev = res.head
	return res
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Caches[key]; ok {
		this.moveToHead(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// 如果没有
	if _, ok := this.Caches[key]; !ok {
		node := &LinkedNode{
			key: key, value: value,
		}
		this.Caches[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			this.size--
			delete(this.Caches, removed.key)
		}
	} else {
		// 如果已经有这个key了
		node := this.Caches[key]
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) addToHead(node *LinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *LinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *LinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *LinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
