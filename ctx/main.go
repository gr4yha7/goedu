package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	// go func() {
	// 	s := bufio.NewScanner(os.Stdin)
	// 	s.Scan()
	// 	cancel()
	// 	// 	time.Sleep(time.Second)
	// }()
	// time.AfterFunc(time.Second, cancel)

	doStuff(ctx, 5*time.Second, "hello")
}

func doStuff(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
	// ctx.Done()
}
