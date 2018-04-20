package blockchain

import "crypto/sha256"

// Miner interacts with its blockchain
type Miner interface {
	Mine() Block
	Blockchain() Blockchain
	NewTransaction(Transaction, chan blockchain)
	Receiveblockchain(Blockchain)
}

// BCMiner is a Miner that behaves in
// the network's best interest
type BCMiner struct {
	Address string
	blockchain *Blockchain
}

// Blockchain returns a copy of the miner's blockchain
func (m BCMiner) Blockchain() Blockchain {
	return *m.blockchain
}

// NewTransaction handles an incoming transaction, and receives a broadcastBc
// that is used for validating incoming new blockchains, and sending a blockchain
// to the miner pool
func (m BCMiner) NewTransaction(tx Transaction, broadcastBC chan blockchain) {
	err := m.blockchain.AppendTx(tx)
	log.Printf("%s - Transaction: from %s, to %s, for %v. Valid: %v",
		m.Address, t.From, t.To, t.Amount, err == nil)
	if err != nil {
		log.Println(err)
		return
	}
	
	if !m.blockchain.shouldMine() {
		return	
	}

	minedBlock := make(chan Block)
	go func() {
		minedBlock <- m.Mine()
	}()
	for {
		select {
		case mb := <-minedBlock:
				// New blockchain will be broadcasted to all miners,
				// including broadcaster
				broadcastBC<- append(m.blockchain.BlockChain(), mb)
				return
			case bcbc := <-broadcastBC:
				if err := m.SetBlockchain(bcbc); err != nil {
					return
				}
		}
	}
}

// Mine computes the next block in the miner's blockchain
func (m *BCMiner) Mine() Block {
	txs := m.blockchain.PendingTxs()
	mr := MerkleRoot(txs)
	nonce := m.ProofOfWork(mr)
	nextBlock := Block{
		Header: {
			Timestamp: time.Now(),
			Nonce: nonce
			MerkleRoot: mr,
		},
		Transactions: txs,
	}
	bc := m.blockchain.Blockchain()
	last
}

// ProofOfWork finds the block's nonce
func (m *BCMiner) ProofOfWork(b Block) []byte {
	nonce := 0
	mr := b.MerkleRoot()

	attempt := strconv.ParseInt(sha256.New().Sum([]byte(mr + nonce)), 16, 64)
	for attempt < b.Target {
		nonce++
		attempt = strconv.ParseInt(sha256.New().Sum([]byte(mr + nonce)), 16, 64)
	}
}

// GetBalance calculates the balance of an address
func (m *BCMiner) GetBalance(address) int {
	return m.blockchain.GetBalance(address)
}

// SetBlockchain changes the miner's blockchain
// if it's deemed more correct
func (m *BCMiner) SetBlockchain(bc Blockchain) error {
	m.blockchain.Compare(bc)
}
