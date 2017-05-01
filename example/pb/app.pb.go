// Code generated by protoc-gen-go.
// source: example/pb/app.proto
// DO NOT EDIT!

/*
Package app is a generated protocol buffer package.

It is generated from these files:
	example/pb/app.proto

It has these top-level messages:
	EmptyMessage
	AuthRequestMessage
	AuthResponseMessage
	KeyValMessage
	SettingsMessage
	UserMessage
*/
package app

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EmptyMessage struct {
}

func (m *EmptyMessage) Reset()                    { *m = EmptyMessage{} }
func (m *EmptyMessage) String() string            { return proto.CompactTextString(m) }
func (*EmptyMessage) ProtoMessage()               {}
func (*EmptyMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type AuthRequestMessage struct {
	GrantType string `protobuf:"bytes,1,opt,name=grant_type,json=grantType" json:"grant_type,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Password  string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
}

func (m *AuthRequestMessage) Reset()                    { *m = AuthRequestMessage{} }
func (m *AuthRequestMessage) String() string            { return proto.CompactTextString(m) }
func (*AuthRequestMessage) ProtoMessage()               {}
func (*AuthRequestMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AuthRequestMessage) GetGrantType() string {
	if m != nil {
		return m.GrantType
	}
	return ""
}

func (m *AuthRequestMessage) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AuthRequestMessage) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthResponseMessage struct {
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	TokenType   string `protobuf:"bytes,2,opt,name=token_type,json=tokenType" json:"token_type,omitempty"`
	ExpiresIn   int64  `protobuf:"varint,3,opt,name=expires_in,json=expiresIn" json:"expires_in,omitempty"`
}

func (m *AuthResponseMessage) Reset()                    { *m = AuthResponseMessage{} }
func (m *AuthResponseMessage) String() string            { return proto.CompactTextString(m) }
func (*AuthResponseMessage) ProtoMessage()               {}
func (*AuthResponseMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AuthResponseMessage) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *AuthResponseMessage) GetTokenType() string {
	if m != nil {
		return m.TokenType
	}
	return ""
}

func (m *AuthResponseMessage) GetExpiresIn() int64 {
	if m != nil {
		return m.ExpiresIn
	}
	return 0
}

type KeyValMessage struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *KeyValMessage) Reset()                    { *m = KeyValMessage{} }
func (m *KeyValMessage) String() string            { return proto.CompactTextString(m) }
func (*KeyValMessage) ProtoMessage()               {}
func (*KeyValMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *KeyValMessage) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValMessage) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SettingsMessage struct {
	Initialized bool   `protobuf:"varint,1,opt,name=initialized" json:"initialized,omitempty"`
	Debug       bool   `protobuf:"varint,2,opt,name=debug" json:"debug,omitempty"`
	LogLevel    string `protobuf:"bytes,3,opt,name=logLevel" json:"logLevel,omitempty"`
}

func (m *SettingsMessage) Reset()                    { *m = SettingsMessage{} }
func (m *SettingsMessage) String() string            { return proto.CompactTextString(m) }
func (*SettingsMessage) ProtoMessage()               {}
func (*SettingsMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SettingsMessage) GetInitialized() bool {
	if m != nil {
		return m.Initialized
	}
	return false
}

func (m *SettingsMessage) GetDebug() bool {
	if m != nil {
		return m.Debug
	}
	return false
}

func (m *SettingsMessage) GetLogLevel() string {
	if m != nil {
		return m.LogLevel
	}
	return ""
}

type UserMessage struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password" json:"password,omitempty"`
	Role     string `protobuf:"bytes,5,opt,name=role" json:"role,omitempty"`
	Phone    string `protobuf:"bytes,6,opt,name=phone" json:"phone,omitempty"`
}

func (m *UserMessage) Reset()                    { *m = UserMessage{} }
func (m *UserMessage) String() string            { return proto.CompactTextString(m) }
func (*UserMessage) ProtoMessage()               {}
func (*UserMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UserMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserMessage) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserMessage) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserMessage) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *UserMessage) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func init() {
	proto.RegisterType((*EmptyMessage)(nil), "app.EmptyMessage")
	proto.RegisterType((*AuthRequestMessage)(nil), "app.AuthRequestMessage")
	proto.RegisterType((*AuthResponseMessage)(nil), "app.AuthResponseMessage")
	proto.RegisterType((*KeyValMessage)(nil), "app.KeyValMessage")
	proto.RegisterType((*SettingsMessage)(nil), "app.SettingsMessage")
	proto.RegisterType((*UserMessage)(nil), "app.UserMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Auth service

type AuthClient interface {
	Auth(ctx context.Context, in *AuthRequestMessage, opts ...grpc.CallOption) (*AuthResponseMessage, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Auth(ctx context.Context, in *AuthRequestMessage, opts ...grpc.CallOption) (*AuthResponseMessage, error) {
	out := new(AuthResponseMessage)
	err := grpc.Invoke(ctx, "/app.Auth/Auth", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthServer interface {
	Auth(context.Context, *AuthRequestMessage) (*AuthResponseMessage, error)
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.Auth/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Auth(ctx, req.(*AuthRequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "app.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _Auth_Auth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/pb/app.proto",
}

// Client API for KeyVal service

type KeyValClient interface {
	KeyValCreate(ctx context.Context, in *KeyValMessage, opts ...grpc.CallOption) (*KeyValMessage, error)
	KeyValRead(ctx context.Context, in *KeyValMessage, opts ...grpc.CallOption) (*KeyValMessage, error)
	KeyValUpdate(ctx context.Context, in *KeyValMessage, opts ...grpc.CallOption) (*KeyValMessage, error)
	KeyValDelete(ctx context.Context, in *KeyValMessage, opts ...grpc.CallOption) (*KeyValMessage, error)
}

type keyValClient struct {
	cc *grpc.ClientConn
}

func NewKeyValClient(cc *grpc.ClientConn) KeyValClient {
	return &keyValClient{cc}
}

func (c *keyValClient) KeyValCreate(ctx context.Context, in *KeyValMessage, opts ...grpc.CallOption) (*KeyValMessage, error) {
	out := new(KeyValMessage)
	err := grpc.Invoke(ctx, "/app.KeyVal/KeyValCreate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValClient) KeyValRead(ctx context.Context, in *KeyValMessage, opts ...grpc.CallOption) (*KeyValMessage, error) {
	out := new(KeyValMessage)
	err := grpc.Invoke(ctx, "/app.KeyVal/KeyValRead", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValClient) KeyValUpdate(ctx context.Context, in *KeyValMessage, opts ...grpc.CallOption) (*KeyValMessage, error) {
	out := new(KeyValMessage)
	err := grpc.Invoke(ctx, "/app.KeyVal/KeyValUpdate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValClient) KeyValDelete(ctx context.Context, in *KeyValMessage, opts ...grpc.CallOption) (*KeyValMessage, error) {
	out := new(KeyValMessage)
	err := grpc.Invoke(ctx, "/app.KeyVal/KeyValDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KeyVal service

type KeyValServer interface {
	KeyValCreate(context.Context, *KeyValMessage) (*KeyValMessage, error)
	KeyValRead(context.Context, *KeyValMessage) (*KeyValMessage, error)
	KeyValUpdate(context.Context, *KeyValMessage) (*KeyValMessage, error)
	KeyValDelete(context.Context, *KeyValMessage) (*KeyValMessage, error)
}

func RegisterKeyValServer(s *grpc.Server, srv KeyValServer) {
	s.RegisterService(&_KeyVal_serviceDesc, srv)
}

func _KeyVal_KeyValCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValServer).KeyValCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.KeyVal/KeyValCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValServer).KeyValCreate(ctx, req.(*KeyValMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyVal_KeyValRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValServer).KeyValRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.KeyVal/KeyValRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValServer).KeyValRead(ctx, req.(*KeyValMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyVal_KeyValUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValServer).KeyValUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.KeyVal/KeyValUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValServer).KeyValUpdate(ctx, req.(*KeyValMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyVal_KeyValDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValServer).KeyValDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.KeyVal/KeyValDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValServer).KeyValDelete(ctx, req.(*KeyValMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeyVal_serviceDesc = grpc.ServiceDesc{
	ServiceName: "app.KeyVal",
	HandlerType: (*KeyValServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "KeyValCreate",
			Handler:    _KeyVal_KeyValCreate_Handler,
		},
		{
			MethodName: "KeyValRead",
			Handler:    _KeyVal_KeyValRead_Handler,
		},
		{
			MethodName: "KeyValUpdate",
			Handler:    _KeyVal_KeyValUpdate_Handler,
		},
		{
			MethodName: "KeyValDelete",
			Handler:    _KeyVal_KeyValDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/pb/app.proto",
}

func init() { proto.RegisterFile("example/pb/app.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0x95, 0x3f, 0x8d, 0x92, 0x49, 0x5a, 0x2a, 0xb7, 0x88, 0x55, 0xa0, 0x52, 0xf1, 0x09,
	0xe5, 0xd0, 0x15, 0xe5, 0x80, 0xd4, 0x1b, 0x02, 0x0e, 0x08, 0xa8, 0xd0, 0xf6, 0xcf, 0x35, 0x72,
	0xba, 0xa3, 0x8d, 0x15, 0xc7, 0x36, 0x6b, 0x6f, 0xe8, 0x82, 0xb8, 0x70, 0xe7, 0xc4, 0xa3, 0xc1,
	0x23, 0xf0, 0x20, 0xc8, 0xf6, 0x1a, 0x1a, 0xa5, 0x07, 0xe0, 0x14, 0xcf, 0x37, 0xce, 0xef, 0xfb,
	0x34, 0xe3, 0x85, 0x7d, 0xbc, 0x66, 0x4b, 0x2d, 0x30, 0xd5, 0xb3, 0x94, 0x69, 0x7d, 0xa4, 0x4b,
	0x65, 0x15, 0xe9, 0x30, 0xad, 0xc7, 0x0f, 0x0a, 0xa5, 0x0a, 0x81, 0x29, 0xd3, 0x3c, 0x65, 0x52,
	0x2a, 0xcb, 0x2c, 0x57, 0xd2, 0x84, 0x2b, 0x74, 0x07, 0x46, 0x2f, 0x97, 0xda, 0xd6, 0x6f, 0xd1,
	0x18, 0x56, 0x20, 0x5d, 0x00, 0x79, 0x56, 0xd9, 0x79, 0x86, 0xef, 0x2b, 0x34, 0xb6, 0x51, 0xc9,
	0x01, 0x40, 0x51, 0x32, 0x69, 0xa7, 0xb6, 0xd6, 0x98, 0xb4, 0x0e, 0x5b, 0x8f, 0x06, 0xd9, 0xc0,
	0x2b, 0xe7, 0xb5, 0x46, 0x32, 0x86, 0x7e, 0x65, 0xb0, 0x94, 0x6c, 0x89, 0x49, 0xdb, 0x37, 0x7f,
	0xd7, 0xae, 0xa7, 0x99, 0x31, 0x1f, 0x54, 0x99, 0x27, 0x9d, 0xd0, 0x8b, 0x35, 0x5d, 0xc1, 0x5e,
	0x30, 0x33, 0x5a, 0x49, 0x83, 0xd1, 0xed, 0x21, 0x8c, 0xd8, 0xd5, 0x15, 0x1a, 0x33, 0xb5, 0x6a,
	0x81, 0xb2, 0xf1, 0x1b, 0x06, 0xed, 0xdc, 0x49, 0x2e, 0x90, 0xef, 0x85, 0x40, 0xc1, 0x73, 0xe0,
	0x15, 0x1f, 0xe8, 0x00, 0x00, 0xaf, 0x35, 0x2f, 0xd1, 0x4c, 0xb9, 0xf4, 0xb6, 0x9d, 0x6c, 0xd0,
	0x28, 0xaf, 0x24, 0x7d, 0x0a, 0xdb, 0xaf, 0xb1, 0xbe, 0x64, 0x22, 0x3a, 0xee, 0x42, 0x67, 0x81,
	0x75, 0x63, 0xe4, 0x8e, 0x64, 0x1f, 0xb6, 0x56, 0x4c, 0x54, 0x91, 0x1d, 0x0a, 0x8a, 0x70, 0xe7,
	0x0c, 0xad, 0xe5, 0xb2, 0x30, 0xf1, 0xaf, 0x87, 0x30, 0xe4, 0x92, 0x5b, 0xce, 0x04, 0xff, 0x88,
	0xb9, 0x47, 0xf4, 0xb3, 0x9b, 0x92, 0x43, 0xe5, 0x38, 0xab, 0x0a, 0x8f, 0xea, 0x67, 0xa1, 0x70,
	0x73, 0x11, 0xaa, 0x78, 0x83, 0x2b, 0x14, 0x71, 0x2e, 0xb1, 0xa6, 0x5f, 0x5b, 0x30, 0xbc, 0x30,
	0x58, 0x46, 0x8f, 0x1d, 0x68, 0xf3, 0xbc, 0x49, 0xd7, 0xe6, 0x9e, 0x88, 0x4b, 0xc6, 0x45, 0x0c,
	0xe7, 0x0b, 0x42, 0xa0, 0xeb, 0x37, 0x10, 0x68, 0xdd, 0x8d, 0xe9, 0x77, 0xd7, 0xa7, 0xef, 0xee,
	0x97, 0x4a, 0x60, 0xb2, 0x15, 0xee, 0xbb, 0xb3, 0x23, 0xeb, 0xb9, 0x92, 0x98, 0xf4, 0x02, 0xd9,
	0x17, 0xc7, 0x97, 0xd0, 0x75, 0x7b, 0x22, 0xa7, 0xcd, 0xef, 0xbd, 0x23, 0xf7, 0xc6, 0x36, 0xdf,
	0xc9, 0x38, 0xb9, 0xd1, 0x58, 0xdb, 0x29, 0xdd, 0xfb, 0xf2, 0xfd, 0xe7, 0xb7, 0xf6, 0x36, 0xed,
	0xa7, 0xab, 0xc7, 0x29, 0xab, 0xec, 0xfc, 0xa4, 0x35, 0x39, 0xfe, 0xd1, 0x86, 0x5e, 0x58, 0x04,
	0x39, 0x83, 0x51, 0x38, 0x3d, 0x2f, 0x91, 0x59, 0x24, 0xc4, 0x93, 0xd6, 0xb6, 0x34, 0xbe, 0x45,
	0xa3, 0xf7, 0x3d, 0xf7, 0xee, 0x78, 0xd7, 0x71, 0x17, 0x58, 0xaf, 0x98, 0x48, 0x3f, 0x2d, 0xb0,
	0xfe, 0x7c, 0xd2, 0x9a, 0x90, 0x53, 0x80, 0x70, 0x3b, 0x43, 0x96, 0xff, 0x35, 0x32, 0xf1, 0x48,
	0x42, 0x36, 0x90, 0x7f, 0x42, 0x5e, 0xe8, 0xfc, 0x3f, 0x42, 0xd2, 0x5b, 0x43, 0xbe, 0x8b, 0xd0,
	0x17, 0x28, 0xf0, 0x1f, 0xa0, 0x4d, 0xcc, 0xc9, 0x06, 0x74, 0xd6, 0xf3, 0x9f, 0xf6, 0x93, 0x5f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x1a, 0x3e, 0x86, 0x15, 0x04, 0x00, 0x00,
}
