//
// Copyright 2023 The Chainloop Authors.
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
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: controlplane/v1/workflow.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type WorkflowServiceCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// The ID of the workload schema to be associated with, if no provided one will be created for you
	SchemaId    string `protobuf:"bytes,3,opt,name=schema_id,json=schemaId,proto3" json:"schema_id,omitempty"`
	Team        string `protobuf:"bytes,4,opt,name=team,proto3" json:"team,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *WorkflowServiceCreateRequest) Reset() {
	*x = WorkflowServiceCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controlplane_v1_workflow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowServiceCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowServiceCreateRequest) ProtoMessage() {}

func (x *WorkflowServiceCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controlplane_v1_workflow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowServiceCreateRequest.ProtoReflect.Descriptor instead.
func (*WorkflowServiceCreateRequest) Descriptor() ([]byte, []int) {
	return file_controlplane_v1_workflow_proto_rawDescGZIP(), []int{0}
}

func (x *WorkflowServiceCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WorkflowServiceCreateRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *WorkflowServiceCreateRequest) GetSchemaId() string {
	if x != nil {
		return x.SchemaId
	}
	return ""
}

func (x *WorkflowServiceCreateRequest) GetTeam() string {
	if x != nil {
		return x.Team
	}
	return ""
}

func (x *WorkflowServiceCreateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type WorkflowServiceUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// "optional" allow us to detect if the value is explicitly set
	// and not just the default balue
	Name        *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Project     *string `protobuf:"bytes,3,opt,name=project,proto3,oneof" json:"project,omitempty"`
	Team        *string `protobuf:"bytes,4,opt,name=team,proto3,oneof" json:"team,omitempty"`
	Public      *bool   `protobuf:"varint,5,opt,name=public,proto3,oneof" json:"public,omitempty"`
	Description *string `protobuf:"bytes,6,opt,name=description,proto3,oneof" json:"description,omitempty"`
}

func (x *WorkflowServiceUpdateRequest) Reset() {
	*x = WorkflowServiceUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controlplane_v1_workflow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowServiceUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowServiceUpdateRequest) ProtoMessage() {}

func (x *WorkflowServiceUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controlplane_v1_workflow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowServiceUpdateRequest.ProtoReflect.Descriptor instead.
func (*WorkflowServiceUpdateRequest) Descriptor() ([]byte, []int) {
	return file_controlplane_v1_workflow_proto_rawDescGZIP(), []int{1}
}

func (x *WorkflowServiceUpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WorkflowServiceUpdateRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *WorkflowServiceUpdateRequest) GetProject() string {
	if x != nil && x.Project != nil {
		return *x.Project
	}
	return ""
}

func (x *WorkflowServiceUpdateRequest) GetTeam() string {
	if x != nil && x.Team != nil {
		return *x.Team
	}
	return ""
}

func (x *WorkflowServiceUpdateRequest) GetPublic() bool {
	if x != nil && x.Public != nil {
		return *x.Public
	}
	return false
}

func (x *WorkflowServiceUpdateRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

type WorkflowServiceUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *WorkflowItem `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *WorkflowServiceUpdateResponse) Reset() {
	*x = WorkflowServiceUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controlplane_v1_workflow_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowServiceUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowServiceUpdateResponse) ProtoMessage() {}

