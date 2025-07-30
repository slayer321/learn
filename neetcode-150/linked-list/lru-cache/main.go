package main

import "fmt"

type Node struct {
	prev *Node
	next *Node
	key  int
	val  int
}

type LRUCache struct {
	capacity int
	hashMap  map[int]*Node
	head     *Node
	tail     *Node
}

func Constructor(capacity int) LRUCache {
	var head, tail *Node
	tail = &Node{
		key: -1,
		val: -1,
	}

	head = &Node{
		key: -1,
		val: -1,
	}
	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		hashMap:  make(map[int]*Node),
		head:     head,
		tail:     tail,
	}

}

func (this *LRUCache) Get(key int) int {

	if val, ok := this.hashMap[key]; ok {

		this.remove(val)
		this.addNewNodeAtHead(val)

		// currHeadNext := this.head.next
		// this.head.next = val
		// val.prev = this.head
		// val.next = currHeadNext

		// tailPrev := this.tail.prev
		// tailsPrevPrev := tailPrev.prev
		// tailsPrevPrev.next = this.tail

		return val.val
	}

	return -1

}

func (this *LRUCache) addNewNodeAtHead(newNode *Node) {

	headsNext := this.head.next
	headsNext.prev = newNode
	newNode.next = this.head.next
	this.head.next = newNode
	newNode.prev = this.head

}

func (this *LRUCache) remove(node *Node) {
	prev := node.prev
	next := node.next

	prev.next = next
	next.prev = prev
}

func (this *LRUCache) deleteTailNode() {
	currPrev := this.tail.prev

	if currPrev == this.head {
		return
	}

	this.remove(currPrev)

	delete(this.hashMap, currPrev.key)
}

func (this *LRUCache) Put(key int, value int) {
	newNode := &Node{
		key: key,
		val: value,
	}

	_, ok := this.hashMap[key]
	if len(this.hashMap) < this.capacity && !ok {

		this.addNewNodeAtHead(newNode)
		this.hashMap[newNode.key] = newNode

	} else if len(this.hashMap) < this.capacity && ok {

		currVal := this.hashMap[key]

		currVal.val = value

		this.remove(currVal)
		this.addNewNodeAtHead(currVal)

	} else if len(this.hashMap) == this.capacity {

		if !ok {
			this.deleteTailNode()

			this.addNewNodeAtHead(newNode)
			this.hashMap[newNode.key] = newNode

		} else if ok {
			currVal := this.hashMap[key]

			currVal.val = value
			this.remove(currVal)
			this.addNewNodeAtHead(currVal)
		}
	}

}
func main() {

	obj := Constructor(2)

	obj.Put(1, 1)
	obj.Put(2, 2)

	fmt.Printf("obj.Get(1): %v\n", obj.Get(1))
	obj.Put(3, 3)
	fmt.Printf("obj.Get(2): %v\n", obj.Get(2))

}
