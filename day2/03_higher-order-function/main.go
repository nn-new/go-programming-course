package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("vim-go")
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(compute(math.Pow))
	fmt.Println(compute(hypot))
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
