package main

import (
	"log"

	"github.com/hashicorp/consul-template/child"
)

func main() {
	input := &child.NewInput{
		Command: "./my-script.sh",
	}

	proc, err := child.New(input)
	if err != nil {
		panic(err)
	}

	if err := proc.Start(); err != nil {
		log.Fatalf("could not start proc: %s", err)
	}

	select {
	case exitCode := <-proc.ExitCh():
		log.Fatalf("exit code: %d", exitCode)
	}
}
