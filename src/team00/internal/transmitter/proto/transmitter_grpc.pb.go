// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: transmitter.proto

package __

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TransmitterService_TransmitStream_FullMethodName = "/transmitter.TransmitterService/TransmitStream"
)

// TransmitterServiceClient is the client API for TransmitterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransmitterServiceClient interface {
	TransmitStream(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (TransmitterService_TransmitStreamClient, error)
}

type transmitterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransmitterServiceClient(cc grpc.ClientConnInterface) TransmitterServiceClient {
	return &transmitterServiceClient{cc}
}

func (c *transmitterServiceClient) TransmitStream(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (TransmitterService_TransmitStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &TransmitterService_ServiceDesc.Streams[0], TransmitterService_TransmitStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &transmitterServiceTransmitStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TransmitterService_TransmitStreamClient interface {
	Recv() (*Transmission, error)
	grpc.ClientStream
}

type transmitterServiceTransmitStreamClient struct {
	grpc.ClientStream
}

func (x *transmitterServiceTransmitStreamClient) Recv() (*Transmission, error) {
	m := new(Transmission)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TransmitterServiceServer is the server API for TransmitterService service.
// All implementations must embed UnimplementedTransmitterServiceServer
// for forward compatibility
type TransmitterServiceServer interface {
	TransmitStream(*empty.Empty, TransmitterService_TransmitStreamServer) error
	mustEmbedUnimplementedTransmitterServiceServer()
}

// UnimplementedTransmitterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTransmitterServiceServer struct {
}

func (UnimplementedTransmitterServiceServer) TransmitStream(*empty.Empty, TransmitterService_TransmitStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method TransmitStream not implemented")
}
func (UnimplementedTransmitterServiceServer) mustEmbedUnimplementedTransmitterServiceServer() {}

// UnsafeTransmitterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransmitterServiceServer will
// result in compilation errors.
type UnsafeTransmitterServiceServer interface {
	mustEmbedUnimplementedTransmitterServiceServer()
}

func RegisterTransmitterServiceServer(s grpc.ServiceRegistrar, srv TransmitterServiceServer) {
	s.RegisterService(&TransmitterService_ServiceDesc, srv)
}

func _TransmitterService_TransmitStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TransmitterServiceServer).TransmitStream(m, &transmitterServiceTransmitStreamServer{stream})
}

type TransmitterService_TransmitStreamServer interface {
	Send(*Transmission) error
	grpc.ServerStream
}

type transmitterServiceTransmitStreamServer struct {
	grpc.ServerStream
}

func (x *transmitterServiceTransmitStreamServer) Send(m *Transmission) error {
	return x.ServerStream.SendMsg(m)
}

// TransmitterService_ServiceDesc is the grpc.ServiceDesc for TransmitterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransmitterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transmitter.TransmitterService",
	HandlerType: (*TransmitterServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TransmitStream",
			Handler:       _TransmitterService_TransmitStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "transmitter.proto",
}
