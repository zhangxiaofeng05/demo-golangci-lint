package main

import "strings"

func staticcheck(a, b string) bool {
	return strings.ToUpper(a) == strings.ToUpper(b)
	//return strings.EqualFold(a, b)
}
