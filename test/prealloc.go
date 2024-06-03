package main

func prealloc() {
	var intSlice []int
	for _, p := range []int{10, 11, 12, 13} {
		//p := p // FIX variable into the local variable
		// 在循环中使用循环变量p的指针，是指向数组最后一个数字的指针
		//log.Printf("&p: %v", &p)
		intSlice = append(intSlice, p)
	}
	// log.Printf("intSlice: %v", intSlice)
}
