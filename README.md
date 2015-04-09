# Beard

[![GoDoc](https://godoc.org/github.com/yleichanok/beard?status.svg)](https://godoc.org/github.com/yleichanok/beard)

Contains implementations of various algorithms in golang. Implementations do not depend on data types and based on a standard Go `sort` package.

## Testing

To execute all tests run:

    go test

Default array size for benchmarking purposes is set to 100000 elements. It can be changed in `sort_test.go` (see const `SIZE`). To run benchmarks do:

    go test -bench=.

Results for an array of 100,000 integer elements:

| Name              | Iterations |           ns/op |
| :---------------- | ---------: | --------------: |
| Quick sort        |        100 |        23392532 |
| Go native sort    |        100 |        27036809 |
| Heap sort         |         50 |        38473021 |
| Comb sort         |         50 |        39060081 |
| Shell sort        |         20 |        97606027 |
| Insertion sort    |          1 |     26890184579 |
| Selection sort    |          1 |     29874953237 |
| Gnome sort        |          1 |     40047292762 |
| Odd-even sort     |          1 |     48786066289 |
| Cocktail sort     |          1 |     49116488340 |
| Bubble sort       |          1 |     82878359824 |
| Stooge sort       |          - |               - |