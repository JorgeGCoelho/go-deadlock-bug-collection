package main

import (
	"context"
)

// Goroutine leak due to failure to drain the channel
// daemon/health.go:184
func monitor(ctx context.Context, stop chan struct{}) {
	results := make(chan struct{})

	go func() {
		results <- struct{}{}
		close(results)
	}()

	select {
	case <-stop:
		// results is not emptied, goroutine leak
		return
	case _ = <-results:
	case <-ctx.Done():
		<-results
	}
}

func main() {
	ctx := context.Background()
	stop := make(chan struct{})
	monitor(ctx, stop)
	close(stop)
}
