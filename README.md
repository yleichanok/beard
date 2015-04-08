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
| Quick sort        |        100 |        23158016 |
| Go native sort    |        100 |        26214720 |
| Comb sort         |         50 |        38533825 |
| Shell sort        |         20 |        81318971 |
| Insertion sort    |          1 |     25602297311 |
| Selection sort    |          1 |     27555065033 |
| Gnome sort        |          1 |     39905122346 |
| Odd-even sort     |          1 |     49141880163 |
| Cocktail sort     |          1 |     50132621173 |
| Bubble sort       |          1 |     83721916771 |
| Stooge sort       |          - |               - |