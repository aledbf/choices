// Code generated by protoc-gen-go.
// source: elwin.proto
// DO NOT EDIT!

/*
Package elwin is a generated protocol buffer package.

It is generated from these files:
	elwin.proto

It has these top-level messages:
	Identifier
	Experiments
	Experiment
	Param
*/
package elwin

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

type Identifier struct {
	TeamID string `protobuf:"bytes,1,opt,name=teamID" json:"teamID,omitempty"`
	UserID string `protobuf:"bytes,2,opt,name=userID" json:"userID,omitempty"`
}

func (m *Identifier) Reset()                    { *m = Identifier{} }
func (m *Identifier) String() string            { return proto.CompactTextString(m) }
func (*Identifier) ProtoMessage()               {}
func (*Identifier) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Experiments struct {
	Experiments map[string]*Experiment `protobuf:"bytes,1,rep,name=experiments" json:"experiments,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Experiments) Reset()                    { *m = Experiments{} }
func (m *Experiments) String() string            { return proto.CompactTextString(m) }
func (*Experiments) ProtoMessage()               {}
func (*Experiments) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Experiments) GetExperiments() map[string]*Experiment {
	if m != nil {
		return m.Experiments
	}
	return nil
}

type Experiment struct {
	Params []*Param `protobuf:"bytes,2,rep,name=params" json:"params,omitempty"`
}

func (m *Experiment) Reset()                    { *m = Experiment{} }
func (m *Experiment) String() string            { return proto.CompactTextString(m) }
func (*Experiment) ProtoMessage()               {}
func (*Experiment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Experiment) GetParams() []*Param {
	if m != nil {
		return m.Params
	}
	return nil
}

type Param struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Param) Reset()                    { *m = Param{} }
func (m *Param) String() string            { return proto.CompactTextString(m) }
func (*Param) ProtoMessage()               {}
func (*Param) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*Identifier)(nil), "elwin.Identifier")
	proto.RegisterType((*Experiments)(nil), "elwin.Experiments")
	proto.RegisterType((*Experiment)(nil), "elwin.Experiment")
	proto.RegisterType((*Param)(nil), "elwin.Param")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Elwin service

type ElwinClient interface {
	GetNamespaces(ctx context.Context, in *Identifier, opts ...grpc.CallOption) (*Experiments, error)
}

type elwinClient struct {
	cc *grpc.ClientConn
}

func NewElwinClient(cc *grpc.ClientConn) ElwinClient {
	return &elwinClient{cc}
}

func (c *elwinClient) GetNamespaces(ctx context.Context, in *Identifier, opts ...grpc.CallOption) (*Experiments, error) {
	out := new(Experiments)
	err := grpc.Invoke(ctx, "/elwin.Elwin/GetNamespaces", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Elwin service

type ElwinServer interface {
	GetNamespaces(context.Context, *Identifier) (*Experiments, error)
}

func RegisterElwinServer(s *grpc.Server, srv ElwinServer) {
	s.RegisterService(&_Elwin_serviceDesc, srv)
}

func _Elwin_GetNamespaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Identifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElwinServer).GetNamespaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/elwin.Elwin/GetNamespaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElwinServer).GetNamespaces(ctx, req.(*Identifier))
	}
	return interceptor(ctx, in, info, handler)
}

var _Elwin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "elwin.Elwin",
	HandlerType: (*ElwinServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNamespaces",
			Handler:    _Elwin_GetNamespaces_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("elwin.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x91, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x4d, 0x4b, 0x02, 0xbd, 0x51, 0xa8, 0x17, 0x11, 0xe9, 0x4a, 0x46, 0xc1, 0xae, 0x2a,
	0xc4, 0x4d, 0x11, 0x37, 0x82, 0x41, 0xba, 0x11, 0xcd, 0x1b, 0x8c, 0x7a, 0x85, 0x60, 0x27, 0x1d,
	0x66, 0xa6, 0x6a, 0x9f, 0xc8, 0xd7, 0x74, 0xfe, 0x24, 0x83, 0x71, 0x37, 0xe7, 0xdc, 0x9f, 0xef,
	0x70, 0x07, 0x4a, 0x5a, 0x7f, 0xb6, 0xdd, 0x42, 0xaa, 0x8d, 0xd9, 0x60, 0xee, 0x05, 0xbb, 0x01,
	0x58, 0xbd, 0x52, 0x67, 0xda, 0xb7, 0x96, 0x14, 0x1e, 0x43, 0x61, 0x88, 0x8b, 0xd5, 0xdd, 0x49,
	0x76, 0x9a, 0xcd, 0x27, 0x4d, 0x54, 0xce, 0xdf, 0x6a, 0x52, 0xd6, 0x1f, 0x05, 0x3f, 0x28, 0xf6,
	0x9d, 0x41, 0x59, 0x7f, 0x49, 0x52, 0xad, 0xb0, 0x3b, 0x34, 0xd6, 0x96, 0xd1, 0x4b, 0xbb, 0x64,
	0x3c, 0x2f, 0xab, 0xb3, 0x45, 0xe0, 0x26, 0x8d, 0xe9, 0xbb, 0xee, 0x8c, 0xda, 0x35, 0xe9, 0xdc,
	0xec, 0x09, 0xa6, 0x7f, 0x1b, 0x70, 0x0a, 0xe3, 0x77, 0xda, 0xc5, 0x5c, 0xee, 0x89, 0x17, 0x90,
	0x7f, 0xf0, 0xf5, 0x96, 0x7c, 0xa6, 0xb2, 0x3a, 0x1c, 0x60, 0x9a, 0x50, 0xbf, 0x1e, 0x2d, 0x33,
	0x56, 0x01, 0xf4, 0x05, 0x3c, 0x87, 0x42, 0x72, 0xc5, 0x85, 0xb6, 0xb3, 0x2e, 0xe2, 0x7e, 0x9c,
	0x7d, 0x74, 0x66, 0x13, 0x6b, 0xec, 0x12, 0x72, 0x6f, 0xfc, 0xc3, 0x3e, 0x4a, 0xd9, 0x93, 0x08,
	0xaa, 0x6e, 0x21, 0xaf, 0xdd, 0x1e, 0x5c, 0xc2, 0xc1, 0x3d, 0x99, 0x07, 0x2e, 0x48, 0x4b, 0xfe,
	0x42, 0x1a, 0x7f, 0xc3, 0xf5, 0xb7, 0x9e, 0xe1, 0xf0, 0x2c, 0x6c, 0xef, 0xb9, 0xf0, 0xbf, 0x73,
	0xf5, 0x13, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x74, 0xe8, 0xb0, 0xac, 0x01, 0x00, 0x00,
}
