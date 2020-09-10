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
