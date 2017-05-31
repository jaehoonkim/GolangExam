package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("--- start ---")

	ctx, cancel := context.WithCancel(context.Background())

	go Render1(ctx)
	//go Render2(ctx)

	go func() {
		select {
		case <-time.After(10 * time.Second):
			cancel()
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Printf("[main] %v\n", ctx.Err())
	}
	fmt.Println("--- end ---")
}

func Render1(ctx context.Context) {
	n := 0

	childCtx, _ := context.WithCancel(ctx)

	go Render2(childCtx)
	for {
		select {
		case <-time.After(2 * time.Second):
			n++
			fmt.Printf("[Render1] %d\n", n)
		case <-ctx.Done():
			fmt.Printf("[Render1] %v\n", ctx.Err())
			return
			//case <-time.After(7 * time.Second):
			//	cancel()
		}
	}
}

func Render2(ctx context.Context) {
	n := 999

	for {
		select {
		case <-time.After(3 * time.Second):
			n++
			fmt.Printf("[Render2] %d\n", n)
		case <-ctx.Done():
			fmt.Printf("[Render2] %v\n", ctx.Err())
			return
		}
	}
}
