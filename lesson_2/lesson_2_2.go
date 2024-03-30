package main

import "fmt"

func main() {
	fmt.Printf("Результат: %d остаток от деления: %d тип результата: %[1]T", 16/3, 16%3)
}
