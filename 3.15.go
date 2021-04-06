package main

import (
	"fmt"
)

func main() {
	var y, x float64
	fmt.Println("\nВведите Y")
	fmt.Scan(&y)
	fmt.Println("\nВведите X")
	fmt.Scan(&x)
	xc := x
	if x == y {
		fmt.Printf("Число X равно числу Y")
	} else if x < y {
		x := (x + y) / 2
		y := 2 * xc * y
		fmt.Printf("\n меньшее число X = %.2f, \n Большее число Y = %.2f", x, y)
	} else {
		x := 2 * x * y
		y := (xc + y) / 2
		fmt.Printf("\n меньшее число Y = %.2f, \n Большее число X = %.2f", y, x)
	}
}
