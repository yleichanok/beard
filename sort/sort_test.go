package sort

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

// ----------------------------------------------------------------------------
// Preparing test...
// ----------------------------------------------------------------------------

// Total amount of elements in a test array.
const SIZE int = 10000

var list = make([]int, SIZE)
var prng = rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

// Wrapper for sorting by interger values.
type ByIntValue []int

func (arr ByIntValue) Len() int {
	return len(arr)
}
func (arr ByIntValue) Compare(i, j int) int8 {
	if arr[i] < arr[j] {
		return -1
	}
	if arr[i] > arr[j] {
		return 1
	}
	return 0
}
func (arr ByIntValue) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// Copies the value from one array to another.
// For merge sort only.
func copy(from, to Sortable, fromIndex, toIndex int) {
	to.(ByIntValue)[toIndex] = from.(ByIntValue)[fromIndex]
}

// Returns a value by index.
// For cycle sort only.
func get(arr Sortable, index int) interface{} {
	return arr.(ByIntValue)[index]
}

// Compares an element of array with item.
// For cycle sort only.
func compare(arr Sortable, index int, item interface{}) int8 {
	if arr.(ByIntValue)[index] < item.(int) {
		return -1
	}
	if arr.(ByIntValue)[index] > item.(int) {
		return 1
	}
	return 0
}

// Swaps an element of array and item.
// For cycle sort only.
func swap(arr Sortable, index int, item interface{}) interface{} {
	arr.(ByIntValue)[index], item = item.(int), arr.(ByIntValue)[index]
	return item
}

// ----------------------------------------------------------------------------
// Testing...
// ----------------------------------------------------------------------------

func TestBubbleSort(t *testing.T) {
	arr := []int{
		2, 5, 8, 1, 9, 3, 6, 9, 1, 3, 9, 4, 7, 1, 0, 5, -1, -3, 1, -5,
	}
	arrSorted := []int{
		-5, -3, -1, 0, 1, 1, 1, 1, 2, 3, 3, 4, 5, 5, 6, 7, 8, 9, 9, 9,
	}

	BubbleSort(ByIntValue(arr))
	if reflect.DeepEqual(arr, arrSorted) == false {
		t.Error("Bubble sort failed for %v", arr)
		return
	}

	fmt.Println("Bubble sort - OK")

	return
}

func TestCocktailSort(t *testing.T) {
	arr := []int{
		2, 5, 8, 1, 9, 3, 6, 9, 1, 3, 9, 4, 7, 1, 0, 5, -1, -3, 1, -5,
	}
	arrSorted := []int{
		-5, -3, -1, 0, 1, 1, 1, 1, 2, 3, 3, 4, 5, 5, 6, 7, 8, 9, 9, 9,
	}

	CocktailSort(ByIntValue(arr))
	if reflect.DeepEqual(arr, arrSorted) == false {
		t.Error("Cocktail sort failed for %v", arr)
		return
	}

	fmt.Println("Cocktail sort - OK")

	return
}

func TestCombSort(t *testing.T) {
	arr := []int{
		2, 5, 8, 1, 9, 3, 6, 9, 1, 3, 9, 4, 7, 1, 0, 5, -1, -3, 1, -5,
	}
	arrSorted := []int{
		-5, -3, -1, 0, 1, 1, 1, 1, 2, 3, 3, 4, 5, 5, 6, 7, 8, 9, 9, 9,
	}

	CombSort(ByIntValue(arr))
	if reflect.DeepEqual(arr, arrSorted) == false {
		t.Error("Comb sort failed for %v", arr)
		return
	}

	fmt.Println("Comb sort - OK")

	return
}

func TestCycleSort(t *testing.T) {
	arr := []int{
		2, 5, 8, 1, 9, 3, 6, 9, 1, 3, 9, 4, 7, 1, 0, 5, -1, -3, 1, -5,
	}
	arrSorted := []int{
		-5, -3, -1, 0, 1, 1, 1, 1, 2, 3, 3, 4, 5, 5, 6, 7, 8, 9, 9, 9,
	}

	CycleSort(ByIntValue(arr), get, compare, swap)
	if reflect.DeepEqual(arr, arrSorted) == false {
		t.Error("Cycle sort failed for %v", arr)
		return
	}

	fmt.Println("Cycle sort - OK")

	return
}

