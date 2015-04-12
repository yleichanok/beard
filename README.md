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
| Quick sort        |        100 |        24046186 |
| Go native sort    |        100 |        26385257 |
| Heap sort         |         50 |        39348495 |
| Comb sort         |         50 |        38874305 |
| Shell sort        |         20 |        75871708 |
| Merge sort        |         10 |       151857156 |
| Insertion sort    |          1 |     26730997119 |
| Selection sort    |          1 |     33175800770 |
| Gnome sort        |          1 |     41226572490 |
| Odd-even sort     |          1 |     49186325185 |
| Cocktail sort     |          1 |     50297429161 |
| Bubble sort       |          1 |     84889191452 |
| Stooge sort       |          - |               - |