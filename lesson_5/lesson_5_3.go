package main

import "fmt"

func main() {
	val := 1
	ptr := &val
	*ptr = 2
	fmt.Println(val)
}
