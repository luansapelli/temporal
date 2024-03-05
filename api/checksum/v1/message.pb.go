// The MIT License
//
// Copyright (c) 2019 Temporal Technologies, Inc.
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
// source: temporal/server/api/checksum/v1/message.proto

package checksum

import (
	reflect "reflect"
	sync "sync"

	v11 "go.temporal.io/api/enums/v1"
	v1 "go.temporal.io/server/api/enums/v1"
	v12 "go.temporal.io/server/api/history/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MutableStateChecksumPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CancelRequested                   bool                        `protobuf:"varint,1,opt,name=cancel_requested,json=cancelRequested,proto3" json:"cancel_requested,omitempty"`
	State                             v1.WorkflowExecutionState   `protobuf:"varint,2,opt,name=state,proto3,enum=temporal.server.api.enums.v1.WorkflowExecutionState" json:"state,omitempty"`
	Status                            v11.WorkflowExecutionStatus `protobuf:"varint,3,opt,name=status,proto3,enum=temporal.api.enums.v1.WorkflowExecutionStatus" json:"status,omitempty"`
	LastWriteVersion                  int64                       `protobuf:"varint,4,opt,name=last_write_version,json=lastWriteVersion,proto3" json:"last_write_version,omitempty"`
	LastWriteEventId                  int64                       `protobuf:"varint,5,opt,name=last_write_event_id,json=lastWriteEventId,proto3" json:"last_write_event_id,omitempty"`
	LastFirstEventId                  int64                       `protobuf:"varint,6,opt,name=last_first_event_id,json=lastFirstEventId,proto3" json:"last_first_event_id,omitempty"`
	NextEventId                       int64                       `protobuf:"varint,7,opt,name=next_event_id,json=nextEventId,proto3" json:"next_event_id,omitempty"`
	LastProcessedEventId              int64                       `protobuf:"varint,8,opt,name=last_processed_event_id,json=lastProcessedEventId,proto3" json:"last_processed_event_id,omitempty"`
	SignalCount                       int64                       `protobuf:"varint,9,opt,name=signal_count,json=signalCount,proto3" json:"signal_count,omitempty"`
	ActivityCount                     int64                       `protobuf:"varint,21,opt,name=activity_count,json=activityCount,proto3" json:"activity_count,omitempty"`
	ChildExecutionCount               int64                       `protobuf:"varint,22,opt,name=child_execution_count,json=childExecutionCount,proto3" json:"child_execution_count,omitempty"`
	UserTimerCount                    int64                       `protobuf:"varint,23,opt,name=user_timer_count,json=userTimerCount,proto3" json:"user_timer_count,omitempty"`
	RequestCancelExternalCount        int64                       `protobuf:"varint,24,opt,name=request_cancel_external_count,json=requestCancelExternalCount,proto3" json:"request_cancel_external_count,omitempty"`
	SignalExternalCount               int64                       `protobuf:"varint,25,opt,name=signal_external_count,json=signalExternalCount,proto3" json:"signal_external_count,omitempty"`
	WorkflowTaskAttempt               int32                       `protobuf:"varint,10,opt,name=workflow_task_attempt,json=workflowTaskAttempt,proto3" json:"workflow_task_attempt,omitempty"`
	WorkflowTaskVersion               int64                       `protobuf:"varint,11,opt,name=workflow_task_version,json=workflowTaskVersion,proto3" json:"workflow_task_version,omitempty"`
	WorkflowTaskScheduledEventId      int64                       `protobuf:"varint,12,opt,name=workflow_task_scheduled_event_id,json=workflowTaskScheduledEventId,proto3" json:"workflow_task_scheduled_event_id,omitempty"`
	WorkflowTaskStartedEventId        int64                       `protobuf:"varint,13,opt,name=workflow_task_started_event_id,json=workflowTaskStartedEventId,proto3" json:"workflow_task_started_event_id,omitempty"`
	PendingTimerStartedEventIds       []int64                     `protobuf:"varint,14,rep,packed,name=pending_timer_started_event_ids,json=pendingTimerStartedEventIds,proto3" json:"pending_timer_started_event_ids,omitempty"`
	PendingActivityScheduledEventIds  []int64                     `protobuf:"varint,15,rep,packed,name=pending_activity_scheduled_event_ids,json=pendingActivityScheduledEventIds,proto3" json:"pending_activity_scheduled_event_ids,omitempty"`
	PendingSignalInitiatedEventIds    []int64                     `protobuf:"varint,16,rep,packed,name=pending_signal_initiated_event_ids,json=pendingSignalInitiatedEventIds,proto3" json:"pending_signal_initiated_event_ids,omitempty"`
	PendingReqCancelInitiatedEventIds []int64                     `protobuf:"varint,17,rep,packed,name=pending_req_cancel_initiated_event_ids,json=pendingReqCancelInitiatedEventIds,proto3" json:"pending_req_cancel_initiated_event_ids,omitempty"`
	PendingChildInitiatedEventIds     []int64                     `protobuf:"varint,18,rep,packed,name=pending_child_initiated_event_ids,json=pendingChildInitiatedEventIds,proto3" json:"pending_child_initiated_event_ids,omitempty"`
	StickyTaskQueueName               string                      `protobuf:"bytes,19,opt,name=sticky_task_queue_name,json=stickyTaskQueueName,proto3" json:"sticky_task_queue_name,omitempty"`
	VersionHistories                  *v12.VersionHistories       `protobuf:"bytes,20,opt,name=version_histories,json=versionHistories,proto3" json:"version_histories,omitempty"`
}

