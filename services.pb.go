// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.15.2
// source: playground-apps/services.proto

package grpc_playground_apps_go

import (
	context "context"
	_ "github.com/napptive/grpc-playground-common-go"
	grpc_playground_environments_go "github.com/napptive/grpc-playground-environments-go"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_playground_apps_services_proto protoreflect.FileDescriptor

var file_playground_apps_services_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d, 0x61, 0x70, 0x70,
	0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x70,
	0x73, 0x1a, 0x1e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d, 0x61, 0x70,
	0x70, 0x73, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd5, 0x01, 0x0a, 0x04, 0x41, 0x70,
	0x70, 0x73, 0x12, 0x6d, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x2e, 0x70, 0x6c, 0x61,
	0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x1a, 0x20, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x30, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6c, 0x69, 0x73,
	0x74, 0x12, 0x5e, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x2e, 0x70, 0x6c, 0x61, 0x79,
	0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x6c, 0x61,
	0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11,
	0x22, 0x0c, 0x2f, 0x76, 0x30, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x3a, 0x01,
	0x2a, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6e, 0x61, 0x70, 0x70, 0x74, 0x69, 0x76, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x6c,
	0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d, 0x61, 0x70, 0x70, 0x73, 0x2d, 0x67, 0x6f,
	0x3b, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x5f, 0x61, 0x70, 0x70, 0x73, 0x5f, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_playground_apps_services_proto_goTypes = []interface{}{
	(*grpc_playground_environments_go.EnvironmentSelector)(nil), // 0: playground_environments.EnvironmentSelector
	(*AppInfoRequest)(nil),  // 1: playground_apps.AppInfoRequest
	(*AppListResponse)(nil), // 2: playground_apps.AppListResponse
	(*Application)(nil),     // 3: playground_apps.Application
}
var file_playground_apps_services_proto_depIdxs = []int32{
	0, // 0: playground_apps.Apps.List:input_type -> playground_environments.EnvironmentSelector
	1, // 1: playground_apps.Apps.Info:input_type -> playground_apps.AppInfoRequest
	2, // 2: playground_apps.Apps.List:output_type -> playground_apps.AppListResponse
	3, // 3: playground_apps.Apps.Info:output_type -> playground_apps.Application
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_playground_apps_services_proto_init() }
func file_playground_apps_services_proto_init() {
	if File_playground_apps_services_proto != nil {
		return
	}
	file_playground_apps_entities_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_playground_apps_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_playground_apps_services_proto_goTypes,
		DependencyIndexes: file_playground_apps_services_proto_depIdxs,
	}.Build()
	File_playground_apps_services_proto = out.File
	file_playground_apps_services_proto_rawDesc = nil
	file_playground_apps_services_proto_goTypes = nil
	file_playground_apps_services_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AppsClient is the client API for Apps service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AppsClient interface {
	// List returns a subset of applications found in the user cluster.
	List(ctx context.Context, in *grpc_playground_environments_go.EnvironmentSelector, opts ...grpc.CallOption) (*AppListResponse, error)
	// Get returns an application
	Info(ctx context.Context, in *AppInfoRequest, opts ...grpc.CallOption) (*Application, error)
}

type appsClient struct {
	cc grpc.ClientConnInterface
}

func NewAppsClient(cc grpc.ClientConnInterface) AppsClient {
	return &appsClient{cc}
}

func (c *appsClient) List(ctx context.Context, in *grpc_playground_environments_go.EnvironmentSelector, opts ...grpc.CallOption) (*AppListResponse, error) {
	out := new(AppListResponse)
	err := c.cc.Invoke(ctx, "/playground_apps.Apps/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appsClient) Info(ctx context.Context, in *AppInfoRequest, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := c.cc.Invoke(ctx, "/playground_apps.Apps/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppsServer is the server API for Apps service.
type AppsServer interface {
	// List returns a subset of applications found in the user cluster.
	List(context.Context, *grpc_playground_environments_go.EnvironmentSelector) (*AppListResponse, error)
	// Get returns an application
	Info(context.Context, *AppInfoRequest) (*Application, error)
}

// UnimplementedAppsServer can be embedded to have forward compatible implementations.
type UnimplementedAppsServer struct {
}

func (*UnimplementedAppsServer) List(context.Context, *grpc_playground_environments_go.EnvironmentSelector) (*AppListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedAppsServer) Info(context.Context, *AppInfoRequest) (*Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}

func RegisterAppsServer(s *grpc.Server, srv AppsServer) {
	s.RegisterService(&_Apps_serviceDesc, srv)
}

func _Apps_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(grpc_playground_environments_go.EnvironmentSelector)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/playground_apps.Apps/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppsServer).List(ctx, req.(*grpc_playground_environments_go.EnvironmentSelector))
	}
	return interceptor(ctx, in, info, handler)
}

func _Apps_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppsServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/playground_apps.Apps/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppsServer).Info(ctx, req.(*AppInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Apps_serviceDesc = grpc.ServiceDesc{
	ServiceName: "playground_apps.Apps",
	HandlerType: (*AppsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Apps_List_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _Apps_Info_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "playground-apps/services.proto",
}
