// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: train_ticket.proto

package train

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FirstName     string                 `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName      string                 `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_train_ticket_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_train_ticket_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_train_ticket_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type Ticket struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	From          string                 `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To            string                 `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	User          *User                  `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Price         float64                `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	Seat          string                 `protobuf:"bytes,5,opt,name=seat,proto3" json:"seat,omitempty"`
	Section       string                 `protobuf:"bytes,6,opt,name=section,proto3" json:"section,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Ticket) Reset() {
	*x = Ticket{}
	mi := &file_train_ticket_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ticket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ticket) ProtoMessage() {}

func (x *Ticket) ProtoReflect() protoreflect.Message {
	mi := &file_train_ticket_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ticket.ProtoReflect.Descriptor instead.
func (*Ticket) Descriptor() ([]byte, []int) {
	return file_train_ticket_proto_rawDescGZIP(), []int{1}
}

func (x *Ticket) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Ticket) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Ticket) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Ticket) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Ticket) GetSeat() string {
	if x != nil {
		return x.Seat
	}
	return ""
}

func (x *Ticket) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

type PurchaseRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	From          string                 `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To            string                 `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	User          *User                  `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	PricePaid     float64                `protobuf:"fixed64,4,opt,name=price_paid,json=pricePaid,proto3" json:"price_paid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PurchaseRequest) Reset() {
	*x = PurchaseRequest{}
	mi := &file_train_ticket_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PurchaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchaseRequest) ProtoMessage() {}

func (x *PurchaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_train_ticket_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchaseRequest.ProtoReflect.Descriptor instead.
func (*PurchaseRequest) Descriptor() ([]byte, []int) {
	return file_train_ticket_proto_rawDescGZIP(), []int{2}
}

func (x *PurchaseRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *PurchaseRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *PurchaseRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *PurchaseRequest) GetPricePaid() float64 {
	if x != nil {
		return x.PricePaid
	}
	return 0
}

type PurchaseResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ticket        *Ticket                `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PurchaseResponse) Reset() {
	*x = PurchaseResponse{}
	mi := &file_train_ticket_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PurchaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchaseResponse) ProtoMessage() {}

func (x *PurchaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_train_ticket_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchaseResponse.ProtoReflect.Descriptor instead.
func (*PurchaseResponse) Descriptor() ([]byte, []int) {
	return file_train_ticket_proto_rawDescGZIP(), []int{3}
}

func (x *PurchaseResponse) GetTicket() *Ticket {
	if x != nil {
		return x.Ticket
	}
	return nil
}

func (x *PurchaseResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Receipt struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	From          string                 `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To            string                 `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	User          *User                  `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	PricePaid     float64                `protobuf:"fixed64,4,opt,name=price_paid,json=pricePaid,proto3" json:"price_paid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Receipt) Reset() {
	*x = Receipt{}
	mi := &file_train_ticket_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Receipt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Receipt) ProtoMessage() {}

func (x *Receipt) ProtoReflect() protoreflect.Message {
	mi := &file_train_ticket_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Receipt.ProtoReflect.Descriptor instead.
func (*Receipt) Descriptor() ([]byte, []int) {
	return file_train_ticket_proto_rawDescGZIP(), []int{4}
}

func (x *Receipt) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Receipt) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Receipt) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Receipt) GetPricePaid() float64 {
	if x != nil {
		return x.PricePaid
	}
	return 0
}

type ReceiptRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReceiptRequest) Reset() {
	*x = ReceiptRequest{}
	mi := &file_train_ticket_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReceiptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiptRequest) ProtoMessage() {}

