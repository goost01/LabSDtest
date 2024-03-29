// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package helloworld

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

// GameClient is the client API for Game service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameClient interface {
	// Sends Join game
	JoinGame(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinReply, error)
	BeginGame(ctx context.Context, in *BeginRequest, opts ...grpc.CallOption) (*BeginReply, error)
	// Sends end game
	EndGame(ctx context.Context, in *EndRequest, opts ...grpc.CallOption) (*EndReply, error)
	// Begin Stage
	BeginStage(ctx context.Context, in *BeginStageRequest, opts ...grpc.CallOption) (*BeginStageReply, error)
	BeginRound(ctx context.Context, in *PingMsg, opts ...grpc.CallOption) (*BeginRoundReply, error)
	SendJugadaE1(ctx context.Context, in *JugadaE1, opts ...grpc.CallOption) (*PlayerStatusE1, error)
	SendJugadaE2(ctx context.Context, in *JugadaE2, opts ...grpc.CallOption) (*PingMsg, error)
	IsAlready(ctx context.Context, in *PingMsg, opts ...grpc.CallOption) (*PlayerStatusE2, error)
}

type gameClient struct {
	cc grpc.ClientConnInterface
}

func NewGameClient(cc grpc.ClientConnInterface) GameClient {
	return &gameClient{cc}
}

func (c *gameClient) JoinGame(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinReply, error) {
	out := new(JoinReply)
	err := c.cc.Invoke(ctx, "/helloworld.Game/JoinGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) BeginGame(ctx context.Context, in *BeginRequest, opts ...grpc.CallOption) (*BeginReply, error) {
	out := new(BeginReply)
	err := c.cc.Invoke(ctx, "/helloworld.Game/BeginGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) EndGame(ctx context.Context, in *EndRequest, opts ...grpc.CallOption) (*EndReply, error) {
	out := new(EndReply)
	err := c.cc.Invoke(ctx, "/helloworld.Game/EndGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) BeginStage(ctx context.Context, in *BeginStageRequest, opts ...grpc.CallOption) (*BeginStageReply, error) {
	out := new(BeginStageReply)
	err := c.cc.Invoke(ctx, "/helloworld.Game/BeginStage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) BeginRound(ctx context.Context, in *PingMsg, opts ...grpc.CallOption) (*BeginRoundReply, error) {
	out := new(BeginRoundReply)
	err := c.cc.Invoke(ctx, "/helloworld.Game/BeginRound", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) SendJugadaE1(ctx context.Context, in *JugadaE1, opts ...grpc.CallOption) (*PlayerStatusE1, error) {
	out := new(PlayerStatusE1)
	err := c.cc.Invoke(ctx, "/helloworld.Game/SendJugadaE1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) SendJugadaE2(ctx context.Context, in *JugadaE2, opts ...grpc.CallOption) (*PingMsg, error) {
	out := new(PingMsg)
	err := c.cc.Invoke(ctx, "/helloworld.Game/SendJugadaE2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) IsAlready(ctx context.Context, in *PingMsg, opts ...grpc.CallOption) (*PlayerStatusE2, error) {
	out := new(PlayerStatusE2)
	err := c.cc.Invoke(ctx, "/helloworld.Game/IsAlready", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServer is the server API for Game service.
// All implementations must embed UnimplementedGameServer
// for forward compatibility
type GameServer interface {
	// Sends Join game
	JoinGame(context.Context, *JoinRequest) (*JoinReply, error)
	BeginGame(context.Context, *BeginRequest) (*BeginReply, error)
	// Sends end game
	EndGame(context.Context, *EndRequest) (*EndReply, error)
	// Begin Stage
	BeginStage(context.Context, *BeginStageRequest) (*BeginStageReply, error)
	BeginRound(context.Context, *PingMsg) (*BeginRoundReply, error)
	SendJugadaE1(context.Context, *JugadaE1) (*PlayerStatusE1, error)
	SendJugadaE2(context.Context, *JugadaE2) (*PingMsg, error)
	IsAlready(context.Context, *PingMsg) (*PlayerStatusE2, error)
	mustEmbedUnimplementedGameServer()
}

// UnimplementedGameServer must be embedded to have forward compatible implementations.
type UnimplementedGameServer struct {
}

func (UnimplementedGameServer) JoinGame(context.Context, *JoinRequest) (*JoinReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinGame not implemented")
}
func (UnimplementedGameServer) BeginGame(context.Context, *BeginRequest) (*BeginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BeginGame not implemented")
}
func (UnimplementedGameServer) EndGame(context.Context, *EndRequest) (*EndReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndGame not implemented")
}
func (UnimplementedGameServer) BeginStage(context.Context, *BeginStageRequest) (*BeginStageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BeginStage not implemented")
}
func (UnimplementedGameServer) BeginRound(context.Context, *PingMsg) (*BeginRoundReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BeginRound not implemented")
}
func (UnimplementedGameServer) SendJugadaE1(context.Context, *JugadaE1) (*PlayerStatusE1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendJugadaE1 not implemented")
}
func (UnimplementedGameServer) SendJugadaE2(context.Context, *JugadaE2) (*PingMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendJugadaE2 not implemented")
}
func (UnimplementedGameServer) IsAlready(context.Context, *PingMsg) (*PlayerStatusE2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAlready not implemented")
}
func (UnimplementedGameServer) mustEmbedUnimplementedGameServer() {}

// UnsafeGameServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameServer will
// result in compilation errors.
type UnsafeGameServer interface {
	mustEmbedUnimplementedGameServer()
}

func RegisterGameServer(s grpc.ServiceRegistrar, srv GameServer) {
	s.RegisterService(&Game_ServiceDesc, srv)
}

func _Game_JoinGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).JoinGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Game/JoinGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).JoinGame(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_BeginGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BeginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).BeginGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Game/BeginGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).BeginGame(ctx, req.(*BeginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_EndGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).EndGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Game/EndGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).EndGame(ctx, req.(*EndRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_BeginStage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BeginStageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).BeginStage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Game/BeginStage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).BeginStage(ctx, req.(*BeginStageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_BeginRound_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).BeginRound(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Game/BeginRound",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).BeginRound(ctx, req.(*PingMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_SendJugadaE1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JugadaE1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).SendJugadaE1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Game/SendJugadaE1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).SendJugadaE1(ctx, req.(*JugadaE1))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_SendJugadaE2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JugadaE2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).SendJugadaE2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Game/SendJugadaE2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).SendJugadaE2(ctx, req.(*JugadaE2))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_IsAlready_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).IsAlready(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Game/IsAlready",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).IsAlready(ctx, req.(*PingMsg))
	}
	return interceptor(ctx, in, info, handler)
}

// Game_ServiceDesc is the grpc.ServiceDesc for Game service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Game_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Game",
	HandlerType: (*GameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinGame",
			Handler:    _Game_JoinGame_Handler,
		},
		{
			MethodName: "BeginGame",
			Handler:    _Game_BeginGame_Handler,
		},
		{
			MethodName: "EndGame",
			Handler:    _Game_EndGame_Handler,
		},
		{
			MethodName: "BeginStage",
			Handler:    _Game_BeginStage_Handler,
		},
		{
			MethodName: "BeginRound",
			Handler:    _Game_BeginRound_Handler,
		},
		{
			MethodName: "SendJugadaE1",
			Handler:    _Game_SendJugadaE1_Handler,
		},
		{
			MethodName: "SendJugadaE2",
			Handler:    _Game_SendJugadaE2_Handler,
		},
		{
			MethodName: "IsAlready",
			Handler:    _Game_IsAlready_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld/helloworld.proto",
}
