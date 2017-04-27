// Code generated by protoc-gen-go.
// source: user.proto
// DO NOT EDIT!

/*
Package userpb is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	User
	SignUpRequest
	Token
	SignInRequest
	UpdateProfileRequest
	GetUsersRequest
	GetUsersResponse
	ChangePasswordRequest
*/
package userpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

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

type User struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username" json:"username,omitempty"`
	Photo    []byte `protobuf:"bytes,4,opt,name=photo,proto3" json:"photo,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetPhoto() []byte {
	if m != nil {
		return m.Photo
	}
	return nil
}

type SignUpRequest struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	Photo    []byte `protobuf:"bytes,7,opt,name=photo,proto3" json:"photo,omitempty"`
}

func (m *SignUpRequest) Reset()                    { *m = SignUpRequest{} }
func (m *SignUpRequest) String() string            { return proto.CompactTextString(m) }
func (*SignUpRequest) ProtoMessage()               {}
func (*SignUpRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SignUpRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SignUpRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SignUpRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *SignUpRequest) GetPhoto() []byte {
	if m != nil {
		return m.Photo
	}
	return nil
}

type Token struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type SignInRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *SignInRequest) Reset()                    { *m = SignInRequest{} }
func (m *SignInRequest) String() string            { return proto.CompactTextString(m) }
func (*SignInRequest) ProtoMessage()               {}
func (*SignInRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SignInRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SignInRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UpdateProfileRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Photo    []byte `protobuf:"bytes,3,opt,name=photo,proto3" json:"photo,omitempty"`
}

func (m *UpdateProfileRequest) Reset()                    { *m = UpdateProfileRequest{} }
func (m *UpdateProfileRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateProfileRequest) ProtoMessage()               {}
func (*UpdateProfileRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateProfileRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UpdateProfileRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateProfileRequest) GetPhoto() []byte {
	if m != nil {
		return m.Photo
	}
	return nil
}

type GetUsersRequest struct {
	Ids      []string `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
	Name     string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Username string   `protobuf:"bytes,3,opt,name=username" json:"username,omitempty"`
}

