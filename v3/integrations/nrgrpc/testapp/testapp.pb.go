// Code generated by protoc-gen-go. DO NOT EDIT.
// source: testapp.proto

package testapp

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Message struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_98d4e818d9f182b1, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "Message")
}

func init() { proto.RegisterFile("testapp.proto", fileDescriptor_98d4e818d9f182b1) }

var fileDescriptor_98d4e818d9f182b1 = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x49, 0x2d, 0x2e,
	0x49, 0x2c, 0x28, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x92, 0xe5, 0x62, 0xf7, 0x4d, 0x2d,
	0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x12, 0xe2, 0x62, 0x29, 0x49, 0xad, 0x28, 0x91, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x8d, 0xfa, 0x98, 0xb8, 0xf8, 0x43, 0x52, 0x8b, 0x4b, 0x1c, 0x0b,
	0x0a, 0x72, 0x32, 0x93, 0x13, 0x4b, 0x32, 0xf3, 0xf3, 0x84, 0x54, 0xb8, 0x78, 0x5c, 0xf2, 0x43,
	0xf3, 0x12, 0x8b, 0x2a, 0xc1, 0x84, 0x10, 0x87, 0x1e, 0xd4, 0x04, 0x29, 0x38, 0x4b, 0x89, 0x41,
	0x48, 0x9d, 0x8b, 0x17, 0xaa, 0x2a, 0xb8, 0xa4, 0x28, 0x35, 0x31, 0x17, 0xbb, 0x32, 0x03, 0x46,
	0x88, 0x42, 0x88, 0x1a, 0x3c, 0xe6, 0x69, 0x30, 0x0a, 0x69, 0x71, 0xf1, 0xc1, 0x14, 0xe2, 0x33,
	0x52, 0x83, 0xd1, 0x80, 0x51, 0x48, 0x93, 0x4b, 0x10, 0xd9, 0x8d, 0xae, 0x45, 0x45, 0xf9, 0x45,
	0x38, 0x1c, 0xaa, 0xc3, 0x25, 0x84, 0xe2, 0x50, 0x3c, 0x6a, 0x0d, 0x18, 0x93, 0xd8, 0xc0, 0xc1,
	0x66, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xca, 0x5d, 0x32, 0x47, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestApplicationClient is the client API for TestApplication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestApplicationClient interface {
	DoUnaryUnary(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	DoUnaryStream(ctx context.Context, in *Message, opts ...grpc.CallOption) (TestApplication_DoUnaryStreamClient, error)
	DoStreamUnary(ctx context.Context, opts ...grpc.CallOption) (TestApplication_DoStreamUnaryClient, error)
	DoStreamStream(ctx context.Context, opts ...grpc.CallOption) (TestApplication_DoStreamStreamClient, error)
	DoUnaryUnaryError(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	DoUnaryStreamError(ctx context.Context, in *Message, opts ...grpc.CallOption) (TestApplication_DoUnaryStreamErrorClient, error)
}

type testApplicationClient struct {
	cc *grpc.ClientConn
}

func NewTestApplicationClient(cc *grpc.ClientConn) TestApplicationClient {
	return &testApplicationClient{cc}
}

func (c *testApplicationClient) DoUnaryUnary(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/TestApplication/DoUnaryUnary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testApplicationClient) DoUnaryStream(ctx context.Context, in *Message, opts ...grpc.CallOption) (TestApplication_DoUnaryStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestApplication_serviceDesc.Streams[0], "/TestApplication/DoUnaryStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &testApplicationDoUnaryStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestApplication_DoUnaryStreamClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type testApplicationDoUnaryStreamClient struct {
	grpc.ClientStream
}

func (x *testApplicationDoUnaryStreamClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testApplicationClient) DoStreamUnary(ctx context.Context, opts ...grpc.CallOption) (TestApplication_DoStreamUnaryClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestApplication_serviceDesc.Streams[1], "/TestApplication/DoStreamUnary", opts...)
	if err != nil {
		return nil, err
	}
	x := &testApplicationDoStreamUnaryClient{stream}
	return x, nil
}

type TestApplication_DoStreamUnaryClient interface {
	Send(*Message) error
	CloseAndRecv() (*Message, error)
	grpc.ClientStream
}

type testApplicationDoStreamUnaryClient struct {
	grpc.ClientStream
}

func (x *testApplicationDoStreamUnaryClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testApplicationDoStreamUnaryClient) CloseAndRecv() (*Message, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testApplicationClient) DoStreamStream(ctx context.Context, opts ...grpc.CallOption) (TestApplication_DoStreamStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestApplication_serviceDesc.Streams[2], "/TestApplication/DoStreamStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &testApplicationDoStreamStreamClient{stream}
	return x, nil
}

type TestApplication_DoStreamStreamClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type testApplicationDoStreamStreamClient struct {
	grpc.ClientStream
}

func (x *testApplicationDoStreamStreamClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testApplicationDoStreamStreamClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testApplicationClient) DoUnaryUnaryError(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/TestApplication/DoUnaryUnaryError", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testApplicationClient) DoUnaryStreamError(ctx context.Context, in *Message, opts ...grpc.CallOption) (TestApplication_DoUnaryStreamErrorClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestApplication_serviceDesc.Streams[3], "/TestApplication/DoUnaryStreamError", opts...)
	if err != nil {
		return nil, err
	}
	x := &testApplicationDoUnaryStreamErrorClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestApplication_DoUnaryStreamErrorClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type testApplicationDoUnaryStreamErrorClient struct {
	grpc.ClientStream
}

func (x *testApplicationDoUnaryStreamErrorClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TestApplicationServer is the server API for TestApplication service.
type TestApplicationServer interface {
	DoUnaryUnary(context.Context, *Message) (*Message, error)
	DoUnaryStream(*Message, TestApplication_DoUnaryStreamServer) error
	DoStreamUnary(TestApplication_DoStreamUnaryServer) error
	DoStreamStream(TestApplication_DoStreamStreamServer) error
	DoUnaryUnaryError(context.Context, *Message) (*Message, error)
	DoUnaryStreamError(*Message, TestApplication_DoUnaryStreamErrorServer) error
}

func RegisterTestApplicationServer(s *grpc.Server, srv TestApplicationServer) {
	s.RegisterService(&_TestApplication_serviceDesc, srv)
}

func _TestApplication_DoUnaryUnary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApplicationServer).DoUnaryUnary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TestApplication/DoUnaryUnary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApplicationServer).DoUnaryUnary(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestApplication_DoUnaryStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Message)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestApplicationServer).DoUnaryStream(m, &testApplicationDoUnaryStreamServer{stream})
}

