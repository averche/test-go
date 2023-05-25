package main

import (
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestRun(t *testing.T) {
	if err := run(); err != nil {
		t.Fatal(err)
	}
}

func TestRunParallel(t *testing.T) {
	g := &errgroup.Group{}

	for i := 0; i < 100; i++ {
		g.Go(run)
	}

	if err := g.Wait(); err != nil {
		t.Fatal(err)
	}
}
