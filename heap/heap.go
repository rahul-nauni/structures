package heap

import "fmt"

type MaxHeap struct {
	array []int
}

// Push the item onto the heap
func (h *MaxHeap) Push(item int) {
	h.array = append(h.array, item)
	h.HeapifyUp(len(h.array) - 1)
}

// Pop function returns the maximum element from heap
func (h *MaxHeap) Pop() int {
	if len(h.array) == 0 {
		fmt.Println("Cannot remove from empty array")
		return -1
	}
	maxVal := h.array[0]
	lastIndex := len(h.array) - 1
	// Put the last element at root
	h.array[0] = h.array[lastIndex]
	// Slice the array to get rid of last element
	h.array = h.array[:lastIndex]
	h.HeapifyDown(0)

	return maxVal
}

// HeapifyUp function heapifies the heap from bottom to top
func (h *MaxHeap) HeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// HeapifyDown function heapifies the heap from top to bottom
func (h *MaxHeap) HeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0
	// loop while index has at least one child
	for l <= lastIndex {
		// Left child is the only child
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}
		// compare current index with childToCompare
		// swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

// Get parent index
func parent(index int) int {
	return (index - 1) / 2
}

// Get left child index
func left(index int) int {
	return 2*index + 1
}

// Get right child index
func right(index int) int {
	return 2*index + 2
}

// Swap values in an array
func (h *MaxHeap) swap(low, high int) {
	h.array[low], h.array[high] = h.array[high], h.array[low]
}

//func main() {
//	heap := &MaxHeap{}
//
//	items := []int{1, 2, 3, 9, 4, 8}
//	for _, item := range items {
//		heap.Push(item)
//	}
//	fmt.Println(heap.array)
//
//	for i := 0; i < len(items); i++ {
//		fmt.Println(heap.Pop())
//	}
//}
