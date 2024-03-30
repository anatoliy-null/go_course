package main

import "fmt"

func main() {
	div := 16 / 3
	mod := 16 - div*3
	fmt.Printf("Результат: %d остаток от деления: %d тип результата: %[1]T", div, mod)
}
