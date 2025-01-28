package main

import "fmt"

func main() {
	fmt.Println("Resultado da soma de 1 + 2:", sum(1, 2))
}

func sum(a, b int) int {
	return a + b
}
