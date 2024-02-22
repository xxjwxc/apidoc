// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package shares

import (
	context "context"
	micro "github.com/gmsec/micro"
	client "github.com/gmsec/micro/client"
	server "github.com/gmsec/micro/server"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	common "rpc/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface
var _ server.Server
var _ client.Client
var _ micro.Service

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WeixinClient is the client API for Weixin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WeixinClient interface {
	// Oauth 微信授权获取登录信息
	Oauth(ctx context.Context, in *OauthReq, opts ...grpc.CallOption) (*OauthResp, error)
	// GetQrcode 获取微信二维码
	GetQrcode(ctx context.Context, in *GetQrcodeReq, opts ...grpc.CallOption) (*GetQrcodeResp, error)
	// GetUserInfo 获取用户信息
	GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error)
	// UpsetUserInfo 更新用户信息
	UpsetUserInfo(ctx context.Context, in *UpsetUserInfoReq, opts ...grpc.CallOption) (*common.Empty, error)
	// ReLogin 是否要重新登录
	ReLogin(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*ReLoginResp, error)
	// ReLogin 是否要重新登录
	GetJsSign(ctx context.Context, in *GetJsSignReq, opts ...grpc.CallOption) (*GetJsSignResp, error)
}

type weixinClient struct {
	cc client.Client
}

// GetWeixinName get client name(package.class)
func GetWeixinName() string {
	return "shares.Weixin"
}

// GetWeixinClient get client by clientname
func GetWeixinClient() WeixinClient {
	cc := micro.GetClient(GetWeixinName())
	return &weixinClient{cc}
}

// GetWeixinClientByName get client by custom name
func GetWeixinClientByName(name string) WeixinClient {
	cc := micro.GetClient(name)
	return &weixinClient{cc}
}

func NewWeixinClient(cc client.Client) WeixinClient {
	return &weixinClient{cc}
}

func (c *weixinClient) Oauth(ctx context.Context, in *OauthReq, opts ...grpc.CallOption) (*OauthResp, error) {
	conn, err := c.cc.Next()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	out := new(OauthResp)
	err = conn.Invoke(ctx, "/shares.Weixin/Oauth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weixinClient) GetQrcode(ctx context.Context, in *GetQrcodeReq, opts ...grpc.CallOption) (*GetQrcodeResp, error) {
	conn, err := c.cc.Next()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	out := new(GetQrcodeResp)
	err = conn.Invoke(ctx, "/shares.Weixin/GetQrcode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weixinClient) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error) {
	conn, err := c.cc.Next()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	out := new(GetUserInfoResp)
	err = conn.Invoke(ctx, "/shares.Weixin/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weixinClient) UpsetUserInfo(ctx context.Context, in *UpsetUserInfoReq, opts ...grpc.CallOption) (*common.Empty, error) {
	conn, err := c.cc.Next()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	out := new(common.Empty)
	err = conn.Invoke(ctx, "/shares.Weixin/UpsetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weixinClient) ReLogin(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*ReLoginResp, error) {
	conn, err := c.cc.Next()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	out := new(ReLoginResp)
	err = conn.Invoke(ctx, "/shares.Weixin/ReLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weixinClient) GetJsSign(ctx context.Context, in *GetJsSignReq, opts ...grpc.CallOption) (*GetJsSignResp, error) {
	conn, err := c.cc.Next()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	out := new(GetJsSignResp)
	err = conn.Invoke(ctx, "/shares.Weixin/GetJsSign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WeixinServer is the server API for Weixin service.
type WeixinServer interface {
	// Oauth 微信授权获取登录信息
	Oauth(context.Context, *OauthReq) (*OauthResp, error)
	// GetQrcode 获取微信二维码
	GetQrcode(context.Context, *GetQrcodeReq) (*GetQrcodeResp, error)
	// GetUserInfo 获取用户信息
	GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoResp, error)
	// UpsetUserInfo 更新用户信息
	UpsetUserInfo(context.Context, *UpsetUserInfoReq) (*common.Empty, error)
	// ReLogin 是否要重新登录
	ReLogin(context.Context, *GetUserInfoReq) (*ReLoginResp, error)
	// ReLogin 是否要重新登录
	GetJsSign(context.Context, *GetJsSignReq) (*GetJsSignResp, error)
}

// UnimplementedWeixinServer can be embedded to have forward compatible implementations.
type UnimplementedWeixinServer struct {
}

func (*UnimplementedWeixinServer) Oauth(context.Context, *OauthReq) (*OauthResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Oauth not implemented")
}
func (*UnimplementedWeixinServer) GetQrcode(context.Context, *GetQrcodeReq) (*GetQrcodeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQrcode not implemented")
}
func (*UnimplementedWeixinServer) GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (*UnimplementedWeixinServer) UpsetUserInfo(context.Context, *UpsetUserInfoReq) (*common.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsetUserInfo not implemented")
}
func (*UnimplementedWeixinServer) ReLogin(context.Context, *GetUserInfoReq) (*ReLoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReLogin not implemented")
}
func (*UnimplementedWeixinServer) GetJsSign(context.Context, *GetJsSignReq) (*GetJsSignResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJsSign not implemented")
}

func RegisterWeixinServer(s server.Server, srv WeixinServer) {
	s.GetServer().RegisterService(&_Weixin_serviceDesc, srv)
}

func _Weixin_Oauth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OauthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeixinServer).Oauth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shares.Weixin/Oauth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeixinServer).Oauth(ctx, req.(*OauthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Weixin_GetQrcode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQrcodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeixinServer).GetQrcode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shares.Weixin/GetQrcode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeixinServer).GetQrcode(ctx, req.(*GetQrcodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Weixin_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeixinServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shares.Weixin/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeixinServer).GetUserInfo(ctx, req.(*GetUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Weixin_UpsetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsetUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeixinServer).UpsetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shares.Weixin/UpsetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeixinServer).UpsetUserInfo(ctx, req.(*UpsetUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Weixin_ReLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeixinServer).ReLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shares.Weixin/ReLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeixinServer).ReLogin(ctx, req.(*GetUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Weixin_GetJsSign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJsSignReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeixinServer).GetJsSign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shares.Weixin/GetJsSign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeixinServer).GetJsSign(ctx, req.(*GetJsSignReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Weixin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shares.Weixin",
	HandlerType: (*WeixinServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Oauth",
			Handler:    _Weixin_Oauth_Handler,
		},
		{
			MethodName: "GetQrcode",
			Handler:    _Weixin_GetQrcode_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _Weixin_GetUserInfo_Handler,
		},
		{
			MethodName: "UpsetUserInfo",
			Handler:    _Weixin_UpsetUserInfo_Handler,
		},
		{
			MethodName: "ReLogin",
			Handler:    _Weixin_ReLogin_Handler,
		},
		{
			MethodName: "GetJsSign",
			Handler:    _Weixin_GetJsSign_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shares/weixin.proto",
}
