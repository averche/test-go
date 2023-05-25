package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"time"

	"github.com/hashicorp/consul-template/child"
)

const WeStoppedTheProcess = -0xC0FFEE

func main() {
	code, err := run()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("exit code:", code)
}

// 0s : start the script
// 2s : stop the script
// 2s : the library attempts to kill the script
func run() (int, error) {
	process, err := child.New(&child.NewInput{
		Command:     "./my-script.sh",
		Stdout:      os.Stdout,
		KillSignal:  syscall.SIGTERM,
		KillTimeout: 5 * time.Second,
	})
	if err != nil {
		return -2, err
	}

	if err := process.Start(); err != nil {
		return -3, fmt.Errorf("could not start the process: %s", err)
	}

	select {
	case <-time.After(2 * time.Second):
		process.Stop()
		return WeStoppedTheProcess, nil

	case exitCode := <-process.ExitCh():
		return exitCode, nil
	}
}
