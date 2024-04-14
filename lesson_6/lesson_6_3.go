package main

import (
	"fmt"
	"strings"
)

type contract struct {
	ID     int
	Number string
	Date   string
}

func (c contract) String() string {
	num := strings.Replace(c.Number, "\n", "\\n", -1)
	return fmt.Sprintf("Договор № %s от %s", num, c.Date)
}

func main() {
	ctr := contract{
		ID:     1,
		Number: "#000A\n101",
		Date:   "2024-01-31",
	}
	fmt.Println(ctr)
}
