# Results

```shell-session
$ go test -bench=.
goos: darwin
goarch: arm64
pkg: asd/return-by-benchmarks
Benchmark_Add-10                   	34738862	        34.62  ns/op
Benchmark_AddByValue-10            	27114303	        44.45  ns/op
Benchmark_Add_InPlace-10           	255989691	         4.631 ns/op
Benchmark_AddByValue_InPlace-10    	350486653	         3.440 ns/op
Benchmark_AddReturnCopy-10         	12094290	       102.6   ns/op
PASS
ok  	asd/return-by-benchmarks	8.046s
```
