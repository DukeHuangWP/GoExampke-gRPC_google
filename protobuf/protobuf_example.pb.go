// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: protobuf_example.proto

package protobufExample

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RequestData) Reset() {
	*x = RequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestData) ProtoMessage() {}

func (x *RequestData) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestData.ProtoReflect.Descriptor instead.
func (*RequestData) Descriptor() ([]byte, []int) {
	return file_protobuf_example_proto_rawDescGZIP(), []int{0}
}

func (x *RequestData) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type ResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ResponseData) Reset() {
	*x = ResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseData) ProtoMessage() {}

func (x *ResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseData.ProtoReflect.Descriptor instead.
func (*ResponseData) Descriptor() ([]byte, []int) {
	return file_protobuf_example_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseData) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_protobuf_example_proto protoreflect.FileDescriptor

var file_protobuf_example_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x45,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0x21, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xb6, 0x02,
	0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x12, 0x45, 0x0a, 0x0a, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x12, 0x19,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x45,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00,
	0x30, 0x01, 0x12, 0x49, 0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x28, 0x01, 0x12, 0x47, 0x0a,
	0x09, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_protobuf_example_proto_rawDescOnce sync.Once
	file_protobuf_example_proto_rawDescData = file_protobuf_example_proto_rawDesc
)

func file_protobuf_example_proto_rawDescGZIP() []byte {
	file_protobuf_example_proto_rawDescOnce.Do(func() {
		file_protobuf_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_example_proto_rawDescData)
	})
	return file_protobuf_example_proto_rawDescData
}

var file_protobuf_example_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protobuf_example_proto_goTypes = []interface{}{
	(*RequestData)(nil),  // 0: protoExample.RequestData
	(*ResponseData)(nil), // 1: protoExample.ResponseData
}
var file_protobuf_example_proto_depIdxs = []int32{
	0, // 0: protoExample.serviceExample.UnaryTrans:input_type -> protoExample.RequestData
	0, // 1: protoExample.serviceExample.ServerStream:input_type -> protoExample.RequestData
	0, // 2: protoExample.serviceExample.ClientStream:input_type -> protoExample.RequestData
	0, // 3: protoExample.serviceExample.AllStream:input_type -> protoExample.RequestData
	1, // 4: protoExample.serviceExample.UnaryTrans:output_type -> protoExample.ResponseData
	1, // 5: protoExample.serviceExample.ServerStream:output_type -> protoExample.ResponseData
	1, // 6: protoExample.serviceExample.ClientStream:output_type -> protoExample.ResponseData
	0, // 7: protoExample.serviceExample.AllStream:output_type -> protoExample.RequestData
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_example_proto_init() }
func file_protobuf_example_proto_init() {
	if File_protobuf_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobuf_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_example_proto_goTypes,
		DependencyIndexes: file_protobuf_example_proto_depIdxs,
		MessageInfos:      file_protobuf_example_proto_msgTypes,
	}.Build()
	File_protobuf_example_proto = out.File
	file_protobuf_example_proto_rawDesc = nil
	file_protobuf_example_proto_goTypes = nil
	file_protobuf_example_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServiceExampleClient is the client API for ServiceExample service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceExampleClient interface {
	UnaryTrans(ctx context.Context, in *RequestData, opts ...grpc.CallOption) (*ResponseData, error)
	ServerStream(ctx context.Context, in *RequestData, opts ...grpc.CallOption) (ServiceExample_ServerStreamClient, error)
	ClientStream(ctx context.Context, opts ...grpc.CallOption) (ServiceExample_ClientStreamClient, error)
	AllStream(ctx context.Context, opts ...grpc.CallOption) (ServiceExample_AllStreamClient, error)
}

type serviceExampleClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceExampleClient(cc grpc.ClientConnInterface) ServiceExampleClient {
	return &serviceExampleClient{cc}
}

