package main

import (
)

type TransactionRequest struct {
	SenderPublicKey    string `json:"senderPublicKey"`
	RecipientPublicKey string `json:"recipientPublicKey"`
	Signature   string `json:"signature"`
	Amount             float32 `json:"amount"`
}