func (x *MutableStateChecksumPayload) Reset() {
	*x = MutableStateChecksumPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_checksum_v1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutableStateChecksumPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutableStateChecksumPayload) ProtoMessage() {}

func (x *MutableStateChecksumPayload) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_checksum_v1_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutableStateChecksumPayload.ProtoReflect.Descriptor instead.
func (*MutableStateChecksumPayload) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_checksum_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *MutableStateChecksumPayload) GetCancelRequested() bool {
	if x != nil {
		return x.CancelRequested
	}
	return false
}

func (x *MutableStateChecksumPayload) GetState() v1.WorkflowExecutionState {
	if x != nil {
		return x.State
	}
	return v1.WorkflowExecutionState(0)
}

func (x *MutableStateChecksumPayload) GetStatus() v11.WorkflowExecutionStatus {
	if x != nil {
		return x.Status
	}
	return v11.WorkflowExecutionStatus(0)
}

func (x *MutableStateChecksumPayload) GetLastWriteVersion() int64 {
	if x != nil {
		return x.LastWriteVersion
	}
	return 0
}

func (x *MutableStateChecksumPayload) GetLastWriteEventId() int64 {
	if x != nil {
		return x.LastWriteEventId
	}
	return 0
}

func (x *MutableStateChecksumPayload) GetLastFirstEventId() int64 {
	if x != nil {
		return x.LastFirstEventId
	}
	return 0
}

func (x *MutableStateChecksumPayload) GetNextEventId() int64 {
	if x != nil {
		return x.NextEventId
	}
	return 0
}

func (x *MutableStateChecksumPayload) GetLastProcessedEventId() int64 {
	if x != nil {
		return x.LastProcessedEventId
	}
	return 0
}

func (x *MutableStateChecksumPayload) GetSignalCount() int64 {
	if x != nil {
		return x.SignalCount
	}
	return 0
}

func (x *MutableStateChecksumPayload) GetActivityCount() int64 {
	if x != nil {
		return x.ActivityCount
	}
	return 0
}

func (x *MutableStateChecksumPayload) GetChildExecutionCount() int64 {
	if x != nil {
		return x.ChildExecutionCount
	}
	return 0
}

func (x *MutableStateChecksumPayload) GetUserTimerCount() int64 {
	if x != nil {
		return x.UserTimerCount
	}
	return 0
}

func (x *MutableStateChecksumPayload) GetRequestCancelExternalCount() int64 {
	if x != nil {
		return x.RequestCancelExternalCount
	}
	return 0
}

func (x *MutableStateChecksumPayload) GetSignalExternalCount() int64 {
	if x != nil {
		return x.SignalExternalCount
	}
	return 0
}

func (x *MutableStateChecksumPayload) GetWorkflowTaskAttempt() int32 {
	if x != nil {
		return x.WorkflowTaskAttempt
	}
	return 0
}