func (c *serviceExampleClient) UnaryTrans(ctx context.Context, in *RequestData, opts ...grpc.CallOption) (*ResponseData, error) {
	out := new(ResponseData)
	err := c.cc.Invoke(ctx, "/protoExample.serviceExample/UnaryTrans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceExampleClient) ServerStream(ctx context.Context, in *RequestData, opts ...grpc.CallOption) (ServiceExample_ServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ServiceExample_serviceDesc.Streams[0], "/protoExample.serviceExample/ServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceExampleServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServiceExample_ServerStreamClient interface {
	Recv() (*ResponseData, error)
	grpc.ClientStream
}

type serviceExampleServerStreamClient struct {
	grpc.ClientStream
}

func (x *serviceExampleServerStreamClient) Recv() (*ResponseData, error) {
	m := new(ResponseData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceExampleClient) ClientStream(ctx context.Context, opts ...grpc.CallOption) (ServiceExample_ClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ServiceExample_serviceDesc.Streams[1], "/protoExample.serviceExample/ClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceExampleClientStreamClient{stream}
	return x, nil
}

type ServiceExample_ClientStreamClient interface {
	Send(*RequestData) error
	CloseAndRecv() (*ResponseData, error)
	grpc.ClientStream
}

type serviceExampleClientStreamClient struct {
	grpc.ClientStream
}

func (x *serviceExampleClientStreamClient) Send(m *RequestData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serviceExampleClientStreamClient) CloseAndRecv() (*ResponseData, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ResponseData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceExampleClient) AllStream(ctx context.Context, opts ...grpc.CallOption) (ServiceExample_AllStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ServiceExample_serviceDesc.Streams[2], "/protoExample.serviceExample/AllStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceExampleAllStreamClient{stream}
	return x, nil
}

type ServiceExample_AllStreamClient interface {
	Send(*RequestData) error
	Recv() (*RequestData, error)
	grpc.ClientStream
}

type serviceExampleAllStreamClient struct {
	grpc.ClientStream
}

func (x *serviceExampleAllStreamClient) Send(m *RequestData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serviceExampleAllStreamClient) Recv() (*RequestData, error) {
	m := new(RequestData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServiceExampleServer is the server API for ServiceExample service.
type ServiceExampleServer interface {
	UnaryTrans(context.Context, *RequestData) (*ResponseData, error)
	ServerStream(*RequestData, ServiceExample_ServerStreamServer) error
	ClientStream(ServiceExample_ClientStreamServer) error
	AllStream(ServiceExample_AllStreamServer) error
}

// UnimplementedServiceExampleServer can be embedded to have forward compatible implementations.
type UnimplementedServiceExampleServer struct {
}

func (*UnimplementedServiceExampleServer) UnaryTrans(context.Context, *RequestData) (*ResponseData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnaryTrans not implemented")
}
func (*UnimplementedServiceExampleServer) ServerStream(*RequestData, ServiceExample_ServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStream not implemented")
}
func (*UnimplementedServiceExampleServer) ClientStream(ServiceExample_ClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStream not implemented")
}
func (*UnimplementedServiceExampleServer) AllStream(ServiceExample_AllStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method AllStream not implemented")
}

func RegisterServiceExampleServer(s *grpc.Server, srv ServiceExampleServer) {
	s.RegisterService(&_ServiceExample_serviceDesc, srv)
}

func _ServiceExample_UnaryTrans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceExampleServer).UnaryTrans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoExample.serviceExample/UnaryTrans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceExampleServer).UnaryTrans(ctx, req.(*RequestData))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceExample_ServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestData)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceExampleServer).ServerStream(m, &serviceExampleServerStreamServer{stream})
}

type ServiceExample_ServerStreamServer interface {
	Send(*ResponseData) error
	grpc.ServerStream
}

type serviceExampleServerStreamServer struct {
	grpc.ServerStream
}

func (x *serviceExampleServerStreamServer) Send(m *ResponseData) error {
	return x.ServerStream.SendMsg(m)
}

func _ServiceExample_ClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServiceExampleServer).ClientStream(&serviceExampleClientStreamServer{stream})
}

type ServiceExample_ClientStreamServer interface {
	SendAndClose(*ResponseData) error
	Recv() (*RequestData, error)
	grpc.ServerStream
}

type serviceExampleClientStreamServer struct {
	grpc.ServerStream
}

func (x *serviceExampleClientStreamServer) SendAndClose(m *ResponseData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serviceExampleClientStreamServer) Recv() (*RequestData, error) {
	m := new(RequestData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ServiceExample_AllStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServiceExampleServer).AllStream(&serviceExampleAllStreamServer{stream})
}

type ServiceExample_AllStreamServer interface {
	Send(*RequestData) error
	Recv() (*RequestData, error)
	grpc.ServerStream
}

type serviceExampleAllStreamServer struct {
	grpc.ServerStream
}

func (x *serviceExampleAllStreamServer) Send(m *RequestData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serviceExampleAllStreamServer) Recv() (*RequestData, error) {
	m := new(RequestData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ServiceExample_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protoExample.serviceExample",
	HandlerType: (*ServiceExampleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryTrans",
			Handler:    _ServiceExample_UnaryTrans_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStream",
			Handler:       _ServiceExample_ServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientStream",
			Handler:       _ServiceExample_ClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "AllStream",
			Handler:       _ServiceExample_AllStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protobuf_example.proto",
}
