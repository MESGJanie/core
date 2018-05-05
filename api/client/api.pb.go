// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/mesg-foundation/core/api/client/api.proto

/*
Package client is a generated protocol buffer package.

It is generated from these files:
	github.com/mesg-foundation/core/api/client/api.proto

It has these top-level messages:
	ListenEventRequest
	ExecuteTaskRequest
	ListenResultRequest
	StartServiceRequest
	StopServiceRequest
	EventData
	ExecuteTaskReply
	ResultData
	StartServiceReply
	StopServiceReply
*/
package client

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import service "github.com/mesg-foundation/core/service"

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

type ListenEventRequest struct {
	Service *service.Service `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
}

func (m *ListenEventRequest) Reset()                    { *m = ListenEventRequest{} }
func (m *ListenEventRequest) String() string            { return proto.CompactTextString(m) }
func (*ListenEventRequest) ProtoMessage()               {}
func (*ListenEventRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ListenEventRequest) GetService() *service.Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type ExecuteTaskRequest struct {
	Service  *service.Service `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	TaskKey  string           `protobuf:"bytes,2,opt,name=taskKey" json:"taskKey,omitempty"`
	TaskData string           `protobuf:"bytes,3,opt,name=taskData" json:"taskData,omitempty"`
}

func (m *ExecuteTaskRequest) Reset()                    { *m = ExecuteTaskRequest{} }
func (m *ExecuteTaskRequest) String() string            { return proto.CompactTextString(m) }
func (*ExecuteTaskRequest) ProtoMessage()               {}
func (*ExecuteTaskRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ExecuteTaskRequest) GetService() *service.Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *ExecuteTaskRequest) GetTaskKey() string {
	if m != nil {
		return m.TaskKey
	}
	return ""
}

func (m *ExecuteTaskRequest) GetTaskData() string {
	if m != nil {
		return m.TaskData
	}
	return ""
}

type ListenResultRequest struct {
	Service *service.Service `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
}

func (m *ListenResultRequest) Reset()                    { *m = ListenResultRequest{} }
func (m *ListenResultRequest) String() string            { return proto.CompactTextString(m) }
func (*ListenResultRequest) ProtoMessage()               {}
func (*ListenResultRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListenResultRequest) GetService() *service.Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type StartServiceRequest struct {
	Service *service.Service `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
}

