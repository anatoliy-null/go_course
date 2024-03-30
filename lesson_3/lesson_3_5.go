package main

import "fmt"

func main() {
	const (
		monday = iota + 1
		tuesday
		wednesday
		thursday
		friday
	)
	fmt.Println(monday, tuesday, wednesday, thursday, friday)
}
