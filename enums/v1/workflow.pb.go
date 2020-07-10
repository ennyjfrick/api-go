// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
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

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/api/enums/v1/workflow.proto

package enums

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
	strconv "strconv"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type WorkflowIdReusePolicy int32

const (
	WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED WorkflowIdReusePolicy = 0
	// Allow start a workflow execution using the same workflow Id, when workflow not running.
	WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE WorkflowIdReusePolicy = 1
	// Allow start a workflow execution using the same workflow Id, when workflow not running, and the last execution close state is in
	// [terminated, cancelled, timed out, failed].
	WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY WorkflowIdReusePolicy = 2
	// Do not allow start a workflow execution using the same workflow Id at all.
	WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE WorkflowIdReusePolicy = 3
)

var WorkflowIdReusePolicy_name = map[int32]string{
	0: "Unspecified",
	1: "AllowDuplicate",
	2: "AllowDuplicateFailedOnly",
	3: "RejectDuplicate",
}

var WorkflowIdReusePolicy_value = map[string]int32{
	"Unspecified":              0,
	"AllowDuplicate":           1,
	"AllowDuplicateFailedOnly": 2,
	"RejectDuplicate":          3,
}

func (WorkflowIdReusePolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_939fa9511cc117f0, []int{0}
}

type ParentClosePolicy int32

const (
	PARENT_CLOSE_POLICY_UNSPECIFIED ParentClosePolicy = 0
	// Terminate means terminating the child workflow.
	PARENT_CLOSE_POLICY_TERMINATE ParentClosePolicy = 1
	// Abandon means not doing anything on the child workflow.
	PARENT_CLOSE_POLICY_ABANDON ParentClosePolicy = 2
	// Cancel means requesting cancellation on the child workflow.
	PARENT_CLOSE_POLICY_REQUEST_CANCEL ParentClosePolicy = 3
)

var ParentClosePolicy_name = map[int32]string{
	0: "Unspecified",
	1: "Terminate",
	2: "Abandon",
	3: "RequestCancel",
}

var ParentClosePolicy_value = map[string]int32{
	"Unspecified":   0,
	"Terminate":     1,
	"Abandon":       2,
	"RequestCancel": 3,
}

func (ParentClosePolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_939fa9511cc117f0, []int{1}
}

type ContinueAsNewInitiator int32

const (
	CONTINUE_AS_NEW_INITIATOR_UNSPECIFIED   ContinueAsNewInitiator = 0
	CONTINUE_AS_NEW_INITIATOR_DECIDER       ContinueAsNewInitiator = 1
	CONTINUE_AS_NEW_INITIATOR_RETRY         ContinueAsNewInitiator = 2
	CONTINUE_AS_NEW_INITIATOR_CRON_SCHEDULE ContinueAsNewInitiator = 3
)

var ContinueAsNewInitiator_name = map[int32]string{
	0: "Unspecified",
	1: "Decider",
	2: "Retry",
	3: "CronSchedule",
}

var ContinueAsNewInitiator_value = map[string]int32{
	"Unspecified":  0,
	"Decider":      1,
	"Retry":        2,
	"CronSchedule": 3,
}

func (ContinueAsNewInitiator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_939fa9511cc117f0, []int{2}
}

// (-- api-linter: core::0216::synonyms=disabled
//     aip.dev/not-precedent: There is WorkflowExecutionState already in another package. --)
type WorkflowExecutionStatus int32

