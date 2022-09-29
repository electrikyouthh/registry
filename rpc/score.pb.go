// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.3
// source: google/cloud/apigeeregistry/v1/scoring/score.proto

// (-- api-linter: core::0215::versioned-packages=disabled
//     aip.dev/not-precedent: Support protos for the apigeeregistry.v1 API. --)

package rpc

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Stores the score for a resource.
// Stored as an artifact against the resource whose score it represents.
type Score struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Artifact identifier. This will be auto-generated based on the id of the
	// ScoreDefinition used to calculate this.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Artifact kind. May be used in YAML representations to identify the type of
	// this artifact.
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// A human-friendly name for the score (populated from ScoreDefinition).
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// A more detailed description of the score (populated from ScoreDefinition).
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// A Uri which points to an explanation of the formula used here.
	Uri string `protobuf:"bytes,5,opt,name=uri,proto3" json:"uri,omitempty"`
	// A human-friendly name to display the uri.
	UriDisplayName string `protobuf:"bytes,6,opt,name=uri_display_name,json=uriDisplayName,proto3" json:"uri_display_name,omitempty"`
	// Full resource name of the ScoreDefinition artifact which was used
	// to generate this score.
	DefinitionName string `protobuf:"bytes,7,opt,name=definition_name,json=definitionName,proto3" json:"definition_name,omitempty"`
	// Stores the severity associated with the score value.
	Severity Severity `protobuf:"varint,8,opt,name=severity,proto3,enum=google.cloud.apigeeregistry.v1.scoring.Severity" json:"severity,omitempty"`
	// Stores the actual score of a particular resource.
	//
	// Types that are assignable to Value:
	//	*Score_PercentValue
	//	*Score_IntegerValue
	//	*Score_BooleanValue
	Value isScore_Value `protobuf_oneof:"value"`
}

func (x *Score) Reset() {
	*x = Score{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_apigeeregistry_v1_scoring_score_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Score) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Score) ProtoMessage() {}

func (x *Score) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_apigeeregistry_v1_scoring_score_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Score.ProtoReflect.Descriptor instead.
func (*Score) Descriptor() ([]byte, []int) {
	return file_google_cloud_apigeeregistry_v1_scoring_score_proto_rawDescGZIP(), []int{0}
}

func (x *Score) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Score) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Score) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Score) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Score) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *Score) GetUriDisplayName() string {
	if x != nil {
		return x.UriDisplayName
	}
	return ""
}

func (x *Score) GetDefinitionName() string {
	if x != nil {
		return x.DefinitionName
	}
	return ""
}

func (x *Score) GetSeverity() Severity {
	if x != nil {
		return x.Severity
	}
	return Severity_SEVERITY_UNSPECIFIED
}

func (m *Score) GetValue() isScore_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Score) GetPercentValue() *PercentValue {
	if x, ok := x.GetValue().(*Score_PercentValue); ok {
		return x.PercentValue
	}
	return nil
}

func (x *Score) GetIntegerValue() *IntegerValue {
	if x, ok := x.GetValue().(*Score_IntegerValue); ok {
		return x.IntegerValue
	}
	return nil
}

func (x *Score) GetBooleanValue() *BooleanValue {
	if x, ok := x.GetValue().(*Score_BooleanValue); ok {
		return x.BooleanValue
	}
	return nil
}

type isScore_Value interface {
	isScore_Value()
}

type Score_PercentValue struct {
	// This is set if the score is a percentage.
	PercentValue *PercentValue `protobuf:"bytes,9,opt,name=percent_value,json=percentValue,proto3,oneof"`
}

type Score_IntegerValue struct {
	// This is set if the score is an integer.
	IntegerValue *IntegerValue `protobuf:"bytes,10,opt,name=integer_value,json=integerValue,proto3,oneof"`
}

type Score_BooleanValue struct {
	// This is set if the score is a boolean.
	BooleanValue *BooleanValue `protobuf:"bytes,11,opt,name=boolean_value,json=booleanValue,proto3,oneof"`
}

func (*Score_PercentValue) isScore_Value() {}

