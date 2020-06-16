package main

import (
	"fmt"
	"math"
)

func maximum(numeros ...float64) float64 {
	max := math.Inf(-1) // Empezar con un valor realmente bajo
	// proceso en for on variaciones de los valores
	for _, number := range numeros {
		if number > max {
			max = number
		}
	}
	return max
}
func main() {
	fmt.Println(maximum(71.8, 56.2, 89.4))
}
