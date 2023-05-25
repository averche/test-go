package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"syscall"
	"time"

	"github.com/hashicorp/consul-template/child"
)

const WeStoppedTheProcess = -0xC0FFEE

func main() {
	code, err := shortRun()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("exit code:", code)
}

func shortRun() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	proc, err := child.New(&child.NewInput{
		Command: "./my-script-short.sh",
	})
	if err != nil {
		return -1, err
	}

	if err := proc.Start(); err != nil {
		return -1, fmt.Errorf("could not start proc: %s", err)
	}

	select {
	case <-ctx.Done():
		return -1, fmt.Errorf("did not get an exit code after 5 seconds")

	case exitCode := <-proc.ExitCh():
		return exitCode, nil
	}
}

func longRun() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	proc, err := child.New(&child.NewInput{
		Command:      "./my-script-long.sh",
		Stdin:        os.Stdin,
		Stdout:       os.Stdout,
		Stderr:       os.Stderr,
		Timeout:      0, // let it run forever
		Env:          os.Environ(),
		ReloadSignal: nil, // can't reload w/ new env vars
		KillSignal:   syscall.SIGTERM,
		KillTimeout:  10 * time.Second,
		Splay:        0,
	})
	if err != nil {
		return -1, err
	}

	if err := proc.Start(); err != nil {
		return -1, fmt.Errorf("could not start proc: %s", err)
	}

	select {
	case <-ctx.Done():
		proc.Stop()
		return WeStoppedTheProcess, nil

	case exitCode := <-proc.ExitCh():
		return exitCode, nil
	}
}
