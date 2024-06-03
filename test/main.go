package main

import (
	"context"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/zhangxiaofeng05/com/db/com_mysql"
)

func init() {
	// default: LstdFlags = Ldate | Ltime
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	halfDsn := com_mysql.GetEnv()
	dsn := fmt.Sprintf("%s/%s?parseTime=true", halfDsn, "testdb")
	sqlxDB, err := com_mysql.Sqlx(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer sqlxDB.Close()

	ctx := context.Background()

	//sqlClose(ctx, sqlxDB)

	start(ctx, sqlxDB)
}

func sqlClose(ctx context.Context, sqlxDB *sqlx.DB) {
	/**
	查看数据库最大预编译数量：show variables like '%max_prepared_stmt_count%';
	查看当前预编译数量和关闭数量：show global status like '%com_stmt%';
	*/

	for i := 0; i < 148; i++ {
		sqlclosecheck(ctx, sqlxDB)
	}
}

func start(ctx context.Context, sqlxDB *sqlx.DB) {
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
		// 检查源代码是否存在安全问题 gosec
		gosec()
		//gosecGood()
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
		//sqlclosecheck(ctx, sqlxDB)
	}

	log.Printf("hello")
}
