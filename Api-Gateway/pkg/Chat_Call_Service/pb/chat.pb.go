// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.0--rc2
// source: pkg/Chat_Call_Service/pb/chat.proto

package pb

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

type SingleOneToOneChat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageID  string `protobuf:"bytes,1,opt,name=messageID,proto3" json:"messageID,omitempty"`
	SenderID   string `protobuf:"bytes,2,opt,name=SenderID,proto3" json:"SenderID,omitempty"`
	RecieverID string `protobuf:"bytes,3,opt,name=RecieverID,proto3" json:"RecieverID,omitempty"`
	Content    string `protobuf:"bytes,4,opt,name=Content,proto3" json:"Content,omitempty"`
	Status     string `protobuf:"bytes,5,opt,name=Status,proto3" json:"Status,omitempty"`
	TimeStamp  string `protobuf:"bytes,6,opt,name=TimeStamp,proto3" json:"TimeStamp,omitempty"`
}

func (x *SingleOneToOneChat) Reset() {
	*x = SingleOneToOneChat{}
	mi := &file_pkg_Chat_Call_Service_pb_chat_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SingleOneToOneChat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleOneToOneChat) ProtoMessage() {}

func (x *SingleOneToOneChat) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_Chat_Call_Service_pb_chat_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleOneToOneChat.ProtoReflect.Descriptor instead.
func (*SingleOneToOneChat) Descriptor() ([]byte, []int) {
	return file_pkg_Chat_Call_Service_pb_chat_proto_rawDescGZIP(), []int{0}
}

func (x *SingleOneToOneChat) GetMessageID() string {
	if x != nil {
		return x.MessageID
	}
	return ""
}

func (x *SingleOneToOneChat) GetSenderID() string {
	if x != nil {
		return x.SenderID
	}
	return ""
}

func (x *SingleOneToOneChat) GetRecieverID() string {
	if x != nil {
		return x.RecieverID
	}
	return ""
}

func (x *SingleOneToOneChat) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SingleOneToOneChat) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *SingleOneToOneChat) GetTimeStamp() string {
	if x != nil {
		return x.TimeStamp
	}
	return ""
}

type ResponseUserOneToOneChat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chat         []*SingleOneToOneChat `protobuf:"bytes,1,rep,name=Chat,proto3" json:"Chat,omitempty"`
	ErrorMessage string                `protobuf:"bytes,2,opt,name=ErrorMessage,proto3" json:"ErrorMessage,omitempty"`
}

func (x *ResponseUserOneToOneChat) Reset() {
	*x = ResponseUserOneToOneChat{}
	mi := &file_pkg_Chat_Call_Service_pb_chat_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResponseUserOneToOneChat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseUserOneToOneChat) ProtoMessage() {}

func (x *ResponseUserOneToOneChat) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_Chat_Call_Service_pb_chat_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseUserOneToOneChat.ProtoReflect.Descriptor instead.
func (*ResponseUserOneToOneChat) Descriptor() ([]byte, []int) {
	return file_pkg_Chat_Call_Service_pb_chat_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseUserOneToOneChat) GetChat() []*SingleOneToOneChat {
	if x != nil {
		return x.Chat
	}
	return nil
}

