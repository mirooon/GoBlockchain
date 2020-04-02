package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"bytes"
	"io/ioutil"
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
	panic(err.Error)
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

type SubmitTransactionRequest struct {
	SenderPublicKey string `json:"senderPublicKey"`
	RecipientPublicKey string `json:"recipientPublicKey"`
	Signature string `json:"signature"`
	Amount float32 `json:"amount"`
	BlockchainNode string `json:"blockchainNode"`
}

func submitTransaction(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields() // catch unwanted fields

	submitTransactionRequest := SubmitTransactionRequest{}

	err := d.Decode(&submitTransactionRequest)
	if err != nil {
		// bad JSON or unrecognized json field
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if d.More() {
		http.Error(w, "extraneous data after JSON object", http.StatusBadRequest)
		return
	}

	log.Printf("%+v\n", "submitTransactionRequest")
	log.Printf("%+v\n", submitTransactionRequest)
	transactionJSON, err := json.Marshal(submitTransactionRequest)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	resp, err := http.Post("http://" + submitTransactionRequest.BlockchainNode + "/transaction/submit", "application/json", bytes.NewBuffer(transactionJSON))
	if err != nil {
		fmt.Printf("%v\n", "Problem with connection with node: " + submitTransactionRequest.BlockchainNode)
		fmt.Printf("%v\n", "err")
		fmt.Printf("%v\n", err)
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	log.Printf("%+v\n", "responseJSON")
	log.Printf("%+v\n", resp)
	// data := fmt.Sprintf(responseJSON)
	w.Write([]byte(body))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/wallet/generate", generateWallet)
	mux.HandleFunc("/transaction/create", createTransaction)
	mux.HandleFunc("/transaction/submit", submitTransaction)
	log.Printf("Listening on port 8080")
	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
