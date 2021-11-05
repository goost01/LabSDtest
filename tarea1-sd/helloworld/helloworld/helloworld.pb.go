// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.0
// source: helloworld/helloworld.proto

package helloworld

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
		mi := &file_helloworld_helloworld_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_helloworld_proto_msgTypes[0]
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
	return file_helloworld_helloworld_proto_rawDescGZIP(), []int{0}
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
		mi := &file_helloworld_helloworld_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_helloworld_proto_msgTypes[1]
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
	return file_helloworld_helloworld_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// The request message
type JoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *JoinRequest) Reset() {
	*x = JoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_helloworld_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRequest) ProtoMessage() {}

func (x *JoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_helloworld_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRequest.ProtoReflect.Descriptor instead.
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_helloworld_proto_rawDescGZIP(), []int{2}
}

func (x *JoinRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *JoinRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message
type JoinReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *JoinReply) Reset() {
	*x = JoinReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_helloworld_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinReply) ProtoMessage() {}

func (x *JoinReply) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_helloworld_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinReply.ProtoReflect.Descriptor instead.
func (*JoinReply) Descriptor() ([]byte, []int) {
	return file_helloworld_helloworld_proto_rawDescGZIP(), []int{3}
}

func (x *JoinReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type BeginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *BeginRequest) Reset() {
	*x = BeginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_helloworld_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeginRequest) ProtoMessage() {}

func (x *BeginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_helloworld_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeginRequest.ProtoReflect.Descriptor instead.
func (*BeginRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_helloworld_proto_rawDescGZIP(), []int{4}
}

func (x *BeginRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type BeginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message bool  `protobuf:"varint,1,opt,name=message,proto3" json:"message,omitempty"`
	Online  int32 `protobuf:"varint,2,opt,name=online,proto3" json:"online,omitempty"`
}

func (x *BeginReply) Reset() {
	*x = BeginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_helloworld_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeginReply) ProtoMessage() {}

func (x *BeginReply) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_helloworld_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeginReply.ProtoReflect.Descriptor instead.
func (*BeginReply) Descriptor() ([]byte, []int) {
	return file_helloworld_helloworld_proto_rawDescGZIP(), []int{5}
}

func (x *BeginReply) GetMessage() bool {
	if x != nil {
		return x.Message
	}
	return false
}

func (x *BeginReply) GetOnline() int32 {
	if x != nil {
		return x.Online
	}
	return 0
}

// End request message
type EndRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EndRequest) Reset() {
	*x = EndRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_helloworld_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndRequest) ProtoMessage() {}

func (x *EndRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_helloworld_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndRequest.ProtoReflect.Descriptor instead.
func (*EndRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_helloworld_proto_rawDescGZIP(), []int{6}
}

func (x *EndRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// End reply message
type EndReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EndReply) Reset() {
	*x = EndReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_helloworld_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndReply) ProtoMessage() {}

func (x *EndReply) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_helloworld_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndReply.ProtoReflect.Descriptor instead.
func (*EndReply) Descriptor() ([]byte, []int) {
	return file_helloworld_helloworld_proto_rawDescGZIP(), []int{7}
}

func (x *EndReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// BeginStageRequest
type BeginStageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stage string `protobuf:"bytes,1,opt,name=stage,proto3" json:"stage,omitempty"` // enviar la etapa
}

func (x *BeginStageRequest) Reset() {
	*x = BeginStageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_helloworld_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeginStageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeginStageRequest) ProtoMessage() {}

func (x *BeginStageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_helloworld_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeginStageRequest.ProtoReflect.Descriptor instead.
func (*BeginStageRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_helloworld_proto_rawDescGZIP(), []int{8}
}

func (x *BeginStageRequest) GetStage() string {
	if x != nil {
		return x.Stage
	}
	return ""
}

// BeginStageReply
type BeginStageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stage bool `protobuf:"varint,1,opt,name=stage,proto3" json:"stage,omitempty"`
}

func (x *BeginStageReply) Reset() {
	*x = BeginStageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_helloworld_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeginStageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeginStageReply) ProtoMessage() {}

func (x *BeginStageReply) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_helloworld_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeginStageReply.ProtoReflect.Descriptor instead.
func (*BeginStageReply) Descriptor() ([]byte, []int) {
	return file_helloworld_helloworld_proto_rawDescGZIP(), []int{9}
}

func (x *BeginStageReply) GetStage() bool {
	if x != nil {
		return x.Stage
	}
	return false
}

// JugadaE1
type JugadaE1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jugada     int32 `protobuf:"varint,1,opt,name=jugada,proto3" json:"jugada,omitempty"`
	SumaActual int32 `protobuf:"varint,2,opt,name=suma_actual,json=sumaActual,proto3" json:"suma_actual,omitempty"`
}

