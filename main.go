package main

import (
	"fmt"
	"strconv"
	"github.com/wormholecowboy/go-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockchain()
	chain.AddBlock("second")
	chain.AddBlock("third")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

    pow := blockchain.NewProof(block)
    fmt.Printf("Proof: %s\n", strconv.FormatBool(pow.Validate()))
    fmt.Println()
	}
}
