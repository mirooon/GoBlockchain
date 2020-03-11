package main

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Transaction struct {
	SenderPublicKey    string `json:"SenderPublicKey"`
	RecipientPublicKey string `json:"RecipientPublicKey"`
	Signature          string `json:"Signature"`
	Amount             string `json:"Amount"`
}

func NewTransaction(senderPublicKey string, recipientPublicKey string, signature string, amount string) *Transaction {
	t := new(Transaction)
	t.SenderPublicKey = senderPublicKey
	t.RecipientPublicKey = recipientPublicKey
	t.Signature = signature
	t.Amount = amount
	return t
}

func (t *Transaction) ShowTransactionInfo() {
	fmt.Printf("%+v\n", t)
}

func ToBytes(o interface{}) []byte {
	return []byte(fmt.Sprintf("%v", o))
}

func SHA256(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(ToBytes(o)))
	return hex.EncodeToString(h.Sum(nil))
}

func (t *Transaction) VerifyTransaction() bool {
	hashTx := SHA256(t)
	rInt, sInt := signatureHexToIntPair(t.Signature)
	fmt.Printf("rInt\n")
	fmt.Printf("%+v\n", rInt)
	fmt.Printf("sInt\n")
	fmt.Printf("%+v\n", sInt)
	publicKeyBytes, err := hex.DecodeString(t.SenderPublicKey)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("publicKeyBytes\n")
	fmt.Printf("%+v\n", publicKeyBytes)
	publicKeyECDSA := ToECDSAPub(publicKeyBytes)
	res := ecdsa.Verify(publicKeyECDSA, ToBytes(hashTx), rInt, sInt)
	fmt.Printf("res\n")
	fmt.Printf("%+v\n", res)
	return res
}
