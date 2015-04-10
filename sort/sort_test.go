package sort

import (
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
const SIZE int = 100000

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
		t.Error("Bubble sort failed for %v", arr)
		return
	}

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
		t.Error("Bubble sort failed for %v", arr)
		return
	}

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
		t.Error("Bubble sort failed for %v", arr)
		return
	}

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
		t.Error("Quick sort failed for %v", arr)
		return
	}

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
		t.Error("Quick sort failed for %v", arr)
		return
	}

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
		t.Error("Quick sort failed for %v", arr)
		return
	}

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
		t.Error("Quick sort failed for %v", arr)
		return
	}

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
		t.Error("Quick sort failed for %v", arr)
		return
	}

	return
}

// ----------------------------------------------------------------------------
// Benchmarking...
// ----------------------------------------------------------------------------

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
