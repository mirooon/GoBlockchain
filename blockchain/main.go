package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func newTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" || r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		d := json.NewDecoder(r.Body)
		d.DisallowUnknownFields() // catch unwanted fields

		transaction := Transaction{}

		err := d.Decode(&transaction)
		fmt.Printf("%v\n", err)
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
		blockchain.ResolveConflictsBetweenNodes()
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

func mine(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" || r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		block := blockchain.Mine()
		response := struct {
			Message      string
			BlockNumber  int
			Transactions []Transaction
			Nonce        int
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

func getChain(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" || r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		chain := blockchain.Chain
		response := struct {
			Chain  []Block
			Length int
		}{
			chain,
			len(chain),
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

func getNodes(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" || r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		nodes := blockchain.Neighbours
		response := struct {
			Nodes []string
		}{
			nodes,
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

func registerNode(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" || r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		d := json.NewDecoder(r.Body)
		d.DisallowUnknownFields()

		var request struct {
			Node string
		}
		err := d.Decode(&request)
		if err != nil {
			// bad JSON or unrecognized json field
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		for _, v := range blockchain.Neighbours {
			if v == request.Node {
				// Found!
			}
		}
		message := "Node successfuly added!"
		if blockchain.AddNeigbourIfNotExist(request.Node) {
			message = "Node already exists!"
		}

		fmt.Printf("%v\n", "Current neighbours")
		fmt.Printf("%v\n", blockchain.Neighbours)
		response := struct {
			Message           string
			AllFollowingNodes []string
		}{
			message,
			blockchain.Neighbours,
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
	mux := http.NewServeMux()
	mux.HandleFunc("/transaction/new", newTransaction)
	mux.HandleFunc("/transactions", getTransactions)
	mux.HandleFunc("/mine", mine)
	mux.HandleFunc("/chain", getChain)
	mux.HandleFunc("/nodes", getNodes)
	mux.HandleFunc("/node/new", registerNode)
	log.Printf("Listening on 5001")
	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":5001", handler))
}
