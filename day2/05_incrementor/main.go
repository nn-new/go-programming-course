package main

import "fmt"

func main() {
	next := incrementor()
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
}

func incrementor() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
