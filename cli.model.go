package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type CLI struct{}

// CLI should check that blockchain exists before calling other functions
func (cli *CLI) Run() {
	cli.validateArgs()

	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)

	createBlockchainData := createBlockchainCmd.String("address", "", "Genesis mining reward address")
	getBalanceData := getBalanceCmd.String("address", "", "Address of which the balance is fetched")

	switch os.Args[1] {
	case "printchain":
		printChainCmd.Parse(os.Args[2:])

	case "createblockchain":
		createBlockchainCmd.Parse(os.Args[2:])

	case "getbalance":
		getBalanceCmd.Parse(os.Args[2:])

	default:
		cli.printUsage()
		os.Exit(1)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}

	if createBlockchainCmd.Parsed() {
		cli.createBlockchain(*createBlockchainData)
	}

	if getBalanceCmd.Parsed() {
		cli.getBalance(*getBalanceData)
	}
}

func (cli *CLI) printChain() {
	bc := NewBlockchain("")
	defer bc.db.Close()
	bci := bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("Prev Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := NewProofOfWork(block)

		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

func (cli *CLI) createBlockchain(address string) {
	bc := CreateBlockchain(address)
	bc.db.Close()
	fmt.Printf("Blockchain created ! Printing blockchain...")
	cli.printChain()
}

func (cli *CLI) getBalance(address string) {
	bc := NewBlockchain(address)
	defer bc.db.Close()

	balance := 0
	UTXOs := bc.FindUTXO(address)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}

// This function could be improved
func (cli *CLI) validateArgs() {
	args := os.Args

	if len(args) < 2 {
		os.Exit(1)
	}
}

func (cli *CLI) printUsage() {
	fmt.Printf("This is the way cli should be used: blah blah blah")
}
