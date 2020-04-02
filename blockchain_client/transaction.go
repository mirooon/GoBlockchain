package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Transaction struct {
	SenderPublicKey    string `json:"senderPublicKey"`
	SenderPrivateKey   string `json:"senderPrivateKey"`
	RecipientPublicKey string `json:"recipientPublicKey"`
	Amount             float32 `json:"amount"`
}

func NewTransaction(senderPublicKey string, senderPrivateKey string, recipientPublicKey string, amount float32) *Transaction {
	t := new(Transaction)
	t.SenderPublicKey = senderPublicKey
	t.SenderPrivateKey = senderPrivateKey
	t.RecipientPublicKey = recipientPublicKey
	t.Amount = amount
	return t
}

func ToBytes(o interface{}) []byte {
	return []byte(fmt.Sprintf("%v", o))
}

func SHA256(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(ToBytes(o)))
	return hex.EncodeToString(h.Sum(nil))
}

func (t *Transaction) SignTransaction() string {
	privateKey, err := hex.DecodeString(t.SenderPrivateKey)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	privateKeyECDSA, err := toECDSA(privateKey)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	
	transactionToSign := struct {
		SenderPublicKey    string
		RecipientPublicKey string
		Amount             float32
		}{t.SenderPublicKey, t.RecipientPublicKey, t.Amount}
		hashTx := SHA256(transactionToSign)
		r, s, err := ecdsa.Sign(rand.Reader, privateKeyECDSA, ToBytes(hashTx))
		if err != nil {
		fmt.Printf("%v\n", err)
	}
	signature := fromIntToHex(r) + fromIntToHex(s)
	return signature
}
