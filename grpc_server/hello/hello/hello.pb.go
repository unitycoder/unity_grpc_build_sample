// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.1
// source: proto/hello.proto

package hello

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

type HelloEnum int32

const (
	HelloEnum_ZERO  HelloEnum = 0
	HelloEnum_ONE   HelloEnum = 1
	HelloEnum_TWO   HelloEnum = 2
	HelloEnum_THREE HelloEnum = 3
)

// Enum value maps for HelloEnum.
var (
	HelloEnum_name = map[int32]string{
		0: "ZERO",
		1: "ONE",
		2: "TWO",
		3: "THREE",
	}
	HelloEnum_value = map[string]int32{
		"ZERO":  0,
		"ONE":   1,
		"TWO":   2,
		"THREE": 3,
	}
)

func (x HelloEnum) Enum() *HelloEnum {
	p := new(HelloEnum)
	*p = x
	return p
}

func (x HelloEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HelloEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_hello_proto_enumTypes[0].Descriptor()
}

func (HelloEnum) Type() protoreflect.EnumType {
	return &file_proto_hello_proto_enumTypes[0]
}

func (x HelloEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HelloEnum.Descriptor instead.
func (HelloEnum) EnumDescriptor() ([]byte, []int) {
	return file_proto_hello_proto_rawDescGZIP(), []int{0}
}

// The request message containing the user's name.
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_proto_hello_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_proto_hello_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type HelloEnumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enum HelloEnum `protobuf:"varint,1,opt,name=enum,proto3,enum=helloworld.HelloEnum" json:"enum,omitempty"`
}

func (x *HelloEnumRequest) Reset() {
	*x = HelloEnumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloEnumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloEnumRequest) ProtoMessage() {}

func (x *HelloEnumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloEnumRequest.ProtoReflect.Descriptor instead.
func (*HelloEnumRequest) Descriptor() ([]byte, []int) {
	return file_proto_hello_proto_rawDescGZIP(), []int{2}
}

func (x *HelloEnumRequest) GetEnum() HelloEnum {
	if x != nil {
		return x.Enum
	}
	return HelloEnum_ZERO
}

type HelloEnumReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enum HelloEnum `protobuf:"varint,1,opt,name=enum,proto3,enum=helloworld.HelloEnum" json:"enum,omitempty"`
}

func (x *HelloEnumReply) Reset() {
	*x = HelloEnumReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloEnumReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloEnumReply) ProtoMessage() {}

func (x *HelloEnumReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloEnumReply.ProtoReflect.Descriptor instead.
func (*HelloEnumReply) Descriptor() ([]byte, []int) {
	return file_proto_hello_proto_rawDescGZIP(), []int{3}
}

func (x *HelloEnumReply) GetEnum() HelloEnum {
	if x != nil {
		return x.Enum
	}
	return HelloEnum_ZERO
}

type HelloOneOfRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Choose:
	//	*HelloOneOfRequest_First
	//	*HelloOneOfRequest_Second
	Choose isHelloOneOfRequest_Choose `protobuf_oneof:"choose"`
}

func (x *HelloOneOfRequest) Reset() {
	*x = HelloOneOfRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloOneOfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloOneOfRequest) ProtoMessage() {}

func (x *HelloOneOfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloOneOfRequest.ProtoReflect.Descriptor instead.
func (*HelloOneOfRequest) Descriptor() ([]byte, []int) {
	return file_proto_hello_proto_rawDescGZIP(), []int{4}
}

func (m *HelloOneOfRequest) GetChoose() isHelloOneOfRequest_Choose {
	if m != nil {
		return m.Choose
	}
	return nil
}

func (x *HelloOneOfRequest) GetFirst() string {
	if x, ok := x.GetChoose().(*HelloOneOfRequest_First); ok {
		return x.First
	}
	return ""
}

func (x *HelloOneOfRequest) GetSecond() int32 {
	if x, ok := x.GetChoose().(*HelloOneOfRequest_Second); ok {
		return x.Second
	}
	return 0
}

type isHelloOneOfRequest_Choose interface {
	isHelloOneOfRequest_Choose()
}

type HelloOneOfRequest_First struct {
	First string `protobuf:"bytes,1,opt,name=first,proto3,oneof"`
}

type HelloOneOfRequest_Second struct {
	Second int32 `protobuf:"varint,2,opt,name=second,proto3,oneof"`
}

func (*HelloOneOfRequest_First) isHelloOneOfRequest_Choose() {}

func (*HelloOneOfRequest_Second) isHelloOneOfRequest_Choose() {}

type HelloOneOfReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Choose:
	//	*HelloOneOfReply_First
	//	*HelloOneOfReply_Second
	Choose isHelloOneOfReply_Choose `protobuf_oneof:"choose"`
}

func (x *HelloOneOfReply) Reset() {
	*x = HelloOneOfReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloOneOfReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloOneOfReply) ProtoMessage() {}

