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

func createValidBlock() models.Block {
	var transactions []models.Transaction
	transactions = append(transactions, createValidTransaction())
	transactions = append(transactions, createValidTransaction())
	transactions = append(transactions, createValidTransaction())

	var b models.Block
	b.PreviousBlockHash = sha256.Sum256([]byte("I am the previous block's header\n"))
	b.Timestamp = time.Now().UnixNano()
	b.NBits = 1
	b.Nonce = nonce.NewToken()
	b.Transactions = transactions

	return b
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
		transactionCopy := transaction
		transactionCopy.Id = uuid.Nil

		if transactionCopy.IsValid() {
			t.Error("Expected transaction to be invalid (Id is 'nil')")
		}
	})

	t.Run("Check SenderAddress", func(t *testing.T) {
		transactionCopy := transaction

		var invalidTransactions = make(map[string]models.Transaction)
		transactionCopy.SenderAddress = models.Address{}
		invalidTransactions["nil"] = transactionCopy
		transactionCopy.SenderAddress = models.Address{Street: "FooStreet", HouseNumber: "1", Town: "FooTown"}
		invalidTransactions["missing 'Name'"] = transactionCopy
		transactionCopy.SenderAddress = models.Address{Name: "Foo", HouseNumber: "1", Town: "FooTown"}
		invalidTransactions["missing 'Street'"] = transactionCopy
		transactionCopy.SenderAddress = models.Address{Name: "Foo", Street: "FooStreet", Town: "FooTown"}
		invalidTransactions["missing 'HouseNumber'"] = transactionCopy
		transactionCopy.SenderAddress = models.Address{Name: "Foo", Street: "FooStreet", HouseNumber: "1"}
		invalidTransactions["missing 'Town'"] = transactionCopy

		for key, invalidTransaction := range invalidTransactions {
			if invalidTransaction.IsValid() {
				t.Errorf("Expected transaction to be invalid (SenderAddress is %s)", key)
			}
		}
	})

	t.Run("Check RecipientAddress", func(t *testing.T) {
		transactionCopy := transaction

		var invalidTransactions = make(map[string]models.Transaction)
		transactionCopy.RecipientAddress = models.Address{}
		invalidTransactions["nil"] = transactionCopy
		transactionCopy.RecipientAddress = models.Address{Street: "FooStreet", HouseNumber: "1", Town: "FooTown"}
		invalidTransactions["missing 'Name'"] = transactionCopy
		transactionCopy.RecipientAddress = models.Address{Name: "Foo", HouseNumber: "1", Town: "FooTown"}
		invalidTransactions["missing 'Street'"] = transactionCopy
		transactionCopy.RecipientAddress = models.Address{Name: "Foo", Street: "FooStreet", Town: "FooTown"}
		invalidTransactions["missing 'HouseNumber'"] = transactionCopy
		transactionCopy.RecipientAddress = models.Address{Name: "Foo", Street: "FooStreet", HouseNumber: "1"}
		invalidTransactions["missing 'Town'"] = transactionCopy

		for key, invalidTransaction := range invalidTransactions {
			if invalidTransaction.IsValid() {
				t.Errorf("Expected transaction to be invalid (RecipientAddress is %s)", key)
			}
		}
	})

	t.Run("Check Value", func(t *testing.T) {
		transactionCopy := transaction

		var invalidTransactions = make(map[string]models.Transaction)
		transactionCopy.Value = 0
		invalidTransactions["null"] = transactionCopy
		transactionCopy.Value = -100.21
		invalidTransactions["negative"] = transactionCopy

		for key, invalidTransaction := range invalidTransactions {
			if invalidTransaction.IsValid() {
				t.Errorf("Expected transaction to be invalid (Value is '%s')", key)
			}
		}
	})
}

func TestBlockGetHash(t *testing.T) {
	block := createValidBlock()

	t.Run("It should calculate the merkle root hash", func(t *testing.T) {
		_ = block.GetHash()

		var emptyHash [32]byte
		if block.MerkleRootHash == emptyHash {
			t.Error("Expected merkle root hash to be calculated")
		}
	})

	// TODO: check that calculated hash is not zero

	t.Run("It should return a different hash if the nonce changes", func(t *testing.T) {
		hash1 := block.GetHash()

		block.Nonce = nonce.NewToken()
		hash2 := block.GetHash()

		if hash1 == hash2 {
			t.Error("Expected hash to be different if nonce has changed")
		}
	})
}

func TestValidBlock(t *testing.T) {
	block := createValidBlock()
	if !block.IsValid() {
		t.Error("Expected block to be valid")
	}
}

func TestInvalidBlock(t *testing.T) {
	block := createValidBlock()

	t.Run("Check PreviousBlockHash", func(t *testing.T) {
		invalidBlock := block
		invalidBlock.PreviousBlockHash = [32]byte{}

		if invalidBlock.IsValid() {
			t.Error("Expected block to be invalid (PreviousBlockHash is zero)")
		}
	})

	t.Run("Check Timestamp", func(t *testing.T) {
		invalidBlock := block
		invalidBlock.Timestamp = 0

		if invalidBlock.IsValid() {
			t.Error("Expected block to be invalid (Timestamp is zero)")
		}
	})

	t.Run("Check NBits", func(t *testing.T) {
		invalidBlock := block
		invalidBlock.NBits = 0

		if invalidBlock.IsValid() {
			t.Error("Expected block to be invalid (NBits is zero)")
		}
	})

	t.Run("Check Nonce", func(t *testing.T) {
		invalidBlock := block
		invalidBlock.Nonce = ""

		if invalidBlock.IsValid() {
			t.Error("Expected block to be invalid (Nonce is zero)")
		}
	})

	t.Run("Check Transactions", func(t *testing.T) {
		invalidBlock := block
		invalidBlock.Transactions = []models.Transaction{}

		if invalidBlock.IsValid() {
			t.Error("Expected block to be invalid (Transactions is zero)")
		}
	})
}
