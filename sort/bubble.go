package sort

// Bubble (sinking) sort algorithm.
//
// Sorts by comparing each adjacent pair of items in a list in turn,
// swapping the items if necessary, and repeating the pass through the list
// until no swaps are done.
//
// See http://en.wikipedia.org/wiki/Bubble_sort.
func BubbleSort(arr Sortable) {

	if arr == nil {
		panic("arr is nil")
	}

	arrLen := arr.Len()

	// Quit if nothing to sort.
	if arrLen < 2 {
		return
	}

	swapped := true

	for swapped == true {
		swapped = false

		for i := 0; i < arrLen-1; i++ {
			if arr.Compare(i, i+1) == 1 {
				arr.Swap(i, i+1)
				swapped = true
			}
		}
	}

	return
}
