package main

import (
	randc "crypto/rand"
	"log"
	"math/big"
	randm "math/rand"
	"time"
)

func gosec() {
	r := randm.New(randm.NewSource(time.Now().UnixNano()))
	randomInt := r.Int()
	log.Printf("randomInt: %v", randomInt)
}

func gosecGood() {
	// 指定生成随机整数的范围，[0, max)
	max := big.NewInt(100)
	cryptoRandNum, err := randc.Int(randc.Reader, max)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("cryptoRandNum: %v", cryptoRandNum)
}
