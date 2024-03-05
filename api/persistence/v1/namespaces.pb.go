// The MIT License
//
// Copyright (c) 2020 Temporal Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// plugins:
// 	protoc-gen-go
// 	protoc
// source: temporal/server/api/persistence/v1/namespaces.proto

package persistence

import (
	reflect "reflect"
	sync "sync"

	v1 "go.temporal.io/api/enums/v1"
	v11 "go.temporal.io/api/namespace/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// detail column
type NamespaceDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info                        *NamespaceInfo              `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Config                      *NamespaceConfig            `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	ReplicationConfig           *NamespaceReplicationConfig `protobuf:"bytes,3,opt,name=replication_config,json=replicationConfig,proto3" json:"replication_config,omitempty"`
	ConfigVersion               int64                       `protobuf:"varint,4,opt,name=config_version,json=configVersion,proto3" json:"config_version,omitempty"`
	FailoverNotificationVersion int64                       `protobuf:"varint,5,opt,name=failover_notification_version,json=failoverNotificationVersion,proto3" json:"failover_notification_version,omitempty"`
	FailoverVersion             int64                       `protobuf:"varint,6,opt,name=failover_version,json=failoverVersion,proto3" json:"failover_version,omitempty"`
	FailoverEndTime             *timestamppb.Timestamp      `protobuf:"bytes,7,opt,name=failover_end_time,json=failoverEndTime,proto3" json:"failover_end_time,omitempty"`
}

func (x *NamespaceDetail) Reset() {
	*x = NamespaceDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_persistence_v1_namespaces_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespaceDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceDetail) ProtoMessage() {}

func (x *NamespaceDetail) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_persistence_v1_namespaces_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceDetail.ProtoReflect.Descriptor instead.
func (*NamespaceDetail) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_persistence_v1_namespaces_proto_rawDescGZIP(), []int{0}
}

func (x *NamespaceDetail) GetInfo() *NamespaceInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *NamespaceDetail) GetConfig() *NamespaceConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *NamespaceDetail) GetReplicationConfig() *NamespaceReplicationConfig {
	if x != nil {
		return x.ReplicationConfig
	}
	return nil
}

func (x *NamespaceDetail) GetConfigVersion() int64 {
	if x != nil {
		return x.ConfigVersion
	}
	return 0
}

func (x *NamespaceDetail) GetFailoverNotificationVersion() int64 {
	if x != nil {
		return x.FailoverNotificationVersion
	}
	return 0
}

func (x *NamespaceDetail) GetFailoverVersion() int64 {
	if x != nil {
		return x.FailoverVersion
	}
	return 0
}

func (x *NamespaceDetail) GetFailoverEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.FailoverEndTime
	}
	return nil
}

type NamespaceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	State       v1.NamespaceState `protobuf:"varint,2,opt,name=state,proto3,enum=temporal.api.enums.v1.NamespaceState" json:"state,omitempty"`
	Name        string            `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string            `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Owner       string            `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
	Data        map[string]string `protobuf:"bytes,6,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *NamespaceInfo) Reset() {
	*x = NamespaceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_persistence_v1_namespaces_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespaceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceInfo) ProtoMessage() {}

func (x *NamespaceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_persistence_v1_namespaces_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceInfo.ProtoReflect.Descriptor instead.
func (*NamespaceInfo) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_persistence_v1_namespaces_proto_rawDescGZIP(), []int{1}
}

func (x *NamespaceInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NamespaceInfo) GetState() v1.NamespaceState {
	if x != nil {
		return x.State
	}
	return v1.NamespaceState(0)
}

func (x *NamespaceInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NamespaceInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NamespaceInfo) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *NamespaceInfo) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

type NamespaceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retention                    *durationpb.Duration `protobuf:"bytes,1,opt,name=retention,proto3" json:"retention,omitempty"`
	ArchivalBucket               string               `protobuf:"bytes,2,opt,name=archival_bucket,json=archivalBucket,proto3" json:"archival_bucket,omitempty"`
	BadBinaries                  *v11.BadBinaries     `protobuf:"bytes,3,opt,name=bad_binaries,json=badBinaries,proto3" json:"bad_binaries,omitempty"`
	HistoryArchivalState         v1.ArchivalState     `protobuf:"varint,4,opt,name=history_archival_state,json=historyArchivalState,proto3,enum=temporal.api.enums.v1.ArchivalState" json:"history_archival_state,omitempty"`
	HistoryArchivalUri           string               `protobuf:"bytes,5,opt,name=history_archival_uri,json=historyArchivalUri,proto3" json:"history_archival_uri,omitempty"`
	VisibilityArchivalState      v1.ArchivalState     `protobuf:"varint,6,opt,name=visibility_archival_state,json=visibilityArchivalState,proto3,enum=temporal.api.enums.v1.ArchivalState" json:"visibility_archival_state,omitempty"`
	VisibilityArchivalUri        string               `protobuf:"bytes,7,opt,name=visibility_archival_uri,json=visibilityArchivalUri,proto3" json:"visibility_archival_uri,omitempty"`
	CustomSearchAttributeAliases map[string]string    `protobuf:"bytes,8,rep,name=custom_search_attribute_aliases,json=customSearchAttributeAliases,proto3" json:"custom_search_attribute_aliases,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *NamespaceConfig) Reset() {
	*x = NamespaceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_persistence_v1_namespaces_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespaceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceConfig) ProtoMessage() {}

func (x *NamespaceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_persistence_v1_namespaces_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceConfig.ProtoReflect.Descriptor instead.
func (*NamespaceConfig) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_persistence_v1_namespaces_proto_rawDescGZIP(), []int{2}
}

func (x *NamespaceConfig) GetRetention() *durationpb.Duration {
	if x != nil {
		return x.Retention
	}
	return nil
}

func (x *NamespaceConfig) GetArchivalBucket() string {
	if x != nil {
		return x.ArchivalBucket
	}
	return ""
}

func (x *NamespaceConfig) GetBadBinaries() *v11.BadBinaries {
	if x != nil {
		return x.BadBinaries
	}
	return nil
}

func (x *NamespaceConfig) GetHistoryArchivalState() v1.ArchivalState {
	if x != nil {
		return x.HistoryArchivalState
	}
	return v1.ArchivalState(0)
}

func (x *NamespaceConfig) GetHistoryArchivalUri() string {
	if x != nil {
		return x.HistoryArchivalUri
	}
	return ""
}

func (x *NamespaceConfig) GetVisibilityArchivalState() v1.ArchivalState {
	if x != nil {
		return x.VisibilityArchivalState
	}
	return v1.ArchivalState(0)
}

func (x *NamespaceConfig) GetVisibilityArchivalUri() string {
	if x != nil {
		return x.VisibilityArchivalUri
	}
	return ""
}

func (x *NamespaceConfig) GetCustomSearchAttributeAliases() map[string]string {
	if x != nil {
		return x.CustomSearchAttributeAliases
	}
	return nil
}

type NamespaceReplicationConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActiveClusterName string              `protobuf:"bytes,1,opt,name=active_cluster_name,json=activeClusterName,proto3" json:"active_cluster_name,omitempty"`
	Clusters          []string            `protobuf:"bytes,2,rep,name=clusters,proto3" json:"clusters,omitempty"`
	State             v1.ReplicationState `protobuf:"varint,3,opt,name=state,proto3,enum=temporal.api.enums.v1.ReplicationState" json:"state,omitempty"`
	FailoverHistory   []*FailoverStatus   `protobuf:"bytes,8,rep,name=failover_history,json=failoverHistory,proto3" json:"failover_history,omitempty"`
}

func (x *NamespaceReplicationConfig) Reset() {
	*x = NamespaceReplicationConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_persistence_v1_namespaces_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespaceReplicationConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceReplicationConfig) ProtoMessage() {}

func (x *NamespaceReplicationConfig) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_persistence_v1_namespaces_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceReplicationConfig.ProtoReflect.Descriptor instead.
func (*NamespaceReplicationConfig) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_persistence_v1_namespaces_proto_rawDescGZIP(), []int{3}
}

func (x *NamespaceReplicationConfig) GetActiveClusterName() string {
	if x != nil {
		return x.ActiveClusterName
	}
	return ""
}

func (x *NamespaceReplicationConfig) GetClusters() []string {
	if x != nil {
		return x.Clusters
	}
	return nil
}

func (x *NamespaceReplicationConfig) GetState() v1.ReplicationState {
	if x != nil {
		return x.State
	}
	return v1.ReplicationState(0)
}

func (x *NamespaceReplicationConfig) GetFailoverHistory() []*FailoverStatus {
	if x != nil {
		return x.FailoverHistory
	}
	return nil
}

// Represents a historical replication status of a Namespace
type FailoverStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FailoverTime    *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=failover_time,json=failoverTime,proto3" json:"failover_time,omitempty"`
	FailoverVersion int64                  `protobuf:"varint,2,opt,name=failover_version,json=failoverVersion,proto3" json:"failover_version,omitempty"`
}

