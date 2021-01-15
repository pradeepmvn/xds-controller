// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: example/proto/person.proto

package example

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PersonRequest) Reset() {
	*x = PersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_person_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonRequest) ProtoMessage() {}

func (x *PersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_person_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonRequest.ProtoReflect.Descriptor instead.
func (*PersonRequest) Descriptor() ([]byte, []int) {
	return file_example_proto_person_proto_rawDescGZIP(), []int{0}
}

type PersonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id   int32  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PersonResponse) Reset() {
	*x = PersonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_person_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonResponse) ProtoMessage() {}

func (x *PersonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_person_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonResponse.ProtoReflect.Descriptor instead.
func (*PersonResponse) Descriptor() ([]byte, []int) {
	return file_example_proto_person_proto_rawDescGZIP(), []int{1}
}

func (x *PersonResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PersonResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_example_proto_person_proto protoreflect.FileDescriptor

var file_example_proto_person_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x22, 0x0f, 0x0a, 0x0d, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x34, 0x0a, 0x0e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x32, 0x45, 0x0a, 0x06, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x14, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x72, 0x61, 0x64, 0x65, 0x65, 0x70, 0x6d, 0x76, 0x6e, 0x2f, 0x78, 0x64, 0x73, 0x2d, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_proto_person_proto_rawDescOnce sync.Once
	file_example_proto_person_proto_rawDescData = file_example_proto_person_proto_rawDesc
)

func file_example_proto_person_proto_rawDescGZIP() []byte {
	file_example_proto_person_proto_rawDescOnce.Do(func() {
		file_example_proto_person_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_proto_person_proto_rawDescData)
	})
	return file_example_proto_person_proto_rawDescData
}

var file_example_proto_person_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_example_proto_person_proto_goTypes = []interface{}{
	(*PersonRequest)(nil),  // 0: hello.PersonRequest
	(*PersonResponse)(nil), // 1: hello.PersonResponse
}
var file_example_proto_person_proto_depIdxs = []int32{
	0, // 0: hello.Person.GetDetails:input_type -> hello.PersonRequest
	1, // 1: hello.Person.GetDetails:output_type -> hello.PersonResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_example_proto_person_proto_init() }
func file_example_proto_person_proto_init() {
	if File_example_proto_person_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_proto_person_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonRequest); i {
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
		file_example_proto_person_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonResponse); i {
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
			RawDescriptor: file_example_proto_person_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_proto_person_proto_goTypes,
		DependencyIndexes: file_example_proto_person_proto_depIdxs,
		MessageInfos:      file_example_proto_person_proto_msgTypes,
	}.Build()
	File_example_proto_person_proto = out.File
	file_example_proto_person_proto_rawDesc = nil
	file_example_proto_person_proto_goTypes = nil
	file_example_proto_person_proto_depIdxs = nil
}
