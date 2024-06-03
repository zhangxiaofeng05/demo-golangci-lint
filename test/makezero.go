package main

func makezero(nums []int) []int {
	values := make([]int, len(nums)) // satisfy prealloc
	values = append(values, nums...)
	return values
}
