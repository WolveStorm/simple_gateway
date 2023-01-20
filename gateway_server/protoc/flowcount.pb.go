// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: flowcount.proto

package protoc

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

type FlowCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName string `protobuf:"bytes,1,opt,name=ServiceName,proto3" json:"ServiceName,omitempty"`
}

func (x *FlowCountRequest) Reset() {
	*x = FlowCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flowcount_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlowCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowCountRequest) ProtoMessage() {}

func (x *FlowCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flowcount_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowCountRequest.ProtoReflect.Descriptor instead.
func (*FlowCountRequest) Descriptor() ([]byte, []int) {
	return file_flowcount_proto_rawDescGZIP(), []int{0}
}

func (x *FlowCountRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

type FlowCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Qpd            int32   `protobuf:"varint,1,opt,name=qpd,proto3" json:"qpd,omitempty"`
	Qps            int32   `protobuf:"varint,2,opt,name=qps,proto3" json:"qps,omitempty"`
	YesterdayCount []int32 `protobuf:"varint,3,rep,packed,name=yesterdayCount,proto3" json:"yesterdayCount,omitempty"`
	TodayCount     []int32 `protobuf:"varint,4,rep,packed,name=todayCount,proto3" json:"todayCount,omitempty"`
}

func (x *FlowCountResponse) Reset() {
	*x = FlowCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flowcount_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlowCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowCountResponse) ProtoMessage() {}

func (x *FlowCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flowcount_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowCountResponse.ProtoReflect.Descriptor instead.
func (*FlowCountResponse) Descriptor() ([]byte, []int) {
	return file_flowcount_proto_rawDescGZIP(), []int{1}
}

func (x *FlowCountResponse) GetQpd() int32 {
	if x != nil {
		return x.Qpd
	}
	return 0
}

func (x *FlowCountResponse) GetQps() int32 {
	if x != nil {
		return x.Qps
	}
	return 0
}

func (x *FlowCountResponse) GetYesterdayCount() []int32 {
	if x != nil {
		return x.YesterdayCount
	}
	return nil
}

func (x *FlowCountResponse) GetTodayCount() []int32 {
	if x != nil {
		return x.TodayCount
	}
	return nil
}

var File_flowcount_proto protoreflect.FileDescriptor

var file_flowcount_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x66, 0x6c, 0x6f, 0x77, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x34, 0x0a, 0x10, 0x46, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x7f, 0x0a, 0x11, 0x46, 0x6c, 0x6f, 0x77, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x71, 0x70, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x71, 0x70, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x71, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x71, 0x70, 0x73,
	0x12, 0x26, 0x0a, 0x0e, 0x79, 0x65, 0x73, 0x74, 0x65, 0x72, 0x64, 0x61, 0x79, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0e, 0x79, 0x65, 0x73, 0x74, 0x65, 0x72,
	0x64, 0x61, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x64, 0x61,
	0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f,
	0x64, 0x61, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x84, 0x01, 0x0a, 0x09, 0x46, 0x6c, 0x6f,
	0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x46, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x11, 0x2e,
	0x46, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x46,
	0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x11, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x46, 0x6c,
	0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flowcount_proto_rawDescOnce sync.Once
	file_flowcount_proto_rawDescData = file_flowcount_proto_rawDesc
)

func file_flowcount_proto_rawDescGZIP() []byte {
	file_flowcount_proto_rawDescOnce.Do(func() {
		file_flowcount_proto_rawDescData = protoimpl.X.CompressGZIP(file_flowcount_proto_rawDescData)
	})
	return file_flowcount_proto_rawDescData
}

var file_flowcount_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_flowcount_proto_goTypes = []interface{}{
	(*FlowCountRequest)(nil),  // 0: FlowCountRequest
	(*FlowCountResponse)(nil), // 1: FlowCountResponse
}
var file_flowcount_proto_depIdxs = []int32{
	0, // 0: FlowCount.GetServiceFlowCount:input_type -> FlowCountRequest
	0, // 1: FlowCount.GetUserFlowCount:input_type -> FlowCountRequest
	1, // 2: FlowCount.GetServiceFlowCount:output_type -> FlowCountResponse
	1, // 3: FlowCount.GetUserFlowCount:output_type -> FlowCountResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_flowcount_proto_init() }
func file_flowcount_proto_init() {
	if File_flowcount_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flowcount_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlowCountRequest); i {
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
		file_flowcount_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlowCountResponse); i {
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
			RawDescriptor: file_flowcount_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_flowcount_proto_goTypes,
		DependencyIndexes: file_flowcount_proto_depIdxs,
		MessageInfos:      file_flowcount_proto_msgTypes,
	}.Build()
	File_flowcount_proto = out.File
	file_flowcount_proto_rawDesc = nil
	file_flowcount_proto_goTypes = nil
	file_flowcount_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FlowCountClient is the client API for FlowCount service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FlowCountClient interface {
	GetServiceFlowCount(ctx context.Context, in *FlowCountRequest, opts ...grpc.CallOption) (*FlowCountResponse, error)
	GetUserFlowCount(ctx context.Context, in *FlowCountRequest, opts ...grpc.CallOption) (*FlowCountResponse, error)
}

type flowCountClient struct {
	cc grpc.ClientConnInterface
}

func NewFlowCountClient(cc grpc.ClientConnInterface) FlowCountClient {
	return &flowCountClient{cc}
}

func (c *flowCountClient) GetServiceFlowCount(ctx context.Context, in *FlowCountRequest, opts ...grpc.CallOption) (*FlowCountResponse, error) {
	out := new(FlowCountResponse)
	err := c.cc.Invoke(ctx, "/FlowCount/GetServiceFlowCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowCountClient) GetUserFlowCount(ctx context.Context, in *FlowCountRequest, opts ...grpc.CallOption) (*FlowCountResponse, error) {
	out := new(FlowCountResponse)
	err := c.cc.Invoke(ctx, "/FlowCount/GetUserFlowCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FlowCountServer is the server API for FlowCount service.
type FlowCountServer interface {
	GetServiceFlowCount(context.Context, *FlowCountRequest) (*FlowCountResponse, error)
	GetUserFlowCount(context.Context, *FlowCountRequest) (*FlowCountResponse, error)
}

// UnimplementedFlowCountServer can be embedded to have forward compatible implementations.
type UnimplementedFlowCountServer struct {
}

func (*UnimplementedFlowCountServer) GetServiceFlowCount(context.Context, *FlowCountRequest) (*FlowCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServiceFlowCount not implemented")
}
func (*UnimplementedFlowCountServer) GetUserFlowCount(context.Context, *FlowCountRequest) (*FlowCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFlowCount not implemented")
}

func RegisterFlowCountServer(s *grpc.Server, srv FlowCountServer) {
	s.RegisterService(&_FlowCount_serviceDesc, srv)
}

func _FlowCount_GetServiceFlowCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlowCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowCountServer).GetServiceFlowCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FlowCount/GetServiceFlowCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowCountServer).GetServiceFlowCount(ctx, req.(*FlowCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowCount_GetUserFlowCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlowCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowCountServer).GetUserFlowCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FlowCount/GetUserFlowCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowCountServer).GetUserFlowCount(ctx, req.(*FlowCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FlowCount_serviceDesc = grpc.ServiceDesc{
	ServiceName: "FlowCount",
	HandlerType: (*FlowCountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServiceFlowCount",
			Handler:    _FlowCount_GetServiceFlowCount_Handler,
		},
		{
			MethodName: "GetUserFlowCount",
			Handler:    _FlowCount_GetUserFlowCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flowcount.proto",
}
