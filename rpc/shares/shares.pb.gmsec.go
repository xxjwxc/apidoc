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

// SharesClient is the client API for Shares service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SharesClient interface {
	// Oauth 微信授权
	Oauth(ctx context.Context, in *OauthReq, opts ...grpc.CallOption) (*OauthResp, error)
}

type sharesClient struct {
	cc client.Client
}

// GetSharesName get client name(package.class)
func GetSharesName() string {
	return "shares.shares"
}

// GetSharesClient get client by clientname
func GetSharesClient() SharesClient {
	cc := micro.GetClient(GetSharesName())
	return &sharesClient{cc}
}

// GetSharesClientByName get client by custom name
func GetSharesClientByName(name string) SharesClient {
	cc := micro.GetClient(name)
	return &sharesClient{cc}
}

func NewSharesClient(cc client.Client) SharesClient {
	return &sharesClient{cc}
}

func (c *sharesClient) Oauth(ctx context.Context, in *OauthReq, opts ...grpc.CallOption) (*OauthResp, error) {
	conn, err := c.cc.Next()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	out := new(OauthResp)
	err = conn.Invoke(ctx, "/shares.shares/Oauth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SharesServer is the server API for Shares service.
type SharesServer interface {
	// Oauth 微信授权
	Oauth(context.Context, *OauthReq) (*OauthResp, error)
}

// UnimplementedSharesServer can be embedded to have forward compatible implementations.
type UnimplementedSharesServer struct {
}

func (*UnimplementedSharesServer) Oauth(context.Context, *OauthReq) (*OauthResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Oauth not implemented")
}

func RegisterSharesServer(s server.Server, srv SharesServer) {
	s.GetServer().RegisterService(&_Shares_serviceDesc, srv)
}

func _Shares_Oauth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OauthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharesServer).Oauth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shares.shares/Oauth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharesServer).Oauth(ctx, req.(*OauthReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Shares_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shares.shares",
	HandlerType: (*SharesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Oauth",
			Handler:    _Shares_Oauth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shares/shares.proto",
}

// WeixinClient is the client API for Weixin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WeixinClient interface {
	// Oauth 微信授权
	Oauth(ctx context.Context, in *OauthReq, opts ...grpc.CallOption) (*OauthResp, error)
}

type weixinClient struct {
	cc client.Client
}

// GetWeixinName get client name(package.class)
func GetWeixinName() string {
	return "shares.weixin"
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
	err = conn.Invoke(ctx, "/shares.weixin/Oauth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WeixinServer is the server API for Weixin service.
type WeixinServer interface {
	// Oauth 微信授权
	Oauth(context.Context, *OauthReq) (*OauthResp, error)
}

// UnimplementedWeixinServer can be embedded to have forward compatible implementations.
type UnimplementedWeixinServer struct {
}

func (*UnimplementedWeixinServer) Oauth(context.Context, *OauthReq) (*OauthResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Oauth not implemented")
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
		FullMethod: "/shares.weixin/Oauth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeixinServer).Oauth(ctx, req.(*OauthReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Weixin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shares.weixin",
	HandlerType: (*WeixinServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Oauth",
			Handler:    _Weixin_Oauth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shares/shares.proto",
}
