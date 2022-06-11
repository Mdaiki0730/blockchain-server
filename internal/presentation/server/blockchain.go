package server

import (
	"context"
	"encoding/json"

	"garicoin/api/proto/blockchain/blockchainpb"
	"garicoin/internal/application/command"
	"garicoin/internal/application/usecase"
	"garicoin/pkg/constant"

	"google.golang.org/protobuf/types/known/emptypb"
)

type blockchainManagementServer struct{
	application usecase.BlockchainAppIF
}

func NewBlockchainManagementServer(app usecase.BlockchainAppIF) blockchainpb.BlockchainManagementServer {
	return &blockchainManagementServer{app}
}

func (bms *blockchainManagementServer) GetChain(ctx context.Context, req *emptypb.Empty) (*blockchainpb.ChainResponse, error) {
	result, err := bms.application.GetChain(ctx)
	if err != nil {
		return nil, err
	}

	res := &blockchainpb.ChainResponse{}
	b, _ := json.Marshal(result)
	json.Unmarshal(b, res)
	return res, nil
}

func (bms *blockchainManagementServer) CreateTransactions(ctx context.Context, req *blockchainpb.CreateTransactionRequest) (*blockchainpb.StatusResponse, error) {
	cmd := command.TransactionCreate{}
	b, _ := json.Marshal(req)
	json.Unmarshal(b, &cmd)

	err := bms.application.CreateTransactions(ctx, cmd)
	if err != nil {
		return nil, err
	}

	status := constant.SUCCESS_MESSAGE
	res := &blockchainpb.StatusResponse{Status: &status}
	return res, nil
}

func (bms *blockchainManagementServer) GetTransactions(ctx context.Context, req *emptypb.Empty) (*blockchainpb.TransactionResponse, error) {
	result, err := bms.application.GetTransactions(ctx)
	if err != nil {
		return nil, err
	}

	res := &blockchainpb.TransactionResponse{}
	b, _ := json.Marshal(result)
	json.Unmarshal(b, res)
	return res, nil
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
