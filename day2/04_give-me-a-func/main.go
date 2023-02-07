package main

import "fmt"

func main() {
	myFunc := giveMeAFunc()
	myFunc("Hello world")
}

func giveMeAFunc() func(string) {
	return func(message string) {
		fmt.Println(message)
	}
}
