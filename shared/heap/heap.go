package heap

// _comparable interface
type _comparable[T any] interface {
	IsLess(v T) bool
}

// Heap implements a min heap
type Heap[T _comparable[T]] struct {
	capacity int
	elements []T
	length   int
}

// New returns a new heap of the given capacity
func New[T _comparable[T]](capacity int) *Heap[T] {
	return &Heap[T]{
		capacity: capacity,
		length:   0,
		elements: make([]T, capacity),
	}
}

// Elements returns all the elements in the heap
func (h *Heap[T]) Elements() []T {
	return h.elements[:h.length]
}

// parentIndex returns the index of the parent for the given index
func (h *Heap[T]) parentIndex(curIndex int) int {
	return (curIndex - 1) / 2
}

// leftIndex returns the index of the left child for the given index
func (h *Heap[T]) leftIndex(curIndex int) int {
	return (2 * curIndex) + 1
}

// rightIndex returns the index of the right child for the given index
func (h *Heap[T]) rightIndex(curIndex int) int {
	return (2 * curIndex) + 2
}

// Swap the values in the given indices
func (h *Heap[T]) Swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

// Len returns the current length of heap
func (h *Heap[T]) Len() int {
	return h.length
}

// IsFull returns true if the heap is full
func (h *Heap[T]) IsFull() bool {
	return h.length == h.capacity
}

// IsEmpty returns true if the heap is empty
func (h *Heap[T]) IsEmpty() bool {
	return h.length == 0
}

// Peek returns the root element in the heap
func (h *Heap[T]) Peek() (T, bool) {
	var v T
	if h.IsEmpty() {
		return v, false
	}
	return h.elements[0], true
}

// Push adds an element to the heap
func (h *Heap[T]) Push(element T) bool {
	if h.IsFull() {
		return false
	}

	i := h.length
	h.elements[i] = element
	h.length++

	for i != 0 && h.elements[i].IsLess(h.elements[h.parentIndex(i)]) {
		h.Swap(i, h.parentIndex(i))
		i = h.parentIndex(i)
	}
	return true
}

// Pop removes the root element and returns it; if the heap is empty returns false
func (h *Heap[T]) Pop() (T, bool) {
	var v T
	if h.IsEmpty() {
		return v, false
	}
	root := h.elements[0]
	h.elements[0] = h.elements[h.length-1]
	h.length--
	h.fix(0)
	return root, true
}

// fix the  heap in case the order of elements needs to be fixed
func (h *Heap[T]) fix(index int) {
	largest := index
	left := h.leftIndex(index)
	right := h.rightIndex(index)

	if left < h.length && h.elements[left].IsLess(h.elements[largest]) {
		largest = left
	}
	if right < h.length && h.elements[right].IsLess(h.elements[largest]) {
		largest = right
	}

	if largest != index {
		h.Swap(index, largest)
		h.fix(largest)
	}
}
