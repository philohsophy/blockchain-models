package models_test

import (
	"crypto/sha256"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/LarryBattle/nonce-golang"
	"github.com/google/uuid"
	"github.com/philohsophy/blockchain-models/models"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func createValidTransaction() models.Transaction {
	recipientAddress := models.Address{Name: "Foo", Street: "FooStreet", HouseNumber: "1", Town: "FooTown"}
	senderAddress := models.Address{Name: "Bar", Street: "BarStreet", HouseNumber: "1", Town: "BarTown"}

	var t models.Transaction
	t.Id = uuid.New()
	t.RecipientAddress = recipientAddress
	t.SenderAddress = senderAddress

	rand.Seed(time.Now().UnixNano())
	t.Value = float32(rand.Intn(999-1)+1) + 0.5
	return t
}

func TestValidTransaction(t *testing.T) {
	transaction := createValidTransaction()
	if !transaction.IsValid() {
		t.Error("Expected transaction to be valid")
	}
}

func TestInvalidTransaction(t *testing.T) {
	transaction := createValidTransaction()

	t.Run("Check Id", func(t *testing.T) {
		invalidTransaction := models.Transaction{SenderAddress: transaction.SenderAddress, RecipientAddress: transaction.RecipientAddress, Value: transaction.Value}

		if invalidTransaction.IsValid() {
			t.Error("Expected transaction to be invalid (Id is 'nil')")
		}
	})

	t.Run("Check SenderAddress", func(t *testing.T) {
		transaction := models.Transaction{Id: transaction.Id, RecipientAddress: transaction.RecipientAddress, Value: transaction.Value}

		var invalidTransactions = make(map[string]models.Transaction)
		invalidTransactions["nil"] = transaction
		transaction.SenderAddress = models.Address{Street: "FooStreet", HouseNumber: "1", Town: "FooTown"}
		invalidTransactions["missing 'Name'"] = transaction
		transaction.SenderAddress = models.Address{Name: "Foo", HouseNumber: "1", Town: "FooTown"}
		invalidTransactions["missing 'Street'"] = transaction
		transaction.SenderAddress = models.Address{Name: "Foo", Street: "FooStreet", Town: "FooTown"}
		invalidTransactions["missing 'HouseNumber'"] = transaction
		transaction.SenderAddress = models.Address{Name: "Foo", Street: "FooStreet", HouseNumber: "1"}
		invalidTransactions["missing 'Town'"] = transaction

		for key, invalidTransaction := range invalidTransactions {
			if invalidTransaction.IsValid() {
				t.Errorf("Expected transaction to be invalid (RecipientAddress is %s)", key)
			}
		}
	})

	t.Run("Check RecipientAddress", func(t *testing.T) {
		transaction := models.Transaction{Id: transaction.Id, SenderAddress: transaction.SenderAddress, Value: transaction.Value}

		var invalidTransactions = make(map[string]models.Transaction)
		invalidTransactions["nil"] = transaction
		transaction.RecipientAddress = models.Address{Street: "FooStreet", HouseNumber: "1", Town: "FooTown"}
		invalidTransactions["missing 'Name'"] = transaction
		transaction.RecipientAddress = models.Address{Name: "Foo", HouseNumber: "1", Town: "FooTown"}
		invalidTransactions["missing 'Street'"] = transaction
		transaction.RecipientAddress = models.Address{Name: "Foo", Street: "FooStreet", Town: "FooTown"}
		invalidTransactions["missing 'HouseNumber'"] = transaction
		transaction.RecipientAddress = models.Address{Name: "Foo", Street: "FooStreet", HouseNumber: "1"}
		invalidTransactions["missing 'Town'"] = transaction

		for key, invalidTransaction := range invalidTransactions {
			if invalidTransaction.IsValid() {
				t.Errorf("Expected transaction to be invalid (RecipientAddress is %s)", key)
			}
		}
	})

	t.Run("Check Value", func(t *testing.T) {
		transaction := models.Transaction{Id: transaction.Id, SenderAddress: transaction.SenderAddress, RecipientAddress: transaction.RecipientAddress}

		var invalidTransactions = make(map[string]models.Transaction)
		invalidTransactions["nil"] = transaction
		transaction.Value = 0.00
		invalidTransactions["null"] = transaction
		transaction.Value = -100.21
		invalidTransactions["negative"] = transaction

		for key, invalidTransaction := range invalidTransactions {
			if invalidTransaction.IsValid() {
				t.Errorf("Expected transaction to be invalid (Value is '%s')", key)
			}
		}
	})
}

func TestBlockGetHash(t *testing.T) {
	var transactions []models.Transaction
	transactions = append(transactions, createValidTransaction())
	transactions = append(transactions, createValidTransaction())
	transactions = append(transactions, createValidTransaction())

	var block models.Block
	block.PreviousBlockHash = sha256.Sum256([]byte("I am the previous block's header\n"))
	block.Timestamp = time.Now().UnixNano()
	block.NBits = 1
	block.Nonce = nonce.NewToken()
	block.Transactions = transactions

	t.Run("It should calculate the merkle root hash", func(t *testing.T) {
		_ = block.GetHash()

		var emptyHash [32]byte
		if block.MerkleRootHash == emptyHash {
			t.Error("Expected merkle root hash to be calculated")
		}
	})

	t.Run("It should return a different hash if the nonce changes", func(t *testing.T) {
		hash1 := block.GetHash()

		block.Nonce = nonce.NewToken()
		hash2 := block.GetHash()

		if hash1 == hash2 {
			t.Error("Expected hash to be different if nonce has changed")
		}
	})
}
