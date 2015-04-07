package sort

// Comb sort algorithm.
//
// An in-place sort algorithm that repeatedly reorders different pairs of items.
// On each pass swap pairs of items separated by the increment or gap,
// if needed, and reduce the gap (divide it by about 1.3).
// The gap starts at about 3/4 of the number of items. Continue until the gap
// is 1 and a pass had no swaps.
//
// See http://en.wikipedia.org/wiki/Comb_sort.
func CombSort(arr Sortable) {

	if arr == nil {
		panic("arr is nil")
	}

	arrLen := arr.Len()

	// Quit if nothing to sort.
	if arrLen < 2 {
		return
	}

	gap := arrLen
	shrink := 1.3
	swapped := true

	for gap > 1 || swapped == true {
		swapped = false
		gap = int(float64(gap) / shrink)
		if gap < 1 {
			gap = 1
		}

		for i := 0; i+gap < arrLen; i++ {
			if arr.Greater(i, i+gap) {
				arr.Swap(i, i+gap)
				swapped = true
			}
		}
	}

	return
}
