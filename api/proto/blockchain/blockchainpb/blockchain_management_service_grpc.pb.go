// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: blockchain_management_service.proto

package blockchainpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BlockchainManagementClient is the client API for BlockchainManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlockchainManagementClient interface {
	GetChain(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ChainResponse, error)
	CreateTransactions(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	GetTransactions(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TransactionResponse, error)
	Amount(ctx context.Context, in *BlockchainAddressRequest, opts ...grpc.CallOption) (*AmountResponse, error)
}

type blockchainManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewBlockchainManagementClient(cc grpc.ClientConnInterface) BlockchainManagementClient {
	return &blockchainManagementClient{cc}
}

func (c *blockchainManagementClient) GetChain(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ChainResponse, error) {
	out := new(ChainResponse)
	err := c.cc.Invoke(ctx, "/proto.BlockchainManagement/GetChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainManagementClient) CreateTransactions(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/proto.BlockchainManagement/CreateTransactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainManagementClient) GetTransactions(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.BlockchainManagement/GetTransactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainManagementClient) Amount(ctx context.Context, in *BlockchainAddressRequest, opts ...grpc.CallOption) (*AmountResponse, error) {
	out := new(AmountResponse)
	err := c.cc.Invoke(ctx, "/proto.BlockchainManagement/Amount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockchainManagementServer is the server API for BlockchainManagement service.
// All implementations should embed UnimplementedBlockchainManagementServer
// for forward compatibility
type BlockchainManagementServer interface {
	GetChain(context.Context, *emptypb.Empty) (*ChainResponse, error)
	CreateTransactions(context.Context, *CreateTransactionRequest) (*StatusResponse, error)
	GetTransactions(context.Context, *emptypb.Empty) (*TransactionResponse, error)
	Amount(context.Context, *BlockchainAddressRequest) (*AmountResponse, error)
}

// UnimplementedBlockchainManagementServer should be embedded to have forward compatible implementations.
type UnimplementedBlockchainManagementServer struct {
}

func (UnimplementedBlockchainManagementServer) GetChain(context.Context, *emptypb.Empty) (*ChainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChain not implemented")
}
func (UnimplementedBlockchainManagementServer) CreateTransactions(context.Context, *CreateTransactionRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransactions not implemented")
}
func (UnimplementedBlockchainManagementServer) GetTransactions(context.Context, *emptypb.Empty) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactions not implemented")
}
func (UnimplementedBlockchainManagementServer) Amount(context.Context, *BlockchainAddressRequest) (*AmountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Amount not implemented")
}

// UnsafeBlockchainManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlockchainManagementServer will
// result in compilation errors.
type UnsafeBlockchainManagementServer interface {
	mustEmbedUnimplementedBlockchainManagementServer()
}

func RegisterBlockchainManagementServer(s grpc.ServiceRegistrar, srv BlockchainManagementServer) {
	s.RegisterService(&BlockchainManagement_ServiceDesc, srv)
}

func _BlockchainManagement_GetChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainManagementServer).GetChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BlockchainManagement/GetChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainManagementServer).GetChain(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockchainManagement_CreateTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainManagementServer).CreateTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BlockchainManagement/CreateTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainManagementServer).CreateTransactions(ctx, req.(*CreateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockchainManagement_GetTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainManagementServer).GetTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BlockchainManagement/GetTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainManagementServer).GetTransactions(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockchainManagement_Amount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockchainAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainManagementServer).Amount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BlockchainManagement/Amount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainManagementServer).Amount(ctx, req.(*BlockchainAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BlockchainManagement_ServiceDesc is the grpc.ServiceDesc for BlockchainManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlockchainManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.BlockchainManagement",
	HandlerType: (*BlockchainManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChain",
			Handler:    _BlockchainManagement_GetChain_Handler,
		},
		{
			MethodName: "CreateTransactions",
			Handler:    _BlockchainManagement_CreateTransactions_Handler,
		},
		{
			MethodName: "GetTransactions",
			Handler:    _BlockchainManagement_GetTransactions_Handler,
		},
		{
			MethodName: "Amount",
			Handler:    _BlockchainManagement_Amount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blockchain_management_service.proto",
}
