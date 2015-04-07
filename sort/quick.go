package sort

/**
 * Quick (partition-exchange) sort algorithm.
 *
 * Picks an element from the array (the pivot),
 * partition the remaining elements into those greater than and less than this pivot,
 * and recursively sorts the partitions.
 *
 * See http://en.wikipedia.org/wiki/Quicksort.
 */
func QuickSort(arr Sortable) {
	if arr == nil {
		panic("arr is nil")
	}

	arrLen := arr.Len()

	// Quit if nothing to sort.
	if arrLen < 2 {
		return
	}

	quickSort(arr, 0, arrLen-1)
	return
}

func quickSort(arr Sortable, left, right int) {
	i := left
	j := right
	pivotIndex := (left + right) / 2

	for i <= j {
		for arr.Less(i, pivotIndex) {
			i++
		}
		for arr.Greater(j, pivotIndex) {
			j--
		}
		if i <= j {
			if pivotIndex == i {
				pivotIndex = j
			}
			if pivotIndex == j {
				pivotIndex = i
			}

			arr.Swap(i, j)
			i++
			j--
		}
	}

	if left < j {
		quickSort(arr, left, j)
	}
	if i < right {
		quickSort(arr, i, right)
	}
	return
}