func TestHeapSort(t *testing.T) {
	arr := []int{
		2, 5, 8, 1, 9, 3, 6, 9, 1, 3, 9, 4, 7, 1, 0, 5, -1, -3, 1, -5,
	}
	arrSorted := []int{
		-5, -3, -1, 0, 1, 1, 1, 1, 2, 3, 3, 4, 5, 5, 6, 7, 8, 9, 9, 9,
	}

	HeapSort(ByIntValue(arr))
	if reflect.DeepEqual(arr, arrSorted) == false {
		t.Error("Quick sort failed for %v", arr)
		return
	}

	fmt.Println("Quick sort - OK")

	return
}

func TestGnomeSort(t *testing.T) {
	arr := []int{
		2, 5, 8, 1, 9, 3, 6, 9, 1, 3, 9, 4, 7, 1, 0, 5, -1, -3, 1, -5,
	}
	arrSorted := []int{
		-5, -3, -1, 0, 1, 1, 1, 1, 2, 3, 3, 4, 5, 5, 6, 7, 8, 9, 9, 9,
	}

	GnomeSort(ByIntValue(arr))
	if reflect.DeepEqual(arr, arrSorted) == false {
		t.Error("Gnome sort failed for %v", arr)
		return
	}

	fmt.Println("Gnome sort - OK")

	return
}

func TestInsertionSort(t *testing.T) {
	arr := []int{
		2, 5, 8, 1, 9, 3, 6, 9, 1, 3, 9, 4, 7, 1, 0, 5, -1, -3, 1, -5,
	}
	arrSorted := []int{
		-5, -3, -1, 0, 1, 1, 1, 1, 2, 3, 3, 4, 5, 5, 6, 7, 8, 9, 9, 9,
	}

	InsertionSort(ByIntValue(arr))
	if reflect.DeepEqual(arr, arrSorted) == false {
		t.Error("Insertion sort failed for %v", arr)
		return
	}

	fmt.Println("Insertion sort - OK")

	return
}

func TestMergeSort(t *testing.T) {
	arr := []int{
		2, 5, 8, 1, 9, 3, 6, 9, 1, 3, 9, 4, 7, 1, 0, 5, -1, -3, 1, -5,
	}
	arrSorted := []int{
		-5, -3, -1, 0, 1, 1, 1, 1, 2, 3, 3, 4, 5, 5, 6, 7, 8, 9, 9, 9,
	}

	tmp := make([]int, len(arr))

	MergeSort(ByIntValue(arr), ByIntValue(tmp), copy)
	if reflect.DeepEqual(arr, arrSorted) == false {
		t.Error("Merge sort failed for %v", arr)
		return
	}

	fmt.Println("Merge sort - OK")

	return
}

func TestOddEvenSort(t *testing.T) {
	arr := []int{
		2, 5, 8, 1, 9, 3, 6, 9, 1, 3, 9, 4, 7, 1, 0, 5, -1, -3, 1, -5,
	}
	arrSorted := []int{
		-5, -3, -1, 0, 1, 1, 1, 1, 2, 3, 3, 4, 5, 5, 6, 7, 8, 9, 9, 9,
	}

	OddEvenSort(ByIntValue(arr))
	if reflect.DeepEqual(arr, arrSorted) == false {
		t.Error("Odd-even sort failed for %v", arr)
		return
	}

	fmt.Println("Odd-even sort - OK")

	return
}

func TestQuickSort(t *testing.T) {
	arr := []int{
		2, 5, 8, 1, 9, 3, 6, 9, 1, 3, 9, 4, 7, 1, 0, 5, -1, -3, 1, -5,
	}
	arrSorted := []int{
		-5, -3, -1, 0, 1, 1, 1, 1, 2, 3, 3, 4, 5, 5, 6, 7, 8, 9, 9, 9,
	}

	QuickSort(ByIntValue(arr))
	if reflect.DeepEqual(arr, arrSorted) == false {
		t.Error("Quick sort failed for %v", arr)
		return
	}

	fmt.Println("Quick sort - OK")

	return
}

