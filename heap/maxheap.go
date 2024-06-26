package heap

// MaxHeap :
type MaxHeap struct {
	heap []int
}

// NewMaxHeap :
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		heap: make([]any, 0),
	}
}

// Insert :
func (h *MaxHeap) Insert(value int) {
	h.heap = append(h.heap, value)
	h.heapifyUp()
}

// ExtractMax :
func (h *MaxHeap) ExtractMax() int {
	if len(h.heap) == 0 {
		return -1
	}
	max := h.heap[0]
	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	h.heapifyDown()
	return max
}

func (h *MaxHeap) heapifyUp() {
	index := len(h.heap) - 1
	for h.hasParent(index) && h.parent(index) < h.heap[index] {
		h.swap(h.parentIndex(index), index)
		index = h.parentIndex(index)
	}
}

func (h *MaxHeap) heapifyDown() {
	index := 0
	for h.hasLeftChild(index) {
		biggerChildIndex := h.leftChildIndex(index)
		if h.hasRightChild(index) && h.rightChild(index) > h.leftChild(index) {
			biggerChildIndex = h.rightChildIndex(index)
		}
		if h.heap[index] > h.heap[biggerChildIndex] {
			break
		}
		h.swap(index, biggerChildIndex)
		index = biggerChildIndex
	}
}

func (h *MaxHeap) hasParent(index int) bool {
	return index > 0
}

func (h *MaxHeap) hasLeftChild(index int) bool {
	return h.leftChildIndex(index) < len(h.heap)
}

func (h *MaxHeap) hasRightChild(index int) bool {
	return h.rightChildIndex(index) < len(h.heap)
}

func (h *MaxHeap) parentIndex(index int) int {
	return (index - 1) / 2
}

func (h *MaxHeap) leftChildIndex(index int) int {
	return 2*index + 1
}

func (h *MaxHeap) rightChildIndex(index int) int {
	return 2*index + 2
}

func (h *MaxHeap) parent(index int) int {
	return h.heap[h.parentIndex(index)]
}

func (h *MaxHeap) leftChild(index int) int {
	return h.heap[h.leftChildIndex(index)]
}

func (h *MaxHeap) rightChild(index int) int {
	return h.heap[h.rightChildIndex(index)]
}

func (h *MaxHeap) swap(index1, index2 int) {
	h.heap[index1], h.heap[index2] = h.heap[index2], h.heap[index1]
}

// Size :
func (h *MaxHeap) Size() int {
	return len(h.heap)
}

// Peek :
func (h *MaxHeap) Peek() int {
	if len(h.heap) == 0 {
		return -1
	}
	return h.heap[0]
}

// IsEmpty :
func (h *MaxHeap) IsEmpty() bool {
	return len(h.heap) == 0
}

// Clear :
func (h *MaxHeap) Clear() {
	h.heap = make([]int, 0)
}

// Values :
func (h *MaxHeap) Values() []int {
	return h.heap
}

// HeapSort :
func HeapSort(arr []int) []int {
	h := NewMaxHeap()
	for _, v := range arr {
		h.Insert(v)
	}
	sorted := make([]int, 0)
	for !h.IsEmpty() {
		sorted = append(sorted, h.ExtractMax())
	}
	return sorted
}

// Heapify :
func Heapify(arr []int) []int {
	h := NewMaxHeap()
	h.heap = arr
	for i := len(arr)/2 - 1; i >= 0; i-- {
		h.heapifyDown()
	}
	return h.heap
}

// HeapifyUp :
func HeapifyUp(arr []int) []int {
	h := NewMaxHeap()
	h.heap = arr
	for i := 0; i < len(arr); i++ {
		h.heapifyUp()
	}
	return h.heap
}

// HeapifyDown :
func HeapifyDown(arr []int) []int {
	h := NewMaxHeap()
	h.heap = arr
	for i := len(arr) - 1; i >= 0; i-- {
		h.heapifyDown()
	}
	return h.heap
}

// BuildHeap :
func BuildHeap(arr []int) []int {
	h := NewMaxHeap()
	for _, v := range arr {
		h.Insert(v)
	}
	return h.heap
}

// BuildHeapInPlace :
func BuildHeapInPlace(arr []int) []int {
	h := NewMaxHeap()
	h.heap = arr
	for i := len(arr)/2 - 1; i >= 0; i-- {
		h.heapifyDown()
	}
	return h.heap
}

// BuildHeapInPlaceUp :
func BuildHeapInPlaceUp(arr []int) []int {
	h := NewMaxHeap()
	h.heap = arr
	for i := 0; i < len(arr); i++ {
		h.heapifyUp()
	}
	return h.heap
}

// BuildHeapInPlaceDown :
func BuildHeapInPlaceDown(arr []int) []int {
	h := NewMaxHeap()
	h.heap = arr
	for i := len(arr) - 1; i >= 0; i-- {
		h.heapifyDown()
	}
	return h.heap
}

// HeapSortInPlace :
func HeapSortInPlace(arr []int) []int {
	h := NewMaxHeap()
	h.heap = arr
	for i := len(arr) - 1; i >= 0; i-- {
		h.heapifyDown()
	}
	sorted := make([]int, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		h.swap(0, i)
		sorted = append(sorted, h.heap[i])
		h.heap = h.heap[:len(h.heap)-1]
		h.heapifyDown()
	}
	return sorted
}

// HeapSortInPlaceUp :
func HeapSortInPlaceUp(arr []int) []int {
	h := NewMaxHeap()
	h.heap = arr
	for i := 0; i < len(arr); i++ {
		h.heapifyUp()
	}
	sorted := make([]int, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		h.swap(0, i)
		sorted = append(sorted, h.heap[i])
		h.heap = h.heap[:len(h.heap)-1]
		h.heapifyDown()
	}
	return sorted
}

// HeapSortInPlaceDown :
func HeapSortInPlaceDown(arr []int) []int {
	h := NewMaxHeap()
	h.heap = arr
	for i := len(arr) - 1; i >= 0; i-- {
		h.heapifyDown()
	}
	sorted := make([]int, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		h.swap(0, i)
		sorted = append(sorted, h.heap[i])
		h.heap = h.heap[:len(h.heap)-1]
		h.heapifyDown()
	}
	return sorted
}

// HeapifyInPlace :
func HeapifyInPlace(arr []int) []int {
	h := NewMaxHeap()
	h.heap = arr
	for i := len(arr)/2 - 1; i >= 0; i-- {
		h.heapifyDown()
	}
	return h.heap
}

// HeapifyInPlaceUp :
func HeapifyInPlaceUp(arr []int) []int {
	h := NewMaxHeap()
	h.heap = arr
	for i := 0; i < len(arr); i++ {
		h.heapifyUp()
	}
	return h.heap
}

// HeapifyInPlaceDown :
func HeapifyInPlaceDown(arr []int) []int {
	h := NewMaxHeap()
	h.heap = arr
	for i := len(arr) - 1; i >= 0; i-- {
		h.heapifyDown()
	}
	return h.heap
}
