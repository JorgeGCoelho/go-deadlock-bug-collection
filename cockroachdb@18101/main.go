package cockroachdb_18101

import "context"

// The sending goroutine, will leak/block when the receiving/parent goroutine cancels the context

func main() {
	ctx := context.Background()

	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)

	channel := make(chan struct{})
	go func() {
		defer close(channel)
		channel <- struct{}{}
		channel <- struct{}{}
	}()

	for _ := range channel {
		cancel()
		break
	}
}
