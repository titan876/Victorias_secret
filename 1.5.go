package main

import "fmt"

func main() {
	var a, n, p int
	fmt.Print("Введите числа A, N\n")
	fmt.Scan(&a, &n)
	p = 1
	for i := 0; i < n; i++ {
		p = p * (a + n)
		fmt.Printf("%d Pначениe: %d; \n", i+1, p)
	}
	fmt.Println("Ответ: ", p)
}
