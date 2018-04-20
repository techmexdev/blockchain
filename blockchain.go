package blockchain

import (
	"math"
	"time"
)

// Blockchain is accessed through
// public methods
type Blockchain struct {
	chain      []Block
	pendingTxs []Transaction
	targetBits int
}

// NewBlockchain creates a blockchain
// with 25 coins given to satoshi
func NewBlockchain() *Blockchain {
	txs := []Transaction{
		{From: "0000", To: "Satoshi", Amount: 25},
	}
	genesis := Block{
		Header: {
			Timestamp:     time.Now(),
			MerkleRoot:    MerkleRoot(txs),
			Nonce:         0000,
			HashPrevBlock: "0000",
			TargetBits:    24,
		},
		Transactions: txs,
	}

	return &Blockchain{
		chain: []Block{{genesis}},
	}
}

// GetBlockchain returns a copy of the blockchain
func (bc Blockchain) Blockchain() []Block {
	return m.chain
}

func (bc Blockchain) PendingTxs() []Transaction {
	return m.pendingTxs
}

// NewTransaction appends a transaction to
// the pending transactions
func (bc *Blockchain) NewTransaction(tx Transaction) {
	bc.pendingTxs = append(bc.pendingTxs, tx)
}

// AppendBlock adds the block to the chain
// if it's valid
func (bc *Blockchain) AppendBlock(b Block) {
	if 
	hdrPrevBlock := bc.chain[len(bc.chain)-1].Header
	b.HashPrevBlock = hash(hdrPrevBlock)
	m.chain = append(m.chain, b)
}

// Halve payout every 5 blocks
func (bc Blockchain) payout() int {
	return 50 / math.Ceil(len(bc.Chain)/5)
}

// Get balance calculates the balance from an address
func (bc Blockchain) GetBalance(address string) int {
	balance := 0
	for _, b := range m.Blockchain.Chain {
		for _, t := range b.Transactions {
			if t.From == address {
				balance -= t.Amount
			}
			if t.To == address {
				balance += t.Amount
			}
		}
	}
	return balance
}

// ShouldMine decides when next block should be mined
func (bc Blockchain) ShouldMine() bool {
	return len(bc.pendingTxs) >= 3
}

// ShouldReplace decides if a blockchain is more valid	
func (bc Blockchain) ShouldReplace(rbc Blockcain) bool {
	return ValidateBlockchain(rbc) && len(rbc.Chain) <= len(bc.Chain)
}

// Validate validates a blockchain
func ValidateBlockchain(bc Blockchain) error {
	for _, b := range bc.Chain[1:] {
		if !b.Validate() {
			return false
		}
	}
	return true
}

