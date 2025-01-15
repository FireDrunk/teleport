// Copyright 2024 Gravitational, Inc
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
// source: teleport/gitserver/v1/git_server_service.proto

package gitserverv1

import (
	types "github.com/gravitational/teleport/api/types"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// CreateGitServerRequest is a request to create a Git server.
type CreateGitServerRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Server is the Git server to create.
	Server        *types.ServerV2 `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateGitServerRequest) Reset() {
	*x = CreateGitServerRequest{}
	mi := &file_teleport_gitserver_v1_git_server_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateGitServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGitServerRequest) ProtoMessage() {}

func (x *CreateGitServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_gitserver_v1_git_server_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGitServerRequest.ProtoReflect.Descriptor instead.
func (*CreateGitServerRequest) Descriptor() ([]byte, []int) {
	return file_teleport_gitserver_v1_git_server_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateGitServerRequest) GetServer() *types.ServerV2 {
	if x != nil {
		return x.Server
	}
	return nil
}

// GetGitServerRequest is a request to get a Git server.
type GetGitServerRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Name is the uuid of the server.
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetGitServerRequest) Reset() {
	*x = GetGitServerRequest{}
	mi := &file_teleport_gitserver_v1_git_server_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetGitServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGitServerRequest) ProtoMessage() {}

func (x *GetGitServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_gitserver_v1_git_server_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGitServerRequest.ProtoReflect.Descriptor instead.
func (*GetGitServerRequest) Descriptor() ([]byte, []int) {
	return file_teleport_gitserver_v1_git_server_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetGitServerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// ListGitServersRequest is the request to list Git servers.
type ListGitServersRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The maximum number of items to return.
	// The server may impose a different page size at its discretion.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The page_token is the next_page_token value returned from a previous List request, if any.
	PageToken     string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListGitServersRequest) Reset() {
	*x = ListGitServersRequest{}
	mi := &file_teleport_gitserver_v1_git_server_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListGitServersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGitServersRequest) ProtoMessage() {}

func (x *ListGitServersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_gitserver_v1_git_server_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGitServersRequest.ProtoReflect.Descriptor instead.
func (*ListGitServersRequest) Descriptor() ([]byte, []int) {
	return file_teleport_gitserver_v1_git_server_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListGitServersRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListGitServersRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// ListGitServersResponse is the response to ListGitServers.
type ListGitServersResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The page of Git servers that matched the request.
	Servers []*types.ServerV2 `protobuf:"bytes,1,rep,name=servers,proto3" json:"servers,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no
	// more results in the list.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListGitServersResponse) Reset() {
	*x = ListGitServersResponse{}
	mi := &file_teleport_gitserver_v1_git_server_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListGitServersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGitServersResponse) ProtoMessage() {}

