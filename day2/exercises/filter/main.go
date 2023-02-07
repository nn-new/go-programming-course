package main

import "fmt"

func filter(f func(int) bool, data []int) []int {
	var result []int
	for _, item := range data {
		if f(item) {
			result = append(result, item)
		}
	}
	return result
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	even := func(x int) bool {
		return x%2 == 0
	}
	fmt.Println(filter(even, data)) // [2, 4, 6, 8, 10]
}