const (
	WORKFLOW_EXECUTION_STATUS_UNSPECIFIED WorkflowExecutionStatus = 0
	// Value 1 is hardcoded in SQL persistence.
	WORKFLOW_EXECUTION_STATUS_RUNNING          WorkflowExecutionStatus = 1
	WORKFLOW_EXECUTION_STATUS_COMPLETED        WorkflowExecutionStatus = 2
	WORKFLOW_EXECUTION_STATUS_FAILED           WorkflowExecutionStatus = 3
	WORKFLOW_EXECUTION_STATUS_CANCELED         WorkflowExecutionStatus = 4
	WORKFLOW_EXECUTION_STATUS_TERMINATED       WorkflowExecutionStatus = 5
	WORKFLOW_EXECUTION_STATUS_CONTINUED_AS_NEW WorkflowExecutionStatus = 6
	WORKFLOW_EXECUTION_STATUS_TIMED_OUT        WorkflowExecutionStatus = 7
)

var WorkflowExecutionStatus_name = map[int32]string{
	0: "Unspecified",
	1: "Running",
	2: "Completed",
	3: "Failed",
	4: "Canceled",
	5: "Terminated",
	6: "ContinuedAsNew",
	7: "TimedOut",
}

var WorkflowExecutionStatus_value = map[string]int32{
	"Unspecified":    0,
	"Running":        1,
	"Completed":      2,
	"Failed":         3,
	"Canceled":       4,
	"Terminated":     5,
	"ContinuedAsNew": 6,
	"TimedOut":       7,
}

func (WorkflowExecutionStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_939fa9511cc117f0, []int{3}
}

type PendingActivityState int32

const (
	PENDING_ACTIVITY_STATE_UNSPECIFIED      PendingActivityState = 0
	PENDING_ACTIVITY_STATE_SCHEDULED        PendingActivityState = 1
	PENDING_ACTIVITY_STATE_STARTED          PendingActivityState = 2
	PENDING_ACTIVITY_STATE_CANCEL_REQUESTED PendingActivityState = 3
)

var PendingActivityState_name = map[int32]string{
	0: "Unspecified",
	1: "Scheduled",
	2: "Started",
	3: "CancelRequested",
}

var PendingActivityState_value = map[string]int32{
	"Unspecified":     0,
	"Scheduled":       1,
	"Started":         2,
	"CancelRequested": 3,
}

func (PendingActivityState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_939fa9511cc117f0, []int{4}
}

type HistoryEventFilterType int32

const (
	HISTORY_EVENT_FILTER_TYPE_UNSPECIFIED HistoryEventFilterType = 0
	HISTORY_EVENT_FILTER_TYPE_ALL_EVENT   HistoryEventFilterType = 1
	HISTORY_EVENT_FILTER_TYPE_CLOSE_EVENT HistoryEventFilterType = 2
)

var HistoryEventFilterType_name = map[int32]string{
	0: "Unspecified",
	1: "AllEvent",
	2: "CloseEvent",
}

var HistoryEventFilterType_value = map[string]int32{
	"Unspecified": 0,
	"AllEvent":    1,
	"CloseEvent":  2,
}

func (HistoryEventFilterType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_939fa9511cc117f0, []int{5}
}

type RetryState int32

const (
	RETRY_STATE_UNSPECIFIED              RetryState = 0
	RETRY_STATE_IN_PROGRESS              RetryState = 1
	RETRY_STATE_NON_RETRYABLE_FAILURE    RetryState = 2
	RETRY_STATE_TIMEOUT                  RetryState = 3
	RETRY_STATE_MAXIMUM_ATTEMPTS_REACHED RetryState = 4
	RETRY_STATE_RETRY_POLICY_NOT_SET     RetryState = 5
	RETRY_STATE_INTERNAL_SERVER_ERROR    RetryState = 6
	RETRY_STATE_CANCEL_REQUESTED         RetryState = 7
)

var RetryState_name = map[int32]string{
	0: "Unspecified",
	1: "InProgress",
	2: "NonRetryableFailure",
	3: "Timeout",
	4: "MaximumAttemptsReached",
	5: "RetryPolicyNotSet",
	6: "InternalServerError",
	7: "CancelRequested",
}

var RetryState_value = map[string]int32{
	"Unspecified":            0,
	"InProgress":             1,
	"NonRetryableFailure":    2,
	"Timeout":                3,
	"MaximumAttemptsReached": 4,
	"RetryPolicyNotSet":      5,
	"InternalServerError":    6,
	"CancelRequested":        7,
}

