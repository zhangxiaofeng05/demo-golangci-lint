package main

import (
	"context"
	"github.com/zhangxiaofeng05/golangcilint/test/example"
	"log"
	"time"
)

func init() {
	// default: LstdFlags = Ldate | Ltime
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	{
		//错误检查 Error return value is not checked (errcheck)
		//errCheck()
	}

	{
		//无效分配 ineffectual assignment to aa (ineffassign)
		//aa := ineffassign()
		//aa = 2 // 上一个 aa 变量没有使用，现在重新赋值了
		//log.Printf("ineffassign: %v", aa)
	}

	{
		//静态检查 staticcheck
		//https://staticcheck.dev/docs/checks
		eq := staticcheck("aa", "AA")
		log.Printf("staticcheck: %v", eq)
	}

	{
		// 类型检查 typecheck - 看函数的测试代码
		example.TypeCheck("typecheckParam")
	}

	{
		//http响应body是否关闭 bodyclose
		bodyclose()
	}

	{
		// 检查函数是否使用上下文的context contextcheck
		contextcheck(context.Background())
	}

	{
		durationcheck(5)               // Waiting for 5 seconds
		durationcheck(5 * time.Second) // Waiting for 1388888h53m20s
	}

	log.Printf("hello")
}

func durationcheck(someDuration time.Duration) {
	timeToWait := someDuration * time.Second
	log.Printf("Waiting for %s", timeToWait)
}
