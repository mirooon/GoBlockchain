package main

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Blockchain struct {
	Transactions []Transaction
	Chain        []Block
	NodeId       string
}

func NewBlockchain() Blockchain {
	b := new(Blockchain)
	b.NodeId = uuid.New().String()
	fmt.Println("New blockchain created!")
	genesisBlock := CreateGenesisBlock()
	b.Chain = append(b.Chain, genesisBlock)
	return *b
}

func (bc *Blockchain) AddBlock(nonce string, prevBlockHash string) Block {
	newBlock := NewBlock(len(bc.Chain), time.Now(), bc.Transactions, nonce, prevBlockHash)
	bc.Transactions = nil
	bc.Chain = append(bc.Chain, newBlock)
	fmt.Println("Block added!")
	return newBlock
}

func (bc *Blockchain) AddTransaction(transaction Transaction) {
	bc.Transactions = append(bc.Transactions, transaction)
	fmt.Println("Transaction added!")
}

func (bc *Blockchain) ProofOfWork() string {
	fmt.Println("Mining!")
	return "1234"
}

func (bc *Blockchain) SubmitTransaction(senderPublicKey string, recipientPublicKey string, signature string, amount int) {
	fmt.Println("Transaction added!")
}

func (bc *Blockchain) Hash(block Block) string {
	fmt.Println("Hash block!")
	return "abc"
}

func asSha256(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))

	return fmt.Sprintf("%x", h.Sum(nil))
}

func proofOfWork() string {
	return "1234"
}