func (RetryState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_939fa9511cc117f0, []int{6}
}

type TimeoutType int32

const (
	TIMEOUT_TYPE_UNSPECIFIED       TimeoutType = 0
	TIMEOUT_TYPE_START_TO_CLOSE    TimeoutType = 1
	TIMEOUT_TYPE_SCHEDULE_TO_START TimeoutType = 2
	TIMEOUT_TYPE_SCHEDULE_TO_CLOSE TimeoutType = 3
	TIMEOUT_TYPE_HEARTBEAT         TimeoutType = 4
)

var TimeoutType_name = map[int32]string{
	0: "Unspecified",
	1: "StartToClose",
	2: "ScheduleToStart",
	3: "ScheduleToClose",
	4: "Heartbeat",
}

var TimeoutType_value = map[string]int32{
	"Unspecified":     0,
	"StartToClose":    1,
	"ScheduleToStart": 2,
	"ScheduleToClose": 3,
	"Heartbeat":       4,
}

func (TimeoutType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_939fa9511cc117f0, []int{7}
}

func init() {
	proto.RegisterEnum("temporal.api.enums.v1.WorkflowIdReusePolicy", WorkflowIdReusePolicy_name, WorkflowIdReusePolicy_value)
	proto.RegisterEnum("temporal.api.enums.v1.ParentClosePolicy", ParentClosePolicy_name, ParentClosePolicy_value)
	proto.RegisterEnum("temporal.api.enums.v1.ContinueAsNewInitiator", ContinueAsNewInitiator_name, ContinueAsNewInitiator_value)
	proto.RegisterEnum("temporal.api.enums.v1.WorkflowExecutionStatus", WorkflowExecutionStatus_name, WorkflowExecutionStatus_value)
	proto.RegisterEnum("temporal.api.enums.v1.PendingActivityState", PendingActivityState_name, PendingActivityState_value)
	proto.RegisterEnum("temporal.api.enums.v1.HistoryEventFilterType", HistoryEventFilterType_name, HistoryEventFilterType_value)
	proto.RegisterEnum("temporal.api.enums.v1.RetryState", RetryState_name, RetryState_value)
	proto.RegisterEnum("temporal.api.enums.v1.TimeoutType", TimeoutType_name, TimeoutType_value)
}

func init() {
	proto.RegisterFile("temporal/api/enums/v1/workflow.proto", fileDescriptor_939fa9511cc117f0)
}

