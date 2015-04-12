package sort

// Merge sort algorithm.
//
// Splits the items to be sorted into two groups, recursively sorts each group,
// and merges them into a final, sorted sequence. Run time is Θ(n log n).
//
// See http://en.wikipedia.org/wiki/Merge_sort.
func MergeSort(arr, tmp Sortable, copy func(from, to Sortable, fromIndex, toIndex int)) {

	if arr == nil {
		panic("arr is nil")
	}

	arrLen := arr.Len()

	// Quit if nothing to sort.
	if arrLen < 2 {
		return
	}

	lo := 0
	hi := arrLen - 1

	mergeSort(arr, tmp, lo, hi, copy)
	return
}

func mergeSort(arr, tmp Sortable, lo, hi int, copy func(from, to Sortable, fromIndex, toIndex int)) {

	if lo >= hi {
		return
	}

	mid := (lo + hi) / 2

	mergeSort(arr, tmp, lo, mid, copy)
	mergeSort(arr, tmp, mid+1, hi, copy)

	i0 := lo
	i1 := mid + 1
	for i := lo; i <= hi; i++ {
		if i0 <= mid && (i1 > hi || arr.Compare(i0, i1) == -1) {
			copy(arr, tmp, i0, i)
			i0++
		} else {
			copy(arr, tmp, i1, i)
			i1++
		}
	}

	// Copy array back.
	for i := lo; i <= hi; i++ {
		copy(tmp, arr, i, i)
	}
	return
}
