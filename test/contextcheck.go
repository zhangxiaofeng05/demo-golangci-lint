package main

import (
	"context"
	"log"
)

func contextcheck(ctx context.Context) {
	// 设置超时时间
	//service(ctx)
	service(context.Background())
}

func service(ctx context.Context) {
	log.Println("contextcheck test")
}
