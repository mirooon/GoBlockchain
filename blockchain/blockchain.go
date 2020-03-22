package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Blockchain struct {
	Transactions []Transaction
	Chain        []Block
	Neighbours   []string //nodes' ips
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
	// fmt.Println("Block added!")
	return newBlock
}

func (bc *Blockchain) AddTransaction(transaction Transaction) {
	bc.Transactions = append(bc.Transactions, transaction)
	// fmt.Println("Transaction added!")
}

var MINING_DIFFICULTY int = 1

func (bc *Blockchain) ValidChain(chain []Block) bool {
	currentBlock := chain[0]
	currentIndex := 1

	for currentIndex < len(chain) {
		block := chain[currentIndex]
		if block.PreviousHash != bc.Hash(currentBlock) {
			return false
		}

		if block.Transactions != nil && len(block.Transactions) > 0 {
			transactions := block.Transactions[:len(block.Transactions)-1] //resign from last (reward) transaction
			if bc.ValidProof(transactions, block.PreviousHash, block.Nonce) == false {
				return false
			}
		}
		currentBlock = block
		currentIndex++
	}
	return true
}
func (bc *Blockchain) ResolveConflictsBetweenNodes() bool {
	var currentChain []Block
	maxLength := len(bc.Chain)

	for _, ip := range bc.Neighbours {
		resp, err := http.Get("http://" + ip + "/chain")
		if err != nil {
			fmt.Printf("%v\n", "Problem with connection with ip: "+ip)
			continue
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%v\n", "Problem with connection with ip: "+ip)
			continue
		}
		var responseObj struct {
			Chain  []Block
			Length int
		}
		err = json.Unmarshal(body, &responseObj)
		if err != nil {
			panic(err)
		}
		length := responseObj.Length
		chain := responseObj.Chain
		if length > maxLength && bc.ValidChain(chain) {
			currentChain = chain
			maxLength = length
		}
	}

	if currentChain != nil {
		bc.Chain = currentChain
		return true
	}

	return false
}
func (bc *Blockchain) ValidProof(transactions []Transaction, lastBlockHash string, nonce int) bool {
	var guess strings.Builder

	guess.WriteString(asSha256(transactions))
	guess.WriteString(lastBlockHash)
	guess.WriteString(string(nonce))
	guessHash := asSha256(guess.String())
	return strings.HasPrefix(guessHash, strings.Repeat("0", MINING_DIFFICULTY))
}

func (bc *Blockchain) ProofOfWork(lastBlockHash string) int {
	nonce := 0
	for bc.ValidProof(bc.Transactions, lastBlockHash, nonce) == false {
		nonce += 1
	}

	return nonce
}

var MINING_SENDER string = "The blockchain"
var MINING_REWARD float32 = 1

func (bc *Blockchain) Mine() Block {
	lastBlockHash := bc.Hash(bc.Chain[len(bc.Chain)-1])
	nonce := bc.ProofOfWork(lastBlockHash)

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
}

func (bc *Blockchain) AddNeigbourIfNotExist(node string) bool {
	addNeighbour := true
	for _, v := range bc.Neighbours {
		if v == node {
			addNeighbour = false
		}
	}
	if addNeighbour {
		bc.Neighbours = append(bc.Neighbours, node)
		return true
	}
	return false
}

func (bc *Blockchain) Hash(block Block) string {
	return asSha256(block)
}

func asSha256(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))

	return fmt.Sprintf("%x", h.Sum(nil))
}
