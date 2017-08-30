// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/types/types.proto

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	pkg/types/types.proto

It has these top-level messages:
	ObjectMetadata
	Workflow
	FunctionInvocation
	FunctionInvocationSpec
	FunctionInvocationStatus
	WorkflowInvocationSpec
	WorkflowInvocationStatus
	WorkflowInvocation
	WorkflowSpec
	Task
	TaskDependencyParameters
	WorkflowStatus
	TaskTypeDef
	TypedValue
*/
package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FunctionInvocationStatus_Status int32

const (
	FunctionInvocationStatus_UNKNOWN     FunctionInvocationStatus_Status = 0
	FunctionInvocationStatus_SCHEDULED   FunctionInvocationStatus_Status = 1
	FunctionInvocationStatus_IN_PROGRESS FunctionInvocationStatus_Status = 2
	FunctionInvocationStatus_SUCCEEDED   FunctionInvocationStatus_Status = 3
	FunctionInvocationStatus_FAILED      FunctionInvocationStatus_Status = 4
	FunctionInvocationStatus_ABORTED     FunctionInvocationStatus_Status = 5
	FunctionInvocationStatus_SKIPPED     FunctionInvocationStatus_Status = 6
)

var FunctionInvocationStatus_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "SCHEDULED",
	2: "IN_PROGRESS",
	3: "SUCCEEDED",
	4: "FAILED",
	5: "ABORTED",
	6: "SKIPPED",
}
var FunctionInvocationStatus_Status_value = map[string]int32{
	"UNKNOWN":     0,
	"SCHEDULED":   1,
	"IN_PROGRESS": 2,
	"SUCCEEDED":   3,
	"FAILED":      4,
	"ABORTED":     5,
	"SKIPPED":     6,
}

func (x FunctionInvocationStatus_Status) String() string {
	return proto.EnumName(FunctionInvocationStatus_Status_name, int32(x))
}
func (FunctionInvocationStatus_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4, 0}
}

type WorkflowInvocationStatus_Status int32

const (
	WorkflowInvocationStatus_UNKNOWN     WorkflowInvocationStatus_Status = 0
	WorkflowInvocationStatus_SCHEDULED   WorkflowInvocationStatus_Status = 1
	WorkflowInvocationStatus_IN_PROGRESS WorkflowInvocationStatus_Status = 2
	WorkflowInvocationStatus_SUCCEEDED   WorkflowInvocationStatus_Status = 3
	WorkflowInvocationStatus_FAILED      WorkflowInvocationStatus_Status = 4
	WorkflowInvocationStatus_ABORTED     WorkflowInvocationStatus_Status = 5
)

var WorkflowInvocationStatus_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "SCHEDULED",
	2: "IN_PROGRESS",
	3: "SUCCEEDED",
	4: "FAILED",
	5: "ABORTED",
}
var WorkflowInvocationStatus_Status_value = map[string]int32{
	"UNKNOWN":     0,
	"SCHEDULED":   1,
	"IN_PROGRESS": 2,
	"SUCCEEDED":   3,
	"FAILED":      4,
	"ABORTED":     5,
}

func (x WorkflowInvocationStatus_Status) String() string {
	return proto.EnumName(WorkflowInvocationStatus_Status_name, int32(x))
}
func (WorkflowInvocationStatus_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{6, 0}
}

type TaskDependencyParameters_DependencyType int32

const (
	TaskDependencyParameters_DATA    TaskDependencyParameters_DependencyType = 0
	TaskDependencyParameters_CONTROL TaskDependencyParameters_DependencyType = 1
)

var TaskDependencyParameters_DependencyType_name = map[int32]string{
	0: "DATA",
	1: "CONTROL",
}
var TaskDependencyParameters_DependencyType_value = map[string]int32{
	"DATA":    0,
	"CONTROL": 1,
}

func (x TaskDependencyParameters_DependencyType) String() string {
	return proto.EnumName(TaskDependencyParameters_DependencyType_name, int32(x))
}
func (TaskDependencyParameters_DependencyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{10, 0}
}

type WorkflowStatus_Status int32

