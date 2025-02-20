package main

import (
	httpproxytest "asd/http-proxy-test"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestRun(t *testing.T) {
	g := &errgroup.Group{}

	for i := 0; i < 100; i++ {
		g.Go(func() error {
			_, err := run()
			return err
		})
	}

	if err := g.Wait(); err != nil {
		t.Fatal(err)
	}

	e.Fatal(httpproxytest.TestHelper())
}
