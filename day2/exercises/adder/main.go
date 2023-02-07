package main

import "fmt"

func adder(f func(int) int) int {
	return f(2)
}

func main() {
	add10 := func(x int) int {
		return x + 10
	}

	fmt.Println(adder(add10)) // 12
}
