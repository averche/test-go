# Results

```shell-session
$ go test -bench=.

goos: darwin
goarch: arm64
pkg: asd/vector-test
BenchmarkAdd-10                 34659668                34.52 ns/op
BenchmarkAddByValue-10          34189576                34.50 ns/op
BenchmarkAddReturnCopy-10       293998209                4.071 ns/op
PASS
ok      asd/vector-test 5.621s
```
