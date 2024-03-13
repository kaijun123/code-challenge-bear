// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: bear/bear/query.proto

package bear

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

const (
	Query_Params_FullMethodName             = "/bear.bear.Query/Params"
	Query_ShowBear_FullMethodName           = "/bear.bear.Query/ShowBear"
	Query_ListBear_FullMethodName           = "/bear.bear.Query/ListBear"
	Query_ListBearRole_FullMethodName       = "/bear.bear.Query/ListBearRole"
	Query_ListBearBackground_FullMethodName = "/bear.bear.Query/ListBearBackground"
	Query_ListBearClothes_FullMethodName    = "/bear.bear.Query/ListBearClothes"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of ShowBear items.
	ShowBear(ctx context.Context, in *QueryShowBearRequest, opts ...grpc.CallOption) (*QueryShowBearResponse, error)
	// Queries a list of ListBear items.
	ListBear(ctx context.Context, in *QueryListBearRequest, opts ...grpc.CallOption) (*QueryListBearResponse, error)
	// Queries a list of ListBearRole items.
	ListBearRole(ctx context.Context, in *QueryListBearRoleRequest, opts ...grpc.CallOption) (*QueryListBearRoleResponse, error)
	// Queries a list of ListBearBackground items.
	ListBearBackground(ctx context.Context, in *QueryListBearBackgroundRequest, opts ...grpc.CallOption) (*QueryListBearBackgroundResponse, error)
	// Queries a list of ListBearClothes items.
	ListBearClothes(ctx context.Context, in *QueryListBearClothesRequest, opts ...grpc.CallOption) (*QueryListBearClothesResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ShowBear(ctx context.Context, in *QueryShowBearRequest, opts ...grpc.CallOption) (*QueryShowBearResponse, error) {
	out := new(QueryShowBearResponse)
	err := c.cc.Invoke(ctx, Query_ShowBear_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListBear(ctx context.Context, in *QueryListBearRequest, opts ...grpc.CallOption) (*QueryListBearResponse, error) {
	out := new(QueryListBearResponse)
	err := c.cc.Invoke(ctx, Query_ListBear_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListBearRole(ctx context.Context, in *QueryListBearRoleRequest, opts ...grpc.CallOption) (*QueryListBearRoleResponse, error) {
	out := new(QueryListBearRoleResponse)
	err := c.cc.Invoke(ctx, Query_ListBearRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListBearBackground(ctx context.Context, in *QueryListBearBackgroundRequest, opts ...grpc.CallOption) (*QueryListBearBackgroundResponse, error) {
	out := new(QueryListBearBackgroundResponse)
	err := c.cc.Invoke(ctx, Query_ListBearBackground_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListBearClothes(ctx context.Context, in *QueryListBearClothesRequest, opts ...grpc.CallOption) (*QueryListBearClothesResponse, error) {
	out := new(QueryListBearClothesResponse)
	err := c.cc.Invoke(ctx, Query_ListBearClothes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of ShowBear items.
	ShowBear(context.Context, *QueryShowBearRequest) (*QueryShowBearResponse, error)
	// Queries a list of ListBear items.
	ListBear(context.Context, *QueryListBearRequest) (*QueryListBearResponse, error)
	// Queries a list of ListBearRole items.
	ListBearRole(context.Context, *QueryListBearRoleRequest) (*QueryListBearRoleResponse, error)
	// Queries a list of ListBearBackground items.
	ListBearBackground(context.Context, *QueryListBearBackgroundRequest) (*QueryListBearBackgroundResponse, error)
	// Queries a list of ListBearClothes items.
	ListBearClothes(context.Context, *QueryListBearClothesRequest) (*QueryListBearClothesResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) ShowBear(context.Context, *QueryShowBearRequest) (*QueryShowBearResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowBear not implemented")
}
func (UnimplementedQueryServer) ListBear(context.Context, *QueryListBearRequest) (*QueryListBearResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBear not implemented")
}
func (UnimplementedQueryServer) ListBearRole(context.Context, *QueryListBearRoleRequest) (*QueryListBearRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBearRole not implemented")
}
func (UnimplementedQueryServer) ListBearBackground(context.Context, *QueryListBearBackgroundRequest) (*QueryListBearBackgroundResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBearBackground not implemented")
}
func (UnimplementedQueryServer) ListBearClothes(context.Context, *QueryListBearClothesRequest) (*QueryListBearClothesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBearClothes not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ShowBear_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryShowBearRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ShowBear(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ShowBear_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ShowBear(ctx, req.(*QueryShowBearRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ListBear_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryListBearRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListBear(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ListBear_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListBear(ctx, req.(*QueryListBearRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ListBearRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryListBearRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListBearRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ListBearRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListBearRole(ctx, req.(*QueryListBearRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ListBearBackground_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryListBearBackgroundRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListBearBackground(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ListBearBackground_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListBearBackground(ctx, req.(*QueryListBearBackgroundRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ListBearClothes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryListBearClothesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListBearClothes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ListBearClothes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListBearClothes(ctx, req.(*QueryListBearClothesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bear.bear.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "ShowBear",
			Handler:    _Query_ShowBear_Handler,
		},
		{
			MethodName: "ListBear",
			Handler:    _Query_ListBear_Handler,
		},
		{
			MethodName: "ListBearRole",
			Handler:    _Query_ListBearRole_Handler,
		},
		{
			MethodName: "ListBearBackground",
			Handler:    _Query_ListBearBackground_Handler,
		},
		{
			MethodName: "ListBearClothes",
			Handler:    _Query_ListBearClothes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bear/bear/query.proto",
}
