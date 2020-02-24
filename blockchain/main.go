package main

import (
	"fmt"
)

func main() {

	blockchain := NewBlockchain()
	blockchain.AddBlock("Second Block!")
	fmt.Printf("%+v\n", blockchain)
	fmt.Printf("%+v\n", *blockchain.chain[1])
}
