package main

import "errors"

var (
	ErrUnknown = errors.New("a unknown error")
)

func errCheck() error {
	return ErrUnknown
}