func (x *ReceiptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_train_ticket_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiptRequest.ProtoReflect.Descriptor instead.
func (*ReceiptRequest) Descriptor() ([]byte, []int) {
	return file_train_ticket_proto_rawDescGZIP(), []int{5}
}

func (x *ReceiptRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type ReceiptResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Receipt       *Receipt               `protobuf:"bytes,1,opt,name=receipt,proto3" json:"receipt,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReceiptResponse) Reset() {
	*x = ReceiptResponse{}
	mi := &file_train_ticket_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReceiptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiptResponse) ProtoMessage() {}

func (x *ReceiptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_train_ticket_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiptResponse.ProtoReflect.Descriptor instead.
func (*ReceiptResponse) Descriptor() ([]byte, []int) {
	return file_train_ticket_proto_rawDescGZIP(), []int{6}
}

func (x *ReceiptResponse) GetReceipt() *Receipt {
	if x != nil {
		return x.Receipt
	}
	return nil
}

func (x *ReceiptResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UsersBySectionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Section       string                 `protobuf:"bytes,1,opt,name=section,proto3" json:"section,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UsersBySectionRequest) Reset() {
	*x = UsersBySectionRequest{}
	mi := &file_train_ticket_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UsersBySectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersBySectionRequest) ProtoMessage() {}

func (x *UsersBySectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_train_ticket_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersBySectionRequest.ProtoReflect.Descriptor instead.
func (*UsersBySectionRequest) Descriptor() ([]byte, []int) {
	return file_train_ticket_proto_rawDescGZIP(), []int{7}
}

func (x *UsersBySectionRequest) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

type UserSeatInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Seat          string                 `protobuf:"bytes,2,opt,name=seat,proto3" json:"seat,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserSeatInfo) Reset() {
	*x = UserSeatInfo{}
	mi := &file_train_ticket_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserSeatInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSeatInfo) ProtoMessage() {}

func (x *UserSeatInfo) ProtoReflect() protoreflect.Message {
	mi := &file_train_ticket_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSeatInfo.ProtoReflect.Descriptor instead.
func (*UserSeatInfo) Descriptor() ([]byte, []int) {
	return file_train_ticket_proto_rawDescGZIP(), []int{8}
}

func (x *UserSeatInfo) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserSeatInfo) GetSeat() string {
	if x != nil {
		return x.Seat
	}
	return ""
}

type UsersBySectionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserSeatInfo  []*UserSeatInfo        `protobuf:"bytes,1,rep,name=user_seat_info,json=userSeatInfo,proto3" json:"user_seat_info,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UsersBySectionResponse) Reset() {
	*x = UsersBySectionResponse{}
	mi := &file_train_ticket_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UsersBySectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersBySectionResponse) ProtoMessage() {}

func (x *UsersBySectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_train_ticket_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersBySectionResponse.ProtoReflect.Descriptor instead.
func (*UsersBySectionResponse) Descriptor() ([]byte, []int) {
	return file_train_ticket_proto_rawDescGZIP(), []int{9}
}

func (x *UsersBySectionResponse) GetUserSeatInfo() []*UserSeatInfo {
	if x != nil {
		return x.UserSeatInfo
	}
	return nil
}

func (x *UsersBySectionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type RemoveUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveUserRequest) Reset() {
	*x = RemoveUserRequest{}
	mi := &file_train_ticket_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserRequest) ProtoMessage() {}

func (x *RemoveUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_train_ticket_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserRequest.ProtoReflect.Descriptor instead.
func (*RemoveUserRequest) Descriptor() ([]byte, []int) {
	return file_train_ticket_proto_rawDescGZIP(), []int{10}
}

func (x *RemoveUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type RemoveUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveUserResponse) Reset() {
	*x = RemoveUserResponse{}
	mi := &file_train_ticket_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserResponse) ProtoMessage() {}

func (x *RemoveUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_train_ticket_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserResponse.ProtoReflect.Descriptor instead.
func (*RemoveUserResponse) Descriptor() ([]byte, []int) {
	return file_train_ticket_proto_rawDescGZIP(), []int{11}
}

func (x *RemoveUserResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RemoveUserResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ModifySeatRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	NewSeat       string                 `protobuf:"bytes,2,opt,name=new_seat,json=newSeat,proto3" json:"new_seat,omitempty"`
	NewSection    string                 `protobuf:"bytes,3,opt,name=new_section,json=newSection,proto3" json:"new_section,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ModifySeatRequest) Reset() {
	*x = ModifySeatRequest{}
	mi := &file_train_ticket_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ModifySeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifySeatRequest) ProtoMessage() {}

func (x *ModifySeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_train_ticket_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifySeatRequest.ProtoReflect.Descriptor instead.
func (*ModifySeatRequest) Descriptor() ([]byte, []int) {
	return file_train_ticket_proto_rawDescGZIP(), []int{12}
}

func (x *ModifySeatRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ModifySeatRequest) GetNewSeat() string {
	if x != nil {
		return x.NewSeat
	}
	return ""
}

func (x *ModifySeatRequest) GetNewSection() string {
	if x != nil {
		return x.NewSection
	}
	return ""
}

type ModifySeatResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	NewSeat       string                 `protobuf:"bytes,2,opt,name=new_seat,json=newSeat,proto3" json:"new_seat,omitempty"`
	NewSection    string                 `protobuf:"bytes,3,opt,name=new_section,json=newSection,proto3" json:"new_section,omitempty"`
	Message       string                 `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ModifySeatResponse) Reset() {
	*x = ModifySeatResponse{}
	mi := &file_train_ticket_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ModifySeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifySeatResponse) ProtoMessage() {}

func (x *ModifySeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_train_ticket_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifySeatResponse.ProtoReflect.Descriptor instead.
func (*ModifySeatResponse) Descriptor() ([]byte, []int) {
	return file_train_ticket_proto_rawDescGZIP(), []int{13}
}

func (x *ModifySeatResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ModifySeatResponse) GetNewSeat() string {
	if x != nil {
		return x.NewSeat
	}
	return ""
}

func (x *ModifySeatResponse) GetNewSection() string {
	if x != nil {
		return x.NewSection
	}
	return ""
}

func (x *ModifySeatResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_train_ticket_proto protoreflect.FileDescriptor

var file_train_ticket_proto_rawDesc = string([]byte{
	0x0a, 0x12, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x22, 0x58, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x91, 0x01, 0x0a, 0x06, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x74, 0x6f, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x65, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x65, 0x61, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x75, 0x0a, 0x0f, 0x50, 0x75, 0x72,
	0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f,
	0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x50, 0x61, 0x69, 0x64,
	0x22, 0x53, 0x0a, 0x10, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x6d, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x74, 0x6f, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x70,
	0x61, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x50, 0x61, 0x69, 0x64, 0x22, 0x26, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x55, 0x0a, 0x0f,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x28, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74,
	0x52, 0x07, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x31, 0x0a, 0x15, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x53, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x43, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x61, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x65, 0x61, 0x74, 0x22, 0x6d, 0x0a, 0x16, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x61, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x61, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x29, 0x0a, 0x11, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x48, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x65, 0x0a, 0x11, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x65,
	0x77, 0x5f, 0x73, 0x65, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65,
	0x77, 0x53, 0x65, 0x61, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x5f, 0x73, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x65, 0x77, 0x53,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x84, 0x01, 0x0a, 0x12, 0x4d, 0x6f, 0x64, 0x69, 0x66,
	0x79, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x5f, 0x73,
	0x65, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x77, 0x53, 0x65,
	0x61, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x5f, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x65, 0x77, 0x53, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xec, 0x02,
	0x0a, 0x12, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x2e, 0x50,
	0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x15, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x42, 0x79, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x79, 0x53, 0x65, 0x61, 0x74, 0x12, 0x18, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x2e,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79,
	0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0e, 0x5a, 0x0c,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_train_ticket_proto_rawDescOnce sync.Once
	file_train_ticket_proto_rawDescData []byte
)

func file_train_ticket_proto_rawDescGZIP() []byte {
	file_train_ticket_proto_rawDescOnce.Do(func() {
		file_train_ticket_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_train_ticket_proto_rawDesc), len(file_train_ticket_proto_rawDesc)))
	})
	return file_train_ticket_proto_rawDescData
}

var file_train_ticket_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_train_ticket_proto_goTypes = []any{
	(*User)(nil),                   // 0: train.User
	(*Ticket)(nil),                 // 1: train.Ticket
	(*PurchaseRequest)(nil),        // 2: train.PurchaseRequest
	(*PurchaseResponse)(nil),       // 3: train.PurchaseResponse
	(*Receipt)(nil),                // 4: train.Receipt
	(*ReceiptRequest)(nil),         // 5: train.ReceiptRequest
	(*ReceiptResponse)(nil),        // 6: train.ReceiptResponse
	(*UsersBySectionRequest)(nil),  // 7: train.UsersBySectionRequest
	(*UserSeatInfo)(nil),           // 8: train.UserSeatInfo
	(*UsersBySectionResponse)(nil), // 9: train.UsersBySectionResponse
	(*RemoveUserRequest)(nil),      // 10: train.RemoveUserRequest
	(*RemoveUserResponse)(nil),     // 11: train.RemoveUserResponse
	(*ModifySeatRequest)(nil),      // 12: train.ModifySeatRequest
	(*ModifySeatResponse)(nil),     // 13: train.ModifySeatResponse
}
var file_train_ticket_proto_depIdxs = []int32{
	0,  // 0: train.Ticket.user:type_name -> train.User
	0,  // 1: train.PurchaseRequest.user:type_name -> train.User
	1,  // 2: train.PurchaseResponse.ticket:type_name -> train.Ticket
	0,  // 3: train.Receipt.user:type_name -> train.User
	4,  // 4: train.ReceiptResponse.receipt:type_name -> train.Receipt
	0,  // 5: train.UserSeatInfo.user:type_name -> train.User
	8,  // 6: train.UsersBySectionResponse.user_seat_info:type_name -> train.UserSeatInfo
	2,  // 7: train.TrainTicketService.PurchaseTicket:input_type -> train.PurchaseRequest
	5,  // 8: train.TrainTicketService.GetReceipt:input_type -> train.ReceiptRequest
	7,  // 9: train.TrainTicketService.GetUsersBySection:input_type -> train.UsersBySectionRequest
	10, // 10: train.TrainTicketService.RemoveUser:input_type -> train.RemoveUserRequest
	12, // 11: train.TrainTicketService.ModifySeat:input_type -> train.ModifySeatRequest
	3,  // 12: train.TrainTicketService.PurchaseTicket:output_type -> train.PurchaseResponse
	6,  // 13: train.TrainTicketService.GetReceipt:output_type -> train.ReceiptResponse
	9,  // 14: train.TrainTicketService.GetUsersBySection:output_type -> train.UsersBySectionResponse
	11, // 15: train.TrainTicketService.RemoveUser:output_type -> train.RemoveUserResponse
	13, // 16: train.TrainTicketService.ModifySeat:output_type -> train.ModifySeatResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_train_ticket_proto_init() }
func file_train_ticket_proto_init() {
	if File_train_ticket_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_train_ticket_proto_rawDesc), len(file_train_ticket_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_train_ticket_proto_goTypes,
		DependencyIndexes: file_train_ticket_proto_depIdxs,
		MessageInfos:      file_train_ticket_proto_msgTypes,
	}.Build()
	File_train_ticket_proto = out.File
	file_train_ticket_proto_goTypes = nil
	file_train_ticket_proto_depIdxs = nil
}
