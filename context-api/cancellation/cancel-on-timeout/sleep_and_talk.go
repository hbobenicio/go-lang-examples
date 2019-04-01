package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func sleepAndTalk(ctx context.Context, t time.Duration, msg string) {
	go func(ctx context.Context) {
		<-ctx.Done()
		log.Fatalln("cancellation happened")
	}(ctx)

	time.Sleep(t)
	fmt.Println(msg)
}
