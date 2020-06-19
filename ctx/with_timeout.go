package main

import (
	"context"
	"fmt"
	"time"
)

func CtxWithTimeout(ctx context.Context, expires time.Duration, message string) {
	ctx, cancel := context.WithTimeout(ctx, expires)
	defer cancel()

	time.Sleep(5 * time.Second)
	fmt.Println(message)
}
