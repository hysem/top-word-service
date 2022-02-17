package heap

type _comparable[T any] interface {
	IsLess(v T) bool
}

// Heap implements a min heap
type Heap[T _comparable[T]] struct {
	capacity int
	elements []T
	length   int
}

func New[T _comparable[T]](capacity int) *Heap[T] {
	return &Heap[T]{
		capacity: capacity,
		length:   0,
		elements: make([]T, capacity),
	}
}

func (h *Heap[T]) Elements() []T {
	return h.elements[:h.length]
}

func (h *Heap[T]) parentIndex(curIndex int) int {
	return (curIndex - 1) / 2
}

func (h *Heap[T]) leftIndex(curIndex int) int {
	return (2 * curIndex) + 1
}

func (h *Heap[T]) rightIndex(curIndex int) int {
	return (2 * curIndex) + 2
}

func (h *Heap[T]) Swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

func (h *Heap[T]) Len() int {
	return h.length
}

func (h *Heap[T]) IsFull() bool {
	return h.length == h.capacity
}
func (h *Heap[T]) IsEmpty() bool {
	return h.length == 0
}

func (h *Heap[T]) Peek() (T, bool) {
	var v T
	if h.IsEmpty() {
		return v, false
	}
	return h.elements[0], true
}

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
