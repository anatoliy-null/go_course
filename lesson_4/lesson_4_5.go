package main

import "fmt"

func hello() {
	defer fmt.Println("Завершение работы")
	fmt.Println("Hello, Go!")
}

func main() {
	hello()
}
