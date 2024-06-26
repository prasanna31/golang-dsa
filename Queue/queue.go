package queue

// Queue represents a queue data structure
type Queue struct {
	elements []any
}

// InitQueue initializes the queue
func InitQueue() *Queue {
	return &Queue{}
}

// Enqueue adds an element to the end of the queue
func (q *Queue) Enqueue(value any) {
	q.elements = append(q.elements, value)
}

// Dequeue removes and returns the front element from the queue
func (q *Queue) Dequeue() any {
	if q.IsEmpty() {
		return nil
	}

	front := q.elements[0]
	q.elements = q.elements[1:]

	return front
}

// PeekQueue returns the front element of the queue without removing it
func (q *Queue) PeekQueue() any {
	if q.IsEmpty() {
		return nil
	}

	return q.elements[0]
}

// IsEmptyQueue checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

// Display displays the queue
func (q *Queue) Display() {
	for _, element := range q.elements {
		print(element, " ")
	}
	println()
}

// Length returns the number of elements in the queue
func (q *Queue) Length() int {
	return len(q.elements)
}

// Clear removes all elements from the queue
func (q *Queue) Clear() {
	q.elements = []any{}
}
