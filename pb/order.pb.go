// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: pb/order.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Order_OrderStatus int32

const (
	Order_UNKNOWN  Order_OrderStatus = 0
	Order_SETTLED  Order_OrderStatus = 1
	Order_PAID     Order_OrderStatus = 3
	Order_CANCELED Order_OrderStatus = 4
)

// Enum value maps for Order_OrderStatus.
var (
	Order_OrderStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "SETTLED",
		3: "PAID",
		4: "CANCELED",
	}
	Order_OrderStatus_value = map[string]int32{
		"UNKNOWN":  0,
		"SETTLED":  1,
		"PAID":     3,
		"CANCELED": 4,
	}
)

func (x Order_OrderStatus) Enum() *Order_OrderStatus {
	p := new(Order_OrderStatus)
	*p = x
	return p
}

func (x Order_OrderStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Order_OrderStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_order_proto_enumTypes[0].Descriptor()
}

func (Order_OrderStatus) Type() protoreflect.EnumType {
	return &file_pb_order_proto_enumTypes[0]
}

func (x Order_OrderStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Order_OrderStatus.Descriptor instead.
func (Order_OrderStatus) EnumDescriptor() ([]byte, []int) {
	return file_pb_order_proto_rawDescGZIP(), []int{1, 0}
}

type OrderItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId         uint64 `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	ItemName        string `protobuf:"bytes,3,opt,name=item_name,json=itemName,proto3" json:"item_name,omitempty"`
	Quantity        uint32 `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	UnitPrice       uint32 `protobuf:"varint,5,opt,name=unit_price,json=unitPrice,proto3" json:"unit_price,omitempty"`
	ItemDescription string `protobuf:"bytes,6,opt,name=item_description,json=itemDescription,proto3" json:"item_description,omitempty"`
}

func (x *OrderItem) Reset() {
	*x = OrderItem{}
	mi := &file_pb_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItem) ProtoMessage() {}

func (x *OrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_pb_order_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItem.ProtoReflect.Descriptor instead.
func (*OrderItem) Descriptor() ([]byte, []int) {
	return file_pb_order_proto_rawDescGZIP(), []int{0}
}

func (x *OrderItem) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderItem) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *OrderItem) GetItemName() string {
	if x != nil {
		return x.ItemName
	}
	return ""
}

func (x *OrderItem) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *OrderItem) GetUnitPrice() uint32 {
	if x != nil {
		return x.UnitPrice
	}
	return 0
}

func (x *OrderItem) GetItemDescription() string {
	if x != nil {
		return x.ItemDescription
	}
	return ""
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Quantity  uint32 `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	TotalBill uint32 `protobuf:"varint,3,opt,name=total_bill,json=totalBill,proto3" json:"total_bill,omitempty"`
	// string ipg_method = 4 [default = "saman", json_name = "ipg"];
	IpgMethod     string       `protobuf:"bytes,4,opt,name=ipg_method,json=ipgMethod,proto3" json:"ipg_method,omitempty"`
	VoucherCode   string       `protobuf:"bytes,5,opt,name=voucher_code,json=voucherCode,proto3" json:"voucher_code,omitempty"`
	VoucherAmount string       `protobuf:"bytes,6,opt,name=voucher_amount,json=voucherAmount,proto3" json:"voucher_amount,omitempty"`
	Items         []*OrderItem `protobuf:"bytes,7,rep,name=items,proto3" json:"items,omitempty"`
	// OrderStatus OBSOLETE_status = 8;
	Status  Order_OrderStatus `protobuf:"varint,8,opt,name=status,proto3,enum=Order_OrderStatus" json:"status,omitempty"`
	User    *User             `protobuf:"bytes,10,opt,name=user,proto3" json:"user,omitempty"`
	Details []*anypb.Any      `protobuf:"bytes,11,rep,name=details,proto3" json:"details,omitempty"`
	// Types that are assignable to ExternalPayment:
	//
	//	*Order_DigiPay
	//	*Order_SnappPay
	ExternalPayment isOrder_ExternalPayment `protobuf_oneof:"external_payment"`
	Headers         map[string]string       `protobuf:"bytes,14,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_pb_order_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_pb_order_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_pb_order_proto_rawDescGZIP(), []int{1}
}

func (x *Order) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Order) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Order) GetTotalBill() uint32 {
	if x != nil {
		return x.TotalBill
	}
	return 0
}

func (x *Order) GetIpgMethod() string {
	if x != nil {
		return x.IpgMethod
	}
	return ""
}

func (x *Order) GetVoucherCode() string {
	if x != nil {
		return x.VoucherCode
	}
	return ""
}

func (x *Order) GetVoucherAmount() string {
	if x != nil {
		return x.VoucherAmount
	}
	return ""
}

func (x *Order) GetItems() []*OrderItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Order) GetStatus() Order_OrderStatus {
	if x != nil {
		return x.Status
	}
	return Order_UNKNOWN
}

func (x *Order) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Order) GetDetails() []*anypb.Any {
	if x != nil {
		return x.Details
	}
	return nil
}

func (m *Order) GetExternalPayment() isOrder_ExternalPayment {
	if m != nil {
		return m.ExternalPayment
	}
	return nil
}

func (x *Order) GetDigiPay() uint32 {
	if x, ok := x.GetExternalPayment().(*Order_DigiPay); ok {
		return x.DigiPay
	}
	return 0
}

func (x *Order) GetSnappPay() uint32 {
	if x, ok := x.GetExternalPayment().(*Order_SnappPay); ok {
		return x.SnappPay
	}
	return 0
}