func (x *ListGitServersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_gitserver_v1_git_server_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGitServersResponse.ProtoReflect.Descriptor instead.
func (*ListGitServersResponse) Descriptor() ([]byte, []int) {
	return file_teleport_gitserver_v1_git_server_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListGitServersResponse) GetServers() []*types.ServerV2 {
	if x != nil {
		return x.Servers
	}
	return nil
}

func (x *ListGitServersResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// UpdateGitServerRequest is the request to update a Git server.
type UpdateGitServerRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Server is the Git server to update.
	Server        *types.ServerV2 `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateGitServerRequest) Reset() {
	*x = UpdateGitServerRequest{}
	mi := &file_teleport_gitserver_v1_git_server_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateGitServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGitServerRequest) ProtoMessage() {}

func (x *UpdateGitServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_gitserver_v1_git_server_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGitServerRequest.ProtoReflect.Descriptor instead.
func (*UpdateGitServerRequest) Descriptor() ([]byte, []int) {
	return file_teleport_gitserver_v1_git_server_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateGitServerRequest) GetServer() *types.ServerV2 {
	if x != nil {
		return x.Server
	}
	return nil
}

// UpsertGitServerRequest is the request to upsert a Git server.
type UpsertGitServerRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Server is the Git server to upsert.
	Server        *types.ServerV2 `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpsertGitServerRequest) Reset() {
	*x = UpsertGitServerRequest{}
	mi := &file_teleport_gitserver_v1_git_server_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertGitServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertGitServerRequest) ProtoMessage() {}

func (x *UpsertGitServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_gitserver_v1_git_server_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertGitServerRequest.ProtoReflect.Descriptor instead.
func (*UpsertGitServerRequest) Descriptor() ([]byte, []int) {
	return file_teleport_gitserver_v1_git_server_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpsertGitServerRequest) GetServer() *types.ServerV2 {
	if x != nil {
		return x.Server
	}
	return nil
}

// DeleteGitServerRequest is the request to delete a Git server.
type DeleteGitServerRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Name is the uuid of the server.
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteGitServerRequest) Reset() {
	*x = DeleteGitServerRequest{}
	mi := &file_teleport_gitserver_v1_git_server_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteGitServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGitServerRequest) ProtoMessage() {}

func (x *DeleteGitServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_gitserver_v1_git_server_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGitServerRequest.ProtoReflect.Descriptor instead.
func (*DeleteGitServerRequest) Descriptor() ([]byte, []int) {
	return file_teleport_gitserver_v1_git_server_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteGitServerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// CreateGitHubAuthRequestRequest is the request for CreateGitHubAuthRequest.
type CreateGitHubAuthRequestRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Request is the basic GitHub auth request.
	Request *types.GithubAuthRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	// Organization is the GitHub organization that the user is accessing.
	Organization  string `protobuf:"bytes,2,opt,name=organization,proto3" json:"organization,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateGitHubAuthRequestRequest) Reset() {
	*x = CreateGitHubAuthRequestRequest{}
	mi := &file_teleport_gitserver_v1_git_server_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateGitHubAuthRequestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGitHubAuthRequestRequest) ProtoMessage() {}

func (x *CreateGitHubAuthRequestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_gitserver_v1_git_server_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGitHubAuthRequestRequest.ProtoReflect.Descriptor instead.
func (*CreateGitHubAuthRequestRequest) Descriptor() ([]byte, []int) {
	return file_teleport_gitserver_v1_git_server_service_proto_rawDescGZIP(), []int{7}
}

func (x *CreateGitHubAuthRequestRequest) GetRequest() *types.GithubAuthRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *CreateGitHubAuthRequestRequest) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

var File_teleport_gitserver_v1_git_server_service_proto protoreflect.FileDescriptor

var file_teleport_gitserver_v1_git_server_service_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x67, 0x69, 0x74, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x69, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x15, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x67, 0x69, 0x74, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6c,
	0x65, 0x67, 0x61, 0x63, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x47, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x56, 0x32, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x29, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x47, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x53, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x69, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x6b, 0x0a, 0x16, 0x4c, 0x69,
	0x73, 0x74, 0x47, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x56, 0x32, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12,
	0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x41, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x47, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x56, 0x32, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x41, 0x0a, 0x16, 0x55, 0x70,
	0x73, 0x65, 0x72, 0x74, 0x47, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x56, 0x32, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x2c, 0x0a,
	0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x78, 0x0a, 0x1e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x69, 0x74, 0x48, 0x75, 0x62, 0x41, 0x75, 0x74, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a,
	0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x41, 0x75, 0x74,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x8d, 0x05, 0x0a, 0x10, 0x47, 0x69, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0f, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x47, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x2d, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x67, 0x69, 0x74, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x69, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x56, 0x32, 0x12, 0x4b, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x47, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x2a, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x67, 0x69, 0x74, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x56, 0x32, 0x12, 0x6d, 0x0a, 0x0e, 0x4c, 0x69,
	0x73, 0x74, 0x47, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x2c, 0x2e, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x67, 0x69, 0x74, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x74, 0x65, 0x6c,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x67, 0x69, 0x74, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0f, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x47, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x2d, 0x2e, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x67, 0x69, 0x74, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x69, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x56, 0x32, 0x12, 0x51, 0x0a, 0x0f,
	0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x47, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12,
	0x2d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x67, 0x69, 0x74, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x47, 0x69,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x56, 0x32, 0x12,
	0x58, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x12, 0x2d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x67, 0x69,
	0x74, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x47, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x6a, 0x0a, 0x17, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x47, 0x69, 0x74, 0x48, 0x75, 0x62, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x67, 0x69, 0x74, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x47, 0x69, 0x74, 0x48, 0x75, 0x62, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x56, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2f, 0x67, 0x69, 0x74, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x3b, 0x67, 0x69, 0x74, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_gitserver_v1_git_server_service_proto_rawDescOnce sync.Once
	file_teleport_gitserver_v1_git_server_service_proto_rawDescData = file_teleport_gitserver_v1_git_server_service_proto_rawDesc
)

func file_teleport_gitserver_v1_git_server_service_proto_rawDescGZIP() []byte {
	file_teleport_gitserver_v1_git_server_service_proto_rawDescOnce.Do(func() {
		file_teleport_gitserver_v1_git_server_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_gitserver_v1_git_server_service_proto_rawDescData)
	})
	return file_teleport_gitserver_v1_git_server_service_proto_rawDescData
}

var file_teleport_gitserver_v1_git_server_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_teleport_gitserver_v1_git_server_service_proto_goTypes = []any{
	(*CreateGitServerRequest)(nil),         // 0: teleport.gitserver.v1.CreateGitServerRequest
	(*GetGitServerRequest)(nil),            // 1: teleport.gitserver.v1.GetGitServerRequest
	(*ListGitServersRequest)(nil),          // 2: teleport.gitserver.v1.ListGitServersRequest
	(*ListGitServersResponse)(nil),         // 3: teleport.gitserver.v1.ListGitServersResponse
	(*UpdateGitServerRequest)(nil),         // 4: teleport.gitserver.v1.UpdateGitServerRequest
	(*UpsertGitServerRequest)(nil),         // 5: teleport.gitserver.v1.UpsertGitServerRequest
	(*DeleteGitServerRequest)(nil),         // 6: teleport.gitserver.v1.DeleteGitServerRequest
	(*CreateGitHubAuthRequestRequest)(nil), // 7: teleport.gitserver.v1.CreateGitHubAuthRequestRequest
	(*types.ServerV2)(nil),                 // 8: types.ServerV2
	(*types.GithubAuthRequest)(nil),        // 9: types.GithubAuthRequest
	(*emptypb.Empty)(nil),                  // 10: google.protobuf.Empty
}
var file_teleport_gitserver_v1_git_server_service_proto_depIdxs = []int32{
	8,  // 0: teleport.gitserver.v1.CreateGitServerRequest.server:type_name -> types.ServerV2
	8,  // 1: teleport.gitserver.v1.ListGitServersResponse.servers:type_name -> types.ServerV2
	8,  // 2: teleport.gitserver.v1.UpdateGitServerRequest.server:type_name -> types.ServerV2
	8,  // 3: teleport.gitserver.v1.UpsertGitServerRequest.server:type_name -> types.ServerV2
	9,  // 4: teleport.gitserver.v1.CreateGitHubAuthRequestRequest.request:type_name -> types.GithubAuthRequest
	0,  // 5: teleport.gitserver.v1.GitServerService.CreateGitServer:input_type -> teleport.gitserver.v1.CreateGitServerRequest
	1,  // 6: teleport.gitserver.v1.GitServerService.GetGitServer:input_type -> teleport.gitserver.v1.GetGitServerRequest
	2,  // 7: teleport.gitserver.v1.GitServerService.ListGitServers:input_type -> teleport.gitserver.v1.ListGitServersRequest
	4,  // 8: teleport.gitserver.v1.GitServerService.UpdateGitServer:input_type -> teleport.gitserver.v1.UpdateGitServerRequest
	5,  // 9: teleport.gitserver.v1.GitServerService.UpsertGitServer:input_type -> teleport.gitserver.v1.UpsertGitServerRequest
	6,  // 10: teleport.gitserver.v1.GitServerService.DeleteGitServer:input_type -> teleport.gitserver.v1.DeleteGitServerRequest
	7,  // 11: teleport.gitserver.v1.GitServerService.CreateGitHubAuthRequest:input_type -> teleport.gitserver.v1.CreateGitHubAuthRequestRequest
	8,  // 12: teleport.gitserver.v1.GitServerService.CreateGitServer:output_type -> types.ServerV2
	8,  // 13: teleport.gitserver.v1.GitServerService.GetGitServer:output_type -> types.ServerV2
	3,  // 14: teleport.gitserver.v1.GitServerService.ListGitServers:output_type -> teleport.gitserver.v1.ListGitServersResponse
	8,  // 15: teleport.gitserver.v1.GitServerService.UpdateGitServer:output_type -> types.ServerV2
	8,  // 16: teleport.gitserver.v1.GitServerService.UpsertGitServer:output_type -> types.ServerV2
	10, // 17: teleport.gitserver.v1.GitServerService.DeleteGitServer:output_type -> google.protobuf.Empty
	9,  // 18: teleport.gitserver.v1.GitServerService.CreateGitHubAuthRequest:output_type -> types.GithubAuthRequest
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_teleport_gitserver_v1_git_server_service_proto_init() }
func file_teleport_gitserver_v1_git_server_service_proto_init() {
	if File_teleport_gitserver_v1_git_server_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teleport_gitserver_v1_git_server_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teleport_gitserver_v1_git_server_service_proto_goTypes,
		DependencyIndexes: file_teleport_gitserver_v1_git_server_service_proto_depIdxs,
		MessageInfos:      file_teleport_gitserver_v1_git_server_service_proto_msgTypes,
	}.Build()
	File_teleport_gitserver_v1_git_server_service_proto = out.File
	file_teleport_gitserver_v1_git_server_service_proto_rawDesc = nil
	file_teleport_gitserver_v1_git_server_service_proto_goTypes = nil
	file_teleport_gitserver_v1_git_server_service_proto_depIdxs = nil
}