func (x *ResponseUserOneToOneChat) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type RequestUserOneToOneChat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderID   string `protobuf:"bytes,1,opt,name=SenderID,proto3" json:"SenderID,omitempty"`
	RecieverID string `protobuf:"bytes,2,opt,name=RecieverID,proto3" json:"RecieverID,omitempty"`
	Limit      string `protobuf:"bytes,3,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Offset     string `protobuf:"bytes,4,opt,name=Offset,proto3" json:"Offset,omitempty"`
}

func (x *RequestUserOneToOneChat) Reset() {
	*x = RequestUserOneToOneChat{}
	mi := &file_pkg_Chat_Call_Service_pb_chat_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestUserOneToOneChat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestUserOneToOneChat) ProtoMessage() {}

func (x *RequestUserOneToOneChat) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_Chat_Call_Service_pb_chat_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestUserOneToOneChat.ProtoReflect.Descriptor instead.
func (*RequestUserOneToOneChat) Descriptor() ([]byte, []int) {
	return file_pkg_Chat_Call_Service_pb_chat_proto_rawDescGZIP(), []int{2}
}

func (x *RequestUserOneToOneChat) GetSenderID() string {
	if x != nil {
		return x.SenderID
	}
	return ""
}

func (x *RequestUserOneToOneChat) GetRecieverID() string {
	if x != nil {
		return x.RecieverID
	}
	return ""
}

func (x *RequestUserOneToOneChat) GetLimit() string {
	if x != nil {
		return x.Limit
	}
	return ""
}

func (x *RequestUserOneToOneChat) GetOffset() string {
	if x != nil {
		return x.Offset
	}
	return ""
}

type RequestRecentChatProfiles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderID string `protobuf:"bytes,1,opt,name=SenderID,proto3" json:"SenderID,omitempty"`
	Limit    string `protobuf:"bytes,2,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Offset   string `protobuf:"bytes,3,opt,name=Offset,proto3" json:"Offset,omitempty"`
}

func (x *RequestRecentChatProfiles) Reset() {
	*x = RequestRecentChatProfiles{}
	mi := &file_pkg_Chat_Call_Service_pb_chat_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestRecentChatProfiles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestRecentChatProfiles) ProtoMessage() {}

func (x *RequestRecentChatProfiles) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_Chat_Call_Service_pb_chat_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestRecentChatProfiles.ProtoReflect.Descriptor instead.
func (*RequestRecentChatProfiles) Descriptor() ([]byte, []int) {
	return file_pkg_Chat_Call_Service_pb_chat_proto_rawDescGZIP(), []int{3}
}

func (x *RequestRecentChatProfiles) GetSenderID() string {
	if x != nil {
		return x.SenderID
	}
	return ""
}

func (x *RequestRecentChatProfiles) GetLimit() string {
	if x != nil {
		return x.Limit
	}
	return ""
}

func (x *RequestRecentChatProfiles) GetOffset() string {
	if x != nil {
		return x.Offset
	}
	return ""
}

type SingelUserAndLastChat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID               string `protobuf:"bytes,5,opt,name=UserID,proto3" json:"UserID,omitempty"`
	UserName             string `protobuf:"bytes,1,opt,name=UserName,proto3" json:"UserName,omitempty"`
	UserProfileURL       string `protobuf:"bytes,2,opt,name=UserProfileURL,proto3" json:"UserProfileURL,omitempty"`
	LastMessageContent   string `protobuf:"bytes,3,opt,name=LastMessageContent,proto3" json:"LastMessageContent,omitempty"`
	LastMessageTimeStamp string `protobuf:"bytes,4,opt,name=LastMessageTimeStamp,proto3" json:"LastMessageTimeStamp,omitempty"`
}

func (x *SingelUserAndLastChat) Reset() {
	*x = SingelUserAndLastChat{}
	mi := &file_pkg_Chat_Call_Service_pb_chat_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SingelUserAndLastChat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingelUserAndLastChat) ProtoMessage() {}

func (x *SingelUserAndLastChat) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_Chat_Call_Service_pb_chat_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingelUserAndLastChat.ProtoReflect.Descriptor instead.
func (*SingelUserAndLastChat) Descriptor() ([]byte, []int) {
	return file_pkg_Chat_Call_Service_pb_chat_proto_rawDescGZIP(), []int{4}
}

func (x *SingelUserAndLastChat) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *SingelUserAndLastChat) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *SingelUserAndLastChat) GetUserProfileURL() string {
	if x != nil {
		return x.UserProfileURL
	}
	return ""
}

func (x *SingelUserAndLastChat) GetLastMessageContent() string {
	if x != nil {
		return x.LastMessageContent
	}
	return ""
}

