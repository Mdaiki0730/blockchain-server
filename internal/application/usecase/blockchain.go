package usecase

import (
	"context"
	"encoding/json"
	"errors"

	"garicoin/internal/application/command"
	"garicoin/internal/application/result"
	"garicoin/internal/domain/model"
	"garicoin/pkg/config"
	"garicoin/pkg/converter"
)

type BlockchainAppIF interface {
	GetChain(ctx context.Context) (*result.Blockchain, error)
	CreateTransactions(ctx context.Context, cmd command.TransactionCreate) error
	GetTransactions(ctx context.Context) (*result.Transactions, error)
	Amount(ctx context.Context, blockchainAddress string) (float64, error)
}

type blockchainApp struct{}

var cache map[string]*model.Blockchain = make(map[string]*model.Blockchain)

func NewBlockchainApp() BlockchainAppIF {
	ba := &blockchainApp{}
	bc := ba.GetBlockchain()
	bc.StartMining()
	return ba
}

func (ba *blockchainApp) GetBlockchain() *model.Blockchain {
	bc, ok := cache["blockchain"]
	if !ok {
		bc = model.NewBlockchain(config.Global.MinersBlockchainAddress)
		cache["blockchain"] = bc
	}
	return bc
}

func (ba *blockchainApp) GetChain(ctx context.Context) (*result.Blockchain, error) {
	bc := ba.GetBlockchain()

	res := &result.Blockchain{}
	b, _ := json.Marshal(bc)
	json.Unmarshal(b, res)
	return res, nil
}

func (ba *blockchainApp) CreateTransactions(ctx context.Context, cmd command.TransactionCreate) error {
	bc := ba.GetBlockchain()
	publicKey, err := converter.PublicKeyFromString(*cmd.SenderPublicKey)
	if err != nil {
		return err
	}
	signature, err := model.NewSignature(*cmd.Signature)
	if err != nil {
		return err
	}
	isCreated := bc.CreateTransaction(*cmd.SenderBlockchainAddress, *cmd.RecipientBlockchainAddress, *cmd.Value, publicKey, signature)
	if !isCreated {
		return errors.New("ERROR: can't create transactions")
	}
	return nil
}

func (ba *blockchainApp) GetTransactions(ctx context.Context) (*result.Transactions, error) {
	bc := ba.GetBlockchain()
	transactions := bc.TransactionPool()

	res := &result.Transactions{}
	b, _ := json.Marshal(transactions)
	json.Unmarshal(b, &res.Transactions)
	res.Length = len(transactions)
	return res, nil
}

func (ba *blockchainApp) Amount(ctx context.Context, blockchainAddress string) (float64, error) {
	amount := ba.GetBlockchain().CalculateTotalAmount(blockchainAddress)
	return amount, nil
}