func (x *HelloOneOfReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloOneOfReply.ProtoReflect.Descriptor instead.
func (*HelloOneOfReply) Descriptor() ([]byte, []int) {
	return file_proto_hello_proto_rawDescGZIP(), []int{5}
}

func (m *HelloOneOfReply) GetChoose() isHelloOneOfReply_Choose {
	if m != nil {
		return m.Choose
	}
	return nil
}

func (x *HelloOneOfReply) GetFirst() string {
	if x, ok := x.GetChoose().(*HelloOneOfReply_First); ok {
		return x.First
	}
	return ""
}

func (x *HelloOneOfReply) GetSecond() int32 {
	if x, ok := x.GetChoose().(*HelloOneOfReply_Second); ok {
		return x.Second
	}
	return 0
}

type isHelloOneOfReply_Choose interface {
	isHelloOneOfReply_Choose()
}

type HelloOneOfReply_First struct {
	First string `protobuf:"bytes,1,opt,name=first,proto3,oneof"`
}

type HelloOneOfReply_Second struct {
	Second int32 `protobuf:"varint,2,opt,name=second,proto3,oneof"`
}

func (*HelloOneOfReply_First) isHelloOneOfReply_Choose() {}

func (*HelloOneOfReply_Second) isHelloOneOfReply_Choose() {}

var File_proto_hello_proto protoreflect.FileDescriptor

var file_proto_hello_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x22,
	0x22, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3d, 0x0a, 0x10, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x29, 0x0a, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x45, 0x6e, 0x75, 0x6d, 0x52, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x22, 0x3b, 0x0a, 0x0e, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x29, 0x0a, 0x04,
	0x65, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x45, 0x6e, 0x75,
	0x6d, 0x52, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x22, 0x4f, 0x0a, 0x11, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x05,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x06, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x42, 0x08,
	0x0a, 0x06, 0x63, 0x68, 0x6f, 0x6f, 0x73, 0x65, 0x22, 0x4d, 0x0a, 0x0f, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x05, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x06, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x42, 0x08, 0x0a,
	0x06, 0x63, 0x68, 0x6f, 0x6f, 0x73, 0x65, 0x2a, 0x32, 0x0a, 0x09, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x45, 0x6e, 0x75, 0x6d, 0x12, 0x08, 0x0a, 0x04, 0x5a, 0x45, 0x52, 0x4f, 0x10, 0x00, 0x12, 0x07,
	0x0a, 0x03, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x57, 0x4f, 0x10, 0x02,
	0x12, 0x09, 0x0a, 0x05, 0x54, 0x48, 0x52, 0x45, 0x45, 0x10, 0x03, 0x32, 0x95, 0x01, 0x0a, 0x07,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x12, 0x18, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0c, 0x53, 0x61, 0x79, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x1c, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_hello_proto_rawDescOnce sync.Once
	file_proto_hello_proto_rawDescData = file_proto_hello_proto_rawDesc
)

func file_proto_hello_proto_rawDescGZIP() []byte {
	file_proto_hello_proto_rawDescOnce.Do(func() {
		file_proto_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_hello_proto_rawDescData)
	})
	return file_proto_hello_proto_rawDescData
}

var file_proto_hello_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_hello_proto_goTypes = []interface{}{
	(HelloEnum)(0),            // 0: helloworld.HelloEnum
	(*HelloRequest)(nil),      // 1: helloworld.HelloRequest
	(*HelloReply)(nil),        // 2: helloworld.HelloReply
	(*HelloEnumRequest)(nil),  // 3: helloworld.HelloEnumRequest
	(*HelloEnumReply)(nil),    // 4: helloworld.HelloEnumReply
	(*HelloOneOfRequest)(nil), // 5: helloworld.HelloOneOfRequest
	(*HelloOneOfReply)(nil),   // 6: helloworld.HelloOneOfReply
}
var file_proto_hello_proto_depIdxs = []int32{
	0, // 0: helloworld.HelloEnumRequest.enum:type_name -> helloworld.HelloEnum
	0, // 1: helloworld.HelloEnumReply.enum:type_name -> helloworld.HelloEnum
	1, // 2: helloworld.Greeter.SayHello:input_type -> helloworld.HelloRequest
	3, // 3: helloworld.Greeter.SayHelloEnum:input_type -> helloworld.HelloEnumRequest
	2, // 4: helloworld.Greeter.SayHello:output_type -> helloworld.HelloReply
	4, // 5: helloworld.Greeter.SayHelloEnum:output_type -> helloworld.HelloEnumReply
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_hello_proto_init() }
func file_proto_hello_proto_init() {
	if File_proto_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_proto_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
		file_proto_hello_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloEnumRequest); i {
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
		file_proto_hello_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloEnumReply); i {
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
		file_proto_hello_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloOneOfRequest); i {
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
		file_proto_hello_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloOneOfReply); i {
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
	file_proto_hello_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*HelloOneOfRequest_First)(nil),
		(*HelloOneOfRequest_Second)(nil),
	}
	file_proto_hello_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*HelloOneOfReply_First)(nil),
		(*HelloOneOfReply_Second)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_hello_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_hello_proto_goTypes,
		DependencyIndexes: file_proto_hello_proto_depIdxs,
		EnumInfos:         file_proto_hello_proto_enumTypes,
		MessageInfos:      file_proto_hello_proto_msgTypes,
	}.Build()
	File_proto_hello_proto = out.File
	file_proto_hello_proto_rawDesc = nil
	file_proto_hello_proto_goTypes = nil
	file_proto_hello_proto_depIdxs = nil
}
