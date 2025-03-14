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

// AnalyClient is the client API for Analy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AnalyClient interface {
	// AnalyCode 分析一直股票
	AnalyCode(ctx context.Context, in *AnalyCodeReq, opts ...grpc.CallOption) (*AnalyCodeResp, error)
	// GetAllSp 获取票的特色数据
	GetSp(ctx context.Context, in *GetAllSpReq, opts ...grpc.CallOption) (*GetAllSpResp, error)
	GetBaseInfo(ctx context.Context, in *GetAllSpReq, opts ...grpc.CallOption) (*GetBaseInfoResp, error)
	GetSampleWdj(ctx context.Context, in *GetSampleWdjReq, opts ...grpc.CallOption) (*GetSampleWdjResp, error)
	GetGjzb(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*GjzbInfo, error)
	// GetBxNxKlineInfo 获取北向，南向当日净流入情况
	GetBxNxKlineInfo(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*GetBxNxKlineInfoResp, error)
	GetCjrl(ctx context.Context, in *GetCjrlReq, opts ...grpc.CallOption) (*GetCjrlResp, error)
	GetGdData(ctx context.Context, in *GetGdDataReq, opts ...grpc.CallOption) (*GetGdDataResp, error)
}

type analyClient struct {
	cc client.Client
}

// GetAnalyName get client name(package.class)
func GetAnalyName() string {
	return "shares.Analy"
}

// GetAnalyClient get client by clientname
func GetAnalyClient() AnalyClient {
	cc := micro.GetClient(GetAnalyName())
	return &analyClient{cc}
}

// GetAnalyClientByName get client by custom name
func GetAnalyClientByName(name string) AnalyClient {
	cc := micro.GetClient(name)
	return &analyClient{cc}
}

func NewAnalyClient(cc client.Client) AnalyClient {
	return &analyClient{cc}
}

