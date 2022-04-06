package moby_d55998b

// d55998be81973be5b083eed56b223c8ce98ce073

// This bug involves two channels, one that transfers the possible error of a background/asynchronous goroutine and the other
// signals that an event has happened (connection being hijacked in this case), allowing the parent goroutine to proceed

// A deadlock occurs when the second case
// TODO: Not I'm not sure

func main() {
	hijacked := make(chan bool)

	errCh := make(chan error)
	go func() {
		errCh <- func(chan bool) error {
			<-hijacked
			// Stuff
			return nil
		}(hijacked)
	}()

	// Acknowledge the hijack before starting
	select {
	case <-hijacked:
	case err := <-errCh:
		// This case is executed
		if err != nil {

		}
	}

	if errCh != nil {
		// no more err to receive; Block
		if err := <-errCh; err != nil {

		}
	}
}
