package main

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Transaction struct {
	SenderPublicKey    string  `json:"senderPublicKey"`
	RecipientPublicKey string  `json:"recipientPublicKey"`
	Signature          string  `json:"signature"`
	Amount             float32 `json:"amount"`
}

func NewTransaction(senderPublicKey string, recipientPublicKey string, signature string, amount float32) *Transaction {
	t := new(Transaction)
	t.SenderPublicKey = senderPublicKey
	t.RecipientPublicKey = recipientPublicKey
	t.Signature = signature
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

func (t *Transaction) VerifyTransaction() bool {
	signedTransaction := struct {
		SenderPublicKey    string
		RecipientPublicKey string
		Amount             float32
	}{t.SenderPublicKey, t.RecipientPublicKey, t.Amount}
	hashTx := SHA256(signedTransaction)

	rInt, sInt := signatureHexToIntPair(t.Signature)

	publicKeyBytes, err := hex.DecodeString(t.SenderPublicKey)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	publicKeyECDSA := ToECDSAPub(publicKeyBytes)
	res := ecdsa.Verify(publicKeyECDSA, ToBytes(hashTx), rInt, sInt)
	return res
}
