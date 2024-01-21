package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Cryptoblock struct {
	Hash []byte
	Data []byte
	PrevHash []byte	
}


func (c *Cryptoblock) BuildHash() {
	details := bytes.Join([][] byte{c.Data, c.PrevHash}, []byte{})
	hash := sha256.Sum256(details)
	c.Hash = hash[:]
}

func BuildBlock(data string, prevHash []byte) *Cryptoblock {
	block := &Cryptoblock{[]byte{}, []byte(data), prevHash}
	block.BuildHash()
	return block
}

type Blockchain struct {
	blocks []*Cryptoblock
}

func (chain *Blockchain) Addblock(data string) {
	prevBlock := chain.blocks[len(chain.blocks) - 1]
	new := BuildBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

// Inception() can also be termed as the Genesis block
func Inception() *Cryptoblock {
	return BuildBlock("Inception ", []byte{})
}

func InitBlockChain() *Blockchain {
	return &Blockchain{[]*Cryptoblock{Inception()}}
}

func main()  {
	chain := InitBlockChain()

	chain.Addblock("First block after Inception")
	chain.Addblock("Second block after Inception")
	chain.Addblock("Third block after Inception")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)	
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}