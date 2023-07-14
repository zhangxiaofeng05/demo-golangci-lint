package main

import (
	"encoding/json"
	"log"
)

func musttag() {
	aa := struct {
		Name string
	}{
		Name: "testName",
	}
	marshal, err := json.Marshal(aa)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("marshal: %v", string(marshal))

}
