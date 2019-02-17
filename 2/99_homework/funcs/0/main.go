package main

import (
	"fmt"
	"math"
)

// TODO: Реализовать вычисление Квадратного корня
func Sqrt(x float64) float64 {
	z := x
	for i := 5; i >= 0; i-- {
		z = z - (((z * z) - x) / (2 * z))
	}
	return z
}

func main() {
	fmt.Println(math.Sqrt(2))
	fmt.Println(Sqrt(2))
}
