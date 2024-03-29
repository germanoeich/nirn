// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package service

import (
	context "context"
	cache "github.com/germanoeich/nirn/proto/discord/v1/cache"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GatewayCacheClient is the client API for GatewayCache service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayCacheClient interface {
	// Guilds
	GetGuild(ctx context.Context, in *cache.GetGuildRequest, opts ...grpc.CallOption) (*cache.GetGuildResponse, error)
	// Channels
	ListGuildChannels(ctx context.Context, in *cache.ListGuildChannelsRequest, opts ...grpc.CallOption) (*cache.ListGuildChannelsResponse, error)
	GetGuildChannel(ctx context.Context, in *cache.GetGuildChannelRequest, opts ...grpc.CallOption) (*cache.GetGuildChannelResponse, error)
	// Guild Members
	ListGuildMembers(ctx context.Context, in *cache.ListGuildMembersRequest, opts ...grpc.CallOption) (*cache.ListGuildMembersResponse, error)
	GetGuildMember(ctx context.Context, in *cache.GetGuildMemberRequest, opts ...grpc.CallOption) (*cache.GetGuildMemberResponse, error)
	FindGuildMembers(ctx context.Context, in *cache.FindGuildMembersRequest, opts ...grpc.CallOption) (*cache.FindGuildMembersResponse, error)
	// Guild Member Presence
	GetGuildMemberPresence(ctx context.Context, in *cache.GetGuildMemberPresenceRequest, opts ...grpc.CallOption) (*cache.GetGuildMemberPresenceResponse, error)
	// Guild Member Roles
	ListGuildRoles(ctx context.Context, in *cache.ListGuildRolesRequest, opts ...grpc.CallOption) (*cache.ListGuildRolesResponse, error)
	GetGuildRole(ctx context.Context, in *cache.GetGuildRoleRequest, opts ...grpc.CallOption) (*cache.GetGuildRoleResponse, error)
	// Emojis
	ListGuildEmojis(ctx context.Context, in *cache.ListGuildEmojisRequest, opts ...grpc.CallOption) (*cache.ListGuildEmojisResponse, error)
	GetGuildEmoji(ctx context.Context, in *cache.GetGuildEmojiRequest, opts ...grpc.CallOption) (*cache.GetGuildEmojiResponse, error)
	// VoiceStates
	GetGuildMemberVoiceState(ctx context.Context, in *cache.GetGuildMemberVoiceStateRequest, opts ...grpc.CallOption) (*cache.GetGuildMemberVoiceStateResponse, error)
	ListGuildChannelVoiceStates(ctx context.Context, in *cache.ListGuildChannelVoiceStatesRequest, opts ...grpc.CallOption) (*cache.ListGuildChannelVoiceStatesResponse, error)
	// GetUser
	GetUser(ctx context.Context, in *cache.GetUserRequest, opts ...grpc.CallOption) (*cache.GetUserResponse, error)
}

type gatewayCacheClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayCacheClient(cc grpc.ClientConnInterface) GatewayCacheClient {
	return &gatewayCacheClient{cc}
}

func (c *gatewayCacheClient) GetGuild(ctx context.Context, in *cache.GetGuildRequest, opts ...grpc.CallOption) (*cache.GetGuildResponse, error) {
	out := new(cache.GetGuildResponse)
	err := c.cc.Invoke(ctx, "/pylon.gateway.v1.service.GatewayCache/GetGuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayCacheClient) ListGuildChannels(ctx context.Context, in *cache.ListGuildChannelsRequest, opts ...grpc.CallOption) (*cache.ListGuildChannelsResponse, error) {
	out := new(cache.ListGuildChannelsResponse)
	err := c.cc.Invoke(ctx, "/pylon.gateway.v1.service.GatewayCache/ListGuildChannels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayCacheClient) GetGuildChannel(ctx context.Context, in *cache.GetGuildChannelRequest, opts ...grpc.CallOption) (*cache.GetGuildChannelResponse, error) {
	out := new(cache.GetGuildChannelResponse)
	err := c.cc.Invoke(ctx, "/pylon.gateway.v1.service.GatewayCache/GetGuildChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayCacheClient) ListGuildMembers(ctx context.Context, in *cache.ListGuildMembersRequest, opts ...grpc.CallOption) (*cache.ListGuildMembersResponse, error) {
	out := new(cache.ListGuildMembersResponse)
	err := c.cc.Invoke(ctx, "/pylon.gateway.v1.service.GatewayCache/ListGuildMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayCacheClient) GetGuildMember(ctx context.Context, in *cache.GetGuildMemberRequest, opts ...grpc.CallOption) (*cache.GetGuildMemberResponse, error) {
	out := new(cache.GetGuildMemberResponse)
	err := c.cc.Invoke(ctx, "/pylon.gateway.v1.service.GatewayCache/GetGuildMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayCacheClient) FindGuildMembers(ctx context.Context, in *cache.FindGuildMembersRequest, opts ...grpc.CallOption) (*cache.FindGuildMembersResponse, error) {
	out := new(cache.FindGuildMembersResponse)
	err := c.cc.Invoke(ctx, "/pylon.gateway.v1.service.GatewayCache/FindGuildMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayCacheClient) GetGuildMemberPresence(ctx context.Context, in *cache.GetGuildMemberPresenceRequest, opts ...grpc.CallOption) (*cache.GetGuildMemberPresenceResponse, error) {
	out := new(cache.GetGuildMemberPresenceResponse)
	err := c.cc.Invoke(ctx, "/pylon.gateway.v1.service.GatewayCache/GetGuildMemberPresence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayCacheClient) ListGuildRoles(ctx context.Context, in *cache.ListGuildRolesRequest, opts ...grpc.CallOption) (*cache.ListGuildRolesResponse, error) {
	out := new(cache.ListGuildRolesResponse)
	err := c.cc.Invoke(ctx, "/pylon.gateway.v1.service.GatewayCache/ListGuildRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayCacheClient) GetGuildRole(ctx context.Context, in *cache.GetGuildRoleRequest, opts ...grpc.CallOption) (*cache.GetGuildRoleResponse, error) {
	out := new(cache.GetGuildRoleResponse)
	err := c.cc.Invoke(ctx, "/pylon.gateway.v1.service.GatewayCache/GetGuildRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayCacheClient) ListGuildEmojis(ctx context.Context, in *cache.ListGuildEmojisRequest, opts ...grpc.CallOption) (*cache.ListGuildEmojisResponse, error) {
	out := new(cache.ListGuildEmojisResponse)
	err := c.cc.Invoke(ctx, "/pylon.gateway.v1.service.GatewayCache/ListGuildEmojis", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayCacheClient) GetGuildEmoji(ctx context.Context, in *cache.GetGuildEmojiRequest, opts ...grpc.CallOption) (*cache.GetGuildEmojiResponse, error) {
	out := new(cache.GetGuildEmojiResponse)
	err := c.cc.Invoke(ctx, "/pylon.gateway.v1.service.GatewayCache/GetGuildEmoji", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayCacheClient) GetGuildMemberVoiceState(ctx context.Context, in *cache.GetGuildMemberVoiceStateRequest, opts ...grpc.CallOption) (*cache.GetGuildMemberVoiceStateResponse, error) {
	out := new(cache.GetGuildMemberVoiceStateResponse)
	err := c.cc.Invoke(ctx, "/pylon.gateway.v1.service.GatewayCache/GetGuildMemberVoiceState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayCacheClient) ListGuildChannelVoiceStates(ctx context.Context, in *cache.ListGuildChannelVoiceStatesRequest, opts ...grpc.CallOption) (*cache.ListGuildChannelVoiceStatesResponse, error) {
	out := new(cache.ListGuildChannelVoiceStatesResponse)
	err := c.cc.Invoke(ctx, "/pylon.gateway.v1.service.GatewayCache/ListGuildChannelVoiceStates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayCacheClient) GetUser(ctx context.Context, in *cache.GetUserRequest, opts ...grpc.CallOption) (*cache.GetUserResponse, error) {
	out := new(cache.GetUserResponse)
	err := c.cc.Invoke(ctx, "/pylon.gateway.v1.service.GatewayCache/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayCacheServer is the server API for GatewayCache service.
// All implementations must embed UnimplementedGatewayCacheServer
// for forward compatibility
type GatewayCacheServer interface {
	// Guilds
	GetGuild(context.Context, *cache.GetGuildRequest) (*cache.GetGuildResponse, error)
	// Channels
	ListGuildChannels(context.Context, *cache.ListGuildChannelsRequest) (*cache.ListGuildChannelsResponse, error)
	GetGuildChannel(context.Context, *cache.GetGuildChannelRequest) (*cache.GetGuildChannelResponse, error)
	// Guild Members
	ListGuildMembers(context.Context, *cache.ListGuildMembersRequest) (*cache.ListGuildMembersResponse, error)
	GetGuildMember(context.Context, *cache.GetGuildMemberRequest) (*cache.GetGuildMemberResponse, error)
	FindGuildMembers(context.Context, *cache.FindGuildMembersRequest) (*cache.FindGuildMembersResponse, error)
	// Guild Member Presence
	GetGuildMemberPresence(context.Context, *cache.GetGuildMemberPresenceRequest) (*cache.GetGuildMemberPresenceResponse, error)
	// Guild Member Roles
	ListGuildRoles(context.Context, *cache.ListGuildRolesRequest) (*cache.ListGuildRolesResponse, error)
	GetGuildRole(context.Context, *cache.GetGuildRoleRequest) (*cache.GetGuildRoleResponse, error)
	// Emojis
	ListGuildEmojis(context.Context, *cache.ListGuildEmojisRequest) (*cache.ListGuildEmojisResponse, error)
	GetGuildEmoji(context.Context, *cache.GetGuildEmojiRequest) (*cache.GetGuildEmojiResponse, error)
	// VoiceStates
	GetGuildMemberVoiceState(context.Context, *cache.GetGuildMemberVoiceStateRequest) (*cache.GetGuildMemberVoiceStateResponse, error)
	ListGuildChannelVoiceStates(context.Context, *cache.ListGuildChannelVoiceStatesRequest) (*cache.ListGuildChannelVoiceStatesResponse, error)
	// GetUser
	GetUser(context.Context, *cache.GetUserRequest) (*cache.GetUserResponse, error)
	mustEmbedUnimplementedGatewayCacheServer()
}

// UnimplementedGatewayCacheServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayCacheServer struct {
}

func (UnimplementedGatewayCacheServer) GetGuild(context.Context, *cache.GetGuildRequest) (*cache.GetGuildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGuild not implemented")
}
func (UnimplementedGatewayCacheServer) ListGuildChannels(context.Context, *cache.ListGuildChannelsRequest) (*cache.ListGuildChannelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGuildChannels not implemented")
}
func (UnimplementedGatewayCacheServer) GetGuildChannel(context.Context, *cache.GetGuildChannelRequest) (*cache.GetGuildChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGuildChannel not implemented")
}
func (UnimplementedGatewayCacheServer) ListGuildMembers(context.Context, *cache.ListGuildMembersRequest) (*cache.ListGuildMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGuildMembers not implemented")
}
func (UnimplementedGatewayCacheServer) GetGuildMember(context.Context, *cache.GetGuildMemberRequest) (*cache.GetGuildMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGuildMember not implemented")
}
func (UnimplementedGatewayCacheServer) FindGuildMembers(context.Context, *cache.FindGuildMembersRequest) (*cache.FindGuildMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindGuildMembers not implemented")
}
func (UnimplementedGatewayCacheServer) GetGuildMemberPresence(context.Context, *cache.GetGuildMemberPresenceRequest) (*cache.GetGuildMemberPresenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGuildMemberPresence not implemented")
}
func (UnimplementedGatewayCacheServer) ListGuildRoles(context.Context, *cache.ListGuildRolesRequest) (*cache.ListGuildRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGuildRoles not implemented")
}
func (UnimplementedGatewayCacheServer) GetGuildRole(context.Context, *cache.GetGuildRoleRequest) (*cache.GetGuildRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGuildRole not implemented")
}
func (UnimplementedGatewayCacheServer) ListGuildEmojis(context.Context, *cache.ListGuildEmojisRequest) (*cache.ListGuildEmojisResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGuildEmojis not implemented")
}
func (UnimplementedGatewayCacheServer) GetGuildEmoji(context.Context, *cache.GetGuildEmojiRequest) (*cache.GetGuildEmojiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGuildEmoji not implemented")
}
func (UnimplementedGatewayCacheServer) GetGuildMemberVoiceState(context.Context, *cache.GetGuildMemberVoiceStateRequest) (*cache.GetGuildMemberVoiceStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGuildMemberVoiceState not implemented")
}
func (UnimplementedGatewayCacheServer) ListGuildChannelVoiceStates(context.Context, *cache.ListGuildChannelVoiceStatesRequest) (*cache.ListGuildChannelVoiceStatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGuildChannelVoiceStates not implemented")
}
func (UnimplementedGatewayCacheServer) GetUser(context.Context, *cache.GetUserRequest) (*cache.GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedGatewayCacheServer) mustEmbedUnimplementedGatewayCacheServer() {}

// UnsafeGatewayCacheServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayCacheServer will
// result in compilation errors.
type UnsafeGatewayCacheServer interface {
	mustEmbedUnimplementedGatewayCacheServer()
}

func RegisterGatewayCacheServer(s grpc.ServiceRegistrar, srv GatewayCacheServer) {
	s.RegisterService(&GatewayCache_ServiceDesc, srv)
}

func _GatewayCache_GetGuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(cache.GetGuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayCacheServer).GetGuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pylon.gateway.v1.service.GatewayCache/GetGuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayCacheServer).GetGuild(ctx, req.(*cache.GetGuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayCache_ListGuildChannels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(cache.ListGuildChannelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayCacheServer).ListGuildChannels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pylon.gateway.v1.service.GatewayCache/ListGuildChannels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayCacheServer).ListGuildChannels(ctx, req.(*cache.ListGuildChannelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayCache_GetGuildChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(cache.GetGuildChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayCacheServer).GetGuildChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pylon.gateway.v1.service.GatewayCache/GetGuildChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayCacheServer).GetGuildChannel(ctx, req.(*cache.GetGuildChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayCache_ListGuildMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(cache.ListGuildMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayCacheServer).ListGuildMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pylon.gateway.v1.service.GatewayCache/ListGuildMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayCacheServer).ListGuildMembers(ctx, req.(*cache.ListGuildMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayCache_GetGuildMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(cache.GetGuildMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayCacheServer).GetGuildMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pylon.gateway.v1.service.GatewayCache/GetGuildMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayCacheServer).GetGuildMember(ctx, req.(*cache.GetGuildMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayCache_FindGuildMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(cache.FindGuildMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayCacheServer).FindGuildMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pylon.gateway.v1.service.GatewayCache/FindGuildMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayCacheServer).FindGuildMembers(ctx, req.(*cache.FindGuildMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayCache_GetGuildMemberPresence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(cache.GetGuildMemberPresenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayCacheServer).GetGuildMemberPresence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pylon.gateway.v1.service.GatewayCache/GetGuildMemberPresence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayCacheServer).GetGuildMemberPresence(ctx, req.(*cache.GetGuildMemberPresenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayCache_ListGuildRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(cache.ListGuildRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayCacheServer).ListGuildRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pylon.gateway.v1.service.GatewayCache/ListGuildRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayCacheServer).ListGuildRoles(ctx, req.(*cache.ListGuildRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayCache_GetGuildRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(cache.GetGuildRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayCacheServer).GetGuildRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pylon.gateway.v1.service.GatewayCache/GetGuildRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayCacheServer).GetGuildRole(ctx, req.(*cache.GetGuildRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayCache_ListGuildEmojis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(cache.ListGuildEmojisRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayCacheServer).ListGuildEmojis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pylon.gateway.v1.service.GatewayCache/ListGuildEmojis",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayCacheServer).ListGuildEmojis(ctx, req.(*cache.ListGuildEmojisRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayCache_GetGuildEmoji_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(cache.GetGuildEmojiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayCacheServer).GetGuildEmoji(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pylon.gateway.v1.service.GatewayCache/GetGuildEmoji",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayCacheServer).GetGuildEmoji(ctx, req.(*cache.GetGuildEmojiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayCache_GetGuildMemberVoiceState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(cache.GetGuildMemberVoiceStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayCacheServer).GetGuildMemberVoiceState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pylon.gateway.v1.service.GatewayCache/GetGuildMemberVoiceState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayCacheServer).GetGuildMemberVoiceState(ctx, req.(*cache.GetGuildMemberVoiceStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayCache_ListGuildChannelVoiceStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(cache.ListGuildChannelVoiceStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayCacheServer).ListGuildChannelVoiceStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pylon.gateway.v1.service.GatewayCache/ListGuildChannelVoiceStates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayCacheServer).ListGuildChannelVoiceStates(ctx, req.(*cache.ListGuildChannelVoiceStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayCache_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(cache.GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayCacheServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pylon.gateway.v1.service.GatewayCache/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayCacheServer).GetUser(ctx, req.(*cache.GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GatewayCache_ServiceDesc is the grpc.ServiceDesc for GatewayCache service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GatewayCache_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pylon.gateway.v1.service.GatewayCache",
	HandlerType: (*GatewayCacheServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGuild",
			Handler:    _GatewayCache_GetGuild_Handler,
		},
		{
			MethodName: "ListGuildChannels",
			Handler:    _GatewayCache_ListGuildChannels_Handler,
		},
		{
			MethodName: "GetGuildChannel",
			Handler:    _GatewayCache_GetGuildChannel_Handler,
		},
		{
			MethodName: "ListGuildMembers",
			Handler:    _GatewayCache_ListGuildMembers_Handler,
		},
		{
			MethodName: "GetGuildMember",
			Handler:    _GatewayCache_GetGuildMember_Handler,
		},
		{
			MethodName: "FindGuildMembers",
			Handler:    _GatewayCache_FindGuildMembers_Handler,
		},
		{
			MethodName: "GetGuildMemberPresence",
			Handler:    _GatewayCache_GetGuildMemberPresence_Handler,
		},
		{
			MethodName: "ListGuildRoles",
			Handler:    _GatewayCache_ListGuildRoles_Handler,
		},
		{
			MethodName: "GetGuildRole",
			Handler:    _GatewayCache_GetGuildRole_Handler,
		},
		{
			MethodName: "ListGuildEmojis",
			Handler:    _GatewayCache_ListGuildEmojis_Handler,
		},
		{
			MethodName: "GetGuildEmoji",
			Handler:    _GatewayCache_GetGuildEmoji_Handler,
		},
		{
			MethodName: "GetGuildMemberVoiceState",
			Handler:    _GatewayCache_GetGuildMemberVoiceState_Handler,
		},
		{
			MethodName: "ListGuildChannelVoiceStates",
			Handler:    _GatewayCache_ListGuildChannelVoiceStates_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _GatewayCache_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway/v1/cache_service.proto",
}