func (x *MutableStateChecksumPayload) GetWorkflowTaskVersion() int64 {
	if x != nil {
		return x.WorkflowTaskVersion
	}
	return 0
}

func (x *MutableStateChecksumPayload) GetWorkflowTaskScheduledEventId() int64 {
	if x != nil {
		return x.WorkflowTaskScheduledEventId
	}
	return 0
}

func (x *MutableStateChecksumPayload) GetWorkflowTaskStartedEventId() int64 {
	if x != nil {
		return x.WorkflowTaskStartedEventId
	}
	return 0
}

func (x *MutableStateChecksumPayload) GetPendingTimerStartedEventIds() []int64 {
	if x != nil {
		return x.PendingTimerStartedEventIds
	}
	return nil
}

func (x *MutableStateChecksumPayload) GetPendingActivityScheduledEventIds() []int64 {
	if x != nil {
		return x.PendingActivityScheduledEventIds
	}
	return nil
}

func (x *MutableStateChecksumPayload) GetPendingSignalInitiatedEventIds() []int64 {
	if x != nil {
		return x.PendingSignalInitiatedEventIds
	}
	return nil
}

func (x *MutableStateChecksumPayload) GetPendingReqCancelInitiatedEventIds() []int64 {
	if x != nil {
		return x.PendingReqCancelInitiatedEventIds
	}
	return nil
}

func (x *MutableStateChecksumPayload) GetPendingChildInitiatedEventIds() []int64 {
	if x != nil {
		return x.PendingChildInitiatedEventIds
	}
	return nil
}

func (x *MutableStateChecksumPayload) GetStickyTaskQueueName() string {
	if x != nil {
		return x.StickyTaskQueueName
	}
	return ""
}

func (x *MutableStateChecksumPayload) GetVersionHistories() *v12.VersionHistories {
	if x != nil {
		return x.VersionHistories
	}
	return nil
}

var File_temporal_server_api_checksum_v1_message_proto protoreflect.FileDescriptor

var file_temporal_server_api_checksum_v1_message_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x74, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x24, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2c, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x0c, 0x0a, 0x1b, 0x4d, 0x75, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2d, 0x0a, 0x10, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x63, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x42, 0x02, 0x68, 0x00,
	0x12, 0x4e, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34,
	0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x42, 0x02, 0x68, 0x00, 0x12, 0x4a, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x74, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x42, 0x02, 0x68, 0x00, 0x12, 0x30, 0x0a, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10,
	0x6c, 0x61, 0x73, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42,
	0x02, 0x68, 0x00, 0x12, 0x31, 0x0a, 0x13, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x10, 0x6c, 0x61, 0x73, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x42, 0x02, 0x68, 0x00, 0x12, 0x31, 0x0a, 0x13, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x10, 0x6c, 0x61, 0x73, 0x74, 0x46, 0x69, 0x72, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x42, 0x02, 0x68, 0x00, 0x12, 0x26, 0x0a, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x42, 0x02, 0x68, 0x00, 0x12, 0x39, 0x0a, 0x17, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x6c, 0x61, 0x73, 0x74, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x42, 0x02,
	0x68, 0x00, 0x12, 0x25, 0x0a, 0x0c, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x42, 0x02, 0x68, 0x00, 0x12, 0x29, 0x0a, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x02, 0x68, 0x00, 0x12,
	0x36, 0x0a, 0x15, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x63, 0x68,
	0x69, 0x6c, 0x64, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x42, 0x02, 0x68, 0x00, 0x12, 0x2c, 0x0a, 0x10, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x17, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x75,
	0x73, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x02, 0x68, 0x00,
	0x12, 0x45, 0x0a, 0x1d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x18, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42,
	0x02, 0x68, 0x00, 0x12, 0x36, 0x0a, 0x15, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x19, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x13, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x02, 0x68, 0x00, 0x12, 0x36, 0x0a, 0x15, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x54,
	0x61, 0x73, 0x6b, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x42, 0x02, 0x68, 0x00, 0x12, 0x36, 0x0a,
	0x15, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x02,
	0x68, 0x00, 0x12, 0x4a, 0x0a, 0x20, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1c, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x42, 0x02, 0x68, 0x00, 0x12, 0x46, 0x0a, 0x1e, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x65, 0x64, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x1a, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x42, 0x02, 0x68, 0x00, 0x12, 0x48,
	0x0a, 0x1f, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x5f, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x0e, 0x20, 0x03, 0x28, 0x03, 0x52, 0x1b, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x73, 0x42, 0x02, 0x68, 0x00, 0x12, 0x52, 0x0a, 0x24, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x64, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x20, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x73,
	0x42, 0x02, 0x68, 0x00, 0x12, 0x4e, 0x0a, 0x22, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x03, 0x52, 0x1e, 0x70,
	0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x42, 0x02, 0x68, 0x00,
	0x12, 0x55, 0x0a, 0x26, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x71, 0x5f, 0x63,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28, 0x03, 0x52, 0x21, 0x70,
	0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x49, 0x6e,
	0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x42, 0x02,
	0x68, 0x00, 0x12, 0x4c, 0x0a, 0x21, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x68, 0x69,
	0x6c, 0x64, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x12, 0x20, 0x03, 0x28, 0x03, 0x52, 0x1d, 0x70, 0x65, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65,
	0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x42, 0x02, 0x68, 0x00, 0x12, 0x37, 0x0a, 0x16,
	0x73, 0x74, 0x69, 0x63, 0x6b, 0x79, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x73, 0x74, 0x69, 0x63,
	0x6b, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x51, 0x75, 0x65, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x02,
	0x68, 0x00, 0x12, 0x61, 0x0a, 0x11, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x74, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x10, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x42, 0x02, 0x68, 0x00,
	0x42, 0x30, 0x5a, 0x2e, 0x67, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69,
	0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x73, 0x75, 0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_server_api_checksum_v1_message_proto_rawDescOnce sync.Once
	file_temporal_server_api_checksum_v1_message_proto_rawDescData = file_temporal_server_api_checksum_v1_message_proto_rawDesc
)

