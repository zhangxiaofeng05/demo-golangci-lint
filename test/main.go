package main

import (
	"context"
	"fmt"
	"github.com/zhangxiaofeng05/com/db/commysql"
	"github.com/zhangxiaofeng05/golangcilint/test/example"
	"log"
	"time"
)

func init() {
	// default: LstdFlags = Ldate | Ltime
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	halfDsn := commysql.GetEnv()
	dsn := fmt.Sprintf("%s/%s?parseTime=true", halfDsn, "testdb")
	sqlxDB, err := commysql.Sqlx(dsn)
	if err != nil {
		log.Fatal(err)
	}

	{
		//错误检查 Error return value is not checked (errcheck)
		errCheck()
	}

	{
		//无效分配 ineffectual assignment to aa (ineffassign)
		aa := ineffassign()
		aa = 2 // 上一个 aa 变量没有使用，现在重新赋值了
		log.Printf("ineffassign: %v", aa)
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
		// 检查两个持续时间相乘 durationcheck
		durationcheck(5)               // Waiting for 5 seconds
		durationcheck(5 * time.Second) // Waiting for 1388888h53m20s
	}

	{
		// 错误lint errorlint
		errorlint()
	}

	{
		// 循环引用 exportloopref
		exportloopref()
	}

	{
		// 检查源代码是否存在安全问题 gosec
		gosec()
		gosecGood()
	}

	{
		// 确保切片不会以非零长度初始化的 linter makezero
		nums := []int{1, 2, 3, 4}
		copyNums := makezero(nums)
		log.Printf("copyNums: %v", copyNums)
	}

	{
		// 在(un)marshaled结构体时，强制标识tag musttag
		// 防止字段重命名，破坏约定
		musttag()
	}

	{
		//检查当 err 不为 nil 时返回 nil nilerr
		_ = nilerr()
	}

	{
		// 检查 sql.Rows 和 sql.Stmt 是否已关闭 sqlclosecheck
		sqlclosecheck(sqlxDB)
	}

	log.Printf("hello")
}
