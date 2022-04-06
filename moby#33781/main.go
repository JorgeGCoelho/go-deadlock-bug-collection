package moby_33781

import (
	"context"
)

// Goroutine leak due to failure to drain the channel
// daemon/health.go:184
func main(ctx context.Context, stop chan struct{}) {
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
