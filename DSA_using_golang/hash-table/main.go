package main

import "fmt"

const ArraySize = 7

// HashTablestruct
type HashTable struct {
	array [ArraySize]*bucket
}

// Insert
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)

}

// Search
func (h *HashTable) Search(key string) {
	index := hash(key)
	val := h.array[index].search(key)
	if val {
		fmt.Printf("Key %s is there in hash table\n", key)
	} else {
		fmt.Printf("Key %s not found in hash table\n", key)
	}
}

// Delete
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)

}

// bucket stuct
type bucket struct {
	head *bucketNode
}

// bucketNode struct
type bucketNode struct {
	key  string
	next *bucketNode
}

// insert will take in a key , create a node with the key and insert the node in the bucket
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "Key already exists")
	}
}

// delete
func (b *bucket) delete(k string) {

	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head

	for previousNode.next != nil {
		if previousNode.next.key == k {
			previousNode.next = previousNode.next.next
			return
		}
		previousNode = previousNode.next
	}
}

// search
func (b *bucket) search(k string) bool {
	currentNode := b.head

	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// hash
func hash(key string) int {
	sum := 0
	for v := range key {
		sum += int(v)
	}

	return sum % ArraySize
}

// Init will create bucket in each slot of hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	testHashTable := Init()
	fmt.Println("Key sachin Insert")
	testHashTable.Insert("sachin")
	fmt.Println("Key RAJU Insert")
	testHashTable.Insert("RAJU")
	fmt.Println("Key chandu Insert")
	testHashTable.Insert("chandu")
	testHashTable.Search("sachin")
	testHashTable.Search("vishal")
	testHashTable.Search("RAJU")
	testHashTable.Delete("RAJU")
	testHashTable.Search("RAJU")
}
