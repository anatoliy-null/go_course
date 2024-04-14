package main

import (
	"fmt"
	"strconv"
)

type contract struct {
	ID     int
	Number string
	Date   string
}

func main() {
	ctr := contract{
		ID:     1,
		Number: "#000A\n101",
		Date:   "2024-01-31",
	}

	out := strconv.Quote(fmt.Sprintf("%+v", ctr))
	fmt.Println(out[1 : len(out)-1])
}
