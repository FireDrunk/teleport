// Copyright 2023 Gravitational, Inc
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
// 	protoc-gen-go v1.36.3
// 	protoc        (unknown)
// source: teleport/externalauditstorage/v1/externalauditstorage.proto

package externalauditstoragev1

import (
	v1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/header/v1"
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

// ExternalAuditStorage contains External Audit Storage configuration.
// It contains configuration that allows users to store audit events and session
// recordings on customer-owned infra instead of in Teleport Cloud.
type ExternalAuditStorage struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Header is the header for the resource.
	Header *v1.ResourceHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// Spec is the specification for external audit storage.
	Spec          *ExternalAuditStorageSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExternalAuditStorage) Reset() {
	*x = ExternalAuditStorage{}
	mi := &file_teleport_externalauditstorage_v1_externalauditstorage_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExternalAuditStorage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalAuditStorage) ProtoMessage() {}

func (x *ExternalAuditStorage) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_externalauditstorage_v1_externalauditstorage_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalAuditStorage.ProtoReflect.Descriptor instead.
func (*ExternalAuditStorage) Descriptor() ([]byte, []int) {
	return file_teleport_externalauditstorage_v1_externalauditstorage_proto_rawDescGZIP(), []int{0}
}

func (x *ExternalAuditStorage) GetHeader() *v1.ResourceHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ExternalAuditStorage) GetSpec() *ExternalAuditStorageSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

