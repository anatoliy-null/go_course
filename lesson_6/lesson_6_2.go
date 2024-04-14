package main

import (
	"fmt"
	"strings"
)

func main() {
	type contract struct {
		ID     int
		Number string
		Date   string
	}

	ctr := contract{
		ID:     1,
		Number: "#000A101\t01",
		Date:   "2024-01-31",
	}

	// По ТЗ кажется надо заменить символ табуляции на пробел?
	// Если нет, то строку ниже можно убрать
	ctr.Number = strings.Replace(ctr.Number, "\t", " ", -1)

	fmt.Printf("%+v\n", ctr)
}
