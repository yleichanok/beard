package sort

// Cycle sort algorithm.
//
// Based on the idea that the permutation to be sorted can be factored into cycles,
// which can individually be rotated to give a sorted result.
//
// See http://en.wikipedia.org/wiki/Cycle_sort.
func CycleSort(
	arr Sortable,
	get func(arr Sortable, index int) interface{},
	compare func(arr Sortable, index int, item interface{}) int8,
	swap func(arr Sortable, index int, item interface{}) interface{}) {

	if arr == nil {
		panic("arr is nil")
	}

	arrLen := arr.Len()

	// Quit if nothing to sort.
	if arrLen < 2 {
		return
	}

	for cycleStart := 0; cycleStart < arrLen-1; cycleStart++ {
		item := get(arr, cycleStart)
		pos := cycleStart

		for i := cycleStart + 1; i < arrLen; i++ {
			if compare(arr, i, item) == -1 {
				pos++
			}
		}
		if pos == cycleStart {
			continue
		}

		for compare(arr, pos, item) == 0 {
			pos++
		}

		item = swap(arr, pos, item)

		for pos != cycleStart {
			pos = cycleStart

			for i := cycleStart + 1; i < arrLen; i++ {
				if compare(arr, i, item) == -1 {
					pos++
				}
			}

			for compare(arr, pos, item) == 0 {
				pos++
			}

			item = swap(arr, pos, item)
		}
	}

	return
}
