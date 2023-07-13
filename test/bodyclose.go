package main

import (
	"io"
	"log"
	"net/http"
)

func bodyclose() {
	resp, err := http.Get("https://baidu.com")
	if err != nil {
		log.Printf("%v", err)
		return
	}
	//defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("%v", err)
		return
	}
	log.Printf("resp: %v byte", len(bytes))
}
