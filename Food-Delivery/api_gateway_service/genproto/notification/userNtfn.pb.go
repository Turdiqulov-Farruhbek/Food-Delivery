// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: userNtfn.proto

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

// Enum for notification types
type NotificationType int32

const (
	NotificationType_ORDER_CONFIRMATION NotificationType = 0 // Buyurtmani tasdiqlash bildirishnomasi
	NotificationType_STATUS_UPDATE_USER NotificationType = 1 // Holatni yangilash bildirishnomasi
)

// Enum value maps for NotificationType.
var (
	NotificationType_name = map[int32]string{
		0: "ORDER_CONFIRMATION",
		1: "STATUS_UPDATE_USER",
	}
	NotificationType_value = map[string]int32{
		"ORDER_CONFIRMATION": 0,
		"STATUS_UPDATE_USER": 1,
	}
)

func (x NotificationType) Enum() *NotificationType {
	p := new(NotificationType)
	*p = x
	return p
}

func (x NotificationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotificationType) Descriptor() protoreflect.EnumDescriptor {
	return file_userNtfn_proto_enumTypes[0].Descriptor()
}

func (NotificationType) Type() protoreflect.EnumType {
	return &file_userNtfn_proto_enumTypes[0]
}

func (x NotificationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotificationType.Descriptor instead.
func (NotificationType) EnumDescriptor() ([]byte, []int) {
	return file_userNtfn_proto_rawDescGZIP(), []int{0}
}

// UserNotification message to represent a user notification
type UserNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NotificationId string           `protobuf:"bytes,1,opt,name=notification_id,json=notificationId,proto3" json:"notification_id,omitempty"` // Bildirishnoma ID'si
	UserId         string           `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                         // Bildirishnomani oladigan foydalanuvchi ID'si
	OrderId        string           `protobuf:"bytes,3,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`                      // Bildirishnomaga bog'langan buyurtma ID'si
	Type           NotificationType `protobuf:"varint,4,opt,name=type,proto3,enum=ecommerce.NotificationType" json:"type,omitempty"`          // Bildirishnoma turi
	Message        string           `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`                                     // Bildirishnoma matni
	CreatedAt      string           `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`                // Bildirishnoma yaratilgan vaqt
	IsRead         bool             `protobuf:"varint,7,opt,name=is_read,json=isRead,proto3" json:"is_read,omitempty"`                        // Foydalanuvchi bildirishnomani o'qiganligini bildiradi
}

func (x *UserNotification) Reset() {
	*x = UserNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userNtfn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserNotification) ProtoMessage() {}

func (x *UserNotification) ProtoReflect() protoreflect.Message {
	mi := &file_userNtfn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserNotification.ProtoReflect.Descriptor instead.
func (*UserNotification) Descriptor() ([]byte, []int) {
	return file_userNtfn_proto_rawDescGZIP(), []int{0}
}

func (x *UserNotification) GetNotificationId() string {
	if x != nil {
		return x.NotificationId
	}
	return ""
}

func (x *UserNotification) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserNotification) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *UserNotification) GetType() NotificationType {
	if x != nil {
		return x.Type
	}
	return NotificationType_ORDER_CONFIRMATION
}

func (x *UserNotification) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UserNotification) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *UserNotification) GetIsRead() bool {
	if x != nil {
		return x.IsRead
	}
	return false
}

// Request message for creating a new user notification
type CreateUserNotificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string           `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                // Bildirishnomani oladigan foydalanuvchi ID'si
	OrderId string           `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`             // Bildirishnomaga bog'langan buyurtma ID'si
	Type    NotificationType `protobuf:"varint,3,opt,name=type,proto3,enum=ecommerce.NotificationType" json:"type,omitempty"` // Bildirishnoma turi
	Message string           `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`                            // Bildirishnoma matni
}

func (x *CreateUserNotificationRequest) Reset() {
	*x = CreateUserNotificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userNtfn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserNotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserNotificationRequest) ProtoMessage() {}

func (x *CreateUserNotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userNtfn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserNotificationRequest.ProtoReflect.Descriptor instead.
func (*CreateUserNotificationRequest) Descriptor() ([]byte, []int) {
	return file_userNtfn_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserNotificationRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateUserNotificationRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *CreateUserNotificationRequest) GetType() NotificationType {
	if x != nil {
		return x.Type
	}
	return NotificationType_ORDER_CONFIRMATION
}

func (x *CreateUserNotificationRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Response message for user notification operations
type UserNotificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success      bool              `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`          // Amaliyot muvaffaqiyatli bo'lganligini bildiruvchi flag
	Message      string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`           // Javob xabari
	Notification *UserNotification `protobuf:"bytes,3,opt,name=notification,proto3" json:"notification,omitempty"` // Foydalanuvchi bildirishnomasi ma'lumotlari (Create operatsiyasi uchun)
}

func (x *UserNotificationResponse) Reset() {
	*x = UserNotificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userNtfn_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserNotificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserNotificationResponse) ProtoMessage() {}

func (x *UserNotificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userNtfn_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserNotificationResponse.ProtoReflect.Descriptor instead.
func (*UserNotificationResponse) Descriptor() ([]byte, []int) {
	return file_userNtfn_proto_rawDescGZIP(), []int{2}
}

func (x *UserNotificationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UserNotificationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UserNotificationResponse) GetNotification() *UserNotification {
	if x != nil {
		return x.Notification
	}
	return nil
}

// Request message for reading or deleting a user notification by ID
type UserNotificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NotificationId string `protobuf:"bytes,1,opt,name=notification_id,json=notificationId,proto3" json:"notification_id,omitempty"` // Bildirishnoma ID'si
}

func (x *UserNotificationRequest) Reset() {
	*x = UserNotificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userNtfn_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserNotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserNotificationRequest) ProtoMessage() {}