func (x *FailoverStatus) Reset() {
	*x = FailoverStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_persistence_v1_namespaces_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FailoverStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FailoverStatus) ProtoMessage() {}

func (x *FailoverStatus) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_persistence_v1_namespaces_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FailoverStatus.ProtoReflect.Descriptor instead.
func (*FailoverStatus) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_persistence_v1_namespaces_proto_rawDescGZIP(), []int{4}
}

func (x *FailoverStatus) GetFailoverTime() *timestamppb.Timestamp {
	if x != nil {
		return x.FailoverTime
	}
	return nil
}

func (x *FailoverStatus) GetFailoverVersion() int64 {
	if x != nil {
		return x.FailoverVersion
	}
	return 0
}

var File_temporal_server_api_persistence_v1_namespaces_proto protoreflect.FileDescriptor

var file_temporal_server_api_persistence_v1_namespaces_proto_rawDesc = []byte{
	0x0a, 0x33, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x22, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x25, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x04,
	0x0a, 0x0f, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x12, 0x49, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e,
	0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x42, 0x02, 0x68, 0x00, 0x12, 0x4f, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x42, 0x02, 0x68, 0x00, 0x12, 0x71, 0x0a, 0x12, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3e,
	0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x11, 0x72, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x02, 0x68,
	0x00, 0x12, 0x29, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x02, 0x68, 0x00, 0x12, 0x46, 0x0a, 0x1d, 0x66, 0x61, 0x69,
	0x6c, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1b,
	0x66, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x02, 0x68, 0x00, 0x12, 0x2d, 0x0a,
	0x10, 0x66, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x66, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x02, 0x68, 0x00, 0x12, 0x4a, 0x0a, 0x11, 0x66, 0x61,
	0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f,
	0x66, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x02,
	0x68, 0x00, 0x22, 0xd2, 0x02, 0x0a, 0x0d, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x42, 0x02, 0x68, 0x00, 0x12, 0x3f, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x42, 0x02, 0x68, 0x00, 0x12, 0x16, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x02, 0x68, 0x00, 0x12, 0x24, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x02, 0x68, 0x00, 0x12, 0x18,
	0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x42, 0x02, 0x68, 0x00, 0x12, 0x53, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x42, 0x02, 0x68, 0x00, 0x1a, 0x3f, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x42, 0x02, 0x68, 0x00, 0x12, 0x18, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x02, 0x68,
	0x00, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xfe, 0x05, 0x0a, 0x0f, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3b, 0x0a, 0x09, 0x72, 0x65, 0x74, 0x65, 0x6e,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x02,
	0x68, 0x00, 0x12, 0x2b, 0x0a, 0x0f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x61, 0x6c, 0x5f, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x61, 0x6c, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x02, 0x68, 0x00, 0x12, 0x4d, 0x0a, 0x0c,
	0x62, 0x61, 0x64, 0x5f, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x64, 0x42,
	0x69, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x52, 0x0b, 0x62, 0x61, 0x64, 0x42, 0x69, 0x6e, 0x61, 0x72,
	0x69, 0x65, 0x73, 0x42, 0x02, 0x68, 0x00, 0x12, 0x5e, 0x0a, 0x16, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x14, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x02,
	0x68, 0x00, 0x12, 0x34, 0x0a, 0x14, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12,
	0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x61, 0x6c, 0x55, 0x72,
	0x69, 0x42, 0x02, 0x68, 0x00, 0x12, 0x64, 0x0a, 0x19, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x17, 0x76, 0x69, 0x73, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x42, 0x02, 0x68, 0x00, 0x12, 0x3a, 0x0a, 0x17, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x69, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x61, 0x6c, 0x55, 0x72, 0x69, 0x42, 0x02, 0x68, 0x00, 0x12, 0xa0,
	0x01, 0x0a, 0x1f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x65, 0x73,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x55, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x41, 0x6c, 0x69,
	0x61, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x1c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x41, 0x6c,
	0x69, 0x61, 0x73, 0x65, 0x73, 0x42, 0x02, 0x68, 0x00, 0x1a, 0x57, 0x0a, 0x21, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x41, 0x6c, 0x69, 0x61, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x42, 0x02, 0x68,
	0x00, 0x12, 0x18, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x02, 0x68, 0x00, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x96, 0x02,
	0x0a, 0x1a, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x32, 0x0a, 0x13, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x02, 0x68, 0x00, 0x12, 0x1e, 0x0a, 0x08, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x42, 0x02, 0x68, 0x00, 0x12, 0x41, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x42, 0x02, 0x68, 0x00, 0x12, 0x61, 0x0a, 0x10, 0x66, 0x61, 0x69,
	0x6c, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x32, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x0f, 0x66, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x42, 0x02, 0x68, 0x00, 0x22, 0x84, 0x01, 0x0a, 0x0e, 0x46, 0x61, 0x69, 0x6c,
	0x6f, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x43, 0x0a, 0x0d, 0x66, 0x61, 0x69,
	0x6c, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x66, 0x61, 0x69, 0x6c, 0x6f,
	0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x02, 0x68, 0x00, 0x12, 0x2d, 0x0a, 0x10, 0x66, 0x61,
	0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0f, 0x66, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x42, 0x02, 0x68, 0x00, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x6f, 0x2e, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x3b, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_server_api_persistence_v1_namespaces_proto_rawDescOnce sync.Once
	file_temporal_server_api_persistence_v1_namespaces_proto_rawDescData = file_temporal_server_api_persistence_v1_namespaces_proto_rawDesc
)

