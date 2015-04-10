package sort

import "math"

// Heap sort algorithm.
//
// Builds a heap, then repeatedly extracts the maximum item. Run time is O(n log n).
//
// See http://en.wikipedia.org/wiki/Heapsort.
func HeapSort(arr Sortable) {

	if arr == nil {
		panic("arr is nil")
	}

	arrLen := arr.Len()

	// Quit if nothing to sort.
	if arrLen < 2 {
		return
	}

	heapify(arr, arrLen)

	end := arrLen - 1

	for end > 0 {
		arr.Swap(end, 0)
		end--
		siftDown(arr, 0, end)
	}
	return
}

// Builds the heap in array so that largest value is at the root.
func heapify(arr Sortable, arrLen int) {

	start := int(math.Floor(float64(arrLen-2) / 2))

	for start >= 0 {
		siftDown(arr, start, arrLen-1)
		start--
	}
	return
}

// Repairs the heap whose root element is at index 'start', assuming the heaps rooted at its children are valid.
func siftDown(arr Sortable, start, end int) {

	root := start

	for root*2+1 <= end {
		child := root*2 + 1
		swap := root

		if arr.Compare(swap, child) == -1 {
			swap = child
		}
		if child+1 <= end && arr.Compare(swap, child+1) == -1 {
			swap = child + 1
		}

		if swap == root {
			return
		} else {
			arr.Swap(root, swap)
			root = swap
		}
	}
	return
}
