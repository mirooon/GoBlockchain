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
		data := fmt.Sprintf(`{"verifyResult": "%B"}`, res)
		w.Write([]byte(data))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main() {

	// blockchain := NewBlockchain()
	// blockchain.AddBlock("Second Block!")
	// fmt.Printf("%+v\n", blockchain)
	// fmt.Printf("%+v\n", *blockchain.chain[1])
	http.HandleFunc("/transaction/new", newTransaction)
	log.Printf("Listening on port 5001")
	log.Fatal(http.ListenAndServe(":5001", nil))
}
