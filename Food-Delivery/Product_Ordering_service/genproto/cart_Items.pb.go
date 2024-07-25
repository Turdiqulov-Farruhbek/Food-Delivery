// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: cart_Items.proto

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

// CartItem message to represent an item in a cart
type CartItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartItemId string `protobuf:"bytes,1,opt,name=cart_item_id,json=cartItemId,proto3" json:"cart_item_id,omitempty"` // Savat mahsuloti noyob ID'si
	CartId     string `protobuf:"bytes,2,opt,name=cart_id,json=cartId,proto3" json:"cart_id,omitempty"`               // Savat ID'si
	ProductId  string `protobuf:"bytes,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`      // Mahsulot ID'si
	Quantity   int32  `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`                        // Mahsulot miqdori
	Options    string `protobuf:"bytes,5,opt,name=options,proto3" json:"options,omitempty"`                           // Variantlar, masalan, o'lcham, qo'shimchalar va hokazo
}

func (x *CartItem) Reset() {
	*x = CartItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_Items_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItem) ProtoMessage() {}

func (x *CartItem) ProtoReflect() protoreflect.Message {
	mi := &file_cart_Items_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItem.ProtoReflect.Descriptor instead.
func (*CartItem) Descriptor() ([]byte, []int) {
	return file_cart_Items_proto_rawDescGZIP(), []int{0}
}

func (x *CartItem) GetCartItemId() string {
	if x != nil {
		return x.CartItemId
	}
	return ""
}

func (x *CartItem) GetCartId() string {
	if x != nil {
		return x.CartId
	}
	return ""
}

func (x *CartItem) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *CartItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *CartItem) GetOptions() string {
	if x != nil {
		return x.Options
	}
	return ""
}

// Request message for creating a new cart item
type CreateCartItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartId    string `protobuf:"bytes,1,opt,name=cart_id,json=cartId,proto3" json:"cart_id,omitempty"`          // Savat ID'si
	ProductId string `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"` // Mahsulot ID'si
	Quantity  int32  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`                   // Mahsulot miqdori
	Options   string `protobuf:"bytes,4,opt,name=options,proto3" json:"options,omitempty"`                      // Variantlar, masalan, o'lcham, qo'shimchalar va hokazo
}

func (x *CreateCartItemRequest) Reset() {
	*x = CreateCartItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_Items_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCartItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCartItemRequest) ProtoMessage() {}

func (x *CreateCartItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cart_Items_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCartItemRequest.ProtoReflect.Descriptor instead.
func (*CreateCartItemRequest) Descriptor() ([]byte, []int) {
	return file_cart_Items_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCartItemRequest) GetCartId() string {
	if x != nil {
		return x.CartId
	}
	return ""
}

func (x *CreateCartItemRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *CreateCartItemRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *CreateCartItemRequest) GetOptions() string {
	if x != nil {
		return x.Options
	}
	return ""
}

// Request message for reading or deleting a cart item by ID
type CartItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartItemId string `protobuf:"bytes,1,opt,name=cart_item_id,json=cartItemId,proto3" json:"cart_item_id,omitempty"` // Savat mahsuloti noyob ID'si
}

func (x *CartItemRequest) Reset() {
	*x = CartItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_Items_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItemRequest) ProtoMessage() {}

func (x *CartItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cart_Items_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItemRequest.ProtoReflect.Descriptor instead.
func (*CartItemRequest) Descriptor() ([]byte, []int) {
	return file_cart_Items_proto_rawDescGZIP(), []int{2}
}

func (x *CartItemRequest) GetCartItemId() string {
	if x != nil {
		return x.CartItemId
	}
	return ""
}

// Request message for updating a cart item
type UpdateCartItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartItemId string `protobuf:"bytes,1,opt,name=cart_item_id,json=cartItemId,proto3" json:"cart_item_id,omitempty"` // Savat mahsuloti noyob ID'si
	CartId     string `protobuf:"bytes,2,opt,name=cart_id,json=cartId,proto3" json:"cart_id,omitempty"`               // Savat ID'si
	ProductId  string `protobuf:"bytes,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`      // Mahsulot ID'si
	Quantity   int32  `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`                        // Mahsulot miqdori
	Options    string `protobuf:"bytes,5,opt,name=options,proto3" json:"options,omitempty"`                           // Variantlar, masalan, o'lcham, qo'shimchalar va hokazo
}

func (x *UpdateCartItemRequest) Reset() {
	*x = UpdateCartItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_Items_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCartItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCartItemRequest) ProtoMessage() {}

func (x *UpdateCartItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cart_Items_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCartItemRequest.ProtoReflect.Descriptor instead.
func (*UpdateCartItemRequest) Descriptor() ([]byte, []int) {
	return file_cart_Items_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateCartItemRequest) GetCartItemId() string {
	if x != nil {
		return x.CartItemId
	}
	return ""
}

func (x *UpdateCartItemRequest) GetCartId() string {
	if x != nil {
		return x.CartId
	}
	return ""
}

func (x *UpdateCartItemRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *UpdateCartItemRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *UpdateCartItemRequest) GetOptions() string {
	if x != nil {
		return x.Options
	}
	return ""
}

// Response message for cart item operations
type CartItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success  bool      `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`                  // Amaliyot muvaffaqiyatli bo'lganligini bildiruvchi flag
	Message  string    `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`                   // Javob xabari
	CartItem *CartItem `protobuf:"bytes,3,opt,name=cart_item,json=cartItem,proto3" json:"cart_item,omitempty"` // Savat mahsuloti ma'lumotlari (Create va Get operatsiyalari uchun)
}

func (x *CartItemResponse) Reset() {
	*x = CartItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_Items_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItemResponse) ProtoMessage() {}

