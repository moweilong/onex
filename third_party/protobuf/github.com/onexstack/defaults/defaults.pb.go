// Copyright 2021 Linka Cloud  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: defaults/defaults.proto

package defaults

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// FieldDefaults encapsulates the default values for each type of field. Depending on the
// field, the correct set should be used to ensure proper defaults generation.
type FieldDefaults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*FieldDefaults_Float
	//	*FieldDefaults_Double
	//	*FieldDefaults_Int32
	//	*FieldDefaults_Int64
	//	*FieldDefaults_Uint32
	//	*FieldDefaults_Uint64
	//	*FieldDefaults_Sint32
	//	*FieldDefaults_Sint64
	//	*FieldDefaults_Fixed32
	//	*FieldDefaults_Fixed64
	//	*FieldDefaults_Sfixed32
	//	*FieldDefaults_Sfixed64
	//	*FieldDefaults_Bool
	//	*FieldDefaults_String_
	//	*FieldDefaults_Bytes
	//	*FieldDefaults_Enum
	//	*FieldDefaults_Message
	//	*FieldDefaults_Duration
	//	*FieldDefaults_Timestamp
	Type isFieldDefaults_Type `protobuf_oneof:"type"`
}

func (x *FieldDefaults) Reset() {
	*x = FieldDefaults{}
	if protoimpl.UnsafeEnabled {
		mi := &file_defaults_defaults_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldDefaults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldDefaults) ProtoMessage() {}

func (x *FieldDefaults) ProtoReflect() protoreflect.Message {
	mi := &file_defaults_defaults_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldDefaults.ProtoReflect.Descriptor instead.
func (*FieldDefaults) Descriptor() ([]byte, []int) {
	return file_defaults_defaults_proto_rawDescGZIP(), []int{0}
}

func (m *FieldDefaults) GetType() isFieldDefaults_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *FieldDefaults) GetFloat() float32 {
	if x, ok := x.GetType().(*FieldDefaults_Float); ok {
		return x.Float
	}
	return 0
}

func (x *FieldDefaults) GetDouble() float64 {
	if x, ok := x.GetType().(*FieldDefaults_Double); ok {
		return x.Double
	}
	return 0
}

func (x *FieldDefaults) GetInt32() int32 {
	if x, ok := x.GetType().(*FieldDefaults_Int32); ok {
		return x.Int32
	}
	return 0
}

func (x *FieldDefaults) GetInt64() int64 {
	if x, ok := x.GetType().(*FieldDefaults_Int64); ok {
		return x.Int64
	}
	return 0
}

func (x *FieldDefaults) GetUint32() uint32 {
	if x, ok := x.GetType().(*FieldDefaults_Uint32); ok {
		return x.Uint32
	}
	return 0
}

func (x *FieldDefaults) GetUint64() uint64 {
	if x, ok := x.GetType().(*FieldDefaults_Uint64); ok {
		return x.Uint64
	}
	return 0
}

func (x *FieldDefaults) GetSint32() int32 {
	if x, ok := x.GetType().(*FieldDefaults_Sint32); ok {
		return x.Sint32
	}
	return 0
}

func (x *FieldDefaults) GetSint64() int64 {
	if x, ok := x.GetType().(*FieldDefaults_Sint64); ok {
		return x.Sint64
	}
	return 0
}

func (x *FieldDefaults) GetFixed32() uint32 {
	if x, ok := x.GetType().(*FieldDefaults_Fixed32); ok {
		return x.Fixed32
	}
	return 0
}

func (x *FieldDefaults) GetFixed64() uint64 {
	if x, ok := x.GetType().(*FieldDefaults_Fixed64); ok {
		return x.Fixed64
	}
	return 0
}

func (x *FieldDefaults) GetSfixed32() int32 {
	if x, ok := x.GetType().(*FieldDefaults_Sfixed32); ok {
		return x.Sfixed32
	}
	return 0
}

func (x *FieldDefaults) GetSfixed64() int64 {
	if x, ok := x.GetType().(*FieldDefaults_Sfixed64); ok {
		return x.Sfixed64
	}
	return 0
}

func (x *FieldDefaults) GetBool() bool {
	if x, ok := x.GetType().(*FieldDefaults_Bool); ok {
		return x.Bool
	}
	return false
}

func (x *FieldDefaults) GetString_() string {
	if x, ok := x.GetType().(*FieldDefaults_String_); ok {
		return x.String_
	}
	return ""
}

func (x *FieldDefaults) GetBytes() []byte {
	if x, ok := x.GetType().(*FieldDefaults_Bytes); ok {
		return x.Bytes
	}
	return nil
}

func (x *FieldDefaults) GetEnum() uint32 {
	if x, ok := x.GetType().(*FieldDefaults_Enum); ok {
		return x.Enum
	}
	return 0
}

func (x *FieldDefaults) GetMessage() *MessageDefaults {
	if x, ok := x.GetType().(*FieldDefaults_Message); ok {
		return x.Message
	}
	return nil
}

func (x *FieldDefaults) GetDuration() string {
	if x, ok := x.GetType().(*FieldDefaults_Duration); ok {
		return x.Duration
	}
	return ""
}

func (x *FieldDefaults) GetTimestamp() string {
	if x, ok := x.GetType().(*FieldDefaults_Timestamp); ok {
		return x.Timestamp
	}
	return ""
}

type isFieldDefaults_Type interface {
	isFieldDefaults_Type()
}

type FieldDefaults_Float struct {
	// Scalar Field Types
	Float float32 `protobuf:"fixed32,1,opt,name=float,oneof"`
}

type FieldDefaults_Double struct {
	Double float64 `protobuf:"fixed64,2,opt,name=double,oneof"`
}

type FieldDefaults_Int32 struct {
	Int32 int32 `protobuf:"varint,3,opt,name=int32,oneof"`
}

type FieldDefaults_Int64 struct {
	Int64 int64 `protobuf:"varint,4,opt,name=int64,oneof"`
}

type FieldDefaults_Uint32 struct {
	Uint32 uint32 `protobuf:"varint,5,opt,name=uint32,oneof"`
}

type FieldDefaults_Uint64 struct {
	Uint64 uint64 `protobuf:"varint,6,opt,name=uint64,oneof"`
}

type FieldDefaults_Sint32 struct {
	Sint32 int32 `protobuf:"zigzag32,7,opt,name=sint32,oneof"`
}

type FieldDefaults_Sint64 struct {
	Sint64 int64 `protobuf:"zigzag64,8,opt,name=sint64,oneof"`
}

type FieldDefaults_Fixed32 struct {
	Fixed32 uint32 `protobuf:"fixed32,9,opt,name=fixed32,oneof"`
}

type FieldDefaults_Fixed64 struct {
	Fixed64 uint64 `protobuf:"fixed64,10,opt,name=fixed64,oneof"`
}

type FieldDefaults_Sfixed32 struct {
	Sfixed32 int32 `protobuf:"fixed32,11,opt,name=sfixed32,oneof"`
}

type FieldDefaults_Sfixed64 struct {
	Sfixed64 int64 `protobuf:"fixed64,12,opt,name=sfixed64,oneof"`
}

type FieldDefaults_Bool struct {
	Bool bool `protobuf:"varint,13,opt,name=bool,oneof"`
}

type FieldDefaults_String_ struct {
	String_ string `protobuf:"bytes,14,opt,name=string,oneof"`
}

type FieldDefaults_Bytes struct {
	Bytes []byte `protobuf:"bytes,15,opt,name=bytes,oneof"`
}

type FieldDefaults_Enum struct {
	// Complex Field Types
	Enum uint32 `protobuf:"varint,16,opt,name=enum,oneof"`
}

type FieldDefaults_Message struct {
	Message *MessageDefaults `protobuf:"bytes,17,opt,name=message,oneof"`
}

type FieldDefaults_Duration struct {
	// Well-Known Field Types
	// any       = 20;
	Duration string `protobuf:"bytes,21,opt,name=duration,oneof"`
}

type FieldDefaults_Timestamp struct {
	Timestamp string `protobuf:"bytes,22,opt,name=timestamp,oneof"`
}

func (*FieldDefaults_Float) isFieldDefaults_Type() {}

func (*FieldDefaults_Double) isFieldDefaults_Type() {}

func (*FieldDefaults_Int32) isFieldDefaults_Type() {}

func (*FieldDefaults_Int64) isFieldDefaults_Type() {}

func (*FieldDefaults_Uint32) isFieldDefaults_Type() {}

func (*FieldDefaults_Uint64) isFieldDefaults_Type() {}

func (*FieldDefaults_Sint32) isFieldDefaults_Type() {}

func (*FieldDefaults_Sint64) isFieldDefaults_Type() {}

func (*FieldDefaults_Fixed32) isFieldDefaults_Type() {}

func (*FieldDefaults_Fixed64) isFieldDefaults_Type() {}

func (*FieldDefaults_Sfixed32) isFieldDefaults_Type() {}

func (*FieldDefaults_Sfixed64) isFieldDefaults_Type() {}

func (*FieldDefaults_Bool) isFieldDefaults_Type() {}

func (*FieldDefaults_String_) isFieldDefaults_Type() {}

func (*FieldDefaults_Bytes) isFieldDefaults_Type() {}

func (*FieldDefaults_Enum) isFieldDefaults_Type() {}

func (*FieldDefaults_Message) isFieldDefaults_Type() {}

func (*FieldDefaults_Duration) isFieldDefaults_Type() {}

func (*FieldDefaults_Timestamp) isFieldDefaults_Type() {}

// MessageDefaults define the default behaviour for this field.
type MessageDefaults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Initialize specify that the message should be initialized
	Initialize *bool `protobuf:"varint,1,opt,name=initialize" json:"initialize,omitempty"`
	// Defaults specifies that the messages' defaults should be applied
	Defaults *bool `protobuf:"varint,2,opt,name=defaults" json:"defaults,omitempty"`
}

