package moby_33293

import "errors"

// ab83b924bc5dde5c2ac81b0befe5e370f3acb7e3

// Simple send to channel without receive
func main() chan error {
	errC := make(chan error)
	if true /* some error condition */ {
		errC <- errors.New("")
		return errC
	}
	return nil
}
