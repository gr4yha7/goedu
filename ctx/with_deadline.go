package main

import (
	"context"
	"fmt"
	"time"
)

func CtxWithDeadline(ctx context.Context, expires time.Time, message string) {
	ctx, cancel := context.WithDeadline(ctx, expires)
	defer cancel()

	time.Sleep(5 * time.Second)
	fmt.Println(message)
}
