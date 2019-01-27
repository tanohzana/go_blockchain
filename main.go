package main

import (
	"fmt"
	"strconv"
)

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 btc to Tanohzana")
	bc.AddBlock("Tanohzana withdraws 1 btc")

	for _, block := range bc.blocks {
		fmt.Printf("Prev hash %x:\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Data: %s\n\n", block.Data)

		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n\n", strconv.FormatBool(pow.Validate()))
	}
}
