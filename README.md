# Beard

Contains implementations of various algorithms in golang.

## Testing

To execute all tests run:

    go test

Default array size for benchmarking purposes is set to 10000 elements. It can be changed in `sort_test.go` (see const `SIZE`). To run benchmarks do:

    go test -bench=.