func (*Score_IntegerValue) isScore_Value() {}

func (*Score_BooleanValue) isScore_Value() {}

// Represents the score which is a percentage.
type PercentValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Stores the value of score.
	Value float32 `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PercentValue) Reset() {
	*x = PercentValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_apigeeregistry_v1_scoring_score_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PercentValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PercentValue) ProtoMessage() {}

func (x *PercentValue) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_apigeeregistry_v1_scoring_score_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PercentValue.ProtoReflect.Descriptor instead.
func (*PercentValue) Descriptor() ([]byte, []int) {
	return file_google_cloud_apigeeregistry_v1_scoring_score_proto_rawDescGZIP(), []int{1}
}

func (x *PercentValue) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

// Represents the score which is a percentage.
type IntegerValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Stores the value of score.
	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	// The following fields will be used to display scores in the UI like "5/10".
	// Stores the minimum value this score can take.
	MinValue int32 `protobuf:"varint,2,opt,name=min_value,json=minValue,proto3" json:"min_value,omitempty"`
	// Stores the minimum value this score can take.
	MaxValue int32 `protobuf:"varint,3,opt,name=max_value,json=maxValue,proto3" json:"max_value,omitempty"`
}

func (x *IntegerValue) Reset() {
	*x = IntegerValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_apigeeregistry_v1_scoring_score_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntegerValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegerValue) ProtoMessage() {}

func (x *IntegerValue) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_apigeeregistry_v1_scoring_score_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntegerValue.ProtoReflect.Descriptor instead.
func (*IntegerValue) Descriptor() ([]byte, []int) {
	return file_google_cloud_apigeeregistry_v1_scoring_score_proto_rawDescGZIP(), []int{2}
}

func (x *IntegerValue) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *IntegerValue) GetMinValue() int32 {
	if x != nil {
		return x.MinValue
	}
	return 0
}

func (x *IntegerValue) GetMaxValue() int32 {
	if x != nil {
		return x.MaxValue
	}
	return 0
}

// Represents the score which is a percentage.
type BooleanValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Stores the value of score.
	Value bool `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	// Stores the value which should be used to display in the UI.
	// This is derived from "display_true" and "display_false"
	// fields of BooleanType.
	DisplayValue string `protobuf:"bytes,2,opt,name=display_value,json=displayValue,proto3" json:"display_value,omitempty"`
}

func (x *BooleanValue) Reset() {
	*x = BooleanValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_apigeeregistry_v1_scoring_score_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BooleanValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BooleanValue) ProtoMessage() {}

func (x *BooleanValue) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_apigeeregistry_v1_scoring_score_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BooleanValue.ProtoReflect.Descriptor instead.
func (*BooleanValue) Descriptor() ([]byte, []int) {
	return file_google_cloud_apigeeregistry_v1_scoring_score_proto_rawDescGZIP(), []int{3}
}

func (x *BooleanValue) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

func (x *BooleanValue) GetDisplayValue() string {
	if x != nil {
		return x.DisplayValue
	}
	return ""
}

var File_google_cloud_apigeeregistry_v1_scoring_score_proto protoreflect.FileDescriptor

var file_google_cloud_apigeeregistry_v1_scoring_score_proto_rawDesc = []byte{
	0x0a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x70, 0x69, 0x67, 0x65, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x67, 0x65, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x67,
	0x65, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x04, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x13,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x69, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x28,
	0x0a, 0x10, 0x75, 0x72, 0x69, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x75, 0x72, 0x69, 0x44, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x0f, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4c, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69,
	0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x67, 0x65, 0x65, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x2e, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65,
	0x72, 0x69, 0x74, 0x79, 0x12, 0x5b, 0x0a, 0x0d, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x67, 0x65,
	0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x63, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x48, 0x00, 0x52, 0x0c, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x5b, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x67, 0x65, 0x65, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00,
	0x52, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x5b,
	0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x67, 0x65, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x62,
	0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x29, 0x0a, 0x0c, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x63, 0x0a, 0x0c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x19, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69,
	0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d,
	0x69, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x4e, 0x0a, 0x0c, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x65, 0x0a, 0x2a, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x67, 0x65, 0x65, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x42, 0x11, 0x53, 0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x67, 0x65, 0x65, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2f, 0x72, 0x70, 0x63, 0x3b, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_apigeeregistry_v1_scoring_score_proto_rawDescOnce sync.Once
	file_google_cloud_apigeeregistry_v1_scoring_score_proto_rawDescData = file_google_cloud_apigeeregistry_v1_scoring_score_proto_rawDesc
)

