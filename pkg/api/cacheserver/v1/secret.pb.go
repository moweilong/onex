// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.3
// source: cacheserver/v1/secret.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SetSecretRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Expire        *durationpb.Duration   `protobuf:"bytes,3,opt,name=expire,proto3,oneof" json:"expire,omitempty"`
	Description   string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetSecretRequest) Reset() {
	*x = SetSecretRequest{}
	mi := &file_cacheserver_v1_secret_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetSecretRequest) ProtoMessage() {}

func (x *SetSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cacheserver_v1_secret_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetSecretRequest.ProtoReflect.Descriptor instead.
func (*SetSecretRequest) Descriptor() ([]byte, []int) {
	return file_cacheserver_v1_secret_proto_rawDescGZIP(), []int{0}
}

func (x *SetSecretRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetSecretRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SetSecretRequest) GetExpire() *durationpb.Duration {
	if x != nil {
		return x.Expire
	}
	return nil
}

func (x *SetSecretRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type DelSecretRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DelSecretRequest) Reset() {
	*x = DelSecretRequest{}
	mi := &file_cacheserver_v1_secret_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DelSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelSecretRequest) ProtoMessage() {}

func (x *DelSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cacheserver_v1_secret_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelSecretRequest.ProtoReflect.Descriptor instead.
func (*DelSecretRequest) Descriptor() ([]byte, []int) {
	return file_cacheserver_v1_secret_proto_rawDescGZIP(), []int{1}
}

func (x *DelSecretRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetSecretRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSecretRequest) Reset() {
	*x = GetSecretRequest{}
	mi := &file_cacheserver_v1_secret_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecretRequest) ProtoMessage() {}

func (x *GetSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cacheserver_v1_secret_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecretRequest.ProtoReflect.Descriptor instead.
func (*GetSecretRequest) Descriptor() ([]byte, []int) {
	return file_cacheserver_v1_secret_proto_rawDescGZIP(), []int{2}
}

func (x *GetSecretRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetSecretResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserID        string                 `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	SecretID      string                 `protobuf:"bytes,3,opt,name=secretID,proto3" json:"secretID,omitempty"`
	SecretKey     string                 `protobuf:"bytes,4,opt,name=secretKey,proto3" json:"secretKey,omitempty"`
	Expires       int64                  `protobuf:"varint,5,opt,name=expires,proto3" json:"expires,omitempty"`
	Status        int32                  `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	Description   string                 `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSecretResponse) Reset() {
	*x = GetSecretResponse{}
	mi := &file_cacheserver_v1_secret_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSecretResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecretResponse) ProtoMessage() {}

func (x *GetSecretResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cacheserver_v1_secret_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecretResponse.ProtoReflect.Descriptor instead.
func (*GetSecretResponse) Descriptor() ([]byte, []int) {
	return file_cacheserver_v1_secret_proto_rawDescGZIP(), []int{3}
}

func (x *GetSecretResponse) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *GetSecretResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetSecretResponse) GetSecretID() string {
	if x != nil {
		return x.SecretID
	}
	return ""
}

func (x *GetSecretResponse) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *GetSecretResponse) GetExpires() int64 {
	if x != nil {
		return x.Expires
	}
	return 0
}

func (x *GetSecretResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetSecretResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetSecretResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *GetSecretResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_cacheserver_v1_secret_proto protoreflect.FileDescriptor

var file_cacheserver_v1_secret_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x61, 0x63, 0x68, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1e,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42,
	0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xfd, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x36,
	0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x06, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x72, 0x03, 0x18, 0x80, 0x02, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x22, 0x24, 0x0a,
	0x10, 0x44, 0x65, 0x6c, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x22, 0x24, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0xc1, 0x02, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x35, 0x5a,
	0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6e, 0x65, 0x78,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x6f, 0x6e, 0x65, 0x78, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cacheserver_v1_secret_proto_rawDescOnce sync.Once
	file_cacheserver_v1_secret_proto_rawDescData = file_cacheserver_v1_secret_proto_rawDesc
)

func file_cacheserver_v1_secret_proto_rawDescGZIP() []byte {
	file_cacheserver_v1_secret_proto_rawDescOnce.Do(func() {
		file_cacheserver_v1_secret_proto_rawDescData = protoimpl.X.CompressGZIP(file_cacheserver_v1_secret_proto_rawDescData)
	})
	return file_cacheserver_v1_secret_proto_rawDescData
}

var file_cacheserver_v1_secret_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cacheserver_v1_secret_proto_goTypes = []any{
	(*SetSecretRequest)(nil),      // 0: cacheserver.v1.SetSecretRequest
	(*DelSecretRequest)(nil),      // 1: cacheserver.v1.DelSecretRequest
	(*GetSecretRequest)(nil),      // 2: cacheserver.v1.GetSecretRequest
	(*GetSecretResponse)(nil),     // 3: cacheserver.v1.GetSecretResponse
	(*durationpb.Duration)(nil),   // 4: google.protobuf.Duration
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_cacheserver_v1_secret_proto_depIdxs = []int32{
	4, // 0: cacheserver.v1.SetSecretRequest.expire:type_name -> google.protobuf.Duration
	5, // 1: cacheserver.v1.GetSecretResponse.createdAt:type_name -> google.protobuf.Timestamp
	5, // 2: cacheserver.v1.GetSecretResponse.updatedAt:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_cacheserver_v1_secret_proto_init() }
func file_cacheserver_v1_secret_proto_init() {
	if File_cacheserver_v1_secret_proto != nil {
		return
	}
	file_cacheserver_v1_secret_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cacheserver_v1_secret_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cacheserver_v1_secret_proto_goTypes,
		DependencyIndexes: file_cacheserver_v1_secret_proto_depIdxs,
		MessageInfos:      file_cacheserver_v1_secret_proto_msgTypes,
	}.Build()
	File_cacheserver_v1_secret_proto = out.File
	file_cacheserver_v1_secret_proto_rawDesc = nil
	file_cacheserver_v1_secret_proto_goTypes = nil
	file_cacheserver_v1_secret_proto_depIdxs = nil
}
