package model

import (
  "crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
  "sync"
  "net/http"

  "garicoin/pkg/constant"
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

func (bc *Blockchain) CreateTransaction(sender string, recipient string, value float64, senderPublicKey *ecdsa.PublicKey, s *Signature) bool {
  isTransacted := bc.AddTransaction(sender, recipient, value, senderPublicKey, s)

  // sync other node

  return isTransacted
}

func (bc *Blockchain) AddTransaction(sender string, recipient string, value float64, senderPublicKey *ecdsa.PublicKey, s *Signature) bool {
	t := NewTransaction(sender, recipient, value)
	if sender == constant.MINING_SENDER {
		bc.transactionPool = append(bc.transactionPool, t)
		return true
	}
	if bc.VerifyTransactionSignature(senderPublicKey, s, t) {
		if bc.CalculateTotalAmount(sender) < value {
		  log.Println("ERROR: Not enough balance in a wallet")
		  return false
		}
		bc.transactionPool = append(bc.transactionPool, t)
		return true
	} else {
		log.Println("ERROR: Verify Transaction")
	}
	return false
}

func (bc *Blockchain) VerifyTransactionSignature(senderPublicKey *ecdsa.PublicKey, s *Signature, t *Transaction) bool {
	m, _ := json.Marshal(t)
	h := sha256.Sum256([]byte(m))
  fmt.Println(t)
	return ecdsa.Verify(senderPublicKey, h[:], s.R, s.S)
}

func (bc *Blockchain) CopyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, 0)
	for _, t := range bc.transactionPool {
		transactions = append(transactions, NewTransaction(t.senderBlockchainAddress, t.recipientBlockchainAddress, t.value))
	}
	return transactions
}

func (bc *Blockchain) ValidProof(nonce int, previousHash [32]byte, transactions []*Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := Block{0, nonce, previousHash, transactions}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashStr[:difficulty] == zeros
}

func (bc *Blockchain) ProofOfWork() int {
	transactions := bc.CopyTransactionPool()
	previousHash := bc.LastBlock().Hash()
	nonce := 0
	for !bc.ValidProof(nonce, previousHash, transactions, constant.MINING_DIFFICULTY) {
		nonce += 1
	}
	return nonce
}

func (bc *Blockchain) Mining() bool {
  bc.mux.Lock()
  defer bc.mux.Unlock()

  // if len(bc.transactionPool) == 0 {
  //   return false
  // }

	bc.AddTransaction(constant.MINING_SENDER, bc.blockchainAddress, constant.MINING_REWARD, nil, nil)
	nonce := bc.ProofOfWork()
	previousHash := bc.LastBlock().Hash()
	bc.CreateBlock(nonce, previousHash)
	log.Println("action=mining, status=success")

  for _, n := range bc.neighbors {
    endpoint := fmt.Sprintf("http://%s/consensus", n)
    client := &http.Client{}
    req, _ := http.NewRequest("PUT", endpoint, nil)
    resp, _ := client.Do(req)
    log.Printf("%v", resp)
  }
	return true
}

func (bc *Blockchain) StartMining() {
  bc.Mining()
  _ = time.AfterFunc(time.Second * constant.MINING_TIMER_SEC, bc.StartMining)
}

func (bc *Blockchain) CalculateTotalAmount(blockchainAddress string) float64 {
	var totalAmount float64 = 0.0
	for _, b := range bc.chain {
		for _, t := range b.transactions {
			value := t.value
			if blockchainAddress == t.recipientBlockchainAddress {
				totalAmount += value
			}
			if blockchainAddress == t.senderBlockchainAddress {
				totalAmount -= value
			}
		}
	}
	return totalAmount
}

func (bc *Blockchain) ValidChain(chain []*Block) bool {
  preBlock := chain[0]
  currentIndex := 1
  for currentIndex < len(chain) {
    b := chain[currentIndex]
    if b.previousHash != preBlock.Hash() {
      return false
    }
    if !bc.ValidProof(b.Nonce(), b.PreviousHash(), b.Transactions(), constant.MINING_DIFFICULTY) {
      return false
    }

    preBlock = b
    currentIndex += 1
  }
  return true
}
