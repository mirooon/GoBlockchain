package main

import (
	"fmt"
)

type Transaction struct {
	senderPublicKey    string
	senderPrivateKey   string
	recipientPublicKey string
	amount             string
}

func NewTransaction(senderPublicKey string, senderPrivateKey string, recipientPublicKey string, amount string) *Transaction {
	t := new(Transaction)
	t.senderPublicKey = senderPublicKey
	t.senderPrivateKey = senderPrivateKey
	t.recipientPublicKey = recipientPublicKey
	t.amount = amount
	return t
}

func (t *Transaction) ShowTransactionInfo() {
	fmt.Printf("%+v\n", t)
}
