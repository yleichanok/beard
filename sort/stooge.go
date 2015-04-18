package sort

// Stooge sort algorithm.
//
// Swaps the top and bottom items if needed, then (recursively) sorts the bottom two-thirds,
// then the top two-thirds, then the bottom two-thirds again.
//
// See http://en.wikipedia.org/wiki/Stooge_sort.
func StoogeSort(arr Sortable) {

	if arr == nil {
		panic("arr is nil")
	}

	arrLen := arr.Len()

	// Quit if nothing to sort.
	if arrLen < 2 {
		return
	}

	stoogeSort(arr, 0, arrLen-1)
	return
}

func stoogeSort(arr Sortable, i, j int) {
	if arr.Compare(i, j) == 1 {
		arr.Swap(i, j)
	}

	if j-i+1 > 2 {
		t := (j - i + 1) / 3
		stoogeSort(arr, i, j-t)
		stoogeSort(arr, i+t, j)
		stoogeSort(arr, i, j-t)
	}
	return
}
