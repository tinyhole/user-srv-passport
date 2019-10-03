// Code generated by protoc-gen-go. DO NOT EDIT.
// source: platform/user_srv_passport.proto

package passport

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	base "github.com/tinyhole/user-srv-passport/idl/base"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Code int32

const (
	Code_Ok     Code = 0
	Code_Failed Code = 1
)

var Code_name = map[int32]string{
	0: "Ok",
	1: "Failed",
}

var Code_value = map[string]int32{
	"Ok":     0,
	"Failed": 1,
}

func (x Code) String() string {
	return proto.EnumName(Code_name, int32(x))
}

func (Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d68c7dc98409413c, []int{0}
}

type BaseRsp struct {
	Code                 Code     `protobuf:"varint,1,opt,name=code,proto3,enum=platform.user.srv.passport.Code" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseRsp) Reset()         { *m = BaseRsp{} }
func (m *BaseRsp) String() string { return proto.CompactTextString(m) }
func (*BaseRsp) ProtoMessage()    {}
func (*BaseRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d68c7dc98409413c, []int{0}
}

func (m *BaseRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseRsp.Unmarshal(m, b)
}
func (m *BaseRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseRsp.Marshal(b, m, deterministic)
}
func (m *BaseRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseRsp.Merge(m, src)
}
func (m *BaseRsp) XXX_Size() int {
	return xxx_messageInfo_BaseRsp.Size(m)
}
func (m *BaseRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseRsp.DiscardUnknown(m)
}

var xxx_messageInfo_BaseRsp proto.InternalMessageInfo

func (m *BaseRsp) GetCode() Code {
	if m != nil {
		return m.Code
	}
	return Code_Ok
}

func (m *BaseRsp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type SignUpReq struct {
	DeviceID             int64           `protobuf:"varint,1,opt,name=DeviceID,proto3" json:"DeviceID,omitempty"`
	Name                 string          `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Mobile               string          `protobuf:"bytes,3,opt,name=Mobile,proto3" json:"Mobile,omitempty"`
	Email                string          `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	Password             string          `protobuf:"bytes,5,opt,name=Password,proto3" json:"Password,omitempty"`
	Code                 string          `protobuf:"bytes,6,opt,name=Code,proto3" json:"Code,omitempty"`
	ActiveType           base.ActiveType `protobuf:"varint,7,opt,name=ActiveType,proto3,enum=base.ActiveType" json:"ActiveType,omitempty"`
	AppID                int32           `protobuf:"varint,8,opt,name=AppID,proto3" json:"AppID,omitempty"`
	Version              string          `protobuf:"bytes,9,opt,name=Version,proto3" json:"Version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SignUpReq) Reset()         { *m = SignUpReq{} }
func (m *SignUpReq) String() string { return proto.CompactTextString(m) }
func (*SignUpReq) ProtoMessage()    {}
func (*SignUpReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d68c7dc98409413c, []int{1}
}

func (m *SignUpReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpReq.Unmarshal(m, b)
}
func (m *SignUpReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpReq.Marshal(b, m, deterministic)
}
func (m *SignUpReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpReq.Merge(m, src)
}
func (m *SignUpReq) XXX_Size() int {
	return xxx_messageInfo_SignUpReq.Size(m)
}
func (m *SignUpReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpReq.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpReq proto.InternalMessageInfo

func (m *SignUpReq) GetDeviceID() int64 {
	if m != nil {
		return m.DeviceID
	}
	return 0
}

func (m *SignUpReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SignUpReq) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *SignUpReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SignUpReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *SignUpReq) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *SignUpReq) GetActiveType() base.ActiveType {
	if m != nil {
		return m.ActiveType
	}
	return base.ActiveType_ActiveTypeEmail
}

func (m *SignUpReq) GetAppID() int32 {
	if m != nil {
		return m.AppID
	}
	return 0
}

