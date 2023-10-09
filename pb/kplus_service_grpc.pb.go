// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// KPlusServiceClient is the client API for KPlusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KPlusServiceClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	// kPLUS Services
	SearchCustomerCID(ctx context.Context, in *KPLUSCustomerRequest, opts ...grpc.CallOption) (*KPLUSCustomerResponse, error)
	CustSavingsList(ctx context.Context, in *KPLUSCustomerRequest, opts ...grpc.CallOption) (*KPLUSCustSavingsListResponse, error)
	GetTransactionHistory(ctx context.Context, in *KPLUSGetTransactionHistoryRequest, opts ...grpc.CallOption) (*KPLUSGetTransactionHistoryResponse, error)
	GenerateColShtperCID(ctx context.Context, in *KPLUSCustomerRequest, opts ...grpc.CallOption) (*KPLUSGenerateColShtperCIDResponse, error)
	K2CCallBackRef(ctx context.Context, in *KPLUSCallBackRefRequest, opts ...grpc.CallOption) (*KPLUSResponse, error)
	GetReferences(ctx context.Context, in *KPLUSGetReferencesRequest, opts ...grpc.CallOption) (*KPLUSGetReferencesResponse, error)
	MultiplePayment(ctx context.Context, in *KPLUSMultiplePaymentRequest, opts ...grpc.CallOption) (*KPLUSResponse, error)
	SearchLoanList(ctx context.Context, in *KPLUSCustomerRequest, opts ...grpc.CallOption) (*KPLUSSearchLoanListResponse, error)
	LoanInfo(ctx context.Context, in *KPLUSAccRequest, opts ...grpc.CallOption) (*KPLUSLoanInfoResponse, error)
	GetSavingForSuperApp(ctx context.Context, in *KPLUSCustomerRequest, opts ...grpc.CallOption) (*KPLUSGetSavingResponse, error)
	FundTransferRequest(ctx context.Context, in *KPLUSFundTransferRequest, opts ...grpc.CallOption) (*KPLUSResponse, error)
}

type kPlusServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKPlusServiceClient(cc grpc.ClientConnInterface) KPlusServiceClient {
	return &kPlusServiceClient{cc}
}

