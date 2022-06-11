package model

import (
	"encoding/json"
	"fmt"
	"strings"
  "sync"
)

type Blockchain struct {
	transactionPool   []*Transaction
	chain             []*Block
	blockchainAddress string
  port              uint16
  mux               sync.Mutex

  neighbors         []string
  muxNeighbors      sync.Mutex
}

func NewBlockchain(blockchainAddress string) *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.blockchainAddress = blockchainAddress
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) Chain() []*Block {
  return bc.chain
}

func (bc *Blockchain) TransactionPool() []*Transaction {
  return bc.transactionPool
}

func (bc *Blockchain) ClearTransactionPool() {
  bc.transactionPool = bc.transactionPool[:0]
}

func (bc *Blockchain) MarshalJSON() ([]byte, error) {
  return json.Marshal(struct {
		Blocks  []*Block `json:"chain"`
	}{
		Blocks: bc.chain,
	})
}

func (bc *Blockchain) UnmarshalJSON(data []byte) error {
  v := &struct {
    Blocks *[]*Block `json:"chain"`
  } {
    Blocks: &bc.chain,
  }
  if err := json.Unmarshal(data, &v); err != nil {
    return err
  }
  return nil
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
  // sync other node
	return b
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}
