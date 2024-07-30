// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: orders.proto

package product

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

// Enum for card status
type CardStatus int32

const (
	CardStatus_OR_PENDING CardStatus = 0
	CardStatus_COMPLETED  CardStatus = 1
	CardStatus_CANCELED   CardStatus = 2
)

// Enum value maps for CardStatus.
var (
	CardStatus_name = map[int32]string{
		0: "OR_PENDING",
		1: "COMPLETED",
		2: "CANCELED",
	}
	CardStatus_value = map[string]int32{
		"OR_PENDING": 0,
		"COMPLETED":  1,
		"CANCELED":   2,
	}
)

func (x CardStatus) Enum() *CardStatus {
	p := new(CardStatus)
	*p = x
	return p
}

func (x CardStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CardStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_orders_proto_enumTypes[0].Descriptor()
}

func (CardStatus) Type() protoreflect.EnumType {
	return &file_orders_proto_enumTypes[0]
}

func (x CardStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CardStatus.Descriptor instead.
func (CardStatus) EnumDescriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{0}
}

// Order message to represent an order
type OrderProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId    string     `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`            // Buyurtma noyob ID'si
	UserId     string     `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`               // Buyurtma bergan foydalanuvchi ID'si
	CreatedAt  string     `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`      // Buyurtma yaratilgan vaqti
	Status     CardStatus `protobuf:"varint,4,opt,name=status,proto3,enum=ecommerce.CardStatus" json:"status,omitempty"`  // Buyurtma holati
	TotalPrice float32    `protobuf:"fixed32,5,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"` // Buyurtma umumiy summasi
}

func (x *OrderProduct) Reset() {
	*x = OrderProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderProduct) ProtoMessage() {}

func (x *OrderProduct) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderProduct.ProtoReflect.Descriptor instead.
func (*OrderProduct) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{0}
}

func (x *OrderProduct) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *OrderProduct) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *OrderProduct) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *OrderProduct) GetStatus() CardStatus {
	if x != nil {
		return x.Status
	}
	return CardStatus_OR_PENDING
}

func (x *OrderProduct) GetTotalPrice() float32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

// Request message for creating a new order
type CreateOrderProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string     `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`               // Buyurtma bergan foydalanuvchi ID'si
	Status     CardStatus `protobuf:"varint,2,opt,name=status,proto3,enum=ecommerce.CardStatus" json:"status,omitempty"`  // Buyurtma holati
	TotalPrice float32    `protobuf:"fixed32,3,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"` // Buyurtma umumiy summasi
}

func (x *CreateOrderProductRequest) Reset() {
	*x = CreateOrderProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderProductRequest) ProtoMessage() {}

func (x *CreateOrderProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderProductRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderProductRequest) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrderProductRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateOrderProductRequest) GetStatus() CardStatus {
	if x != nil {
		return x.Status
	}
	return CardStatus_OR_PENDING
}

func (x *CreateOrderProductRequest) GetTotalPrice() float32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

// Request message for reading or deleting an order by ID
type OrderProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"` // Buyurtma noyob ID'si
}

func (x *OrderProductRequest) Reset() {
	*x = OrderProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderProductRequest) ProtoMessage() {}

func (x *OrderProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderProductRequest.ProtoReflect.Descriptor instead.
func (*OrderProductRequest) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{2}
}

func (x *OrderProductRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

// Request message for updating an order
type UpdateOrderProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId    string     `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`            // Buyurtma noyob ID'si
	Status     CardStatus `protobuf:"varint,2,opt,name=status,proto3,enum=ecommerce.CardStatus" json:"status,omitempty"`  // Buyurtma holati
	TotalPrice float32    `protobuf:"fixed32,3,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"` // Buyurtma umumiy summasi
}

func (x *UpdateOrderProductRequest) Reset() {
	*x = UpdateOrderProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOrderProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderProductRequest) ProtoMessage() {}

func (x *UpdateOrderProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderProductRequest.ProtoReflect.Descriptor instead.
func (*UpdateOrderProductRequest) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateOrderProductRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *UpdateOrderProductRequest) GetStatus() CardStatus {
	if x != nil {
		return x.Status
	}
	return CardStatus_OR_PENDING
}

func (x *UpdateOrderProductRequest) GetTotalPrice() float32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

// Response message for order operations
type OrderProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool          `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // Amaliyot muvaffaqiyatli bo'lganligini bildiruvchi flag
	Message string        `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`  // Javob xabari
	Order   *OrderProduct `protobuf:"bytes,3,opt,name=order,proto3" json:"order,omitempty"`      // Buyurtma ma'lumotlari (Create va Get operatsiyalari uchun)
}

func (x *OrderProductResponse) Reset() {
	*x = OrderProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderProductResponse) ProtoMessage() {}

func (x *OrderProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderProductResponse.ProtoReflect.Descriptor instead.
func (*OrderProductResponse) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{4}
}

func (x *OrderProductResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *OrderProductResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *OrderProductResponse) GetOrder() *OrderProduct {
	if x != nil {
		return x.Order
	}
	return nil
}

// Response message for listing all orders
type OrderProductListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*OrderProduct `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"` // Buyurtmalar ro'yxati
}

func (x *OrderProductListResponse) Reset() {
	*x = OrderProductListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderProductListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderProductListResponse) ProtoMessage() {}

func (x *OrderProductListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderProductListResponse.ProtoReflect.Descriptor instead.
func (*OrderProductListResponse) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{5}
}

func (x *OrderProductListResponse) GetOrders() []*OrderProduct {
	if x != nil {
		return x.Orders
	}
	return nil
}