func (c *kPlusServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/simplebank.KPlusService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kPlusServiceClient) SearchCustomerCID(ctx context.Context, in *KPLUSCustomerRequest, opts ...grpc.CallOption) (*KPLUSCustomerResponse, error) {
	out := new(KPLUSCustomerResponse)
	err := c.cc.Invoke(ctx, "/simplebank.KPlusService/SearchCustomerCID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kPlusServiceClient) CustSavingsList(ctx context.Context, in *KPLUSCustomerRequest, opts ...grpc.CallOption) (*KPLUSCustSavingsListResponse, error) {
	out := new(KPLUSCustSavingsListResponse)
	err := c.cc.Invoke(ctx, "/simplebank.KPlusService/CustSavingsList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kPlusServiceClient) GetTransactionHistory(ctx context.Context, in *KPLUSGetTransactionHistoryRequest, opts ...grpc.CallOption) (*KPLUSGetTransactionHistoryResponse, error) {
	out := new(KPLUSGetTransactionHistoryResponse)
	err := c.cc.Invoke(ctx, "/simplebank.KPlusService/GetTransactionHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kPlusServiceClient) GenerateColShtperCID(ctx context.Context, in *KPLUSCustomerRequest, opts ...grpc.CallOption) (*KPLUSGenerateColShtperCIDResponse, error) {
	out := new(KPLUSGenerateColShtperCIDResponse)
	err := c.cc.Invoke(ctx, "/simplebank.KPlusService/GenerateColShtperCID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kPlusServiceClient) K2CCallBackRef(ctx context.Context, in *KPLUSCallBackRefRequest, opts ...grpc.CallOption) (*KPLUSResponse, error) {
	out := new(KPLUSResponse)
	err := c.cc.Invoke(ctx, "/simplebank.KPlusService/K2CCallBackRef", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kPlusServiceClient) GetReferences(ctx context.Context, in *KPLUSGetReferencesRequest, opts ...grpc.CallOption) (*KPLUSGetReferencesResponse, error) {
	out := new(KPLUSGetReferencesResponse)
	err := c.cc.Invoke(ctx, "/simplebank.KPlusService/GetReferences", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kPlusServiceClient) MultiplePayment(ctx context.Context, in *KPLUSMultiplePaymentRequest, opts ...grpc.CallOption) (*KPLUSResponse, error) {
	out := new(KPLUSResponse)
	err := c.cc.Invoke(ctx, "/simplebank.KPlusService/MultiplePayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kPlusServiceClient) SearchLoanList(ctx context.Context, in *KPLUSCustomerRequest, opts ...grpc.CallOption) (*KPLUSSearchLoanListResponse, error) {
	out := new(KPLUSSearchLoanListResponse)
	err := c.cc.Invoke(ctx, "/simplebank.KPlusService/SearchLoanList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kPlusServiceClient) LoanInfo(ctx context.Context, in *KPLUSAccRequest, opts ...grpc.CallOption) (*KPLUSLoanInfoResponse, error) {
	out := new(KPLUSLoanInfoResponse)
	err := c.cc.Invoke(ctx, "/simplebank.KPlusService/LoanInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kPlusServiceClient) GetSavingForSuperApp(ctx context.Context, in *KPLUSCustomerRequest, opts ...grpc.CallOption) (*KPLUSGetSavingResponse, error) {
	out := new(KPLUSGetSavingResponse)
	err := c.cc.Invoke(ctx, "/simplebank.KPlusService/GetSavingForSuperApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kPlusServiceClient) FundTransferRequest(ctx context.Context, in *KPLUSFundTransferRequest, opts ...grpc.CallOption) (*KPLUSResponse, error) {
	out := new(KPLUSResponse)
	err := c.cc.Invoke(ctx, "/simplebank.KPlusService/FundTransferRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KPlusServiceServer is the server API for KPlusService service.
// All implementations must embed UnimplementedKPlusServiceServer
// for forward compatibility
type KPlusServiceServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	// kPLUS Services
	SearchCustomerCID(context.Context, *KPLUSCustomerRequest) (*KPLUSCustomerResponse, error)
	CustSavingsList(context.Context, *KPLUSCustomerRequest) (*KPLUSCustSavingsListResponse, error)
	GetTransactionHistory(context.Context, *KPLUSGetTransactionHistoryRequest) (*KPLUSGetTransactionHistoryResponse, error)
	GenerateColShtperCID(context.Context, *KPLUSCustomerRequest) (*KPLUSGenerateColShtperCIDResponse, error)
	K2CCallBackRef(context.Context, *KPLUSCallBackRefRequest) (*KPLUSResponse, error)
	GetReferences(context.Context, *KPLUSGetReferencesRequest) (*KPLUSGetReferencesResponse, error)
	MultiplePayment(context.Context, *KPLUSMultiplePaymentRequest) (*KPLUSResponse, error)
	SearchLoanList(context.Context, *KPLUSCustomerRequest) (*KPLUSSearchLoanListResponse, error)
	LoanInfo(context.Context, *KPLUSAccRequest) (*KPLUSLoanInfoResponse, error)
	GetSavingForSuperApp(context.Context, *KPLUSCustomerRequest) (*KPLUSGetSavingResponse, error)
	FundTransferRequest(context.Context, *KPLUSFundTransferRequest) (*KPLUSResponse, error)
	mustEmbedUnimplementedKPlusServiceServer()
}

// UnimplementedKPlusServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKPlusServiceServer struct {
}

func (UnimplementedKPlusServiceServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedKPlusServiceServer) SearchCustomerCID(context.Context, *KPLUSCustomerRequest) (*KPLUSCustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchCustomerCID not implemented")
}
func (UnimplementedKPlusServiceServer) CustSavingsList(context.Context, *KPLUSCustomerRequest) (*KPLUSCustSavingsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CustSavingsList not implemented")
}
func (UnimplementedKPlusServiceServer) GetTransactionHistory(context.Context, *KPLUSGetTransactionHistoryRequest) (*KPLUSGetTransactionHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionHistory not implemented")
}
func (UnimplementedKPlusServiceServer) GenerateColShtperCID(context.Context, *KPLUSCustomerRequest) (*KPLUSGenerateColShtperCIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateColShtperCID not implemented")
}
func (UnimplementedKPlusServiceServer) K2CCallBackRef(context.Context, *KPLUSCallBackRefRequest) (*KPLUSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method K2CCallBackRef not implemented")
}
func (UnimplementedKPlusServiceServer) GetReferences(context.Context, *KPLUSGetReferencesRequest) (*KPLUSGetReferencesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReferences not implemented")
}
func (UnimplementedKPlusServiceServer) MultiplePayment(context.Context, *KPLUSMultiplePaymentRequest) (*KPLUSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MultiplePayment not implemented")
}
func (UnimplementedKPlusServiceServer) SearchLoanList(context.Context, *KPLUSCustomerRequest) (*KPLUSSearchLoanListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchLoanList not implemented")
}
func (UnimplementedKPlusServiceServer) LoanInfo(context.Context, *KPLUSAccRequest) (*KPLUSLoanInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoanInfo not implemented")
}
func (UnimplementedKPlusServiceServer) GetSavingForSuperApp(context.Context, *KPLUSCustomerRequest) (*KPLUSGetSavingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSavingForSuperApp not implemented")
}
func (UnimplementedKPlusServiceServer) FundTransferRequest(context.Context, *KPLUSFundTransferRequest) (*KPLUSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FundTransferRequest not implemented")
}
func (UnimplementedKPlusServiceServer) mustEmbedUnimplementedKPlusServiceServer() {}

// UnsafeKPlusServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KPlusServiceServer will
// result in compilation errors.
type UnsafeKPlusServiceServer interface {
	mustEmbedUnimplementedKPlusServiceServer()
}

func RegisterKPlusServiceServer(s grpc.ServiceRegistrar, srv KPlusServiceServer) {
	s.RegisterService(&KPlusService_ServiceDesc, srv)
}

func _KPlusService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KPlusServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simplebank.KPlusService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KPlusServiceServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KPlusService_SearchCustomerCID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KPLUSCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KPlusServiceServer).SearchCustomerCID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simplebank.KPlusService/SearchCustomerCID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KPlusServiceServer).SearchCustomerCID(ctx, req.(*KPLUSCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KPlusService_CustSavingsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KPLUSCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KPlusServiceServer).CustSavingsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simplebank.KPlusService/CustSavingsList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KPlusServiceServer).CustSavingsList(ctx, req.(*KPLUSCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KPlusService_GetTransactionHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KPLUSGetTransactionHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KPlusServiceServer).GetTransactionHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simplebank.KPlusService/GetTransactionHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KPlusServiceServer).GetTransactionHistory(ctx, req.(*KPLUSGetTransactionHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KPlusService_GenerateColShtperCID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KPLUSCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KPlusServiceServer).GenerateColShtperCID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simplebank.KPlusService/GenerateColShtperCID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KPlusServiceServer).GenerateColShtperCID(ctx, req.(*KPLUSCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KPlusService_K2CCallBackRef_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KPLUSCallBackRefRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KPlusServiceServer).K2CCallBackRef(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simplebank.KPlusService/K2CCallBackRef",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KPlusServiceServer).K2CCallBackRef(ctx, req.(*KPLUSCallBackRefRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KPlusService_GetReferences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KPLUSGetReferencesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KPlusServiceServer).GetReferences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simplebank.KPlusService/GetReferences",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KPlusServiceServer).GetReferences(ctx, req.(*KPLUSGetReferencesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KPlusService_MultiplePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KPLUSMultiplePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KPlusServiceServer).MultiplePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simplebank.KPlusService/MultiplePayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KPlusServiceServer).MultiplePayment(ctx, req.(*KPLUSMultiplePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KPlusService_SearchLoanList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KPLUSCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KPlusServiceServer).SearchLoanList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simplebank.KPlusService/SearchLoanList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KPlusServiceServer).SearchLoanList(ctx, req.(*KPLUSCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KPlusService_LoanInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KPLUSAccRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KPlusServiceServer).LoanInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simplebank.KPlusService/LoanInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KPlusServiceServer).LoanInfo(ctx, req.(*KPLUSAccRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KPlusService_GetSavingForSuperApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KPLUSCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KPlusServiceServer).GetSavingForSuperApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simplebank.KPlusService/GetSavingForSuperApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KPlusServiceServer).GetSavingForSuperApp(ctx, req.(*KPLUSCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KPlusService_FundTransferRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KPLUSFundTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KPlusServiceServer).FundTransferRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simplebank.KPlusService/FundTransferRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KPlusServiceServer).FundTransferRequest(ctx, req.(*KPLUSFundTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KPlusService_ServiceDesc is the grpc.ServiceDesc for KPlusService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KPlusService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "simplebank.KPlusService",
	HandlerType: (*KPlusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _KPlusService_SayHello_Handler,
		},
		{
			MethodName: "SearchCustomerCID",
			Handler:    _KPlusService_SearchCustomerCID_Handler,
		},
		{
			MethodName: "CustSavingsList",
			Handler:    _KPlusService_CustSavingsList_Handler,
		},
		{
			MethodName: "GetTransactionHistory",
			Handler:    _KPlusService_GetTransactionHistory_Handler,
		},
		{
			MethodName: "GenerateColShtperCID",
			Handler:    _KPlusService_GenerateColShtperCID_Handler,
		},
		{
			MethodName: "K2CCallBackRef",
			Handler:    _KPlusService_K2CCallBackRef_Handler,
		},
		{
			MethodName: "GetReferences",
			Handler:    _KPlusService_GetReferences_Handler,
		},
		{
			MethodName: "MultiplePayment",
			Handler:    _KPlusService_MultiplePayment_Handler,
		},
		{
			MethodName: "SearchLoanList",
			Handler:    _KPlusService_SearchLoanList_Handler,
		},
		{
			MethodName: "LoanInfo",
			Handler:    _KPlusService_LoanInfo_Handler,
		},
		{
			MethodName: "GetSavingForSuperApp",
			Handler:    _KPlusService_GetSavingForSuperApp_Handler,
		},
		{
			MethodName: "FundTransferRequest",
			Handler:    _KPlusService_FundTransferRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kplus_service.proto",
}
