package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func newTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" || r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		enableCors(&w)

		d := json.NewDecoder(r.Body)
		d.DisallowUnknownFields() // catch unwanted fields

		transaction := Transaction{}

		err := d.Decode(&transaction)
		if err != nil {
			// bad JSON or unrecognized json field
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if d.More() {
			http.Error(w, "extraneous data after JSON object", http.StatusBadRequest)
			return
		}

		res := transaction.VerifyTransaction()
		data := fmt.Sprintf(`{"verifyResult": "%t"}`, res)
		if res {
			blockchain.AddTransaction(transaction)
		}
		w.Write([]byte(data))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func getTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		enableCors(&w)
		fmt.Printf("blockchain.Transactions\n")
		fmt.Printf("%v\n", blockchain.Transactions)
		jsonTransactions, err := json.Marshal(blockchain.Transactions)
		if err != nil {
			// bad JSON or unrecognized json field
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write([]byte(jsonTransactions))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

var MINING_SENDER string = "The node"
var MINING_REWARD int = 1

func mine(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" || r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		enableCors(&w)
		nonce := blockchain.ProofOfWork()
		blockchain.SubmitTransaction(MINING_SENDER, blockchain.NodeId, "", MINING_REWARD)

		lastBlock := blockchain.Chain[len(blockchain.Chain)-1]
		previousHash := blockchain.Hash(lastBlock)
		block := blockchain.AddBlock(nonce, previousHash)
		fmt.Printf("%v\n", block)
		response := struct {
			Message      string
			BlockNumber  int
			Transactions []Transaction
			Nonce        string
			PreviousHash string
		}{
			"Block created!",
			block.BlockNumber,
			block.Transactions,
			block.Nonce,
			block.PreviousHash,
		}

		jsResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(jsResponse))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

var blockchain Blockchain

func main() {

	blockchain = NewBlockchain()
	// blockchain.AddBlock("Second Block!")
	// fmt.Printf("%+v\n", blockchain)
	// fmt.Printf("%+v\n", *blockchain.chain[1])
	http.HandleFunc("/transaction/new", newTransaction)
	http.HandleFunc("/transactions", getTransactions)
	http.HandleFunc("/mine", mine)
	log.Printf("Listening on port 5001")
	log.Fatal(http.ListenAndServe(":5001", nil))
}
