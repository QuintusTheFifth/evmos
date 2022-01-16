// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evmos/fees/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryWithdrawAddressesRequest is the request type for the Query/WithdrawAddresses RPC method.
type QueryWithdrawAddressesRequest struct {
	// pagination defines an optional pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryWithdrawAddressesRequest) Reset()         { *m = QueryWithdrawAddressesRequest{} }
func (m *QueryWithdrawAddressesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryWithdrawAddressesRequest) ProtoMessage()    {}
func (*QueryWithdrawAddressesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6533658f63002c05, []int{0}
}
func (m *QueryWithdrawAddressesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryWithdrawAddressesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryWithdrawAddressesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryWithdrawAddressesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryWithdrawAddressesRequest.Merge(m, src)
}
func (m *QueryWithdrawAddressesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryWithdrawAddressesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryWithdrawAddressesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryWithdrawAddressesRequest proto.InternalMessageInfo

func (m *QueryWithdrawAddressesRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryWithdrawAddressesResponse is the response type for the Query/WithdrawAddresses RPC
// method.
type QueryWithdrawAddressesResponse struct {
	// hex addresses of contract used for withdrawing tx fees
	WithdrawAddresses []ContractWithdrawAddress `protobuf:"bytes,1,rep,name=withdraw_addresses,json=withdrawAddresses,proto3" json:"withdraw_addresses"`
	// pagination defines the pagination in the response.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryWithdrawAddressesResponse) Reset()         { *m = QueryWithdrawAddressesResponse{} }
func (m *QueryWithdrawAddressesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryWithdrawAddressesResponse) ProtoMessage()    {}
func (*QueryWithdrawAddressesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6533658f63002c05, []int{1}
}
func (m *QueryWithdrawAddressesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryWithdrawAddressesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryWithdrawAddressesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryWithdrawAddressesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryWithdrawAddressesResponse.Merge(m, src)
}
func (m *QueryWithdrawAddressesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryWithdrawAddressesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryWithdrawAddressesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryWithdrawAddressesResponse proto.InternalMessageInfo

func (m *QueryWithdrawAddressesResponse) GetWithdrawAddresses() []ContractWithdrawAddress {
	if m != nil {
		return m.WithdrawAddresses
	}
	return nil
}

func (m *QueryWithdrawAddressesResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryWithdrawAddressRequest is the request type for the Query/WithdrawAddress RPC method.
type QueryWithdrawAddressRequest struct {
	// contract hex address registered on the fee distribution module
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
}

func (m *QueryWithdrawAddressRequest) Reset()         { *m = QueryWithdrawAddressRequest{} }
func (m *QueryWithdrawAddressRequest) String() string { return proto.CompactTextString(m) }
func (*QueryWithdrawAddressRequest) ProtoMessage()    {}
func (*QueryWithdrawAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6533658f63002c05, []int{2}
}
func (m *QueryWithdrawAddressRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryWithdrawAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryWithdrawAddressRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryWithdrawAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryWithdrawAddressRequest.Merge(m, src)
}
func (m *QueryWithdrawAddressRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryWithdrawAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryWithdrawAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryWithdrawAddressRequest proto.InternalMessageInfo

func (m *QueryWithdrawAddressRequest) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

// QueryWithdrawAddressResponse is the response type for the Query/WithdrawAddress RPC
// method.
type QueryWithdrawAddressResponse struct {
	// hex address from the registered owner/withdraw address
	WithdrawAddress string `protobuf:"bytes,1,opt,name=withdraw_address,json=withdrawAddress,proto3" json:"withdraw_address,omitempty"`
}

func (m *QueryWithdrawAddressResponse) Reset()         { *m = QueryWithdrawAddressResponse{} }
func (m *QueryWithdrawAddressResponse) String() string { return proto.CompactTextString(m) }
func (*QueryWithdrawAddressResponse) ProtoMessage()    {}
func (*QueryWithdrawAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6533658f63002c05, []int{3}
}
func (m *QueryWithdrawAddressResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryWithdrawAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryWithdrawAddressResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryWithdrawAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryWithdrawAddressResponse.Merge(m, src)
}
func (m *QueryWithdrawAddressResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryWithdrawAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryWithdrawAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryWithdrawAddressResponse proto.InternalMessageInfo

func (m *QueryWithdrawAddressResponse) GetWithdrawAddress() string {
	if m != nil {
		return m.WithdrawAddress
	}
	return ""
}

// QueryParamsRequest is the request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6533658f63002c05, []int{4}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is the response type for the Query/Params RPC
// method.
type QueryParamsResponse struct {
	// distribution module parameters
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6533658f63002c05, []int{5}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*QueryWithdrawAddressesRequest)(nil), "evmos.fees.v1.QueryWithdrawAddressesRequest")
	proto.RegisterType((*QueryWithdrawAddressesResponse)(nil), "evmos.fees.v1.QueryWithdrawAddressesResponse")
	proto.RegisterType((*QueryWithdrawAddressRequest)(nil), "evmos.fees.v1.QueryWithdrawAddressRequest")
	proto.RegisterType((*QueryWithdrawAddressResponse)(nil), "evmos.fees.v1.QueryWithdrawAddressResponse")
	proto.RegisterType((*QueryParamsRequest)(nil), "evmos.fees.v1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "evmos.fees.v1.QueryParamsResponse")
}

func init() { proto.RegisterFile("evmos/fees/v1/query.proto", fileDescriptor_6533658f63002c05) }

var fileDescriptor_6533658f63002c05 = []byte{
	// 537 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x8d, 0x5b, 0x88, 0xc4, 0x45, 0xa8, 0xed, 0x50, 0xa4, 0x90, 0xb6, 0x6e, 0xb1, 0x54, 0x20,
	0x2d, 0xcc, 0xc8, 0xed, 0x16, 0x81, 0x28, 0x88, 0xd7, 0xaa, 0x64, 0x83, 0x04, 0x0b, 0x34, 0x4e,
	0x86, 0xc9, 0x48, 0xc4, 0xe3, 0x7a, 0x26, 0x0e, 0x15, 0x62, 0xc3, 0x17, 0x20, 0xf1, 0x09, 0xfc,
	0x00, 0x5f, 0x81, 0xba, 0xac, 0xc4, 0xa6, 0x2b, 0x84, 0x12, 0x3e, 0x04, 0x79, 0x66, 0x2c, 0xb0,
	0x9b, 0xa6, 0xd9, 0x59, 0xf7, 0x71, 0xee, 0x39, 0xe7, 0x5e, 0x0f, 0x5c, 0x67, 0x59, 0x5f, 0x2a,
	0xf2, 0x8e, 0x31, 0x45, 0xb2, 0x90, 0x1c, 0x0c, 0x58, 0x7a, 0x88, 0x93, 0x54, 0x6a, 0x89, 0xae,
	0x98, 0x14, 0xce, 0x53, 0x38, 0x0b, 0x9b, 0x5b, 0x1d, 0xa9, 0xf2, 0xd2, 0x88, 0x2a, 0x66, 0xeb,
	0x48, 0x16, 0x46, 0x4c, 0xd3, 0x90, 0x24, 0x94, 0x8b, 0x98, 0x6a, 0x21, 0x63, 0xdb, 0xda, 0x6c,
	0x94, 0x51, 0x0d, 0x84, 0xcd, 0xac, 0x94, 0x33, 0x9c, 0xc5, 0x4c, 0x89, 0x22, 0xb9, 0xcc, 0x25,
	0x97, 0xe6, 0x93, 0xe4, 0x5f, 0x2e, 0xba, 0xca, 0xa5, 0xe4, 0xef, 0x19, 0xa1, 0x89, 0x20, 0x34,
	0x8e, 0xa5, 0x36, 0x93, 0x5c, 0x4f, 0xc0, 0x61, 0xed, 0x65, 0x4e, 0xe6, 0x95, 0xd0, 0xbd, 0x6e,
	0x4a, 0x87, 0x0f, 0xbb, 0xdd, 0x94, 0x29, 0xc5, 0x54, 0x9b, 0x1d, 0x0c, 0x98, 0xd2, 0xe8, 0x09,
	0xc0, 0x3f, 0x7e, 0x0d, 0x6f, 0xc3, 0xbb, 0x7d, 0x79, 0xe7, 0x26, 0xb6, 0x62, 0x70, 0x2e, 0x06,
	0x5b, 0xd1, 0x4e, 0x0c, 0xde, 0xa7, 0x9c, 0xb9, 0xde, 0xf6, 0x7f, 0x9d, 0xc1, 0x0f, 0x0f, 0xfc,
	0xb3, 0x26, 0xa9, 0x44, 0xc6, 0x8a, 0xa1, 0x37, 0x80, 0x86, 0x2e, 0xf9, 0x96, 0x16, 0xd9, 0x86,
	0xb7, 0x31, 0x6f, 0x46, 0x96, 0xec, 0xc4, 0x8f, 0x64, 0xac, 0x53, 0xda, 0xd1, 0x15, 0xb4, 0xbd,
	0x0b, 0x47, 0xbf, 0xd6, 0x6b, 0xed, 0xa5, 0x61, 0x75, 0x08, 0x7a, 0x5a, 0xd2, 0x31, 0x67, 0x74,
	0xdc, 0x3a, 0x57, 0x87, 0x65, 0x56, 0x12, 0xf2, 0x0c, 0x56, 0x26, 0xe9, 0x28, 0xfc, 0x6a, 0xc1,
	0x62, 0xc7, 0x71, 0x2b, 0x44, 0x18, 0xd7, 0x2e, 0xb5, 0x17, 0x8a, 0xb8, 0xeb, 0x08, 0x9e, 0xc3,
	0xea, 0x64, 0x24, 0xe7, 0x47, 0x0b, 0x16, 0xab, 0x7e, 0x14, 0x50, 0x15, 0x7d, 0xc1, 0x32, 0x20,
	0x03, 0xb5, 0x4f, 0x53, 0xda, 0x2f, 0xb8, 0x04, 0x2f, 0xe0, 0x6a, 0x29, 0xea, 0x70, 0x77, 0xa1,
	0x9e, 0x98, 0x88, 0x5b, 0xe7, 0xb5, 0x8a, 0xb7, 0xb6, 0xdc, 0x59, 0xe9, 0x4a, 0x77, 0x4e, 0xe6,
	0xe1, 0xa2, 0x01, 0x43, 0xdf, 0x3c, 0x58, 0x3a, 0xb5, 0x44, 0x74, 0xa7, 0x02, 0x32, 0xf5, 0xaa,
	0x9a, 0x77, 0x67, 0xac, 0xb6, 0x8c, 0x83, 0xf0, 0xf3, 0xcf, 0x3f, 0x5f, 0xe7, 0xb6, 0x51, 0x8b,
	0xd8, 0xfb, 0xef, 0x0a, 0xa5, 0x53, 0x11, 0x0d, 0xf2, 0x85, 0xe4, 0xff, 0xc1, 0xe9, 0xb3, 0x41,
	0xdf, 0x3d, 0x58, 0xa8, 0x00, 0xa2, 0xad, 0x19, 0xa6, 0x16, 0x0c, 0xb7, 0x67, 0xaa, 0x75, 0xfc,
	0x1e, 0x1b, 0x7e, 0xf7, 0xd1, 0xbd, 0x99, 0xf9, 0x91, 0x8f, 0xd5, 0x2b, 0xf9, 0x84, 0x32, 0xa8,
	0x5b, 0xeb, 0xd1, 0x8d, 0x49, 0xc3, 0x4b, 0xbb, 0x6d, 0x06, 0xd3, 0x4a, 0x1c, 0xad, 0x4d, 0x43,
	0x6b, 0x1d, 0xad, 0x9d, 0x41, 0xcb, 0xae, 0x76, 0xef, 0xc1, 0xd1, 0xc8, 0xf7, 0x8e, 0x47, 0xbe,
	0xf7, 0x7b, 0xe4, 0x7b, 0x5f, 0xc6, 0x7e, 0xed, 0x78, 0xec, 0xd7, 0x4e, 0xc6, 0x7e, 0xed, 0xf5,
	0x26, 0x17, 0xba, 0x37, 0x88, 0x70, 0x47, 0xf6, 0x89, 0xee, 0xd1, 0x54, 0x09, 0xe5, 0xa0, 0x3e,
	0xd8, 0x37, 0x48, 0x1f, 0x26, 0x4c, 0x45, 0x75, 0xf3, 0x96, 0xec, 0xfe, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0xba, 0x62, 0x78, 0xbc, 0x0e, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// WithdrawAddresses retrieves all withdrawal addresses of contracts for fee distribution
	WithdrawAddresses(ctx context.Context, in *QueryWithdrawAddressesRequest, opts ...grpc.CallOption) (*QueryWithdrawAddressesResponse, error)
	// WithdrawAddress retrieves a registered withdrawal address of contract for fee distribution
	WithdrawAddress(ctx context.Context, in *QueryWithdrawAddressRequest, opts ...grpc.CallOption) (*QueryWithdrawAddressResponse, error)
	// Params retrieves the distribution module params
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) WithdrawAddresses(ctx context.Context, in *QueryWithdrawAddressesRequest, opts ...grpc.CallOption) (*QueryWithdrawAddressesResponse, error) {
	out := new(QueryWithdrawAddressesResponse)
	err := c.cc.Invoke(ctx, "/evmos.fees.v1.Query/WithdrawAddresses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) WithdrawAddress(ctx context.Context, in *QueryWithdrawAddressRequest, opts ...grpc.CallOption) (*QueryWithdrawAddressResponse, error) {
	out := new(QueryWithdrawAddressResponse)
	err := c.cc.Invoke(ctx, "/evmos.fees.v1.Query/WithdrawAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/evmos.fees.v1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// WithdrawAddresses retrieves all withdrawal addresses of contracts for fee distribution
	WithdrawAddresses(context.Context, *QueryWithdrawAddressesRequest) (*QueryWithdrawAddressesResponse, error)
	// WithdrawAddress retrieves a registered withdrawal address of contract for fee distribution
	WithdrawAddress(context.Context, *QueryWithdrawAddressRequest) (*QueryWithdrawAddressResponse, error)
	// Params retrieves the distribution module params
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) WithdrawAddresses(ctx context.Context, req *QueryWithdrawAddressesRequest) (*QueryWithdrawAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawAddresses not implemented")
}
func (*UnimplementedQueryServer) WithdrawAddress(ctx context.Context, req *QueryWithdrawAddressRequest) (*QueryWithdrawAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawAddress not implemented")
}
func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_WithdrawAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryWithdrawAddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).WithdrawAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evmos.fees.v1.Query/WithdrawAddresses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).WithdrawAddresses(ctx, req.(*QueryWithdrawAddressesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_WithdrawAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryWithdrawAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).WithdrawAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evmos.fees.v1.Query/WithdrawAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).WithdrawAddress(ctx, req.(*QueryWithdrawAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evmos.fees.v1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "evmos.fees.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WithdrawAddresses",
			Handler:    _Query_WithdrawAddresses_Handler,
		},
		{
			MethodName: "WithdrawAddress",
			Handler:    _Query_WithdrawAddress_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "evmos/fees/v1/query.proto",
}

func (m *QueryWithdrawAddressesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryWithdrawAddressesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryWithdrawAddressesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryWithdrawAddressesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryWithdrawAddressesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryWithdrawAddressesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.WithdrawAddresses) > 0 {
		for iNdEx := len(m.WithdrawAddresses) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.WithdrawAddresses[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryWithdrawAddressRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryWithdrawAddressRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryWithdrawAddressRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryWithdrawAddressResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryWithdrawAddressResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryWithdrawAddressResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.WithdrawAddress) > 0 {
		i -= len(m.WithdrawAddress)
		copy(dAtA[i:], m.WithdrawAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.WithdrawAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryWithdrawAddressesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryWithdrawAddressesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.WithdrawAddresses) > 0 {
		for _, e := range m.WithdrawAddresses {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryWithdrawAddressRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryWithdrawAddressResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.WithdrawAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryWithdrawAddressesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryWithdrawAddressesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryWithdrawAddressesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryWithdrawAddressesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryWithdrawAddressesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryWithdrawAddressesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawAddresses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WithdrawAddresses = append(m.WithdrawAddresses, ContractWithdrawAddress{})
			if err := m.WithdrawAddresses[len(m.WithdrawAddresses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryWithdrawAddressRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryWithdrawAddressRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryWithdrawAddressRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryWithdrawAddressResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryWithdrawAddressResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryWithdrawAddressResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WithdrawAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)