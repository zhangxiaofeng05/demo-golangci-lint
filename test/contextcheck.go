package main

import (
	"context"
	"log"
)

func contextcheck(_ context.Context) {
	// 设置超时时间
	//service(ctx)
	service(context.Background())
}

func service(_ context.Context) {
	log.Println("contextcheck test")
}
