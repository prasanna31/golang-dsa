package heap

// MinHeap :
type MinHeap struct {
	heap []int
}

// NewMinHeap :
func NewMinHeap() *MinHeap {
	return &MinHeap{
		heap: make([]int, 0),
	}
}

// Insert :
func (h *MinHeap) Insert(value int) {
	h.heap = append(h.heap, value)
	h.heapifyUp()
}

// ExtractMin :
func (h *MinHeap) ExtractMin() int {
	if len(h.heap) == 0 {
		return -1
	}
	min := h.heap[0]
	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	h.heapifyDown()
	return min
}

func (h *MinHeap) heapifyUp() {
	index := len(h.heap) - 1
	for h.hasParent(index) && h.parent(index) > h.heap[index] {
		h.swap(h.parentIndex(index), index)
		index = h.parentIndex(index)
	}
}

func (h *MinHeap) heapifyDown() {
	index := 0
	for h.hasLeftChild(index) {
		smallerChildIndex := h.leftChildIndex(index)
		if h.hasRightChild(index) && h.rightChild(index) < h.leftChild(index) {
			smallerChildIndex = h.rightChildIndex(index)
		}
		if h.heap[index] < h.heap[smallerChildIndex] {
			break
		}
		h.swap(index, smallerChildIndex)
		index = smallerChildIndex
	}
}

func (h *MinHeap) hasParent(index int) bool {
	return index > 0
}

func (h *MinHeap) hasLeftChild(index int) bool {
	return h.leftChildIndex(index) < len(h.heap)
}

func (h *MinHeap) hasRightChild(index int) bool {
	return h.rightChildIndex(index) < len(h.heap)
}

func (h *MinHeap) parentIndex(index int) int {
	return (index - 1) / 2
}

func (h *MinHeap) leftChildIndex(index int) int {

	return 2*index + 1
}

func (h *MinHeap) rightChildIndex(index int) int {
	return 2*index + 2
}

func (h *MinHeap) parent(index int) int {
	return h.heap[h.parentIndex(index)]
}

func (h *MinHeap) leftChild(index int) int {
	return h.heap[h.leftChildIndex(index)]
}

func (h *MinHeap) rightChild(index int) int {
	return h.heap[h.rightChildIndex(index)]
}

func (h *MinHeap) swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}
