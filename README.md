# Beard

[![GoDoc](https://godoc.org/github.com/yleichanok/beard?status.svg)](https://godoc.org/github.com/yleichanok/beard)

Contains implementations of various algorithms in golang. Implementations do not depend on data types and based on a standard Go `sort` package.

## Testing

To execute all tests run:

    go test

Default array size for benchmarking purposes is set to 100000 elements. It can be changed in `sort_test.go` (see const `SIZE`). To run benchmarks do:

    go test -bench=.

Results for an array of 10,000 integer elements:

| Name              | Iterations |           ns/op |                              |
| :---------------- | ---------: | --------------: | :--------------------------- |
| Quick sort        |       1000 |         2007156 |                              |
| Go native sort    |       1000 |         2014763 |                              |
| Comb sort         |        500 |         3239221 |                              |
| Heap sort         |        500 |         3061923 |                              |
| Shell sort        |        500 |         3102759 |                              |
| Merge sort        |        200 |         9318359 | Type casting used.           |
| Insertion sort    |          5 |       281339567 |                              |
| Selection sort    |          5 |       336679882 |                              |
| Gnome sort        |          5 |       426937305 |                              |
| Cocktail sort     |          2 |       530186142 |                              |
| Odd-even sort     |          2 |       512593399 |                              |
| Bubble sort       |          2 |       860656513 |                              |
| Cycle sort        |          1 |      6266636497 | Type casting used.           |
| Stooge sort       |          - |               - | Believe me, it's slow :)     |
