package main

import "fmt"

func fibonacci(first, second, count int) {
	if count > 0 {
		fmt.Printf("%d ", first)
		fibonacci(second, first+second, count-1)
	}
}

func main() {
	fibonacci(0, 1, 23)
}
