// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.13.0
// source: proto/user.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type DouyinUserRegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"` //注册同户名，最长32个字符
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"` // 密码，最长32个字符
}

func (x *DouyinUserRegisterRequest) Reset() {
	*x = DouyinUserRegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinUserRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinUserRegisterRequest) ProtoMessage() {}

func (x *DouyinUserRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinUserRegisterRequest.ProtoReflect.Descriptor instead.
func (*DouyinUserRegisterRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{0}
}

func (x *DouyinUserRegisterRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *DouyinUserRegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type DouyinUserRegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败。
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`     //返回状态描述
	UserId     int64  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`             //用户id
	Token      string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`                              //用户鉴权
}

func (x *DouyinUserRegisterResponse) Reset() {
	*x = DouyinUserRegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinUserRegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinUserRegisterResponse) ProtoMessage() {}

func (x *DouyinUserRegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinUserRegisterResponse.ProtoReflect.Descriptor instead.
func (*DouyinUserRegisterResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{1}
}

func (x *DouyinUserRegisterResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *DouyinUserRegisterResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *DouyinUserRegisterResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DouyinUserRegisterResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type DouyinUserLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username          string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`                   //登录用户名
	Password          string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`                   // 登录密码
	EncryptedPassword string `protobuf:"bytes,3,opt,name=encryptedPassword,proto3" json:"encryptedPassword,omitempty"` //加密
}

func (x *DouyinUserLoginRequest) Reset() {
	*x = DouyinUserLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinUserLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinUserLoginRequest) ProtoMessage() {}

func (x *DouyinUserLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinUserLoginRequest.ProtoReflect.Descriptor instead.
func (*DouyinUserLoginRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{2}
}

func (x *DouyinUserLoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *DouyinUserLoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *DouyinUserLoginRequest) GetEncryptedPassword() string {
	if x != nil {
		return x.EncryptedPassword
	}
	return ""
}

type DouyinUserLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` // 状态吗，0-成功,其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`     //返回状态描述
	UserId     int64  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`             //用户id
	Token      string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`                              //用户鉴权token
}

func (x *DouyinUserLoginResponse) Reset() {
	*x = DouyinUserLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinUserLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinUserLoginResponse) ProtoMessage() {}

func (x *DouyinUserLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinUserLoginResponse.ProtoReflect.Descriptor instead.
func (*DouyinUserLoginResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{3}
}

func (x *DouyinUserLoginResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *DouyinUserLoginResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *DouyinUserLoginResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DouyinUserLoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                            //用户id
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                         // 用户名称
	FollowCount   int64  `protobuf:"varint,3,opt,name=follow_count,json=followCount,proto3" json:"follow_count,omitempty"`       // 关注总数
	FollowerCount int64  `protobuf:"varint,4,opt,name=follower_count,json=followerCount,proto3" json:"follower_count,omitempty"` // 粉丝总数
	IsFollow      bool   `protobuf:"varint,5,opt,name=is_follow,json=isFollow,proto3" json:"is_follow,omitempty"`                // true-已关注，false-未关注。
	Password      string `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`                                 //用户密码
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_proto_user_proto_rawDescGZIP(), []int{4}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetFollowCount() int64 {
	if x != nil {
		return x.FollowCount
	}
	return 0
}

func (x *User) GetFollowerCount() int64 {
	if x != nil {
		return x.FollowerCount
	}
	return 0
}

func (x *User) GetIsFollow() bool {
	if x != nil {
		return x.IsFollow
	}
	return false
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type DouyinUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` //用户id
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`                  // 用户鉴权token
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`                    //用户名称
}

func (x *DouyinUserRequest) Reset() {
	*x = DouyinUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinUserRequest) ProtoMessage() {}

func (x *DouyinUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinUserRequest.ProtoReflect.Descriptor instead.
func (*DouyinUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{5}
}

func (x *DouyinUserRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DouyinUserRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *DouyinUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DouyinUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` // 状态码，0-成功，其他值，失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`     // 返回状态描述
	User       *User  `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`                                // 用户信息
}

func (x *DouyinUserResponse) Reset() {
	*x = DouyinUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinUserResponse) ProtoMessage() {}

func (x *DouyinUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinUserResponse.ProtoReflect.Descriptor instead.
func (*DouyinUserResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{6}
}

func (x *DouyinUserResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *DouyinUserResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *DouyinUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_proto_user_proto protoreflect.FileDescriptor

var file_proto_user_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x56, 0x0a, 0x1c, 0x64, 0x6f, 0x75, 0x79,
	0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x8e, 0x01, 0x0a, 0x1d, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d,
	0x73, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x81, 0x01, 0x0a, 0x19, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x65, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x65, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x8b, 0x01, 0x0a, 0x1a, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x4d, 0x73, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0xad, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73,
	0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69,
	0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x58, 0x0a, 0x13, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x76, 0x0a,
	0x14, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x1e, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x32, 0xca, 0x02, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x53, 0x72,
	0x76, 0x12, 0x44, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x64,
	0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x64, 0x6f,
	0x75, 0x79, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x64,
	0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x0c, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_user_proto_rawDescOnce sync.Once
	file_proto_user_proto_rawDescData = file_proto_user_proto_rawDesc
)

func file_proto_user_proto_rawDescGZIP() []byte {
	file_proto_user_proto_rawDescOnce.Do(func() {
		file_proto_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_user_proto_rawDescData)
	})
	return file_proto_user_proto_rawDescData
}

var file_proto_user_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_user_proto_goTypes = []interface{}{
	(*DouyinUserRegisterRequest)(nil),  // 0: user.douyin_user_register_request
	(*DouyinUserRegisterResponse)(nil), // 1: user.douyin_user_register_response
	(*DouyinUserLoginRequest)(nil),     // 2: user.douyin_user_login_request
	(*DouyinUserLoginResponse)(nil),    // 3: user.douyin_user_login_response
	(*User)(nil),                       // 4: user.User
	(*DouyinUserRequest)(nil),          // 5: user.douyin_user_request
	(*DouyinUserResponse)(nil),         // 6: user.douyin_user_response
}
var file_proto_user_proto_depIdxs = []int32{
	4, // 0: user.douyin_user_response.user:type_name -> user.User
	5, // 1: user.UserSrv.GetUserById:input_type -> user.douyin_user_request
	5, // 2: user.UserSrv.GetUserInfoByName:input_type -> user.douyin_user_request
	2, // 3: user.UserSrv.UserLoginByName:input_type -> user.douyin_user_login_request
	0, // 4: user.UserSrv.UserRegister:input_type -> user.douyin_user_register_request
	6, // 5: user.UserSrv.GetUserById:output_type -> user.douyin_user_response
	6, // 6: user.UserSrv.GetUserInfoByName:output_type -> user.douyin_user_response
	3, // 7: user.UserSrv.UserLoginByName:output_type -> user.douyin_user_login_response
	1, // 8: user.UserSrv.UserRegister:output_type -> user.douyin_user_register_response
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_user_proto_init() }
func file_proto_user_proto_init() {
	if File_proto_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinUserRegisterRequest); i {
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
		file_proto_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinUserRegisterResponse); i {
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
		file_proto_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinUserLoginRequest); i {
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
		file_proto_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinUserLoginResponse); i {
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
		file_proto_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_proto_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinUserRequest); i {
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
		file_proto_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinUserResponse); i {
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
			RawDescriptor: file_proto_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_user_proto_goTypes,
		DependencyIndexes: file_proto_user_proto_depIdxs,
		MessageInfos:      file_proto_user_proto_msgTypes,
	}.Build()
	File_proto_user_proto = out.File
	file_proto_user_proto_rawDesc = nil
	file_proto_user_proto_goTypes = nil
	file_proto_user_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserSrvClient is the client API for UserSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserSrvClient interface {
	//用户信息
	GetUserById(ctx context.Context, in *DouyinUserRequest, opts ...grpc.CallOption) (*DouyinUserResponse, error)
	GetUserInfoByName(ctx context.Context, in *DouyinUserRequest, opts ...grpc.CallOption) (*DouyinUserResponse, error)
	//用户登录
	UserLoginByName(ctx context.Context, in *DouyinUserLoginRequest, opts ...grpc.CallOption) (*DouyinUserLoginResponse, error)
	//用户注册
	UserRegister(ctx context.Context, in *DouyinUserRegisterRequest, opts ...grpc.CallOption) (*DouyinUserRegisterResponse, error)
}

type userSrvClient struct {
	cc grpc.ClientConnInterface
}

func NewUserSrvClient(cc grpc.ClientConnInterface) UserSrvClient {
	return &userSrvClient{cc}
}

func (c *userSrvClient) GetUserById(ctx context.Context, in *DouyinUserRequest, opts ...grpc.CallOption) (*DouyinUserResponse, error) {
	out := new(DouyinUserResponse)
	err := c.cc.Invoke(ctx, "/user.UserSrv/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSrvClient) GetUserInfoByName(ctx context.Context, in *DouyinUserRequest, opts ...grpc.CallOption) (*DouyinUserResponse, error) {
	out := new(DouyinUserResponse)
	err := c.cc.Invoke(ctx, "/user.UserSrv/GetUserInfoByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSrvClient) UserLoginByName(ctx context.Context, in *DouyinUserLoginRequest, opts ...grpc.CallOption) (*DouyinUserLoginResponse, error) {
	out := new(DouyinUserLoginResponse)
	err := c.cc.Invoke(ctx, "/user.UserSrv/UserLoginByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSrvClient) UserRegister(ctx context.Context, in *DouyinUserRegisterRequest, opts ...grpc.CallOption) (*DouyinUserRegisterResponse, error) {
	out := new(DouyinUserRegisterResponse)
	err := c.cc.Invoke(ctx, "/user.UserSrv/UserRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserSrvServer is the server API for UserSrv service.
type UserSrvServer interface {
	//用户信息
	GetUserById(context.Context, *DouyinUserRequest) (*DouyinUserResponse, error)
	GetUserInfoByName(context.Context, *DouyinUserRequest) (*DouyinUserResponse, error)
	//用户登录
	UserLoginByName(context.Context, *DouyinUserLoginRequest) (*DouyinUserLoginResponse, error)
	//用户注册
	UserRegister(context.Context, *DouyinUserRegisterRequest) (*DouyinUserRegisterResponse, error)
}

// UnimplementedUserSrvServer can be embedded to have forward compatible implementations.
type UnimplementedUserSrvServer struct {
}

func (*UnimplementedUserSrvServer) GetUserById(context.Context, *DouyinUserRequest) (*DouyinUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (*UnimplementedUserSrvServer) GetUserInfoByName(context.Context, *DouyinUserRequest) (*DouyinUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfoByName not implemented")
}
func (*UnimplementedUserSrvServer) UserLoginByName(context.Context, *DouyinUserLoginRequest) (*DouyinUserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLoginByName not implemented")
}
func (*UnimplementedUserSrvServer) UserRegister(context.Context, *DouyinUserRegisterRequest) (*DouyinUserRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegister not implemented")
}

func RegisterUserSrvServer(s *grpc.Server, srv UserSrvServer) {
	s.RegisterService(&_UserSrv_serviceDesc, srv)
}

func _UserSrv_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSrvServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserSrv/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSrvServer).GetUserById(ctx, req.(*DouyinUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSrv_GetUserInfoByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSrvServer).GetUserInfoByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserSrv/GetUserInfoByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSrvServer).GetUserInfoByName(ctx, req.(*DouyinUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSrv_UserLoginByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinUserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSrvServer).UserLoginByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserSrv/UserLoginByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSrvServer).UserLoginByName(ctx, req.(*DouyinUserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSrv_UserRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinUserRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSrvServer).UserRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserSrv/UserRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSrvServer).UserRegister(ctx, req.(*DouyinUserRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserSrv_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserSrv",
	HandlerType: (*UserSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserById",
			Handler:    _UserSrv_GetUserById_Handler,
		},
		{
			MethodName: "GetUserInfoByName",
			Handler:    _UserSrv_GetUserInfoByName_Handler,
		},
		{
			MethodName: "UserLoginByName",
			Handler:    _UserSrv_UserLoginByName_Handler,
		},
		{
			MethodName: "UserRegister",
			Handler:    _UserSrv_UserRegister_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user.proto",
}