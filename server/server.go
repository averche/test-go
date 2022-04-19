package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		/* */ log.Println("handler: begin")
		defer log.Println("handler: end")

		ctx := r.Context()

		select {
		case <-time.After(5 * time.Second):
			fmt.Fprintln(w, "hello")
		case <-ctx.Done():
			log.Println(ctx.Err().Error())
			http.Error(w, ctx.Err().Error(), http.StatusInternalServerError)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