func (m *SignUpReq) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type SignUpRsp struct {
	UserID               int64    `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	BaseRsp              *BaseRsp `protobuf:"bytes,3,opt,name=BaseRsp,proto3" json:"BaseRsp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignUpRsp) Reset()         { *m = SignUpRsp{} }
func (m *SignUpRsp) String() string { return proto.CompactTextString(m) }
func (*SignUpRsp) ProtoMessage()    {}
func (*SignUpRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d68c7dc98409413c, []int{2}
}

func (m *SignUpRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpRsp.Unmarshal(m, b)
}
func (m *SignUpRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpRsp.Marshal(b, m, deterministic)
}
func (m *SignUpRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpRsp.Merge(m, src)
}
func (m *SignUpRsp) XXX_Size() int {
	return xxx_messageInfo_SignUpRsp.Size(m)
}
func (m *SignUpRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpRsp.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpRsp proto.InternalMessageInfo

func (m *SignUpRsp) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *SignUpRsp) GetBaseRsp() *BaseRsp {
	if m != nil {
		return m.BaseRsp
	}
	return nil
}

type WeChatSignInReq struct {
	DeviceID             int64    `protobuf:"varint,1,opt,name=DeviceID,proto3" json:"DeviceID,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=Code,proto3" json:"Code,omitempty"`
	AppID                string   `protobuf:"bytes,3,opt,name=AppID,proto3" json:"AppID,omitempty"`
	Secret               string   `protobuf:"bytes,4,opt,name=Secret,proto3" json:"Secret,omitempty"`
	ExtraInfo            string   `protobuf:"bytes,5,opt,name=ExtraInfo,proto3" json:"ExtraInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WeChatSignInReq) Reset()         { *m = WeChatSignInReq{} }
func (m *WeChatSignInReq) String() string { return proto.CompactTextString(m) }
func (*WeChatSignInReq) ProtoMessage()    {}
func (*WeChatSignInReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d68c7dc98409413c, []int{3}
}

func (m *WeChatSignInReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeChatSignInReq.Unmarshal(m, b)
}
func (m *WeChatSignInReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeChatSignInReq.Marshal(b, m, deterministic)
}
func (m *WeChatSignInReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeChatSignInReq.Merge(m, src)
}
func (m *WeChatSignInReq) XXX_Size() int {
	return xxx_messageInfo_WeChatSignInReq.Size(m)
}
func (m *WeChatSignInReq) XXX_DiscardUnknown() {
	xxx_messageInfo_WeChatSignInReq.DiscardUnknown(m)
}

var xxx_messageInfo_WeChatSignInReq proto.InternalMessageInfo

func (m *WeChatSignInReq) GetDeviceID() int64 {
	if m != nil {
		return m.DeviceID
	}
	return 0
}

func (m *WeChatSignInReq) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *WeChatSignInReq) GetAppID() string {
	if m != nil {
		return m.AppID
	}
	return ""
}

func (m *WeChatSignInReq) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *WeChatSignInReq) GetExtraInfo() string {
	if m != nil {
		return m.ExtraInfo
	}
	return ""
}

type WeChatSignInRsp struct {
	UserID               int64    `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	IsFirstSignIn        bool     `protobuf:"varint,2,opt,name=IsFirstSignIn,proto3" json:"IsFirstSignIn,omitempty"`
	BaseRsp              *BaseRsp `protobuf:"bytes,3,opt,name=BaseRsp,proto3" json:"BaseRsp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WeChatSignInRsp) Reset()         { *m = WeChatSignInRsp{} }
func (m *WeChatSignInRsp) String() string { return proto.CompactTextString(m) }
func (*WeChatSignInRsp) ProtoMessage()    {}
func (*WeChatSignInRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d68c7dc98409413c, []int{4}
}

func (m *WeChatSignInRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeChatSignInRsp.Unmarshal(m, b)
}
func (m *WeChatSignInRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeChatSignInRsp.Marshal(b, m, deterministic)
}
func (m *WeChatSignInRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeChatSignInRsp.Merge(m, src)
}
func (m *WeChatSignInRsp) XXX_Size() int {
	return xxx_messageInfo_WeChatSignInRsp.Size(m)
}
func (m *WeChatSignInRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_WeChatSignInRsp.DiscardUnknown(m)
}

var xxx_messageInfo_WeChatSignInRsp proto.InternalMessageInfo

func (m *WeChatSignInRsp) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *WeChatSignInRsp) GetIsFirstSignIn() bool {
	if m != nil {
		return m.IsFirstSignIn
	}
	return false
}

func (m *WeChatSignInRsp) GetBaseRsp() *BaseRsp {
	if m != nil {
		return m.BaseRsp
	}
	return nil
}

func init() {
	proto.RegisterEnum("platform.user.srv.passport.Code", Code_name, Code_value)
	proto.RegisterType((*BaseRsp)(nil), "platform.user.srv.passport.BaseRsp")
	proto.RegisterType((*SignUpReq)(nil), "platform.user.srv.passport.SignUpReq")
	proto.RegisterType((*SignUpRsp)(nil), "platform.user.srv.passport.SignUpRsp")
	proto.RegisterType((*WeChatSignInReq)(nil), "platform.user.srv.passport.WeChatSignInReq")
	proto.RegisterType((*WeChatSignInRsp)(nil), "platform.user.srv.passport.WeChatSignInRsp")
}

func init() { proto.RegisterFile("platform/user_srv_passport.proto", fileDescriptor_d68c7dc98409413c) }

var fileDescriptor_d68c7dc98409413c = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0xf9, 0x70, 0x92, 0xe1, 0xa3, 0xd1, 0x08, 0x55, 0x2b, 0xab, 0x87, 0xc8, 0x80, 0x14,
	0x81, 0xe4, 0xa0, 0xc0, 0x95, 0x43, 0xdb, 0xb4, 0x92, 0x0f, 0x7c, 0xb9, 0xb4, 0x48, 0x5c, 0xaa,
	0x4d, 0x32, 0x6d, 0x2d, 0x92, 0xec, 0xb2, 0x63, 0x0c, 0xfc, 0x08, 0xf8, 0x65, 0xfc, 0x23, 0x2e,
	0x68, 0xd7, 0x1f, 0x4d, 0x2b, 0x11, 0x45, 0xe2, 0x62, 0xed, 0x1b, 0xbf, 0x79, 0x3b, 0xef, 0xed,
	0xda, 0x30, 0xd0, 0x0b, 0x99, 0x5d, 0x28, 0xb3, 0x1c, 0x7d, 0x65, 0x32, 0xe7, 0x6c, 0xf2, 0x73,
	0x2d, 0x99, 0xb5, 0x32, 0x59, 0xa4, 0x8d, 0xca, 0x14, 0x06, 0x15, 0x23, 0xb2, 0x8c, 0x88, 0x4d,
	0x1e, 0x55, 0x8c, 0x60, 0x67, 0x2a, 0x99, 0x46, 0xf6, 0x51, 0x90, 0xc3, 0xf7, 0xd0, 0x39, 0x90,
	0x4c, 0x09, 0x6b, 0x7c, 0x09, 0xad, 0x99, 0x9a, 0x93, 0xf0, 0x06, 0xde, 0xf0, 0xc1, 0x78, 0x10,
	0xfd, 0x5b, 0x26, 0x3a, 0x54, 0x73, 0x4a, 0x1c, 0x1b, 0xfb, 0xd0, 0x5c, 0xf2, 0xa5, 0x68, 0x0c,
	0xbc, 0x61, 0x2f, 0xb1, 0xcb, 0xf0, 0x8f, 0x07, 0xbd, 0x93, 0xf4, 0x72, 0x75, 0xaa, 0x13, 0xfa,
	0x82, 0x01, 0x74, 0x27, 0x94, 0xa7, 0x33, 0x8a, 0x27, 0x4e, 0xb9, 0x99, 0xd4, 0x18, 0x11, 0x5a,
	0x6f, 0xe4, 0x92, 0xca, 0x66, 0xb7, 0xc6, 0x5d, 0xf0, 0x5f, 0xab, 0x69, 0xba, 0x20, 0xd1, 0x74,
	0xd5, 0x12, 0xe1, 0x43, 0x68, 0x1f, 0x2d, 0x65, 0xba, 0x10, 0x2d, 0x57, 0x2e, 0x80, 0x55, 0x7f,
	0x27, 0x99, 0xbf, 0x29, 0x33, 0x17, 0x6d, 0xf7, 0xa2, 0xc6, 0x56, 0xdd, 0xce, 0x29, 0xfc, 0x42,
	0xdd, 0xae, 0xf1, 0x39, 0xc0, 0xfe, 0x2c, 0x4b, 0x73, 0xfa, 0xf0, 0x43, 0x93, 0xe8, 0x38, 0xa7,
	0xfd, 0xc8, 0xe5, 0x71, 0x5d, 0x4f, 0xd6, 0x38, 0x76, 0xdf, 0x7d, 0xad, 0xe3, 0x89, 0xe8, 0x0e,
	0xbc, 0x61, 0x3b, 0x29, 0x00, 0x0a, 0xe8, 0x9c, 0x91, 0xe1, 0x54, 0xad, 0x44, 0xcf, 0xc9, 0x57,
	0x30, 0x9c, 0xd6, 0xe6, 0x59, 0x5b, 0x33, 0xa7, 0x4c, 0xa6, 0xb6, 0x5e, 0x22, 0x7c, 0x55, 0xa7,
	0xee, 0x5c, 0xde, 0x1d, 0x3f, 0xda, 0x94, 0x76, 0x49, 0x4d, 0xaa, 0x9e, 0xf0, 0xa7, 0x07, 0x3b,
	0x1f, 0xe9, 0xf0, 0x4a, 0x66, 0x76, 0xab, 0x78, 0xb5, 0x45, 0xce, 0x2e, 0x89, 0xc6, 0x5a, 0x12,
	0xb5, 0xaf, 0x22, 0xe6, 0xd2, 0xd7, 0x2e, 0xf8, 0x27, 0x34, 0x33, 0x94, 0x95, 0x31, 0x97, 0x08,
	0xf7, 0xa0, 0x77, 0xf4, 0x3d, 0x33, 0x32, 0x5e, 0x5d, 0xa8, 0x32, 0xe8, 0xeb, 0x42, 0xf8, 0xeb,
	0xf6, 0x3c, 0x1b, 0xac, 0x3f, 0x86, 0xfb, 0x31, 0x1f, 0xa7, 0x86, 0x4b, 0xae, 0x1b, 0xaa, 0x9b,
	0xdc, 0x2c, 0xfe, 0x67, 0x40, 0x4f, 0x83, 0xc2, 0x30, 0xfa, 0xd0, 0x78, 0xfb, 0xb9, 0x7f, 0x07,
	0x01, 0xfc, 0x63, 0x99, 0x2e, 0x68, 0xde, 0xf7, 0xc6, 0xbf, 0xbd, 0xe2, 0xce, 0xd8, 0x4e, 0x3c,
	0x03, 0xbf, 0x38, 0x2d, 0x7c, 0xb2, 0x69, 0x83, 0xfa, 0x3a, 0x07, 0xdb, 0xd0, 0x58, 0xe3, 0x15,
	0xdc, 0x5b, 0x0f, 0x04, 0x9f, 0x6d, 0x6a, 0xbb, 0x75, 0x94, 0xc1, 0xf6, 0x64, 0xd6, 0x07, 0x7b,
	0x9f, 0x82, 0x1b, 0x7f, 0x84, 0x11, 0x9b, 0x7c, 0x54, 0xb1, 0xa7, 0xbe, 0xfb, 0xca, 0x5f, 0xfc,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xf1, 0x05, 0x7e, 0x36, 0x04, 0x00, 0x00,
}