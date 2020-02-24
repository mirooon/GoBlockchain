package main

import (
	"fmt"
)

type Transaction struct {
	sth string
}

func NewTransaction() *Transaction {
	t := new(Transaction)
	t.sth = "TODO"
	return t
}

func (t *Transaction) ShowTransactionInfo() {
	fmt.Printf("%+v\n", t)
}