func (x *Order) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

type isOrder_ExternalPayment interface {
	isOrder_ExternalPayment()
}

type Order_DigiPay struct {
	DigiPay uint32 `protobuf:"varint,12,opt,name=digi_pay,json=digiPay,proto3,oneof"`
}

type Order_SnappPay struct {
	SnappPay uint32 `protobuf:"varint,13,opt,name=snapp_pay,json=snappPay,proto3,oneof"`
}

func (*Order_DigiPay) isOrder_ExternalPayment() {}

func (*Order_SnappPay) isOrder_ExternalPayment() {}

type GetOrderFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetOrderFilter) Reset() {
	*x = GetOrderFilter{}
	mi := &file_pb_order_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderFilter) ProtoMessage() {}

func (x *GetOrderFilter) ProtoReflect() protoreflect.Message {
	mi := &file_pb_order_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderFilter.ProtoReflect.Descriptor instead.
func (*GetOrderFilter) Descriptor() ([]byte, []int) {
	return file_pb_order_proto_rawDescGZIP(), []int{2}
}

func (x *GetOrderFilter) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *GetOrderResponse) Reset() {
	*x = GetOrderResponse{}
	mi := &file_pb_order_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderResponse) ProtoMessage() {}

func (x *GetOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_order_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderResponse.ProtoReflect.Descriptor instead.
func (*GetOrderResponse) Descriptor() ([]byte, []int) {
	return file_pb_order_proto_rawDescGZIP(), []int{3}
}

func (x *GetOrderResponse) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

var File_pb_order_proto protoreflect.FileDescriptor

var file_pb_order_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x62, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0d, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x01, 0x0a, 0x09, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x6e, 0x69, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x69,
	0x74, 0x65, 0x6d, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x74, 0x65, 0x6d, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xdd, 0x04, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x62, 0x69, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x69,
	0x70, 0x67, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x69, 0x70, 0x67, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x6f,
	0x75, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x19, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x2e, 0x0a,
	0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1b, 0x0a,
	0x08, 0x64, 0x69, 0x67, 0x69, 0x5f, 0x70, 0x61, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x48,
	0x00, 0x52, 0x07, 0x64, 0x69, 0x67, 0x69, 0x50, 0x61, 0x79, 0x12, 0x1d, 0x0a, 0x09, 0x73, 0x6e,
	0x61, 0x70, 0x70, 0x5f, 0x70, 0x61, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52,
	0x08, 0x73, 0x6e, 0x61, 0x70, 0x70, 0x50, 0x61, 0x79, 0x12, 0x2d, 0x0a, 0x07, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x46, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x45, 0x54, 0x54, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x08, 0x0a,
	0x04, 0x50, 0x41, 0x49, 0x44, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41, 0x4e, 0x43, 0x45,
	0x4c, 0x45, 0x44, 0x10, 0x04, 0x2a, 0x05, 0x46, 0x4f, 0x4c, 0x41, 0x4e, 0x42, 0x12, 0x0a, 0x10,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x4a, 0x04, 0x08, 0x09, 0x10, 0x0a, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x06,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x32, 0x3e, 0x0a, 0x0c,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_order_proto_rawDescOnce sync.Once
	file_pb_order_proto_rawDescData = file_pb_order_proto_rawDesc
)

func file_pb_order_proto_rawDescGZIP() []byte {
	file_pb_order_proto_rawDescOnce.Do(func() {
		file_pb_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_order_proto_rawDescData)
	})
	return file_pb_order_proto_rawDescData
}

var file_pb_order_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_order_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pb_order_proto_goTypes = []any{
	(Order_OrderStatus)(0),   // 0: Order.OrderStatus
	(*OrderItem)(nil),        // 1: OrderItem
	(*Order)(nil),            // 2: Order
	(*GetOrderFilter)(nil),   // 3: GetOrderFilter
	(*GetOrderResponse)(nil), // 4: GetOrderResponse
	nil,                      // 5: Order.HeadersEntry
	(*User)(nil),             // 6: User
	(*anypb.Any)(nil),        // 7: google.protobuf.Any
}
var file_pb_order_proto_depIdxs = []int32{
	1, // 0: Order.items:type_name -> OrderItem
	0, // 1: Order.status:type_name -> Order.OrderStatus
	6, // 2: Order.user:type_name -> User
	7, // 3: Order.details:type_name -> google.protobuf.Any
	5, // 4: Order.headers:type_name -> Order.HeadersEntry
	2, // 5: GetOrderResponse.orders:type_name -> Order
	3, // 6: OrderService.GetOrder:input_type -> GetOrderFilter
	4, // 7: OrderService.GetOrder:output_type -> GetOrderResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pb_order_proto_init() }
func file_pb_order_proto_init() {
	if File_pb_order_proto != nil {
		return
	}
	file_pb_user_proto_init()
	file_pb_order_proto_msgTypes[1].OneofWrappers = []any{
		(*Order_DigiPay)(nil),
		(*Order_SnappPay)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_order_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_order_proto_goTypes,
		DependencyIndexes: file_pb_order_proto_depIdxs,
		EnumInfos:         file_pb_order_proto_enumTypes,
		MessageInfos:      file_pb_order_proto_msgTypes,
	}.Build()
	File_pb_order_proto = out.File
	file_pb_order_proto_rawDesc = nil
	file_pb_order_proto_goTypes = nil
	file_pb_order_proto_depIdxs = nil
}
