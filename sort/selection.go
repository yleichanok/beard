package sort

/**
 * Selection sort algorithm.
 *
 * Repeatedly searches remaining items to find the least one
 * and moves it to its final location. The run time is Θ(n²),
 * where n is the number of elements. The number of swaps is O(n).
 *
 * See http://en.wikipedia.org/wiki/Selection_sort.
 */
func SelectionSort(arr Sortable) {

	if arr == nil {
		panic("arr is nil")
	}

	arrLen := arr.Len()

	// Quit if nothing to sort.
	if arrLen < 2 {
		return
	}

	for i := 0; i < arrLen-1; i++ {
		min := i
		for j := i + 1; j < arrLen; j++ {
			if arr.Less(j, min) {
				min = j
			}
		}
		if min != i {
			arr.Swap(i, min)
		}
	}

	return
}
