// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: courierNtfn.proto

// import "adminAlert.proto";

package notification

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

// Enum for courier notification types
type CourierNotificationType int32

const (
	CourierNotificationType_NEW_RECOMMENDATION CourierNotificationType = 0 // Yangi tavsiya bildirishnomasi
	CourierNotificationType_ORDER_ASSIGNMENT   CourierNotificationType = 1 // Buyurtma topshiriq bildirishnomasi
	CourierNotificationType_STATUS_UPDATE      CourierNotificationType = 2 // Holatni yangilash bildirishnomasi
)

// Enum value maps for CourierNotificationType.
var (
	CourierNotificationType_name = map[int32]string{
		0: "NEW_RECOMMENDATION",
		1: "ORDER_ASSIGNMENT",
		2: "STATUS_UPDATE",
	}
	CourierNotificationType_value = map[string]int32{
		"NEW_RECOMMENDATION": 0,
		"ORDER_ASSIGNMENT":   1,
		"STATUS_UPDATE":      2,
	}
)

func (x CourierNotificationType) Enum() *CourierNotificationType {
	p := new(CourierNotificationType)
	*p = x
	return p
}

func (x CourierNotificationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CourierNotificationType) Descriptor() protoreflect.EnumDescriptor {
	return file_courierNtfn_proto_enumTypes[0].Descriptor()
}

func (CourierNotificationType) Type() protoreflect.EnumType {
	return &file_courierNtfn_proto_enumTypes[0]
}

func (x CourierNotificationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CourierNotificationType.Descriptor instead.
func (CourierNotificationType) EnumDescriptor() ([]byte, []int) {
	return file_courierNtfn_proto_rawDescGZIP(), []int{0}
}

// CourierNotification message to represent a courier notification
type CourierNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NotificationId string                  `protobuf:"bytes,1,opt,name=notification_id,json=notificationId,proto3" json:"notification_id,omitempty"` // Bildirishnoma ID'si
	CourierId      string                  `protobuf:"bytes,2,opt,name=courier_id,json=courierId,proto3" json:"courier_id,omitempty"`                // Bildirishnomani oladigan kuryer ID'si
	OrderId        string                  `protobuf:"bytes,3,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`                      // Bildirishnomaga bog'langan buyurtma ID'si
	Type           CourierNotificationType `protobuf:"varint,4,opt,name=type,proto3,enum=ecommerce.CourierNotificationType" json:"type,omitempty"`   // Bildirishnoma turi
	Message        string                  `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`                                     // Bildirishnoma matni
	CreatedAt      string                  `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`                // Bildirishnoma yaratilgan vaqt
	IsRead         bool                    `protobuf:"varint,7,opt,name=is_read,json=isRead,proto3" json:"is_read,omitempty"`                        // Kuryer bildirishnomani o'qiganligini bildiradi
}

func (x *CourierNotification) Reset() {
	*x = CourierNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_courierNtfn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourierNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourierNotification) ProtoMessage() {}

func (x *CourierNotification) ProtoReflect() protoreflect.Message {
	mi := &file_courierNtfn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourierNotification.ProtoReflect.Descriptor instead.
func (*CourierNotification) Descriptor() ([]byte, []int) {
	return file_courierNtfn_proto_rawDescGZIP(), []int{0}
}

func (x *CourierNotification) GetNotificationId() string {
	if x != nil {
		return x.NotificationId
	}
	return ""
}

func (x *CourierNotification) GetCourierId() string {
	if x != nil {
		return x.CourierId
	}
	return ""
}

func (x *CourierNotification) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *CourierNotification) GetType() CourierNotificationType {
	if x != nil {
		return x.Type
	}
	return CourierNotificationType_NEW_RECOMMENDATION
}

func (x *CourierNotification) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CourierNotification) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *CourierNotification) GetIsRead() bool {
	if x != nil {
		return x.IsRead
	}
	return false
}

// Request message for creating a new courier notification
type CreateCourierNotificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourierId string                  `protobuf:"bytes,1,opt,name=courier_id,json=courierId,proto3" json:"courier_id,omitempty"`              // Bildirishnomani oladigan kuryer ID'si
	OrderId   string                  `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`                    // Bildirishnomaga bog'langan buyurtma ID'si
	Type      CourierNotificationType `protobuf:"varint,3,opt,name=type,proto3,enum=ecommerce.CourierNotificationType" json:"type,omitempty"` // Bildirishnoma turi
	Message   string                  `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`                                   // Bildirishnoma matni
}

