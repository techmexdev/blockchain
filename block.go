package blockchain

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"time"

	"github.com/cbergoon/merkletree"
)

// Block represents a
// bitcoin blockchain block
type Block struct {
	Header
	Transactions []Transaction
}

// Header represents a block's header
type Header struct {
	Timestamp     time.Time
	Target        int
	Nonce         int
	HashPrevBlock string
	// MerkleRoot is the hash of the block's
	// transactions merkle tree root
	MerkleRoot string
}

// MerkleRoot creates a merkle root from
// transactions
func MerkleRoot(txs []Transaction) ([]byte, error) {
	t, err := merkletree.NewTree(merkleTree.Content(txs))
	if err != nil {
		return nil, err
	}
	return t.MerkleRoot(), nil
}

// Validate validates the block is valid
func (b Block) Validate() error {
	if b.Header.MerkleRoot+b.Header.Nonce >= target {
		return errors.New("invalid nonce: %v")
	}
}

// Hash returns the hashed header
func (h Header) Hash() []byte {
	sha256.New().Sum(bytes.Join([][]byte{
		h.Timestamp, h.Target, h.Nonce, h.HashPrevBlock, h.MerkleRoot,
	}, []byte{}))
}