func (c *analyClient) AnalyCode(ctx context.Context, in *AnalyCodeReq, opts ...grpc.CallOption) (*AnalyCodeResp, error) {
	conn, err := c.cc.Next()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	out := new(AnalyCodeResp)
	err = conn.Invoke(ctx, "/shares.Analy/AnalyCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyClient) GetSp(ctx context.Context, in *GetAllSpReq, opts ...grpc.CallOption) (*GetAllSpResp, error) {
	conn, err := c.cc.Next()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	out := new(GetAllSpResp)
	err = conn.Invoke(ctx, "/shares.Analy/GetSp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyClient) GetBaseInfo(ctx context.Context, in *GetAllSpReq, opts ...grpc.CallOption) (*GetBaseInfoResp, error) {
	conn, err := c.cc.Next()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	out := new(GetBaseInfoResp)
	err = conn.Invoke(ctx, "/shares.Analy/GetBaseInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyClient) GetSampleWdj(ctx context.Context, in *GetSampleWdjReq, opts ...grpc.CallOption) (*GetSampleWdjResp, error) {
	conn, err := c.cc.Next()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	out := new(GetSampleWdjResp)
	err = conn.Invoke(ctx, "/shares.Analy/GetSampleWdj", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyClient) GetGjzb(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*GjzbInfo, error) {
	conn, err := c.cc.Next()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	out := new(GjzbInfo)
	err = conn.Invoke(ctx, "/shares.Analy/GetGjzb", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyClient) GetBxNxKlineInfo(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*GetBxNxKlineInfoResp, error) {
	conn, err := c.cc.Next()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	out := new(GetBxNxKlineInfoResp)
	err = conn.Invoke(ctx, "/shares.Analy/GetBxNxKlineInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyClient) GetCjrl(ctx context.Context, in *GetCjrlReq, opts ...grpc.CallOption) (*GetCjrlResp, error) {
	conn, err := c.cc.Next()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	out := new(GetCjrlResp)
	err = conn.Invoke(ctx, "/shares.Analy/GetCjrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyClient) GetGdData(ctx context.Context, in *GetGdDataReq, opts ...grpc.CallOption) (*GetGdDataResp, error) {
	conn, err := c.cc.Next()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	out := new(GetGdDataResp)
	err = conn.Invoke(ctx, "/shares.Analy/GetGdData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnalyServer is the server API for Analy service.
type AnalyServer interface {
	// AnalyCode 分析一直股票
	AnalyCode(context.Context, *AnalyCodeReq) (*AnalyCodeResp, error)
	// GetAllSp 获取票的特色数据
	GetSp(context.Context, *GetAllSpReq) (*GetAllSpResp, error)
	GetBaseInfo(context.Context, *GetAllSpReq) (*GetBaseInfoResp, error)
	GetSampleWdj(context.Context, *GetSampleWdjReq) (*GetSampleWdjResp, error)
	GetGjzb(context.Context, *common.Empty) (*GjzbInfo, error)
	// GetBxNxKlineInfo 获取北向，南向当日净流入情况
	GetBxNxKlineInfo(context.Context, *common.Empty) (*GetBxNxKlineInfoResp, error)
	GetCjrl(context.Context, *GetCjrlReq) (*GetCjrlResp, error)
	GetGdData(context.Context, *GetGdDataReq) (*GetGdDataResp, error)
}

// UnimplementedAnalyServer can be embedded to have forward compatible implementations.
type UnimplementedAnalyServer struct {
}

func (*UnimplementedAnalyServer) AnalyCode(context.Context, *AnalyCodeReq) (*AnalyCodeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnalyCode not implemented")
}
func (*UnimplementedAnalyServer) GetSp(context.Context, *GetAllSpReq) (*GetAllSpResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSp not implemented")
}
func (*UnimplementedAnalyServer) GetBaseInfo(context.Context, *GetAllSpReq) (*GetBaseInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBaseInfo not implemented")
}
func (*UnimplementedAnalyServer) GetSampleWdj(context.Context, *GetSampleWdjReq) (*GetSampleWdjResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSampleWdj not implemented")
}
func (*UnimplementedAnalyServer) GetGjzb(context.Context, *common.Empty) (*GjzbInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGjzb not implemented")
}
func (*UnimplementedAnalyServer) GetBxNxKlineInfo(context.Context, *common.Empty) (*GetBxNxKlineInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBxNxKlineInfo not implemented")
}
func (*UnimplementedAnalyServer) GetCjrl(context.Context, *GetCjrlReq) (*GetCjrlResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCjrl not implemented")
}
func (*UnimplementedAnalyServer) GetGdData(context.Context, *GetGdDataReq) (*GetGdDataResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGdData not implemented")
}

func RegisterAnalyServer(s server.Server, srv AnalyServer) {
	s.GetServer().RegisterService(&_Analy_serviceDesc, srv)
}

func _Analy_AnalyCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnalyCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyServer).AnalyCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shares.Analy/AnalyCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyServer).AnalyCode(ctx, req.(*AnalyCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analy_GetSp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllSpReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyServer).GetSp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shares.Analy/GetSp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyServer).GetSp(ctx, req.(*GetAllSpReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analy_GetBaseInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllSpReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyServer).GetBaseInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shares.Analy/GetBaseInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyServer).GetBaseInfo(ctx, req.(*GetAllSpReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analy_GetSampleWdj_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSampleWdjReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyServer).GetSampleWdj(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shares.Analy/GetSampleWdj",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyServer).GetSampleWdj(ctx, req.(*GetSampleWdjReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analy_GetGjzb_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyServer).GetGjzb(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shares.Analy/GetGjzb",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyServer).GetGjzb(ctx, req.(*common.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analy_GetBxNxKlineInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyServer).GetBxNxKlineInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shares.Analy/GetBxNxKlineInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyServer).GetBxNxKlineInfo(ctx, req.(*common.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analy_GetCjrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCjrlReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyServer).GetCjrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shares.Analy/GetCjrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyServer).GetCjrl(ctx, req.(*GetCjrlReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analy_GetGdData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGdDataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyServer).GetGdData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shares.Analy/GetGdData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyServer).GetGdData(ctx, req.(*GetGdDataReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Analy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shares.Analy",
	HandlerType: (*AnalyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AnalyCode",
			Handler:    _Analy_AnalyCode_Handler,
		},
		{
			MethodName: "GetSp",
			Handler:    _Analy_GetSp_Handler,
		},
		{
			MethodName: "GetBaseInfo",
			Handler:    _Analy_GetBaseInfo_Handler,
		},
		{
			MethodName: "GetSampleWdj",
			Handler:    _Analy_GetSampleWdj_Handler,
		},
		{
			MethodName: "GetGjzb",
			Handler:    _Analy_GetGjzb_Handler,
		},
		{
			MethodName: "GetBxNxKlineInfo",
			Handler:    _Analy_GetBxNxKlineInfo_Handler,
		},
		{
			MethodName: "GetCjrl",
			Handler:    _Analy_GetCjrl_Handler,
		},
		{
			MethodName: "GetGdData",
			Handler:    _Analy_GetGdData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shares/analy.proto",
}