const (
	WorkflowStatus_UNKNOWN WorkflowStatus_Status = 0
	WorkflowStatus_PARSING WorkflowStatus_Status = 1
	WorkflowStatus_READY   WorkflowStatus_Status = 2
	WorkflowStatus_FAILED  WorkflowStatus_Status = 3
	WorkflowStatus_DELETED WorkflowStatus_Status = 4
)

var WorkflowStatus_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "PARSING",
	2: "READY",
	3: "FAILED",
	4: "DELETED",
}
var WorkflowStatus_Status_value = map[string]int32{
	"UNKNOWN": 0,
	"PARSING": 1,
	"READY":   2,
	"FAILED":  3,
	"DELETED": 4,
}

func (x WorkflowStatus_Status) String() string {
	return proto.EnumName(WorkflowStatus_Status_name, int32(x))
}
func (WorkflowStatus_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{11, 0} }

// Common
type ObjectMetadata struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	//    string namespace = 2;
	//    string name = 3;
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=createdAt" json:"createdAt,omitempty"`
}

func (m *ObjectMetadata) Reset()                    { *m = ObjectMetadata{} }
func (m *ObjectMetadata) String() string            { return proto.CompactTextString(m) }
func (*ObjectMetadata) ProtoMessage()               {}
func (*ObjectMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ObjectMetadata) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ObjectMetadata) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

// Workflow Model
type Workflow struct {
	Metadata *ObjectMetadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Spec     *WorkflowSpec   `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	Status   *WorkflowStatus `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *Workflow) Reset()                    { *m = Workflow{} }
