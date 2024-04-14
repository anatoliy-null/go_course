package main

import "fmt"

func change(ptr *string) {
	*ptr = "измененное значение"
}

func main() {
	val := "локальная переменная"
	change(&val)
	fmt.Println(val)
}
