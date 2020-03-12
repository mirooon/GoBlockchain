package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"log"
	// "github.com/ethereum/go-ethereum/crypto/sha3"
)

type Wallet struct {
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
}

// func (wallet *Wallet) GetHexPrivateKey() string {
// 	privateKeyBytes := wallet.PrivateKey.D.Bytes()
// 	encodedPrivateKeyStr := hex.EncodeToString(privateKeyBytes)
// 	return encodedPrivateKeyStr
// }

// func (wallet *Wallet) GetHexPublicKey() string {
// 	publicKeyBytes := elliptic.Marshal(wallet.PublicKey, wallet.PublicKey.X, wallet.PublicKey.Y)
// 	encodedPublicKeyStr := hex.EncodeToString(publicKeyBytes)
// 	return encodedPublicKeyStr
// }

// func HexKeysToWallet(publicKey string, privateKey string) Wallet {
// 	decodedPublicKeyBytes := hex.DecodeString(publicKey)
// 	ecdsaPublicKey := elliptic.Unmarshal()
// }

func NewKeyPair() (string, string) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := privateKey.D.Bytes()
	privateKeyHex := hex.EncodeToString(privateKeyBytes)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	publicKeyBytes := FromECDSAPub(publicKeyECDSA)
	publicKeyHex := hex.EncodeToString(publicKeyBytes)

	return privateKeyHex, publicKeyHex
}

func MakeWallet() *Wallet {
	private, public := NewKeyPair()
	wallet := Wallet{private, public}
	return &wallet
}