func (m *Workflow) String() string            { return proto.CompactTextString(m) }
func (*Workflow) ProtoMessage()               {}
func (*Workflow) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Workflow) GetMetadata() *ObjectMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Workflow) GetSpec() *WorkflowSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Workflow) GetStatus() *WorkflowStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// Function Invocation Model
type FunctionInvocation struct {
	Metadata *ObjectMetadata           `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Spec     *FunctionInvocationSpec   `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	Status   *FunctionInvocationStatus `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *FunctionInvocation) Reset()                    { *m = FunctionInvocation{} }
func (m *FunctionInvocation) String() string            { return proto.CompactTextString(m) }
func (*FunctionInvocation) ProtoMessage()               {}
func (*FunctionInvocation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FunctionInvocation) GetMetadata() *ObjectMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *FunctionInvocation) GetSpec() *FunctionInvocationSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *FunctionInvocation) GetStatus() *FunctionInvocationStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type FunctionInvocationSpec struct {
	// Id of the function to be invoked (no ambiguatity at this point
	Type   *TaskTypeDef           `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	TaskId string                 `protobuf:"bytes,2,opt,name=taskId" json:"taskId,omitempty"`
	Inputs map[string]*TypedValue `protobuf:"bytes,3,rep,name=inputs" json:"inputs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *FunctionInvocationSpec) Reset()                    { *m = FunctionInvocationSpec{} }
func (m *FunctionInvocationSpec) String() string            { return proto.CompactTextString(m) }
func (*FunctionInvocationSpec) ProtoMessage()               {}
func (*FunctionInvocationSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *FunctionInvocationSpec) GetType() *TaskTypeDef {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *FunctionInvocationSpec) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *FunctionInvocationSpec) GetInputs() map[string]*TypedValue {
	if m != nil {
		return m.Inputs
	}
	return nil
}

type FunctionInvocationStatus struct {
	Status    FunctionInvocationStatus_Status `protobuf:"varint,1,opt,name=status,enum=FunctionInvocationStatus_Status" json:"status,omitempty"`
	UpdatedAt *google_protobuf.Timestamp      `protobuf:"bytes,2,opt,name=updatedAt" json:"updatedAt,omitempty"`
	Output    *TypedValue                     `protobuf:"bytes,3,opt,name=output" json:"output,omitempty"`
}

func (m *FunctionInvocationStatus) Reset()                    { *m = FunctionInvocationStatus{} }
func (m *FunctionInvocationStatus) String() string            { return proto.CompactTextString(m) }
func (*FunctionInvocationStatus) ProtoMessage()               {}
func (*FunctionInvocationStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *FunctionInvocationStatus) GetStatus() FunctionInvocationStatus_Status {
	if m != nil {
		return m.Status
	}
	return FunctionInvocationStatus_UNKNOWN
}

func (m *FunctionInvocationStatus) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *FunctionInvocationStatus) GetOutput() *TypedValue {
	if m != nil {
		return m.Output
	}
	return nil
}

// Workflow Invocation Model
type WorkflowInvocationSpec struct {
	WorkflowId string                 `protobuf:"bytes,1,opt,name=workflowId" json:"workflowId,omitempty"`
	Inputs     map[string]*TypedValue `protobuf:"bytes,3,rep,name=inputs" json:"inputs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *WorkflowInvocationSpec) Reset()                    { *m = WorkflowInvocationSpec{} }
func (m *WorkflowInvocationSpec) String() string            { return proto.CompactTextString(m) }
func (*WorkflowInvocationSpec) ProtoMessage()               {}
func (*WorkflowInvocationSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *WorkflowInvocationSpec) GetWorkflowId() string {
	if m != nil {
		return m.WorkflowId
	}
	return ""
}

func (m *WorkflowInvocationSpec) GetInputs() map[string]*TypedValue {
	if m != nil {
		return m.Inputs
	}
	return nil
}

type WorkflowInvocationStatus struct {
	Status    WorkflowInvocationStatus_Status `protobuf:"varint,1,opt,name=status,enum=WorkflowInvocationStatus_Status" json:"status,omitempty"`
	UpdatedAt *google_protobuf.Timestamp      `protobuf:"bytes,2,opt,name=updatedAt" json:"updatedAt,omitempty"`
	Tasks     map[string]*FunctionInvocation  `protobuf:"bytes,3,rep,name=tasks" json:"tasks,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Output    *TypedValue                     `protobuf:"bytes,4,opt,name=output" json:"output,omitempty"`
}

func (m *WorkflowInvocationStatus) Reset()                    { *m = WorkflowInvocationStatus{} }
func (m *WorkflowInvocationStatus) String() string            { return proto.CompactTextString(m) }
func (*WorkflowInvocationStatus) ProtoMessage()               {}
func (*WorkflowInvocationStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *WorkflowInvocationStatus) GetStatus() WorkflowInvocationStatus_Status {
	if m != nil {
		return m.Status
	}
	return WorkflowInvocationStatus_UNKNOWN
}

func (m *WorkflowInvocationStatus) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *WorkflowInvocationStatus) GetTasks() map[string]*FunctionInvocation {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *WorkflowInvocationStatus) GetOutput() *TypedValue {
	if m != nil {
		return m.Output
	}
	return nil
}

type WorkflowInvocation struct {
	Metadata *ObjectMetadata           `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Spec     *WorkflowInvocationSpec   `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	Status   *WorkflowInvocationStatus `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *WorkflowInvocation) Reset()                    { *m = WorkflowInvocation{} }
func (m *WorkflowInvocation) String() string            { return proto.CompactTextString(m) }
func (*WorkflowInvocation) ProtoMessage()               {}
func (*WorkflowInvocation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *WorkflowInvocation) GetMetadata() *ObjectMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *WorkflowInvocation) GetSpec() *WorkflowInvocationSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *WorkflowInvocation) GetStatus() *WorkflowInvocationStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// Workflow Definition
//
// The workflowDefinition contains the definition of a workflow.
//
// Ideally the source code (json, yaml) can be converted directly to this message.
// Naming, triggers and versioning of the workflow itself is out of the scope of this data structure, which is delegated
// to the user/system upon the creation of a workflow.
type WorkflowSpec struct {
	// apiVersion describes what version is of the workflow definition.
	//
	// By default the workflow engine will assume the latest version to be used.
	ApiVersion string `protobuf:"bytes,1,opt,name=apiVersion" json:"apiVersion,omitempty"`
	// TODO Parameters
	// Actions
	//
	// Dependency graph is build into the tasks
	Tasks map[string]*Task `protobuf:"bytes,2,rep,name=tasks" json:"tasks,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// From which task should the workflow return the output? Future: multiple? Implicit?
	OutputTask string `protobuf:"bytes,3,opt,name=outputTask" json:"outputTask,omitempty"`
}

func (m *WorkflowSpec) Reset()                    { *m = WorkflowSpec{} }
func (m *WorkflowSpec) String() string            { return proto.CompactTextString(m) }
func (*WorkflowSpec) ProtoMessage()               {}
func (*WorkflowSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *WorkflowSpec) GetApiVersion() string {
	if m != nil {
		return m.ApiVersion
	}
	return ""
}

func (m *WorkflowSpec) GetTasks() map[string]*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *WorkflowSpec) GetOutputTask() string {
	if m != nil {
		return m.OutputTask
	}
	return ""
}

// A task is the primitive unit of a workflow, representing an action that needs to be performed in order to continue.
//
// A task as a number of inputs and exactly two outputs
// Id is specified outside of task
type Task struct {
	// Same as the string in the map<string, Task> definition. One needs to be set.
	//
	// If there is a conflict the map key gets precendence over this field.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Type defines the type of the task that needs to be performed
	//
	// By default the following types are supported:
	// - `function`: execute a function.
	// - `exit`: Once this task is executed, the controller will exit the workflow, canceling all running tasks.
	Type string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	// Name/identifier of the function
	Name   string                 `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Inputs map[string]*TypedValue `protobuf:"bytes,4,rep,name=inputs" json:"inputs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Dependencies for this task to execute
	Dependencies map[string]*TaskDependencyParameters `protobuf:"bytes,5,rep,name=dependencies" json:"dependencies,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Number of dependencies to wait for
	DependenciesAwait int32 `protobuf:"varint,6,opt,name=dependencies_await,json=dependenciesAwait" json:"dependencies_await,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Task) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Task) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Task) GetInputs() map[string]*TypedValue {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *Task) GetDependencies() map[string]*TaskDependencyParameters {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func (m *Task) GetDependenciesAwait() int32 {
	if m != nil {
		return m.DependenciesAwait
	}
	return 0
}

type TaskDependencyParameters struct {
	Type  TaskDependencyParameters_DependencyType `protobuf:"varint,1,opt,name=type,enum=TaskDependencyParameters_DependencyType" json:"type,omitempty"`
	Alias string                                  `protobuf:"bytes,2,opt,name=alias" json:"alias,omitempty"`
}

func (m *TaskDependencyParameters) Reset()                    { *m = TaskDependencyParameters{} }
func (m *TaskDependencyParameters) String() string            { return proto.CompactTextString(m) }
func (*TaskDependencyParameters) ProtoMessage()               {}
func (*TaskDependencyParameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *TaskDependencyParameters) GetType() TaskDependencyParameters_DependencyType {
	if m != nil {
		return m.Type
	}
	return TaskDependencyParameters_DATA
}

func (m *TaskDependencyParameters) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

// Internal
type WorkflowStatus struct {
	Status        WorkflowStatus_Status      `protobuf:"varint,1,opt,name=status,enum=WorkflowStatus_Status" json:"status,omitempty"`
	UpdatedAt     *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=updatedAt" json:"updatedAt,omitempty"`
	ResolvedTasks map[string]*TaskTypeDef    `protobuf:"bytes,3,rep,name=resolvedTasks" json:"resolvedTasks,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *WorkflowStatus) Reset()                    { *m = WorkflowStatus{} }
func (m *WorkflowStatus) String() string            { return proto.CompactTextString(m) }
func (*WorkflowStatus) ProtoMessage()               {}
func (*WorkflowStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *WorkflowStatus) GetStatus() WorkflowStatus_Status {
	if m != nil {
		return m.Status
	}
	return WorkflowStatus_UNKNOWN
}

func (m *WorkflowStatus) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *WorkflowStatus) GetResolvedTasks() map[string]*TaskTypeDef {
	if m != nil {
		return m.ResolvedTasks
	}
	return nil
}

type TaskTypeDef struct {
	Src      string `protobuf:"bytes,1,opt,name=src" json:"src,omitempty"`
	Runtime  string `protobuf:"bytes,2,opt,name=runtime" json:"runtime,omitempty"`
	Resolved string `protobuf:"bytes,3,opt,name=resolved" json:"resolved,omitempty"`
}

func (m *TaskTypeDef) Reset()                    { *m = TaskTypeDef{} }
func (m *TaskTypeDef) String() string            { return proto.CompactTextString(m) }
func (*TaskTypeDef) ProtoMessage()               {}
func (*TaskTypeDef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *TaskTypeDef) GetSrc() string {
	if m != nil {
		return m.Src
	}
	return ""
}

func (m *TaskTypeDef) GetRuntime() string {
	if m != nil {
		return m.Runtime
	}
	return ""
}

func (m *TaskTypeDef) GetResolved() string {
	if m != nil {
		return m.Resolved
	}
	return ""
}

// Copy of protobuf's Any, to avoid protobuf requirement of a protobuf-based type.
type TypedValue struct {
	Type  string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *TypedValue) Reset()                    { *m = TypedValue{} }
func (m *TypedValue) String() string            { return proto.CompactTextString(m) }
func (*TypedValue) ProtoMessage()               {}
func (*TypedValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *TypedValue) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *TypedValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*ObjectMetadata)(nil), "ObjectMetadata")
	proto.RegisterType((*Workflow)(nil), "Workflow")
	proto.RegisterType((*FunctionInvocation)(nil), "FunctionInvocation")
	proto.RegisterType((*FunctionInvocationSpec)(nil), "FunctionInvocationSpec")
	proto.RegisterType((*FunctionInvocationStatus)(nil), "FunctionInvocationStatus")
	proto.RegisterType((*WorkflowInvocationSpec)(nil), "WorkflowInvocationSpec")
	proto.RegisterType((*WorkflowInvocationStatus)(nil), "WorkflowInvocationStatus")
	proto.RegisterType((*WorkflowInvocation)(nil), "WorkflowInvocation")
	proto.RegisterType((*WorkflowSpec)(nil), "WorkflowSpec")
	proto.RegisterType((*Task)(nil), "Task")
	proto.RegisterType((*TaskDependencyParameters)(nil), "TaskDependencyParameters")
	proto.RegisterType((*WorkflowStatus)(nil), "WorkflowStatus")
	proto.RegisterType((*TaskTypeDef)(nil), "TaskTypeDef")
	proto.RegisterType((*TypedValue)(nil), "TypedValue")
	proto.RegisterEnum("FunctionInvocationStatus_Status", FunctionInvocationStatus_Status_name, FunctionInvocationStatus_Status_value)
	proto.RegisterEnum("WorkflowInvocationStatus_Status", WorkflowInvocationStatus_Status_name, WorkflowInvocationStatus_Status_value)
	proto.RegisterEnum("TaskDependencyParameters_DependencyType", TaskDependencyParameters_DependencyType_name, TaskDependencyParameters_DependencyType_value)
	proto.RegisterEnum("WorkflowStatus_Status", WorkflowStatus_Status_name, WorkflowStatus_Status_value)
}

func init() { proto.RegisterFile("pkg/types/types.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 988 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x0e, 0xa9, 0x1f, 0x5b, 0x23, 0x5b, 0xa6, 0xb7, 0xa9, 0xc3, 0xa8, 0x40, 0xaa, 0x30, 0x05,
	0xa2, 0x20, 0x28, 0x8d, 0xaa, 0x40, 0x61, 0x24, 0x05, 0x0a, 0xc5, 0xa4, 0x13, 0x21, 0x8e, 0x24,
	0x50, 0x72, 0x82, 0x1a, 0x28, 0x02, 0x5a, 0x5c, 0x1b, 0x8c, 0x24, 0x92, 0x20, 0x97, 0x36, 0x74,
	0xee, 0x43, 0xf4, 0x54, 0x14, 0x7d, 0x90, 0xa2, 0x8f, 0xd1, 0x4b, 0x5f, 0xa4, 0xb7, 0x62, 0x97,
	0x4b, 0x69, 0x29, 0x91, 0x08, 0x8a, 0x3a, 0x17, 0x41, 0xbb, 0xf3, 0xed, 0xcc, 0xb7, 0xf3, 0xcd,
	0x0c, 0x17, 0x3e, 0x0f, 0xa6, 0x57, 0x87, 0x64, 0x11, 0xe0, 0x28, 0xf9, 0xd5, 0x83, 0xd0, 0x27,
	0x7e, 0xf3, 0xcb, 0x2b, 0xdf, 0xbf, 0x9a, 0xe1, 0x43, 0xb6, 0xba, 0x88, 0x2f, 0x0f, 0x89, 0x3b,
	0xc7, 0x11, 0xb1, 0xe7, 0x41, 0x02, 0xd0, 0xce, 0xa1, 0x31, 0xb8, 0xf8, 0x80, 0x27, 0xe4, 0x0d,
	0x26, 0xb6, 0x63, 0x13, 0x1b, 0x35, 0x40, 0x76, 0x1d, 0x55, 0x6a, 0x49, 0xed, 0x9a, 0x25, 0xbb,
	0x0e, 0x3a, 0x82, 0xda, 0x24, 0xc4, 0x36, 0xc1, 0x4e, 0x97, 0xa8, 0xa5, 0x96, 0xd4, 0xae, 0x77,
	0x9a, 0x7a, 0xe2, 0x56, 0x4f, 0xdd, 0xea, 0xe3, 0xd4, 0xad, 0xb5, 0x02, 0x6b, 0x3f, 0x4b, 0xb0,
	0xfd, 0xce, 0x0f, 0xa7, 0x97, 0x33, 0xff, 0x06, 0x3d, 0x85, 0xed, 0x39, 0x0f, 0xc1, 0x9c, 0xd7,
	0x3b, 0x7b, 0x7a, 0x36, 0xb2, 0xb5, 0x04, 0xa0, 0x87, 0x50, 0x8e, 0x02, 0x3c, 0x51, 0x65, 0x06,
	0xdc, 0xd5, 0x53, 0x2f, 0xa3, 0x00, 0x4f, 0x2c, 0x66, 0x42, 0x8f, 0xa1, 0x1a, 0x11, 0x9b, 0xc4,
	0x11, 0xe7, 0xb4, 0xb7, 0x02, 0xb1, 0x6d, 0x8b, 0x9b, 0xb5, 0xdf, 0x25, 0x40, 0x27, 0xb1, 0x37,
	0x21, 0xae, 0xef, 0xf5, 0xbc, 0x6b, 0x7f, 0x62, 0xd3, 0x7f, 0xff, 0x8d, 0xcf, 0xd3, 0x0c, 0x9f,
	0x7b, 0xfa, 0xa6, 0x3f, 0x81, 0xd9, 0x37, 0x6b, 0xcc, 0xee, 0xe7, 0xc1, 0xb3, 0x1c, 0xff, 0x92,
	0xe0, 0x20, 0xdf, 0x27, 0x6a, 0x41, 0x99, 0x0a, 0xca, 0x39, 0xee, 0xe8, 0x63, 0x3b, 0x9a, 0x8e,
	0x17, 0x01, 0x36, 0xf0, 0xa5, 0xc5, 0x2c, 0xe8, 0x00, 0xaa, 0xc4, 0x8e, 0xa6, 0x3d, 0x87, 0xd1,
	0xab, 0x59, 0x7c, 0x85, 0x9e, 0x43, 0xd5, 0xf5, 0x82, 0x98, 0x50, 0x1e, 0xa5, 0x76, 0xbd, 0xf3,
	0xa8, 0x80, 0xb6, 0xde, 0x63, 0x28, 0xd3, 0x23, 0xe1, 0xc2, 0xe2, 0x47, 0x9a, 0x27, 0x50, 0x17,
	0xb6, 0x91, 0x02, 0xa5, 0x29, 0x5e, 0xf0, 0xaa, 0xa0, 0x7f, 0xd1, 0x43, 0xa8, 0x5c, 0xdb, 0xb3,
	0x18, 0xf3, 0x9c, 0xd4, 0x75, 0x4a, 0xca, 0x79, 0x4b, 0xb7, 0xac, 0xc4, 0xf2, 0x4c, 0x3e, 0x92,
	0xb4, 0xdf, 0x64, 0x50, 0x8b, 0xae, 0x8f, 0x8e, 0x96, 0x99, 0xa2, 0x8e, 0x1b, 0x9d, 0x56, 0x61,
	0xa6, 0xf4, 0x6c, 0xc2, 0x68, 0x51, 0xc6, 0x81, 0xc3, 0x8b, 0x52, 0xfe, 0x78, 0x51, 0x2e, 0xc1,
	0xe8, 0x11, 0x54, 0xfd, 0x98, 0x04, 0x71, 0x5a, 0xcb, 0x19, 0xe2, 0xdc, 0xa4, 0x7d, 0x80, 0x2a,
	0xa7, 0x58, 0x87, 0xad, 0xb3, 0xfe, 0xeb, 0xfe, 0xe0, 0x5d, 0x5f, 0xb9, 0x83, 0x76, 0xa1, 0x36,
	0x3a, 0x7e, 0x65, 0x1a, 0x67, 0xa7, 0xa6, 0xa1, 0x48, 0x68, 0x0f, 0xea, 0xbd, 0xfe, 0xfb, 0xa1,
	0x35, 0x78, 0x69, 0x99, 0xa3, 0x91, 0x22, 0x33, 0xfb, 0xd9, 0xf1, 0xb1, 0x69, 0x1a, 0xa6, 0xa1,
	0x94, 0x10, 0x40, 0xf5, 0xa4, 0xdb, 0xa3, 0xd8, 0x32, 0xf5, 0xd3, 0x7d, 0x31, 0xb0, 0xc6, 0xa6,
	0xa1, 0x54, 0xe8, 0x62, 0xf4, 0xba, 0x37, 0x1c, 0x9a, 0x86, 0x52, 0xd5, 0xfe, 0x90, 0xe0, 0x20,
	0x2d, 0xdd, 0x35, 0xed, 0x1f, 0x00, 0xdc, 0xa4, 0x96, 0xb4, 0x25, 0x85, 0x9d, 0x1c, 0x85, 0xf3,
	0x1d, 0x7d, 0x52, 0x85, 0x7f, 0x2d, 0x81, 0x9a, 0x13, 0xb6, 0x48, 0xe1, 0x22, 0xe8, 0xed, 0x29,
	0xfc, 0x0c, 0x2a, 0xb4, 0x03, 0xd2, 0xa4, 0x7c, 0x55, 0x1c, 0x92, 0xf6, 0x12, 0xcf, 0x4a, 0x72,
	0x44, 0xa8, 0x8e, 0x72, 0x61, 0x75, 0x34, 0xdf, 0x00, 0xac, 0x4e, 0xe6, 0x24, 0xee, 0x49, 0x36,
	0x71, 0x9f, 0xe5, 0x54, 0xb5, 0x98, 0xc0, 0x9f, 0x3e, 0x69, 0xb1, 0xb1, 0xf9, 0xb7, 0x99, 0x81,
	0xff, 0x37, 0xff, 0xf2, 0xcb, 0xac, 0x70, 0xfe, 0x15, 0x09, 0xb0, 0x9c, 0x7f, 0x7f, 0x4a, 0xb0,
	0x23, 0xce, 0x78, 0x5a, 0xf9, 0x76, 0xe0, 0xbe, 0xc5, 0x61, 0xe4, 0xfa, 0x5e, 0x5a, 0xf9, 0xab,
	0x1d, 0xa4, 0xa7, 0x1a, 0xcb, 0x4c, 0x63, 0x35, 0xf3, 0x85, 0xc8, 0xd1, 0xf5, 0x01, 0x40, 0x22,
	0x1e, 0x35, 0x31, 0x5e, 0x35, 0x4b, 0xd8, 0x69, 0xfe, 0xf0, 0x11, 0x49, 0xbf, 0xc8, 0x4a, 0x5a,
	0x61, 0x21, 0x44, 0x11, 0xff, 0x91, 0xa1, 0x4c, 0xf7, 0x36, 0x3e, 0x9f, 0x88, 0xcf, 0xef, 0x64,
	0x36, 0x27, 0x13, 0x1b, 0x41, 0xd9, 0xb3, 0xe7, 0x98, 0xf3, 0x60, 0xff, 0xd1, 0x93, 0x65, 0x2f,
	0x97, 0xd9, 0x95, 0xf6, 0x59, 0x88, 0xbc, 0xce, 0x45, 0xcf, 0x61, 0xc7, 0xc1, 0x01, 0xf6, 0x1c,
	0xec, 0x4d, 0x5c, 0x1c, 0xa9, 0x15, 0x76, 0xe0, 0x5e, 0x72, 0xc0, 0x10, 0x2c, 0xc9, 0xb1, 0x0c,
	0x18, 0x7d, 0x0d, 0x48, 0x5c, 0xbf, 0xb7, 0x6f, 0x6c, 0x97, 0xa8, 0xd5, 0x96, 0xd4, 0xae, 0x58,
	0xfb, 0xa2, 0xa5, 0x4b, 0x0d, 0xb7, 0x35, 0x25, 0x9a, 0xe7, 0xb0, 0xbf, 0xc1, 0x2c, 0xc7, 0xdb,
	0x61, 0xd6, 0xdb, 0x7d, 0x76, 0xa7, 0xe5, 0xc1, 0xc5, 0xd0, 0x0e, 0xed, 0x39, 0x26, 0x38, 0x8c,
	0xc4, 0xdc, 0xff, 0x22, 0x81, 0x5a, 0x84, 0x43, 0xdf, 0x0b, 0xdf, 0xcf, 0x46, 0xa7, 0x5d, 0xe8,
	0x70, 0x95, 0xb8, 0x05, 0xbd, 0x01, 0x57, 0xea, 0x2e, 0x54, 0xec, 0x99, 0x6b, 0x47, 0x5c, 0xbe,
	0x64, 0xa1, 0x3d, 0x86, 0x46, 0x16, 0x8d, 0xb6, 0xa1, 0x6c, 0x74, 0xc7, 0x5d, 0xe5, 0x0e, 0xed,
	0xbd, 0xe3, 0x41, 0x7f, 0x6c, 0x0d, 0x4e, 0x15, 0x49, 0xfb, 0x5b, 0x86, 0x46, 0xf6, 0x59, 0x82,
	0xf4, 0xb5, 0x89, 0x78, 0xb0, 0xf6, 0x6e, 0xb9, 0xbd, 0x39, 0xf8, 0x0a, 0x76, 0x43, 0x1c, 0xf9,
	0xb3, 0x6b, 0xec, 0x8c, 0x85, 0x79, 0xa8, 0xad, 0x07, 0xb4, 0x44, 0x50, 0x52, 0x32, 0xd9, 0x83,
	0xcd, 0x3e, 0xa0, 0x4d, 0x50, 0x8e, 0x7a, 0x5a, 0x56, 0xbd, 0xec, 0x63, 0x45, 0x10, 0xec, 0x24,
	0x7f, 0xe2, 0xd5, 0x61, 0x6b, 0xd8, 0xb5, 0x46, 0xbd, 0xfe, 0x4b, 0x45, 0x42, 0x35, 0xa8, 0x58,
	0x66, 0xd7, 0xf8, 0x51, 0x91, 0x85, 0xd1, 0x56, 0xa2, 0x18, 0xc3, 0x3c, 0x35, 0xe9, 0x68, 0x2b,
	0x6b, 0x67, 0x50, 0x17, 0x22, 0x50, 0x42, 0x51, 0x38, 0x49, 0x09, 0x45, 0xe1, 0x04, 0xa9, 0xb0,
	0x15, 0xc6, 0x1e, 0x7d, 0xf3, 0x72, 0x01, 0xd3, 0x25, 0x6a, 0xc2, 0x76, 0x7a, 0x47, 0xde, 0x86,
	0xcb, 0xb5, 0xf6, 0x1d, 0xc0, 0xaa, 0x88, 0x97, 0x0d, 0x2c, 0x09, 0x0d, 0x7c, 0x57, 0xbc, 0xe8,
	0x0e, 0xbf, 0xda, 0x8b, 0xad, 0xf3, 0x0a, 0x7b, 0x7b, 0x5f, 0x54, 0x99, 0x30, 0xdf, 0xfe, 0x1b,
	0x00, 0x00, 0xff, 0xff, 0x62, 0x8f, 0xa2, 0x8f, 0x95, 0x0b, 0x00, 0x00,
}
