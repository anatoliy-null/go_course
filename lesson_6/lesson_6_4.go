package main

import "fmt"

type contacts struct {
	Addss string
	Phone string
}

type user struct {
	ID   int
	Name string
	contacts
}

type employee struct {
	ID   int
	Name string
	contacts
}

func main() {
	u := user{
		ID:   1,
		Name: "Иванов",
		contacts: contacts{
			Addss: "ivanov@company.ru",
			Phone: "+7 999 999 99 99",
		},
	}

	e := employee{
		ID:   1,
		Name: "Петров",
		contacts: contacts{
			Addss: "petrov@company.ru",
			Phone: "+7 988 888 88 88",
		},
	}

	fmt.Println(u.Addss, u.Phone, e.Addss, e.Phone)
}
