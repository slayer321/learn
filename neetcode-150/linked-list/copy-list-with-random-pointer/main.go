package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	myMap := make(map[*Node]*Node)

	myMap[nil] = nil

	tmp := head

	for tmp != nil {

		myMap[tmp] = &Node{
			Val: tmp.Val,
		}

		tmp = tmp.Next
	}

	// fmt.Println(myMap)

	tmp1 := head
	for tmp1 != nil {
		copy := myMap[tmp1]
		copy.Next = myMap[tmp1.Next]
		copy.Random = myMap[tmp1.Random]

		tmp1 = tmp1.Next

	}

	return myMap[head]
}
