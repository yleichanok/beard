package sort

// Shell sorting algorithm.
//
// The first diminishing increment sort.
// On each pass i sets of n/i items are sorted, typically with insertion sort.
// On each succeeding pass, i is reduced until it is 1 for the last pass.
// A good series of i values is important to efficiency.
//
// See http://en.wikipedia.org/wiki/Shellsort.
func ShellSort(arr Sortable) {

	if arr == nil {
		panic("arr is nil")
	}

	arrLen := arr.Len()

	// Quit if nothing to sort.
	if arrLen < 2 {
		return
	}

	gaps := []int{701, 301, 132, 57, 23, 10, 4, 1}

	for k := 0; k < len(gaps); k++ {
		gap := gaps[k]
		for i := gap; i < arrLen; i++ {
			j := i - gap
			for j >= 0 && arr.Compare(j, j+gap) == 1 {
				arr.Swap(j, j+gap)
				j -= gap
			}
		}
	}
	return
}