func (x *SingelUserAndLastChat) GetLastMessageTimeStamp() string {
	if x != nil {
		return x.LastMessageTimeStamp
	}
	return ""
}

type ResponseRecentChatProfiles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorMessage string                   `protobuf:"bytes,1,opt,name=ErrorMessage,proto3" json:"ErrorMessage,omitempty"`
	ActualData   []*SingelUserAndLastChat `protobuf:"bytes,2,rep,name=ActualData,proto3" json:"ActualData,omitempty"`
}

func (x *ResponseRecentChatProfiles) Reset() {
	*x = ResponseRecentChatProfiles{}
	mi := &file_pkg_Chat_Call_Service_pb_chat_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResponseRecentChatProfiles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseRecentChatProfiles) ProtoMessage() {}

func (x *ResponseRecentChatProfiles) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_Chat_Call_Service_pb_chat_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseRecentChatProfiles.ProtoReflect.Descriptor instead.
func (*ResponseRecentChatProfiles) Descriptor() ([]byte, []int) {
	return file_pkg_Chat_Call_Service_pb_chat_proto_rawDescGZIP(), []int{5}
}

func (x *ResponseRecentChatProfiles) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *ResponseRecentChatProfiles) GetActualData() []*SingelUserAndLastChat {
	if x != nil {
		return x.ActualData
	}
	return nil
}

var File_pkg_Chat_Call_Service_pb_chat_proto protoreflect.FileDescriptor

var file_pkg_Chat_Call_Service_pb_chat_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x6b, 0x67, 0x2f, 0x43, 0x68, 0x61, 0x74, 0x5f, 0x43, 0x61, 0x6c, 0x6c, 0x5f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x68, 0x61, 0x74, 0x4e, 0x63, 0x61, 0x6c, 0x6c,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x01, 0x0a, 0x12, 0x53, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x4f, 0x6e, 0x65, 0x54, 0x6f, 0x4f, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x53,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x63, 0x69, 0x65,
	0x76, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x52, 0x65, 0x63,
	0x69, 0x65, 0x76, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d,
	0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x69,
	0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x77, 0x0a, 0x18, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x6e, 0x65, 0x54, 0x6f, 0x4f, 0x6e, 0x65, 0x43,
	0x68, 0x61, 0x74, 0x12, 0x37, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x4e, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4f, 0x6e, 0x65, 0x54, 0x6f, 0x4f,
	0x6e, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x22, 0x0a, 0x0c,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x83, 0x01, 0x0a, 0x17, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x4f, 0x6e, 0x65, 0x54, 0x6f, 0x4f, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x63, 0x69,
	0x65, 0x76, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x52, 0x65,
	0x63, 0x69, 0x65, 0x76, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x65, 0x0a, 0x19, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0xd7, 0x01,
	0x0a, 0x15, 0x53, 0x69, 0x6e, 0x67, 0x65, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x4c,
	0x61, 0x73, 0x74, 0x43, 0x68, 0x61, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x55, 0x52, 0x4c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x55, 0x52, 0x4c, 0x12, 0x2e, 0x0a, 0x12, 0x4c, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x12, 0x4c, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x14, 0x4c, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x14, 0x4c, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x88, 0x01, 0x0a, 0x1a, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x74, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x46, 0x0a, 0x0a, 0x41, 0x63,
	0x74, 0x75, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x4e, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x69, 0x6e, 0x67, 0x65, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x4c, 0x61,
	0x73, 0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x0a, 0x41, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x44, 0x61,
	0x74, 0x61, 0x32, 0xed, 0x01, 0x0a, 0x10, 0x43, 0x68, 0x61, 0x74, 0x4e, 0x43, 0x61, 0x6c, 0x6c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x67, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4f, 0x6e,
	0x65, 0x54, 0x6f, 0x4f, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x74, 0x73, 0x12, 0x28, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x4e, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x6e, 0x65, 0x54, 0x6f, 0x4f, 0x6e,
	0x65, 0x43, 0x68, 0x61, 0x74, 0x1a, 0x29, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x4e, 0x63, 0x61, 0x6c,
	0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x4f, 0x6e, 0x65, 0x54, 0x6f, 0x4f, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x74,
	0x12, 0x70, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61,
	0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x2a, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x4e, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x74, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x1a, 0x2b, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x4e, 0x63, 0x61, 0x6c,
	0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x42, 0x1c, 0x5a, 0x1a, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x43, 0x68, 0x61, 0x74,
	0x5f, 0x43, 0x61, 0x6c, 0x6c, 0x5f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_Chat_Call_Service_pb_chat_proto_rawDescOnce sync.Once
	file_pkg_Chat_Call_Service_pb_chat_proto_rawDescData = file_pkg_Chat_Call_Service_pb_chat_proto_rawDesc
)

