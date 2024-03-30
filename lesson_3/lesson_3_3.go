package main

import "fmt"

const test = "глобальная константа"

func main() {
	const test = "локальная константа"
	fmt.Printf("В консоль будет выведена %s!", test)
}
