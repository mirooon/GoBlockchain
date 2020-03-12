package main

import (
	"fmt"
	"time"
)

type Block struct {
	BlockNumber  int
	Data         string
	Timestamp    time.Time
	Transactions []Transaction
	PreviousHash string
}

func CreateGenesisBlock() Block {
	b := new(Block)
	b.BlockNumber = 0
	b.Data = "Genesis Block"
	b.Timestamp = time.Now()
	// b.Transactions = []
	b.PreviousHash = "0000000000000000000000000000000000000000000000000000000000000000"
	fmt.Println("Genesis block created!")
	return *b
}

func NewBlock(data string, blockNumber int, previousHash string, transactions []Transaction) Block {
	b := new(Block)
	b.BlockNumber = blockNumber
	b.Data = data
	b.Timestamp = time.Now()
	b.Transactions = transactions
	b.PreviousHash = previousHash
	fmt.Println("New block created!")
	return *b
}

func (b *Block) ShowBlockInfo() {
	fmt.Printf("%+v\n", b)
}
