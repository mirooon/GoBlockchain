package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"fmt"
	"math/big"
)

func ToECDSAPub(pub []byte) *ecdsa.PublicKey {
	if pub == nil || len(pub) == 0 {
		return nil
	}
	x, y := elliptic.Unmarshal(elliptic.P256(), pub)
	fmt.Printf("x\n")
	fmt.Printf("%+v\n", x)
	fmt.Printf("y\n")
	fmt.Printf("%+v\n", y)
	publicKey := ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}
	return &publicKey
}

func signatureHexToIntPair(signature string) (*big.Int, *big.Int) {
	rInt := new(big.Int)
	rInt.SetString(signature[:64], 16)
	sInt := new(big.Int)
	sInt.SetString(signature[64:], 16)
	return rInt, sInt
}
