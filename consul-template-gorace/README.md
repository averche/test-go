```sh
$ go test --race

...

==================
WARNING: DATA RACE
Write at 0x00c0005923c0 by goroutine 210:
  github.com/hashicorp/consul-template/child.(*Child).kill.func1()
      /Users/avean/go/pkg/mod/github.com/hashicorp/consul-template@v0.32.0/child/child.go:439 +0x84
  runtime.deferreturn()
      /opt/homebrew/Cellar/go/1.20.4/libexec/src/runtime/panic.go:476 +0x30
  github.com/hashicorp/consul-template/child.(*Child).internalStop()
      /Users/avean/go/pkg/mod/github.com/hashicorp/consul-template@v0.32.0/child/child.go:296 +0x140
  github.com/hashicorp/consul-template/child.(*Child).Stop()
      /Users/avean/go/pkg/mod/github.com/hashicorp/consul-template@v0.32.0/child/child.go:274 +0x2dc
  asd/consul-template-child.longRun()
      /Users/avean/github/averche/test-go/consul-template-child/main.go:76 +0x2d0
  asd/consul-template-child.TestLongRun.func1()
      /Users/avean/github/averche/test-go/consul-template-child/main_test.go:30 +0x20
  golang.org/x/sync/errgroup.(*Group).Go.func1()
      /Users/avean/go/pkg/mod/golang.org/x/sync@v0.2.0/errgroup/errgroup.go:75 +0x6c

Previous read at 0x00c0005923c0 by goroutine 443:
  github.com/hashicorp/consul-template/child.(*Child).kill.func2()
      /Users/avean/go/pkg/mod/github.com/hashicorp/consul-template@v0.32.0/child/child.go:457 +0x70

Goroutine 210 (running) created at:
  golang.org/x/sync/errgroup.(*Group).Go()
      /Users/avean/go/pkg/mod/golang.org/x/sync@v0.2.0/errgroup/errgroup.go:72 +0x10c
  asd/consul-template-child.TestLongRun()
      /Users/avean/github/averche/test-go/consul-template-child/main_test.go:29 +0x8c
  testing.tRunner()
      /opt/homebrew/Cellar/go/1.20.4/libexec/src/testing/testing.go:1576 +0x188
  testing.(*T).Run.func1()
      /opt/homebrew/Cellar/go/1.20.4/libexec/src/testing/testing.go:1629 +0x40

Goroutine 443 (running) created at:
  github.com/hashicorp/consul-template/child.(*Child).kill()
      /Users/avean/go/pkg/mod/github.com/hashicorp/consul-template@v0.32.0/child/child.go:455 +0x334
  github.com/hashicorp/consul-template/child.(*Child).internalStop()
      /Users/avean/go/pkg/mod/github.com/hashicorp/consul-template@v0.32.0/child/child.go:296 +0x140
  github.com/hashicorp/consul-template/child.(*Child).Stop()
      /Users/avean/go/pkg/mod/github.com/hashicorp/consul-template@v0.32.0/child/child.go:274 +0x2dc
  asd/consul-template-child.longRun()
      /Users/avean/github/averche/test-go/consul-template-child/main.go:76 +0x2d0
  asd/consul-template-child.TestLongRun.func1()
      /Users/avean/github/averche/test-go/consul-template-child/main_test.go:30 +0x20
  golang.org/x/sync/errgroup.(*Group).Go.func1()
      /Users/avean/go/pkg/mod/golang.org/x/sync@v0.2.0/errgroup/errgroup.go:75 +0x6c
==================
--- FAIL: TestLongRun (15.04s)
    testing.go:1446: race detected during execution of test
FAIL
exit status 1
FAIL    asd/consul-template-child       17.093s
```
