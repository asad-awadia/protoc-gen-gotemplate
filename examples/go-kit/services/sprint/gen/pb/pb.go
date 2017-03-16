// Code generated by protoc-gen-go.
// source: services/sprint/sprint.proto
// DO NOT EDIT!

/*
Package sprint is a generated protocol buffer package.

It is generated from these files:
	services/sprint/sprint.proto

It has these top-level messages:
	AddSprintRequest
	AddSprintResponse
	CloseSprintRequest
	CloseSprintResponse
	GetSprintRequest
	GetSprintResponse
	Sprint
*/
package sprint

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

type AddSprintRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *AddSprintRequest) Reset()                    { *m = AddSprintRequest{} }
func (m *AddSprintRequest) String() string            { return proto.CompactTextString(m) }
func (*AddSprintRequest) ProtoMessage()               {}
func (*AddSprintRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddSprintRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AddSprintResponse struct {
	Sprint *Sprint `protobuf:"bytes,1,opt,name=sprint" json:"sprint,omitempty"`
	ErrMsg string  `protobuf:"bytes,2,opt,name=err_msg,json=errMsg" json:"err_msg,omitempty"`
}

func (m *AddSprintResponse) Reset()                    { *m = AddSprintResponse{} }
func (m *AddSprintResponse) String() string            { return proto.CompactTextString(m) }
func (*AddSprintResponse) ProtoMessage()               {}
func (*AddSprintResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddSprintResponse) GetSprint() *Sprint {
	if m != nil {
		return m.Sprint
	}
	return nil
}

func (m *AddSprintResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type CloseSprintRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *CloseSprintRequest) Reset()                    { *m = CloseSprintRequest{} }
func (m *CloseSprintRequest) String() string            { return proto.CompactTextString(m) }
func (*CloseSprintRequest) ProtoMessage()               {}
func (*CloseSprintRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CloseSprintRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CloseSprintResponse struct {
	ErrMsg string `protobuf:"bytes,1,opt,name=err_msg,json=errMsg" json:"err_msg,omitempty"`
}

func (m *CloseSprintResponse) Reset()                    { *m = CloseSprintResponse{} }
func (m *CloseSprintResponse) String() string            { return proto.CompactTextString(m) }
func (*CloseSprintResponse) ProtoMessage()               {}
func (*CloseSprintResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CloseSprintResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type GetSprintRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetSprintRequest) Reset()                    { *m = GetSprintRequest{} }
func (m *GetSprintRequest) String() string            { return proto.CompactTextString(m) }
func (*GetSprintRequest) ProtoMessage()               {}
func (*GetSprintRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetSprintRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetSprintResponse struct {
	Sprint *Sprint `protobuf:"bytes,1,opt,name=sprint" json:"sprint,omitempty"`
	ErrMsg string  `protobuf:"bytes,2,opt,name=err_msg,json=errMsg" json:"err_msg,omitempty"`
}

func (m *GetSprintResponse) Reset()                    { *m = GetSprintResponse{} }
func (m *GetSprintResponse) String() string            { return proto.CompactTextString(m) }
func (*GetSprintResponse) ProtoMessage()               {}
func (*GetSprintResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetSprintResponse) GetSprint() *Sprint {
	if m != nil {
		return m.Sprint
	}
	return nil
}

func (m *GetSprintResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type Sprint struct {
	Id        string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	CreatedAt uint32 `protobuf:"varint,2,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *Sprint) Reset()                    { *m = Sprint{} }
func (m *Sprint) String() string            { return proto.CompactTextString(m) }
func (*Sprint) ProtoMessage()               {}
func (*Sprint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Sprint) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Sprint) GetCreatedAt() uint32 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Sprint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*AddSprintRequest)(nil), "sprint.AddSprintRequest")
	proto.RegisterType((*AddSprintResponse)(nil), "sprint.AddSprintResponse")
	proto.RegisterType((*CloseSprintRequest)(nil), "sprint.CloseSprintRequest")
	proto.RegisterType((*CloseSprintResponse)(nil), "sprint.CloseSprintResponse")
	proto.RegisterType((*GetSprintRequest)(nil), "sprint.GetSprintRequest")
	proto.RegisterType((*GetSprintResponse)(nil), "sprint.GetSprintResponse")
	proto.RegisterType((*Sprint)(nil), "sprint.Sprint")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SprintService service

type SprintServiceClient interface {
	AddSprint(ctx context.Context, in *AddSprintRequest, opts ...grpc.CallOption) (*AddSprintResponse, error)
	CloseSprint(ctx context.Context, in *CloseSprintRequest, opts ...grpc.CallOption) (*CloseSprintResponse, error)
	GetSprint(ctx context.Context, in *GetSprintRequest, opts ...grpc.CallOption) (*GetSprintResponse, error)
}

type sprintServiceClient struct {
	cc *grpc.ClientConn
}

func NewSprintServiceClient(cc *grpc.ClientConn) SprintServiceClient {
	return &sprintServiceClient{cc}
}

func (c *sprintServiceClient) AddSprint(ctx context.Context, in *AddSprintRequest, opts ...grpc.CallOption) (*AddSprintResponse, error) {
	out := new(AddSprintResponse)
	err := grpc.Invoke(ctx, "/sprint.SprintService/AddSprint", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sprintServiceClient) CloseSprint(ctx context.Context, in *CloseSprintRequest, opts ...grpc.CallOption) (*CloseSprintResponse, error) {
	out := new(CloseSprintResponse)
	err := grpc.Invoke(ctx, "/sprint.SprintService/CloseSprint", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sprintServiceClient) GetSprint(ctx context.Context, in *GetSprintRequest, opts ...grpc.CallOption) (*GetSprintResponse, error) {
	out := new(GetSprintResponse)
	err := grpc.Invoke(ctx, "/sprint.SprintService/GetSprint", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SprintService service

type SprintServiceServer interface {
	AddSprint(context.Context, *AddSprintRequest) (*AddSprintResponse, error)
	CloseSprint(context.Context, *CloseSprintRequest) (*CloseSprintResponse, error)
	GetSprint(context.Context, *GetSprintRequest) (*GetSprintResponse, error)
}

func RegisterSprintServiceServer(s *grpc.Server, srv SprintServiceServer) {
	s.RegisterService(&_SprintService_serviceDesc, srv)
}

func _SprintService_AddSprint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSprintRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SprintServiceServer).AddSprint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sprint.SprintService/AddSprint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SprintServiceServer).AddSprint(ctx, req.(*AddSprintRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SprintService_CloseSprint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseSprintRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SprintServiceServer).CloseSprint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sprint.SprintService/CloseSprint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SprintServiceServer).CloseSprint(ctx, req.(*CloseSprintRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SprintService_GetSprint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSprintRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SprintServiceServer).GetSprint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sprint.SprintService/GetSprint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SprintServiceServer).GetSprint(ctx, req.(*GetSprintRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SprintService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sprint.SprintService",
	HandlerType: (*SprintServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddSprint",
			Handler:    _SprintService_AddSprint_Handler,
		},
		{
			MethodName: "CloseSprint",
			Handler:    _SprintService_CloseSprint_Handler,
		},
		{
			MethodName: "GetSprint",
			Handler:    _SprintService_GetSprint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/sprint/sprint.proto",
}

func init() { proto.RegisterFile("services/sprint/sprint.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x6d, 0xaa, 0x44, 0x32, 0xa5, 0xa5, 0x1d, 0x0f, 0xc6, 0xa8, 0x20, 0x8b, 0x14, 0x4f, 0x11,
	0xea, 0x2f, 0x68, 0x3d, 0x28, 0x88, 0x97, 0xb4, 0xf7, 0x12, 0xbb, 0x43, 0x09, 0xd8, 0x24, 0xee,
	0xac, 0xfe, 0x5f, 0xff, 0x89, 0xb0, 0xbb, 0xcd, 0x57, 0x8b, 0x27, 0x4f, 0xc9, 0xec, 0xbc, 0x7d,
	0x6f, 0xe6, 0xbd, 0x85, 0x6b, 0x26, 0xf5, 0x9d, 0x6d, 0x88, 0x1f, 0xb8, 0x54, 0x59, 0xae, 0xdd,
	0x27, 0x2e, 0x55, 0xa1, 0x0b, 0xf4, 0x6d, 0x25, 0xa6, 0x30, 0x9e, 0x4b, 0xb9, 0x34, 0x45, 0x42,
	0x9f, 0x5f, 0xc4, 0x1a, 0x11, 0x4e, 0xf3, 0x74, 0x47, 0xa1, 0x77, 0xeb, 0xdd, 0x07, 0x89, 0xf9,
	0x17, 0x2b, 0x98, 0x34, 0x70, 0x5c, 0x16, 0x39, 0x13, 0x4e, 0xc1, 0xd1, 0x18, 0xe8, 0x60, 0x36,
	0x8a, 0x9d, 0x86, 0xc3, 0xb9, 0x2e, 0x5e, 0xc0, 0x19, 0x29, 0xb5, 0xde, 0xf1, 0x36, 0xec, 0x1b,
	0x4e, 0x9f, 0x94, 0x7a, 0xe3, 0xad, 0xb8, 0x03, 0x7c, 0xfa, 0x28, 0x98, 0xda, 0xfa, 0x23, 0xe8,
	0x67, 0xd2, 0xa9, 0xf7, 0x33, 0x29, 0x62, 0x38, 0x6f, 0xa1, 0x9c, 0x7a, 0x83, 0xd5, 0x6b, 0xb1,
	0x0a, 0x18, 0x3f, 0x93, 0xfe, 0x9b, 0x73, 0x05, 0x93, 0x06, 0xe6, 0xbf, 0xf6, 0x79, 0x05, 0xdf,
	0x42, 0xbb, 0x7a, 0x78, 0x03, 0xb0, 0x51, 0x94, 0x6a, 0x92, 0xeb, 0x54, 0x9b, 0x5b, 0xc3, 0x24,
	0x70, 0x27, 0xf3, 0xda, 0xf2, 0x93, 0xda, 0xf2, 0xd9, 0x8f, 0x07, 0x43, 0xcb, 0xb6, 0xb4, 0x49,
	0xe2, 0x02, 0x82, 0x2a, 0x04, 0x0c, 0xf7, 0xc3, 0x75, 0xf3, 0x8b, 0x2e, 0x8f, 0x74, 0xec, 0x86,
	0xa2, 0x87, 0x2f, 0x30, 0x68, 0x98, 0x89, 0xd1, 0x1e, 0x7b, 0x98, 0x43, 0x74, 0x75, 0xb4, 0x57,
	0x31, 0x2d, 0x20, 0xa8, 0x2c, 0xac, 0xa7, 0xe9, 0x3a, 0x5f, 0x4f, 0x73, 0xe0, 0xb7, 0xe8, 0xbd,
	0xfb, 0xe6, 0x35, 0x3e, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x9e, 0xb2, 0x1e, 0xad, 0x02,
	0x00, 0x00,
}
