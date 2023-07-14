package main

import "fmt"

func errorlint() error {
	// bad
	return fmt.Errorf("oh noes: %v", ErrUnknown)

	// good
	//return fmt.Errorf("oh noes: %w", ErrUnknown)
}
