// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: protobuf/auth.proto

package authpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	SendOTPEmail(ctx context.Context, in *OTPEmailRequest, opts ...grpc.CallOption) (*OTPEmailResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) SendOTPEmail(ctx context.Context, in *OTPEmailRequest, opts ...grpc.CallOption) (*OTPEmailResponse, error) {
	out := new(OTPEmailResponse)
	err := c.cc.Invoke(ctx, "/authpb.AuthService/SendOTPEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/authpb.AuthService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error) {
	out := new(RefreshResponse)
	err := c.cc.Invoke(ctx, "/authpb.AuthService/Refresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	SendOTPEmail(context.Context, *OTPEmailRequest) (*OTPEmailResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) SendOTPEmail(context.Context, *OTPEmailRequest) (*OTPEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOTPEmail not implemented")
}
func (UnimplementedAuthServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServiceServer) Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_SendOTPEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OTPEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SendOTPEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authpb.AuthService/SendOTPEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SendOTPEmail(ctx, req.(*OTPEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authpb.AuthService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authpb.AuthService/Refresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Refresh(ctx, req.(*RefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authpb.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendOTPEmail",
			Handler:    _AuthService_SendOTPEmail_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _AuthService_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/auth.proto",
}
