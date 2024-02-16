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
// source: workflowcontract/v1/crafting_schema.proto

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

type CraftingSchema_Runner_RunnerType int32

const (
	CraftingSchema_Runner_RUNNER_TYPE_UNSPECIFIED CraftingSchema_Runner_RunnerType = 0
	CraftingSchema_Runner_GITHUB_ACTION           CraftingSchema_Runner_RunnerType = 1
	CraftingSchema_Runner_GITLAB_PIPELINE         CraftingSchema_Runner_RunnerType = 2
	CraftingSchema_Runner_AZURE_PIPELINE          CraftingSchema_Runner_RunnerType = 3
	CraftingSchema_Runner_JENKINS_JOB             CraftingSchema_Runner_RunnerType = 4
	CraftingSchema_Runner_CIRCLECI_BUILD          CraftingSchema_Runner_RunnerType = 5
	CraftingSchema_Runner_DAGGER_PIPELINE         CraftingSchema_Runner_RunnerType = 6
)

// Enum value maps for CraftingSchema_Runner_RunnerType.
var (
	CraftingSchema_Runner_RunnerType_name = map[int32]string{
		0: "RUNNER_TYPE_UNSPECIFIED",
		1: "GITHUB_ACTION",
		2: "GITLAB_PIPELINE",
		3: "AZURE_PIPELINE",
		4: "JENKINS_JOB",
		5: "CIRCLECI_BUILD",
		6: "DAGGER_PIPELINE",
	}
	CraftingSchema_Runner_RunnerType_value = map[string]int32{
		"RUNNER_TYPE_UNSPECIFIED": 0,
		"GITHUB_ACTION":           1,
		"GITLAB_PIPELINE":         2,
		"AZURE_PIPELINE":          3,
		"JENKINS_JOB":             4,
		"CIRCLECI_BUILD":          5,
		"DAGGER_PIPELINE":         6,
	}
)

func (x CraftingSchema_Runner_RunnerType) Enum() *CraftingSchema_Runner_RunnerType {
	p := new(CraftingSchema_Runner_RunnerType)
	*p = x
	return p
}

func (x CraftingSchema_Runner_RunnerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CraftingSchema_Runner_RunnerType) Descriptor() protoreflect.EnumDescriptor {
	return file_workflowcontract_v1_crafting_schema_proto_enumTypes[0].Descriptor()
}

func (CraftingSchema_Runner_RunnerType) Type() protoreflect.EnumType {
	return &file_workflowcontract_v1_crafting_schema_proto_enumTypes[0]
}

