package sort

/**
 * Odd-even (brick) sort algorithm.
 *
 * Compares all (odd, even)-indexed pairs of adjacent elements
 * in the list and, if a pair is in the wrong order
 * (the first is larger than the second) the elements are switched.
 * The next step repeats this for (even, odd)-indexed pairs
 * (of adjacent elements). Then it alternates between (odd, even)
 * and (even, odd) steps until the list is sorted.
 *
 * See http://en.wikipedia.org/wiki/Odd-even_sort.
 */
func OddEvenSort(arr Sortable) {

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

		for i := 1; i < arrLen-1; i += 2 {
			if arr.Greater(i, i+1) {
				arr.Swap(i, i+1)
				swapped = true
			}
		}
		for i := 0; i < arrLen-1; i += 2 {
			if arr.Greater(i, i+1) {
				arr.Swap(i, i+1)
				swapped = true
			}
		}
	}

	return
}
