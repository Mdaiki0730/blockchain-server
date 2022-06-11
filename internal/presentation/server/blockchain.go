package server

import (
	"context"

	"garicoin/api/proto/blockchain/blockchainpb"

	"google.golang.org/protobuf/types/known/emptypb"
)

type blockchainManagementServer struct{}

func NewBlockchainManagementServer() blockchainpb.BlockchainManagementServer {
	return &blockchainManagementServer{}
}

func (bms *blockchainManagementServer) GetChain(context.Context, *emptypb.Empty) (*blockchainpb.ChainResponse, error) {
	return nil, nil
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
