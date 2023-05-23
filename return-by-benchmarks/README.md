# Results

```shell-session
$ go test -bench=.

goos: darwin
goarch: arm64
pkg: asd/return-by-benchmarks
BenchmarkAdd-10                 28339845                35.37 ns/op
BenchmarkAddByValue-10          34687761                34.94 ns/op
BenchmarkAddReturnCopy-10       34478547                35.80 ns/op
PASS
ok      asd/return-by-benchmarks        3.671s
```
