// Basically copy of blockchain_test.go with some extra steps
package core

import (
	"math/big"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/params"
)

// So we can deterministically seed different blockchains
var (
	canonicalSeed = 1
	forkSeed      = 2
)

// newCanonical creates a chain database, and injects a deterministic canonical
// chain. Depending on the full flag, if creates either a full block chain or a
// header only chain.
func newCanonical(engine consensus.Engine, n int, full bool) (ethdb.Database, *BlockChain, error) {
	var (
		db      = rawdb.NewMemoryDatabase()
		genesis = (&Genesis{BaseFee: big.NewInt(params.InitialBaseFee)}).MustCommit(db)
	)

	// Initialize a fresh chain with only a genesis block
	blockchain, _ := NewBlockChain(db, nil, params.AllEthashProtocolChanges, engine, vm.Config{}, nil, nil)
	// Create and inject the requested chain
	if n == 0 {
		return db, blockchain, nil
	}
	if full {
		// Full block-chain requested
		blocks := makeBlockChain(genesis, n, engine, db, canonicalSeed)
		_, err := blockchain.InsertChain(blocks)
		return db, blockchain, err
	}
	// Header-only chain requested
	headers := makeHeaderChain(genesis.Header(), n, engine, db, canonicalSeed)
	_, err := blockchain.InsertHeaderChain(headers, 1)
	return db, blockchain, err
}

// Custom exported functions

func ExpNewCanonical(engine consensus.Engine, n int, full bool) (ethdb.Database, *BlockChain, error) {
	return newCanonical(engine, n, full)
}