func file_google_cloud_apigeeregistry_v1_scoring_score_proto_rawDescGZIP() []byte {
	file_google_cloud_apigeeregistry_v1_scoring_score_proto_rawDescOnce.Do(func() {
		file_google_cloud_apigeeregistry_v1_scoring_score_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_apigeeregistry_v1_scoring_score_proto_rawDescData)
	})
	return file_google_cloud_apigeeregistry_v1_scoring_score_proto_rawDescData
}

var file_google_cloud_apigeeregistry_v1_scoring_score_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_apigeeregistry_v1_scoring_score_proto_goTypes = []interface{}{
	(*Score)(nil),        // 0: google.cloud.apigeeregistry.v1.scoring.Score
	(*PercentValue)(nil), // 1: google.cloud.apigeeregistry.v1.scoring.PercentValue
	(*IntegerValue)(nil), // 2: google.cloud.apigeeregistry.v1.scoring.IntegerValue
	(*BooleanValue)(nil), // 3: google.cloud.apigeeregistry.v1.scoring.BooleanValue
	(Severity)(0),        // 4: google.cloud.apigeeregistry.v1.scoring.Severity
}
var file_google_cloud_apigeeregistry_v1_scoring_score_proto_depIdxs = []int32{
	4, // 0: google.cloud.apigeeregistry.v1.scoring.Score.severity:type_name -> google.cloud.apigeeregistry.v1.scoring.Severity
	1, // 1: google.cloud.apigeeregistry.v1.scoring.Score.percent_value:type_name -> google.cloud.apigeeregistry.v1.scoring.PercentValue
	2, // 2: google.cloud.apigeeregistry.v1.scoring.Score.integer_value:type_name -> google.cloud.apigeeregistry.v1.scoring.IntegerValue
	3, // 3: google.cloud.apigeeregistry.v1.scoring.Score.boolean_value:type_name -> google.cloud.apigeeregistry.v1.scoring.BooleanValue
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_apigeeregistry_v1_scoring_score_proto_init() }
func file_google_cloud_apigeeregistry_v1_scoring_score_proto_init() {
	if File_google_cloud_apigeeregistry_v1_scoring_score_proto != nil {
		return
	}
	file_google_cloud_apigeeregistry_v1_scoring_severity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_apigeeregistry_v1_scoring_score_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Score); i {
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
		file_google_cloud_apigeeregistry_v1_scoring_score_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PercentValue); i {
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
		file_google_cloud_apigeeregistry_v1_scoring_score_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntegerValue); i {
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
		file_google_cloud_apigeeregistry_v1_scoring_score_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BooleanValue); i {
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
	file_google_cloud_apigeeregistry_v1_scoring_score_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Score_PercentValue)(nil),
		(*Score_IntegerValue)(nil),
		(*Score_BooleanValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_apigeeregistry_v1_scoring_score_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_apigeeregistry_v1_scoring_score_proto_goTypes,
		DependencyIndexes: file_google_cloud_apigeeregistry_v1_scoring_score_proto_depIdxs,
		MessageInfos:      file_google_cloud_apigeeregistry_v1_scoring_score_proto_msgTypes,
	}.Build()
	File_google_cloud_apigeeregistry_v1_scoring_score_proto = out.File
	file_google_cloud_apigeeregistry_v1_scoring_score_proto_rawDesc = nil
	file_google_cloud_apigeeregistry_v1_scoring_score_proto_goTypes = nil
	file_google_cloud_apigeeregistry_v1_scoring_score_proto_depIdxs = nil
}