// Request message for listing all orders (with optional filters)
type OrderProductListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string     `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`              // Foydalanuvchi ID'si bo'yicha buyurtmalarni filtrlash (ixtiyoriy)
	Status CardStatus `protobuf:"varint,2,opt,name=status,proto3,enum=ecommerce.CardStatus" json:"status,omitempty"` // Buyurtma holati bo'yicha buyurtmalarni filtrlash (ixtiyoriy)
}

func (x *OrderProductListRequest) Reset() {
	*x = OrderProductListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderProductListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderProductListRequest) ProtoMessage() {}

func (x *OrderProductListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderProductListRequest.ProtoReflect.Descriptor instead.
func (*OrderProductListRequest) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{6}
}

func (x *OrderProductListRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *OrderProductListRequest) GetStatus() CardStatus {
	if x != nil {
		return x.Status
	}
	return CardStatus_OR_PENDING
}

var File_orders_proto protoreflect.FileDescriptor

var file_orders_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x22, 0xb1, 0x01, 0x0a, 0x0c, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2d, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x84, 0x01,
	0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65,
	0x2e, 0x43, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x22, 0x30, 0x0a, 0x13, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x86, 0x01, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x15, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x61, 0x72, 0x64,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22,
	0x79, 0x0a, 0x14, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x4b, 0x0a, 0x18, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72,
	0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x61, 0x0a, 0x17, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x39, 0x0a, 0x0a, 0x43, 0x61,
	0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x52, 0x5f, 0x50,
	0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4d, 0x50,
	0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41, 0x4e, 0x43, 0x45,
	0x4c, 0x45, 0x44, 0x10, 0x02, 0x32, 0xb5, 0x03, 0x0a, 0x13, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a,
	0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x24, 0x2e, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x1e, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x54, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x24, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63,
	0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63,
	0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63,
	0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x12, 0x22, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x72, 0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x13, 0x5a,
	0x11, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orders_proto_rawDescOnce sync.Once
	file_orders_proto_rawDescData = file_orders_proto_rawDesc
)

func file_orders_proto_rawDescGZIP() []byte {
	file_orders_proto_rawDescOnce.Do(func() {
		file_orders_proto_rawDescData = protoimpl.X.CompressGZIP(file_orders_proto_rawDescData)
	})
	return file_orders_proto_rawDescData
}

var file_orders_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_orders_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_orders_proto_goTypes = []any{
	(CardStatus)(0),                   // 0: ecommerce.CardStatus
	(*OrderProduct)(nil),              // 1: ecommerce.OrderProduct
	(*CreateOrderProductRequest)(nil), // 2: ecommerce.CreateOrderProductRequest
	(*OrderProductRequest)(nil),       // 3: ecommerce.OrderProductRequest
	(*UpdateOrderProductRequest)(nil), // 4: ecommerce.UpdateOrderProductRequest
	(*OrderProductResponse)(nil),      // 5: ecommerce.OrderProductResponse
	(*OrderProductListResponse)(nil),  // 6: ecommerce.OrderProductListResponse
	(*OrderProductListRequest)(nil),   // 7: ecommerce.OrderProductListRequest
}
var file_orders_proto_depIdxs = []int32{
	0,  // 0: ecommerce.OrderProduct.status:type_name -> ecommerce.CardStatus
	0,  // 1: ecommerce.CreateOrderProductRequest.status:type_name -> ecommerce.CardStatus
	0,  // 2: ecommerce.UpdateOrderProductRequest.status:type_name -> ecommerce.CardStatus
	1,  // 3: ecommerce.OrderProductResponse.order:type_name -> ecommerce.OrderProduct
	1,  // 4: ecommerce.OrderProductListResponse.orders:type_name -> ecommerce.OrderProduct
	0,  // 5: ecommerce.OrderProductListRequest.status:type_name -> ecommerce.CardStatus
	2,  // 6: ecommerce.OrderProductService.CreateOrder:input_type -> ecommerce.CreateOrderProductRequest
	3,  // 7: ecommerce.OrderProductService.GetOrder:input_type -> ecommerce.OrderProductRequest
	4,  // 8: ecommerce.OrderProductService.UpdateOrder:input_type -> ecommerce.UpdateOrderProductRequest
	3,  // 9: ecommerce.OrderProductService.DeleteOrder:input_type -> ecommerce.OrderProductRequest
	7,  // 10: ecommerce.OrderProductService.ListOrders:input_type -> ecommerce.OrderProductListRequest
	5,  // 11: ecommerce.OrderProductService.CreateOrder:output_type -> ecommerce.OrderProductResponse
	5,  // 12: ecommerce.OrderProductService.GetOrder:output_type -> ecommerce.OrderProductResponse
	5,  // 13: ecommerce.OrderProductService.UpdateOrder:output_type -> ecommerce.OrderProductResponse
	5,  // 14: ecommerce.OrderProductService.DeleteOrder:output_type -> ecommerce.OrderProductResponse
	6,  // 15: ecommerce.OrderProductService.ListOrders:output_type -> ecommerce.OrderProductListResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_orders_proto_init() }
func file_orders_proto_init() {
	if File_orders_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orders_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*OrderProduct); i {
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
		file_orders_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateOrderProductRequest); i {
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
		file_orders_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*OrderProductRequest); i {
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
		file_orders_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateOrderProductRequest); i {
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
		file_orders_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*OrderProductResponse); i {
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
		file_orders_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*OrderProductListResponse); i {
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
		file_orders_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*OrderProductListRequest); i {
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
			RawDescriptor: file_orders_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orders_proto_goTypes,
		DependencyIndexes: file_orders_proto_depIdxs,
		EnumInfos:         file_orders_proto_enumTypes,
		MessageInfos:      file_orders_proto_msgTypes,
	}.Build()
	File_orders_proto = out.File
	file_orders_proto_rawDesc = nil
	file_orders_proto_goTypes = nil
	file_orders_proto_depIdxs = nil
}