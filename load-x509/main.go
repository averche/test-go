package main

import (
	"crypto/tls"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <cert-file> <key-file>\n", os.Args[0])
	}

	log.Println("Certificate file:", os.Args[1])
	log.Println("Certificate key:", os.Args[2])

	_, err := tls.LoadX509KeyPair(
		os.Args[1],
		os.Args[2],
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Success!")
}
