package sort

// Common interface for sortable arrays to abstract functions.
type Sortable interface {

	// Returns array's length.
	Len() int

	// Determines whether arr[i] is less than arr[j].
	Less(i, j int) bool

	// Determines whether arr[i] is greater than arr[j].
	Greater(i, j int) bool

	// Swaps arr[i] and arr[j] with each other.
	Swap(i, j int)
}
