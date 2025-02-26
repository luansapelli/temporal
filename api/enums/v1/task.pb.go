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
// source: temporal/server/api/enums/v1/task.proto

package enums

import (
	reflect "reflect"
	"strconv"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// TaskSource is the source from which a task was produced.
type TaskSource int32

const (
	TASK_SOURCE_UNSPECIFIED TaskSource = 0
	// Task produced by history service.
	TASK_SOURCE_HISTORY TaskSource = 1
	// Task produced from matching db backlog.
	TASK_SOURCE_DB_BACKLOG TaskSource = 2
)

// Enum value maps for TaskSource.
var (
	TaskSource_name = map[int32]string{
		0: "TASK_SOURCE_UNSPECIFIED",
		1: "TASK_SOURCE_HISTORY",
		2: "TASK_SOURCE_DB_BACKLOG",
	}
	TaskSource_value = map[string]int32{
		"TASK_SOURCE_UNSPECIFIED": 0,
		"TASK_SOURCE_HISTORY":     1,
		"TASK_SOURCE_DB_BACKLOG":  2,
	}
)

func (x TaskSource) Enum() *TaskSource {
	p := new(TaskSource)
	*p = x
	return p
}

func (x TaskSource) String() string {
	switch x {
	case TASK_SOURCE_UNSPECIFIED:
		return "Unspecified"
	case TASK_SOURCE_HISTORY:
		return "History"
	case TASK_SOURCE_DB_BACKLOG:
		return "DbBacklog"
	default:
		return strconv.Itoa(int(x))
	}

}

func (TaskSource) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_server_api_enums_v1_task_proto_enumTypes[0].Descriptor()
}

func (TaskSource) Type() protoreflect.EnumType {
	return &file_temporal_server_api_enums_v1_task_proto_enumTypes[0]
}

