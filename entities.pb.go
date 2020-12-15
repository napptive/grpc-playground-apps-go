// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: playground-apps/entities.proto

package grpc_playground_apps_go

import (
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

// Application top level entity.
type Application struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// AppId with the application identifier.
	AppId string `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	// Name of the app.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// VisualId contains a hash unique to the instance for representation purposes.
	VisualId string `protobuf:"bytes,3,opt,name=visual_id,json=visualId,proto3" json:"visual_id,omitempty"`
	// StatusName with the name of the overall status of the application.
	StatusName string `protobuf:"bytes,4,opt,name=status_name,json=statusName,proto3" json:"status_name,omitempty"`
	// Description of the application.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Application) Reset() {
	*x = Application{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playground_apps_entities_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application) ProtoMessage() {}

func (x *Application) ProtoReflect() protoreflect.Message {
	mi := &file_playground_apps_entities_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Application.ProtoReflect.Descriptor instead.
func (*Application) Descriptor() ([]byte, []int) {
	return file_playground_apps_entities_proto_rawDescGZIP(), []int{0}
}

func (x *Application) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *Application) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Application) GetVisualId() string {
	if x != nil {
		return x.VisualId
	}
	return ""
}

func (x *Application) GetStatusName() string {
	if x != nil {
		return x.StatusName
	}
	return ""
}

func (x *Application) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// AppListResponse with a subset of applications.
type AppListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Entries in the returned page.
	Entries []*Application `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	// From indicates the index of the first entry returned.
	From int32 `protobuf:"varint,2,opt,name=from,proto3" json:"from,omitempty"`
	// To indicates the index of the second entry returned.
	To int32 `protobuf:"varint,3,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *AppListResponse) Reset() {
	*x = AppListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playground_apps_entities_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppListResponse) ProtoMessage() {}

func (x *AppListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_playground_apps_entities_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppListResponse.ProtoReflect.Descriptor instead.
func (*AppListResponse) Descriptor() ([]byte, []int) {
	return file_playground_apps_entities_proto_rawDescGZIP(), []int{1}
}

func (x *AppListResponse) GetEntries() []*Application {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *AppListResponse) GetFrom() int32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *AppListResponse) GetTo() int32 {
	if x != nil {
		return x.To
	}
	return 0
}

var File_playground_apps_entities_proto protoreflect.FileDescriptor

var file_playground_apps_entities_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d, 0x61, 0x70, 0x70,
	0x73, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x70,
	0x73, 0x22, 0x98, 0x01, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6d, 0x0a, 0x0f,
	0x41, 0x70, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x36, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x61, 0x70,
	0x70, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07,
	0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x74, 0x6f, 0x42, 0x45, 0x5a, 0x43, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x61, 0x70, 0x70, 0x74, 0x69,
	0x76, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x2d, 0x61, 0x70, 0x70, 0x73, 0x2d, 0x67, 0x6f, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x5f,
	0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_playground_apps_entities_proto_rawDescOnce sync.Once
	file_playground_apps_entities_proto_rawDescData = file_playground_apps_entities_proto_rawDesc
)

func file_playground_apps_entities_proto_rawDescGZIP() []byte {
	file_playground_apps_entities_proto_rawDescOnce.Do(func() {
		file_playground_apps_entities_proto_rawDescData = protoimpl.X.CompressGZIP(file_playground_apps_entities_proto_rawDescData)
	})
	return file_playground_apps_entities_proto_rawDescData
}

var file_playground_apps_entities_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_playground_apps_entities_proto_goTypes = []interface{}{
	(*Application)(nil),     // 0: playground_apps.Application
	(*AppListResponse)(nil), // 1: playground_apps.AppListResponse
}
var file_playground_apps_entities_proto_depIdxs = []int32{
	0, // 0: playground_apps.AppListResponse.entries:type_name -> playground_apps.Application
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_playground_apps_entities_proto_init() }
func file_playground_apps_entities_proto_init() {
	if File_playground_apps_entities_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_playground_apps_entities_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Application); i {
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
		file_playground_apps_entities_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppListResponse); i {
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
			RawDescriptor: file_playground_apps_entities_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_playground_apps_entities_proto_goTypes,
		DependencyIndexes: file_playground_apps_entities_proto_depIdxs,
		MessageInfos:      file_playground_apps_entities_proto_msgTypes,
	}.Build()
	File_playground_apps_entities_proto = out.File
	file_playground_apps_entities_proto_rawDesc = nil
	file_playground_apps_entities_proto_goTypes = nil
	file_playground_apps_entities_proto_depIdxs = nil
}