// ExternalAuditStorageConfigSpec is the specification of external audit storage.
type ExternalAuditStorageSpec struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// IntegrationName is name of an existing AWS OIDC integration used to
	// authenticate to the external AWS account.
	IntegrationName string `protobuf:"bytes,1,opt,name=integration_name,json=integrationName,proto3" json:"integration_name,omitempty"`
	// Region is the AWS region where the infrastructure is hosted.
	Region string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	// SessionRecordingsURI is the S3 path used to store session recordings.
	SessionRecordingsUri string `protobuf:"bytes,3,opt,name=session_recordings_uri,json=sessionRecordingsUri,proto3" json:"session_recordings_uri,omitempty"`
	// AuditEventsLongTermURI is the S3 path used to store batched parquet files
	// with audit events.
	AuditEventsLongTermUri string `protobuf:"bytes,4,opt,name=audit_events_long_term_uri,json=auditEventsLongTermUri,proto3" json:"audit_events_long_term_uri,omitempty"`
	// AthenaResultsURI is the S3 path used to store temporary results of Athena
	// queries.
	AthenaResultsUri string `protobuf:"bytes,5,opt,name=athena_results_uri,json=athenaResultsUri,proto3" json:"athena_results_uri,omitempty"`
	// AthenaWorkgroup is the workgroup used for Athena audit log queries.
	AthenaWorkgroup string `protobuf:"bytes,6,opt,name=athena_workgroup,json=athenaWorkgroup,proto3" json:"athena_workgroup,omitempty"`
	// GlueDatabase is the database used for Athena audit log queries.
	GlueDatabase string `protobuf:"bytes,7,opt,name=glue_database,json=glueDatabase,proto3" json:"glue_database,omitempty"`
	// GlueTable is the table used for Athena audit log queries.
	GlueTable string `protobuf:"bytes,8,opt,name=glue_table,json=glueTable,proto3" json:"glue_table,omitempty"`
	// PolicyName is the name of the IAM policy attached to the OIDC integration
	// role.
	PolicyName    string `protobuf:"bytes,9,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExternalAuditStorageSpec) Reset() {
	*x = ExternalAuditStorageSpec{}
	mi := &file_teleport_externalauditstorage_v1_externalauditstorage_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExternalAuditStorageSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalAuditStorageSpec) ProtoMessage() {}

func (x *ExternalAuditStorageSpec) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_externalauditstorage_v1_externalauditstorage_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalAuditStorageSpec.ProtoReflect.Descriptor instead.
func (*ExternalAuditStorageSpec) Descriptor() ([]byte, []int) {
	return file_teleport_externalauditstorage_v1_externalauditstorage_proto_rawDescGZIP(), []int{1}
}

func (x *ExternalAuditStorageSpec) GetIntegrationName() string {
	if x != nil {
		return x.IntegrationName
	}
	return ""
}

func (x *ExternalAuditStorageSpec) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *ExternalAuditStorageSpec) GetSessionRecordingsUri() string {
	if x != nil {
		return x.SessionRecordingsUri
	}
	return ""
}

func (x *ExternalAuditStorageSpec) GetAuditEventsLongTermUri() string {
	if x != nil {
		return x.AuditEventsLongTermUri
	}
	return ""
}

func (x *ExternalAuditStorageSpec) GetAthenaResultsUri() string {
	if x != nil {
		return x.AthenaResultsUri
	}
	return ""
}

func (x *ExternalAuditStorageSpec) GetAthenaWorkgroup() string {
	if x != nil {
		return x.AthenaWorkgroup
	}
	return ""
}

func (x *ExternalAuditStorageSpec) GetGlueDatabase() string {
	if x != nil {
		return x.GlueDatabase
	}
	return ""
}

func (x *ExternalAuditStorageSpec) GetGlueTable() string {
	if x != nil {
		return x.GlueTable
	}
	return ""
}

func (x *ExternalAuditStorageSpec) GetPolicyName() string {
	if x != nil {
		return x.PolicyName
	}
	return ""
}

var File_teleport_externalauditstorage_v1_externalauditstorage_proto protoreflect.FileDescriptor

var file_teleport_externalauditstorage_v1_externalauditstorage_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x27, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x14, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x12, 0x3a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x4e, 0x0a,
	0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x74, 0x65,
	0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x8d, 0x03,
	0x0a, 0x18, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a,
	0x16, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x67, 0x73, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x73,
	0x55, 0x72, 0x69, 0x12, 0x3a, 0x0a, 0x1a, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x5f, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x5f, 0x75, 0x72,
	0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x61, 0x75, 0x64, 0x69, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x4c, 0x6f, 0x6e, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x55, 0x72, 0x69, 0x12,
	0x2c, 0x0a, 0x12, 0x61, 0x74, 0x68, 0x65, 0x6e, 0x61, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x74, 0x68,
	0x65, 0x6e, 0x61, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x55, 0x72, 0x69, 0x12, 0x29, 0x0a,
	0x10, 0x61, 0x74, 0x68, 0x65, 0x6e, 0x61, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x74, 0x68, 0x65, 0x6e, 0x61, 0x57,
	0x6f, 0x72, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x23, 0x0a, 0x0d, 0x67, 0x6c, 0x75, 0x65,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x67, 0x6c, 0x75, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x67, 0x6c, 0x75, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x67, 0x6c, 0x75, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x6c, 0x5a,
	0x6a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_teleport_externalauditstorage_v1_externalauditstorage_proto_rawDescOnce sync.Once
	file_teleport_externalauditstorage_v1_externalauditstorage_proto_rawDescData = file_teleport_externalauditstorage_v1_externalauditstorage_proto_rawDesc
)

func file_teleport_externalauditstorage_v1_externalauditstorage_proto_rawDescGZIP() []byte {
	file_teleport_externalauditstorage_v1_externalauditstorage_proto_rawDescOnce.Do(func() {
		file_teleport_externalauditstorage_v1_externalauditstorage_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_externalauditstorage_v1_externalauditstorage_proto_rawDescData)
	})
	return file_teleport_externalauditstorage_v1_externalauditstorage_proto_rawDescData
}

var file_teleport_externalauditstorage_v1_externalauditstorage_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_teleport_externalauditstorage_v1_externalauditstorage_proto_goTypes = []any{
	(*ExternalAuditStorage)(nil),     // 0: teleport.externalauditstorage.v1.ExternalAuditStorage
	(*ExternalAuditStorageSpec)(nil), // 1: teleport.externalauditstorage.v1.ExternalAuditStorageSpec
	(*v1.ResourceHeader)(nil),        // 2: teleport.header.v1.ResourceHeader
}
var file_teleport_externalauditstorage_v1_externalauditstorage_proto_depIdxs = []int32{
	2, // 0: teleport.externalauditstorage.v1.ExternalAuditStorage.header:type_name -> teleport.header.v1.ResourceHeader
	1, // 1: teleport.externalauditstorage.v1.ExternalAuditStorage.spec:type_name -> teleport.externalauditstorage.v1.ExternalAuditStorageSpec
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_teleport_externalauditstorage_v1_externalauditstorage_proto_init() }
func file_teleport_externalauditstorage_v1_externalauditstorage_proto_init() {
	if File_teleport_externalauditstorage_v1_externalauditstorage_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teleport_externalauditstorage_v1_externalauditstorage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_externalauditstorage_v1_externalauditstorage_proto_goTypes,
		DependencyIndexes: file_teleport_externalauditstorage_v1_externalauditstorage_proto_depIdxs,
		MessageInfos:      file_teleport_externalauditstorage_v1_externalauditstorage_proto_msgTypes,
	}.Build()
	File_teleport_externalauditstorage_v1_externalauditstorage_proto = out.File
	file_teleport_externalauditstorage_v1_externalauditstorage_proto_rawDesc = nil
	file_teleport_externalauditstorage_v1_externalauditstorage_proto_goTypes = nil
	file_teleport_externalauditstorage_v1_externalauditstorage_proto_depIdxs = nil
}
