package main

func main() {
	{
		//错误检查 Error return value is not checked (errcheck)
		errCheck()
	}

	{
		//无效分配 ineffectual assignment to aa (ineffassign)
		aa := ineffassign()
		aa = 2 // 上一个 aa 变量没有使用，现在重新赋值了
		println(aa)
	}

	{
		//静态检查
		//https: //staticcheck.dev/docs/checks
		staticcheck(nil)
	}

	println("hello")
}
