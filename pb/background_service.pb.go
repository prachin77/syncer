// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0--rc1
// source: background_service.proto

package pb

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

type InitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkspacePath     string `protobuf:"bytes,1,opt,name=workspace_path,json=workspacePath,proto3" json:"workspace_path,omitempty"`
	WorkspaceName     string `protobuf:"bytes,2,opt,name=workspace_name,json=workspaceName,proto3" json:"workspace_name,omitempty"`
	WorkspacePassword string `protobuf:"bytes,4,opt,name=workspace_password,json=workspacePassword,proto3" json:"workspace_password,omitempty"`
	Port              string `protobuf:"bytes,5,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *InitRequest) Reset() {
	*x = InitRequest{}
	mi := &file_background_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitRequest) ProtoMessage() {}

func (x *InitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_background_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitRequest.ProtoReflect.Descriptor instead.
func (*InitRequest) Descriptor() ([]byte, []int) {
	return file_background_service_proto_rawDescGZIP(), []int{0}
}

func (x *InitRequest) GetWorkspacePath() string {
	if x != nil {
		return x.WorkspacePath
	}
	return ""
}

func (x *InitRequest) GetWorkspaceName() string {
	if x != nil {
		return x.WorkspaceName
	}
	return ""
}

func (x *InitRequest) GetWorkspacePassword() string {
	if x != nil {
		return x.WorkspacePassword
	}
	return ""
}

func (x *InitRequest) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

type InitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsDirInitialized bool  `protobuf:"varint,1,opt,name=isDirInitialized,proto3" json:"isDirInitialized,omitempty"`
	Port             int32 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *InitResponse) Reset() {
	*x = InitResponse{}
	mi := &file_background_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitResponse) ProtoMessage() {}

func (x *InitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_background_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitResponse.ProtoReflect.Descriptor instead.
func (*InitResponse) Descriptor() ([]byte, []int) {
	return file_background_service_proto_rawDescGZIP(), []int{1}
}

func (x *InitResponse) GetIsDirInitialized() bool {
	if x != nil {
		return x.IsDirInitialized
	}
	return false
}

func (x *InitResponse) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

var File_background_service_proto protoreflect.FileDescriptor

var file_background_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x01, 0x0a, 0x0b, 0x49,
	0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x25, 0x0a, 0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x4e, 0x0a, 0x0c, 0x49,
	0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x69,
	0x73, 0x44, 0x69, 0x72, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x73, 0x44, 0x69, 0x72, 0x49, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x32, 0x41, 0x0a, 0x11, 0x42,
	0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x2c, 0x0a, 0x0d, 0x49, 0x6e, 0x69, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x0c, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0d, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x05,
	0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_background_service_proto_rawDescOnce sync.Once
	file_background_service_proto_rawDescData = file_background_service_proto_rawDesc
)

func file_background_service_proto_rawDescGZIP() []byte {
	file_background_service_proto_rawDescOnce.Do(func() {
		file_background_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_background_service_proto_rawDescData)
	})
	return file_background_service_proto_rawDescData
}

var file_background_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_background_service_proto_goTypes = []any{
	(*InitRequest)(nil),  // 0: InitRequest
	(*InitResponse)(nil), // 1: InitResponse
}
var file_background_service_proto_depIdxs = []int32{
	0, // 0: BackgroundService.InitWorkspace:input_type -> InitRequest
	1, // 1: BackgroundService.InitWorkspace:output_type -> InitResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_background_service_proto_init() }
func file_background_service_proto_init() {
	if File_background_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_background_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_background_service_proto_goTypes,
		DependencyIndexes: file_background_service_proto_depIdxs,
		MessageInfos:      file_background_service_proto_msgTypes,
	}.Build()
	File_background_service_proto = out.File
	file_background_service_proto_rawDesc = nil
	file_background_service_proto_goTypes = nil
	file_background_service_proto_depIdxs = nil
}
