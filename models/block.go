package models

import (
	"crypto/sha256"
)

type Block struct {
	// Header
	PreviousBlockHash [32]byte `json:"previousHash"`
	MerkleRootHash    [32]byte `json:"merkleRoot"`
	Timestamp         int64    `json:"timestamp"`
	NBits             uint8    `json:"nBits"`
	Nonce             string   `json:"nonce"`
	// Transactions
	Transactions []Transaction `json:"Transactions"`
}

func (b *Block) GetHash() [32]byte {
	hash := sha256.Sum256([]byte("hello world\n"))
	return hash
}
