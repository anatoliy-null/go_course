package main

import "fmt"

func main() {
	const (
		one = 1 << iota
		two
		four
		eight
		sixteen
	)
	fmt.Println(one, two, four, eight, sixteen)
}