func file_pkg_Chat_Call_Service_pb_chat_proto_rawDescGZIP() []byte {
	file_pkg_Chat_Call_Service_pb_chat_proto_rawDescOnce.Do(func() {
		file_pkg_Chat_Call_Service_pb_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_Chat_Call_Service_pb_chat_proto_rawDescData)
	})
	return file_pkg_Chat_Call_Service_pb_chat_proto_rawDescData
}

var file_pkg_Chat_Call_Service_pb_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pkg_Chat_Call_Service_pb_chat_proto_goTypes = []any{
	(*SingleOneToOneChat)(nil),         // 0: chatNcall_proto.SingleOneToOneChat
	(*ResponseUserOneToOneChat)(nil),   // 1: chatNcall_proto.ResponseUserOneToOneChat
	(*RequestUserOneToOneChat)(nil),    // 2: chatNcall_proto.RequestUserOneToOneChat
	(*RequestRecentChatProfiles)(nil),  // 3: chatNcall_proto.RequestRecentChatProfiles
	(*SingelUserAndLastChat)(nil),      // 4: chatNcall_proto.SingelUserAndLastChat
	(*ResponseRecentChatProfiles)(nil), // 5: chatNcall_proto.ResponseRecentChatProfiles
}
var file_pkg_Chat_Call_Service_pb_chat_proto_depIdxs = []int32{
	0, // 0: chatNcall_proto.ResponseUserOneToOneChat.Chat:type_name -> chatNcall_proto.SingleOneToOneChat
	4, // 1: chatNcall_proto.ResponseRecentChatProfiles.ActualData:type_name -> chatNcall_proto.SingelUserAndLastChat
	2, // 2: chatNcall_proto.ChatNCallService.GetOneToOneChats:input_type -> chatNcall_proto.RequestUserOneToOneChat
	3, // 3: chatNcall_proto.ChatNCallService.GetRecentChatProfiles:input_type -> chatNcall_proto.RequestRecentChatProfiles
	1, // 4: chatNcall_proto.ChatNCallService.GetOneToOneChats:output_type -> chatNcall_proto.ResponseUserOneToOneChat
	5, // 5: chatNcall_proto.ChatNCallService.GetRecentChatProfiles:output_type -> chatNcall_proto.ResponseRecentChatProfiles
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_Chat_Call_Service_pb_chat_proto_init() }
func file_pkg_Chat_Call_Service_pb_chat_proto_init() {
	if File_pkg_Chat_Call_Service_pb_chat_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_Chat_Call_Service_pb_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_Chat_Call_Service_pb_chat_proto_goTypes,
		DependencyIndexes: file_pkg_Chat_Call_Service_pb_chat_proto_depIdxs,
		MessageInfos:      file_pkg_Chat_Call_Service_pb_chat_proto_msgTypes,
	}.Build()
	File_pkg_Chat_Call_Service_pb_chat_proto = out.File
	file_pkg_Chat_Call_Service_pb_chat_proto_rawDesc = nil
	file_pkg_Chat_Call_Service_pb_chat_proto_goTypes = nil
	file_pkg_Chat_Call_Service_pb_chat_proto_depIdxs = nil
}
