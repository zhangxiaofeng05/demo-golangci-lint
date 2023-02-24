package main

import "errors"

var (
	Unknown = errors.New("a unknown error")
)

func A() error {
	return Unknown
}

func main() {
	// A()
}