func (x CraftingSchema_Runner_RunnerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CraftingSchema_Runner_RunnerType.Descriptor instead.
func (CraftingSchema_Runner_RunnerType) EnumDescriptor() ([]byte, []int) {
	return file_workflowcontract_v1_crafting_schema_proto_rawDescGZIP(), []int{0, 0, 0}
}

type CraftingSchema_Material_MaterialType int32

const (
	CraftingSchema_Material_MATERIAL_TYPE_UNSPECIFIED CraftingSchema_Material_MaterialType = 0
	CraftingSchema_Material_STRING                    CraftingSchema_Material_MaterialType = 1
	CraftingSchema_Material_CONTAINER_IMAGE           CraftingSchema_Material_MaterialType = 2
	CraftingSchema_Material_ARTIFACT                  CraftingSchema_Material_MaterialType = 3
	CraftingSchema_Material_SBOM_CYCLONEDX_JSON       CraftingSchema_Material_MaterialType = 4
	CraftingSchema_Material_SBOM_SPDX_JSON            CraftingSchema_Material_MaterialType = 5
	CraftingSchema_Material_JUNIT_XML                 CraftingSchema_Material_MaterialType = 6
	// https://github.com/openvex/spec
	CraftingSchema_Material_OPENVEX CraftingSchema_Material_MaterialType = 7
	// https://docs.oasis-open.org/csaf/csaf/v2.0/cs03/csaf-v2.0-cs03.html
	CraftingSchema_Material_CSAF_VEX CraftingSchema_Material_MaterialType = 8
	// Static analysis output format
	// https://github.com/microsoft/sarif-tutorials/blob/main/docs/1-Introduction.md
	CraftingSchema_Material_SARIF CraftingSchema_Material_MaterialType = 9
)

// Enum value maps for CraftingSchema_Material_MaterialType.
var (
	CraftingSchema_Material_MaterialType_name = map[int32]string{
		0: "MATERIAL_TYPE_UNSPECIFIED",
		1: "STRING",
		2: "CONTAINER_IMAGE",
		3: "ARTIFACT",
		4: "SBOM_CYCLONEDX_JSON",
		5: "SBOM_SPDX_JSON",
		6: "JUNIT_XML",
		7: "OPENVEX",
		8: "CSAF_VEX",
		9: "SARIF",
	}
	CraftingSchema_Material_MaterialType_value = map[string]int32{
		"MATERIAL_TYPE_UNSPECIFIED": 0,
		"STRING":                    1,
		"CONTAINER_IMAGE":           2,
		"ARTIFACT":                  3,
		"SBOM_CYCLONEDX_JSON":       4,
		"SBOM_SPDX_JSON":            5,
		"JUNIT_XML":                 6,
		"OPENVEX":                   7,
		"CSAF_VEX":                  8,
		"SARIF":                     9,
	}
)

func (x CraftingSchema_Material_MaterialType) Enum() *CraftingSchema_Material_MaterialType {
	p := new(CraftingSchema_Material_MaterialType)
	*p = x
	return p
}

func (x CraftingSchema_Material_MaterialType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CraftingSchema_Material_MaterialType) Descriptor() protoreflect.EnumDescriptor {
	return file_workflowcontract_v1_crafting_schema_proto_enumTypes[1].Descriptor()
}

func (CraftingSchema_Material_MaterialType) Type() protoreflect.EnumType {
	return &file_workflowcontract_v1_crafting_schema_proto_enumTypes[1]
}

func (x CraftingSchema_Material_MaterialType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CraftingSchema_Material_MaterialType.Descriptor instead.
func (CraftingSchema_Material_MaterialType) EnumDescriptor() ([]byte, []int) {
	return file_workflowcontract_v1_crafting_schema_proto_rawDescGZIP(), []int{0, 1, 0}
}

// Schema definition provided by the user to the tool
// that defines the schema of the workflowRun
type CraftingSchema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Version of the schema, do not confuse with the revision of the content
	SchemaVersion string                     `protobuf:"bytes,1,opt,name=schema_version,json=schemaVersion,proto3" json:"schema_version,omitempty"`
	Materials     []*CraftingSchema_Material `protobuf:"bytes,2,rep,name=materials,proto3" json:"materials,omitempty"`
	EnvAllowList  []string                   `protobuf:"bytes,3,rep,name=env_allow_list,json=envAllowList,proto3" json:"env_allow_list,omitempty"`
	Runner        *CraftingSchema_Runner     `protobuf:"bytes,4,opt,name=runner,proto3" json:"runner,omitempty"`
	// List of annotations that can be used to add metadata to the attestation
	// this metadata can be used later on by the integrations engine to filter and interpolate data
	// It works in addition to the annotations defined in the materials and the runner
	Annotations []*Annotation `protobuf:"bytes,5,rep,name=annotations,proto3" json:"annotations,omitempty"`
}

func (x *CraftingSchema) Reset() {
	*x = CraftingSchema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflowcontract_v1_crafting_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CraftingSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CraftingSchema) ProtoMessage() {}

func (x *CraftingSchema) ProtoReflect() protoreflect.Message {
	mi := &file_workflowcontract_v1_crafting_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CraftingSchema.ProtoReflect.Descriptor instead.
func (*CraftingSchema) Descriptor() ([]byte, []int) {
	return file_workflowcontract_v1_crafting_schema_proto_rawDescGZIP(), []int{0}
}

func (x *CraftingSchema) GetSchemaVersion() string {
	if x != nil {
		return x.SchemaVersion
	}
	return ""
}

func (x *CraftingSchema) GetMaterials() []*CraftingSchema_Material {
	if x != nil {
		return x.Materials
	}
	return nil
}

func (x *CraftingSchema) GetEnvAllowList() []string {
	if x != nil {
		return x.EnvAllowList
	}
	return nil
}

func (x *CraftingSchema) GetRunner() *CraftingSchema_Runner {
	if x != nil {
		return x.Runner
	}
	return nil
}

func (x *CraftingSchema) GetAnnotations() []*Annotation {
	if x != nil {
		return x.Annotations
	}
	return nil
}

type Annotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // Single word optionally separated with _
	// This value can be set in the contract or provided during the attestation
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Annotation) Reset() {
	*x = Annotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflowcontract_v1_crafting_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Annotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Annotation) ProtoMessage() {}

func (x *Annotation) ProtoReflect() protoreflect.Message {
	mi := &file_workflowcontract_v1_crafting_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Annotation.ProtoReflect.Descriptor instead.
func (*Annotation) Descriptor() ([]byte, []int) {
	return file_workflowcontract_v1_crafting_schema_proto_rawDescGZIP(), []int{1}
}

func (x *Annotation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Annotation) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type CraftingSchema_Runner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type CraftingSchema_Runner_RunnerType `protobuf:"varint,1,opt,name=type,proto3,enum=workflowcontract.v1.CraftingSchema_Runner_RunnerType" json:"type,omitempty"`
}

func (x *CraftingSchema_Runner) Reset() {
	*x = CraftingSchema_Runner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflowcontract_v1_crafting_schema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CraftingSchema_Runner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CraftingSchema_Runner) ProtoMessage() {}

func (x *CraftingSchema_Runner) ProtoReflect() protoreflect.Message {
	mi := &file_workflowcontract_v1_crafting_schema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CraftingSchema_Runner.ProtoReflect.Descriptor instead.
func (*CraftingSchema_Runner) Descriptor() ([]byte, []int) {
	return file_workflowcontract_v1_crafting_schema_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CraftingSchema_Runner) GetType() CraftingSchema_Runner_RunnerType {
	if x != nil {
		return x.Type
	}
	return CraftingSchema_Runner_RUNNER_TYPE_UNSPECIFIED
}

type CraftingSchema_Material struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     CraftingSchema_Material_MaterialType `protobuf:"varint,1,opt,name=type,proto3,enum=workflowcontract.v1.CraftingSchema_Material_MaterialType" json:"type,omitempty"`
	Name     string                               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // Single word optionally separated with _ or -
	Optional bool                                 `protobuf:"varint,3,opt,name=optional,proto3" json:"optional,omitempty"`
	// If a material is set as output it will get added to the subject in the statement
	Output bool `protobuf:"varint,4,opt,name=output,proto3" json:"output,omitempty"`
	// List of annotations that can be used to add metadata to the material
	// this metadata can be used later on by the integrations engine to filter and interpolate data
	Annotations []*Annotation `protobuf:"bytes,5,rep,name=annotations,proto3" json:"annotations,omitempty"`
}

func (x *CraftingSchema_Material) Reset() {
	*x = CraftingSchema_Material{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflowcontract_v1_crafting_schema_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CraftingSchema_Material) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CraftingSchema_Material) ProtoMessage() {}

func (x *CraftingSchema_Material) ProtoReflect() protoreflect.Message {
	mi := &file_workflowcontract_v1_crafting_schema_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CraftingSchema_Material.ProtoReflect.Descriptor instead.
func (*CraftingSchema_Material) Descriptor() ([]byte, []int) {
	return file_workflowcontract_v1_crafting_schema_proto_rawDescGZIP(), []int{0, 1}
}

func (x *CraftingSchema_Material) GetType() CraftingSchema_Material_MaterialType {
	if x != nil {
		return x.Type
	}
	return CraftingSchema_Material_MATERIAL_TYPE_UNSPECIFIED
}

func (x *CraftingSchema_Material) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CraftingSchema_Material) GetOptional() bool {
	if x != nil {
		return x.Optional
	}
	return false
}

func (x *CraftingSchema_Material) GetOutput() bool {
	if x != nil {
		return x.Output
	}
	return false
}

func (x *CraftingSchema_Material) GetAnnotations() []*Annotation {
	if x != nil {
		return x.Annotations
	}
	return nil
}

var File_workflowcontract_v1_crafting_schema_proto protoreflect.FileDescriptor

var file_workflowcontract_v1_crafting_schema_proto_rawDesc = []byte{
	0x0a, 0x29, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x61, 0x66, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x08, 0x0a, 0x0e, 0x43, 0x72,
	0x61, 0x66, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x30, 0x0a, 0x0e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x0a, 0x02, 0x76, 0x31, 0x52,
	0x0d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4a,
	0x0a, 0x09, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2c, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x61, 0x66, 0x74, 0x69, 0x6e, 0x67,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x52,
	0x09, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x65, 0x6e,
	0x76, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0c, 0x65, 0x6e, 0x76, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x42, 0x0a, 0x06, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x61, 0x66, 0x74, 0x69, 0x6e, 0x67, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x06, 0x72, 0x75,
	0x6e, 0x6e, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0xff, 0x01, 0x0a, 0x06, 0x52, 0x75, 0x6e, 0x6e,
	0x65, 0x72, 0x12, 0x53, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x35, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x61, 0x66, 0x74, 0x69, 0x6e, 0x67, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x52, 0x75, 0x6e,
	0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x20,
	0x00, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x9f, 0x01, 0x0a, 0x0a, 0x52, 0x75, 0x6e, 0x6e,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x52, 0x55, 0x4e, 0x4e, 0x45, 0x52,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x47, 0x49, 0x54, 0x48, 0x55, 0x42, 0x5f, 0x41, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x47, 0x49, 0x54, 0x4c, 0x41, 0x42,
	0x5f, 0x50, 0x49, 0x50, 0x45, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x41,
	0x5a, 0x55, 0x52, 0x45, 0x5f, 0x50, 0x49, 0x50, 0x45, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x03, 0x12,
	0x0f, 0x0a, 0x0b, 0x4a, 0x45, 0x4e, 0x4b, 0x49, 0x4e, 0x53, 0x5f, 0x4a, 0x4f, 0x42, 0x10, 0x04,
	0x12, 0x12, 0x0a, 0x0e, 0x43, 0x49, 0x52, 0x43, 0x4c, 0x45, 0x43, 0x49, 0x5f, 0x42, 0x55, 0x49,
	0x4c, 0x44, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x41, 0x47, 0x47, 0x45, 0x52, 0x5f, 0x50,
	0x49, 0x50, 0x45, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x06, 0x1a, 0xc1, 0x03, 0x0a, 0x08, 0x4d, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x57, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x39, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x61, 0x66, 0x74,
	0x69, 0x6e, 0x67, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x20, 0x00, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x24, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xfa,
	0x42, 0x0d, 0x72, 0x0b, 0x32, 0x09, 0x5e, 0x5b, 0x5c, 0x77, 0x7c, 0x2d, 0x5d, 0x2b, 0x24, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x41, 0x0a, 0x0b, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xbe, 0x01, 0x0a,
	0x0c, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a,
	0x19, 0x4d, 0x41, 0x54, 0x45, 0x52, 0x49, 0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4f, 0x4e, 0x54,
	0x41, 0x49, 0x4e, 0x45, 0x52, 0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x02, 0x12, 0x0c, 0x0a,
	0x08, 0x41, 0x52, 0x54, 0x49, 0x46, 0x41, 0x43, 0x54, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x53,
	0x42, 0x4f, 0x4d, 0x5f, 0x43, 0x59, 0x43, 0x4c, 0x4f, 0x4e, 0x45, 0x44, 0x58, 0x5f, 0x4a, 0x53,
	0x4f, 0x4e, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x42, 0x4f, 0x4d, 0x5f, 0x53, 0x50, 0x44,
	0x58, 0x5f, 0x4a, 0x53, 0x4f, 0x4e, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x4a, 0x55, 0x4e, 0x49,
	0x54, 0x5f, 0x58, 0x4d, 0x4c, 0x10, 0x06, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x50, 0x45, 0x4e, 0x56,
	0x45, 0x58, 0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x53, 0x41, 0x46, 0x5f, 0x56, 0x45, 0x58,
	0x10, 0x08, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x41, 0x52, 0x49, 0x46, 0x10, 0x09, 0x22, 0x46, 0x0a,
	0x0a, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xfa, 0x42, 0x0b, 0x72, 0x09,
	0x32, 0x07, 0x5e, 0x5b, 0x5c, 0x77, 0x5d, 0x2b, 0x24, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x4d, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x6c, 0x6f, 0x6f, 0x70, 0x2d, 0x64, 0x65,
	0x76, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x6c, 0x6f, 0x6f, 0x70, 0x2f, 0x61, 0x70, 0x70, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_workflowcontract_v1_crafting_schema_proto_rawDescOnce sync.Once
	file_workflowcontract_v1_crafting_schema_proto_rawDescData = file_workflowcontract_v1_crafting_schema_proto_rawDesc
)

func file_workflowcontract_v1_crafting_schema_proto_rawDescGZIP() []byte {
	file_workflowcontract_v1_crafting_schema_proto_rawDescOnce.Do(func() {
		file_workflowcontract_v1_crafting_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_workflowcontract_v1_crafting_schema_proto_rawDescData)
	})
	return file_workflowcontract_v1_crafting_schema_proto_rawDescData
}

