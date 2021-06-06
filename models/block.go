package models

import (
	"crypto/sha256"
	"fmt"
)

type header struct {
	PreviousBlockHash [32]byte `json:"previousHash"`
	MerkleRootHash    [32]byte `json:"merkleRoot"`
	Timestamp         int64    `json:"timestamp"`
	NBits             uint8    `json:"nBits"`
	Nonce             string   `json:"nonce"`
}

type Block struct {
	// Use promoted fields for header fields
	header
	Transactions []Transaction `json:"Transactions"`
}

func calculateHashOfStruct(myStruct interface{}) [32]byte {
	// "%+v" to conserve the field names
	serializedHash := fmt.Sprintf("%+v", myStruct)
	return sha256.Sum256([]byte(serializedHash))
}

func (b *Block) GetHash() [32]byte {
	b.MerkleRootHash = calculateHashOfStruct(b.Transactions)
	return calculateHashOfStruct(b.header)
}