func (x *WorkflowServiceUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controlplane_v1_workflow_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowServiceUpdateResponse.ProtoReflect.Descriptor instead.
func (*WorkflowServiceUpdateResponse) Descriptor() ([]byte, []int) {
	return file_controlplane_v1_workflow_proto_rawDescGZIP(), []int{2}
}

func (x *WorkflowServiceUpdateResponse) GetResult() *WorkflowItem {
	if x != nil {
		return x.Result
	}
	return nil
}

type WorkflowServiceCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *WorkflowItem `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *WorkflowServiceCreateResponse) Reset() {
	*x = WorkflowServiceCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controlplane_v1_workflow_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowServiceCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowServiceCreateResponse) ProtoMessage() {}

func (x *WorkflowServiceCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controlplane_v1_workflow_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowServiceCreateResponse.ProtoReflect.Descriptor instead.
func (*WorkflowServiceCreateResponse) Descriptor() ([]byte, []int) {
	return file_controlplane_v1_workflow_proto_rawDescGZIP(), []int{3}
}

func (x *WorkflowServiceCreateResponse) GetResult() *WorkflowItem {
	if x != nil {
		return x.Result
	}
	return nil
}

type WorkflowServiceDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *WorkflowServiceDeleteRequest) Reset() {
	*x = WorkflowServiceDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controlplane_v1_workflow_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowServiceDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowServiceDeleteRequest) ProtoMessage() {}

func (x *WorkflowServiceDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controlplane_v1_workflow_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowServiceDeleteRequest.ProtoReflect.Descriptor instead.
func (*WorkflowServiceDeleteRequest) Descriptor() ([]byte, []int) {
	return file_controlplane_v1_workflow_proto_rawDescGZIP(), []int{4}
}

func (x *WorkflowServiceDeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type WorkflowServiceDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WorkflowServiceDeleteResponse) Reset() {
	*x = WorkflowServiceDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controlplane_v1_workflow_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowServiceDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowServiceDeleteResponse) ProtoMessage() {}

func (x *WorkflowServiceDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controlplane_v1_workflow_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowServiceDeleteResponse.ProtoReflect.Descriptor instead.
func (*WorkflowServiceDeleteResponse) Descriptor() ([]byte, []int) {
	return file_controlplane_v1_workflow_proto_rawDescGZIP(), []int{5}
}

type WorkflowServiceListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WorkflowServiceListRequest) Reset() {
	*x = WorkflowServiceListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controlplane_v1_workflow_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowServiceListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowServiceListRequest) ProtoMessage() {}

func (x *WorkflowServiceListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controlplane_v1_workflow_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowServiceListRequest.ProtoReflect.Descriptor instead.
func (*WorkflowServiceListRequest) Descriptor() ([]byte, []int) {
	return file_controlplane_v1_workflow_proto_rawDescGZIP(), []int{6}
}

type WorkflowServiceListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*WorkflowItem `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *WorkflowServiceListResponse) Reset() {
	*x = WorkflowServiceListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controlplane_v1_workflow_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowServiceListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowServiceListResponse) ProtoMessage() {}

