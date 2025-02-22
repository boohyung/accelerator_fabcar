// Code generated by protoc-gen-go. DO NOT EDIT.
// source: accelerator.proto

package protos // import "github.com/nexledger/accelerator/protos"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 요청 메시지 구조 정의
type TxRequest struct {
	ChannelId            string   `protobuf:"bytes,1,opt,name=channelId" json:"channelId,omitempty"`
	ChaincodeName        string   `protobuf:"bytes,2,opt,name=chaincodeName" json:"chaincodeName,omitempty"`
	Fcn                  string   `protobuf:"bytes,3,opt,name=fcn" json:"fcn,omitempty"`
	Args                 [][]byte `protobuf:"bytes,4,rep,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxRequest) Reset()         { *m = TxRequest{} }
func (m *TxRequest) String() string { return proto.CompactTextString(m) }
func (*TxRequest) ProtoMessage()    {}
func (*TxRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_accelerator_61f348634469893b, []int{0}
}
func (m *TxRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxRequest.Unmarshal(m, b)
}
func (m *TxRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxRequest.Marshal(b, m, deterministic)
}
func (dst *TxRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxRequest.Merge(dst, src)
}
func (m *TxRequest) XXX_Size() int {
	return xxx_messageInfo_TxRequest.Size(m)
}
func (m *TxRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TxRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TxRequest proto.InternalMessageInfo

func (m *TxRequest) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *TxRequest) GetChaincodeName() string {
	if m != nil {
		return m.ChaincodeName
	}
	return ""
}

func (m *TxRequest) GetFcn() string {
	if m != nil {
		return m.Fcn
	}
	return ""
}

func (m *TxRequest) GetArgs() [][]byte {
	if m != nil {
		return m.Args
	}
	return nil
}
// 응답 메시지 구조 정의
type TxResponse struct {
	Payload              []byte                 `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	TxId                 string                 `protobuf:"bytes,2,opt,name=txId" json:"txId,omitempty"`
	Validation           *TransactionValidation `protobuf:"bytes,3,opt,name=validation" json:"validation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *TxResponse) Reset()         { *m = TxResponse{} }
func (m *TxResponse) String() string { return proto.CompactTextString(m) }
func (*TxResponse) ProtoMessage()    {}
func (*TxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_accelerator_61f348634469893b, []int{1}
}
func (m *TxResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxResponse.Unmarshal(m, b)
}
func (m *TxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxResponse.Marshal(b, m, deterministic)
}
func (dst *TxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxResponse.Merge(dst, src)
}
func (m *TxResponse) XXX_Size() int {
	return xxx_messageInfo_TxResponse.Size(m)
}
func (m *TxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TxResponse proto.InternalMessageInfo

func (m *TxResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *TxResponse) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *TxResponse) GetValidation() *TransactionValidation {
	if m != nil {
		return m.Validation
	}
	return nil
}

type TransactionValidation struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionValidation) Reset()         { *m = TransactionValidation{} }
func (m *TransactionValidation) String() string { return proto.CompactTextString(m) }
func (*TransactionValidation) ProtoMessage()    {}
func (*TransactionValidation) Descriptor() ([]byte, []int) {
	return fileDescriptor_accelerator_61f348634469893b, []int{2}
}
func (m *TransactionValidation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionValidation.Unmarshal(m, b)
}
func (m *TransactionValidation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionValidation.Marshal(b, m, deterministic)
}
func (dst *TransactionValidation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionValidation.Merge(dst, src)
}
func (m *TransactionValidation) XXX_Size() int {
	return xxx_messageInfo_TransactionValidation.Size(m)
}
func (m *TransactionValidation) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionValidation.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionValidation proto.InternalMessageInfo

func (m *TransactionValidation) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *TransactionValidation) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*TxRequest)(nil), "TxRequest")
	proto.RegisterType((*TxResponse)(nil), "TxResponse")
	proto.RegisterType((*TransactionValidation)(nil), "TransactionValidation")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AcceleratorServiceClient is the client API for AcceleratorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.

// 프로토콜 버퍼 컴파일러에 의해 생성된 RPC 서비스 인터페이스
type AcceleratorServiceClient interface {
	Execute(ctx context.Context, in *TxRequest, opts ...grpc.CallOption) (*TxResponse, error)
	Query(ctx context.Context, in *TxRequest, opts ...grpc.CallOption) (*TxResponse, error)
}

type acceleratorServiceClient struct {
	cc *grpc.ClientConn
}

func NewAcceleratorServiceClient(cc *grpc.ClientConn) AcceleratorServiceClient {
	return &acceleratorServiceClient{cc}
}

func (c *acceleratorServiceClient) Execute(ctx context.Context, in *TxRequest, opts ...grpc.CallOption) (*TxResponse, error) {
	// fmt.Println(ctx, in, opts) // ok
	// 응답을 넣을 TxResponse 타입의 객체 out 생성
	out := new(TxResponse)
	// Invoke() 정의: https://godoc.org/google.golang.org/grpc
	// "AcceleratorService/Execute": accelerator.proto에 정의되어 있는 서비스/메서드
	err := c.cc.Invoke(ctx, "/AcceleratorService/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *acceleratorServiceClient) Query(ctx context.Context, in *TxRequest, opts ...grpc.CallOption) (*TxResponse, error) {
	out := new(TxResponse)
	err := c.cc.Invoke(ctx, "/AcceleratorService/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AcceleratorServiceServer is the server API for AcceleratorService service.
type AcceleratorServiceServer interface {
	Execute(context.Context, *TxRequest) (*TxResponse, error)
	Query(context.Context, *TxRequest) (*TxResponse, error)
}

func RegisterAcceleratorServiceServer(s *grpc.Server, srv AcceleratorServiceServer) {
	s.RegisterService(&_AcceleratorService_serviceDesc, srv)
}

func _AcceleratorService_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AcceleratorServiceServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AcceleratorService/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AcceleratorServiceServer).Execute(ctx, req.(*TxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AcceleratorService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AcceleratorServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AcceleratorService/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AcceleratorServiceServer).Query(ctx, req.(*TxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AcceleratorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AcceleratorService",
	HandlerType: (*AcceleratorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _AcceleratorService_Execute_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _AcceleratorService_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accelerator.proto",
}

func init() { proto.RegisterFile("accelerator.proto", fileDescriptor_accelerator_61f348634469893b) }

var fileDescriptor_accelerator_61f348634469893b = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0xad, 0x6d, 0x2d, 0x9d, 0x56, 0xd0, 0x01, 0x65, 0x11, 0x0f, 0x65, 0x29, 0x58, 0x2f,
	0x5b, 0xa8, 0xe0, 0x5d, 0xc1, 0x43, 0x0f, 0x0a, 0xae, 0xc5, 0x83, 0x07, 0x61, 0x9a, 0x1d, 0xdb,
	0xc0, 0x36, 0x59, 0x93, 0x6c, 0xd9, 0xfe, 0x7b, 0x49, 0xaa, 0xed, 0x0a, 0xe2, 0x29, 0x2f, 0x6f,
	0x5e, 0xf8, 0x32, 0x33, 0x70, 0x4a, 0x42, 0x70, 0xce, 0x86, 0x9c, 0x36, 0x49, 0x61, 0xb4, 0xd3,
	0x71, 0x09, 0xdd, 0x59, 0x95, 0xf2, 0x67, 0xc9, 0xd6, 0xe1, 0x25, 0x74, 0xc5, 0x92, 0x94, 0xe2,
	0x7c, 0x9a, 0x45, 0x8d, 0x41, 0x63, 0xd4, 0x4d, 0xf7, 0x06, 0x0e, 0xe1, 0x58, 0x2c, 0x49, 0x2a,
	0xa1, 0x33, 0x7e, 0xa2, 0x15, 0x47, 0x87, 0x21, 0xf1, 0xdb, 0xc4, 0x13, 0x68, 0x7e, 0x08, 0x15,
	0x35, 0x43, 0xcd, 0x4b, 0x44, 0x68, 0x91, 0x59, 0xd8, 0xa8, 0x35, 0x68, 0x8e, 0xfa, 0x69, 0xd0,
	0xb1, 0x01, 0xf0, 0x58, 0x5b, 0x68, 0x65, 0x19, 0x23, 0xe8, 0x14, 0xb4, 0xc9, 0x35, 0x6d, 0xa9,
	0xfd, 0xf4, 0xe7, 0xea, 0xdf, 0xba, 0x6a, 0x9a, 0x7d, 0xa3, 0x82, 0xc6, 0x5b, 0x80, 0x35, 0xe5,
	0x32, 0x23, 0x27, 0xf5, 0x16, 0xd4, 0x9b, 0x9c, 0x27, 0x33, 0x43, 0xca, 0x92, 0xf0, 0xde, 0xeb,
	0xae, 0x9a, 0xd6, 0x92, 0xf1, 0x23, 0x9c, 0xfd, 0x19, 0xf2, 0x10, 0xff, 0xfd, 0xc0, 0x6e, 0xa7,
	0x41, 0xe3, 0x00, 0x7a, 0x19, 0x5b, 0x61, 0x64, 0x11, 0x28, 0x5b, 0x7e, 0xdd, 0x9a, 0xbc, 0x03,
	0xde, 0xed, 0xc7, 0xf9, 0xc2, 0x66, 0x2d, 0x05, 0xe3, 0x10, 0x3a, 0x0f, 0x15, 0x8b, 0xd2, 0x31,
	0x42, 0xb2, 0x9b, 0xec, 0x45, 0x2f, 0xd9, 0xb7, 0x1b, 0x1f, 0x60, 0x0c, 0xed, 0xe7, 0x92, 0xcd,
	0xe6, 0x9f, 0xcc, 0xfd, 0xf5, 0xdb, 0xd5, 0x42, 0xba, 0x65, 0x39, 0x4f, 0x84, 0x5e, 0x8d, 0x15,
	0x57, 0x39, 0x67, 0x0b, 0x36, 0xe3, 0xda, 0x0e, 0xc7, 0x61, 0x87, 0x76, 0x7e, 0x14, 0xce, 0x9b,
	0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x16, 0x0c, 0x60, 0xc3, 0xe0, 0x01, 0x00, 0x00,
}
