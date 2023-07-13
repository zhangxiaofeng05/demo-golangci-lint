package main

import "errors"

var (
	Unknown = errors.New("a unknown error")
)

func errCheck() error {
	return Unknown
}
