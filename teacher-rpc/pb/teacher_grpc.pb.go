// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: teacher.proto

package pb

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

// TeachercenterClient is the client API for Teachercenter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TeachercenterClient interface {
	GetTeacherInfo(ctx context.Context, in *GetTeacherInfoReq, opts ...grpc.CallOption) (*GetTeacherInfoResp, error)
}

type teachercenterClient struct {
	cc grpc.ClientConnInterface
}

func NewTeachercenterClient(cc grpc.ClientConnInterface) TeachercenterClient {
	return &teachercenterClient{cc}
}

func (c *teachercenterClient) GetTeacherInfo(ctx context.Context, in *GetTeacherInfoReq, opts ...grpc.CallOption) (*GetTeacherInfoResp, error) {
	out := new(GetTeacherInfoResp)
	err := c.cc.Invoke(ctx, "/pb.teachercenter/GetTeacherInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeachercenterServer is the server API for Teachercenter service.
// All implementations must embed UnimplementedTeachercenterServer
// for forward compatibility
type TeachercenterServer interface {
	GetTeacherInfo(context.Context, *GetTeacherInfoReq) (*GetTeacherInfoResp, error)
	mustEmbedUnimplementedTeachercenterServer()
}

// UnimplementedTeachercenterServer must be embedded to have forward compatible implementations.
type UnimplementedTeachercenterServer struct {
}

func (UnimplementedTeachercenterServer) GetTeacherInfo(context.Context, *GetTeacherInfoReq) (*GetTeacherInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeacherInfo not implemented")
}
func (UnimplementedTeachercenterServer) mustEmbedUnimplementedTeachercenterServer() {}

// UnsafeTeachercenterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TeachercenterServer will
// result in compilation errors.
type UnsafeTeachercenterServer interface {
	mustEmbedUnimplementedTeachercenterServer()
}

func RegisterTeachercenterServer(s grpc.ServiceRegistrar, srv TeachercenterServer) {
	s.RegisterService(&Teachercenter_ServiceDesc, srv)
}

func _Teachercenter_GetTeacherInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeacherInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeachercenterServer).GetTeacherInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.teachercenter/GetTeacherInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeachercenterServer).GetTeacherInfo(ctx, req.(*GetTeacherInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Teachercenter_ServiceDesc is the grpc.ServiceDesc for Teachercenter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Teachercenter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.teachercenter",
	HandlerType: (*TeachercenterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTeacherInfo",
			Handler:    _Teachercenter_GetTeacherInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teacher.proto",
}