func file_temporal_server_api_checksum_v1_message_proto_rawDescGZIP() []byte {
	file_temporal_server_api_checksum_v1_message_proto_rawDescOnce.Do(func() {
		file_temporal_server_api_checksum_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_server_api_checksum_v1_message_proto_rawDescData)
	})
	return file_temporal_server_api_checksum_v1_message_proto_rawDescData
}

var file_temporal_server_api_checksum_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_temporal_server_api_checksum_v1_message_proto_goTypes = []interface{}{
	(*MutableStateChecksumPayload)(nil), // 0: temporal.server.api.checksum.v1.MutableStateChecksumPayload
	(v1.WorkflowExecutionState)(0),      // 1: temporal.server.api.enums.v1.WorkflowExecutionState
	(v11.WorkflowExecutionStatus)(0),    // 2: temporal.api.enums.v1.WorkflowExecutionStatus
	(*v12.VersionHistories)(nil),        // 3: temporal.server.api.history.v1.VersionHistories
}
var file_temporal_server_api_checksum_v1_message_proto_depIdxs = []int32{
	1, // 0: temporal.server.api.checksum.v1.MutableStateChecksumPayload.state:type_name -> temporal.server.api.enums.v1.WorkflowExecutionState
	2, // 1: temporal.server.api.checksum.v1.MutableStateChecksumPayload.status:type_name -> temporal.api.enums.v1.WorkflowExecutionStatus
	3, // 2: temporal.server.api.checksum.v1.MutableStateChecksumPayload.version_histories:type_name -> temporal.server.api.history.v1.VersionHistories
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_temporal_server_api_checksum_v1_message_proto_init() }
func file_temporal_server_api_checksum_v1_message_proto_init() {
	if File_temporal_server_api_checksum_v1_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_temporal_server_api_checksum_v1_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutableStateChecksumPayload); i {
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
			RawDescriptor: file_temporal_server_api_checksum_v1_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_server_api_checksum_v1_message_proto_goTypes,
		DependencyIndexes: file_temporal_server_api_checksum_v1_message_proto_depIdxs,
		MessageInfos:      file_temporal_server_api_checksum_v1_message_proto_msgTypes,
	}.Build()
	File_temporal_server_api_checksum_v1_message_proto = out.File
	file_temporal_server_api_checksum_v1_message_proto_rawDesc = nil
	file_temporal_server_api_checksum_v1_message_proto_goTypes = nil
	file_temporal_server_api_checksum_v1_message_proto_depIdxs = nil
}
