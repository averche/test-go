package main

import (
	"crypto/tls"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <cert> <key>\n", os.Args[0])
	}

	_, err := tls.LoadX509KeyPair(
		os.Args[1],
		os.Args[2],
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Success!")
}
