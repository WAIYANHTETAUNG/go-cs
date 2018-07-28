package stacklinkedlist

// Node for linked list
type Node struct {
	Data interface{}
	next *Node
}

// NewNode is creating node for linked list
func NewNode(data interface{}) *Node {
	return &Node{Data: data, next: nil}
}

// NewLinkedList is creating node for linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// LinkedList is struct of linked object
type LinkedList struct {
	head *Node
}

// Push is to push node to the list
func (linkedlist *LinkedList) Push(node *Node) {
	if linkedlist.head == nil {
		linkedlist.head = node
		return
	}

	temp := linkedlist.head
	for temp != nil {
		if temp.next != nil {
			temp = temp.next
		} else {
			lastNode := temp
			lastNode.next = node
			temp = nil
		}
	}
}

// Pop is to get node at the end of list
func (linkedlist *LinkedList) Pop() *Node {
	if linkedlist.head == nil {
		return nil
	}

	temp := linkedlist.head
	var lastNode Node

	for temp != nil {
		if temp.next != nil {
			lastNode := temp.next.next
			if lastNode == nil {
				temp.next = nil
			}
		}
		temp = temp.next
	}

	return &lastNode
}

// Size is to get number of node list in a linked
func (linkedlist *LinkedList) Size() int {
	length := 0
	temp := linkedlist.head

	for temp != nil {
		length++
		temp = temp.next
	}

	return length
}

// IsEmpty is to check whether stack is empty or not
func (linkedlist *LinkedList) IsEmpty() bool {
	return linkedlist.Size() == 0
}
