package models

import (
	"github.com/google/uuid"
)

type Transaction struct {
	Id               uuid.UUID `json:"id"`
	RecipientAddress Address   `json:"recipientAddress"`
	SenderAddress    Address   `json:"senderAddress"`
	Value            float32   `json:"value"`
}

// Alterantive: use factory function and make structs private
// http://www.golangpatterns.info/object-oriented/constructors
func (t *Transaction) IsValid() bool {
	return t.Id != uuid.Nil && t.SenderAddress.IsValid() && t.RecipientAddress.IsValid() && t.Value > 0
}
