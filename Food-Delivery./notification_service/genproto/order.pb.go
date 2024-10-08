// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.12.4
// source: order.proto

package genproto

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

type OrderCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeliveryAddress string `protobuf:"bytes,1,opt,name=DeliveryAddress,proto3" json:"DeliveryAddress,omitempty"`
	DeliveryTime    string `protobuf:"bytes,2,opt,name=DeliveryTime,proto3" json:"DeliveryTime,omitempty"`
}

func (x *OrderCreate) Reset() {
	*x = OrderCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCreate) ProtoMessage() {}

func (x *OrderCreate) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCreate.ProtoReflect.Descriptor instead.
func (*OrderCreate) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0}
}

func (x *OrderCreate) GetDeliveryAddress() string {
	if x != nil {
		return x.DeliveryAddress
	}
	return ""
}

func (x *OrderCreate) GetDeliveryTime() string {
	if x != nil {
		return x.DeliveryTime
	}
	return ""
}

type OrderCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string       `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Body   *OrderCreate `protobuf:"bytes,2,opt,name=Body,proto3" json:"Body,omitempty"`
}

func (x *OrderCreateReq) Reset() {
	*x = OrderCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCreateReq) ProtoMessage() {}

func (x *OrderCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCreateReq.ProtoReflect.Descriptor instead.
func (*OrderCreateReq) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{1}
}

func (x *OrderCreateReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *OrderCreateReq) GetBody() *OrderCreate {
	if x != nil {
		return x.Body
	}
	return nil
}

type OrderGet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string  `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	UserId       string  `protobuf:"bytes,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	CourierId    string  `protobuf:"bytes,3,opt,name=CourierId,proto3" json:"CourierId,omitempty"`
	Status       string  `protobuf:"bytes,4,opt,name=Status,proto3" json:"Status,omitempty"`
	Total        float32 `protobuf:"fixed32,5,opt,name=Total,proto3" json:"Total,omitempty"`
	Address      string  `protobuf:"bytes,6,opt,name=Address,proto3" json:"Address,omitempty"`
	CreatedAt    string  `protobuf:"bytes,7,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	DeliveryTime string  `protobuf:"bytes,8,opt,name=DeliveryTime,proto3" json:"DeliveryTime,omitempty"`
	IsPaid       string  `protobuf:"bytes,9,opt,name=IsPaid,proto3" json:"IsPaid,omitempty"`
}

func (x *OrderGet) Reset() {
	*x = OrderGet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderGet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderGet) ProtoMessage() {}

func (x *OrderGet) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderGet.ProtoReflect.Descriptor instead.
func (*OrderGet) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{2}
}

func (x *OrderGet) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrderGet) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *OrderGet) GetCourierId() string {
	if x != nil {
		return x.CourierId
	}
	return ""
}

func (x *OrderGet) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *OrderGet) GetTotal() float32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *OrderGet) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *OrderGet) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *OrderGet) GetDeliveryTime() string {
	if x != nil {
		return x.DeliveryTime
	}
	return ""
}

func (x *OrderGet) GetIsPaid() string {
	if x != nil {
		return x.IsPaid
	}
	return ""
}

type OrderUpt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourierId    string  `protobuf:"bytes,1,opt,name=CourierId,proto3" json:"CourierId,omitempty"`
	Status       string  `protobuf:"bytes,2,opt,name=Status,proto3" json:"Status,omitempty"`
	Total        float32 `protobuf:"fixed32,3,opt,name=Total,proto3" json:"Total,omitempty"`
	Address      string  `protobuf:"bytes,4,opt,name=Address,proto3" json:"Address,omitempty"`
	DeliveryTime string  `protobuf:"bytes,5,opt,name=DeliveryTime,proto3" json:"DeliveryTime,omitempty"`
}

func (x *OrderUpt) Reset() {
	*x = OrderUpt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderUpt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderUpt) ProtoMessage() {}

func (x *OrderUpt) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderUpt.ProtoReflect.Descriptor instead.
func (*OrderUpt) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{3}
}

func (x *OrderUpt) GetCourierId() string {
	if x != nil {
		return x.CourierId
	}
	return ""
}

