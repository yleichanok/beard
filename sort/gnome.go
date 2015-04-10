package sort

// Gnome sort algorithm.
//
// Puts items in order by comparing the current item with the previous item.
// If they are in order, move to the next item (or stop if the end is reached).
// If they are out of order, swap them and move to the previous item.
// If there is no previous item, move to the next item.
//
// See http://en.wikipedia.org/wiki/Gnome_sort.
func GnomeSort(arr Sortable) {

	if arr == nil {
		panic("arr is nil")
	}

	arrLen := arr.Len()

	// Quit if nothing to sort.
	if arrLen < 2 {
		return
	}

	for i := 1; i < arrLen; {
		if arr.Compare(i, i-1) >= 0 {
			i++
		} else {
			arr.Swap(i, i-1)
			if i > 1 {
				i--
			}
		}
	}

	return
}