func (x *UserNotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userNtfn_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserNotificationRequest.ProtoReflect.Descriptor instead.
func (*UserNotificationRequest) Descriptor() ([]byte, []int) {
	return file_userNtfn_proto_rawDescGZIP(), []int{3}
}

func (x *UserNotificationRequest) GetNotificationId() string {
	if x != nil {
		return x.NotificationId
	}
	return ""
}

// Response message for listing all user notifications
type UserNotificationListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notifications []*UserNotification `protobuf:"bytes,1,rep,name=notifications,proto3" json:"notifications,omitempty"` // Foydalanuvchi bildirishnomalari ro'yxati
}

func (x *UserNotificationListResponse) Reset() {
	*x = UserNotificationListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userNtfn_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserNotificationListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserNotificationListResponse) ProtoMessage() {}

func (x *UserNotificationListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userNtfn_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserNotificationListResponse.ProtoReflect.Descriptor instead.
func (*UserNotificationListResponse) Descriptor() ([]byte, []int) {
	return file_userNtfn_proto_rawDescGZIP(), []int{4}
}

func (x *UserNotificationListResponse) GetNotifications() []*UserNotification {
	if x != nil {
		return x.Notifications
	}
	return nil
}

var File_userNtfn_proto protoreflect.FileDescriptor

var file_userNtfn_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x74, 0x66, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x1a, 0x10, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf2, 0x01,
	0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x2f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f,
	0x72, 0x65, 0x61, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x52, 0x65,
	0x61, 0x64, 0x22, 0x9e, 0x01, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72,
	0x63, 0x65, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x18, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x42, 0x0a, 0x17, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x27, 0x0a, 0x0f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x61, 0x0a, 0x1c, 0x55, 0x73, 0x65,
	0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0d, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2a, 0x42, 0x0a, 0x10,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x16, 0x0a, 0x12, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52,
	0x4d, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x01,
	0x32, 0x99, 0x03, 0x0a, 0x17, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x67, 0x0a, 0x16,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72,
	0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x2e, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x22, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x10, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x27, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x19, 0x5a, 0x17,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_userNtfn_proto_rawDescOnce sync.Once
	file_userNtfn_proto_rawDescData = file_userNtfn_proto_rawDesc
)

func file_userNtfn_proto_rawDescGZIP() []byte {
	file_userNtfn_proto_rawDescOnce.Do(func() {
		file_userNtfn_proto_rawDescData = protoimpl.X.CompressGZIP(file_userNtfn_proto_rawDescData)
	})
	return file_userNtfn_proto_rawDescData
}

var file_userNtfn_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_userNtfn_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_userNtfn_proto_goTypes = []any{
	(NotificationType)(0),                 // 0: ecommerce.NotificationType
	(*UserNotification)(nil),              // 1: ecommerce.UserNotification
	(*CreateUserNotificationRequest)(nil), // 2: ecommerce.CreateUserNotificationRequest
	(*UserNotificationResponse)(nil),      // 3: ecommerce.UserNotificationResponse
	(*UserNotificationRequest)(nil),       // 4: ecommerce.UserNotificationRequest
	(*UserNotificationListResponse)(nil),  // 5: ecommerce.UserNotificationListResponse
	(*Empty)(nil),                         // 6: ecommerce.Empty
}
var file_userNtfn_proto_depIdxs = []int32{
	0, // 0: ecommerce.UserNotification.type:type_name -> ecommerce.NotificationType
	0, // 1: ecommerce.CreateUserNotificationRequest.type:type_name -> ecommerce.NotificationType
	1, // 2: ecommerce.UserNotificationResponse.notification:type_name -> ecommerce.UserNotification
	1, // 3: ecommerce.UserNotificationListResponse.notifications:type_name -> ecommerce.UserNotification
	2, // 4: ecommerce.UserNotificationService.CreateUserNotification:input_type -> ecommerce.CreateUserNotificationRequest
	4, // 5: ecommerce.UserNotificationService.GetUserNotification:input_type -> ecommerce.UserNotificationRequest
	4, // 6: ecommerce.UserNotificationService.DeleteUserNotification:input_type -> ecommerce.UserNotificationRequest
	6, // 7: ecommerce.UserNotificationService.ListUserNotifications:input_type -> ecommerce.Empty
	3, // 8: ecommerce.UserNotificationService.CreateUserNotification:output_type -> ecommerce.UserNotificationResponse
	3, // 9: ecommerce.UserNotificationService.GetUserNotification:output_type -> ecommerce.UserNotificationResponse
	3, // 10: ecommerce.UserNotificationService.DeleteUserNotification:output_type -> ecommerce.UserNotificationResponse
	5, // 11: ecommerce.UserNotificationService.ListUserNotifications:output_type -> ecommerce.UserNotificationListResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_userNtfn_proto_init() }
func file_userNtfn_proto_init() {
	if File_userNtfn_proto != nil {
		return
	}
	file_adminAlert_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_userNtfn_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*UserNotification); i {
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
		file_userNtfn_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateUserNotificationRequest); i {
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
		file_userNtfn_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UserNotificationResponse); i {
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
		file_userNtfn_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UserNotificationRequest); i {
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
		file_userNtfn_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UserNotificationListResponse); i {
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
			RawDescriptor: file_userNtfn_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_userNtfn_proto_goTypes,
		DependencyIndexes: file_userNtfn_proto_depIdxs,
		EnumInfos:         file_userNtfn_proto_enumTypes,
		MessageInfos:      file_userNtfn_proto_msgTypes,
	}.Build()
	File_userNtfn_proto = out.File
	file_userNtfn_proto_rawDesc = nil
	file_userNtfn_proto_goTypes = nil
	file_userNtfn_proto_depIdxs = nil
}