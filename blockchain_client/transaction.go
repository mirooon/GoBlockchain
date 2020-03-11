package main

import (
	// "crypto"
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

func ToBytes(o interface{}) []byte {
	return []byte(fmt.Sprintf("%v", o))
}

func SHA256(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(ToBytes(o)))
	return hex.EncodeToString(h.Sum(nil))
}

func (t *Transaction) ShowTransactionInfo() {
	fmt.Printf("%+v\n", t)
}

func (t *Transaction) SignTransaction() string {
	fmt.Printf("SenderPrivateKey\n")
	fmt.Printf("%v\n", t.SenderPrivateKey)
	privateKey, err := hex.DecodeString(t.SenderPrivateKey)
	fmt.Printf("PrivateKey\n")
	fmt.Printf("%v\n", len(privateKey))
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	privateKeyECDSA, err := toECDSA(privateKey)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	hashTx := SHA256(t)
	r, s, err := ecdsa.Sign(rand.Reader, privateKeyECDSA, ToBytes(hashTx))
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	signature := fromIntToHex(r) + fromIntToHex(s)
	return signature
}