func file_temporal_server_api_persistence_v1_namespaces_proto_rawDescGZIP() []byte {
	file_temporal_server_api_persistence_v1_namespaces_proto_rawDescOnce.Do(func() {
		file_temporal_server_api_persistence_v1_namespaces_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_server_api_persistence_v1_namespaces_proto_rawDescData)
	})
	return file_temporal_server_api_persistence_v1_namespaces_proto_rawDescData
}

var file_temporal_server_api_persistence_v1_namespaces_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_temporal_server_api_persistence_v1_namespaces_proto_goTypes = []interface{}{
	(*NamespaceDetail)(nil),            // 0: temporal.server.api.persistence.v1.NamespaceDetail
	(*NamespaceInfo)(nil),              // 1: temporal.server.api.persistence.v1.NamespaceInfo
	(*NamespaceConfig)(nil),            // 2: temporal.server.api.persistence.v1.NamespaceConfig
	(*NamespaceReplicationConfig)(nil), // 3: temporal.server.api.persistence.v1.NamespaceReplicationConfig
	(*FailoverStatus)(nil),             // 4: temporal.server.api.persistence.v1.FailoverStatus
	nil,                                // 5: temporal.server.api.persistence.v1.NamespaceInfo.DataEntry
	nil,                                // 6: temporal.server.api.persistence.v1.NamespaceConfig.CustomSearchAttributeAliasesEntry
	(*timestamppb.Timestamp)(nil),      // 7: google.protobuf.Timestamp
	(v1.NamespaceState)(0),             // 8: temporal.api.enums.v1.NamespaceState
	(*durationpb.Duration)(nil),        // 9: google.protobuf.Duration
	(*v11.BadBinaries)(nil),            // 10: temporal.api.namespace.v1.BadBinaries
	(v1.ArchivalState)(0),              // 11: temporal.api.enums.v1.ArchivalState
	(v1.ReplicationState)(0),           // 12: temporal.api.enums.v1.ReplicationState
}
var file_temporal_server_api_persistence_v1_namespaces_proto_depIdxs = []int32{
	1,  // 0: temporal.server.api.persistence.v1.NamespaceDetail.info:type_name -> temporal.server.api.persistence.v1.NamespaceInfo
	2,  // 1: temporal.server.api.persistence.v1.NamespaceDetail.config:type_name -> temporal.server.api.persistence.v1.NamespaceConfig
	3,  // 2: temporal.server.api.persistence.v1.NamespaceDetail.replication_config:type_name -> temporal.server.api.persistence.v1.NamespaceReplicationConfig
	7,  // 3: temporal.server.api.persistence.v1.NamespaceDetail.failover_end_time:type_name -> google.protobuf.Timestamp
	8,  // 4: temporal.server.api.persistence.v1.NamespaceInfo.state:type_name -> temporal.api.enums.v1.NamespaceState
	5,  // 5: temporal.server.api.persistence.v1.NamespaceInfo.data:type_name -> temporal.server.api.persistence.v1.NamespaceInfo.DataEntry
	9,  // 6: temporal.server.api.persistence.v1.NamespaceConfig.retention:type_name -> google.protobuf.Duration
	10, // 7: temporal.server.api.persistence.v1.NamespaceConfig.bad_binaries:type_name -> temporal.api.namespace.v1.BadBinaries
	11, // 8: temporal.server.api.persistence.v1.NamespaceConfig.history_archival_state:type_name -> temporal.api.enums.v1.ArchivalState
	11, // 9: temporal.server.api.persistence.v1.NamespaceConfig.visibility_archival_state:type_name -> temporal.api.enums.v1.ArchivalState
	6,  // 10: temporal.server.api.persistence.v1.NamespaceConfig.custom_search_attribute_aliases:type_name -> temporal.server.api.persistence.v1.NamespaceConfig.CustomSearchAttributeAliasesEntry
	12, // 11: temporal.server.api.persistence.v1.NamespaceReplicationConfig.state:type_name -> temporal.api.enums.v1.ReplicationState
	4,  // 12: temporal.server.api.persistence.v1.NamespaceReplicationConfig.failover_history:type_name -> temporal.server.api.persistence.v1.FailoverStatus
	7,  // 13: temporal.server.api.persistence.v1.FailoverStatus.failover_time:type_name -> google.protobuf.Timestamp
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_temporal_server_api_persistence_v1_namespaces_proto_init() }
func file_temporal_server_api_persistence_v1_namespaces_proto_init() {
	if File_temporal_server_api_persistence_v1_namespaces_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_temporal_server_api_persistence_v1_namespaces_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespaceDetail); i {
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
		file_temporal_server_api_persistence_v1_namespaces_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespaceInfo); i {
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
		file_temporal_server_api_persistence_v1_namespaces_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespaceConfig); i {
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
		file_temporal_server_api_persistence_v1_namespaces_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespaceReplicationConfig); i {
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
		file_temporal_server_api_persistence_v1_namespaces_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FailoverStatus); i {
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
			RawDescriptor: file_temporal_server_api_persistence_v1_namespaces_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_server_api_persistence_v1_namespaces_proto_goTypes,
		DependencyIndexes: file_temporal_server_api_persistence_v1_namespaces_proto_depIdxs,
		MessageInfos:      file_temporal_server_api_persistence_v1_namespaces_proto_msgTypes,
	}.Build()
	File_temporal_server_api_persistence_v1_namespaces_proto = out.File
	file_temporal_server_api_persistence_v1_namespaces_proto_rawDesc = nil
	file_temporal_server_api_persistence_v1_namespaces_proto_goTypes = nil
	file_temporal_server_api_persistence_v1_namespaces_proto_depIdxs = nil
}