func (x *CreateCourierNotificationRequest) Reset() {
	*x = CreateCourierNotificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_courierNtfn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCourierNotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCourierNotificationRequest) ProtoMessage() {}

func (x *CreateCourierNotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_courierNtfn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCourierNotificationRequest.ProtoReflect.Descriptor instead.
func (*CreateCourierNotificationRequest) Descriptor() ([]byte, []int) {
	return file_courierNtfn_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCourierNotificationRequest) GetCourierId() string {
	if x != nil {
		return x.CourierId
	}
	return ""
}

func (x *CreateCourierNotificationRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *CreateCourierNotificationRequest) GetType() CourierNotificationType {
	if x != nil {
		return x.Type
	}
	return CourierNotificationType_NEW_RECOMMENDATION
}

func (x *CreateCourierNotificationRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Response message for courier notification operations
type CourierNotificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success      bool                 `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`          // Amaliyot muvaffaqiyatli bo'lganligini bildiruvchi flag
	Message      string               `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`           // Javob xabari
	Notification *CourierNotification `protobuf:"bytes,3,opt,name=notification,proto3" json:"notification,omitempty"` // Kuryer bildirishnomasi ma'lumotlari (Create operatsiyasi uchun)
}

func (x *CourierNotificationResponse) Reset() {
	*x = CourierNotificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_courierNtfn_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourierNotificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourierNotificationResponse) ProtoMessage() {}

func (x *CourierNotificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_courierNtfn_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourierNotificationResponse.ProtoReflect.Descriptor instead.
func (*CourierNotificationResponse) Descriptor() ([]byte, []int) {
	return file_courierNtfn_proto_rawDescGZIP(), []int{2}
}

func (x *CourierNotificationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CourierNotificationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CourierNotificationResponse) GetNotification() *CourierNotification {
	if x != nil {
		return x.Notification
	}
	return nil
}

// Request message for reading or deleting a courier notification by ID
type CourierNotificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NotificationId string `protobuf:"bytes,1,opt,name=notification_id,json=notificationId,proto3" json:"notification_id,omitempty"` // Bildirishnoma ID'si
}

func (x *CourierNotificationRequest) Reset() {
	*x = CourierNotificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_courierNtfn_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourierNotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourierNotificationRequest) ProtoMessage() {}

func (x *CourierNotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_courierNtfn_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourierNotificationRequest.ProtoReflect.Descriptor instead.
func (*CourierNotificationRequest) Descriptor() ([]byte, []int) {
	return file_courierNtfn_proto_rawDescGZIP(), []int{3}
}

func (x *CourierNotificationRequest) GetNotificationId() string {
	if x != nil {
		return x.NotificationId
	}
	return ""
}

// Response message for listing all courier notifications
type CourierNotificationListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notifications []*CourierNotification `protobuf:"bytes,1,rep,name=notifications,proto3" json:"notifications,omitempty"` // Kuryer bildirishnomalari ro'yxati
}

func (x *CourierNotificationListResponse) Reset() {
	*x = CourierNotificationListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_courierNtfn_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourierNotificationListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourierNotificationListResponse) ProtoMessage() {}

func (x *CourierNotificationListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_courierNtfn_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourierNotificationListResponse.ProtoReflect.Descriptor instead.
func (*CourierNotificationListResponse) Descriptor() ([]byte, []int) {
	return file_courierNtfn_proto_rawDescGZIP(), []int{4}
}

func (x *CourierNotificationListResponse) GetNotifications() []*CourierNotification {
	if x != nil {
		return x.Notifications
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_courierNtfn_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_courierNtfn_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_courierNtfn_proto_rawDescGZIP(), []int{5}
}

var File_courierNtfn_proto protoreflect.FileDescriptor

var file_courierNtfn_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x4e, 0x74, 0x66, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x09, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x22, 0x82,
	0x02, 0x0a, 0x13, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x72, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73,
	0x5f, 0x72, 0x65, 0x61, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x52,
	0x65, 0x61, 0x64, 0x22, 0xae, 0x01, 0x0a, 0x20, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x75, 0x72, 0x69, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x75, 0x72,
	0x69, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f,
	0x75, 0x72, 0x69, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x36, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x22, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x75,
	0x72, 0x69, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x95, 0x01, 0x0a, 0x1b, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x42, 0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x69,
	0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x45, 0x0a, 0x1a,
	0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x22, 0x67, 0x0a, 0x1f, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65,
	0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x07, 0x0a, 0x05,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x2a, 0x5a, 0x0a, 0x17, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x16, 0x0a, 0x12, 0x4e, 0x45, 0x57, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e,
	0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x4f, 0x52, 0x44, 0x45,
	0x52, 0x5f, 0x41, 0x53, 0x53, 0x49, 0x47, 0x4e, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x11,
	0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10,
	0x02, 0x32, 0xbd, 0x03, 0x0a, 0x1a, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x70, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65,
	0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x2e,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x67, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x2e, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e,
	0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6a, 0x0a, 0x19, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x26, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72,
	0x69, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x10, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2a, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63,
	0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x18, 0x5a, 0x16, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_courierNtfn_proto_rawDescOnce sync.Once
	file_courierNtfn_proto_rawDescData = file_courierNtfn_proto_rawDesc
)

func file_courierNtfn_proto_rawDescGZIP() []byte {
	file_courierNtfn_proto_rawDescOnce.Do(func() {
		file_courierNtfn_proto_rawDescData = protoimpl.X.CompressGZIP(file_courierNtfn_proto_rawDescData)
	})
	return file_courierNtfn_proto_rawDescData
}

var file_courierNtfn_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_courierNtfn_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_courierNtfn_proto_goTypes = []any{
	(CourierNotificationType)(0),             // 0: ecommerce.CourierNotificationType
	(*CourierNotification)(nil),              // 1: ecommerce.CourierNotification
	(*CreateCourierNotificationRequest)(nil), // 2: ecommerce.CreateCourierNotificationRequest
	(*CourierNotificationResponse)(nil),      // 3: ecommerce.CourierNotificationResponse
	(*CourierNotificationRequest)(nil),       // 4: ecommerce.CourierNotificationRequest
	(*CourierNotificationListResponse)(nil),  // 5: ecommerce.CourierNotificationListResponse
	(*Empty)(nil),                            // 6: ecommerce.Empty
}
var file_courierNtfn_proto_depIdxs = []int32{
	0, // 0: ecommerce.CourierNotification.type:type_name -> ecommerce.CourierNotificationType
	0, // 1: ecommerce.CreateCourierNotificationRequest.type:type_name -> ecommerce.CourierNotificationType
	1, // 2: ecommerce.CourierNotificationResponse.notification:type_name -> ecommerce.CourierNotification
	1, // 3: ecommerce.CourierNotificationListResponse.notifications:type_name -> ecommerce.CourierNotification
	2, // 4: ecommerce.CourierNotificationService.CreateCourierNotification:input_type -> ecommerce.CreateCourierNotificationRequest
	4, // 5: ecommerce.CourierNotificationService.GetCourierNotification:input_type -> ecommerce.CourierNotificationRequest
	4, // 6: ecommerce.CourierNotificationService.DeleteCourierNotification:input_type -> ecommerce.CourierNotificationRequest
	6, // 7: ecommerce.CourierNotificationService.ListCourierNotifications:input_type -> ecommerce.Empty
	3, // 8: ecommerce.CourierNotificationService.CreateCourierNotification:output_type -> ecommerce.CourierNotificationResponse
	3, // 9: ecommerce.CourierNotificationService.GetCourierNotification:output_type -> ecommerce.CourierNotificationResponse
	3, // 10: ecommerce.CourierNotificationService.DeleteCourierNotification:output_type -> ecommerce.CourierNotificationResponse
	5, // 11: ecommerce.CourierNotificationService.ListCourierNotifications:output_type -> ecommerce.CourierNotificationListResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_courierNtfn_proto_init() }
func file_courierNtfn_proto_init() {
	if File_courierNtfn_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_courierNtfn_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CourierNotification); i {
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
		file_courierNtfn_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateCourierNotificationRequest); i {
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
		file_courierNtfn_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CourierNotificationResponse); i {
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
		file_courierNtfn_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CourierNotificationRequest); i {
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
		file_courierNtfn_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CourierNotificationListResponse); i {
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
		file_courierNtfn_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_courierNtfn_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_courierNtfn_proto_goTypes,
		DependencyIndexes: file_courierNtfn_proto_depIdxs,
		EnumInfos:         file_courierNtfn_proto_enumTypes,
		MessageInfos:      file_courierNtfn_proto_msgTypes,
	}.Build()
	File_courierNtfn_proto = out.File
	file_courierNtfn_proto_rawDesc = nil
	file_courierNtfn_proto_goTypes = nil
	file_courierNtfn_proto_depIdxs = nil
}
