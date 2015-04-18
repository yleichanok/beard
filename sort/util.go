package sort

// Common interface for sortable arrays to abstract functions.
type Sortable interface {

	// Returns array's length.
	Len() int

	// Compares arr[i] to arr[j] and returns
	// -1 if arr[i] < arr[j]
	// 0 if arr[i] == arr[j]
	// 1 if arr[i] > arr[j]
	Compare(i, j int) int8

	// Swaps arr[i] and arr[j] with each other.
	Swap(i, j int)
}
