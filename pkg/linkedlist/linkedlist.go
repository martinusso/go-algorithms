/*
Package linkedlist implements a linear data structure Linked List

A Linked List is a linear collection of data elements, in which linear order is not given by their physical placement in memory. Instead, each element points to the next. It is a data structure consisting of a group of nodes which together represent a sequence. Under the simplest form, each node is composed of data and a reference (in other words, a link) to the next node in the sequence. This structure allows for efficient insertion or removal of elements from any position in the sequence during iteration. More complex variants add additional links, allowing efficient insertion or removal from arbitrary element references. A drawback of linked lists is that access time is linear (and difficult to pipeline). Faster access, such as random access, is not feasible. Arrays have better cache locality as compared to linked lists.

 Reference: https://en.wikipedia.org/wiki/Linked_list
*/
package linkedlist

// Node from the Linkedlist
type Node struct {
	data interface{}
	next *Node
}

// Linkedlist contains the reference to the head element
type Linkedlist struct {
	head *Node
	size int
}

// Len returns the linkedlist´s lenght
func (list *Linkedlist) Len() int {
	return list.size
}

// Push a new element onto the Linkedlist
func (list *Linkedlist) Push(value interface{}) {
	list.head = &Node{value, list.head}
	list.size++
}

// Pop remove the head element from the Linkedlist and return it's value
func (list *Linkedlist) Pop() (data interface{}) {
	if list.head != nil {
		data, list.head = list.head.data, list.head.next
		list.size--
		return data
	}
	return nil
}

// FirstNode returns the first element from the Linkedlist
func (list *Linkedlist) FirstNode() interface{} {
	node := list.head
	for node.next != nil {
		node = node.next
	}
	return node.data
}
