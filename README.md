# Beard

[![GoDoc](https://godoc.org/github.com/yleichanok/beard?status.svg)](https://godoc.org/github.com/yleichanok/beard)

Contains implementations of various algorithms in golang. Implementations do not depend on data types and based on a standard Go `sort` package.

## Testing

To execute all tests run:

    go test

Default array size for benchmarking purposes is set to 10,000 elements. It can be changed in `sort_test.go` (see const `SIZE`). To run benchmarks do:

    go test -bench=.

Results for an array of 10,000 integer elements:

| Name              | Iterations |           ns/op |                              |
| :---------------- | ---------: | --------------: | :--------------------------- |
| Quick sort        |       1000 |         1925249 |                              |
| Go native sort    |       1000 |         1999172 |                              |
| Heap sort         |        500 |         3024846 |                              |
| Comb sort         |        500 |         3117685 |                              |
| Shell sort        |        500 |         3134365 |                              |
| Merge sort        |        200 |         8754628 | Type casting used.           |
| Insertion sort    |          5 |       276740979 |                              |
| Selection sort    |          5 |       325540659 |                              |
| Gnome sort        |          5 |       408988837 |                              |
| Odd-even sort     |          2 |       489537071 |                              |
| Cocktail sort     |          2 |       504772233 |                              |
| Bingo sort        |          2 |       681953480 |                              |
| Bubble sort       |          2 |       831446304 |                              |
| Cycle sort        |          1 |      5932367672 | Type casting used.           |
| Stooge sort       |          - |               - | Believe me, it's slow :)     |
