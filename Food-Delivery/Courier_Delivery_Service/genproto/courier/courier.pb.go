// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: courier.proto

// import "courier_Order.proto";

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

// Enum for courier status
type CourierStatus int32

const (
	CourierStatus_AVAILABLE   CourierStatus = 0 // Kuryer holati: mavjud
	CourierStatus_UNAVAILABLE CourierStatus = 1 // Kuryer holati: mavjud emas
	CourierStatus_ON_DELIVERY CourierStatus = 2 // Kuryer holati: yetkazish jarayonida
)

// Enum value maps for CourierStatus.
var (
	CourierStatus_name = map[int32]string{
		0: "AVAILABLE",
		1: "UNAVAILABLE",
		2: "ON_DELIVERY",
	}
	CourierStatus_value = map[string]int32{
		"AVAILABLE":   0,
		"UNAVAILABLE": 1,
		"ON_DELIVERY": 2,
	}
)

func (x CourierStatus) Enum() *CourierStatus {
	p := new(CourierStatus)
	*p = x
	return p
}

func (x CourierStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CourierStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_courier_proto_enumTypes[0].Descriptor()
}

func (CourierStatus) Type() protoreflect.EnumType {
	return &file_courier_proto_enumTypes[0]
}

func (x CourierStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CourierStatus.Descriptor instead.
func (CourierStatus) EnumDescriptor() ([]byte, []int) {
	return file_courier_proto_rawDescGZIP(), []int{0}
}

// Courier message to represent a courier
type Courier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourierId   string        `protobuf:"bytes,1,opt,name=courier_id,json=courierId,proto3" json:"courier_id,omitempty"`        // Kuryerning noyob ID'si
	Name        string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                   // Kuryerning ismi
	PhoneNumber string        `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`  // Kuryerning telefon raqami
	Email       string        `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`                                 // Kuryerning elektron pochtasi
	Status      CourierStatus `protobuf:"varint,5,opt,name=status,proto3,enum=ecommerce.CourierStatus" json:"status,omitempty"` // Kuryerning holati
}

func (x *Courier) Reset() {
	*x = Courier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_courier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Courier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Courier) ProtoMessage() {}

func (x *Courier) ProtoReflect() protoreflect.Message {
	mi := &file_courier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Courier.ProtoReflect.Descriptor instead.
func (*Courier) Descriptor() ([]byte, []int) {
	return file_courier_proto_rawDescGZIP(), []int{0}
}

func (x *Courier) GetCourierId() string {
	if x != nil {
		return x.CourierId
	}
	return ""
}

func (x *Courier) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Courier) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *Courier) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Courier) GetStatus() CourierStatus {
	if x != nil {
		return x.Status
	}
	return CourierStatus_AVAILABLE
}

// Request message for creating a new courier
type CreateCourierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`                                   // Kuryerning ismi
	PhoneNumber string        `protobuf:"bytes,2,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`  // Kuryerning telefon raqami
	Email       string        `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`                                 // Kuryerning elektron pochtasi
	Status      CourierStatus `protobuf:"varint,4,opt,name=status,proto3,enum=ecommerce.CourierStatus" json:"status,omitempty"` // Kuryerning holati
}

func (x *CreateCourierRequest) Reset() {
	*x = CreateCourierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_courier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCourierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCourierRequest) ProtoMessage() {}

