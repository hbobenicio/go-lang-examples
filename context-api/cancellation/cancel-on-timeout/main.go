package main

import (
	"context"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	sleepAndTalk(ctx, 5*time.Second, "Hello, Context API!")
}