func (x *OrderUpt) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *OrderUpt) GetTotal() float32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *OrderUpt) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *OrderUpt) GetDeliveryTime() string {
	if x != nil {
		return x.DeliveryTime
	}
	return ""
}

type OrderUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string    `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Body *OrderUpt `protobuf:"bytes,2,opt,name=Body,proto3" json:"Body,omitempty"`
}

func (x *OrderUpdate) Reset() {
	*x = OrderUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderUpdate) ProtoMessage() {}

func (x *OrderUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderUpdate.ProtoReflect.Descriptor instead.
func (*OrderUpdate) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{4}
}

func (x *OrderUpdate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrderUpdate) GetBody() *OrderUpt {
	if x != nil {
		return x.Body
	}
	return nil
}

type OrderFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string  `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Status    string  `protobuf:"bytes,2,opt,name=Status,proto3" json:"Status,omitempty"`
	Address   string  `protobuf:"bytes,3,opt,name=Address,proto3" json:"Address,omitempty"`
	CourierId string  `protobuf:"bytes,4,opt,name=CourierId,proto3" json:"CourierId,omitempty"`
	MinTotal  float32 `protobuf:"fixed32,5,opt,name=MinTotal,proto3" json:"MinTotal,omitempty"`
	MaxTotal  float32 `protobuf:"fixed32,6,opt,name=MaxTotal,proto3" json:"MaxTotal,omitempty"`
	TimeFrom  string  `protobuf:"bytes,7,opt,name=TimeFrom,proto3" json:"TimeFrom,omitempty"`
	TimeTo    string  `protobuf:"bytes,8,opt,name=TimeTo,proto3" json:"TimeTo,omitempty"`
	Filter    *Filter `protobuf:"bytes,9,opt,name=Filter,proto3" json:"Filter,omitempty"`
}

func (x *OrderFilter) Reset() {
	*x = OrderFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderFilter) ProtoMessage() {}

func (x *OrderFilter) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderFilter.ProtoReflect.Descriptor instead.
func (*OrderFilter) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{5}
}

func (x *OrderFilter) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *OrderFilter) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *OrderFilter) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *OrderFilter) GetCourierId() string {
	if x != nil {
		return x.CourierId
	}
	return ""
}

func (x *OrderFilter) GetMinTotal() float32 {
	if x != nil {
		return x.MinTotal
	}
	return 0
}

func (x *OrderFilter) GetMaxTotal() float32 {
	if x != nil {
		return x.MaxTotal
	}
	return 0
}

func (x *OrderFilter) GetTimeFrom() string {
	if x != nil {
		return x.TimeFrom
	}
	return ""
}

func (x *OrderFilter) GetTimeTo() string {
	if x != nil {
		return x.TimeTo
	}
	return ""
}

func (x *OrderFilter) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type OrderList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*OrderGet `protobuf:"bytes,1,rep,name=Orders,proto3" json:"Orders,omitempty"`
	Count  int32       `protobuf:"varint,2,opt,name=Count,proto3" json:"Count,omitempty"`
	Limit  int32       `protobuf:"varint,3,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Offset int32       `protobuf:"varint,4,opt,name=Offset,proto3" json:"Offset,omitempty"`
}

func (x *OrderList) Reset() {
	*x = OrderList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderList) ProtoMessage() {}

func (x *OrderList) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderList.ProtoReflect.Descriptor instead.
func (*OrderList) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{6}
}

func (x *OrderList) GetOrders() []*OrderGet {
	if x != nil {
		return x.Orders
	}
	return nil
}

func (x *OrderList) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *OrderList) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *OrderList) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type OrderAssign struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourierId string `protobuf:"bytes,1,opt,name=CourierId,proto3" json:"CourierId,omitempty"`
	OrderId   string `protobuf:"bytes,2,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
}

func (x *OrderAssign) Reset() {
	*x = OrderAssign{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderAssign) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderAssign) ProtoMessage() {}

func (x *OrderAssign) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderAssign.ProtoReflect.Descriptor instead.
func (*OrderAssign) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{7}
}

