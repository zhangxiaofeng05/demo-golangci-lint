package main

func makezero(nums []int) []int {
	values := make([]int, len(nums)) // satisfy prealloc
	for _, n := range nums {
		values = append(values, n)
	}
	return values
}
