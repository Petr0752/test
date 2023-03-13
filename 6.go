package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	stop := make(chan struct{})

	// Способ 1: Использование флага
	flag := false
	go func() {
		for {
			if flag {
				break
			}
			fmt.Println("Goroutine 1")
			time.Sleep(1 * time.Second)
		}
	}()

	// Способ 2: Использование канала
	go func() {
		for {
			select {
			case <-stop:
				return
			default:
				fmt.Println("Goroutine 2")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// Способ 3: Использование пакета `Context`
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("Goroutine 3")
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(5 * time.Second)
	flag = true
	close(stop)
	cancel()
	fmt.Println("main function exited")
}
