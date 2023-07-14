package main

import (
	"log"
	"time"
)

func durationcheck(someDuration time.Duration) {
	timeToWait := someDuration * time.Second
	log.Printf("Waiting for %s", timeToWait)
}
