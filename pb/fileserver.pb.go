// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: pb/fileserver.proto

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

type UploadStatus int32

const (
	UploadStatus_UNKNOWN UploadStatus = 0
	UploadStatus_SUCCESS UploadStatus = 1
	UploadStatus_FAILED  UploadStatus = 2
)

// Enum value maps for UploadStatus.
var (
	UploadStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "SUCCESS",
		2: "FAILED",
	}
	UploadStatus_value = map[string]int32{
		"UNKNOWN": 0,
		"SUCCESS": 1,
		"FAILED":  2,
	}
)

func (x UploadStatus) Enum() *UploadStatus {
	p := new(UploadStatus)
	*p = x
	return p
}

func (x UploadStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UploadStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_fileserver_proto_enumTypes[0].Descriptor()
}

func (UploadStatus) Type() protoreflect.EnumType {
	return &file_pb_fileserver_proto_enumTypes[0]
}

func (x UploadStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UploadStatus.Descriptor instead.
func (UploadStatus) EnumDescriptor() ([]byte, []int) {
	return file_pb_fileserver_proto_rawDescGZIP(), []int{0}
}

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data     []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	FileName string `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Done     bool   `protobuf:"varint,3,opt,name=done,proto3" json:"done,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	mi := &file_pb_fileserver_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_pb_fileserver_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_pb_fileserver_proto_rawDescGZIP(), []int{0}
}

func (x *Chunk) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Chunk) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *Chunk) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

type DownloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
}

func (x *DownloadRequest) Reset() {
	*x = DownloadRequest{}
	mi := &file_pb_fileserver_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadRequest) ProtoMessage() {}

func (x *DownloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_fileserver_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadRequest.ProtoReflect.Descriptor instead.
func (*DownloadRequest) Descriptor() ([]byte, []int) {
	return file_pb_fileserver_proto_rawDescGZIP(), []int{1}
}

func (x *DownloadRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type UploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status UploadStatus `protobuf:"varint,1,opt,name=status,proto3,enum=UploadStatus" json:"status,omitempty"`
}

func (x *UploadResponse) Reset() {
	*x = UploadResponse{}
	mi := &file_pb_fileserver_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResponse) ProtoMessage() {}

func (x *UploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_fileserver_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadResponse.ProtoReflect.Descriptor instead.
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return file_pb_fileserver_proto_rawDescGZIP(), []int{2}
}

func (x *UploadResponse) GetStatus() UploadStatus {
	if x != nil {
		return x.Status
	}
	return UploadStatus_UNKNOWN
}

type EchoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *EchoMessage) Reset() {
	*x = EchoMessage{}
	mi := &file_pb_fileserver_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EchoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoMessage) ProtoMessage() {}

func (x *EchoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pb_fileserver_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoMessage.ProtoReflect.Descriptor instead.
func (*EchoMessage) Descriptor() ([]byte, []int) {
	return file_pb_fileserver_proto_rawDescGZIP(), []int{3}
}

func (x *EchoMessage) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type EchoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EchoMsg string `protobuf:"bytes,1,opt,name=echoMsg,proto3" json:"echoMsg,omitempty"`
}

func (x *EchoResponse) Reset() {
	*x = EchoResponse{}
	mi := &file_pb_fileserver_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EchoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoResponse) ProtoMessage() {}

func (x *EchoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_fileserver_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoResponse.ProtoReflect.Descriptor instead.
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return file_pb_fileserver_proto_rawDescGZIP(), []int{4}
}

func (x *EchoResponse) GetEchoMsg() string {
	if x != nil {
		return x.EchoMsg
	}
	return ""
}

var File_pb_fileserver_proto protoreflect.FileDescriptor

var file_pb_fileserver_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x62, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64,
	0x6f, 0x6e, 0x65, 0x22, 0x2e, 0x0a, 0x0f, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x37, 0x0a, 0x0e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x1f, 0x0a, 0x0b,
	0x45, 0x63, 0x68, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x28, 0x0a,
	0x0c, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x63, 0x68, 0x6f, 0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x65, 0x63, 0x68, 0x6f, 0x4d, 0x73, 0x67, 0x2a, 0x34, 0x0a, 0x0c, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x32, 0x76, 0x0a,
	0x02, 0x46, 0x53, 0x12, 0x23, 0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x06, 0x2e,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x0f, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x26, 0x0a, 0x08, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x10, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x06, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x30, 0x01,
	0x12, 0x23, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x0c, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0d, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_fileserver_proto_rawDescOnce sync.Once
	file_pb_fileserver_proto_rawDescData = file_pb_fileserver_proto_rawDesc
)

func file_pb_fileserver_proto_rawDescGZIP() []byte {
	file_pb_fileserver_proto_rawDescOnce.Do(func() {
		file_pb_fileserver_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_fileserver_proto_rawDescData)
	})
	return file_pb_fileserver_proto_rawDescData
}

var file_pb_fileserver_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_fileserver_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pb_fileserver_proto_goTypes = []any{
	(UploadStatus)(0),       // 0: UploadStatus
	(*Chunk)(nil),           // 1: Chunk
	(*DownloadRequest)(nil), // 2: DownloadRequest
	(*UploadResponse)(nil),  // 3: UploadResponse
	(*EchoMessage)(nil),     // 4: EchoMessage
	(*EchoResponse)(nil),    // 5: EchoResponse
}
var file_pb_fileserver_proto_depIdxs = []int32{
	0, // 0: UploadResponse.status:type_name -> UploadStatus
	1, // 1: FS.Upload:input_type -> Chunk
	2, // 2: FS.Download:input_type -> DownloadRequest
	4, // 3: FS.Echo:input_type -> EchoMessage
	3, // 4: FS.Upload:output_type -> UploadResponse
	1, // 5: FS.Download:output_type -> Chunk
	5, // 6: FS.Echo:output_type -> EchoResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_fileserver_proto_init() }
func file_pb_fileserver_proto_init() {
	if File_pb_fileserver_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_fileserver_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_fileserver_proto_goTypes,
		DependencyIndexes: file_pb_fileserver_proto_depIdxs,
		EnumInfos:         file_pb_fileserver_proto_enumTypes,
		MessageInfos:      file_pb_fileserver_proto_msgTypes,
	}.Build()
	File_pb_fileserver_proto = out.File
	file_pb_fileserver_proto_rawDesc = nil
	file_pb_fileserver_proto_goTypes = nil
	file_pb_fileserver_proto_depIdxs = nil
}
