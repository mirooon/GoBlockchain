package main

import (
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
	fmt.Printf("%s\n", wallet.GetHexPrivateKey())
	fmt.Printf("%s\n", wallet.GetHexPublicKey())

	data := fmt.Sprintf(`{"publicKey": "%s", "privateKey": "%s"}`, wallet.GetHexPrivateKey(), wallet.GetHexPublicKey())
	w.Write([]byte(data))
}

func main() {
	http.HandleFunc("/wallet/generate", generateWallet)
	log.Printf("Listening on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
