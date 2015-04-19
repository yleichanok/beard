package sort

// Bingo sort algorithm.
//
// A variant of selection sort that orders items by first finding the least value,
// then repeatedly moving all items with that value to their final location and find the least value for the next pass.
// This is more efficient than selection sort if there are many duplicate values.
//
// See http://en.wikipedia.org/wiki/Selection_sort.
func BingoSort(arr Sortable) {

	if arr == nil {
		panic("arr is nil")
	}

	arrLen := arr.Len()

	// Quit if nothing to sort.
	if arrLen < 2 {
		return
	}

	i := arrLen - 1

	for i > 0 {
		maxIndex := 0

		for j := i; j > 0; j-- {
			if arr.Compare(maxIndex, j) == -1 {
				maxIndex = j
			}
		}

		arr.Swap(i, maxIndex)

		maxIndex = i
		i--
		j := i

		for j > 0 {
			j--
			if arr.Compare(maxIndex, j) == 0 {
				arr.Swap(i, j)
				i--
			}
		}
	}

	return
}