func TestSelectionSort(t *testing.T) {
	arr := []int{
		2, 5, 8, 1, 9, 3, 6, 9, 1, 3, 9, 4, 7, 1, 0, 5, -1, -3, 1, -5,
	}
	arrSorted := []int{
		-5, -3, -1, 0, 1, 1, 1, 1, 2, 3, 3, 4, 5, 5, 6, 7, 8, 9, 9, 9,
	}

	SelectionSort(ByIntValue(arr))
	if reflect.DeepEqual(arr, arrSorted) == false {
		t.Error("Selection sort failed for %v", arr)
		return
	}

	fmt.Println("Selection sort - OK")

	return
}

func TestShellSort(t *testing.T) {
	arr := []int{
		2, 5, 8, 1, 9, 3, 6, 9, 1, 3, 9, 4, 7, 1, 0, 5, -1, -3, 1, -5,
	}
	arrSorted := []int{
		-5, -3, -1, 0, 1, 1, 1, 1, 2, 3, 3, 4, 5, 5, 6, 7, 8, 9, 9, 9,
	}

	ShellSort(ByIntValue(arr))
	if reflect.DeepEqual(arr, arrSorted) == false {
		t.Error("Shell sort failed for %v", arr)
		return
	}

	fmt.Println("Shell sort - OK")

	return
}

func TestStoogeSort(t *testing.T) {
	arr := []int{
		2, 5, 8, 1, 9, 3, 6, 9, 1, 3, 9, 4, 7, 1, 0, 5, -1, -3, 1, -5,
	}
	arrSorted := []int{
		-5, -3, -1, 0, 1, 1, 1, 1, 2, 3, 3, 4, 5, 5, 6, 7, 8, 9, 9, 9,
	}

	StoogeSort(ByIntValue(arr))
	if reflect.DeepEqual(arr, arrSorted) == false {
		t.Error("Stooge sort failed for %v", arr)
		return
	}

	fmt.Println("Stooge sort - OK")

	return
}

// ----------------------------------------------------------------------------
// Benchmarking...
// ----------------------------------------------------------------------------

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for i := range list {
			list[i] = prng.Int()
		}
		b.StartTimer()

		BubbleSort(ByIntValue(list))
	}
	return
}

func BenchmarkCocktailSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for i := range list {
			list[i] = prng.Int()
		}
		b.StartTimer()

		CocktailSort(ByIntValue(list))
	}
	return
}

func BenchmarkCombSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for i := range list {
			list[i] = prng.Int()
		}
		b.StartTimer()

		CombSort(ByIntValue(list))
	}
	return
}

func BenchmarkCycleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for i := range list {
			list[i] = prng.Int()
		}
		b.StartTimer()

		CycleSort(ByIntValue(list), get, compare, swap)
	}
	return
}

func BenchmarkHeapSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for i := range list {
			list[i] = prng.Int()
		}
		b.StartTimer()

		HeapSort(ByIntValue(list))
	}
	return
}

func BenchmarkGnomeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for i := range list {
			list[i] = prng.Int()
		}
		b.StartTimer()

		GnomeSort(ByIntValue(list))
	}
	return
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for i := range list {
			list[i] = prng.Int()
		}
		b.StartTimer()

		InsertionSort(ByIntValue(list))
	}
	return
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for i := range list {
			list[i] = prng.Int()
		}
		b.StartTimer()

		tmp := make([]int, len(list))
		MergeSort(ByIntValue(list), ByIntValue(tmp), copy)
	}
	return
}

func BenchmarkNativeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for i := range list {
			list[i] = prng.Int()
		}
		b.StartTimer()

		sort.Ints(list)
	}
	return
}

func BenchmarkOddEvenSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for i := range list {
			list[i] = prng.Int()
		}
		b.StartTimer()

		OddEvenSort(ByIntValue(list))
	}
	return
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for i := range list {
			list[i] = prng.Int()
		}
		b.StartTimer()

		QuickSort(ByIntValue(list))
	}
	return
}

func BenchmarkSelectionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for i := range list {
			list[i] = prng.Int()
		}
		b.StartTimer()

		SelectionSort(ByIntValue(list))
	}
	return
}

func BenchmarkShellSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for i := range list {
			list[i] = prng.Int()
		}
		b.StartTimer()

		ShellSort(ByIntValue(list))
	}
	return
}
