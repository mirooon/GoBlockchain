package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func generateWallet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	wallet := MakeWallet()

	data := fmt.Sprintf(`{"publicKey": "%s", "privateKey": "%s"}`, wallet.PublicKey, wallet.PrivateKey)
	w.Write([]byte(data))
}

type CreateTransactionResponse struct {
	Transaction Transaction
	Signature   string
}

func createTransaction(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
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

	log.Printf("%+v\n", transaction)
	response := CreateTransactionResponse{
		Transaction: transaction,
		Signature:   transaction.SignTransaction(),
	}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	// data := fmt.Sprintf(responseJSON)
	w.Write([]byte(responseJSON))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/wallet/generate", generateWallet)
	mux.HandleFunc("/transaction/create", createTransaction)
	log.Printf("Listening on port 8080")
	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
