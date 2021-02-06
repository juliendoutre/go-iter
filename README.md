# go-iter

## End goal

Implement iterators in Go, approaching performances of simple `for` loops, without external dependencies.

## Current state

Benchmarks show a x20 performance gap compared to `for` loops implementations:
```
BenchmarkRangeDivisorsSearch/with_a_loop-8                   888           1298827 ns/op         2917648 B/op         28 allocs/op
BenchmarkRangeDivisorsSearch/with_a_range-8                   48          23422088 ns/op        13829701 B/op     999775 allocs/op
BenchmarkVectorStringSearch/with_a_loop-8                2751242               419 ns/op               0 B/op          0 allocs/op
BenchmarkVectorStringSearch/with_a_vector-8               313200              3710 ns/op              32 B/op          1 allocs/op
```

Pprof leads me to think this is caused by excessive heap allocations due to the `interface{}` type.

After having moved to generated code, memory allocation performances improved a bit:
```
BenchmarkRangeDivisorsSearch/with_a_loop-8         	     156	   7966201 ns/op	 2917640 B/op	      28 allocs/op
BenchmarkRangeDivisorsSearch/with_a_range-8        	      26	  47768131 ns/op	 2917693 B/op	      30 allocs/op
BenchmarkVectorStringSearch/with_a_loop-8          	 1182849	       953 ns/op	       0 B/op	       0 allocs/op
BenchmarkVectorStringSearch/with_a_VectorOfInt-8   	  113923	      9625 ns/op	      32 B/op	       1 allocs/op
```
