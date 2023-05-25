## go test

```sh
$ go test -run TestRunOnce                                                                                                                  27s  13:54:16

2023/05/25 13:54:19 [INFO] (child) spawning: ./my-script.sh
sleeping for 20s
sleeping for 19s
2023/05/25 13:54:21 [INFO] (child) stopping process
received SIGTERM; ignoring it
sleeping for 18s
sleeping for 17s
sleeping for 16s
sleeping for 15s
sleeping for 14s
PASS
ok      asd/consul-template-gorace      7.180s
```

## go test --race

```sh
$ go test -run TestRunOnce --race

2023/05/25 13:55:39 [INFO] (child) spawning: ./my-script.sh
sleeping for 20s
sleeping for 19s
2023/05/25 13:55:41 [INFO] (child) stopping process
received SIGTERM; ignoring it
sleeping for 18s
sleeping for 17s
sleeping for 16s
sleeping for 15s
sleeping for 14s
==================
WARNING: DATA RACE
Write at 0x00c0000f20c0 by goroutine 6:
  github.com/hashicorp/consul-template/child.(*Child).kill.func1()
      /Users/avean/go/pkg/mod/github.com/hashicorp/consul-template@v0.32.0/child/child.go:439 +0x84
  runtime.deferreturn()
      /opt/homebrew/Cellar/go/1.20.4/libexec/src/runtime/panic.go:476 +0x30
  github.com/hashicorp/consul-template/child.(*Child).internalStop()
      /Users/avean/go/pkg/mod/github.com/hashicorp/consul-template@v0.32.0/child/child.go:296 +0x140
  github.com/hashicorp/consul-template/child.(*Child).Stop()
      /Users/avean/go/pkg/mod/github.com/hashicorp/consul-template@v0.32.0/child/child.go:274 +0x1f0
  asd/consul-template-gorace.run()
      /Users/avean/github/averche/test-go/consul-template-gorace/main.go:45 +0x1e4
  asd/consul-template-gorace.TestRunOnce()
      /Users/avean/github/averche/test-go/consul-template-gorace/main_test.go:11 +0x24
  testing.tRunner()
      /opt/homebrew/Cellar/go/1.20.4/libexec/src/testing/testing.go:1576 +0x188
  testing.(*T).Run.func1()
      /opt/homebrew/Cellar/go/1.20.4/libexec/src/testing/testing.go:1629 +0x40

Previous read at 0x00c0000f20c0 by goroutine 9:
  github.com/hashicorp/consul-template/child.(*Child).kill.func2()
      /Users/avean/go/pkg/mod/github.com/hashicorp/consul-template@v0.32.0/child/child.go:457 +0x70

Goroutine 6 (running) created at:
  testing.(*T).Run()
      /opt/homebrew/Cellar/go/1.20.4/libexec/src/testing/testing.go:1629 +0x5e4
  testing.runTests.func1()
      /opt/homebrew/Cellar/go/1.20.4/libexec/src/testing/testing.go:2036 +0x80
  testing.tRunner()
      /opt/homebrew/Cellar/go/1.20.4/libexec/src/testing/testing.go:1576 +0x188
  testing.runTests()
      /opt/homebrew/Cellar/go/1.20.4/libexec/src/testing/testing.go:2034 +0x700
  testing.(*M).Run()
      /opt/homebrew/Cellar/go/1.20.4/libexec/src/testing/testing.go:1906 +0x950
  main.main()
      _testmain.go:49 +0x300

Goroutine 9 (running) created at:
  github.com/hashicorp/consul-template/child.(*Child).kill()
      /Users/avean/go/pkg/mod/github.com/hashicorp/consul-template@v0.32.0/child/child.go:455 +0x334
  github.com/hashicorp/consul-template/child.(*Child).internalStop()
      /Users/avean/go/pkg/mod/github.com/hashicorp/consul-template@v0.32.0/child/child.go:296 +0x140
  github.com/hashicorp/consul-template/child.(*Child).Stop()
      /Users/avean/go/pkg/mod/github.com/hashicorp/consul-template@v0.32.0/child/child.go:274 +0x1f0
  asd/consul-template-gorace.run()
      /Users/avean/github/averche/test-go/consul-template-gorace/main.go:45 +0x1e4
  asd/consul-template-gorace.TestRunOnce()
      /Users/avean/github/averche/test-go/consul-template-gorace/main_test.go:11 +0x24
  testing.tRunner()
      /opt/homebrew/Cellar/go/1.20.4/libexec/src/testing/testing.go:1576 +0x188
  testing.(*T).Run.func1()
      /opt/homebrew/Cellar/go/1.20.4/libexec/src/testing/testing.go:1629 +0x40
==================
--- FAIL: TestRunOnce (7.01s)
    testing.go:1446: race detected during execution of test
FAIL
exit status 1
FAIL	asd/consul-template-gorace	7.251s
```
