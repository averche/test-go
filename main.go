package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(1*time.Second))
	defer cancel()

	ctx = context.WithValue(ctx, "a", "asd")

	sleepAndSay(ctx, 5*time.Second, "hello")
}

func sleepAndSay(ctx context.Context, sleepTime time.Duration, msg string) {
	select {
	case <-time.After(sleepTime):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Println(ctx.Err().Error())
	}
	func2(ctx)
}

func func2(ctx context.Context) {
	func3(ctx)
}

func func3(ctx context.Context) {

}