type TestApplication_DoUnaryStreamServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type testApplicationDoUnaryStreamServer struct {
	grpc.ServerStream
}

func (x *testApplicationDoUnaryStreamServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _TestApplication_DoStreamUnary_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestApplicationServer).DoStreamUnary(&testApplicationDoStreamUnaryServer{stream})
}

type TestApplication_DoStreamUnaryServer interface {
	SendAndClose(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type testApplicationDoStreamUnaryServer struct {
	grpc.ServerStream
}

func (x *testApplicationDoStreamUnaryServer) SendAndClose(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testApplicationDoStreamUnaryServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestApplication_DoStreamStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestApplicationServer).DoStreamStream(&testApplicationDoStreamStreamServer{stream})
}

type TestApplication_DoStreamStreamServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type testApplicationDoStreamStreamServer struct {
	grpc.ServerStream
}

func (x *testApplicationDoStreamStreamServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testApplicationDoStreamStreamServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestApplication_DoUnaryUnaryError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApplicationServer).DoUnaryUnaryError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TestApplication/DoUnaryUnaryError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApplicationServer).DoUnaryUnaryError(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestApplication_DoUnaryStreamError_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Message)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestApplicationServer).DoUnaryStreamError(m, &testApplicationDoUnaryStreamErrorServer{stream})
}

type TestApplication_DoUnaryStreamErrorServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type testApplicationDoUnaryStreamErrorServer struct {
	grpc.ServerStream
}

func (x *testApplicationDoUnaryStreamErrorServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

var _TestApplication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "TestApplication",
	HandlerType: (*TestApplicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoUnaryUnary",
			Handler:    _TestApplication_DoUnaryUnary_Handler,
		},
		{
			MethodName: "DoUnaryUnaryError",
			Handler:    _TestApplication_DoUnaryUnaryError_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DoUnaryStream",
			Handler:       _TestApplication_DoUnaryStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "DoStreamUnary",
			Handler:       _TestApplication_DoStreamUnary_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DoStreamStream",
			Handler:       _TestApplication_DoStreamStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DoUnaryStreamError",
			Handler:       _TestApplication_DoUnaryStreamError_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "testapp.proto",
}
