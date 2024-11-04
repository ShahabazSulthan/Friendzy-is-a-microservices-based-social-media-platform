// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0--rc2
// source: pkg/Auth_Service/pb/auth.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AuthService_UserSignUp_FullMethodName                    = "/auth_proto.AuthService/UserSignUp"
	AuthService_UserOTPVerification_FullMethodName           = "/auth_proto.AuthService/UserOTPVerification"
	AuthService_UserLogin_FullMethodName                     = "/auth_proto.AuthService/UserLogin"
	AuthService_ForgotPasswordRequest_FullMethodName         = "/auth_proto.AuthService/ForgotPasswordRequest"
	AuthService_ResetPassword_FullMethodName                 = "/auth_proto.AuthService/ResetPassword"
	AuthService_VerifyAccessToken_FullMethodName             = "/auth_proto.AuthService/VerifyAccessToken"
	AuthService_AccessRegenerator_FullMethodName             = "/auth_proto.AuthService/AccessRegenerator"
	AuthService_GetUserProfile_FullMethodName                = "/auth_proto.AuthService/GetUserProfile"
	AuthService_EditUserProfile_FullMethodName               = "/auth_proto.AuthService/EditUserProfile"
	AuthService_GetFollowersDetails_FullMethodName           = "/auth_proto.AuthService/GetFollowersDetails"
	AuthService_GetFollowingsDetails_FullMethodName          = "/auth_proto.AuthService/GetFollowingsDetails"
	AuthService_SearchUser_FullMethodName                    = "/auth_proto.AuthService/SearchUser"
	AuthService_SetUserProfileImage_FullMethodName           = "/auth_proto.AuthService/SetUserProfileImage"
	AuthService_GetUserDetailsLiteForPostView_FullMethodName = "/auth_proto.AuthService/GetUserDetailsLiteForPostView"
	AuthService_CheckUserExist_FullMethodName                = "/auth_proto.AuthService/CheckUserExist"
	AuthService_AdminLogin_FullMethodName                    = "/auth_proto.AuthService/AdminLogin"
	AuthService_GetAllUsers_FullMethodName                   = "/auth_proto.AuthService/GetAllUsers"
	AuthService_BlockUser_FullMethodName                     = "/auth_proto.AuthService/BlockUser"
	AuthService_UnblockUser_FullMethodName                   = "/auth_proto.AuthService/UnblockUser"
	AuthService_VerifyAdminToken_FullMethodName              = "/auth_proto.AuthService/VerifyAdminToken"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	UserSignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error)
	UserOTPVerification(ctx context.Context, in *RequestOtpVefification, opts ...grpc.CallOption) (*ResponseOtpVerification, error)
	UserLogin(ctx context.Context, in *RequestUserLogin, opts ...grpc.CallOption) (*ResponseUserLogin, error)
	ForgotPasswordRequest(ctx context.Context, in *RequestForgotPass, opts ...grpc.CallOption) (*ResponseForgotPass, error)
	ResetPassword(ctx context.Context, in *RequestResetPass, opts ...grpc.CallOption) (*ResponseErrorMessage, error)
	VerifyAccessToken(ctx context.Context, in *RequestVerifyAccess, opts ...grpc.CallOption) (*ResponseVerifyAccess, error)
	AccessRegenerator(ctx context.Context, in *RequestAccessGenerator, opts ...grpc.CallOption) (*ResponseAccessGenerator, error)
	GetUserProfile(ctx context.Context, in *RequestGetUserProfile, opts ...grpc.CallOption) (*ResponseUserProfile, error)
	EditUserProfile(ctx context.Context, in *RequestEditUserProfile, opts ...grpc.CallOption) (*ResponseErrorMessage, error)
	GetFollowersDetails(ctx context.Context, in *RequestUserId, opts ...grpc.CallOption) (*ResponseGetUsersDetails, error)
	GetFollowingsDetails(ctx context.Context, in *RequestUserId, opts ...grpc.CallOption) (*ResponseGetUsersDetails, error)
	SearchUser(ctx context.Context, in *RequestUserSearch, opts ...grpc.CallOption) (*ResponseUserSearch, error)
	SetUserProfileImage(ctx context.Context, in *RequestSetProfileImg, opts ...grpc.CallOption) (*ResponseErrorMessage, error)
	GetUserDetailsLiteForPostView(ctx context.Context, in *RequestUserId, opts ...grpc.CallOption) (*ResponseUserDetailsLite, error)
	CheckUserExist(ctx context.Context, in *RequestUserId, opts ...grpc.CallOption) (*ResponseBool, error)
	AdminLogin(ctx context.Context, in *AdminLoginRequest, opts ...grpc.CallOption) (*AdminLoginResponse, error)
	GetAllUsers(ctx context.Context, in *GetAllUsersRequest, opts ...grpc.CallOption) (*GetAllUsersResponse, error)
	BlockUser(ctx context.Context, in *BlockUserRequest, opts ...grpc.CallOption) (*ResponseErrorMessage, error)
	UnblockUser(ctx context.Context, in *UnblockUserRequest, opts ...grpc.CallOption) (*ResponseErrorMessage, error)
	VerifyAdminToken(ctx context.Context, in *RequestVerifyAdmin, opts ...grpc.CallOption) (*ResponseVerifyAdmin, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) UserSignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignUpResponse)
	err := c.cc.Invoke(ctx, AuthService_UserSignUp_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UserOTPVerification(ctx context.Context, in *RequestOtpVefification, opts ...grpc.CallOption) (*ResponseOtpVerification, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseOtpVerification)
	err := c.cc.Invoke(ctx, AuthService_UserOTPVerification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UserLogin(ctx context.Context, in *RequestUserLogin, opts ...grpc.CallOption) (*ResponseUserLogin, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseUserLogin)
	err := c.cc.Invoke(ctx, AuthService_UserLogin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ForgotPasswordRequest(ctx context.Context, in *RequestForgotPass, opts ...grpc.CallOption) (*ResponseForgotPass, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseForgotPass)
	err := c.cc.Invoke(ctx, AuthService_ForgotPasswordRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ResetPassword(ctx context.Context, in *RequestResetPass, opts ...grpc.CallOption) (*ResponseErrorMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseErrorMessage)
	err := c.cc.Invoke(ctx, AuthService_ResetPassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) VerifyAccessToken(ctx context.Context, in *RequestVerifyAccess, opts ...grpc.CallOption) (*ResponseVerifyAccess, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseVerifyAccess)
	err := c.cc.Invoke(ctx, AuthService_VerifyAccessToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AccessRegenerator(ctx context.Context, in *RequestAccessGenerator, opts ...grpc.CallOption) (*ResponseAccessGenerator, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseAccessGenerator)
	err := c.cc.Invoke(ctx, AuthService_AccessRegenerator_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetUserProfile(ctx context.Context, in *RequestGetUserProfile, opts ...grpc.CallOption) (*ResponseUserProfile, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseUserProfile)
	err := c.cc.Invoke(ctx, AuthService_GetUserProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) EditUserProfile(ctx context.Context, in *RequestEditUserProfile, opts ...grpc.CallOption) (*ResponseErrorMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseErrorMessage)
	err := c.cc.Invoke(ctx, AuthService_EditUserProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetFollowersDetails(ctx context.Context, in *RequestUserId, opts ...grpc.CallOption) (*ResponseGetUsersDetails, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseGetUsersDetails)
	err := c.cc.Invoke(ctx, AuthService_GetFollowersDetails_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetFollowingsDetails(ctx context.Context, in *RequestUserId, opts ...grpc.CallOption) (*ResponseGetUsersDetails, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseGetUsersDetails)
	err := c.cc.Invoke(ctx, AuthService_GetFollowingsDetails_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) SearchUser(ctx context.Context, in *RequestUserSearch, opts ...grpc.CallOption) (*ResponseUserSearch, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseUserSearch)
	err := c.cc.Invoke(ctx, AuthService_SearchUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) SetUserProfileImage(ctx context.Context, in *RequestSetProfileImg, opts ...grpc.CallOption) (*ResponseErrorMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseErrorMessage)
	err := c.cc.Invoke(ctx, AuthService_SetUserProfileImage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetUserDetailsLiteForPostView(ctx context.Context, in *RequestUserId, opts ...grpc.CallOption) (*ResponseUserDetailsLite, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseUserDetailsLite)
	err := c.cc.Invoke(ctx, AuthService_GetUserDetailsLiteForPostView_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CheckUserExist(ctx context.Context, in *RequestUserId, opts ...grpc.CallOption) (*ResponseBool, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseBool)
	err := c.cc.Invoke(ctx, AuthService_CheckUserExist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AdminLogin(ctx context.Context, in *AdminLoginRequest, opts ...grpc.CallOption) (*AdminLoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdminLoginResponse)
	err := c.cc.Invoke(ctx, AuthService_AdminLogin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetAllUsers(ctx context.Context, in *GetAllUsersRequest, opts ...grpc.CallOption) (*GetAllUsersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllUsersResponse)
	err := c.cc.Invoke(ctx, AuthService_GetAllUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) BlockUser(ctx context.Context, in *BlockUserRequest, opts ...grpc.CallOption) (*ResponseErrorMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseErrorMessage)
	err := c.cc.Invoke(ctx, AuthService_BlockUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UnblockUser(ctx context.Context, in *UnblockUserRequest, opts ...grpc.CallOption) (*ResponseErrorMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseErrorMessage)
	err := c.cc.Invoke(ctx, AuthService_UnblockUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) VerifyAdminToken(ctx context.Context, in *RequestVerifyAdmin, opts ...grpc.CallOption) (*ResponseVerifyAdmin, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseVerifyAdmin)
	err := c.cc.Invoke(ctx, AuthService_VerifyAdminToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility.
type AuthServiceServer interface {
	UserSignUp(context.Context, *SignUpRequest) (*SignUpResponse, error)
	UserOTPVerification(context.Context, *RequestOtpVefification) (*ResponseOtpVerification, error)
	UserLogin(context.Context, *RequestUserLogin) (*ResponseUserLogin, error)
	ForgotPasswordRequest(context.Context, *RequestForgotPass) (*ResponseForgotPass, error)
	ResetPassword(context.Context, *RequestResetPass) (*ResponseErrorMessage, error)
	VerifyAccessToken(context.Context, *RequestVerifyAccess) (*ResponseVerifyAccess, error)
	AccessRegenerator(context.Context, *RequestAccessGenerator) (*ResponseAccessGenerator, error)
	GetUserProfile(context.Context, *RequestGetUserProfile) (*ResponseUserProfile, error)
	EditUserProfile(context.Context, *RequestEditUserProfile) (*ResponseErrorMessage, error)
	GetFollowersDetails(context.Context, *RequestUserId) (*ResponseGetUsersDetails, error)
	GetFollowingsDetails(context.Context, *RequestUserId) (*ResponseGetUsersDetails, error)
	SearchUser(context.Context, *RequestUserSearch) (*ResponseUserSearch, error)
	SetUserProfileImage(context.Context, *RequestSetProfileImg) (*ResponseErrorMessage, error)
	GetUserDetailsLiteForPostView(context.Context, *RequestUserId) (*ResponseUserDetailsLite, error)
	CheckUserExist(context.Context, *RequestUserId) (*ResponseBool, error)
	AdminLogin(context.Context, *AdminLoginRequest) (*AdminLoginResponse, error)
	GetAllUsers(context.Context, *GetAllUsersRequest) (*GetAllUsersResponse, error)
	BlockUser(context.Context, *BlockUserRequest) (*ResponseErrorMessage, error)
	UnblockUser(context.Context, *UnblockUserRequest) (*ResponseErrorMessage, error)
	VerifyAdminToken(context.Context, *RequestVerifyAdmin) (*ResponseVerifyAdmin, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthServiceServer struct{}

func (UnimplementedAuthServiceServer) UserSignUp(context.Context, *SignUpRequest) (*SignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSignUp not implemented")
}
func (UnimplementedAuthServiceServer) UserOTPVerification(context.Context, *RequestOtpVefification) (*ResponseOtpVerification, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserOTPVerification not implemented")
}
func (UnimplementedAuthServiceServer) UserLogin(context.Context, *RequestUserLogin) (*ResponseUserLogin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedAuthServiceServer) ForgotPasswordRequest(context.Context, *RequestForgotPass) (*ResponseForgotPass, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForgotPasswordRequest not implemented")
}
func (UnimplementedAuthServiceServer) ResetPassword(context.Context, *RequestResetPass) (*ResponseErrorMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedAuthServiceServer) VerifyAccessToken(context.Context, *RequestVerifyAccess) (*ResponseVerifyAccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyAccessToken not implemented")
}
func (UnimplementedAuthServiceServer) AccessRegenerator(context.Context, *RequestAccessGenerator) (*ResponseAccessGenerator, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccessRegenerator not implemented")
}
func (UnimplementedAuthServiceServer) GetUserProfile(context.Context, *RequestGetUserProfile) (*ResponseUserProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}
func (UnimplementedAuthServiceServer) EditUserProfile(context.Context, *RequestEditUserProfile) (*ResponseErrorMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditUserProfile not implemented")
}
func (UnimplementedAuthServiceServer) GetFollowersDetails(context.Context, *RequestUserId) (*ResponseGetUsersDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowersDetails not implemented")
}
func (UnimplementedAuthServiceServer) GetFollowingsDetails(context.Context, *RequestUserId) (*ResponseGetUsersDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowingsDetails not implemented")
}
func (UnimplementedAuthServiceServer) SearchUser(context.Context, *RequestUserSearch) (*ResponseUserSearch, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUser not implemented")
}
func (UnimplementedAuthServiceServer) SetUserProfileImage(context.Context, *RequestSetProfileImg) (*ResponseErrorMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserProfileImage not implemented")
}
func (UnimplementedAuthServiceServer) GetUserDetailsLiteForPostView(context.Context, *RequestUserId) (*ResponseUserDetailsLite, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserDetailsLiteForPostView not implemented")
}
func (UnimplementedAuthServiceServer) CheckUserExist(context.Context, *RequestUserId) (*ResponseBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUserExist not implemented")
}
func (UnimplementedAuthServiceServer) AdminLogin(context.Context, *AdminLoginRequest) (*AdminLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminLogin not implemented")
}
func (UnimplementedAuthServiceServer) GetAllUsers(context.Context, *GetAllUsersRequest) (*GetAllUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}
func (UnimplementedAuthServiceServer) BlockUser(context.Context, *BlockUserRequest) (*ResponseErrorMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockUser not implemented")
}
func (UnimplementedAuthServiceServer) UnblockUser(context.Context, *UnblockUserRequest) (*ResponseErrorMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnblockUser not implemented")
}
func (UnimplementedAuthServiceServer) VerifyAdminToken(context.Context, *RequestVerifyAdmin) (*ResponseVerifyAdmin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyAdminToken not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}
func (UnimplementedAuthServiceServer) testEmbeddedByValue()                     {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	// If the following call pancis, it indicates UnimplementedAuthServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_UserSignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UserSignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_UserSignUp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UserSignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UserOTPVerification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestOtpVefification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UserOTPVerification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_UserOTPVerification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UserOTPVerification(ctx, req.(*RequestOtpVefification))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUserLogin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_UserLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UserLogin(ctx, req.(*RequestUserLogin))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ForgotPasswordRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestForgotPass)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ForgotPasswordRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ForgotPasswordRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ForgotPasswordRequest(ctx, req.(*RequestForgotPass))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestResetPass)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ResetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ResetPassword(ctx, req.(*RequestResetPass))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_VerifyAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestVerifyAccess)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).VerifyAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_VerifyAccessToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).VerifyAccessToken(ctx, req.(*RequestVerifyAccess))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AccessRegenerator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAccessGenerator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AccessRegenerator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_AccessRegenerator_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AccessRegenerator(ctx, req.(*RequestAccessGenerator))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGetUserProfile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GetUserProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetUserProfile(ctx, req.(*RequestGetUserProfile))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_EditUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestEditUserProfile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).EditUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_EditUserProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).EditUserProfile(ctx, req.(*RequestEditUserProfile))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetFollowersDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetFollowersDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GetFollowersDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetFollowersDetails(ctx, req.(*RequestUserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetFollowingsDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetFollowingsDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GetFollowingsDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetFollowingsDetails(ctx, req.(*RequestUserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_SearchUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUserSearch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SearchUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_SearchUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SearchUser(ctx, req.(*RequestUserSearch))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_SetUserProfileImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestSetProfileImg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SetUserProfileImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_SetUserProfileImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SetUserProfileImage(ctx, req.(*RequestSetProfileImg))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetUserDetailsLiteForPostView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetUserDetailsLiteForPostView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GetUserDetailsLiteForPostView_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetUserDetailsLiteForPostView(ctx, req.(*RequestUserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CheckUserExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CheckUserExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_CheckUserExist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CheckUserExist(ctx, req.(*RequestUserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AdminLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AdminLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_AdminLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AdminLogin(ctx, req.(*AdminLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GetAllUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetAllUsers(ctx, req.(*GetAllUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_BlockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).BlockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_BlockUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).BlockUser(ctx, req.(*BlockUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UnblockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnblockUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UnblockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_UnblockUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UnblockUser(ctx, req.(*UnblockUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_VerifyAdminToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestVerifyAdmin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).VerifyAdminToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_VerifyAdminToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).VerifyAdminToken(ctx, req.(*RequestVerifyAdmin))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth_proto.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserSignUp",
			Handler:    _AuthService_UserSignUp_Handler,
		},
		{
			MethodName: "UserOTPVerification",
			Handler:    _AuthService_UserOTPVerification_Handler,
		},
		{
			MethodName: "UserLogin",
			Handler:    _AuthService_UserLogin_Handler,
		},
		{
			MethodName: "ForgotPasswordRequest",
			Handler:    _AuthService_ForgotPasswordRequest_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _AuthService_ResetPassword_Handler,
		},
		{
			MethodName: "VerifyAccessToken",
			Handler:    _AuthService_VerifyAccessToken_Handler,
		},
		{
			MethodName: "AccessRegenerator",
			Handler:    _AuthService_AccessRegenerator_Handler,
		},
		{
			MethodName: "GetUserProfile",
			Handler:    _AuthService_GetUserProfile_Handler,
		},
		{
			MethodName: "EditUserProfile",
			Handler:    _AuthService_EditUserProfile_Handler,
		},
		{
			MethodName: "GetFollowersDetails",
			Handler:    _AuthService_GetFollowersDetails_Handler,
		},
		{
			MethodName: "GetFollowingsDetails",
			Handler:    _AuthService_GetFollowingsDetails_Handler,
		},
		{
			MethodName: "SearchUser",
			Handler:    _AuthService_SearchUser_Handler,
		},
		{
			MethodName: "SetUserProfileImage",
			Handler:    _AuthService_SetUserProfileImage_Handler,
		},
		{
			MethodName: "GetUserDetailsLiteForPostView",
			Handler:    _AuthService_GetUserDetailsLiteForPostView_Handler,
		},
		{
			MethodName: "CheckUserExist",
			Handler:    _AuthService_CheckUserExist_Handler,
		},
		{
			MethodName: "AdminLogin",
			Handler:    _AuthService_AdminLogin_Handler,
		},
		{
			MethodName: "GetAllUsers",
			Handler:    _AuthService_GetAllUsers_Handler,
		},
		{
			MethodName: "BlockUser",
			Handler:    _AuthService_BlockUser_Handler,
		},
		{
			MethodName: "UnblockUser",
			Handler:    _AuthService_UnblockUser_Handler,
		},
		{
			MethodName: "VerifyAdminToken",
			Handler:    _AuthService_VerifyAdminToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/Auth_Service/pb/auth.proto",
}