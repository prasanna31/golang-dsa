package Interface

type golangDSA interface {
	// Array and LinkedList Operations
	InitializeArray()      // Initialize an array/linked list
	InitializeLinkedList() // Initialize a linked list
	InsertMiddle(int, any) // Insert at a given index in array/linked list
	InsertEnd(any)         // Insert at the end of array/linked list
	InsertStart(any)       // Insert at the start of array/linked list
	DeleteByIndex(int)     // Delete by index in array/linked list
	DeleteByValue(any)     // Delete by value in array/linked list
	Get(int)               // Search for a value in array/linked list
	DisplayReverse()       // Display the array/linked list in reverse
	DisplayAtIndex(int)    // Display the value at a given index in array/linked list
	Reverse()              // Reverse the array/linked list
	Sort()                 // Sort the array/linked list

	// Stack Operations
	InitStack()    // Initialize the stack
	Push(any)      // Push an element onto the stack
	Pop() any      // Pop an element from the stack
	Peek() any     // Peek at the top element of the stack
	IsEmpty() bool // Check if the stack is empty

	// Queue Operations
	InitQueue()         // Initialize the queue
	Enqueue(any)        // Enqueue an element into the queue
	Dequeue() any       // Dequeue an element from the queue
	PeekQueue() any     // Peek at the front element of the queue
	IsEmptyQueue() bool // Check if the queue is empty

	// BinaryTree Operations
	Insert(any)  // Insert a value into the binary tree
	Search(any)  // Search for a value in the binary tree
	Delete(any)  // Delete a value from the binary tree
	PreOrder()   // Perform a pre-order traversal of the binary tree
	InOrder()    // Perform an in-order traversal of the binary tree
	PostOrder()  // Perform a post-order traversal of the binary tree
	LevelOrder() // Perform a level-order traversal of the binary tree

	// Graph Operations
	BFS() // Perform a breadth-first search on the graph
	DFS() // Perform a depth-first search on the graph
	Dijikstra()
	BellmanFord()
	FloydWarshall()
	AdejacencyMatrix()
	AdejacencyList()
	IncidenceMatrix()
	IncidenceList()

	// Map Operations
	// Set Operations

	// array/linkedlist/stack
	Display()
	Length() int

	// stack and queue
	Clear()
}
