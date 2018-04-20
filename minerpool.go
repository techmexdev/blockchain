package minerpool

import (
	"math"
)

// MinerPool enables communication between miners
type MinerPool struct {
	miners      []miner.Miner
	broadcastBc chan Blockchain
}

// New creates a MinerPool with the same blockchain
// across miners, and sets up a communication channel for miners
func New(miners []Miner) *MinerPool {
	bcbc := make(chan Blockchain)
	mp := &MinerPool{miners: miners, broadcastBC: bcbc}
	go mp.listen()
	return mp
}

// listen broadcasts incoming blockchains to all miners
func (mp *MinerPool) listen() {
	for {
		bcbc := <-mp.broadcastBc
		for _, m := range m.miners {
			mp.ReceiveBlockchain(bcbc)
		}
	}
}

// NewTransaction notifies all miners of an incoming transaction
func (mp *MinerPool) NewTransaction(tx Transaction) {
	for _, m := range mp.miners {
		go m.NewTransaction(tx, mp.broadcastBc)
	}
}

// GetBalance finds the balance for an address that
// the majority of miners agree on
func (mp *MinerPool) GetBalance(address) {
	var bals map[int]int
	for _, m := range m.miners {
		bal := m.GetBalance(address)
		bals[bal]++
	}
	majBal := 0
	for _, b := range bals {
		majBal = math.Max(majBal, b)
	}
	return majBal
}

// Blockchains retrieves all of the miners blockchains
func (mp *MinerPool) Blockchains() []Blockchain {
	var bcs []Blockchain
	for _, m := range mp.miners {
		bcs = append(bcs, m.Blockchain())
	}
	return bcs
}

// Blockchain retrieves the blockchain that the majority
// of miners agree on
func (mp *MinerPool) Blockchain() Blockchain {
	bcs := mp.Blockchains()

	var bcCount map[string]int
	for _, bc := range bcs {
		bcCount[bc]++
	}

	majBC := 0
	for _, bc = range bcCount {
		majBC = math.Max(majBC, bcCount[bc])
	}

	return majBC
}

// AddMiner adds a miner to the pool
func (mp *MinerPool) AddMiner(m Miner) {
	mp.miners = append(mp.miners, m)
}

// Remove miner removes a miner from the pool
func (mp *MinerPool) RemoveMiner(address string) {
	var miners []Miner
	for _, m := range mp.miners {
		if !m.Address == address {
			miners = append(miners, m)
		}
	}
	mp.miners = miners
}
