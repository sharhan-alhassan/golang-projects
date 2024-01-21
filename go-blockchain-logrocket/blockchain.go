

package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// block Struct that will make up our blockchain
type Block struct {
	data 			map[string]interface{}
	hash 			string
	previousHash 	string
	timestamp 		time.Time
	pow 			int
}

// Blockchain type that contains our blocks
type Blockchain struct {
	genesisBlock 	Block
	chain			[]Block
	difficulty		int
}

// Let’s create a method on our Block type that generates a hash:
func (b Block) calculateHash() string {
	data, _ 		:= json.Marshal(b.data)
	blockData 		:= b.previousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow)
	blockHash 		:= sha256.Sum256([]byte(blockData))

	return fmt.Sprintf("%x", blockHash)
}

// A mine() method for our Block type that repeatedly increments the PoW value and calculates the block hash until we get a valid one
func (b *Block) mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.hash = b.calculateHash()
	}
}


// Initiate Blockchian creation
func NewBlockchain(difficulty int) Blockchain {
	genesisBlock := Block {
		hash: "0",
		timestamp: time.Now(),
	}

	return Blockchain {
		genesisBlock,
		[]Block{genesisBlock},
		difficulty,
	}
}


// Add Block instance to the Blockchain
// Create an addBlock method to the Blockchain type
func (blc *Blockchain) addBlock(from string, to string, amount float64) {
	// Construct block data 
	blockData := map[string]interface{}{
		"from": from,
		"to": to,
		"amount": amount,
	}

	// Fetch last block from Blockchain to extract its hash
	lastBlock := blc.chain[len(blc.chain)-1]

	newBlock := Block{
		data: blockData,
		previousHash: lastBlock.hash,
		timestamp: time.Now(),
	}	
	
	// Mine the block using the blockchain difficulty value
	newBlock.mine(blc.difficulty)

	// Add the new block to the block chain if the mining is successful
	blc.chain = append(blc.chain, newBlock)
}


// functionality that checks if the blockchain is valid
func (b Blockchain) isValid() bool {
	for i := range b.chain[1:] {
		previousBlock := b.chain[i]
		currentBlock := b.chain[i+1]
		if currentBlock.hash != currentBlock.calculateHash() || currentBlock.previousHash != previousBlock.hash {
			return false
		}
	}
	return true
}