func (m *GetUsersRequest) Reset()                    { *m = GetUsersRequest{} }
func (m *GetUsersRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUsersRequest) ProtoMessage()               {}
func (*GetUsersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetUsersRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *GetUsersRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetUsersRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type GetUsersResponse struct {
	Users []*User `protobuf:"bytes,1,rep,name=users" json:"users,omitempty"`
}

func (m *GetUsersResponse) Reset()                    { *m = GetUsersResponse{} }
func (m *GetUsersResponse) String() string            { return proto.CompactTextString(m) }
func (*GetUsersResponse) ProtoMessage()               {}
func (*GetUsersResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type ChangePasswordRequest struct {
	OldPassword string `protobuf:"bytes,1,opt,name=old_password,json=oldPassword" json:"old_password,omitempty"`
	NewPassword string `protobuf:"bytes,2,opt,name=new_password,json=newPassword" json:"new_password,omitempty"`
}

func (m *ChangePasswordRequest) Reset()                    { *m = ChangePasswordRequest{} }
func (m *ChangePasswordRequest) String() string            { return proto.CompactTextString(m) }
func (*ChangePasswordRequest) ProtoMessage()               {}
func (*ChangePasswordRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ChangePasswordRequest) GetOldPassword() string {
	if m != nil {
		return m.OldPassword
	}
	return ""
}

func (m *ChangePasswordRequest) GetNewPassword() string {
	if m != nil {
		return m.NewPassword
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "userpb.User")
	proto.RegisterType((*SignUpRequest)(nil), "userpb.SignUpRequest")
	proto.RegisterType((*Token)(nil), "userpb.Token")
	proto.RegisterType((*SignInRequest)(nil), "userpb.SignInRequest")
	proto.RegisterType((*UpdateProfileRequest)(nil), "userpb.UpdateProfileRequest")
	proto.RegisterType((*GetUsersRequest)(nil), "userpb.GetUsersRequest")
	proto.RegisterType((*GetUsersResponse)(nil), "userpb.GetUsersResponse")
	proto.RegisterType((*ChangePasswordRequest)(nil), "userpb.ChangePasswordRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AuthService service

type AuthServiceClient interface {
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*Token, error)
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*Token, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := grpc.Invoke(ctx, "/userpb.AuthService/SignUp", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := grpc.Invoke(ctx, "/userpb.AuthService/SignIn", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthService service

type AuthServiceServer interface {
	SignUp(context.Context, *SignUpRequest) (*Token, error)
	SignIn(context.Context, *SignInRequest) (*Token, error)
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userpb.AuthService/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userpb.AuthService/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "userpb.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _AuthService_SignUp_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _AuthService_SignIn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

// Client API for UserService service

type UserServiceClient interface {
	UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	// Get Users by id or find users by username or name
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/userpb.UserService/UpdateProfile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	out := new(GetUsersResponse)
	err := grpc.Invoke(ctx, "/userpb.UserService/GetUsers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/userpb.UserService/ChangePassword", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceServer interface {
	UpdateProfile(context.Context, *UpdateProfileRequest) (*google_protobuf.Empty, error)
	// Get Users by id or find users by username or name
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*google_protobuf.Empty, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userpb.UserService/UpdateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateProfile(ctx, req.(*UpdateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userpb.UserService/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUsers(ctx, req.(*GetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userpb.UserService/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "userpb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateProfile",
			Handler:    _UserService_UpdateProfile_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _UserService_GetUsers_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _UserService_ChangePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xdf, 0x6f, 0xd3, 0x30,
	0x10, 0x6e, 0x92, 0xb6, 0x74, 0x97, 0x76, 0x4c, 0xd6, 0x06, 0x51, 0x60, 0x52, 0xf1, 0x53, 0x9f,
	0x32, 0x54, 0x24, 0x5e, 0x11, 0x42, 0xa8, 0xea, 0xdb, 0x94, 0xd2, 0xb7, 0xa1, 0x29, 0x25, 0xb7,
	0x34, 0x22, 0xb3, 0xdd, 0xd8, 0xa1, 0xe2, 0xdf, 0xe5, 0x2f, 0x41, 0x8e, 0xe7, 0xac, 0xe9, 0x5a,
	0xc4, 0xde, 0xee, 0x97, 0xbf, 0xef, 0xee, 0x3e, 0x1f, 0x40, 0x25, 0xb1, 0x8c, 0x44, 0xc9, 0x15,
	0x27, 0x7d, 0x6d, 0x8b, 0x55, 0xf8, 0x26, 0xe3, 0x3c, 0x2b, 0xf0, 0xaa, 0x8e, 0xae, 0xaa, 0xbb,
	0x2b, 0xbc, 0x17, 0xea, 0xb7, 0x29, 0xa2, 0x37, 0xd0, 0x5d, 0x4a, 0x2c, 0xc9, 0x29, 0xb8, 0x79,
	0x1a, 0x38, 0x63, 0x67, 0x72, 0x12, 0xbb, 0x79, 0x4a, 0x08, 0x74, 0x59, 0x72, 0x8f, 0x81, 0x5b,
	0x47, 0x6a, 0x9b, 0x84, 0x30, 0xd0, 0x90, 0x75, 0xdc, 0xab, 0xe3, 0x8d, 0x4f, 0xce, 0xa1, 0x27,
	0xd6, 0x5c, 0xf1, 0xa0, 0x3b, 0x76, 0x26, 0xc3, 0xd8, 0x38, 0x74, 0x03, 0xa3, 0x45, 0x9e, 0xb1,
	0xa5, 0x88, 0x71, 0x53, 0xa1, 0x54, 0x0d, 0xac, 0x73, 0x04, 0xd6, 0xdd, 0x83, 0x0d, 0x61, 0x20,
	0x12, 0x29, 0xb7, 0xbc, 0x4c, 0x2d, 0xa5, 0xf5, 0x1f, 0x29, 0x5f, 0xec, 0x52, 0x5e, 0x42, 0xef,
	0x1b, 0xff, 0x89, 0x4c, 0xa7, 0x95, 0x36, 0x1e, 0xb8, 0x8c, 0x43, 0x67, 0xa6, 0xa3, 0x39, 0xb3,
	0x1d, 0xed, 0xb2, 0x3b, 0xff, 0x60, 0x77, 0xdb, 0xec, 0xf4, 0x06, 0xce, 0x97, 0x22, 0x4d, 0x14,
	0x5e, 0x97, 0xfc, 0x2e, 0x2f, 0xf0, 0x7f, 0xf0, 0x0e, 0x2d, 0xb5, 0x99, 0xc2, 0xdb, 0x9d, 0x62,
	0x01, 0x2f, 0x67, 0xa8, 0xb4, 0x32, 0xd2, 0x02, 0x9f, 0x81, 0x97, 0xa7, 0x32, 0x70, 0xc6, 0xde,
	0xe4, 0x24, 0xd6, 0xe6, 0x73, 0x35, 0xa2, 0x1f, 0xe1, 0xec, 0x11, 0x54, 0x0a, 0xce, 0x24, 0x12,
	0x0a, 0x3d, 0x9d, 0x37, 0xb8, 0xfe, 0x74, 0x18, 0x99, 0x4f, 0x13, 0xe9, 0xaa, 0xd8, 0xa4, 0xe8,
	0x77, 0xb8, 0xf8, 0xb2, 0x4e, 0x58, 0x86, 0xd7, 0x0f, 0xc3, 0xdb, 0x96, 0xde, 0xc1, 0x90, 0x17,
	0xe9, 0x6d, 0xb3, 0x23, 0x33, 0xaf, 0xcf, 0x8b, 0xd4, 0x56, 0xea, 0x12, 0x86, 0xdb, 0xdb, 0xbd,
	0x35, 0xfa, 0x0c, 0xb7, 0xb6, 0x64, 0xba, 0x01, 0xff, 0x73, 0xa5, 0xd6, 0x0b, 0x2c, 0x7f, 0xe5,
	0x3f, 0x90, 0xbc, 0x87, 0xbe, 0xf9, 0x33, 0xe4, 0xc2, 0x36, 0xd3, 0xfa, 0x43, 0xe1, 0xc8, 0x86,
	0x6b, 0x9d, 0x69, 0xc7, 0xbe, 0x98, 0xb3, 0xf6, 0x8b, 0x46, 0xe3, 0x27, 0x2f, 0xa6, 0x7f, 0x1c,
	0xf0, 0xf5, 0x84, 0x96, 0x73, 0x06, 0xa3, 0x96, 0x98, 0xe4, 0x6d, 0xb3, 0x87, 0x03, 0x1a, 0x87,
	0xaf, 0x22, 0x73, 0x52, 0x91, 0x3d, 0xa9, 0xe8, 0xab, 0x3e, 0x29, 0xda, 0x21, 0x9f, 0x60, 0x60,
	0x57, 0x4c, 0x5e, 0x5b, 0x8c, 0x3d, 0x25, 0xc3, 0xe0, 0x69, 0xc2, 0xa8, 0x41, 0x3b, 0x64, 0x0e,
	0xa7, 0xed, 0x5d, 0x93, 0x4b, 0x5b, 0x7d, 0x50, 0x83, 0xe3, 0xbd, 0xac, 0xfa, 0x75, 0xe4, 0xc3,
	0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x94, 0x97, 0x4c, 0x14, 0x04, 0x00, 0x00,
}