var fileDescriptor_939fa9511cc117f0 = []byte{
	// 883 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0x41, 0x6f, 0xe3, 0x44,
	0x14, 0xc7, 0x33, 0x49, 0xb7, 0x2b, 0x0d, 0x42, 0x1a, 0x0c, 0xdb, 0xae, 0xe8, 0xe2, 0xb6, 0xdb,
	0x2e, 0xdd, 0xcd, 0x42, 0xaa, 0x15, 0x1c, 0x90, 0x38, 0x4d, 0xec, 0x97, 0x76, 0xc0, 0x19, 0x9b,
	0xf1, 0xb8, 0xdd, 0x70, 0x19, 0x85, 0xae, 0x59, 0x59, 0xa4, 0x76, 0xe4, 0x38, 0x2d, 0xbd, 0x21,
	0x3e, 0x01, 0x37, 0x8e, 0x5c, 0x38, 0x70, 0xe4, 0xc2, 0x8d, 0x0f, 0xc0, 0xb1, 0xc7, 0x1e, 0x69,
	0x2a, 0x24, 0x8e, 0xfb, 0x11, 0xd0, 0xd8, 0x49, 0x95, 0xa4, 0x71, 0xc4, 0x2d, 0xf2, 0xfc, 0x66,
	0xf2, 0xde, 0xff, 0xff, 0x7f, 0x33, 0x78, 0x37, 0x0b, 0x4f, 0xfb, 0x49, 0xda, 0xed, 0xed, 0x77,
	0xfb, 0xd1, 0x7e, 0x18, 0x0f, 0x4f, 0x07, 0xfb, 0x67, 0x2f, 0xf6, 0xcf, 0x93, 0xf4, 0xbb, 0x6f,
	0x7b, 0xc9, 0x79, 0xa3, 0x9f, 0x26, 0x59, 0x62, 0x3c, 0x98, 0x50, 0x8d, 0x6e, 0x3f, 0x6a, 0xe4,
	0x54, 0xe3, 0xec, 0x45, 0xfd, 0x0a, 0xe1, 0x07, 0xc7, 0x63, 0x92, 0xbd, 0x12, 0xe1, 0x70, 0x10,
	0x7a, 0x49, 0x2f, 0x3a, 0xb9, 0x30, 0x9e, 0xe2, 0xdd, 0x63, 0x57, 0x7c, 0xd9, 0x72, 0xdc, 0x63,
	0xc5, 0x6c, 0x25, 0x20, 0xf0, 0x41, 0x79, 0xae, 0xc3, 0xac, 0x8e, 0x0a, 0xb8, 0xef, 0x81, 0xc5,
	0x5a, 0x0c, 0x6c, 0x52, 0x31, 0x3e, 0xc2, 0x4f, 0x4b, 0x49, 0xea, 0xe8, 0xaf, 0x76, 0xe0, 0x39,
	0xcc, 0xa2, 0x12, 0x08, 0x32, 0x3e, 0xc3, 0x9f, 0xfe, 0x5f, 0x5a, 0xb5, 0x28, 0x73, 0xc0, 0x56,
	0x2e, 0x77, 0x3a, 0xa4, 0x6a, 0x7c, 0x8c, 0x9f, 0x95, 0xee, 0x14, 0xf0, 0x05, 0x58, 0x72, 0xea,
	0x8f, 0x6a, 0xf5, 0x5f, 0x11, 0x7e, 0xc7, 0xeb, 0xa6, 0x61, 0x9c, 0x59, 0xbd, 0xe4, 0xb6, 0xad,
	0x1d, 0xbc, 0xe9, 0x51, 0x01, 0x5c, 0x2a, 0xcb, 0x71, 0xcb, 0x3a, 0xda, 0xc6, 0x1f, 0x2c, 0x82,
	0x24, 0x88, 0x36, 0xe3, 0x45, 0x1b, 0x9b, 0x78, 0x63, 0x11, 0x42, 0x9b, 0x94, 0xdb, 0x2e, 0x27,
	0x55, 0xe3, 0x43, 0xfc, 0x78, 0x11, 0x20, 0xe0, 0xab, 0x00, 0x7c, 0xa9, 0x2c, 0xca, 0x2d, 0x70,
	0x48, 0xad, 0xfe, 0x27, 0xc2, 0x6b, 0x56, 0x12, 0x67, 0x51, 0x3c, 0x0c, 0xe9, 0x80, 0x87, 0xe7,
	0x2c, 0x8e, 0xb2, 0xa8, 0x9b, 0x25, 0xa9, 0xf1, 0x0c, 0x3f, 0xb1, 0x5c, 0x2e, 0x19, 0x0f, 0x40,
	0x51, 0x5f, 0x71, 0x38, 0x56, 0x8c, 0x33, 0xc9, 0xa8, 0x74, 0xc5, 0x5c, 0xc5, 0x4f, 0xf0, 0x76,
	0x39, 0x6a, 0x83, 0xc5, 0x6c, 0x10, 0x04, 0xe9, 0xee, 0xcb, 0x31, 0x01, 0x52, 0x68, 0x9d, 0x9f,
	0xe3, 0xbd, 0x72, 0xc8, 0x12, 0x2e, 0x57, 0xbe, 0x75, 0x08, 0x76, 0xe0, 0x68, 0x95, 0xff, 0xa9,
	0xe2, 0xf5, 0x49, 0x80, 0xe0, 0xfb, 0xf0, 0x64, 0x98, 0x45, 0x49, 0xec, 0x67, 0xdd, 0x6c, 0x38,
	0xd0, 0xf5, 0xdf, 0x1a, 0x06, 0x2f, 0xc1, 0x0a, 0x24, 0xd3, 0x9b, 0x25, 0x95, 0x81, 0x7f, 0xb7,
	0xfe, 0x72, 0x54, 0x04, 0x9c, 0x33, 0x7e, 0x40, 0x90, 0xb1, 0x87, 0x77, 0xca, 0x31, 0xcb, 0x6d,
	0x7b, 0x0e, 0x48, 0xb0, 0x49, 0xd5, 0xd8, 0xc5, 0x5b, 0xe5, 0x60, 0x11, 0x2b, 0x52, 0xd3, 0x1e,
	0x2d, 0x39, 0x2e, 0x77, 0x08, 0x6c, 0xb2, 0x32, 0x33, 0x0b, 0x77, 0xb8, 0xdb, 0x54, 0xd8, 0xe4,
	0x9e, 0xd1, 0xc0, 0xf5, 0x65, 0x05, 0x16, 0xaa, 0xda, 0x63, 0x59, 0xc9, 0xea, 0xf2, 0x86, 0x24,
	0x6b, 0xeb, 0xf4, 0x07, 0x92, 0xdc, 0xaf, 0xff, 0x81, 0xf0, 0x7b, 0x5e, 0x18, 0xbf, 0x8a, 0xe2,
	0xd7, 0xf4, 0x24, 0x8b, 0xce, 0xa2, 0xec, 0x42, 0xab, 0x1c, 0xe6, 0x39, 0x03, 0x6e, 0x33, 0x7e,
	0xa0, 0xa8, 0x25, 0xd9, 0x11, 0x93, 0x9d, 0x7c, 0x3f, 0xcc, 0x29, 0xbc, 0x8b, 0xb7, 0x4a, 0xb8,
	0x89, 0x9b, 0x36, 0x41, 0xc6, 0x63, 0x6c, 0x96, 0x51, 0x92, 0x8a, 0x42, 0xdb, 0xe7, 0x78, 0xaf,
	0x84, 0x29, 0x24, 0x9b, 0x64, 0x5c, 0x4b, 0x5c, 0xff, 0x19, 0xe1, 0xb5, 0xc3, 0x68, 0x90, 0x25,
	0xe9, 0x05, 0x9c, 0x85, 0x71, 0xd6, 0x8a, 0x7a, 0x59, 0x98, 0xca, 0x8b, 0x7e, 0xa8, 0xe3, 0x71,
	0xc8, 0x7c, 0xe9, 0x8a, 0x8e, 0x82, 0x23, 0x3d, 0x28, 0x2d, 0xe6, 0x48, 0x10, 0x4a, 0x76, 0xbc,
	0xf9, 0xe2, 0xf7, 0xf0, 0x4e, 0x39, 0x4a, 0x1d, 0xa7, 0xf8, 0x4a, 0xd0, 0xf2, 0x33, 0x8b, 0x41,
	0x2c, 0xd0, 0x6a, 0xfd, 0x97, 0x2a, 0xc6, 0x22, 0xcc, 0xd2, 0xb1, 0x8e, 0x1b, 0x78, 0x3d, 0x1f,
	0x80, 0x85, 0xe2, 0xcd, 0x2d, 0x32, 0xae, 0x3c, 0xe1, 0x1e, 0x08, 0xf0, 0x7d, 0x82, 0x74, 0x76,
	0xa7, 0x17, 0xb9, 0xcb, 0x8b, 0x51, 0xa2, 0x4d, 0xa7, 0xb8, 0xc2, 0x02, 0x01, 0xa4, 0x6a, 0xac,
	0xe3, 0x77, 0xa7, 0x31, 0x6d, 0xae, 0xb6, 0xb6, 0xa6, 0xd3, 0x35, 0xbd, 0xd0, 0xa6, 0x2f, 0x59,
	0x3b, 0x68, 0x2b, 0x2a, 0x25, 0xb4, 0x3d, 0xe9, 0x2b, 0x01, 0x54, 0x3b, 0x44, 0x56, 0xb4, 0x87,
	0xd3, 0x64, 0xf1, 0x7b, 0x7c, 0xb1, 0x70, 0x57, 0x2a, 0x1f, 0x24, 0xb9, 0x37, 0x5f, 0x0f, 0xe3,
	0x12, 0x04, 0xa7, 0x8e, 0xf2, 0x41, 0x1c, 0x81, 0x50, 0x20, 0x84, 0x2b, 0xc8, 0xaa, 0xb1, 0x85,
	0x1f, 0x4d, 0x63, 0x77, 0xbc, 0xbb, 0x5f, 0xff, 0x1d, 0xe1, 0xb7, 0x64, 0x74, 0x1a, 0x26, 0xc3,
	0x2c, 0x37, 0xec, 0x11, 0x7e, 0x38, 0xae, 0x7a, 0x91, 0x47, 0x9b, 0x78, 0x63, 0x66, 0x35, 0x0f,
	0x8c, 0x92, 0x6e, 0x21, 0x7b, 0x91, 0xad, 0x59, 0x60, 0x9c, 0x3b, 0xcd, 0xe4, 0x30, 0xa9, 0x2e,
	0x65, 0x8a, 0x73, 0x6a, 0xc6, 0xfb, 0x78, 0x6d, 0x86, 0x39, 0x04, 0x2a, 0x64, 0x13, 0xa8, 0x24,
	0x2b, 0xcd, 0x1f, 0xd1, 0xe5, 0xb5, 0x59, 0xb9, 0xba, 0x36, 0x2b, 0x6f, 0xae, 0x4d, 0xf4, 0xc3,
	0xc8, 0x44, 0xbf, 0x8d, 0x4c, 0xf4, 0xd7, 0xc8, 0x44, 0x97, 0x23, 0x13, 0xfd, 0x3d, 0x32, 0xd1,
	0xbf, 0x23, 0xb3, 0xf2, 0x66, 0x64, 0xa2, 0x9f, 0x6e, 0xcc, 0xca, 0xe5, 0x8d, 0x59, 0xb9, 0xba,
	0x31, 0x2b, 0xf8, 0x61, 0x94, 0x34, 0x16, 0x3e, 0x90, 0xcd, 0xb7, 0x27, 0x97, 0x9b, 0xa7, 0x9f,
	0x51, 0x0f, 0x7d, 0xbd, 0xfd, 0x7a, 0x0a, 0x8d, 0x92, 0x99, 0x47, 0xf7, 0xf3, 0xfc, 0xc7, 0x37,
	0xab, 0xf9, 0x93, 0xfb, 0xc9, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x03, 0x75, 0xa0, 0x9a,
	0x07, 0x00, 0x00,
}

func (x WorkflowIdReusePolicy) String() string {
	s, ok := WorkflowIdReusePolicy_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x ParentClosePolicy) String() string {
	s, ok := ParentClosePolicy_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x ContinueAsNewInitiator) String() string {
	s, ok := ContinueAsNewInitiator_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x WorkflowExecutionStatus) String() string {
	s, ok := WorkflowExecutionStatus_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x PendingActivityState) String() string {
	s, ok := PendingActivityState_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x HistoryEventFilterType) String() string {
	s, ok := HistoryEventFilterType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x RetryState) String() string {
	s, ok := RetryState_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x TimeoutType) String() string {
	s, ok := TimeoutType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