var file_workflowcontract_v1_crafting_schema_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_workflowcontract_v1_crafting_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_workflowcontract_v1_crafting_schema_proto_goTypes = []interface{}{
	(CraftingSchema_Runner_RunnerType)(0),     // 0: workflowcontract.v1.CraftingSchema.Runner.RunnerType
	(CraftingSchema_Material_MaterialType)(0), // 1: workflowcontract.v1.CraftingSchema.Material.MaterialType
	(*CraftingSchema)(nil),                    // 2: workflowcontract.v1.CraftingSchema
	(*Annotation)(nil),                        // 3: workflowcontract.v1.Annotation
	(*CraftingSchema_Runner)(nil),             // 4: workflowcontract.v1.CraftingSchema.Runner
	(*CraftingSchema_Material)(nil),           // 5: workflowcontract.v1.CraftingSchema.Material
}
var file_workflowcontract_v1_crafting_schema_proto_depIdxs = []int32{
	5, // 0: workflowcontract.v1.CraftingSchema.materials:type_name -> workflowcontract.v1.CraftingSchema.Material
	4, // 1: workflowcontract.v1.CraftingSchema.runner:type_name -> workflowcontract.v1.CraftingSchema.Runner
	3, // 2: workflowcontract.v1.CraftingSchema.annotations:type_name -> workflowcontract.v1.Annotation
	0, // 3: workflowcontract.v1.CraftingSchema.Runner.type:type_name -> workflowcontract.v1.CraftingSchema.Runner.RunnerType
	1, // 4: workflowcontract.v1.CraftingSchema.Material.type:type_name -> workflowcontract.v1.CraftingSchema.Material.MaterialType
	3, // 5: workflowcontract.v1.CraftingSchema.Material.annotations:type_name -> workflowcontract.v1.Annotation
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_workflowcontract_v1_crafting_schema_proto_init() }
func file_workflowcontract_v1_crafting_schema_proto_init() {
	if File_workflowcontract_v1_crafting_schema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_workflowcontract_v1_crafting_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CraftingSchema); i {
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
		file_workflowcontract_v1_crafting_schema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Annotation); i {
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
		file_workflowcontract_v1_crafting_schema_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CraftingSchema_Runner); i {
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
		file_workflowcontract_v1_crafting_schema_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CraftingSchema_Material); i {
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
			RawDescriptor: file_workflowcontract_v1_crafting_schema_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_workflowcontract_v1_crafting_schema_proto_goTypes,
		DependencyIndexes: file_workflowcontract_v1_crafting_schema_proto_depIdxs,
		EnumInfos:         file_workflowcontract_v1_crafting_schema_proto_enumTypes,
		MessageInfos:      file_workflowcontract_v1_crafting_schema_proto_msgTypes,
	}.Build()
	File_workflowcontract_v1_crafting_schema_proto = out.File
	file_workflowcontract_v1_crafting_schema_proto_rawDesc = nil
	file_workflowcontract_v1_crafting_schema_proto_goTypes = nil
	file_workflowcontract_v1_crafting_schema_proto_depIdxs = nil
}
