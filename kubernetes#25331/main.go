package main

import "sync"

// When it receives an error (via wc.errChan) it might block trying to send (wc.resultChan) due to no receivers.

// pkg/storage/etcd3/watcher.go:103
func (wc *watchChan) run() {
	var resultChanWG sync.WaitGroup
	resultChanWG.Add(1)
	// processEvent processes events from etcd watcher and sends results to resultChan
	go wc.processEvent(&resultChanWG)

	select {
	case err := <-wc.errChan:
		errResult := parseError(err)
		wc.cancel()
		if errResult != nil {
			wc.resultChan <- *errResult
		}
	case <-wc.ctx.Done():
	}
	resultChanWG.Wait()
	close(wc.resultChan)
}
