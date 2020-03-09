package main

import (
	"fmt"
)

type Transaction struct {
	SenderPublicKey    string `json:"senderPublicKey"`
	SenderPrivateKey   string `json:"senderPrivateKey"`
	RecipientPublicKey string `json:"recipientPublicKey"`
	Amount             string `json:"amount"`
}

func NewTransaction(senderPublicKey string, senderPrivateKey string, recipientPublicKey string, amount string) *Transaction {
	t := new(Transaction)
	t.SenderPublicKey = senderPublicKey
	t.SenderPrivateKey = senderPrivateKey
	t.RecipientPublicKey = recipientPublicKey
	t.Amount = amount
	return t
}

func (t *Transaction) ShowTransactionInfo() {
	fmt.Printf("%+v\n", t)
}
