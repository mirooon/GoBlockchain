package main

import (
	"fmt"
	"time"
)

type Block struct {
	BlockNumber  int
	Timestamp    int64
	Transactions []Transaction
	Nonce        string
	PreviousHash string
}

func CreateGenesisBlock() Block {
	b := new(Block)
	b.BlockNumber = 0
	b.Timestamp = time.Now().Unix()
	// b.Transactions = []
	b.PreviousHash = "0000000000000000000000000000000000000000000000000000000000000000"
	fmt.Println("Genesis block created!")
	return *b
}

func NewBlock(blockNumber int, timestamp time.Time, transactions []Transaction, nonce string, previousHash string) Block {
	b := new(Block)
	b.BlockNumber = blockNumber
	b.Timestamp = timestamp.Unix()
	b.Transactions = transactions
	b.Nonce = nonce
	b.PreviousHash = previousHash
	fmt.Println("New block created!")
	return *b
}

func (b *Block) ShowBlockInfo() {
	fmt.Printf("%+v\n", b)
}
