// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package monitoring_backup

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// MonitoringBackupClient is the client API for MonitoringBackup service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MonitoringBackupClient interface {
	SendNotif(ctx context.Context, in *CreatedNotify, opts ...grpc.CallOption) (*ResponseNotif, error)
	GetMonitoringLogs(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*MonitoringLogs, error)
}

type monitoringBackupClient struct {
	cc grpc.ClientConnInterface
}

func NewMonitoringBackupClient(cc grpc.ClientConnInterface) MonitoringBackupClient {
	return &monitoringBackupClient{cc}
}

func (c *monitoringBackupClient) SendNotif(ctx context.Context, in *CreatedNotify, opts ...grpc.CallOption) (*ResponseNotif, error) {
	out := new(ResponseNotif)
	err := c.cc.Invoke(ctx, "/monitoring_backup.MonitoringBackup/SendNotif", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitoringBackupClient) GetMonitoringLogs(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*MonitoringLogs, error) {
	out := new(MonitoringLogs)
	err := c.cc.Invoke(ctx, "/monitoring_backup.MonitoringBackup/GetMonitoringLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MonitoringBackupServer is the server API for MonitoringBackup service.
// All implementations must embed UnimplementedMonitoringBackupServer
// for forward compatibility
type MonitoringBackupServer interface {
	SendNotif(context.Context, *CreatedNotify) (*ResponseNotif, error)
	GetMonitoringLogs(context.Context, *empty.Empty) (*MonitoringLogs, error)
	mustEmbedUnimplementedMonitoringBackupServer()
}

// UnimplementedMonitoringBackupServer must be embedded to have forward compatible implementations.
type UnimplementedMonitoringBackupServer struct {
}

func (UnimplementedMonitoringBackupServer) SendNotif(context.Context, *CreatedNotify) (*ResponseNotif, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendNotif not implemented")
}
func (UnimplementedMonitoringBackupServer) GetMonitoringLogs(context.Context, *empty.Empty) (*MonitoringLogs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMonitoringLogs not implemented")
}
func (UnimplementedMonitoringBackupServer) mustEmbedUnimplementedMonitoringBackupServer() {}

// UnsafeMonitoringBackupServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MonitoringBackupServer will
// result in compilation errors.
type UnsafeMonitoringBackupServer interface {
	mustEmbedUnimplementedMonitoringBackupServer()
}

func RegisterMonitoringBackupServer(s grpc.ServiceRegistrar, srv MonitoringBackupServer) {
	s.RegisterService(&_MonitoringBackup_serviceDesc, srv)
}

func _MonitoringBackup_SendNotif_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatedNotify)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringBackupServer).SendNotif(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitoring_backup.MonitoringBackup/SendNotif",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringBackupServer).SendNotif(ctx, req.(*CreatedNotify))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitoringBackup_GetMonitoringLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringBackupServer).GetMonitoringLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitoring_backup.MonitoringBackup/GetMonitoringLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringBackupServer).GetMonitoringLogs(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _MonitoringBackup_serviceDesc = grpc.ServiceDesc{
	ServiceName: "monitoring_backup.MonitoringBackup",
	HandlerType: (*MonitoringBackupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendNotif",
			Handler:    _MonitoringBackup_SendNotif_Handler,
		},
		{
			MethodName: "GetMonitoringLogs",
			Handler:    _MonitoringBackup_GetMonitoringLogs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "monitoring_backup.proto",
}