func (x *JugadaE1) Reset() {
	*x = JugadaE1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_helloworld_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JugadaE1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JugadaE1) ProtoMessage() {}

func (x *JugadaE1) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_helloworld_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JugadaE1.ProtoReflect.Descriptor instead.
func (*JugadaE1) Descriptor() ([]byte, []int) {
	return file_helloworld_helloworld_proto_rawDescGZIP(), []int{10}
}

func (x *JugadaE1) GetJugada() int32 {
	if x != nil {
		return x.Jugada
	}
	return 0
}

func (x *JugadaE1) GetSumaActual() int32 {
	if x != nil {
		return x.SumaActual
	}
	return 0
}

type PlayerStatusE1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SumaTotal int32 `protobuf:"varint,1,opt,name=suma_total,json=sumaTotal,proto3" json:"suma_total,omitempty"`
	Dead      bool  `protobuf:"varint,2,opt,name=dead,proto3" json:"dead,omitempty"`
	Ronda     int32 `protobuf:"varint,3,opt,name=ronda,proto3" json:"ronda,omitempty"`
}

func (x *PlayerStatusE1) Reset() {
	*x = PlayerStatusE1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_helloworld_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerStatusE1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerStatusE1) ProtoMessage() {}

func (x *PlayerStatusE1) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_helloworld_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerStatusE1.ProtoReflect.Descriptor instead.
func (*PlayerStatusE1) Descriptor() ([]byte, []int) {
	return file_helloworld_helloworld_proto_rawDescGZIP(), []int{11}
}

func (x *PlayerStatusE1) GetSumaTotal() int32 {
	if x != nil {
		return x.SumaTotal
	}
	return 0
}

func (x *PlayerStatusE1) GetDead() bool {
	if x != nil {
		return x.Dead
	}
	return false
}

func (x *PlayerStatusE1) GetRonda() int32 {
	if x != nil {
		return x.Ronda
	}
	return 0
}

var File_helloworld_helloworld_proto protoreflect.FileDescriptor

var file_helloworld_helloworld_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x22, 0x22, 0x0a, 0x0c, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a,
	0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3b, 0x0a, 0x0b, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x25, 0x0a, 0x09, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x28, 0x0a, 0x0c, 0x42, 0x65, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x3e, 0x0a, 0x0a, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x22, 0x26, 0x0a, 0x0a, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x24, 0x0a, 0x08, 0x45,
	0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x29, 0x0a, 0x11, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x22, 0x27, 0x0a, 0x0f,
	0x42, 0x65, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x67, 0x65, 0x22, 0x43, 0x0a, 0x08, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x45,
	0x31, 0x12, 0x16, 0x0a, 0x06, 0x6a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x6a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x75, 0x6d,
	0x61, 0x5f, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x73, 0x75, 0x6d, 0x61, 0x41, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x22, 0x59, 0x0a, 0x0e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x31, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x75, 0x6d, 0x61, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x73, 0x75, 0x6d, 0x61, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x65, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x65, 0x61, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6e, 0x64, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x72, 0x6f, 0x6e, 0x64, 0x61, 0x32, 0x8e, 0x01, 0x0a, 0x07, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65,
	0x72, 0x12, 0x3e, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x18, 0x2e,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x43, 0x0a, 0x0d, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x41, 0x67, 0x61,
	0x69, 0x6e, 0x12, 0x18, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x32, 0xee, 0x03, 0x0a, 0x04, 0x47, 0x61, 0x6d, 0x65, 0x12,
	0x3c, 0x0a, 0x08, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x2e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3f, 0x0a,
	0x09, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x2e, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x39,
	0x0a, 0x07, 0x45, 0x6e, 0x64, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x45,
	0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0c, 0x42, 0x65, 0x67,
	0x69, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x45, 0x31, 0x12, 0x1d, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0c, 0x42, 0x65, 0x67, 0x69, 0x6e,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x45, 0x32, 0x12, 0x1d, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0c, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x45, 0x33, 0x12, 0x1d, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x4a, 0x75, 0x67, 0x61, 0x64,
	0x61, 0x45, 0x31, 0x12, 0x14, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2e, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x45, 0x31, 0x1a, 0x1a, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x45, 0x31, 0x22, 0x00, 0x42, 0x46, 0x0a, 0x1b, 0x69, 0x6f, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x42, 0x0f, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72,
	0x6c, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x14, 0x6c, 0x61, 0x62, 0x31, 0x2f,
	0x67, 0x61, 0x6d, 0x65, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_helloworld_helloworld_proto_rawDescOnce sync.Once
	file_helloworld_helloworld_proto_rawDescData = file_helloworld_helloworld_proto_rawDesc
)

