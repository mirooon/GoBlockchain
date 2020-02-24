package main

import (
	"crypto/sha256"
	"fmt"
)

type Blockchain struct {
	transactions []*Transaction
	chain        []*Block
}

func NewBlockchain() *Blockchain {
	b := new(Blockchain)
	fmt.Println("New blockchain created!")
	genesisBlock := CreateGenesisBlock()
	b.chain = append(b.chain, genesisBlock)
	return b
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.chain[len(bc.chain)-1]
	newBlock := NewBlock(data, len(bc.chain), asSha256(prevBlock), bc.transactions)
	bc.transactions = nil
	bc.chain = append(bc.chain, newBlock)
	fmt.Println("Block added!")
}

func asSha256(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))

	return fmt.Sprintf("%x", h.Sum(nil))
}
