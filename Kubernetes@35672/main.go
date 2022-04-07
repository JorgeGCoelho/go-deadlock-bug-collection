package Kubernetes_35672

import "time"

// stopCh is closed to stop this (and other) goroutines from a parent goroutine
// Seems like a naive implementation of context

// It is possible for the ListAndWatch function to return without stopping the goroutine
// The stopCh in received from the parameter, and seems to be used to stop the whole caching process, not just the
// goroutine.

// ListAndWatch pkg/client/cache/reflector.go:233
func ListAndWatch(stopCh <-chan struct{}) {

	// Channel based on a timer to resync something
	resyncCh := time.NewTimer(1 * time.Minute).C

	go func() {
		for {
			select {
			case <-resyncCh:
			case <-stopCh:
				return
			}
			// Resync something
		}
	}()
}
