// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: protobuf/user.proto

package userpb

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

// GetProfile
type GetProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetProfileRequest) Reset() {
	*x = GetProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileRequest) ProtoMessage() {}

func (x *GetProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileRequest.ProtoReflect.Descriptor instead.
func (*GetProfileRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_user_proto_rawDescGZIP(), []int{0}
}

func (x *GetProfileRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Date string `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *GetProfileResponse) Reset() {
	*x = GetProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileResponse) ProtoMessage() {}

func (x *GetProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileResponse.ProtoReflect.Descriptor instead.
func (*GetProfileResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetProfileResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetProfileResponse) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

// SendOTPEmail
type SendOtpEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
}

func (x *SendOtpEmailRequest) Reset() {
	*x = SendOtpEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendOtpEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendOtpEmailRequest) ProtoMessage() {}

func (x *SendOtpEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendOtpEmailRequest.ProtoReflect.Descriptor instead.
func (*SendOtpEmailRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_user_proto_rawDescGZIP(), []int{2}
}

func (x *SendOtpEmailRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

type SendOtpEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *SendOtpEmailResponse) Reset() {
	*x = SendOtpEmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendOtpEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendOtpEmailResponse) ProtoMessage() {}

func (x *SendOtpEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendOtpEmailResponse.ProtoReflect.Descriptor instead.
func (*SendOtpEmailResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_user_proto_rawDescGZIP(), []int{3}
}

func (x *SendOtpEmailResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// VerifyOtp
type VerifyOtpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Otp   string `protobuf:"bytes,2,opt,name=otp,proto3" json:"otp,omitempty"`
}

func (x *VerifyOtpRequest) Reset() {
	*x = VerifyOtpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyOtpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyOtpRequest) ProtoMessage() {}

func (x *VerifyOtpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyOtpRequest.ProtoReflect.Descriptor instead.
func (*VerifyOtpRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_user_proto_rawDescGZIP(), []int{4}
}

func (x *VerifyOtpRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *VerifyOtpRequest) GetOtp() string {
	if x != nil {
		return x.Otp
	}
	return ""
}

// RefreshTokens
type RefreshTokensRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshUUID string `protobuf:"bytes,1,opt,name=refreshUUID,proto3" json:"refreshUUID,omitempty"`
}

func (x *RefreshTokensRequest) Reset() {
	*x = RefreshTokensRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshTokensRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokensRequest) ProtoMessage() {}

func (x *RefreshTokensRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokensRequest.ProtoReflect.Descriptor instead.
func (*RefreshTokensRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_user_proto_rawDescGZIP(), []int{5}
}

func (x *RefreshTokensRequest) GetRefreshUUID() string {
	if x != nil {
		return x.RefreshUUID
	}
	return ""
}

// Tokens
type TokensResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshUUID string `protobuf:"bytes,1,opt,name=refreshUUID,proto3" json:"refreshUUID,omitempty"`
	AccessJWT   string `protobuf:"bytes,2,opt,name=accessJWT,proto3" json:"accessJWT,omitempty"`
}

func (x *TokensResponse) Reset() {
	*x = TokensResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokensResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokensResponse) ProtoMessage() {}

func (x *TokensResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokensResponse.ProtoReflect.Descriptor instead.
func (*TokensResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_user_proto_rawDescGZIP(), []int{6}
}

func (x *TokensResponse) GetRefreshUUID() string {
	if x != nil {
		return x.RefreshUUID
	}
	return ""
}

func (x *TokensResponse) GetAccessJWT() string {
	if x != nil {
		return x.AccessJWT
	}
	return ""
}

var File_protobuf_user_proto protoreflect.FileDescriptor

var file_protobuf_user_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x22, 0x27, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x22, 0x2b, 0x0a, 0x13, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x74, 0x70, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x22, 0x2c, 0x0a, 0x14, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x74, 0x70, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22,
	0x3a, 0x0a, 0x10, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x74, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x74, 0x70, 0x22, 0x38, 0x0a, 0x14, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x55, 0x55,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x55, 0x55, 0x49, 0x44, 0x22, 0x50, 0x0a, 0x0e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x55, 0x55, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x4a, 0x57, 0x54, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x4a, 0x57, 0x54, 0x32, 0xab, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b,
	0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x74, 0x70, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1b,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x74, 0x70, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x74, 0x70, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x09, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x74, 0x70, 0x12, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70,
	0x62, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0d,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x1c, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x70, 0x62, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_user_proto_rawDescOnce sync.Once
	file_protobuf_user_proto_rawDescData = file_protobuf_user_proto_rawDesc
)

func file_protobuf_user_proto_rawDescGZIP() []byte {
	file_protobuf_user_proto_rawDescOnce.Do(func() {
		file_protobuf_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_user_proto_rawDescData)
	})
	return file_protobuf_user_proto_rawDescData
}

var file_protobuf_user_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_protobuf_user_proto_goTypes = []interface{}{
	(*GetProfileRequest)(nil),    // 0: userpb.GetProfileRequest
	(*GetProfileResponse)(nil),   // 1: userpb.GetProfileResponse
	(*SendOtpEmailRequest)(nil),  // 2: userpb.SendOtpEmailRequest
	(*SendOtpEmailResponse)(nil), // 3: userpb.SendOtpEmailResponse
	(*VerifyOtpRequest)(nil),     // 4: userpb.VerifyOtpRequest
	(*RefreshTokensRequest)(nil), // 5: userpb.RefreshTokensRequest
	(*TokensResponse)(nil),       // 6: userpb.TokensResponse
}
var file_protobuf_user_proto_depIdxs = []int32{
	0, // 0: userpb.UserService.GetProfile:input_type -> userpb.GetProfileRequest
	2, // 1: userpb.UserService.SendOtpEmail:input_type -> userpb.SendOtpEmailRequest
	4, // 2: userpb.UserService.VerifyOtp:input_type -> userpb.VerifyOtpRequest
	5, // 3: userpb.UserService.RefreshTokens:input_type -> userpb.RefreshTokensRequest
	1, // 4: userpb.UserService.GetProfile:output_type -> userpb.GetProfileResponse
	3, // 5: userpb.UserService.SendOtpEmail:output_type -> userpb.SendOtpEmailResponse
	6, // 6: userpb.UserService.VerifyOtp:output_type -> userpb.TokensResponse
	6, // 7: userpb.UserService.RefreshTokens:output_type -> userpb.TokensResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_user_proto_init() }
func file_protobuf_user_proto_init() {
	if File_protobuf_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfileRequest); i {
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
		file_protobuf_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfileResponse); i {
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
		file_protobuf_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendOtpEmailRequest); i {
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
		file_protobuf_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendOtpEmailResponse); i {
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
		file_protobuf_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyOtpRequest); i {
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
		file_protobuf_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshTokensRequest); i {
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
		file_protobuf_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokensResponse); i {
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
			RawDescriptor: file_protobuf_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_user_proto_goTypes,
		DependencyIndexes: file_protobuf_user_proto_depIdxs,
		MessageInfos:      file_protobuf_user_proto_msgTypes,
	}.Build()
	File_protobuf_user_proto = out.File
	file_protobuf_user_proto_rawDesc = nil
	file_protobuf_user_proto_goTypes = nil
	file_protobuf_user_proto_depIdxs = nil
}
