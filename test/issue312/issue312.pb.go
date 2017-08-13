// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: issue312.proto

/*
Package issue312 is a generated protocol buffer package.

It is generated from these files:
	issue312.proto

It has these top-level messages:
*/
package issue312

import proto "github.com/vaniila/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/vaniila/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type TaskState int32

const (
	TaskState_TASK_STAGING  TaskState = 6
	TaskState_TASK_STARTING TaskState = 0
	TaskState_TASK_RUNNING  TaskState = 1
)

var TaskState_name = map[int32]string{
	6: "TASK_STAGING",
	0: "TASK_STARTING",
	1: "TASK_RUNNING",
}
var TaskState_value = map[string]int32{
	"TASK_STAGING":  6,
	"TASK_STARTING": 0,
	"TASK_RUNNING":  1,
}

func (x TaskState) Enum() *TaskState {
	p := new(TaskState)
	*p = x
	return p
}
func (x TaskState) String() string {
	return proto.EnumName(TaskState_name, int32(x))
}
func (x *TaskState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TaskState_value, data, "TaskState")
	if err != nil {
		return err
	}
	*x = TaskState(value)
	return nil
}
func (TaskState) EnumDescriptor() ([]byte, []int) { return fileDescriptorIssue312, []int{0} }

func init() {
	proto.RegisterEnum("issue312.TaskState", TaskState_name, TaskState_value)
}

func init() { proto.RegisterFile("issue312.proto", fileDescriptorIssue312) }

var fileDescriptorIssue312 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x2c, 0x2e, 0x2e,
	0x4d, 0x35, 0x36, 0x34, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0xa5, 0x0c,
	0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xcb, 0x12, 0xf3, 0x32, 0x33,
	0x73, 0x12, 0xf5, 0xc1, 0x6a, 0x92, 0x4a, 0xd3, 0xf4, 0xd3, 0xf3, 0xd3, 0xf3, 0xc1, 0x1c, 0x30,
	0x0b, 0xa2, 0x57, 0xcb, 0x89, 0x8b, 0x33, 0x24, 0xb1, 0x38, 0x3b, 0xb8, 0x24, 0xb1, 0x24, 0x55,
	0x48, 0x80, 0x8b, 0x27, 0xc4, 0x31, 0xd8, 0x3b, 0x3e, 0x38, 0xc4, 0xd1, 0xdd, 0xd3, 0xcf, 0x5d,
	0x80, 0x4d, 0x48, 0x90, 0x8b, 0x17, 0x26, 0x12, 0x14, 0x02, 0x12, 0x62, 0x80, 0x2b, 0x0a, 0x0a,
	0xf5, 0xf3, 0x03, 0x89, 0x30, 0x3a, 0x49, 0x7d, 0x78, 0x28, 0xc7, 0xf8, 0xe3, 0xa1, 0x1c, 0xe3,
	0x8a, 0x47, 0x72, 0x8c, 0x3b, 0x1e, 0xc9, 0x31, 0x46, 0xc1, 0x5d, 0x04, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x09, 0xbe, 0x7d, 0x40, 0xac, 0x00, 0x00, 0x00,
}
