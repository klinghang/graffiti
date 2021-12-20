// Copyright 2021 Google LLC
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
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.2
// source: google/cloud/aiplatform/v1beta1/lineage_subgraph.proto

package aiplatform

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// A subgraph of the overall lineage graph. Event edges connect Artifact and
// Execution nodes.
type LineageSubgraph struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Artifact nodes in the subgraph.
	Artifacts []*Artifact `protobuf:"bytes,1,rep,name=artifacts,proto3" json:"artifacts,omitempty"`
	// The Execution nodes in the subgraph.
	Executions []*Execution `protobuf:"bytes,2,rep,name=executions,proto3" json:"executions,omitempty"`
	// The Event edges between Artifacts and Executions in the subgraph.
	Events []*Event `protobuf:"bytes,3,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *LineageSubgraph) Reset() {
	*x = LineageSubgraph{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_lineage_subgraph_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LineageSubgraph) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineageSubgraph) ProtoMessage() {}

func (x *LineageSubgraph) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_lineage_subgraph_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LineageSubgraph.ProtoReflect.Descriptor instead.
func (*LineageSubgraph) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_lineage_subgraph_proto_rawDescGZIP(), []int{0}
}

func (x *LineageSubgraph) GetArtifacts() []*Artifact {
	if x != nil {
		return x.Artifacts
	}
	return nil
}

func (x *LineageSubgraph) GetExecutions() []*Execution {
	if x != nil {
		return x.Executions
	}
	return nil
}

func (x *LineageSubgraph) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

var File_google_cloud_aiplatform_v1beta1_lineage_subgraph_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1beta1_lineage_subgraph_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x6c, 0x69, 0x6e, 0x65, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x75, 0x62, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb9, 0x02, 0x0a, 0x0f, 0x4c, 0x69, 0x6e, 0x65, 0x61, 0x67, 0x65, 0x53, 0x75, 0x62, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x12, 0x70, 0x0a, 0x09, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x42, 0x27, 0xfa, 0x41, 0x24, 0x0a, 0x22, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x09, 0x61, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x12, 0x74, 0x0a, 0x0a, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x28, 0xfa, 0x41, 0x25, 0x0a, 0x23, 0x61, 0x69, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0a, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3e, 0x0a, 0x06,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x88, 0x01, 0x0a,
	0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x42, 0x14, 0x4c, 0x69, 0x6e, 0x65, 0x61, 0x67, 0x65, 0x53, 0x75, 0x62,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x61, 0x69, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1beta1_lineage_subgraph_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1beta1_lineage_subgraph_proto_rawDescData = file_google_cloud_aiplatform_v1beta1_lineage_subgraph_proto_rawDesc
)

func file_google_cloud_aiplatform_v1beta1_lineage_subgraph_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1beta1_lineage_subgraph_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1beta1_lineage_subgraph_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1beta1_lineage_subgraph_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1beta1_lineage_subgraph_proto_rawDescData
}

var file_google_cloud_aiplatform_v1beta1_lineage_subgraph_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_aiplatform_v1beta1_lineage_subgraph_proto_goTypes = []interface{}{
	(*LineageSubgraph)(nil), // 0: google.cloud.aiplatform.v1beta1.LineageSubgraph
	(*Artifact)(nil),        // 1: google.cloud.aiplatform.v1beta1.Artifact
	(*Execution)(nil),       // 2: google.cloud.aiplatform.v1beta1.Execution
	(*Event)(nil),           // 3: google.cloud.aiplatform.v1beta1.Event
}
var file_google_cloud_aiplatform_v1beta1_lineage_subgraph_proto_depIdxs = []int32{
	1, // 0: google.cloud.aiplatform.v1beta1.LineageSubgraph.artifacts:type_name -> google.cloud.aiplatform.v1beta1.Artifact
	2, // 1: google.cloud.aiplatform.v1beta1.LineageSubgraph.executions:type_name -> google.cloud.aiplatform.v1beta1.Execution
	3, // 2: google.cloud.aiplatform.v1beta1.LineageSubgraph.events:type_name -> google.cloud.aiplatform.v1beta1.Event
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1beta1_lineage_subgraph_proto_init() }
func file_google_cloud_aiplatform_v1beta1_lineage_subgraph_proto_init() {
	if File_google_cloud_aiplatform_v1beta1_lineage_subgraph_proto != nil {
		return
	}
	file_google_cloud_aiplatform_v1beta1_artifact_proto_init()
	file_google_cloud_aiplatform_v1beta1_event_proto_init()
	file_google_cloud_aiplatform_v1beta1_execution_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1beta1_lineage_subgraph_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LineageSubgraph); i {
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
			RawDescriptor: file_google_cloud_aiplatform_v1beta1_lineage_subgraph_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1beta1_lineage_subgraph_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1beta1_lineage_subgraph_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1beta1_lineage_subgraph_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1beta1_lineage_subgraph_proto = out.File
	file_google_cloud_aiplatform_v1beta1_lineage_subgraph_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1beta1_lineage_subgraph_proto_goTypes = nil
	file_google_cloud_aiplatform_v1beta1_lineage_subgraph_proto_depIdxs = nil
}
