package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var ticker *time.Ticker

func main() {
	ctx := context.Background()
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGTERM)

	ticker = time.NewTicker(1 * time.Second)

	go func() {
		for {
			select {
			case t := <-ticker.C:
				fmt.Println("Current time is", t.Format(time.RFC3339Nano))
			}
		}
	}()

	for {
		select {
		case <-sigterm:
			fmt.Println("Sigterm at", time.Now().Format(time.RFC3339Nano))
			closeAllResources(ctx)
			return
		}
	}
}

func closeAllResources(ctx context.Context) {
	ctx, cancelFunc := context.WithTimeout(ctx, 1*time.Second)
	defer cancelFunc()
	ticker.Stop()
}