func (x *WorkflowServiceListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controlplane_v1_workflow_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowServiceListResponse.ProtoReflect.Descriptor instead.
func (*WorkflowServiceListResponse) Descriptor() ([]byte, []int) {
	return file_controlplane_v1_workflow_proto_rawDescGZIP(), []int{7}
}

func (x *WorkflowServiceListResponse) GetResult() []*WorkflowItem {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_controlplane_v1_workflow_proto protoreflect.FileDescriptor

var file_controlplane_v1_workflow_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x27, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x01, 0x0a, 0x1c, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x21, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x61, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x86, 0x02, 0x0a, 0x1c, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x74, 0x65,
	0x61, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x04, 0x74, 0x65, 0x61, 0x6d,
	0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x48, 0x03, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x88, 0x01, 0x01,
	0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x74, 0x65, 0x61, 0x6d, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x56, 0x0a, 0x1d, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x35, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x56, 0x0a, 0x1d, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x38, 0x0a, 0x1c, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1f, 0x0a, 0x1d, 0x57, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x0a, 0x1a, 0x57,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x54, 0x0a, 0x1b, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32,
	0xaf, 0x03, 0x0a, 0x0f, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x67, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x2d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x6c, 0x6f, 0x6f, 0x70, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x6c, 0x6f, 0x6f, 0x70, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controlplane_v1_workflow_proto_rawDescOnce sync.Once
	file_controlplane_v1_workflow_proto_rawDescData = file_controlplane_v1_workflow_proto_rawDesc
)

func file_controlplane_v1_workflow_proto_rawDescGZIP() []byte {
	file_controlplane_v1_workflow_proto_rawDescOnce.Do(func() {
		file_controlplane_v1_workflow_proto_rawDescData = protoimpl.X.CompressGZIP(file_controlplane_v1_workflow_proto_rawDescData)
	})
	return file_controlplane_v1_workflow_proto_rawDescData
}

var file_controlplane_v1_workflow_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_controlplane_v1_workflow_proto_goTypes = []interface{}{
	(*WorkflowServiceCreateRequest)(nil),  // 0: controlplane.v1.WorkflowServiceCreateRequest
	(*WorkflowServiceUpdateRequest)(nil),  // 1: controlplane.v1.WorkflowServiceUpdateRequest
	(*WorkflowServiceUpdateResponse)(nil), // 2: controlplane.v1.WorkflowServiceUpdateResponse
	(*WorkflowServiceCreateResponse)(nil), // 3: controlplane.v1.WorkflowServiceCreateResponse
	(*WorkflowServiceDeleteRequest)(nil),  // 4: controlplane.v1.WorkflowServiceDeleteRequest
	(*WorkflowServiceDeleteResponse)(nil), // 5: controlplane.v1.WorkflowServiceDeleteResponse
	(*WorkflowServiceListRequest)(nil),    // 6: controlplane.v1.WorkflowServiceListRequest
	(*WorkflowServiceListResponse)(nil),   // 7: controlplane.v1.WorkflowServiceListResponse
	(*WorkflowItem)(nil),                  // 8: controlplane.v1.WorkflowItem
}
var file_controlplane_v1_workflow_proto_depIdxs = []int32{
	8, // 0: controlplane.v1.WorkflowServiceUpdateResponse.result:type_name -> controlplane.v1.WorkflowItem
	8, // 1: controlplane.v1.WorkflowServiceCreateResponse.result:type_name -> controlplane.v1.WorkflowItem
	8, // 2: controlplane.v1.WorkflowServiceListResponse.result:type_name -> controlplane.v1.WorkflowItem
	0, // 3: controlplane.v1.WorkflowService.Create:input_type -> controlplane.v1.WorkflowServiceCreateRequest
	1, // 4: controlplane.v1.WorkflowService.Update:input_type -> controlplane.v1.WorkflowServiceUpdateRequest
	6, // 5: controlplane.v1.WorkflowService.List:input_type -> controlplane.v1.WorkflowServiceListRequest
	4, // 6: controlplane.v1.WorkflowService.Delete:input_type -> controlplane.v1.WorkflowServiceDeleteRequest
	3, // 7: controlplane.v1.WorkflowService.Create:output_type -> controlplane.v1.WorkflowServiceCreateResponse
	2, // 8: controlplane.v1.WorkflowService.Update:output_type -> controlplane.v1.WorkflowServiceUpdateResponse
	7, // 9: controlplane.v1.WorkflowService.List:output_type -> controlplane.v1.WorkflowServiceListResponse
	5, // 10: controlplane.v1.WorkflowService.Delete:output_type -> controlplane.v1.WorkflowServiceDeleteResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_controlplane_v1_workflow_proto_init() }
func file_controlplane_v1_workflow_proto_init() {
	if File_controlplane_v1_workflow_proto != nil {
		return
	}
	file_controlplane_v1_response_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_controlplane_v1_workflow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowServiceCreateRequest); i {
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
		file_controlplane_v1_workflow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowServiceUpdateRequest); i {
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
		file_controlplane_v1_workflow_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowServiceUpdateResponse); i {
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
		file_controlplane_v1_workflow_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowServiceCreateResponse); i {
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
		file_controlplane_v1_workflow_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowServiceDeleteRequest); i {
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
		file_controlplane_v1_workflow_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowServiceDeleteResponse); i {
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
		file_controlplane_v1_workflow_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowServiceListRequest); i {
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
		file_controlplane_v1_workflow_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowServiceListResponse); i {
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
	file_controlplane_v1_workflow_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_controlplane_v1_workflow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_controlplane_v1_workflow_proto_goTypes,
		DependencyIndexes: file_controlplane_v1_workflow_proto_depIdxs,
		MessageInfos:      file_controlplane_v1_workflow_proto_msgTypes,
	}.Build()
	File_controlplane_v1_workflow_proto = out.File
	file_controlplane_v1_workflow_proto_rawDesc = nil
	file_controlplane_v1_workflow_proto_goTypes = nil
	file_controlplane_v1_workflow_proto_depIdxs = nil
}