func (x *CartItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cart_Items_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItemResponse.ProtoReflect.Descriptor instead.
func (*CartItemResponse) Descriptor() ([]byte, []int) {
	return file_cart_Items_proto_rawDescGZIP(), []int{4}
}

func (x *CartItemResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CartItemResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CartItemResponse) GetCartItem() *CartItem {
	if x != nil {
		return x.CartItem
	}
	return nil
}

// Response message for listing all cart items
type CartItemListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartItems []*CartItem `protobuf:"bytes,1,rep,name=cart_items,json=cartItems,proto3" json:"cart_items,omitempty"` // Savat mahsulotlari ro'yxati
}

func (x *CartItemListResponse) Reset() {
	*x = CartItemListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_Items_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItemListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItemListResponse) ProtoMessage() {}

func (x *CartItemListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cart_Items_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItemListResponse.ProtoReflect.Descriptor instead.
func (*CartItemListResponse) Descriptor() ([]byte, []int) {
	return file_cart_Items_proto_rawDescGZIP(), []int{5}
}

func (x *CartItemListResponse) GetCartItems() []*CartItem {
	if x != nil {
		return x.CartItems
	}
	return nil
}

var File_cart_Items_proto protoreflect.FileDescriptor

var file_cart_Items_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x1a, 0x0a, 0x63,
	0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x01, 0x0a, 0x08, 0x43, 0x61,
	0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x20, 0x0a, 0x0c, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x69,
	0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61,
	0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x72, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x72, 0x74, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x85, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x61, 0x72, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x33,
	0x0a, 0x0f, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x0c, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x49, 0x64, 0x22, 0xa7, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61,
	0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a,
	0x0c, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x61, 0x72, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x78, 0x0a,
	0x10, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x74,
	0x65, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x63,
	0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x4a, 0x0a, 0x14, 0x43, 0x61, 0x72, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x32, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e,
	0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x09, 0x63, 0x61, 0x72, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x32, 0x8a, 0x03, 0x0a, 0x0f, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x20, 0x2e, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1a, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x72, 0x63, 0x65, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e,
	0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4f, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x20, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65,
	0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x49, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x1a, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e,
	0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x61, 0x72, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0d,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x10, 0x2e,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x1f, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x61, 0x72, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cart_Items_proto_rawDescOnce sync.Once
	file_cart_Items_proto_rawDescData = file_cart_Items_proto_rawDesc
)

func file_cart_Items_proto_rawDescGZIP() []byte {
	file_cart_Items_proto_rawDescOnce.Do(func() {
		file_cart_Items_proto_rawDescData = protoimpl.X.CompressGZIP(file_cart_Items_proto_rawDescData)
	})
	return file_cart_Items_proto_rawDescData
}

var file_cart_Items_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cart_Items_proto_goTypes = []any{
	(*CartItem)(nil),              // 0: ecommerce.CartItem
	(*CreateCartItemRequest)(nil), // 1: ecommerce.CreateCartItemRequest
	(*CartItemRequest)(nil),       // 2: ecommerce.CartItemRequest
	(*UpdateCartItemRequest)(nil), // 3: ecommerce.UpdateCartItemRequest
	(*CartItemResponse)(nil),      // 4: ecommerce.CartItemResponse
	(*CartItemListResponse)(nil),  // 5: ecommerce.CartItemListResponse
	(*Empty)(nil),                 // 6: ecommerce.Empty
}
var file_cart_Items_proto_depIdxs = []int32{
	0, // 0: ecommerce.CartItemResponse.cart_item:type_name -> ecommerce.CartItem
	0, // 1: ecommerce.CartItemListResponse.cart_items:type_name -> ecommerce.CartItem
	1, // 2: ecommerce.CartItemService.CreateCartItem:input_type -> ecommerce.CreateCartItemRequest
	2, // 3: ecommerce.CartItemService.GetCartItem:input_type -> ecommerce.CartItemRequest
	3, // 4: ecommerce.CartItemService.UpdateCartItem:input_type -> ecommerce.UpdateCartItemRequest
	2, // 5: ecommerce.CartItemService.DeleteCartItem:input_type -> ecommerce.CartItemRequest
	6, // 6: ecommerce.CartItemService.ListCartItems:input_type -> ecommerce.Empty
	4, // 7: ecommerce.CartItemService.CreateCartItem:output_type -> ecommerce.CartItemResponse
	4, // 8: ecommerce.CartItemService.GetCartItem:output_type -> ecommerce.CartItemResponse
	4, // 9: ecommerce.CartItemService.UpdateCartItem:output_type -> ecommerce.CartItemResponse
	4, // 10: ecommerce.CartItemService.DeleteCartItem:output_type -> ecommerce.CartItemResponse
	5, // 11: ecommerce.CartItemService.ListCartItems:output_type -> ecommerce.CartItemListResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cart_Items_proto_init() }
func file_cart_Items_proto_init() {
	if File_cart_Items_proto != nil {
		return
	}
	file_cart_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cart_Items_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CartItem); i {
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
		file_cart_Items_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateCartItemRequest); i {
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
		file_cart_Items_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CartItemRequest); i {
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
		file_cart_Items_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateCartItemRequest); i {
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
		file_cart_Items_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CartItemResponse); i {
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
		file_cart_Items_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CartItemListResponse); i {
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
			RawDescriptor: file_cart_Items_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cart_Items_proto_goTypes,
		DependencyIndexes: file_cart_Items_proto_depIdxs,
		MessageInfos:      file_cart_Items_proto_msgTypes,
	}.Build()
	File_cart_Items_proto = out.File
	file_cart_Items_proto_rawDesc = nil
	file_cart_Items_proto_goTypes = nil
	file_cart_Items_proto_depIdxs = nil
}
