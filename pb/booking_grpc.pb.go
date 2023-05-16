// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: services/booking/proto/booking.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BookingServiceClient is the client API for BookingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookingServiceClient interface {
	PostBooking(ctx context.Context, in *BookingInput, opts ...grpc.CallOption) (*Booking, error)
	GetBooking(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*Booking, error)
	GetBookings(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Bookings, error)
	PostTicket(ctx context.Context, in *Ticket, opts ...grpc.CallOption) (*Ticket, error)
	GetTicket(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*Ticket, error)
	GetTickets(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Tickets, error)
}

type bookingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookingServiceClient(cc grpc.ClientConnInterface) BookingServiceClient {
	return &bookingServiceClient{cc}
}

func (c *bookingServiceClient) PostBooking(ctx context.Context, in *BookingInput, opts ...grpc.CallOption) (*Booking, error) {
	out := new(Booking)
	err := c.cc.Invoke(ctx, "/pb.BookingService/PostBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetBooking(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*Booking, error) {
	out := new(Booking)
	err := c.cc.Invoke(ctx, "/pb.BookingService/GetBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetBookings(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Bookings, error) {
	out := new(Bookings)
	err := c.cc.Invoke(ctx, "/pb.BookingService/GetBookings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) PostTicket(ctx context.Context, in *Ticket, opts ...grpc.CallOption) (*Ticket, error) {
	out := new(Ticket)
	err := c.cc.Invoke(ctx, "/pb.BookingService/PostTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetTicket(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*Ticket, error) {
	out := new(Ticket)
	err := c.cc.Invoke(ctx, "/pb.BookingService/GetTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetTickets(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Tickets, error) {
	out := new(Tickets)
	err := c.cc.Invoke(ctx, "/pb.BookingService/GetTickets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookingServiceServer is the server API for BookingService service.
// All implementations must embed UnimplementedBookingServiceServer
// for forward compatibility
type BookingServiceServer interface {
	PostBooking(context.Context, *BookingInput) (*Booking, error)
	GetBooking(context.Context, *UUID) (*Booking, error)
	GetBookings(context.Context, *emptypb.Empty) (*Bookings, error)
	PostTicket(context.Context, *Ticket) (*Ticket, error)
	GetTicket(context.Context, *UUID) (*Ticket, error)
	GetTickets(context.Context, *emptypb.Empty) (*Tickets, error)
	mustEmbedUnimplementedBookingServiceServer()
}

// UnimplementedBookingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookingServiceServer struct {
}

func (UnimplementedBookingServiceServer) PostBooking(context.Context, *BookingInput) (*Booking, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostBooking not implemented")
}
func (UnimplementedBookingServiceServer) GetBooking(context.Context, *UUID) (*Booking, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBooking not implemented")
}
func (UnimplementedBookingServiceServer) GetBookings(context.Context, *emptypb.Empty) (*Bookings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookings not implemented")
}
func (UnimplementedBookingServiceServer) PostTicket(context.Context, *Ticket) (*Ticket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostTicket not implemented")
}
func (UnimplementedBookingServiceServer) GetTicket(context.Context, *UUID) (*Ticket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTicket not implemented")
}
func (UnimplementedBookingServiceServer) GetTickets(context.Context, *emptypb.Empty) (*Tickets, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTickets not implemented")
}
func (UnimplementedBookingServiceServer) mustEmbedUnimplementedBookingServiceServer() {}

// UnsafeBookingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookingServiceServer will
// result in compilation errors.
type UnsafeBookingServiceServer interface {
	mustEmbedUnimplementedBookingServiceServer()
}

func RegisterBookingServiceServer(s grpc.ServiceRegistrar, srv BookingServiceServer) {
	s.RegisterService(&BookingService_ServiceDesc, srv)
}

func _BookingService_PostBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookingInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).PostBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BookingService/PostBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).PostBooking(ctx, req.(*BookingInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BookingService/GetBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetBooking(ctx, req.(*UUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetBookings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetBookings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BookingService/GetBookings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetBookings(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_PostTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ticket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).PostTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BookingService/PostTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).PostTicket(ctx, req.(*Ticket))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BookingService/GetTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetTicket(ctx, req.(*UUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BookingService/GetTickets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetTickets(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// BookingService_ServiceDesc is the grpc.ServiceDesc for BookingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.BookingService",
	HandlerType: (*BookingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostBooking",
			Handler:    _BookingService_PostBooking_Handler,
		},
		{
			MethodName: "GetBooking",
			Handler:    _BookingService_GetBooking_Handler,
		},
		{
			MethodName: "GetBookings",
			Handler:    _BookingService_GetBookings_Handler,
		},
		{
			MethodName: "PostTicket",
			Handler:    _BookingService_PostTicket_Handler,
		},
		{
			MethodName: "GetTicket",
			Handler:    _BookingService_GetTicket_Handler,
		},
		{
			MethodName: "GetTickets",
			Handler:    _BookingService_GetTickets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/booking/proto/booking.proto",
}