func file_helloworld_helloworld_proto_rawDescGZIP() []byte {
	file_helloworld_helloworld_proto_rawDescOnce.Do(func() {
		file_helloworld_helloworld_proto_rawDescData = protoimpl.X.CompressGZIP(file_helloworld_helloworld_proto_rawDescData)
	})
	return file_helloworld_helloworld_proto_rawDescData
}

var file_helloworld_helloworld_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_helloworld_helloworld_proto_goTypes = []interface{}{
	(*HelloRequest)(nil),      // 0: helloworld.HelloRequest
	(*HelloReply)(nil),        // 1: helloworld.HelloReply
	(*JoinRequest)(nil),       // 2: helloworld.JoinRequest
	(*JoinReply)(nil),         // 3: helloworld.JoinReply
	(*BeginRequest)(nil),      // 4: helloworld.BeginRequest
	(*BeginReply)(nil),        // 5: helloworld.BeginReply
	(*EndRequest)(nil),        // 6: helloworld.EndRequest
	(*EndReply)(nil),          // 7: helloworld.EndReply
	(*BeginStageRequest)(nil), // 8: helloworld.BeginStageRequest
	(*BeginStageReply)(nil),   // 9: helloworld.BeginStageReply
	(*JugadaE1)(nil),          // 10: helloworld.JugadaE1
	(*PlayerStatusE1)(nil),    // 11: helloworld.PlayerStatusE1
}
var file_helloworld_helloworld_proto_depIdxs = []int32{
	0,  // 0: helloworld.Greeter.SayHello:input_type -> helloworld.HelloRequest
	0,  // 1: helloworld.Greeter.SayHelloAgain:input_type -> helloworld.HelloRequest
	2,  // 2: helloworld.Game.JoinGame:input_type -> helloworld.JoinRequest
	4,  // 3: helloworld.Game.BeginGame:input_type -> helloworld.BeginRequest
	6,  // 4: helloworld.Game.EndGame:input_type -> helloworld.EndRequest
	8,  // 5: helloworld.Game.BeginStageE1:input_type -> helloworld.BeginStageRequest
	8,  // 6: helloworld.Game.BeginStageE2:input_type -> helloworld.BeginStageRequest
	8,  // 7: helloworld.Game.BeginStageE3:input_type -> helloworld.BeginStageRequest
	10, // 8: helloworld.Game.SendJugadaE1:input_type -> helloworld.JugadaE1
	1,  // 9: helloworld.Greeter.SayHello:output_type -> helloworld.HelloReply
	1,  // 10: helloworld.Greeter.SayHelloAgain:output_type -> helloworld.HelloReply
	3,  // 11: helloworld.Game.JoinGame:output_type -> helloworld.JoinReply
	5,  // 12: helloworld.Game.BeginGame:output_type -> helloworld.BeginReply
	7,  // 13: helloworld.Game.EndGame:output_type -> helloworld.EndReply
	9,  // 14: helloworld.Game.BeginStageE1:output_type -> helloworld.BeginStageReply
	9,  // 15: helloworld.Game.BeginStageE2:output_type -> helloworld.BeginStageReply
	9,  // 16: helloworld.Game.BeginStageE3:output_type -> helloworld.BeginStageReply
	11, // 17: helloworld.Game.SendJugadaE1:output_type -> helloworld.PlayerStatusE1
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_helloworld_helloworld_proto_init() }
func file_helloworld_helloworld_proto_init() {
	if File_helloworld_helloworld_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_helloworld_helloworld_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_helloworld_helloworld_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_helloworld_helloworld_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRequest); i {
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
		file_helloworld_helloworld_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinReply); i {
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
		file_helloworld_helloworld_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeginRequest); i {
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
		file_helloworld_helloworld_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeginReply); i {
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
		file_helloworld_helloworld_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndRequest); i {
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
		file_helloworld_helloworld_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndReply); i {
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
		file_helloworld_helloworld_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeginStageRequest); i {
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
		file_helloworld_helloworld_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeginStageReply); i {
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
		file_helloworld_helloworld_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JugadaE1); i {
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
		file_helloworld_helloworld_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerStatusE1); i {
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
			RawDescriptor: file_helloworld_helloworld_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_helloworld_helloworld_proto_goTypes,
		DependencyIndexes: file_helloworld_helloworld_proto_depIdxs,
		MessageInfos:      file_helloworld_helloworld_proto_msgTypes,
	}.Build()
	File_helloworld_helloworld_proto = out.File
	file_helloworld_helloworld_proto_rawDesc = nil
	file_helloworld_helloworld_proto_goTypes = nil
	file_helloworld_helloworld_proto_depIdxs = nil
}