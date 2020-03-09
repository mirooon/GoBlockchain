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

func generateWallet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	enableCors(&w)

	wallet := MakeWallet()

	data := fmt.Sprintf(`{"publicKey": "%s", "privateKey": "%s"}`, wallet.GetHexPrivateKey(), wallet.GetHexPublicKey())
	w.Write([]byte(data))
}

type CreateTransactionResponse struct {
	Transaction Transaction
	Signature   string
}

func createTransaction(w http.ResponseWriter, r *http.Request) {
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

	log.Printf("%+v\n", transaction)
	response := CreateTransactionResponse{
		Transaction: transaction,
		Signature:   signTransaction(transaction),
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
	http.HandleFunc("/wallet/generate", generateWallet)
	http.HandleFunc("/transaction/create", createTransaction)
	log.Printf("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