func (x TaskSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskSource.Descriptor instead.
func (TaskSource) EnumDescriptor() ([]byte, []int) {
	return file_temporal_server_api_enums_v1_task_proto_rawDescGZIP(), []int{0}
}

type TaskType int32

const (
	TASK_TYPE_UNSPECIFIED                     TaskType = 0
	TASK_TYPE_REPLICATION_HISTORY             TaskType = 1
	TASK_TYPE_REPLICATION_SYNC_ACTIVITY       TaskType = 2
	TASK_TYPE_TRANSFER_WORKFLOW_TASK          TaskType = 3
	TASK_TYPE_TRANSFER_ACTIVITY_TASK          TaskType = 4
	TASK_TYPE_TRANSFER_CLOSE_EXECUTION        TaskType = 5
	TASK_TYPE_TRANSFER_CANCEL_EXECUTION       TaskType = 6
	TASK_TYPE_TRANSFER_START_CHILD_EXECUTION  TaskType = 7
	TASK_TYPE_TRANSFER_SIGNAL_EXECUTION       TaskType = 8
	TASK_TYPE_TRANSFER_RESET_WORKFLOW         TaskType = 10
	TASK_TYPE_WORKFLOW_TASK_TIMEOUT           TaskType = 12
	TASK_TYPE_ACTIVITY_TIMEOUT                TaskType = 13
	TASK_TYPE_USER_TIMER                      TaskType = 14
	TASK_TYPE_WORKFLOW_RUN_TIMEOUT            TaskType = 15
	TASK_TYPE_DELETE_HISTORY_EVENT            TaskType = 16
	TASK_TYPE_ACTIVITY_RETRY_TIMER            TaskType = 17
	TASK_TYPE_WORKFLOW_BACKOFF_TIMER          TaskType = 18
	TASK_TYPE_VISIBILITY_START_EXECUTION      TaskType = 19
	TASK_TYPE_VISIBILITY_UPSERT_EXECUTION     TaskType = 20
	TASK_TYPE_VISIBILITY_CLOSE_EXECUTION      TaskType = 21
	TASK_TYPE_VISIBILITY_DELETE_EXECUTION     TaskType = 22
	TASK_TYPE_TRANSFER_DELETE_EXECUTION       TaskType = 24
	TASK_TYPE_REPLICATION_SYNC_WORKFLOW_STATE TaskType = 25
	TASK_TYPE_ARCHIVAL_ARCHIVE_EXECUTION      TaskType = 26
	TASK_TYPE_CALLBACK                        TaskType = 27
	TASK_TYPE_CALLBACK_BACKOFF                TaskType = 28
)

// Enum value maps for TaskType.
var (
	TaskType_name = map[int32]string{
		0:  "TASK_TYPE_UNSPECIFIED",
		1:  "TASK_TYPE_REPLICATION_HISTORY",
		2:  "TASK_TYPE_REPLICATION_SYNC_ACTIVITY",
		3:  "TASK_TYPE_TRANSFER_WORKFLOW_TASK",
		4:  "TASK_TYPE_TRANSFER_ACTIVITY_TASK",
		5:  "TASK_TYPE_TRANSFER_CLOSE_EXECUTION",
		6:  "TASK_TYPE_TRANSFER_CANCEL_EXECUTION",
		7:  "TASK_TYPE_TRANSFER_START_CHILD_EXECUTION",
		8:  "TASK_TYPE_TRANSFER_SIGNAL_EXECUTION",
		10: "TASK_TYPE_TRANSFER_RESET_WORKFLOW",
		12: "TASK_TYPE_WORKFLOW_TASK_TIMEOUT",
		13: "TASK_TYPE_ACTIVITY_TIMEOUT",
		14: "TASK_TYPE_USER_TIMER",
		15: "TASK_TYPE_WORKFLOW_RUN_TIMEOUT",
		16: "TASK_TYPE_DELETE_HISTORY_EVENT",
		17: "TASK_TYPE_ACTIVITY_RETRY_TIMER",
		18: "TASK_TYPE_WORKFLOW_BACKOFF_TIMER",
		19: "TASK_TYPE_VISIBILITY_START_EXECUTION",
		20: "TASK_TYPE_VISIBILITY_UPSERT_EXECUTION",
		21: "TASK_TYPE_VISIBILITY_CLOSE_EXECUTION",
		22: "TASK_TYPE_VISIBILITY_DELETE_EXECUTION",
		24: "TASK_TYPE_TRANSFER_DELETE_EXECUTION",
		25: "TASK_TYPE_REPLICATION_SYNC_WORKFLOW_STATE",
		26: "TASK_TYPE_ARCHIVAL_ARCHIVE_EXECUTION",
		27: "TASK_TYPE_CALLBACK",
		28: "TASK_TYPE_CALLBACK_BACKOFF",
	}
	TaskType_value = map[string]int32{
		"TASK_TYPE_UNSPECIFIED":                     0,
		"TASK_TYPE_REPLICATION_HISTORY":             1,
		"TASK_TYPE_REPLICATION_SYNC_ACTIVITY":       2,
		"TASK_TYPE_TRANSFER_WORKFLOW_TASK":          3,
		"TASK_TYPE_TRANSFER_ACTIVITY_TASK":          4,
		"TASK_TYPE_TRANSFER_CLOSE_EXECUTION":        5,
		"TASK_TYPE_TRANSFER_CANCEL_EXECUTION":       6,
		"TASK_TYPE_TRANSFER_START_CHILD_EXECUTION":  7,
		"TASK_TYPE_TRANSFER_SIGNAL_EXECUTION":       8,
		"TASK_TYPE_TRANSFER_RESET_WORKFLOW":         10,
		"TASK_TYPE_WORKFLOW_TASK_TIMEOUT":           12,
		"TASK_TYPE_ACTIVITY_TIMEOUT":                13,
		"TASK_TYPE_USER_TIMER":                      14,
		"TASK_TYPE_WORKFLOW_RUN_TIMEOUT":            15,
		"TASK_TYPE_DELETE_HISTORY_EVENT":            16,
		"TASK_TYPE_ACTIVITY_RETRY_TIMER":            17,
		"TASK_TYPE_WORKFLOW_BACKOFF_TIMER":          18,
		"TASK_TYPE_VISIBILITY_START_EXECUTION":      19,
		"TASK_TYPE_VISIBILITY_UPSERT_EXECUTION":     20,
		"TASK_TYPE_VISIBILITY_CLOSE_EXECUTION":      21,
		"TASK_TYPE_VISIBILITY_DELETE_EXECUTION":     22,
		"TASK_TYPE_TRANSFER_DELETE_EXECUTION":       24,
		"TASK_TYPE_REPLICATION_SYNC_WORKFLOW_STATE": 25,
		"TASK_TYPE_ARCHIVAL_ARCHIVE_EXECUTION":      26,
		"TASK_TYPE_CALLBACK":                        27,
		"TASK_TYPE_CALLBACK_BACKOFF":                28,
	}
)

func (x TaskType) Enum() *TaskType {
	p := new(TaskType)
	*p = x
	return p
}

func (x TaskType) String() string {
	switch x {
	case TASK_TYPE_UNSPECIFIED:
		return "Unspecified"
	case TASK_TYPE_REPLICATION_HISTORY:
		return "ReplicationHistory"
	case TASK_TYPE_REPLICATION_SYNC_ACTIVITY:
		return "ReplicationSyncActivity"
	case TASK_TYPE_TRANSFER_WORKFLOW_TASK:
		return "TransferWorkflowTask"
	case TASK_TYPE_TRANSFER_ACTIVITY_TASK:
		return "TransferActivityTask"
	case TASK_TYPE_TRANSFER_CLOSE_EXECUTION:
		return "TransferCloseExecution"
	case TASK_TYPE_TRANSFER_CANCEL_EXECUTION:
		return "TransferCancelExecution"

		// Deprecated: Use TaskType.Descriptor instead.
	case TASK_TYPE_TRANSFER_START_CHILD_EXECUTION:
		return "TransferStartChildExecution"
	case TASK_TYPE_TRANSFER_SIGNAL_EXECUTION:
		return "TransferSignalExecution"
	case TASK_TYPE_TRANSFER_RESET_WORKFLOW:
		return "TransferResetWorkflow"
	case TASK_TYPE_WORKFLOW_TASK_TIMEOUT:
		return "WorkflowTaskTimeout"
	case TASK_TYPE_ACTIVITY_TIMEOUT:
		return "ActivityTimeout"
	case TASK_TYPE_USER_TIMER:
		return "UserTimer"
	case TASK_TYPE_WORKFLOW_RUN_TIMEOUT:
		return "WorkflowRunTimeout"
	case TASK_TYPE_DELETE_HISTORY_EVENT:
		return "DeleteHistoryEvent"
	case TASK_TYPE_ACTIVITY_RETRY_TIMER:
		return "ActivityRetryTimer"
	case TASK_TYPE_WORKFLOW_BACKOFF_TIMER:
		return "WorkflowBackoffTimer"
	case TASK_TYPE_VISIBILITY_START_EXECUTION:
		return "VisibilityStartExecution"
	case TASK_TYPE_VISIBILITY_UPSERT_EXECUTION:
		return "VisibilityUpsertExecution"
	case TASK_TYPE_VISIBILITY_CLOSE_EXECUTION:
		return "VisibilityCloseExecution"
	case TASK_TYPE_VISIBILITY_DELETE_EXECUTION:
		return "VisibilityDeleteExecution"
	case TASK_TYPE_TRANSFER_DELETE_EXECUTION:
		return "TransferDeleteExecution"
	case TASK_TYPE_REPLICATION_SYNC_WORKFLOW_STATE:
		return "ReplicationSyncWorkflowState"
	case TASK_TYPE_ARCHIVAL_ARCHIVE_EXECUTION:
		return "ArchivalArchiveExecution"
	case TASK_TYPE_CALLBACK:
		return "Callback"
	case TASK_TYPE_CALLBACK_BACKOFF:
		return "CallbackBackoff"
	default:
		return strconv.Itoa(int(x))
	}

}

func (TaskType) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_server_api_enums_v1_task_proto_enumTypes[1].Descriptor()
}

