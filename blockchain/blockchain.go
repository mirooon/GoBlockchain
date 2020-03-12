package main

import (
	"crypto/sha256"
	"fmt"
)

type Blockchain struct {
	Transactions []Transaction
	Chain        []Block
}

func NewBlockchain() Blockchain {
	b := new(Blockchain)
	fmt.Println("New blockchain created!")
	genesisBlock := CreateGenesisBlock()
	b.Chain = append(b.Chain, genesisBlock)
	return *b
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := NewBlock(data, len(bc.Chain), asSha256(prevBlock), bc.Transactions)
	bc.Transactions = nil
	bc.Chain = append(bc.Chain, newBlock)
	fmt.Println("Block added!")
}

func (bc *Blockchain) AddTransaction(transaction Transaction) {
	bc.Transactions = append(bc.Transactions, transaction)
	fmt.Println("Transaction added!")
}

func asSha256(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))

	return fmt.Sprintf("%x", h.Sum(nil))
}
