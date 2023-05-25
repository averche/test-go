package main

import (
	"fmt"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestRunOnce(t *testing.T) {
	c, err := run()
	if err != nil {
		t.Fatal(err)
	}
	if c != SuccessfullyStoppedTheProcess {
		t.Fatalf("unexpected return code: %d", c)
	}
}

func TestRunParallel(t *testing.T) {
	g := &errgroup.Group{}

	for i := 0; i < 10; i++ {
		g.Go(func() error {
			c, err := run()
			if err != nil {
				return err
			}
			if c != SuccessfullyStoppedTheProcess {
				return fmt.Errorf("unexpected return code: %d", c)
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		t.Fatal(err)
	}
}
