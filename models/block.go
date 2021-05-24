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
	Header       header
	Transactions []Transaction `json:"Transactions"`
}

func calculateHashOfStruct(myStruct interface{}) [32]byte {
	// "%+v" to conserve the field names
	return sha256.Sum256([]byte(fmt.Sprintf("%+v", myStruct)))
}

func (b *Block) GetHash() [32]byte {
	b.Header.MerkleRootHash = calculateHashOfStruct(b.Transactions)
	hash := calculateHashOfStruct(*&b.Header)
	return hash
}
