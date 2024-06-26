package LinkedList

// Node represents a node in the linked list
type Node struct {
	Value any
	Next  *Node
}

// LinkedList represents a linked list
type LinkedList struct {
	Head *Node
}

// initializeLinkedList initializes a linked list
func initializeLinkedList() *LinkedList {
	return &LinkedList{}
}

// InsertEnd inserts a value at the end of the linked list
func (ll *LinkedList) InsertEnd(value any) {
	newNode := &Node{Value: value}

	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
}

// InsertStart inserts a value at the start of the linked list
func (ll *LinkedList) InsertStart(value any) {
	newNode := &Node{Value: value}

	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	newNode.Next = ll.Head
	ll.Head = newNode
}

// InsertMiddle inserts a value at a specified index in the linked list
func (ll *LinkedList) InsertMiddle(index int, value any) {
	if index == 0 {
		ll.InsertStart(value)
		return
	}

	newNode := &Node{Value: value}
	current := ll.Head

	for i := 0; i < index-1; i++ {
		if current == nil {
			return
		}
		current = current.Next
	}

	newNode.Next = current.Next
	current.Next = newNode
}

// DeleteByIndex deletes a node at a specified index in the linked list
func (ll *LinkedList) DeleteByIndex(index int) {
	if ll.Head == nil {
		return
	}

	if index == 0 {
		ll.Head = ll.Head.Next
		return
	}

	current := ll.Head
	for i := 0; i < index-1; i++ {
		if current == nil {
			return
		}
		current = current.Next
	}

	if current.Next == nil {
		return
	}

	current.Next = current.Next.Next
}

// DeleteByValue deletes the first occurrence of a value in the linked list
func (ll *LinkedList) DeleteByValue(value any) {
	if ll.Head == nil {
		return
	}

	if ll.Head.Value == value {
		ll.Head = ll.Head.Next
		return
	}

	current := ll.Head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

// Get returns the value at a specified index in the linked list
func (ll *LinkedList) Get(index int) any {
	current := ll.Head
	for i := 0; i < index; i++ {
		if current == nil {
			return nil
		}
		current = current.Next
	}
	return current.Value
}

// Display displays the linked list
func (ll *LinkedList) Display() {
	current := ll.Head
	for current != nil {
		print(current.Value, " ")
		current = current.Next
	}
	println()
}

// DisplayReverse displays the linked list in reverse
func (ll *LinkedList) DisplayReverse() {
	var printReverse func(*Node)
	printReverse = func(node *Node) {
		if node == nil {
			return
		}
		printReverse(node.Next)
		print(node.Value, " ")
	}
	printReverse(ll.Head)
	println()
}

// DisplayAtIndex displays the value at a specified index in the linked list
func (ll *LinkedList) DisplayAtIndex(index int) {
	current := ll.Head
	for i := 0; i < index; i++ {
		if current == nil {
			return
		}
		current = current.Next
	}
	println(current.Value)
}

// Reverse reverses the linked list
func (ll *LinkedList) Reverse() {
	var prev, next *Node
	current := ll.Head

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	ll.Head = prev
}

// Comparator defines a function that compares two elements and returns true if a <= b
type Comparator func(a, b any) bool

// intComparator is a comparator function for int type elements
func intComparator(a, b any) bool {
	return a.(int) > b.(int)
}

// Sort sorts the linked list
func (ll *LinkedList) Sort(cmp Comparator) {
	if ll.Head == nil || ll.Head.Next == nil {
		return
	}

	sorted := false
	for !sorted {
		sorted = true
		current := ll.Head
		for current.Next != nil {
			if cmp(current.Value, current.Next.Value) {
				current.Value, current.Next.Value = current.Next.Value, current.Value
				sorted = false
			}
			current = current.Next
		}
	}
}

// fetches length of the linked list
func (ll *LinkedList) Length() int {
	count := 0
	current := ll.Head
	for current != nil {
		count++
		current = current.Next
	}
	return count
}
