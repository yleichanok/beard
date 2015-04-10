package sort

// Insertion sort algotithm.
//
// Sorts by repeatedly taking the next item and inserting it into the final data structure
// in its proper order with respect to items already inserted.
//
// See http://en.wikipedia.org/wiki/Insertion_sort.
func InsertionSort(arr Sortable) {

	if arr == nil {
		panic("arr is nil")
	}

	arrLen := arr.Len()

	// Quit if nothing to sort.
	if arrLen < 2 {
		return
	}

	for i := 1; i < arrLen; i++ {
		j := i
		for j > 0 && arr.Compare(j-1, j) == 1 {
			arr.Swap(j-1, j)
			j--
		}
	}

	return
}
