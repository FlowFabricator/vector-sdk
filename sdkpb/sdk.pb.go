// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: sdkpb/sdk.proto

package sdkpb

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

type RunRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plugin        string `protobuf:"bytes,1,opt,name=plugin,proto3" json:"plugin,omitempty"`
	PluginVersion string `protobuf:"bytes,2,opt,name=pluginVersion,proto3" json:"pluginVersion,omitempty"`
	Action        string `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	ArgsMapAsJson []byte `protobuf:"bytes,4,opt,name=argsMapAsJson,proto3" json:"argsMapAsJson,omitempty"`
}

func (x *RunRequest) Reset() {
	*x = RunRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdkpb_sdk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunRequest) ProtoMessage() {}

func (x *RunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sdkpb_sdk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunRequest.ProtoReflect.Descriptor instead.
func (*RunRequest) Descriptor() ([]byte, []int) {
	return file_sdkpb_sdk_proto_rawDescGZIP(), []int{0}
}

func (x *RunRequest) GetPlugin() string {
	if x != nil {
		return x.Plugin
	}
	return ""
}

func (x *RunRequest) GetPluginVersion() string {
	if x != nil {
		return x.PluginVersion
	}
	return ""
}

func (x *RunRequest) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *RunRequest) GetArgsMapAsJson() []byte {
	if x != nil {
		return x.ArgsMapAsJson
	}
	return nil
}

type RunResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExitCode uint32 `protobuf:"varint,1,opt,name=exitCode,proto3" json:"exitCode,omitempty"`
	Details  string `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
	Data     []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RunResponse) Reset() {
	*x = RunResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdkpb_sdk_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunResponse) ProtoMessage() {}

func (x *RunResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sdkpb_sdk_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunResponse.ProtoReflect.Descriptor instead.
func (*RunResponse) Descriptor() ([]byte, []int) {
	return file_sdkpb_sdk_proto_rawDescGZIP(), []int{1}
}

func (x *RunResponse) GetExitCode() uint32 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

func (x *RunResponse) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

func (x *RunResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type CallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RunRequest  *RunRequest `protobuf:"bytes,1,opt,name=runRequest,proto3" json:"runRequest,omitempty"`
	Roles       []string    `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
	StateName   string      `protobuf:"bytes,3,opt,name=stateName,proto3" json:"stateName,omitempty"`
	ExecutionId string      `protobuf:"bytes,4,opt,name=executionId,proto3" json:"executionId,omitempty"`
}

func (x *CallRequest) Reset() {
	*x = CallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdkpb_sdk_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallRequest) ProtoMessage() {}

func (x *CallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sdkpb_sdk_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallRequest.ProtoReflect.Descriptor instead.
func (*CallRequest) Descriptor() ([]byte, []int) {
	return file_sdkpb_sdk_proto_rawDescGZIP(), []int{2}
}

func (x *CallRequest) GetRunRequest() *RunRequest {
	if x != nil {
		return x.RunRequest
	}
	return nil
}

func (x *CallRequest) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *CallRequest) GetStateName() string {
	if x != nil {
		return x.StateName
	}
	return ""
}

func (x *CallRequest) GetExecutionId() string {
	if x != nil {
		return x.ExecutionId
	}
	return ""
}

type CallResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExitCode uint32 `protobuf:"varint,1,opt,name=exitCode,proto3" json:"exitCode,omitempty"`
	Details  string `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
	Data     []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CallResponse) Reset() {
	*x = CallResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdkpb_sdk_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallResponse) ProtoMessage() {}

func (x *CallResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sdkpb_sdk_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallResponse.ProtoReflect.Descriptor instead.
func (*CallResponse) Descriptor() ([]byte, []int) {
	return file_sdkpb_sdk_proto_rawDescGZIP(), []int{3}
}

func (x *CallResponse) GetExitCode() uint32 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

func (x *CallResponse) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

func (x *CallResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_sdkpb_sdk_proto protoreflect.FileDescriptor

var file_sdkpb_sdk_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x64, 0x6b, 0x70, 0x62, 0x2f, 0x73, 0x64, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x73, 0x64, 0x6b, 0x70, 0x62, 0x22, 0x88, 0x01, 0x0a, 0x0a, 0x52, 0x75, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12,
	0x24, 0x0a, 0x0d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a,
	0x0d, 0x61, 0x72, 0x67, 0x73, 0x4d, 0x61, 0x70, 0x41, 0x73, 0x4a, 0x73, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x61, 0x72, 0x67, 0x73, 0x4d, 0x61, 0x70, 0x41, 0x73, 0x4a,
	0x73, 0x6f, 0x6e, 0x22, 0x57, 0x0a, 0x0b, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x96, 0x01, 0x0a,
	0x0b, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x0a,
	0x72, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x73, 0x64, 0x6b, 0x70, 0x62, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x0a, 0x72, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x58, 0x0a, 0x0c, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32,
	0x64, 0x0a, 0x03, 0x53, 0x44, 0x4b, 0x12, 0x2f, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x12,
	0x2e, 0x73, 0x64, 0x6b, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x73, 0x64, 0x6b, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12, 0x11,
	0x2e, 0x73, 0x64, 0x6b, 0x70, 0x62, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x73, 0x64, 0x6b, 0x70, 0x62, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x73, 0x64, 0x6b, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sdkpb_sdk_proto_rawDescOnce sync.Once
	file_sdkpb_sdk_proto_rawDescData = file_sdkpb_sdk_proto_rawDesc
)

func file_sdkpb_sdk_proto_rawDescGZIP() []byte {
	file_sdkpb_sdk_proto_rawDescOnce.Do(func() {
		file_sdkpb_sdk_proto_rawDescData = protoimpl.X.CompressGZIP(file_sdkpb_sdk_proto_rawDescData)
	})
	return file_sdkpb_sdk_proto_rawDescData
}

var file_sdkpb_sdk_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_sdkpb_sdk_proto_goTypes = []interface{}{
	(*RunRequest)(nil),   // 0: sdkpb.RunRequest
	(*RunResponse)(nil),  // 1: sdkpb.RunResponse
	(*CallRequest)(nil),  // 2: sdkpb.CallRequest
	(*CallResponse)(nil), // 3: sdkpb.CallResponse
}
var file_sdkpb_sdk_proto_depIdxs = []int32{
	0, // 0: sdkpb.CallRequest.runRequest:type_name -> sdkpb.RunRequest
	2, // 1: sdkpb.SDK.Call:input_type -> sdkpb.CallRequest
	0, // 2: sdkpb.SDK.Run:input_type -> sdkpb.RunRequest
	3, // 3: sdkpb.SDK.Call:output_type -> sdkpb.CallResponse
	1, // 4: sdkpb.SDK.Run:output_type -> sdkpb.RunResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sdkpb_sdk_proto_init() }
func file_sdkpb_sdk_proto_init() {
	if File_sdkpb_sdk_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sdkpb_sdk_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunRequest); i {
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
		file_sdkpb_sdk_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunResponse); i {
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
		file_sdkpb_sdk_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallRequest); i {
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
		file_sdkpb_sdk_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallResponse); i {
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
			RawDescriptor: file_sdkpb_sdk_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sdkpb_sdk_proto_goTypes,
		DependencyIndexes: file_sdkpb_sdk_proto_depIdxs,
		MessageInfos:      file_sdkpb_sdk_proto_msgTypes,
	}.Build()
	File_sdkpb_sdk_proto = out.File
	file_sdkpb_sdk_proto_rawDesc = nil
	file_sdkpb_sdk_proto_goTypes = nil
	file_sdkpb_sdk_proto_depIdxs = nil
}
