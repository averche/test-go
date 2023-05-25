package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"time"

	"github.com/hashicorp/consul-template/child"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

// 0s : start the script
// 2s : stop the script
// 2s : the library attempts to kill the script
func run() error {
	process, err := child.New(&child.NewInput{
		Command:     "./my-script.sh",
		Stdin:       os.Stdin,
		Stdout:      os.Stdout,
		Stderr:      os.Stderr,
		Env:         os.Environ(),
		KillSignal:  syscall.SIGTERM,
		KillTimeout: 5 * time.Second,
	})
	if err != nil {
		return err
	}

	if err := process.Start(); err != nil {
		return fmt.Errorf("could not start process: %s", err)
	}

	time.Sleep(2 * time.Second)

	process.Stop()

	return nil
}
