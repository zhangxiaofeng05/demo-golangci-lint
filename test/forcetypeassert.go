package main

func forcetypeassert() {
	var a interface{} = 1
	b := a.(int)
	_ = b
}
