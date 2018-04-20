package blockchain

import (
	"github.com/cbergoon/merkletree"
)

// Transaction implements merkletree.Content
type Transaction struct {
	From string
	To string
	Amount int
}

// CalculateHash hashes the values of the transaction
func (t Transaction) CalculateHash() []byte {
	h := sha256.New()
	h.Sum([]byte(t.From+t.To+t.Amount))
}

// Equals tests for equality of two transactions
func (t TestContent) Equals(other merkletree.Content) bool {
	return (
		t.From == other.(Transaction).From &&
		t.To == other.(Transaction).To &&
		t.Amount == other.(Transaction).Amount
	)
}