func (TaskType) Type() protoreflect.EnumType {
	return &file_temporal_server_api_enums_v1_task_proto_enumTypes[1]
}

func (x TaskType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

func (TaskType) EnumDescriptor() ([]byte, []int) {
	return file_temporal_server_api_enums_v1_task_proto_rawDescGZIP(), []int{1}
}

var File_temporal_server_api_enums_v1_task_proto protoreflect.FileDescriptor

var file_temporal_server_api_enums_v1_task_proto_rawDesc = []byte{
	0x0a, 0x27, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x74, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2a, 0x5e, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x53, 0x4f,
	0x55, 0x52, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43,
	0x45, 0x5f, 0x48, 0x49, 0x53, 0x54, 0x4f, 0x52, 0x59, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x54,
	0x41, 0x53, 0x4b, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x44, 0x42, 0x5f, 0x42, 0x41,
	0x43, 0x4b, 0x4c, 0x4f, 0x47, 0x10, 0x02, 0x2a, 0xf3, 0x07, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x21, 0x0a, 0x1d, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x50,
	0x4c, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x48, 0x49, 0x53, 0x54, 0x4f, 0x52, 0x59,
	0x10, 0x01, 0x12, 0x27, 0x0a, 0x23, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x52, 0x45, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x59, 0x4e, 0x43,
	0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x10, 0x02, 0x12, 0x24, 0x0a, 0x20, 0x54,
	0x41, 0x53, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45,
	0x52, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x10,
	0x03, 0x12, 0x24, 0x0a, 0x20, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54,
	0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59,
	0x5f, 0x54, 0x41, 0x53, 0x4b, 0x10, 0x04, 0x12, 0x26, 0x0a, 0x22, 0x54, 0x41, 0x53, 0x4b, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x43, 0x4c,
	0x4f, 0x53, 0x45, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x05, 0x12,
	0x27, 0x0a, 0x23, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x52, 0x41,
	0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x5f, 0x45, 0x58, 0x45,
	0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x06, 0x12, 0x2c, 0x0a, 0x28, 0x54, 0x41, 0x53, 0x4b,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x53,
	0x54, 0x41, 0x52, 0x54, 0x5f, 0x43, 0x48, 0x49, 0x4c, 0x44, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x07, 0x12, 0x27, 0x0a, 0x23, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x53, 0x49, 0x47,
	0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x08, 0x12,
	0x25, 0x0a, 0x21, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x52, 0x41,
	0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x53, 0x45, 0x54, 0x5f, 0x57, 0x4f, 0x52, 0x4b,
	0x46, 0x4c, 0x4f, 0x57, 0x10, 0x0a, 0x12, 0x23, 0x0a, 0x1f, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x54, 0x41, 0x53,
	0x4b, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x0c, 0x12, 0x1e, 0x0a, 0x1a, 0x54,
	0x41, 0x53, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54,
	0x59, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x0d, 0x12, 0x18, 0x0a, 0x14, 0x54,
	0x41, 0x53, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x54, 0x49,
	0x4d, 0x45, 0x52, 0x10, 0x0e, 0x12, 0x22, 0x0a, 0x1e, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x52, 0x55, 0x4e, 0x5f,
	0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x0f, 0x12, 0x22, 0x0a, 0x1e, 0x54, 0x41, 0x53,
	0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x48, 0x49,
	0x53, 0x54, 0x4f, 0x52, 0x59, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x10, 0x10, 0x12, 0x22, 0x0a,
	0x1e, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56,
	0x49, 0x54, 0x59, 0x5f, 0x52, 0x45, 0x54, 0x52, 0x59, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x52, 0x10,
	0x11, 0x12, 0x24, 0x0a, 0x20, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57,
	0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x4f, 0x46, 0x46, 0x5f,
	0x54, 0x49, 0x4d, 0x45, 0x52, 0x10, 0x12, 0x12, 0x28, 0x0a, 0x24, 0x54, 0x41, 0x53, 0x4b, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x56, 0x49, 0x53, 0x49, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f,
	0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0x13, 0x12, 0x29, 0x0a, 0x25, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x56,
	0x49, 0x53, 0x49, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x50, 0x53, 0x45, 0x52, 0x54,
	0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x14, 0x12, 0x28, 0x0a, 0x24,
	0x54, 0x41, 0x53, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x56, 0x49, 0x53, 0x49, 0x42, 0x49,
	0x4c, 0x49, 0x54, 0x59, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x15, 0x12, 0x29, 0x0a, 0x25, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x56, 0x49, 0x53, 0x49, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x44,
	0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0x16, 0x12, 0x27, 0x0a, 0x23, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54,
	0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x45,
	0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x18, 0x12, 0x2d, 0x0a, 0x29, 0x54, 0x41,
	0x53, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f,
	0x57, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x10, 0x19, 0x12, 0x28, 0x0a, 0x24, 0x54, 0x41, 0x53,
	0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x52, 0x43, 0x48, 0x49, 0x56, 0x41, 0x4c, 0x5f,
	0x41, 0x52, 0x43, 0x48, 0x49, 0x56, 0x45, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f,
	0x4e, 0x10, 0x1a, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x42, 0x41, 0x43, 0x4b, 0x10, 0x1b, 0x12, 0x1e, 0x0a, 0x1a, 0x54,
	0x41, 0x53, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x42, 0x41, 0x43,
	0x4b, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x4f, 0x46, 0x46, 0x10, 0x1c, 0x22, 0x04, 0x08, 0x09, 0x10,
	0x09, 0x22, 0x04, 0x08, 0x0b, 0x10, 0x0b, 0x22, 0x04, 0x08, 0x17, 0x10, 0x17, 0x42, 0x2a, 0x5a,
	0x28, 0x67, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x76, 0x31, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_temporal_server_api_enums_v1_task_proto_rawDescOnce sync.Once
	file_temporal_server_api_enums_v1_task_proto_rawDescData = file_temporal_server_api_enums_v1_task_proto_rawDesc
)

func file_temporal_server_api_enums_v1_task_proto_rawDescGZIP() []byte {
	file_temporal_server_api_enums_v1_task_proto_rawDescOnce.Do(func() {
		file_temporal_server_api_enums_v1_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_server_api_enums_v1_task_proto_rawDescData)
	})
	return file_temporal_server_api_enums_v1_task_proto_rawDescData
}

var file_temporal_server_api_enums_v1_task_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_temporal_server_api_enums_v1_task_proto_goTypes = []interface{}{
	(TaskSource)(0), // 0: temporal.server.api.enums.v1.TaskSource
	(TaskType)(0),   // 1: temporal.server.api.enums.v1.TaskType
}
var file_temporal_server_api_enums_v1_task_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_temporal_server_api_enums_v1_task_proto_init() }
func file_temporal_server_api_enums_v1_task_proto_init() {
	if File_temporal_server_api_enums_v1_task_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_temporal_server_api_enums_v1_task_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_server_api_enums_v1_task_proto_goTypes,
		DependencyIndexes: file_temporal_server_api_enums_v1_task_proto_depIdxs,
		EnumInfos:         file_temporal_server_api_enums_v1_task_proto_enumTypes,
	}.Build()
	File_temporal_server_api_enums_v1_task_proto = out.File
	file_temporal_server_api_enums_v1_task_proto_rawDesc = nil
	file_temporal_server_api_enums_v1_task_proto_goTypes = nil
	file_temporal_server_api_enums_v1_task_proto_depIdxs = nil
}
