package main

type Node struct {
	name string
}

type SingleLinkedList struct {
	count int
	next *Chain
}

type Chain struct {
	node *interface
	next *Chain
}

func main() {
	i := createNode("Test")
}

func createLinkedList(v *interface) (*SingleLinkedList) {
	nextChain:=Chain{
		node: v,	
		next: nil
	}
	return SingleLinkedList{
count:1,
next:&nextChain,
}
}

func createNode(s string) Node {
	return Node{
		name: s,
	}
}
