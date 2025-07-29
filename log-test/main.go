package main

import (
	"log"
	"os"
	"time"
)

func main() {
	f, err := os.OpenFile("observations.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("could not open log file: %v", err)
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)

	for i := 1; ; i++ {
		logger.Println(i)
		time.Sleep(time.Second)
	}
}