func (x *OrderAssign) GetCourierId() string {
	if x != nil {
		return x.CourierId
	}
	return ""
}

func (x *OrderAssign) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type OrderReview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string    `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	UserId       string    `protobuf:"bytes,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	CourierId    string    `protobuf:"bytes,3,opt,name=CourierId,proto3" json:"CourierId,omitempty"`
	Status       string    `protobuf:"bytes,4,opt,name=Status,proto3" json:"Status,omitempty"`
	Total        float32   `protobuf:"fixed32,5,opt,name=Total,proto3" json:"Total,omitempty"`
	Address      string    `protobuf:"bytes,6,opt,name=Address,proto3" json:"Address,omitempty"`
	CreatedAt    string    `protobuf:"bytes,7,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	DeliveryTime string    `protobuf:"bytes,8,opt,name=DeliveryTime,proto3" json:"DeliveryTime,omitempty"`
	IsPaid       string    `protobuf:"bytes,9,opt,name=IsPaid,proto3" json:"IsPaid,omitempty"`
	Items        *ItemList `protobuf:"bytes,10,opt,name=Items,proto3" json:"Items,omitempty"`
}

func (x *OrderReview) Reset() {
	*x = OrderReview{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderReview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderReview) ProtoMessage() {}

func (x *OrderReview) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderReview.ProtoReflect.Descriptor instead.
func (*OrderReview) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{8}
}

func (x *OrderReview) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrderReview) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *OrderReview) GetCourierId() string {
	if x != nil {
		return x.CourierId
	}
	return ""
}

func (x *OrderReview) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *OrderReview) GetTotal() float32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *OrderReview) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *OrderReview) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *OrderReview) GetDeliveryTime() string {
	if x != nil {
		return x.DeliveryTime
	}
	return ""
}

func (x *OrderReview) GetIsPaid() string {
	if x != nil {
		return x.IsPaid
	}
	return ""
}

func (x *OrderReview) GetItems() *ItemList {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_order_proto protoreflect.FileDescriptor

var file_order_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73,
	0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x28, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x54, 0x0a, 0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x04, 0x42, 0x6f, 0x64,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x04, 0x42, 0x6f, 0x64, 0x79, 0x22, 0xf2, 0x01, 0x0a, 0x08, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x47,
	0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f,
	0x75, 0x72, 0x69, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43,
	0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x22,
	0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x73, 0x50, 0x61, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x49, 0x73, 0x50, 0x61, 0x69, 0x64, 0x22, 0x94, 0x01, 0x0a, 0x08, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x55, 0x70, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x75, 0x72, 0x69,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x6f, 0x75, 0x72,
	0x69, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x22, 0x0a,
	0x0c, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x46, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64,
	0x12, 0x27, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x55, 0x70, 0x74, 0x52, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x8c, 0x02, 0x0a, 0x0b, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x69, 0x6e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x08, 0x4d, 0x69, 0x6e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x4d, 0x61, 0x78, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x08, 0x4d, 0x61, 0x78, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x69, 0x6d,
	0x65, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x69, 0x6d,
	0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x6f, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x6f, 0x12, 0x29, 0x0a,
	0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x52, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x7c, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x06, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x06, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x45, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0xa0, 0x02,
	0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x49, 0x73, 0x50, 0x61, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x49,
	0x73, 0x50, 0x61, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x32, 0xc1, 0x03, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x39, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x19, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x73, 0x75,
	0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x13, 0x2e, 0x73, 0x75, 0x62, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x47, 0x65, 0x74, 0x12, 0x36,
	0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x16, 0x2e,
	0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x0f, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x3a, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x14, 0x2e,
	0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0b, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x16, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x1a, 0x0f, 0x2e, 0x73, 0x75, 0x62,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x0b, 0x50,
	0x69, 0x63, 0x6b, 0x55, 0x70, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x73, 0x75, 0x62,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x73, 0x75,
	0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x36, 0x0a, 0x0b,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x73, 0x75,
	0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x73,
	0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_proto_rawDescOnce sync.Once
	file_order_proto_rawDescData = file_order_proto_rawDesc
)

func file_order_proto_rawDescGZIP() []byte {
	file_order_proto_rawDescOnce.Do(func() {
		file_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_proto_rawDescData)
	})
	return file_order_proto_rawDescData
}

var file_order_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_order_proto_goTypes = []interface{}{
	(*OrderCreate)(nil),    // 0: submodule.OrderCreate
	(*OrderCreateReq)(nil), // 1: submodule.OrderCreateReq
	(*OrderGet)(nil),       // 2: submodule.OrderGet
	(*OrderUpt)(nil),       // 3: submodule.OrderUpt
	(*OrderUpdate)(nil),    // 4: submodule.OrderUpdate
	(*OrderFilter)(nil),    // 5: submodule.OrderFilter
	(*OrderList)(nil),      // 6: submodule.OrderList
	(*OrderAssign)(nil),    // 7: submodule.OrderAssign
	(*OrderReview)(nil),    // 8: submodule.OrderReview
	(*Filter)(nil),         // 9: submodule.Filter
	(*ItemList)(nil),       // 10: submodule.ItemList
	(*ById)(nil),           // 11: submodule.ById
	(*Void)(nil),           // 12: submodule.Void
}
var file_order_proto_depIdxs = []int32{
	0,  // 0: submodule.OrderCreateReq.Body:type_name -> submodule.OrderCreate
	3,  // 1: submodule.OrderUpdate.Body:type_name -> submodule.OrderUpt
	9,  // 2: submodule.OrderFilter.Filter:type_name -> submodule.Filter
	2,  // 3: submodule.OrderList.Orders:type_name -> submodule.OrderGet
	10, // 4: submodule.OrderReview.Items:type_name -> submodule.ItemList
	1,  // 5: submodule.OrderService.CreateOrder:input_type -> submodule.OrderCreateReq
	11, // 6: submodule.OrderService.GetOrder:input_type -> submodule.ById
	4,  // 7: submodule.OrderService.UpdateOrder:input_type -> submodule.OrderUpdate
	11, // 8: submodule.OrderService.DeleteOrder:input_type -> submodule.ById
	5,  // 9: submodule.OrderService.ListOrders:input_type -> submodule.OrderFilter
	7,  // 10: submodule.OrderService.AssignOrder:input_type -> submodule.OrderAssign
	11, // 11: submodule.OrderService.PickUpOrder:input_type -> submodule.ById
	11, // 12: submodule.OrderService.ReviewOrder:input_type -> submodule.ById
	12, // 13: submodule.OrderService.CreateOrder:output_type -> submodule.Void
	2,  // 14: submodule.OrderService.GetOrder:output_type -> submodule.OrderGet
	12, // 15: submodule.OrderService.UpdateOrder:output_type -> submodule.Void
	12, // 16: submodule.OrderService.DeleteOrder:output_type -> submodule.Void
	6,  // 17: submodule.OrderService.ListOrders:output_type -> submodule.OrderList
	12, // 18: submodule.OrderService.AssignOrder:output_type -> submodule.Void
	12, // 19: submodule.OrderService.PickUpOrder:output_type -> submodule.Void
	8,  // 20: submodule.OrderService.ReviewOrder:output_type -> submodule.OrderReview
	13, // [13:21] is the sub-list for method output_type
	5,  // [5:13] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_order_proto_init() }
func file_order_proto_init() {
	if File_order_proto != nil {
		return
	}
	file_common_proto_init()
	file_item_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCreate); i {
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
		file_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCreateReq); i {
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
		file_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderGet); i {
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
		file_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderUpt); i {
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
		file_order_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderUpdate); i {
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
		file_order_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderFilter); i {
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
		file_order_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderList); i {
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
		file_order_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderAssign); i {
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
		file_order_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderReview); i {
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
			RawDescriptor: file_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_proto_goTypes,
		DependencyIndexes: file_order_proto_depIdxs,
		MessageInfos:      file_order_proto_msgTypes,
	}.Build()
	File_order_proto = out.File
	file_order_proto_rawDesc = nil
	file_order_proto_goTypes = nil
	file_order_proto_depIdxs = nil
}
