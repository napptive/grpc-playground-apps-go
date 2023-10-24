// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.15.2
// source: playground-apps/services.proto

package grpc_playground_apps_go

import (
	context "context"
	grpc_playground_common_go "github.com/napptive/grpc-playground-common-go"
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
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x8e, 0x09, 0x0a, 0x04, 0x41, 0x70,
	0x70, 0x73, 0x12, 0x7b, 0x0a, 0x06, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x12, 0x29, 0x2e, 0x70,
	0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76, 0x30,
	0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x3a, 0x01, 0x2a, 0x12,
	0x6e, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x29, 0x2e, 0x70, 0x6c, 0x61, 0x79,
	0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76, 0x30,
	0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x3a, 0x01, 0x2a, 0x12,
	0x86, 0x01, 0x0a, 0x0b, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x2c, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x1a, 0x27, 0x2e,
	0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x2e,
	0x41, 0x70, 0x70, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15,
	0x2f, 0x76, 0x30, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x2f, 0x6c, 0x69, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x5e, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1f, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x61, 0x70,
	0x70, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x61,
	0x70, 0x70, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x76, 0x30, 0x2f, 0x61, 0x70, 0x70,
	0x73, 0x2f, 0x67, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x6e, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75,
	0x6d, 0x65, 0x12, 0x29, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f,
	0x61, 0x70, 0x70, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x4f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76, 0x30, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72,
	0x65, 0x73, 0x75, 0x6d, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x5f, 0x0a, 0x04, 0x4c, 0x6f, 0x67, 0x73,
	0x12, 0x1b, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x61, 0x70,
	0x70, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x2e,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x73, 0x22,
	0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x76, 0x30, 0x2f, 0x61, 0x70, 0x70,
	0x73, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x60, 0x0a, 0x05, 0x53, 0x63, 0x61,
	0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f,
	0x61, 0x70, 0x70, 0x73, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x76, 0x30, 0x2f, 0x61, 0x70,
	0x70, 0x73, 0x2f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x87, 0x01, 0x0a, 0x19,
	0x53, 0x74, 0x6f, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x31, 0x2e, 0x70, 0x6c, 0x61, 0x79,
	0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x70,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70,
	0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x4f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x76, 0x30, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x74,
	0x6f, 0x70, 0x3a, 0x01, 0x2a, 0x12, 0x82, 0x01, 0x0a, 0x12, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x2a, 0x2e, 0x70,
	0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22,
	0x16, 0x2f, 0x76, 0x30, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x6e, 0x0a, 0x06, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x29, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76, 0x30, 0x2f, 0x61, 0x70, 0x70, 0x73,
	0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x61, 0x70, 0x70, 0x74, 0x69, 0x76,
	0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x2d, 0x61, 0x70, 0x70, 0x73, 0x2d, 0x67, 0x6f, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70,
	0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x5f, 0x67,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_playground_apps_services_proto_goTypes = []interface{}{
	(*DeployApplicationRequest)(nil),                            // 0: playground_apps.DeployApplicationRequest
	(*RemoveApplicationRequest)(nil),                            // 1: playground_apps.RemoveApplicationRequest
	(*grpc_playground_environments_go.EnvironmentSelector)(nil), // 2: playground_environments.EnvironmentSelector
	(*AppInfoRequest)(nil),                                      // 3: playground_apps.AppInfoRequest
	(*ResumeApplicationRequest)(nil),                            // 4: playground_apps.ResumeApplicationRequest
	(*LogRequest)(nil),                                          // 5: playground_apps.LogRequest
	(*ScaleRequest)(nil),                                        // 6: playground_apps.ScaleRequest
	(*StopApplicationComponentsRequest)(nil),                    // 7: playground_apps.StopApplicationComponentsRequest
	(*ValidateRepoAccessRequest)(nil),                           // 8: playground_apps.ValidateRepoAccessRequest
	(*UpdateApplicationRequest)(nil),                            // 9: playground_apps.UpdateApplicationRequest
	(*DeployApplicationResponse)(nil),                           // 10: playground_apps.DeployApplicationResponse
	(*grpc_playground_common_go.OpResponse)(nil),                // 11: playground_common.OpResponse
	(*AppSummaryListResponse)(nil),                              // 12: playground_apps.AppSummaryListResponse
	(*Application)(nil),                                         // 13: playground_apps.Application
	(*ApplicationLogs)(nil),                                     // 14: playground_apps.ApplicationLogs
}
var file_playground_apps_services_proto_depIdxs = []int32{
	0,  // 0: playground_apps.Apps.Deploy:input_type -> playground_apps.DeployApplicationRequest
	1,  // 1: playground_apps.Apps.Remove:input_type -> playground_apps.RemoveApplicationRequest
	2,  // 2: playground_apps.Apps.SummaryList:input_type -> playground_environments.EnvironmentSelector
	3,  // 3: playground_apps.Apps.Info:input_type -> playground_apps.AppInfoRequest
	4,  // 4: playground_apps.Apps.Resume:input_type -> playground_apps.ResumeApplicationRequest
	5,  // 5: playground_apps.Apps.Logs:input_type -> playground_apps.LogRequest
	6,  // 6: playground_apps.Apps.Scale:input_type -> playground_apps.ScaleRequest
	7,  // 7: playground_apps.Apps.StopApplicationComponents:input_type -> playground_apps.StopApplicationComponentsRequest
	8,  // 8: playground_apps.Apps.ValidateRepoAccess:input_type -> playground_apps.ValidateRepoAccessRequest
	9,  // 9: playground_apps.Apps.Update:input_type -> playground_apps.UpdateApplicationRequest
	10, // 10: playground_apps.Apps.Deploy:output_type -> playground_apps.DeployApplicationResponse
	11, // 11: playground_apps.Apps.Remove:output_type -> playground_common.OpResponse
	12, // 12: playground_apps.Apps.SummaryList:output_type -> playground_apps.AppSummaryListResponse
	13, // 13: playground_apps.Apps.Info:output_type -> playground_apps.Application
	11, // 14: playground_apps.Apps.Resume:output_type -> playground_common.OpResponse
	14, // 15: playground_apps.Apps.Logs:output_type -> playground_apps.ApplicationLogs
	11, // 16: playground_apps.Apps.Scale:output_type -> playground_common.OpResponse
	11, // 17: playground_apps.Apps.StopApplicationComponents:output_type -> playground_common.OpResponse
	11, // 18: playground_apps.Apps.ValidateRepoAccess:output_type -> playground_common.OpResponse
	11, // 19: playground_apps.Apps.Update:output_type -> playground_common.OpResponse
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
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
	// Deploy an application.
	Deploy(ctx context.Context, in *DeployApplicationRequest, opts ...grpc.CallOption) (*DeployApplicationResponse, error)
	// Remove an application instance.
	Remove(ctx context.Context, in *RemoveApplicationRequest, opts ...grpc.CallOption) (*grpc_playground_common_go.OpResponse, error)
	// SummaryList returns a subset of application summaries found in the user cluster.
	SummaryList(ctx context.Context, in *grpc_playground_environments_go.EnvironmentSelector, opts ...grpc.CallOption) (*AppSummaryListResponse, error)
	// Get returns an application
	Info(ctx context.Context, in *AppInfoRequest, opts ...grpc.CallOption) (*Application, error)
	// Resume a suspended workflow application.
	Resume(ctx context.Context, in *ResumeApplicationRequest, opts ...grpc.CallOption) (*grpc_playground_common_go.OpResponse, error)
	// Logs returns the application logs (filtering by component if required)
	Logs(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*ApplicationLogs, error)
	// Scale scales the application components
	Scale(ctx context.Context, in *ScaleRequest, opts ...grpc.CallOption) (*grpc_playground_common_go.OpResponse, error)
	// StopApplicationComponents stops application components adding the stop trait
	StopApplicationComponents(ctx context.Context, in *StopApplicationComponentsRequest, opts ...grpc.CallOption) (*grpc_playground_common_go.OpResponse, error)
	// ValidateRepoAccess to validate if a repository is accessible with the credentials received
	ValidateRepoAccess(ctx context.Context, in *ValidateRepoAccessRequest, opts ...grpc.CallOption) (*grpc_playground_common_go.OpResponse, error)
	// Update an application.
	Update(ctx context.Context, in *UpdateApplicationRequest, opts ...grpc.CallOption) (*grpc_playground_common_go.OpResponse, error)
}

type appsClient struct {
	cc grpc.ClientConnInterface
}

func NewAppsClient(cc grpc.ClientConnInterface) AppsClient {
	return &appsClient{cc}
}

func (c *appsClient) Deploy(ctx context.Context, in *DeployApplicationRequest, opts ...grpc.CallOption) (*DeployApplicationResponse, error) {
	out := new(DeployApplicationResponse)
	err := c.cc.Invoke(ctx, "/playground_apps.Apps/Deploy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appsClient) Remove(ctx context.Context, in *RemoveApplicationRequest, opts ...grpc.CallOption) (*grpc_playground_common_go.OpResponse, error) {
	out := new(grpc_playground_common_go.OpResponse)
	err := c.cc.Invoke(ctx, "/playground_apps.Apps/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appsClient) SummaryList(ctx context.Context, in *grpc_playground_environments_go.EnvironmentSelector, opts ...grpc.CallOption) (*AppSummaryListResponse, error) {
	out := new(AppSummaryListResponse)
	err := c.cc.Invoke(ctx, "/playground_apps.Apps/SummaryList", in, out, opts...)
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

func (c *appsClient) Resume(ctx context.Context, in *ResumeApplicationRequest, opts ...grpc.CallOption) (*grpc_playground_common_go.OpResponse, error) {
	out := new(grpc_playground_common_go.OpResponse)
	err := c.cc.Invoke(ctx, "/playground_apps.Apps/Resume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appsClient) Logs(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*ApplicationLogs, error) {
	out := new(ApplicationLogs)
	err := c.cc.Invoke(ctx, "/playground_apps.Apps/Logs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appsClient) Scale(ctx context.Context, in *ScaleRequest, opts ...grpc.CallOption) (*grpc_playground_common_go.OpResponse, error) {
	out := new(grpc_playground_common_go.OpResponse)
	err := c.cc.Invoke(ctx, "/playground_apps.Apps/Scale", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appsClient) StopApplicationComponents(ctx context.Context, in *StopApplicationComponentsRequest, opts ...grpc.CallOption) (*grpc_playground_common_go.OpResponse, error) {
	out := new(grpc_playground_common_go.OpResponse)
	err := c.cc.Invoke(ctx, "/playground_apps.Apps/StopApplicationComponents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appsClient) ValidateRepoAccess(ctx context.Context, in *ValidateRepoAccessRequest, opts ...grpc.CallOption) (*grpc_playground_common_go.OpResponse, error) {
	out := new(grpc_playground_common_go.OpResponse)
	err := c.cc.Invoke(ctx, "/playground_apps.Apps/ValidateRepoAccess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appsClient) Update(ctx context.Context, in *UpdateApplicationRequest, opts ...grpc.CallOption) (*grpc_playground_common_go.OpResponse, error) {
	out := new(grpc_playground_common_go.OpResponse)
	err := c.cc.Invoke(ctx, "/playground_apps.Apps/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppsServer is the server API for Apps service.
type AppsServer interface {
	// Deploy an application.
	Deploy(context.Context, *DeployApplicationRequest) (*DeployApplicationResponse, error)
	// Remove an application instance.
	Remove(context.Context, *RemoveApplicationRequest) (*grpc_playground_common_go.OpResponse, error)
	// SummaryList returns a subset of application summaries found in the user cluster.
	SummaryList(context.Context, *grpc_playground_environments_go.EnvironmentSelector) (*AppSummaryListResponse, error)
	// Get returns an application
	Info(context.Context, *AppInfoRequest) (*Application, error)
	// Resume a suspended workflow application.
	Resume(context.Context, *ResumeApplicationRequest) (*grpc_playground_common_go.OpResponse, error)
	// Logs returns the application logs (filtering by component if required)
	Logs(context.Context, *LogRequest) (*ApplicationLogs, error)
	// Scale scales the application components
	Scale(context.Context, *ScaleRequest) (*grpc_playground_common_go.OpResponse, error)
	// StopApplicationComponents stops application components adding the stop trait
	StopApplicationComponents(context.Context, *StopApplicationComponentsRequest) (*grpc_playground_common_go.OpResponse, error)
	// ValidateRepoAccess to validate if a repository is accessible with the credentials received
	ValidateRepoAccess(context.Context, *ValidateRepoAccessRequest) (*grpc_playground_common_go.OpResponse, error)
	// Update an application.
	Update(context.Context, *UpdateApplicationRequest) (*grpc_playground_common_go.OpResponse, error)
}

// UnimplementedAppsServer can be embedded to have forward compatible implementations.
type UnimplementedAppsServer struct {
}

func (*UnimplementedAppsServer) Deploy(context.Context, *DeployApplicationRequest) (*DeployApplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deploy not implemented")
}
func (*UnimplementedAppsServer) Remove(context.Context, *RemoveApplicationRequest) (*grpc_playground_common_go.OpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (*UnimplementedAppsServer) SummaryList(context.Context, *grpc_playground_environments_go.EnvironmentSelector) (*AppSummaryListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SummaryList not implemented")
}
func (*UnimplementedAppsServer) Info(context.Context, *AppInfoRequest) (*Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (*UnimplementedAppsServer) Resume(context.Context, *ResumeApplicationRequest) (*grpc_playground_common_go.OpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Resume not implemented")
}
func (*UnimplementedAppsServer) Logs(context.Context, *LogRequest) (*ApplicationLogs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logs not implemented")
}
func (*UnimplementedAppsServer) Scale(context.Context, *ScaleRequest) (*grpc_playground_common_go.OpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Scale not implemented")
}
func (*UnimplementedAppsServer) StopApplicationComponents(context.Context, *StopApplicationComponentsRequest) (*grpc_playground_common_go.OpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopApplicationComponents not implemented")
}
func (*UnimplementedAppsServer) ValidateRepoAccess(context.Context, *ValidateRepoAccessRequest) (*grpc_playground_common_go.OpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateRepoAccess not implemented")
}
func (*UnimplementedAppsServer) Update(context.Context, *UpdateApplicationRequest) (*grpc_playground_common_go.OpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func RegisterAppsServer(s *grpc.Server, srv AppsServer) {
	s.RegisterService(&_Apps_serviceDesc, srv)
}

func _Apps_Deploy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeployApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppsServer).Deploy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/playground_apps.Apps/Deploy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppsServer).Deploy(ctx, req.(*DeployApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Apps_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppsServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/playground_apps.Apps/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppsServer).Remove(ctx, req.(*RemoveApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Apps_SummaryList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(grpc_playground_environments_go.EnvironmentSelector)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppsServer).SummaryList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/playground_apps.Apps/SummaryList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppsServer).SummaryList(ctx, req.(*grpc_playground_environments_go.EnvironmentSelector))
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

func _Apps_Resume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResumeApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppsServer).Resume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/playground_apps.Apps/Resume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppsServer).Resume(ctx, req.(*ResumeApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Apps_Logs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppsServer).Logs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/playground_apps.Apps/Logs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppsServer).Logs(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Apps_Scale_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScaleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppsServer).Scale(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/playground_apps.Apps/Scale",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppsServer).Scale(ctx, req.(*ScaleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Apps_StopApplicationComponents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopApplicationComponentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppsServer).StopApplicationComponents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/playground_apps.Apps/StopApplicationComponents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppsServer).StopApplicationComponents(ctx, req.(*StopApplicationComponentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Apps_ValidateRepoAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateRepoAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppsServer).ValidateRepoAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/playground_apps.Apps/ValidateRepoAccess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppsServer).ValidateRepoAccess(ctx, req.(*ValidateRepoAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Apps_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/playground_apps.Apps/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppsServer).Update(ctx, req.(*UpdateApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Apps_serviceDesc = grpc.ServiceDesc{
	ServiceName: "playground_apps.Apps",
	HandlerType: (*AppsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Deploy",
			Handler:    _Apps_Deploy_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Apps_Remove_Handler,
		},
		{
			MethodName: "SummaryList",
			Handler:    _Apps_SummaryList_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _Apps_Info_Handler,
		},
		{
			MethodName: "Resume",
			Handler:    _Apps_Resume_Handler,
		},
		{
			MethodName: "Logs",
			Handler:    _Apps_Logs_Handler,
		},
		{
			MethodName: "Scale",
			Handler:    _Apps_Scale_Handler,
		},
		{
			MethodName: "StopApplicationComponents",
			Handler:    _Apps_StopApplicationComponents_Handler,
		},
		{
			MethodName: "ValidateRepoAccess",
			Handler:    _Apps_ValidateRepoAccess_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Apps_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "playground-apps/services.proto",
}
