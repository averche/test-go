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

func main() {
	code, err := run()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("exit code:", code)
}

func run() (int, error) {
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
