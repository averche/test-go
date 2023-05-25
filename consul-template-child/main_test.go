package main

import (
	"fmt"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestShortRun(t *testing.T) {
	g := &errgroup.Group{}

	for i := 0; i < 100; i++ {
		g.Go(func() error {
			_, err := shortRun()
			return err
		})
	}

	if err := g.Wait(); err != nil {
		t.Fatal(err)
	}
}

func TestLongRun(t *testing.T) {
	g := &errgroup.Group{}

	for i := 0; i < 100; i++ {
		g.Go(func() error {
			c, err := longRun()
			if err != nil {
				return err
			}
			if c != WeStoppedTheProcess {
				return fmt.Errorf("unexpected return code: %d", c)
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		t.Fatal(err)
	}
}
