package linkedlist

import (
	"testing"
)

func TestNodeStoresData(t *testing.T) {
	data := "nodeContent"
	node := Node{data: data}
	if node.data != data {
		t.Errorf("Expected %s got %s", data, node.data)
	}
}

func TestNodeStoresNext(t *testing.T) {
	node1 := Node{data: "node1"}
	node2 := Node{data: "node2", next: &node1}
	if node2.next.data != node1.data {
		t.Errorf("Expected node1.Next is equals node2")
	}
}

func TestLinkedListPush(t *testing.T) {
	linkedlist := new(Linkedlist)

	linkedlist.Push(1)
	if linkedlist.head.data != 1 {
		t.Errorf("Expected %d got %d", 1, linkedlist.head.data)
	}

	linkedlist.Push(2)
	if linkedlist.head.data != 2 {
		t.Errorf("Expected %d got %d", 2, linkedlist.head.data)
	}
}

func TestLinkedListPop(t *testing.T) {
	linkedlist := new(Linkedlist)

	linkedlist.Push(1)
	linkedlist.Push(2)

	popped := linkedlist.Pop()

	// checks the popped value
	if popped != 2 {
		t.Errorf("Expected %d got %d", 2, popped)
	}

	// checks the head node
	if linkedlist.head.data != 1 {
		t.Errorf("Expected %d got %d", 1, linkedlist.head.data)
	}
}

func TestLinkedListLen(t *testing.T) {
	linkedlist := new(Linkedlist)

	linkedlist.Push(1)
	if linkedlist.Len() != 1 {
		t.Errorf("Expected %d got %d", 1, linkedlist.Len())
	}

	linkedlist.Push(2)
	if linkedlist.Len() != 2 {
		t.Errorf("Expected %d got %d", 2, linkedlist.Len())
	}

	linkedlist.Pop()
	if linkedlist.Len() != 1 {
		t.Errorf("Expected %d got %d", 1, linkedlist.Len())
	}
}

func TestLinkedListFirstNode(t *testing.T) {
	linkedlist := new(Linkedlist)

	linkedlist.Push(1)
	linkedlist.Push(2)
	linkedlist.Push(3)

	firstNode := linkedlist.FirstNode()

	if firstNode != 1 {
		t.Errorf("Expected %d got %d", 1, firstNode)
	}
}
