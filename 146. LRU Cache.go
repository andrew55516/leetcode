package main

//type LRUCache struct {
//	capacity int
//	cache    map[int]int
//	used     map[int]int
//	i        int
//}
//
//func Constructor146(capacity int) LRUCache {
//	return LRUCache{
//		capacity: capacity,
//		cache:    make(map[int]int, capacity),
//		used:     make(map[int]int, capacity),
//	}
//}
//
//func (this *LRUCache) Get(key int) int {
//	if _, ok := this.cache[key]; !ok {
//		return -1
//	}
//	this.used[key] = this.i
//	this.i++
//	return this.cache[key]
//}
//
//func (this *LRUCache) Put(key int, value int) {
//	if _, ok := this.cache[key]; !ok {
//		if this.capacity == len(this.cache) {
//			min := -1
//			for k, v := range this.used {
//				if min == -1 || v < this.used[min] {
//					min = k
//				}
//			}
//			delete(this.used, min)
//			delete(this.cache, min)
//		}
//		this.cache[key] = value
//	}
//	this.cache[key] = value
//	this.used[key] = this.i
//	this.i++
//
//}

type Node146 struct {
	Key   int
	Value int
	Next  *Node146
	Prev  *Node146
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node146
	head     *Node146
	tail     *Node146
}

func Constructor146(capacity int) LRUCache {
	cache := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node146, capacity),
		head:     &Node146{-1, -1, nil, nil},
		tail:     &Node146{-1, -1, nil, nil},
	}

	cache.head.Next = cache.tail
	cache.tail.Prev = cache.head

	return cache
}

func (this *LRUCache) DeleteNode(node *Node146) {
	prev := node.Prev
	next := node.Next
	prev.Next = next
	next.Prev = prev
}

func (this *LRUCache) AddNode(node *Node146) {
	next := this.head.Next
	this.head.Next = node
	node.Next = next
	node.Prev = this.head
	next.Prev = node
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	this.DeleteNode(this.cache[key])
	this.AddNode(this.cache[key])
	return this.cache[key].Value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; ok {
		this.cache[key].Value = value
		this.DeleteNode(this.cache[key])
		this.AddNode(this.cache[key])
		return
	}

	if this.capacity == len(this.cache) {
		d := this.tail.Prev
		delete(this.cache, d.Key)
		this.DeleteNode(d)
	}

	this.cache[key] = &Node146{Key: key, Value: value}
	this.AddNode(this.cache[key])
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
