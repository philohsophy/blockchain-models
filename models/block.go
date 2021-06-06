package models

import (
	"crypto/sha256"
	"fmt"
	"reflect"
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

func (b *Block) IsValid() bool {
	// Ref: https://stackoverflow.com/a/45222521
	// Using reflect.DeepEqual instead of cmp.Equal
	// since this is a custom struct that is comparable via reflect
	// and to avoid having to install the cmp module
	return b.PreviousBlockHash != [32]byte{} && b.Timestamp != 0 && b.NBits != 0 && b.Nonce != "" && !reflect.DeepEqual(b.Transactions, []Transaction{})
}
