// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package analysispb

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

// AnalysisServiceClient is the client API for AnalysisService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnalysisServiceClient interface {
	GetYoY(ctx context.Context, in *GetYoYRequest, opts ...grpc.CallOption) (*GetYoYResponse, error)
	GetMoM(ctx context.Context, in *GetMoMRequest, opts ...grpc.CallOption) (*GetMoMResponse, error)
}

type analysisServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAnalysisServiceClient(cc grpc.ClientConnInterface) AnalysisServiceClient {
	return &analysisServiceClient{cc}
}

func (c *analysisServiceClient) GetYoY(ctx context.Context, in *GetYoYRequest, opts ...grpc.CallOption) (*GetYoYResponse, error) {
	out := new(GetYoYResponse)
	err := c.cc.Invoke(ctx, "/analysis.v1.AnalysisService/GetYoY", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analysisServiceClient) GetMoM(ctx context.Context, in *GetMoMRequest, opts ...grpc.CallOption) (*GetMoMResponse, error) {
	out := new(GetMoMResponse)
	err := c.cc.Invoke(ctx, "/analysis.v1.AnalysisService/GetMoM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnalysisServiceServer is the server API for AnalysisService service.
// All implementations must embed UnimplementedAnalysisServiceServer
// for forward compatibility
type AnalysisServiceServer interface {
	GetYoY(context.Context, *GetYoYRequest) (*GetYoYResponse, error)
	GetMoM(context.Context, *GetMoMRequest) (*GetMoMResponse, error)
	mustEmbedUnimplementedAnalysisServiceServer()
}

// UnimplementedAnalysisServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAnalysisServiceServer struct {
}

func (UnimplementedAnalysisServiceServer) GetYoY(context.Context, *GetYoYRequest) (*GetYoYResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetYoY not implemented")
}
func (UnimplementedAnalysisServiceServer) GetMoM(context.Context, *GetMoMRequest) (*GetMoMResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMoM not implemented")
}
func (UnimplementedAnalysisServiceServer) mustEmbedUnimplementedAnalysisServiceServer() {}

// UnsafeAnalysisServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnalysisServiceServer will
// result in compilation errors.
type UnsafeAnalysisServiceServer interface {
	mustEmbedUnimplementedAnalysisServiceServer()
}

func RegisterAnalysisServiceServer(s grpc.ServiceRegistrar, srv AnalysisServiceServer) {
	s.RegisterService(&AnalysisService_ServiceDesc, srv)
}

func _AnalysisService_GetYoY_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetYoYRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalysisServiceServer).GetYoY(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analysis.v1.AnalysisService/GetYoY",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalysisServiceServer).GetYoY(ctx, req.(*GetYoYRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnalysisService_GetMoM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMoMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalysisServiceServer).GetMoM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analysis.v1.AnalysisService/GetMoM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalysisServiceServer).GetMoM(ctx, req.(*GetMoMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AnalysisService_ServiceDesc is the grpc.ServiceDesc for AnalysisService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AnalysisService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "analysis.v1.AnalysisService",
	HandlerType: (*AnalysisServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetYoY",
			Handler:    _AnalysisService_GetYoY_Handler,
		},
		{
			MethodName: "GetMoM",
			Handler:    _AnalysisService_GetMoM_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "analysis.proto",
}