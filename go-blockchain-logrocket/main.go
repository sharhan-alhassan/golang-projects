package main

import (
	"fmt"
)


func main()  {
	// create a new blockchain instance with a mining difficulty of 2
	blockchain := NewBlockchain(2)

	// record transactions on the blockchain for Alice, Bob, and John
	blockchain.addBlock("Alice", "Bob", 5)
	blockchain.addBlock("John", "Bob", 2)

	// check if the blockchain is valid; expecting true
	fmt.Println(blockchain.isValid())

}