func (m *StartServiceRequest) Reset()                    { *m = StartServiceRequest{} }
func (m *StartServiceRequest) String() string            { return proto.CompactTextString(m) }
func (*StartServiceRequest) ProtoMessage()               {}
func (*StartServiceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *StartServiceRequest) GetService() *service.Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type StopServiceRequest struct {
	Service *service.Service `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
}

func (m *StopServiceRequest) Reset()                    { *m = StopServiceRequest{} }
func (m *StopServiceRequest) String() string            { return proto.CompactTextString(m) }
func (*StopServiceRequest) ProtoMessage()               {}
func (*StopServiceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StopServiceRequest) GetService() *service.Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type EventData struct {
	Error     string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	EventKey  string `protobuf:"bytes,2,opt,name=eventKey" json:"eventKey,omitempty"`
	EventData string `protobuf:"bytes,3,opt,name=eventData" json:"eventData,omitempty"`
}

func (m *EventData) Reset()                    { *m = EventData{} }
func (m *EventData) String() string            { return proto.CompactTextString(m) }
func (*EventData) ProtoMessage()               {}
func (*EventData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *EventData) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *EventData) GetEventKey() string {
	if m != nil {
		return m.EventKey
	}
	return ""
}

func (m *EventData) GetEventData() string {
	if m != nil {
		return m.EventData
	}
	return ""
}

type ExecuteTaskReply struct {
	ExecutionID string `protobuf:"bytes,1,opt,name=executionID" json:"executionID,omitempty"`
	Error       string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *ExecuteTaskReply) Reset()                    { *m = ExecuteTaskReply{} }
func (m *ExecuteTaskReply) String() string            { return proto.CompactTextString(m) }
func (*ExecuteTaskReply) ProtoMessage()               {}
func (*ExecuteTaskReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ExecuteTaskReply) GetExecutionID() string {
	if m != nil {
		return m.ExecutionID
	}
	return ""
}

func (m *ExecuteTaskReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type ResultData struct {
	Error       string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	ExecutionID string `protobuf:"bytes,2,opt,name=executionID" json:"executionID,omitempty"`
	TaskKey     string `protobuf:"bytes,3,opt,name=taskKey" json:"taskKey,omitempty"`
	OutputKey   string `protobuf:"bytes,4,opt,name=outputKey" json:"outputKey,omitempty"`
	OutputData  string `protobuf:"bytes,5,opt,name=outputData" json:"outputData,omitempty"`
}

func (m *ResultData) Reset()                    { *m = ResultData{} }
func (m *ResultData) String() string            { return proto.CompactTextString(m) }
func (*ResultData) ProtoMessage()               {}
func (*ResultData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ResultData) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ResultData) GetExecutionID() string {
	if m != nil {
		return m.ExecutionID
	}
	return ""
}

func (m *ResultData) GetTaskKey() string {
	if m != nil {
		return m.TaskKey
	}
	return ""
}

func (m *ResultData) GetOutputKey() string {
	if m != nil {
		return m.OutputKey
	}
	return ""
}

func (m *ResultData) GetOutputData() string {
	if m != nil {
		return m.OutputData
	}
	return ""
}

type StartServiceReply struct {
	Error string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *StartServiceReply) Reset()                    { *m = StartServiceReply{} }
func (m *StartServiceReply) String() string            { return proto.CompactTextString(m) }
func (*StartServiceReply) ProtoMessage()               {}
func (*StartServiceReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *StartServiceReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type StopServiceReply struct {
	Error string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *StopServiceReply) Reset()                    { *m = StopServiceReply{} }
func (m *StopServiceReply) String() string            { return proto.CompactTextString(m) }
func (*StopServiceReply) ProtoMessage()               {}
func (*StopServiceReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *StopServiceReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*ListenEventRequest)(nil), "api.ListenEventRequest")
	proto.RegisterType((*ExecuteTaskRequest)(nil), "api.ExecuteTaskRequest")
	proto.RegisterType((*ListenResultRequest)(nil), "api.ListenResultRequest")
	proto.RegisterType((*StartServiceRequest)(nil), "api.StartServiceRequest")
	proto.RegisterType((*StopServiceRequest)(nil), "api.StopServiceRequest")
	proto.RegisterType((*EventData)(nil), "api.EventData")
	proto.RegisterType((*ExecuteTaskReply)(nil), "api.ExecuteTaskReply")
	proto.RegisterType((*ResultData)(nil), "api.ResultData")
	proto.RegisterType((*StartServiceReply)(nil), "api.StartServiceReply")
	proto.RegisterType((*StopServiceReply)(nil), "api.StopServiceReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Client service

type ClientClient interface {
	ListenEvent(ctx context.Context, in *ListenEventRequest, opts ...grpc.CallOption) (Client_ListenEventClient, error)
	ExecuteTask(ctx context.Context, in *ExecuteTaskRequest, opts ...grpc.CallOption) (*ExecuteTaskReply, error)
	ListenResult(ctx context.Context, in *ListenResultRequest, opts ...grpc.CallOption) (Client_ListenResultClient, error)
	StartService(ctx context.Context, in *StartServiceRequest, opts ...grpc.CallOption) (*StartServiceReply, error)
	StopService(ctx context.Context, in *StopServiceRequest, opts ...grpc.CallOption) (*StopServiceReply, error)
}

type clientClient struct {
	cc *grpc.ClientConn
}

func NewClientClient(cc *grpc.ClientConn) ClientClient {
	return &clientClient{cc}
}

func (c *clientClient) ListenEvent(ctx context.Context, in *ListenEventRequest, opts ...grpc.CallOption) (Client_ListenEventClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Client_serviceDesc.Streams[0], c.cc, "/api.Client/ListenEvent", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientListenEventClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Client_ListenEventClient interface {
	Recv() (*EventData, error)
	grpc.ClientStream
}

type clientListenEventClient struct {
	grpc.ClientStream
}

func (x *clientListenEventClient) Recv() (*EventData, error) {
	m := new(EventData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clientClient) ExecuteTask(ctx context.Context, in *ExecuteTaskRequest, opts ...grpc.CallOption) (*ExecuteTaskReply, error) {
	out := new(ExecuteTaskReply)
	err := grpc.Invoke(ctx, "/api.Client/ExecuteTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) ListenResult(ctx context.Context, in *ListenResultRequest, opts ...grpc.CallOption) (Client_ListenResultClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Client_serviceDesc.Streams[1], c.cc, "/api.Client/ListenResult", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientListenResultClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Client_ListenResultClient interface {
	Recv() (*ResultData, error)
	grpc.ClientStream
}

type clientListenResultClient struct {
	grpc.ClientStream
}

func (x *clientListenResultClient) Recv() (*ResultData, error) {
	m := new(ResultData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clientClient) StartService(ctx context.Context, in *StartServiceRequest, opts ...grpc.CallOption) (*StartServiceReply, error) {
	out := new(StartServiceReply)
	err := grpc.Invoke(ctx, "/api.Client/StartService", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) StopService(ctx context.Context, in *StopServiceRequest, opts ...grpc.CallOption) (*StopServiceReply, error) {
	out := new(StopServiceReply)
	err := grpc.Invoke(ctx, "/api.Client/StopService", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Client service

type ClientServer interface {
	ListenEvent(*ListenEventRequest, Client_ListenEventServer) error
	ExecuteTask(context.Context, *ExecuteTaskRequest) (*ExecuteTaskReply, error)
	ListenResult(*ListenResultRequest, Client_ListenResultServer) error
	StartService(context.Context, *StartServiceRequest) (*StartServiceReply, error)
	StopService(context.Context, *StopServiceRequest) (*StopServiceReply, error)
}

func RegisterClientServer(s *grpc.Server, srv ClientServer) {
	s.RegisterService(&_Client_serviceDesc, srv)
}

func _Client_ListenEvent_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenEventRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClientServer).ListenEvent(m, &clientListenEventServer{stream})
}

type Client_ListenEventServer interface {
	Send(*EventData) error
	grpc.ServerStream
}

type clientListenEventServer struct {
	grpc.ServerStream
}

func (x *clientListenEventServer) Send(m *EventData) error {
	return x.ServerStream.SendMsg(m)
}

func _Client_ExecuteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).ExecuteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Client/ExecuteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).ExecuteTask(ctx, req.(*ExecuteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_ListenResult_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenResultRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClientServer).ListenResult(m, &clientListenResultServer{stream})
}

type Client_ListenResultServer interface {
	Send(*ResultData) error
	grpc.ServerStream
}

type clientListenResultServer struct {
	grpc.ServerStream
}

func (x *clientListenResultServer) Send(m *ResultData) error {
	return x.ServerStream.SendMsg(m)
}

func _Client_StartService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).StartService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Client/StartService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).StartService(ctx, req.(*StartServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_StopService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).StopService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Client/StopService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).StopService(ctx, req.(*StopServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Client_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Client",
	HandlerType: (*ClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecuteTask",
			Handler:    _Client_ExecuteTask_Handler,
		},
		{
			MethodName: "StartService",
			Handler:    _Client_StartService_Handler,
		},
		{
			MethodName: "StopService",
			Handler:    _Client_StopService_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListenEvent",
			Handler:       _Client_ListenEvent_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListenResult",
			Handler:       _Client_ListenResult_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/mesg-foundation/core/api/client/api.proto",
}

func init() {
	proto.RegisterFile("github.com/mesg-foundation/core/api/client/api.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4b, 0x6f, 0xd3, 0x40,
	0x10, 0x8e, 0x13, 0x9a, 0xd6, 0xe3, 0x0a, 0xc2, 0x94, 0x87, 0x65, 0x55, 0x28, 0xda, 0x53, 0x40,
	0x22, 0x41, 0x05, 0x2e, 0x48, 0xa8, 0x0f, 0xda, 0x03, 0x8f, 0x93, 0xc3, 0x09, 0x4e, 0x5b, 0x33,
	0x94, 0x55, 0x53, 0xaf, 0xb1, 0xd7, 0x11, 0xf9, 0x2f, 0xfc, 0x21, 0xfe, 0x15, 0xf2, 0xae, 0x1f,
	0xeb, 0x38, 0x80, 0x14, 0x4e, 0xbb, 0xf3, 0xdc, 0x99, 0xef, 0x9b, 0x59, 0x78, 0x71, 0x25, 0xd4,
	0xb7, 0xfc, 0x72, 0x1a, 0xc9, 0x9b, 0xd9, 0x0d, 0x65, 0x57, 0x4f, 0xbf, 0xca, 0x3c, 0xfe, 0xc2,
	0x95, 0x90, 0xf1, 0x2c, 0x92, 0x29, 0xcd, 0x78, 0x22, 0x66, 0xd1, 0x42, 0x50, 0xac, 0x8a, 0xeb,
	0x34, 0x49, 0xa5, 0x92, 0x38, 0xe0, 0x89, 0x08, 0x5e, 0xfe, 0x2b, 0x34, 0xa3, 0x74, 0x29, 0xa2,
	0xfa, 0x34, 0xb1, 0xec, 0x04, 0xf0, 0x83, 0xc8, 0x14, 0xc5, 0x17, 0x4b, 0x8a, 0x55, 0x48, 0xdf,
	0x73, 0xca, 0x14, 0x3e, 0x81, 0xdd, 0xd2, 0xcd, 0x77, 0xc6, 0xce, 0xc4, 0x3b, 0x1a, 0x4d, 0xab,
	0xb0, 0xb9, 0x39, 0xc3, 0xca, 0x81, 0x2d, 0x01, 0x2f, 0x7e, 0x50, 0x94, 0x2b, 0xfa, 0xc8, 0xb3,
	0xeb, 0x2d, 0x32, 0xa0, 0x0f, 0xbb, 0x8a, 0x67, 0xd7, 0xef, 0x69, 0xe5, 0xf7, 0xc7, 0xce, 0xc4,
	0x0d, 0x2b, 0x11, 0x03, 0xd8, 0x2b, 0xae, 0xe7, 0x5c, 0x71, 0x7f, 0xa0, 0x4d, 0xb5, 0xcc, 0x4e,
	0xe1, 0xc0, 0x54, 0x1e, 0x52, 0x96, 0x2f, 0xb6, 0x2a, 0xfd, 0x14, 0x0e, 0xe6, 0x8a, 0xa7, 0xaa,
	0x32, 0x6c, 0x91, 0xe2, 0x04, 0x70, 0xae, 0x64, 0xf2, 0x1f, 0x19, 0x3e, 0x83, 0xab, 0xb1, 0x2f,
	0x9a, 0xc2, 0x7b, 0xb0, 0x43, 0x69, 0x2a, 0x53, 0x1d, 0xe6, 0x86, 0x46, 0x28, 0x60, 0xa0, 0xc2,
	0xa5, 0x41, 0xa8, 0x96, 0xf1, 0x10, 0x5c, 0xaa, 0xc2, 0x4b, 0x8c, 0x1a, 0x05, 0x7b, 0x07, 0xa3,
	0x16, 0x39, 0xc9, 0x62, 0x85, 0x63, 0xf0, 0x48, 0xeb, 0x84, 0x8c, 0xdf, 0x9e, 0x97, 0x2f, 0xd9,
	0xaa, 0xa6, 0x8a, 0xbe, 0x55, 0x05, 0xfb, 0xe9, 0x00, 0x18, 0xac, 0xff, 0x52, 0xea, 0x5a, 0xf2,
	0x7e, 0x37, 0xb9, 0xc5, 0xf6, 0xa0, 0xcd, 0xf6, 0x21, 0xb8, 0x32, 0x57, 0x49, 0xae, 0xfb, 0xbc,
	0x65, 0x5a, 0xa9, 0x15, 0xf8, 0x08, 0xc0, 0x08, 0xba, 0xd3, 0x1d, 0x6d, 0xb6, 0x34, 0xec, 0x31,
	0xdc, 0x6d, 0x93, 0x59, 0xf4, 0xba, 0xb1, 0x48, 0x36, 0x81, 0x51, 0x8b, 0xb4, 0x3f, 0x7a, 0x1e,
	0xfd, 0xea, 0xc3, 0xf0, 0x8d, 0xde, 0x37, 0x7c, 0x05, 0x9e, 0xb5, 0x29, 0xf8, 0x70, 0x5a, 0x2c,
	0x60, 0x77, 0x77, 0x82, 0xdb, 0xda, 0x50, 0x53, 0xca, 0x7a, 0xcf, 0x1c, 0x3c, 0x06, 0xcf, 0xa2,
	0xa1, 0x8c, 0xed, 0x6e, 0x4d, 0x70, 0xbf, 0x6b, 0x48, 0x16, 0x2b, 0xd6, 0xc3, 0xd7, 0xb0, 0x6f,
	0x0f, 0x3b, 0xfa, 0xd6, 0xeb, 0xad, 0xf9, 0x0f, 0xee, 0x68, 0x4b, 0xc3, 0x93, 0x7e, 0xff, 0x0c,
	0xf6, 0x6d, 0x6c, 0xca, 0xf0, 0x0d, 0xb3, 0x1f, 0x3c, 0xd8, 0x60, 0x31, 0x25, 0x1c, 0x83, 0x67,
	0x81, 0x56, 0xf6, 0xd0, 0x9d, 0xfd, 0xb2, 0x87, 0x75, 0x7c, 0x59, 0xef, 0x6c, 0xef, 0xd3, 0xd0,
	0x7c, 0x5d, 0x97, 0x43, 0xfd, 0xf7, 0x3c, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x01, 0xe1,
	0x59, 0xef, 0x04, 0x00, 0x00,
}