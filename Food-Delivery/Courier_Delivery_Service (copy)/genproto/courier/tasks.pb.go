// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: tasks.proto

package courier

import (
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

// Task status enumeration
type TaskStatus int32

const (
	TaskStatus_TODO        TaskStatus = 0
	TaskStatus_IN_PROGRESS TaskStatus = 1
	TaskStatus_DONE        TaskStatus = 2
)

// Enum value maps for TaskStatus.
var (
	TaskStatus_name = map[int32]string{
		0: "TODO",
		1: "IN_PROGRESS",
		2: "DONE",
	}
	TaskStatus_value = map[string]int32{
		"TODO":        0,
		"IN_PROGRESS": 1,
		"DONE":        2,
	}
)

func (x TaskStatus) Enum() *TaskStatus {
	p := new(TaskStatus)
	*p = x
	return p
}

func (x TaskStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_tasks_proto_enumTypes[0].Descriptor()
}

func (TaskStatus) Type() protoreflect.EnumType {
	return &file_tasks_proto_enumTypes[0]
}

func (x TaskStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskStatus.Descriptor instead.
func (TaskStatus) EnumDescriptor() ([]byte, []int) {
	return file_tasks_proto_rawDescGZIP(), []int{0}
}

// Request message for creating a task
type CreateTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string     `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string     `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Status      TaskStatus `protobuf:"varint,3,opt,name=status,proto3,enum=taskmanagement.TaskStatus" json:"status,omitempty"`
	AssignedTo  int32      `protobuf:"varint,4,opt,name=assigned_to,json=assignedTo,proto3" json:"assigned_to,omitempty"`
	DueDate     string     `protobuf:"bytes,5,opt,name=due_date,json=dueDate,proto3" json:"due_date,omitempty"` // ISO 8601 date format (e.g., "2023-12-31")
}

func (x *CreateTaskRequest) Reset() {
	*x = CreateTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskRequest) ProtoMessage() {}

func (x *CreateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskRequest.ProtoReflect.Descriptor instead.
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return file_tasks_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTaskRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateTaskRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateTaskRequest) GetStatus() TaskStatus {
	if x != nil {
		return x.Status
	}
	return TaskStatus_TODO
}

func (x *CreateTaskRequest) GetAssignedTo() int32 {
	if x != nil {
		return x.AssignedTo
	}
	return 0
}

func (x *CreateTaskRequest) GetDueDate() string {
	if x != nil {
		return x.DueDate
	}
	return ""
}

// Response message for task operations
type TaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string     `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string     `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Status      TaskStatus `protobuf:"varint,4,opt,name=status,proto3,enum=taskmanagement.TaskStatus" json:"status,omitempty"`
	AssignedTo  int32      `protobuf:"varint,5,opt,name=assigned_to,json=assignedTo,proto3" json:"assigned_to,omitempty"`
	DueDate     string     `protobuf:"bytes,6,opt,name=due_date,json=dueDate,proto3" json:"due_date,omitempty"`       // ISO 8601 date format (e.g., "2023-12-31")
	CreatedAt   string     `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"` // ISO 8601 timestamp format (e.g., "2023-12-31T12:34:56Z")
	UpdatedAt   string     `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"` // ISO 8601 timestamp format (e.g., "2023-12-31T12:34:56Z")
}

func (x *TaskResponse) Reset() {
	*x = TaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasks_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResponse) ProtoMessage() {}

