package main

import "fmt"

func main() {
	fmt.Println("Hola mundo")

	result := suma(30, 3)
	fmt.Println(result)
}

func suma(x int, y int) int {
	return x + y
}
