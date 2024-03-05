// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: protos/ticket.proto

package protos

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
	TicketService_PurchaseTicket_FullMethodName    = "/reservationapp.TicketService/PurchaseTicket"
	TicketService_ViewSeatBySection_FullMethodName = "/reservationapp.TicketService/ViewSeatBySection"
	TicketService_ViewSeatByEmail_FullMethodName   = "/reservationapp.TicketService/ViewSeatByEmail"
	TicketService_RemoveUser_FullMethodName        = "/reservationapp.TicketService/RemoveUser"
	TicketService_ModifySeat_FullMethodName        = "/reservationapp.TicketService/ModifySeat"
)

// TicketServiceClient is the client API for TicketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TicketServiceClient interface {
	PurchaseTicket(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*PurchaseResponse, error)
	ViewSeatBySection(ctx context.Context, in *SectionID, opts ...grpc.CallOption) (TicketService_ViewSeatBySectionClient, error)
	ViewSeatByEmail(ctx context.Context, in *Email, opts ...grpc.CallOption) (*SeatInfo, error)
	RemoveUser(ctx context.Context, in *Email, opts ...grpc.CallOption) (*Status, error)
	ModifySeat(ctx context.Context, in *ModifySeatRequest, opts ...grpc.CallOption) (*PurchaseResponse, error)
}

type ticketServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTicketServiceClient(cc grpc.ClientConnInterface) TicketServiceClient {
	return &ticketServiceClient{cc}
}

func (c *ticketServiceClient) PurchaseTicket(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*PurchaseResponse, error) {
	out := new(PurchaseResponse)
	err := c.cc.Invoke(ctx, TicketService_PurchaseTicket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) ViewSeatBySection(ctx context.Context, in *SectionID, opts ...grpc.CallOption) (TicketService_ViewSeatBySectionClient, error) {
	stream, err := c.cc.NewStream(ctx, &TicketService_ServiceDesc.Streams[0], TicketService_ViewSeatBySection_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &ticketServiceViewSeatBySectionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TicketService_ViewSeatBySectionClient interface {
	Recv() (*SeatInfo, error)
	grpc.ClientStream
}

type ticketServiceViewSeatBySectionClient struct {
	grpc.ClientStream
}

func (x *ticketServiceViewSeatBySectionClient) Recv() (*SeatInfo, error) {
	m := new(SeatInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ticketServiceClient) ViewSeatByEmail(ctx context.Context, in *Email, opts ...grpc.CallOption) (*SeatInfo, error) {
	out := new(SeatInfo)
	err := c.cc.Invoke(ctx, TicketService_ViewSeatByEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) RemoveUser(ctx context.Context, in *Email, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, TicketService_RemoveUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) ModifySeat(ctx context.Context, in *ModifySeatRequest, opts ...grpc.CallOption) (*PurchaseResponse, error) {
	out := new(PurchaseResponse)
	err := c.cc.Invoke(ctx, TicketService_ModifySeat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TicketServiceServer is the server API for TicketService service.
// All implementations must embed UnimplementedTicketServiceServer
// for forward compatibility
type TicketServiceServer interface {
	PurchaseTicket(context.Context, *PurchaseRequest) (*PurchaseResponse, error)
	ViewSeatBySection(*SectionID, TicketService_ViewSeatBySectionServer) error
	ViewSeatByEmail(context.Context, *Email) (*SeatInfo, error)
	RemoveUser(context.Context, *Email) (*Status, error)
	ModifySeat(context.Context, *ModifySeatRequest) (*PurchaseResponse, error)
	mustEmbedUnimplementedTicketServiceServer()
}

// UnimplementedTicketServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTicketServiceServer struct {
}

func (UnimplementedTicketServiceServer) PurchaseTicket(context.Context, *PurchaseRequest) (*PurchaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurchaseTicket not implemented")
}
func (UnimplementedTicketServiceServer) ViewSeatBySection(*SectionID, TicketService_ViewSeatBySectionServer) error {
	return status.Errorf(codes.Unimplemented, "method ViewSeatBySection not implemented")
}
func (UnimplementedTicketServiceServer) ViewSeatByEmail(context.Context, *Email) (*SeatInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewSeatByEmail not implemented")
}
func (UnimplementedTicketServiceServer) RemoveUser(context.Context, *Email) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUser not implemented")
}
func (UnimplementedTicketServiceServer) ModifySeat(context.Context, *ModifySeatRequest) (*PurchaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifySeat not implemented")
}
func (UnimplementedTicketServiceServer) mustEmbedUnimplementedTicketServiceServer() {}

// UnsafeTicketServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TicketServiceServer will
// result in compilation errors.
type UnsafeTicketServiceServer interface {
	mustEmbedUnimplementedTicketServiceServer()
}

func RegisterTicketServiceServer(s grpc.ServiceRegistrar, srv TicketServiceServer) {
	s.RegisterService(&TicketService_ServiceDesc, srv)
}

func _TicketService_PurchaseTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurchaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).PurchaseTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_PurchaseTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).PurchaseTicket(ctx, req.(*PurchaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_ViewSeatBySection_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SectionID)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TicketServiceServer).ViewSeatBySection(m, &ticketServiceViewSeatBySectionServer{stream})
}

type TicketService_ViewSeatBySectionServer interface {
	Send(*SeatInfo) error
	grpc.ServerStream
}

type ticketServiceViewSeatBySectionServer struct {
	grpc.ServerStream
}

func (x *ticketServiceViewSeatBySectionServer) Send(m *SeatInfo) error {
	return x.ServerStream.SendMsg(m)
}

func _TicketService_ViewSeatByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Email)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).ViewSeatByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_ViewSeatByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).ViewSeatByEmail(ctx, req.(*Email))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_RemoveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Email)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).RemoveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_RemoveUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).RemoveUser(ctx, req.(*Email))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_ModifySeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifySeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).ModifySeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_ModifySeat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).ModifySeat(ctx, req.(*ModifySeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TicketService_ServiceDesc is the grpc.ServiceDesc for TicketService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TicketService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reservationapp.TicketService",
	HandlerType: (*TicketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PurchaseTicket",
			Handler:    _TicketService_PurchaseTicket_Handler,
		},
		{
			MethodName: "ViewSeatByEmail",
			Handler:    _TicketService_ViewSeatByEmail_Handler,
		},
		{
			MethodName: "RemoveUser",
			Handler:    _TicketService_RemoveUser_Handler,
		},
		{
			MethodName: "ModifySeat",
			Handler:    _TicketService_ModifySeat_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ViewSeatBySection",
			Handler:       _TicketService_ViewSeatBySection_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/ticket.proto",
}
