package server

import (
	"context"
	"encoding/json"

	"garicoin/api/proto/blockchain/blockchainpb"
	"garicoin/internal/domain/model"
	"garicoin/pkg/config"

	"google.golang.org/protobuf/types/known/emptypb"
)

type blockchainManagementServer struct{}

var cache map[string]*model.Blockchain = make(map[string]*model.Blockchain)

func NewBlockchainManagementServer() blockchainpb.BlockchainManagementServer {
	return &blockchainManagementServer{}
}

func (bms *blockchainManagementServer) GetBlockchain() *model.Blockchain {
  bc, ok := cache["blockchain"]
  if !ok {
    bc = model.NewBlockchain(config.Global.MinersBlockchainAddress)
    cache["blockchain"] = bc
  }
  return bc
}

func (bms *blockchainManagementServer) GetChain(context.Context, *emptypb.Empty) (*blockchainpb.ChainResponse, error) {
	bc := bms.GetBlockchain()

	res := &blockchainpb.ChainResponse{}
	b, _ := json.Marshal(bc)
	json.Unmarshal(b, res)
	return res, nil
}

func (bms *blockchainManagementServer) CreateTransactions(context.Context, *blockchainpb.CreateTransactionRequest) (*blockchainpb.StatusResponse, error) {
	return nil, nil
}

func (bms *blockchainManagementServer) GetTransactions(context.Context, *emptypb.Empty) (*blockchainpb.TransactionResponse, error) {
	return nil, nil
}

func (bms *blockchainManagementServer) Mine(context.Context, *blockchainpb.BlockchainAddressRequest) (*blockchainpb.StatusResponse, error) {
	return nil, nil
}

func (bms *blockchainManagementServer) StartMine(context.Context, *blockchainpb.BlockchainAddressRequest) (*blockchainpb.StatusResponse, error) {
	return nil, nil
}

func (bms *blockchainManagementServer) Amount(context.Context, *blockchainpb.BlockchainAddressRequest) (*blockchainpb.AmountResponse, error) {
	return nil, nil
}
