package sort

// Cocktail (bidirectional bubble, shaker) sort algorithm.
//
// A variant of bubble sort that compares each adjacent pair of items in a list in turn,
// swapping them if necessary, and alternately passes through the list from the beginning
// to the end then from the end to the beginning. It stops when a pass does no swaps.
//
// See http://en.wikipedia.org/wiki/Cocktail_sort.
func CocktailSort(arr Sortable) {

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

		for i := 0; i < arrLen-2; i++ {
			if arr.Greater(i, i+1) {
				arr.Swap(i, i+1)
				swapped = true
			}
		}

		if swapped == false {
			break
		}

		swapped = false

		for i := arrLen - 2; i >= 0; i-- {
			if arr.Greater(i, i+1) {
				arr.Swap(i, i+1)
				swapped = true
			}
		}
	}

	return
}
