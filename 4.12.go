package main

import (
	"fmt"
	"math"
)

func main() {
	var min float64
	fmt.Printf("\n \t||============================||")
	fmt.Printf("\n \t||      x      ||      y      ||")
	fmt.Printf("\n \t||============================||")
	for x := -1.2; x < 1.5; x += 0.1 {
		y := math.Cos(math.Tan(math.Pow(x, 2))) + (1 / 4)
		if y < min {
			min = y
		}
		fmt.Printf("\n \t||    %4.1f    ||    %6.2f    ||", x, y)
	}
	fmt.Printf("\n \t||============================||")
	fmt.Printf("\n\n Наименьшее значение Y: %.2f", min)
}