func (x *CreateCourierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_courier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCourierRequest.ProtoReflect.Descriptor instead.
func (*CreateCourierRequest) Descriptor() ([]byte, []int) {
	return file_courier_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCourierRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCourierRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *CreateCourierRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateCourierRequest) GetStatus() CourierStatus {
	if x != nil {
		return x.Status
	}
	return CourierStatus_AVAILABLE
}

// Response message for courier operations
type CourierResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // Amaliyot muvaffaqiyatli bo'lganligini bildiruvchi flag
	Message string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`  // Javob xabari
	Courier *Courier `protobuf:"bytes,3,opt,name=courier,proto3" json:"courier,omitempty"`  // Kuryer ma'lumotlari (Create va Get operatsiyalari uchun)
}

func (x *CourierResponse) Reset() {
	*x = CourierResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_courier_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourierResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourierResponse) ProtoMessage() {}

func (x *CourierResponse) ProtoReflect() protoreflect.Message {
	mi := &file_courier_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourierResponse.ProtoReflect.Descriptor instead.
func (*CourierResponse) Descriptor() ([]byte, []int) {
	return file_courier_proto_rawDescGZIP(), []int{2}
}

func (x *CourierResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CourierResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CourierResponse) GetCourier() *Courier {
	if x != nil {
		return x.Courier
	}
	return nil
}

// Request message for reading or deleting a courier by ID
type CourierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourierId string `protobuf:"bytes,1,opt,name=courier_id,json=courierId,proto3" json:"courier_id,omitempty"` // Kuryerning noyob ID'si
}

func (x *CourierRequest) Reset() {
	*x = CourierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_courier_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourierRequest) ProtoMessage() {}

func (x *CourierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_courier_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourierRequest.ProtoReflect.Descriptor instead.
func (*CourierRequest) Descriptor() ([]byte, []int) {
	return file_courier_proto_rawDescGZIP(), []int{3}
}

func (x *CourierRequest) GetCourierId() string {
	if x != nil {
		return x.CourierId
	}
	return ""
}

// Request message for updating a courier
type UpdateCourierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourierId   string        `protobuf:"bytes,1,opt,name=courier_id,json=courierId,proto3" json:"courier_id,omitempty"`        // Kuryerning noyob ID'si
	Name        string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                   // Kuryerning ismi
	PhoneNumber string        `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`  // Kuryerning telefon raqami
	Email       string        `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`                                 // Kuryerning elektron pochtasi
	Status      CourierStatus `protobuf:"varint,5,opt,name=status,proto3,enum=ecommerce.CourierStatus" json:"status,omitempty"` // Kuryerning holati
}

func (x *UpdateCourierRequest) Reset() {
	*x = UpdateCourierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_courier_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCourierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCourierRequest) ProtoMessage() {}

func (x *UpdateCourierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_courier_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCourierRequest.ProtoReflect.Descriptor instead.
func (*UpdateCourierRequest) Descriptor() ([]byte, []int) {
	return file_courier_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCourierRequest) GetCourierId() string {
	if x != nil {
		return x.CourierId
	}
	return ""
}

func (x *UpdateCourierRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateCourierRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *UpdateCourierRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateCourierRequest) GetStatus() CourierStatus {
	if x != nil {
		return x.Status
	}
	return CourierStatus_AVAILABLE
}

// Response message for listing all couriers
type CourierListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Couriers []*Courier `protobuf:"bytes,1,rep,name=couriers,proto3" json:"couriers,omitempty"` // Kuryerlar ro'yxati
}

func (x *CourierListResponse) Reset() {
	*x = CourierListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_courier_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourierListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourierListResponse) ProtoMessage() {}

func (x *CourierListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_courier_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourierListResponse.ProtoReflect.Descriptor instead.
func (*CourierListResponse) Descriptor() ([]byte, []int) {
	return file_courier_proto_rawDescGZIP(), []int{5}
}

func (x *CourierListResponse) GetCouriers() []*Courier {
	if x != nil {
		return x.Couriers
	}
	return nil
}

type EmptyCourier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyCourier) Reset() {
	*x = EmptyCourier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_courier_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyCourier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyCourier) ProtoMessage() {}

func (x *EmptyCourier) ProtoReflect() protoreflect.Message {
	mi := &file_courier_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyCourier.ProtoReflect.Descriptor instead.
func (*EmptyCourier) Descriptor() ([]byte, []int) {
	return file_courier_proto_rawDescGZIP(), []int{6}
}

var File_courier_proto protoreflect.FileDescriptor

var file_courier_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x22, 0xa7, 0x01, 0x0a, 0x07, 0x43,
	0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x72,
	0x69, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x18, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x43,
	0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x95, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x73, 0x0a, 0x0f,
	0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65,
	0x2e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65,
	0x72, 0x22, 0x2f, 0x0a, 0x0e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72,
	0x49, 0x64, 0x22, 0xb4, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75,
	0x72, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x72, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x45, 0x0a, 0x13, 0x43, 0x6f, 0x75,
	0x72, 0x69, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2e, 0x0a, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x43,
	0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x73,
	0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72,
	0x2a, 0x40, 0x0a, 0x0d, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x00,
	0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10,
	0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59,
	0x10, 0x02, 0x32, 0x82, 0x03, 0x0a, 0x0e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72,
	0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x72, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65,
	0x72, 0x12, 0x19, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x6f,
	0x75, 0x72, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72,
	0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x72, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x43,
	0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47,
	0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x73, 0x12, 0x17,
	0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x1a, 0x1e, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x72, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x13, 0x5a, 0x11, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_courier_proto_rawDescOnce sync.Once
	file_courier_proto_rawDescData = file_courier_proto_rawDesc
)

func file_courier_proto_rawDescGZIP() []byte {
	file_courier_proto_rawDescOnce.Do(func() {
		file_courier_proto_rawDescData = protoimpl.X.CompressGZIP(file_courier_proto_rawDescData)
	})
	return file_courier_proto_rawDescData
}

var file_courier_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_courier_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_courier_proto_goTypes = []any{
	(CourierStatus)(0),           // 0: ecommerce.CourierStatus
	(*Courier)(nil),              // 1: ecommerce.Courier
	(*CreateCourierRequest)(nil), // 2: ecommerce.CreateCourierRequest
	(*CourierResponse)(nil),      // 3: ecommerce.CourierResponse
	(*CourierRequest)(nil),       // 4: ecommerce.CourierRequest
	(*UpdateCourierRequest)(nil), // 5: ecommerce.UpdateCourierRequest
	(*CourierListResponse)(nil),  // 6: ecommerce.CourierListResponse
	(*EmptyCourier)(nil),         // 7: ecommerce.EmptyCourier
}
var file_courier_proto_depIdxs = []int32{
	0,  // 0: ecommerce.Courier.status:type_name -> ecommerce.CourierStatus
	0,  // 1: ecommerce.CreateCourierRequest.status:type_name -> ecommerce.CourierStatus
	1,  // 2: ecommerce.CourierResponse.courier:type_name -> ecommerce.Courier
	0,  // 3: ecommerce.UpdateCourierRequest.status:type_name -> ecommerce.CourierStatus
	1,  // 4: ecommerce.CourierListResponse.couriers:type_name -> ecommerce.Courier
	2,  // 5: ecommerce.CourierService.CreateCourier:input_type -> ecommerce.CreateCourierRequest
	4,  // 6: ecommerce.CourierService.GetCourier:input_type -> ecommerce.CourierRequest
	5,  // 7: ecommerce.CourierService.UpdateCourier:input_type -> ecommerce.UpdateCourierRequest
	4,  // 8: ecommerce.CourierService.DeleteCourier:input_type -> ecommerce.CourierRequest
	7,  // 9: ecommerce.CourierService.ListCouriers:input_type -> ecommerce.EmptyCourier
	3,  // 10: ecommerce.CourierService.CreateCourier:output_type -> ecommerce.CourierResponse
	3,  // 11: ecommerce.CourierService.GetCourier:output_type -> ecommerce.CourierResponse
	3,  // 12: ecommerce.CourierService.UpdateCourier:output_type -> ecommerce.CourierResponse
	3,  // 13: ecommerce.CourierService.DeleteCourier:output_type -> ecommerce.CourierResponse
	6,  // 14: ecommerce.CourierService.ListCouriers:output_type -> ecommerce.CourierListResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_courier_proto_init() }
func file_courier_proto_init() {
	if File_courier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_courier_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Courier); i {
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
		file_courier_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateCourierRequest); i {
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
		file_courier_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CourierResponse); i {
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
		file_courier_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CourierRequest); i {
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
		file_courier_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateCourierRequest); i {
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
		file_courier_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CourierListResponse); i {
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
		file_courier_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*EmptyCourier); i {
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
			RawDescriptor: file_courier_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_courier_proto_goTypes,
		DependencyIndexes: file_courier_proto_depIdxs,
		EnumInfos:         file_courier_proto_enumTypes,
		MessageInfos:      file_courier_proto_msgTypes,
	}.Build()
	File_courier_proto = out.File
	file_courier_proto_rawDesc = nil
	file_courier_proto_goTypes = nil
	file_courier_proto_depIdxs = nil
}
