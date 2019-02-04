package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func (b *Block) setHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.HashTransactions(), b.PrevBlockHash, timestamp}, []byte{})

	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

func NewBlock(transactions []*Transaction, PrevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), transactions, PrevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)

	nonce, hash := pow.Run()

	block.Nonce = nonce
	block.Hash = hash[:]

	return block
}

func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encode := gob.NewEncoder(&result)

	encode.Encode(b)

	return result.Bytes()
}

func (block *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range block.Transactions {
		txHashes = append(txHashes, tx.ID)
	}

	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}

func Deserialize(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	decoder.Decode(&block)

	return &block
}
