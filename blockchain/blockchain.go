package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
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

func (bc *Blockchain) AddBlock(nonce int, prevBlockHash string) Block {
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

var MINING_DIFFICULTY int = 1

func (bc *Blockchain) ValidProof(lastBlockHash string, nonce int) bool {
	var guess strings.Builder
	guess.WriteString(string(asSha256(bc.Transactions)))
	guess.WriteString(lastBlockHash)
	guess.WriteString(string(nonce))
	guessHash := asSha256(guess.String())
	return strings.HasPrefix(guessHash, strings.Repeat("0", MINING_DIFFICULTY))
}

func (bc *Blockchain) ProofOfWork() int {
	lastBlockHash := asSha256(bc.Chain[len(bc.Chain)-1])
	nonce := 0
	for bc.ValidProof(lastBlockHash, nonce) == false {
		nonce += 1
		fmt.Printf("%v\n", "nonce")
		fmt.Printf("%v\n", nonce)
	}
	return nonce
}

var MINING_SENDER string = "The node"
var MINING_REWARD float32 = 1

func (bc *Blockchain) Mine() Block {
	nonce := bc.ProofOfWork()
	fmt.Printf("%v\n", "nonce")
	fmt.Printf("%v\n", nonce)

	bc.AddRewardTransaction(MINING_SENDER, bc.NodeId, "", MINING_REWARD)

	lastBlock := bc.Chain[len(bc.Chain)-1]
	previousHash := bc.Hash(lastBlock)
	return bc.AddBlock(nonce, previousHash)
}

func (bc *Blockchain) AddRewardTransaction(senderPublicKey string, recipientPublicKey string, signature string, amount float32) {
	transaction := Transaction{
		SenderPublicKey:    senderPublicKey,
		RecipientPublicKey: recipientPublicKey,
		Signature:          signature,
		Amount:             amount,
	}
	if senderPublicKey == MINING_SENDER {
		bc.AddTransaction(transaction)
	} else if transaction.VerifyTransaction() {
		bc.AddTransaction(transaction)
	}
	fmt.Println("Reward transaction added!")
}

func (bc *Blockchain) Hash(block Block) string {
	return asSha256(block)
}

func asSha256(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))

	return fmt.Sprintf("%x", h.Sum(nil))
}

func proofOfWork() string {
	return "1234"
}
