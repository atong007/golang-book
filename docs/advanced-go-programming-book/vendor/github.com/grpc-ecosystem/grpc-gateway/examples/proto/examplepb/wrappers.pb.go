// Code generated by protoc-gen-go. DO NOT EDIT.
// source: examples/proto/examplepb/wrappers.proto

package examplepb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type Wrappers struct {
	StringValue          *wrappers.StringValue `protobuf:"bytes,1,opt,name=string_value,json=stringValue" json:"string_value,omitempty"`
	Int32Value           *wrappers.Int32Value  `protobuf:"bytes,2,opt,name=int32_value,json=int32Value" json:"int32_value,omitempty"`
	Int64Value           *wrappers.Int64Value  `protobuf:"bytes,3,opt,name=int64_value,json=int64Value" json:"int64_value,omitempty"`
	FloatValue           *wrappers.FloatValue  `protobuf:"bytes,4,opt,name=float_value,json=floatValue" json:"float_value,omitempty"`
	DoubleValue          *wrappers.DoubleValue `protobuf:"bytes,5,opt,name=double_value,json=doubleValue" json:"double_value,omitempty"`
	BoolValue            *wrappers.BoolValue   `protobuf:"bytes,6,opt,name=bool_value,json=boolValue" json:"bool_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Wrappers) Reset()         { *m = Wrappers{} }
func (m *Wrappers) String() string { return proto.CompactTextString(m) }
func (*Wrappers) ProtoMessage()    {}
func (*Wrappers) Descriptor() ([]byte, []int) {
	return fileDescriptor_wrappers_c39e75a579a3cddc, []int{0}
}
func (m *Wrappers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Wrappers.Unmarshal(m, b)
}
func (m *Wrappers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Wrappers.Marshal(b, m, deterministic)
}
func (dst *Wrappers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Wrappers.Merge(dst, src)
}
func (m *Wrappers) XXX_Size() int {
	return xxx_messageInfo_Wrappers.Size(m)
}
func (m *Wrappers) XXX_DiscardUnknown() {
	xxx_messageInfo_Wrappers.DiscardUnknown(m)
}

var xxx_messageInfo_Wrappers proto.InternalMessageInfo

func (m *Wrappers) GetStringValue() *wrappers.StringValue {
	if m != nil {
		return m.StringValue
	}
	return nil
}

func (m *Wrappers) GetInt32Value() *wrappers.Int32Value {
	if m != nil {
		return m.Int32Value
	}
	return nil
}

func (m *Wrappers) GetInt64Value() *wrappers.Int64Value {
	if m != nil {
		return m.Int64Value
	}
	return nil
}

func (m *Wrappers) GetFloatValue() *wrappers.FloatValue {
	if m != nil {
		return m.FloatValue
	}
	return nil
}

func (m *Wrappers) GetDoubleValue() *wrappers.DoubleValue {
	if m != nil {
		return m.DoubleValue
	}
	return nil
}

func (m *Wrappers) GetBoolValue() *wrappers.BoolValue {
	if m != nil {
		return m.BoolValue
	}
	return nil
}

func init() {
	proto.RegisterType((*Wrappers)(nil), "grpc.gateway.examples.examplepb.Wrappers")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WrappersServiceClient is the client API for WrappersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WrappersServiceClient interface {
	Create(ctx context.Context, in *Wrappers, opts ...grpc.CallOption) (*Wrappers, error)
}

type wrappersServiceClient struct {
	cc *grpc.ClientConn
}

func NewWrappersServiceClient(cc *grpc.ClientConn) WrappersServiceClient {
	return &wrappersServiceClient{cc}
}

func (c *wrappersServiceClient) Create(ctx context.Context, in *Wrappers, opts ...grpc.CallOption) (*Wrappers, error) {
	out := new(Wrappers)
	err := c.cc.Invoke(ctx, "/grpc.gateway.examples.examplepb.WrappersService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WrappersServiceServer is the server API for WrappersService service.
type WrappersServiceServer interface {
	Create(context.Context, *Wrappers) (*Wrappers, error)
}

func RegisterWrappersServiceServer(s *grpc.Server, srv WrappersServiceServer) {
	s.RegisterService(&_WrappersService_serviceDesc, srv)
}

func _WrappersService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Wrappers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WrappersServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.examplepb.WrappersService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WrappersServiceServer).Create(ctx, req.(*Wrappers))
	}
	return interceptor(ctx, in, info, handler)
}

var _WrappersService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.gateway.examples.examplepb.WrappersService",
	HandlerType: (*WrappersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _WrappersService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "examples/proto/examplepb/wrappers.proto",
}

func init() {
	proto.RegisterFile("examples/proto/examplepb/wrappers.proto", fileDescriptor_wrappers_c39e75a579a3cddc)
}

var fileDescriptor_wrappers_c39e75a579a3cddc = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd2, 0xcf, 0x4a, 0x3b, 0x31,
	0x10, 0x07, 0x70, 0xb6, 0xfd, 0xfd, 0x8a, 0xcd, 0x0a, 0xc2, 0xe2, 0x41, 0xd7, 0x62, 0xa5, 0x17,
	0xff, 0x1c, 0xb2, 0xd8, 0x96, 0x82, 0x22, 0x08, 0x55, 0x04, 0xaf, 0x16, 0x14, 0xbc, 0x48, 0xd2,
	0xa6, 0x4b, 0x20, 0xee, 0x84, 0x6c, 0xda, 0xea, 0x49, 0xf4, 0x11, 0xf4, 0xd1, 0x7c, 0x05, 0x1f,
	0x44, 0x36, 0x9d, 0x6c, 0xc1, 0x52, 0xf4, 0x36, 0xb3, 0xf9, 0x7e, 0x06, 0x86, 0x1d, 0xb2, 0x2f,
	0x9e, 0xd8, 0xa3, 0x56, 0x22, 0x4f, 0xb4, 0x01, 0x0b, 0x09, 0xb6, 0x9a, 0x27, 0x33, 0xc3, 0xb4,
	0x16, 0x26, 0xa7, 0xee, 0x21, 0x6a, 0xa6, 0x46, 0x0f, 0x69, 0xca, 0xac, 0x98, 0xb1, 0x67, 0xea,
	0x15, 0x2d, 0xf3, 0x71, 0x23, 0x05, 0x48, 0x95, 0x48, 0x98, 0x96, 0x09, 0xcb, 0x32, 0xb0, 0xcc,
	0x4a, 0xc8, 0x90, 0xc7, 0xbb, 0xf8, 0xea, 0x3a, 0x3e, 0x19, 0xff, 0x18, 0xdf, 0x7a, 0xad, 0x92,
	0xb5, 0x3b, 0xfc, 0x14, 0x9d, 0x93, 0xf5, 0xdc, 0x1a, 0x99, 0xa5, 0x0f, 0x53, 0xa6, 0x26, 0x62,
	0x2b, 0xd8, 0x0b, 0x0e, 0xc2, 0x76, 0x83, 0xce, 0x67, 0x50, 0x3f, 0x83, 0x0e, 0x5c, 0xe8, 0xb6,
	0xc8, 0xdc, 0x84, 0xf9, 0xa2, 0x89, 0xce, 0x48, 0x28, 0x33, 0xdb, 0x69, 0xa3, 0xaf, 0x38, 0xbf,
	0xb3, 0xe4, 0xaf, 0x8b, 0xcc, 0x9c, 0x13, 0x59, 0xd6, 0xa8, 0x7b, 0x5d, 0xd4, 0xd5, 0xd5, 0xba,
	0xd7, 0x5d, 0x68, 0xac, 0x0b, 0x3d, 0x56, 0xc0, 0x2c, 0xea, 0x7f, 0x2b, 0xf4, 0x55, 0x91, 0x41,
	0x3d, 0x2e, 0xeb, 0x62, 0xf5, 0x11, 0x4c, 0xb8, 0x12, 0xc8, 0xff, 0xaf, 0x58, 0xfd, 0xd2, 0x85,
	0x70, 0xf5, 0xd1, 0xa2, 0x89, 0x4e, 0x08, 0xe1, 0x00, 0x0a, 0x79, 0xcd, 0xf1, 0x78, 0x89, 0xf7,
	0x01, 0xd4, 0x1c, 0xd7, 0xb9, 0x2f, 0xdb, 0xef, 0x01, 0xd9, 0xf0, 0xff, 0x60, 0x20, 0xcc, 0x54,
	0x0e, 0x45, 0xf4, 0x42, 0x6a, 0x17, 0x46, 0x30, 0x2b, 0xa2, 0x43, 0xfa, 0xcb, 0x05, 0x50, 0x6f,
	0xe3, 0xbf, 0x47, 0x5b, 0xcd, 0xb7, 0xcf, 0xaf, 0x8f, 0xca, 0x76, 0x6b, 0x33, 0x99, 0x1e, 0xfb,
	0xe3, 0x2b, 0x6f, 0xe3, 0x34, 0x38, 0xea, 0x87, 0xf7, 0xf5, 0x92, 0xf1, 0x9a, 0x5b, 0xa0, 0xf3,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0x3f, 0xa4, 0x0c, 0xe1, 0xb6, 0x02, 0x00, 0x00,
}
