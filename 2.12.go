package main

import (
	"fmt"
	"math"
)

func main() {
	var y float64
	fmt.Println("Введите Y")
	fmt.Scan(&y)
	for x := -3.2; x <= 3.2; x += 0.1 {
		z := 1 + math.Cos(math.Pow(x, 2)+y)/1 + math.Pow(x, 2)*(math.Pow(y, 2))
		fmt.Printf("\n x = %.1f, z = %.3f", x, z)
	}
}