func (x *MessageDefaults) Reset() {
	*x = MessageDefaults{}
	if protoimpl.UnsafeEnabled {
		mi := &file_defaults_defaults_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageDefaults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageDefaults) ProtoMessage() {}

func (x *MessageDefaults) ProtoReflect() protoreflect.Message {
	mi := &file_defaults_defaults_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageDefaults.ProtoReflect.Descriptor instead.
func (*MessageDefaults) Descriptor() ([]byte, []int) {
	return file_defaults_defaults_proto_rawDescGZIP(), []int{1}
}

func (x *MessageDefaults) GetInitialize() bool {
	if x != nil && x.Initialize != nil {
		return *x.Initialize
	}
	return false
}

func (x *MessageDefaults) GetDefaults() bool {
	if x != nil && x.Defaults != nil {
		return *x.Defaults
	}
	return false
}

var file_defaults_defaults_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         1171,
		Name:          "defaults.disabled",
		Tag:           "varint,1171,opt,name=disabled",
		Filename:      "defaults/defaults.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         1172,
		Name:          "defaults.ignored",
		Tag:           "varint,1172,opt,name=ignored",
		Filename:      "defaults/defaults.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         1173,
		Name:          "defaults.unexported",
		Tag:           "varint,1173,opt,name=unexported",
		Filename:      "defaults/defaults.proto",
	},
	{
		ExtendedType:  (*descriptorpb.OneofOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         1171,
		Name:          "defaults.oneof",
		Tag:           "bytes,1171,opt,name=oneof",
		Filename:      "defaults/defaults.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*FieldDefaults)(nil),
		Field:         1171,
		Name:          "defaults.value",
		Tag:           "bytes,1171,opt,name=value",
		Filename:      "defaults/defaults.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// Disabled nullifies any defaults for this message, including any
	// message fields associated with it that do support defaults.
	//
	// optional bool disabled = 1171;
	E_Disabled = &file_defaults_defaults_proto_extTypes[0]
	// Ignore skips generation of default methods for this message.
	//
	// optional bool ignored = 1172;
	E_Ignored = &file_defaults_defaults_proto_extTypes[1]
	// Unexported generate an unexported defaults method, this can
	// be useful when we want both the generated defaults and a custom
	// defaults method that will call the unexported method.
	//
	// optional bool unexported = 1173;
	E_Unexported = &file_defaults_defaults_proto_extTypes[2]
)

// Extension fields to descriptorpb.OneofOptions.
var (
	// optional string oneof = 1171;
	E_Oneof = &file_defaults_defaults_proto_extTypes[3]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// Value specify the default value to set on this field. By default,
	// none is set on a field.
	//
	// optional defaults.FieldDefaults value = 1171;
	E_Value = &file_defaults_defaults_proto_extTypes[4]
)

var File_defaults_defaults_proto protoreflect.FileDescriptor

var file_defaults_defaults_proto_rawDesc = []byte{
	0x0a, 0x17, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x2f, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x04, 0x0a, 0x0d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x12,
	0x18, 0x0a, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x48,
	0x00, 0x52, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x05, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x12, 0x16, 0x0a, 0x05, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x00, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x18, 0x0a, 0x06, 0x75, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x06, 0x75, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x12, 0x18, 0x0a, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x18, 0x0a,
	0x06, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x07, 0x20, 0x01, 0x28, 0x11, 0x48, 0x00, 0x52,
	0x06, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x18, 0x08, 0x20, 0x01, 0x28, 0x12, 0x48, 0x00, 0x52, 0x06, 0x73, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x12, 0x1a, 0x0a, 0x07, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x07, 0x48, 0x00, 0x52, 0x07, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x1a, 0x0a,
	0x07, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x06, 0x48, 0x00,
	0x52, 0x07, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x1c, 0x0a, 0x08, 0x73, 0x66, 0x69,
	0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0f, 0x48, 0x00, 0x52, 0x08, 0x73,
	0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x1c, 0x0a, 0x08, 0x73, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x36, 0x34, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x10, 0x48, 0x00, 0x52, 0x08, 0x73, 0x66, 0x69,
	0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x14, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x12, 0x18, 0x0a, 0x06, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x14, 0x0a,
	0x04, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x04, 0x65,
	0x6e, 0x75, 0x6d, 0x12, 0x35, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x48,
	0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x4a, 0x04, 0x08, 0x12, 0x10, 0x15, 0x22, 0x4d, 0x0a, 0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x73, 0x3a, 0x3c, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x93, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x3a, 0x3a, 0x0a, 0x07, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x64, 0x12, 0x1f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x94, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x64, 0x3a,
	0x40, 0x0a, 0x0a, 0x75, 0x6e, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x1f, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x95,
	0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x75, 0x6e, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x64, 0x3a, 0x34, 0x0a, 0x05, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4f, 0x6e, 0x65,
	0x6f, 0x66, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x93, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x3a, 0x4d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x93, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x73, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x6f, 0x2e, 0x6c, 0x69, 0x6e,
	0x6b, 0x61, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x2f, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x73, 0x3b, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73,
}

var (
	file_defaults_defaults_proto_rawDescOnce sync.Once
	file_defaults_defaults_proto_rawDescData = file_defaults_defaults_proto_rawDesc
)

func file_defaults_defaults_proto_rawDescGZIP() []byte {
	file_defaults_defaults_proto_rawDescOnce.Do(func() {
		file_defaults_defaults_proto_rawDescData = protoimpl.X.CompressGZIP(file_defaults_defaults_proto_rawDescData)
	})
	return file_defaults_defaults_proto_rawDescData
}

var file_defaults_defaults_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_defaults_defaults_proto_goTypes = []interface{}{
	(*FieldDefaults)(nil),               // 0: defaults.FieldDefaults
	(*MessageDefaults)(nil),             // 1: defaults.MessageDefaults
	(*descriptorpb.MessageOptions)(nil), // 2: google.protobuf.MessageOptions
	(*descriptorpb.OneofOptions)(nil),   // 3: google.protobuf.OneofOptions
	(*descriptorpb.FieldOptions)(nil),   // 4: google.protobuf.FieldOptions
}
var file_defaults_defaults_proto_depIdxs = []int32{
	1, // 0: defaults.FieldDefaults.message:type_name -> defaults.MessageDefaults
	2, // 1: defaults.disabled:extendee -> google.protobuf.MessageOptions
	2, // 2: defaults.ignored:extendee -> google.protobuf.MessageOptions
	2, // 3: defaults.unexported:extendee -> google.protobuf.MessageOptions
	3, // 4: defaults.oneof:extendee -> google.protobuf.OneofOptions
	4, // 5: defaults.value:extendee -> google.protobuf.FieldOptions
	0, // 6: defaults.value:type_name -> defaults.FieldDefaults
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	6, // [6:7] is the sub-list for extension type_name
	1, // [1:6] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_defaults_defaults_proto_init() }
func file_defaults_defaults_proto_init() {
	if File_defaults_defaults_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_defaults_defaults_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldDefaults); i {
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
		file_defaults_defaults_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageDefaults); i {
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
	file_defaults_defaults_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*FieldDefaults_Float)(nil),
		(*FieldDefaults_Double)(nil),
		(*FieldDefaults_Int32)(nil),
		(*FieldDefaults_Int64)(nil),
		(*FieldDefaults_Uint32)(nil),
		(*FieldDefaults_Uint64)(nil),
		(*FieldDefaults_Sint32)(nil),
		(*FieldDefaults_Sint64)(nil),
		(*FieldDefaults_Fixed32)(nil),
		(*FieldDefaults_Fixed64)(nil),
		(*FieldDefaults_Sfixed32)(nil),
		(*FieldDefaults_Sfixed64)(nil),
		(*FieldDefaults_Bool)(nil),
		(*FieldDefaults_String_)(nil),
		(*FieldDefaults_Bytes)(nil),
		(*FieldDefaults_Enum)(nil),
		(*FieldDefaults_Message)(nil),
		(*FieldDefaults_Duration)(nil),
		(*FieldDefaults_Timestamp)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_defaults_defaults_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 5,
			NumServices:   0,
		},
		GoTypes:           file_defaults_defaults_proto_goTypes,
		DependencyIndexes: file_defaults_defaults_proto_depIdxs,
		MessageInfos:      file_defaults_defaults_proto_msgTypes,
		ExtensionInfos:    file_defaults_defaults_proto_extTypes,
	}.Build()
	File_defaults_defaults_proto = out.File
	file_defaults_defaults_proto_rawDesc = nil
	file_defaults_defaults_proto_goTypes = nil
	file_defaults_defaults_proto_depIdxs = nil
}
