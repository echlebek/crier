// Code generated by protoc-gen-go. DO NOT EDIT.
// source: crier.proto

/*
Package crier is a generated protocol buffer package.

It is generated from these files:
	crier.proto

It has these top-level messages:
	Cry
	Hark
*/
package crier

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

type Cry struct {
	Time   int64  `protobuf:"varint,1,opt,name=Time" json:"Time,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=Status" json:"Status,omitempty"`
}

func (m *Cry) Reset()                    { *m = Cry{} }
func (m *Cry) String() string            { return proto.CompactTextString(m) }
func (*Cry) ProtoMessage()               {}
func (*Cry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Cry) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Cry) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type Hark struct {
	Harkening string `protobuf:"bytes,1,opt,name=Harkening" json:"Harkening,omitempty"`
}

func (m *Hark) Reset()                    { *m = Hark{} }
func (m *Hark) String() string            { return proto.CompactTextString(m) }
func (*Hark) ProtoMessage()               {}
func (*Hark) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Hark) GetHarkening() string {
	if m != nil {
		return m.Harkening
	}
	return ""
}

func init() {
	proto.RegisterType((*Cry)(nil), "Cry")
	proto.RegisterType((*Hark)(nil), "Hark")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Crier service

type CrierClient interface {
	ListenForCries(ctx context.Context, opts ...grpc.CallOption) (Crier_ListenForCriesClient, error)
}

type crierClient struct {
	cc *grpc.ClientConn
}

func NewCrierClient(cc *grpc.ClientConn) CrierClient {
	return &crierClient{cc}
}

func (c *crierClient) ListenForCries(ctx context.Context, opts ...grpc.CallOption) (Crier_ListenForCriesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Crier_serviceDesc.Streams[0], c.cc, "/Crier/ListenForCries", opts...)
	if err != nil {
		return nil, err
	}
	x := &crierListenForCriesClient{stream}
	return x, nil
}

type Crier_ListenForCriesClient interface {
	Send(*Hark) error
	Recv() (*Cry, error)
	grpc.ClientStream
}

type crierListenForCriesClient struct {
	grpc.ClientStream
}

func (x *crierListenForCriesClient) Send(m *Hark) error {
	return x.ClientStream.SendMsg(m)
}

func (x *crierListenForCriesClient) Recv() (*Cry, error) {
	m := new(Cry)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Crier service

type CrierServer interface {
	ListenForCries(Crier_ListenForCriesServer) error
}

func RegisterCrierServer(s *grpc.Server, srv CrierServer) {
	s.RegisterService(&_Crier_serviceDesc, srv)
}

func _Crier_ListenForCries_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CrierServer).ListenForCries(&crierListenForCriesServer{stream})
}

type Crier_ListenForCriesServer interface {
	Send(*Cry) error
	Recv() (*Hark, error)
	grpc.ServerStream
}

type crierListenForCriesServer struct {
	grpc.ServerStream
}

func (x *crierListenForCriesServer) Send(m *Cry) error {
	return x.ServerStream.SendMsg(m)
}

func (x *crierListenForCriesServer) Recv() (*Hark, error) {
	m := new(Hark)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Crier_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Crier",
	HandlerType: (*CrierServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListenForCries",
			Handler:       _Crier_ListenForCries_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "crier.proto",
}

func init() { proto.RegisterFile("crier.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2e, 0xca, 0x4c,
	0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x32, 0xe4, 0x62, 0x76, 0x2e, 0xaa, 0x14, 0x12,
	0xe2, 0x62, 0x09, 0xc9, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x02, 0xb3, 0x85,
	0xc4, 0xb8, 0xd8, 0x82, 0x4b, 0x12, 0x4b, 0x4a, 0x8b, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83,
	0xa0, 0x3c, 0x25, 0x15, 0x2e, 0x16, 0x8f, 0xc4, 0xa2, 0x6c, 0x21, 0x19, 0x2e, 0x4e, 0x10, 0x9d,
	0x9a, 0x97, 0x99, 0x97, 0x0e, 0xd6, 0xc8, 0x19, 0x84, 0x10, 0x30, 0xd2, 0xe1, 0x62, 0x75, 0x06,
	0xd9, 0x23, 0xa4, 0xcc, 0xc5, 0xe7, 0x93, 0x59, 0x5c, 0x92, 0x9a, 0xe7, 0x96, 0x5f, 0x04, 0x12,
	0x29, 0x16, 0x62, 0xd5, 0x03, 0x29, 0x93, 0x62, 0xd1, 0x73, 0x2e, 0xaa, 0x54, 0x62, 0xd0, 0x60,
	0x34, 0x60, 0x4c, 0x62, 0x03, 0xbb, 0xc6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x4f, 0xa3,
	0x1a, 0x9c, 0x00, 0x00, 0x00,
}
