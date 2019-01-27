package main

import "fmt"

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 btc to Tanohzana")
	bc.AddBlock("Tanohzana withdraws 1 btc")

	for _, block := range bc.blocks {
		fmt.Printf("Prev hash %x:\n", block.prevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Data: %s\n\n", block.Data)
	}
}