func (x *TaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskResponse.ProtoReflect.Descriptor instead.
func (*TaskResponse) Descriptor() ([]byte, []int) {
	return file_tasks_proto_rawDescGZIP(), []int{1}
}

func (x *TaskResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TaskResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TaskResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TaskResponse) GetStatus() TaskStatus {
	if x != nil {
		return x.Status
	}
	return TaskStatus_TODO
}

func (x *TaskResponse) GetAssignedTo() int32 {
	if x != nil {
		return x.AssignedTo
	}
	return 0
}

func (x *TaskResponse) GetDueDate() string {
	if x != nil {
		return x.DueDate
	}
	return ""
}

func (x *TaskResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *TaskResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

// Request message for getting a task by ID
type GetTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTaskRequest) Reset() {
	*x = GetTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasks_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskRequest) ProtoMessage() {}

func (x *GetTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskRequest.ProtoReflect.Descriptor instead.
func (*GetTaskRequest) Descriptor() ([]byte, []int) {
	return file_tasks_proto_rawDescGZIP(), []int{2}
}

func (x *GetTaskRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Request message for updating a task
type UpdateTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string     `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string     `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Status      TaskStatus `protobuf:"varint,4,opt,name=status,proto3,enum=taskmanagement.TaskStatus" json:"status,omitempty"`
	AssignedTo  int32      `protobuf:"varint,5,opt,name=assigned_to,json=assignedTo,proto3" json:"assigned_to,omitempty"`
	DueDate     string     `protobuf:"bytes,6,opt,name=due_date,json=dueDate,proto3" json:"due_date,omitempty"` // ISO 8601 date format (e.g., "2023-12-31")
}

func (x *UpdateTaskRequest) Reset() {
	*x = UpdateTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasks_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskRequest) ProtoMessage() {}

func (x *UpdateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskRequest.ProtoReflect.Descriptor instead.
func (*UpdateTaskRequest) Descriptor() ([]byte, []int) {
	return file_tasks_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateTaskRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateTaskRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateTaskRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateTaskRequest) GetStatus() TaskStatus {
	if x != nil {
		return x.Status
	}
	return TaskStatus_TODO
}

func (x *UpdateTaskRequest) GetAssignedTo() int32 {
	if x != nil {
		return x.AssignedTo
	}
	return 0
}

func (x *UpdateTaskRequest) GetDueDate() string {
	if x != nil {
		return x.DueDate
	}
	return ""
}

// Request message for deleting a task by ID
type DeleteTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTaskRequest) Reset() {
	*x = DeleteTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasks_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskRequest) ProtoMessage() {}

func (x *DeleteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskRequest.ProtoReflect.Descriptor instead.
func (*DeleteTaskRequest) Descriptor() ([]byte, []int) {
	return file_tasks_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteTaskRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Response message for task deletion
type DeleteTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteTaskResponse) Reset() {
	*x = DeleteTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasks_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskResponse) ProtoMessage() {}

func (x *DeleteTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskResponse.ProtoReflect.Descriptor instead.
func (*DeleteTaskResponse) Descriptor() ([]byte, []int) {
	return file_tasks_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTaskResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DeleteTaskResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Request message for getting all tasks assigned to a user
type GetAllTasksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssignedTo int32 `protobuf:"varint,1,opt,name=assigned_to,json=assignedTo,proto3" json:"assigned_to,omitempty"`
}

func (x *GetAllTasksRequest) Reset() {
	*x = GetAllTasksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasks_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTasksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTasksRequest) ProtoMessage() {}

func (x *GetAllTasksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTasksRequest.ProtoReflect.Descriptor instead.
func (*GetAllTasksRequest) Descriptor() ([]byte, []int) {
	return file_tasks_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllTasksRequest) GetAssignedTo() int32 {
	if x != nil {
		return x.AssignedTo
	}
	return 0
}

// Response message for getting all tasks
type GetAllTasksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*TaskResponse `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *GetAllTasksResponse) Reset() {
	*x = GetAllTasksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasks_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTasksResponse) ProtoMessage() {}

func (x *GetAllTasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tasks_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTasksResponse.ProtoReflect.Descriptor instead.
func (*GetAllTasksResponse) Descriptor() ([]byte, []int) {
	return file_tasks_proto_rawDescGZIP(), []int{7}
}

func (x *GetAllTasksResponse) GetTasks() []*TaskResponse {
	if x != nil {
		return x.Tasks
	}
	return nil
}

var File_tasks_proto protoreflect.FileDescriptor

var file_tasks_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74,
	0x61, 0x73, 0x6b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xbb, 0x01,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x74, 0x61,
	0x73, 0x6b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x6f,
	0x12, 0x19, 0x0a, 0x08, 0x64, 0x75, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x64, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x84, 0x02, 0x0a, 0x0c,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x61,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x75, 0x65,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x75, 0x65,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0xcb, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x5f, 0x74, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x61, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x75, 0x65, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x75, 0x65, 0x44, 0x61,
	0x74, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x48, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x35, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x61, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x6f, 0x22, 0x49, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x32, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x2a, 0x31, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x4f, 0x44, 0x4f, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x49,
	0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04,
	0x44, 0x4f, 0x4e, 0x45, 0x10, 0x02, 0x32, 0xa1, 0x03, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x12, 0x21, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x1e, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d,
	0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x21, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x21, 0x2e, 0x74, 0x61,
	0x73, 0x6b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x56, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x61, 0x73, 0x6b,
	0x73, 0x12, 0x22, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x61, 0x73,
	0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x13, 0x5a, 0x11, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tasks_proto_rawDescOnce sync.Once
	file_tasks_proto_rawDescData = file_tasks_proto_rawDesc
)

func file_tasks_proto_rawDescGZIP() []byte {
	file_tasks_proto_rawDescOnce.Do(func() {
		file_tasks_proto_rawDescData = protoimpl.X.CompressGZIP(file_tasks_proto_rawDescData)
	})
	return file_tasks_proto_rawDescData
}

var file_tasks_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tasks_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_tasks_proto_goTypes = []any{
	(TaskStatus)(0),             // 0: taskmanagement.TaskStatus
	(*CreateTaskRequest)(nil),   // 1: taskmanagement.CreateTaskRequest
	(*TaskResponse)(nil),        // 2: taskmanagement.TaskResponse
	(*GetTaskRequest)(nil),      // 3: taskmanagement.GetTaskRequest
	(*UpdateTaskRequest)(nil),   // 4: taskmanagement.UpdateTaskRequest
	(*DeleteTaskRequest)(nil),   // 5: taskmanagement.DeleteTaskRequest
	(*DeleteTaskResponse)(nil),  // 6: taskmanagement.DeleteTaskResponse
	(*GetAllTasksRequest)(nil),  // 7: taskmanagement.GetAllTasksRequest
	(*GetAllTasksResponse)(nil), // 8: taskmanagement.GetAllTasksResponse
}
var file_tasks_proto_depIdxs = []int32{
	0, // 0: taskmanagement.CreateTaskRequest.status:type_name -> taskmanagement.TaskStatus
	0, // 1: taskmanagement.TaskResponse.status:type_name -> taskmanagement.TaskStatus
	0, // 2: taskmanagement.UpdateTaskRequest.status:type_name -> taskmanagement.TaskStatus
	2, // 3: taskmanagement.GetAllTasksResponse.tasks:type_name -> taskmanagement.TaskResponse
	1, // 4: taskmanagement.TaskService.CreateTask:input_type -> taskmanagement.CreateTaskRequest
	3, // 5: taskmanagement.TaskService.GetTask:input_type -> taskmanagement.GetTaskRequest
	4, // 6: taskmanagement.TaskService.UpdateTask:input_type -> taskmanagement.UpdateTaskRequest
	5, // 7: taskmanagement.TaskService.DeleteTask:input_type -> taskmanagement.DeleteTaskRequest
	7, // 8: taskmanagement.TaskService.GetAllTasks:input_type -> taskmanagement.GetAllTasksRequest
	2, // 9: taskmanagement.TaskService.CreateTask:output_type -> taskmanagement.TaskResponse
	2, // 10: taskmanagement.TaskService.GetTask:output_type -> taskmanagement.TaskResponse
	2, // 11: taskmanagement.TaskService.UpdateTask:output_type -> taskmanagement.TaskResponse
	6, // 12: taskmanagement.TaskService.DeleteTask:output_type -> taskmanagement.DeleteTaskResponse
	8, // 13: taskmanagement.TaskService.GetAllTasks:output_type -> taskmanagement.GetAllTasksResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tasks_proto_init() }
func file_tasks_proto_init() {
	if File_tasks_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tasks_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateTaskRequest); i {
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
		file_tasks_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TaskResponse); i {
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
		file_tasks_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetTaskRequest); i {
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
		file_tasks_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateTaskRequest); i {
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
		file_tasks_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteTaskRequest); i {
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
		file_tasks_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteTaskResponse); i {
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
		file_tasks_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllTasksRequest); i {
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
		file_tasks_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllTasksResponse); i {
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
			RawDescriptor: file_tasks_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tasks_proto_goTypes,
		DependencyIndexes: file_tasks_proto_depIdxs,
		EnumInfos:         file_tasks_proto_enumTypes,
		MessageInfos:      file_tasks_proto_msgTypes,
	}.Build()
	File_tasks_proto = out.File
	file_tasks_proto_rawDesc = nil
	file_tasks_proto_goTypes = nil
	file_tasks_proto_depIdxs = nil
}
