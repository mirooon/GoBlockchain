package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
)

type Wallet struct {
	PrivateKey ecdsa.PrivateKey `json:"privateKey"`
	PublicKey  ecdsa.PublicKey  `json:"publicKey"`
}

func (wallet *Wallet) GetHexPrivateKey() string {
	privateKeyBytes := wallet.PrivateKey.D.Bytes()
	encodedPrivateKeyStr := hex.EncodeToString(privateKeyBytes)
	return encodedPrivateKeyStr
}

func (wallet *Wallet) GetHexPublicKey() string {
	publicKeyBytes := elliptic.Marshal(wallet.PublicKey, wallet.PublicKey.X, wallet.PublicKey.Y)
	encodedPublicKeyStr := hex.EncodeToString(publicKeyBytes)
	return encodedPublicKeyStr
}

func NewKeyPair() (ecdsa.PrivateKey, ecdsa.PublicKey) {
	pubkeyCurve := elliptic.P256() //see http://golang.org/pkg/crypto/elliptic/#P256

	privatekey := new(ecdsa.PrivateKey)
	privatekey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader) // this generates a public & private key pair

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var pubkey ecdsa.PublicKey
	pubkey = privatekey.PublicKey

	return *privatekey, pubkey
}

func MakeWallet() *Wallet {
	private, public := NewKeyPair()
	wallet := Wallet{private, public}
	return &wallet
}
