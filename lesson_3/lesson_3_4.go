package main

import "fmt"

const (
	codeBadRequest = iota + 400
	codeUnauthorized
	codePaymentRequired
	codeForbidden
	codeNotFound
)

func main() {
	fmt.Println(codeBadRequest, codeUnauthorized, codePaymentRequired, codeForbidden, codeNotFound)
}
