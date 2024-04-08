// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: plugin.proto

package libplugin

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

// SshPiperPluginClient is the client API for SshPiperPlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SshPiperPluginClient interface {
	Logs(ctx context.Context, in *StartLogRequest, opts ...grpc.CallOption) (SshPiperPlugin_LogsClient, error)
	ListCallbacks(ctx context.Context, in *ListCallbackRequest, opts ...grpc.CallOption) (*ListCallbackResponse, error)
	NewConnection(ctx context.Context, in *NewConnectionRequest, opts ...grpc.CallOption) (*NewConnectionResponse, error)
	NextAuthMethods(ctx context.Context, in *NextAuthMethodsRequest, opts ...grpc.CallOption) (*NextAuthMethodsResponse, error)
	NoneAuth(ctx context.Context, in *NoneAuthRequest, opts ...grpc.CallOption) (*NoneAuthResponse, error)
	PasswordAuth(ctx context.Context, in *PasswordAuthRequest, opts ...grpc.CallOption) (*PasswordAuthResponse, error)
	PublicKeyAuth(ctx context.Context, in *PublicKeyAuthRequest, opts ...grpc.CallOption) (*PublicKeyAuthResponse, error)
	KeyboardInteractiveAuth(ctx context.Context, opts ...grpc.CallOption) (SshPiperPlugin_KeyboardInteractiveAuthClient, error)
	UpstreamAuthFailureNotice(ctx context.Context, in *UpstreamAuthFailureNoticeRequest, opts ...grpc.CallOption) (*UpstreamAuthFailureNoticeResponse, error)
	Banner(ctx context.Context, in *BannerRequest, opts ...grpc.CallOption) (*BannerResponse, error)
	VerifyHostKey(ctx context.Context, in *VerifyHostKeyRequest, opts ...grpc.CallOption) (*VerifyHostKeyResponse, error)
	PipeCreateErrorNotice(ctx context.Context, in *PipeCreateErrorNoticeRequest, opts ...grpc.CallOption) (*PipeCreateErrorNoticeResponse, error)
	PipeStartNotice(ctx context.Context, in *PipeStartNoticeRequest, opts ...grpc.CallOption) (*PipeStartNoticeResponse, error)
	PipeErrorNotice(ctx context.Context, in *PipeErrorNoticeRequest, opts ...grpc.CallOption) (*PipeErrorNoticeResponse, error)
}

type sshPiperPluginClient struct {
	cc grpc.ClientConnInterface
}

func NewSshPiperPluginClient(cc grpc.ClientConnInterface) SshPiperPluginClient {
	return &sshPiperPluginClient{cc}
}

func (c *sshPiperPluginClient) Logs(ctx context.Context, in *StartLogRequest, opts ...grpc.CallOption) (SshPiperPlugin_LogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &SshPiperPlugin_ServiceDesc.Streams[0], "/libplugin.SshPiperPlugin/Logs", opts...)
	if err != nil {
		return nil, err
	}
	x := &sshPiperPluginLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SshPiperPlugin_LogsClient interface {
	Recv() (*Log, error)
	grpc.ClientStream
}

type sshPiperPluginLogsClient struct {
	grpc.ClientStream
}

func (x *sshPiperPluginLogsClient) Recv() (*Log, error) {
	m := new(Log)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sshPiperPluginClient) ListCallbacks(ctx context.Context, in *ListCallbackRequest, opts ...grpc.CallOption) (*ListCallbackResponse, error) {
	out := new(ListCallbackResponse)
	err := c.cc.Invoke(ctx, "/libplugin.SshPiperPlugin/ListCallbacks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sshPiperPluginClient) NewConnection(ctx context.Context, in *NewConnectionRequest, opts ...grpc.CallOption) (*NewConnectionResponse, error) {
	out := new(NewConnectionResponse)
	err := c.cc.Invoke(ctx, "/libplugin.SshPiperPlugin/NewConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sshPiperPluginClient) NextAuthMethods(ctx context.Context, in *NextAuthMethodsRequest, opts ...grpc.CallOption) (*NextAuthMethodsResponse, error) {
	out := new(NextAuthMethodsResponse)
	err := c.cc.Invoke(ctx, "/libplugin.SshPiperPlugin/NextAuthMethods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sshPiperPluginClient) NoneAuth(ctx context.Context, in *NoneAuthRequest, opts ...grpc.CallOption) (*NoneAuthResponse, error) {
	out := new(NoneAuthResponse)
	err := c.cc.Invoke(ctx, "/libplugin.SshPiperPlugin/NoneAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sshPiperPluginClient) PasswordAuth(ctx context.Context, in *PasswordAuthRequest, opts ...grpc.CallOption) (*PasswordAuthResponse, error) {
	out := new(PasswordAuthResponse)
	err := c.cc.Invoke(ctx, "/libplugin.SshPiperPlugin/PasswordAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sshPiperPluginClient) PublicKeyAuth(ctx context.Context, in *PublicKeyAuthRequest, opts ...grpc.CallOption) (*PublicKeyAuthResponse, error) {
	out := new(PublicKeyAuthResponse)
	err := c.cc.Invoke(ctx, "/libplugin.SshPiperPlugin/PublicKeyAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sshPiperPluginClient) KeyboardInteractiveAuth(ctx context.Context, opts ...grpc.CallOption) (SshPiperPlugin_KeyboardInteractiveAuthClient, error) {
	stream, err := c.cc.NewStream(ctx, &SshPiperPlugin_ServiceDesc.Streams[1], "/libplugin.SshPiperPlugin/KeyboardInteractiveAuth", opts...)
	if err != nil {
		return nil, err
	}
	x := &sshPiperPluginKeyboardInteractiveAuthClient{stream}
	return x, nil
}

type SshPiperPlugin_KeyboardInteractiveAuthClient interface {
	Send(*KeyboardInteractiveAuthMessage) error
	Recv() (*KeyboardInteractiveAuthMessage, error)
	grpc.ClientStream
}

type sshPiperPluginKeyboardInteractiveAuthClient struct {
	grpc.ClientStream
}

func (x *sshPiperPluginKeyboardInteractiveAuthClient) Send(m *KeyboardInteractiveAuthMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sshPiperPluginKeyboardInteractiveAuthClient) Recv() (*KeyboardInteractiveAuthMessage, error) {
	m := new(KeyboardInteractiveAuthMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sshPiperPluginClient) UpstreamAuthFailureNotice(ctx context.Context, in *UpstreamAuthFailureNoticeRequest, opts ...grpc.CallOption) (*UpstreamAuthFailureNoticeResponse, error) {
	out := new(UpstreamAuthFailureNoticeResponse)
	err := c.cc.Invoke(ctx, "/libplugin.SshPiperPlugin/UpstreamAuthFailureNotice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sshPiperPluginClient) Banner(ctx context.Context, in *BannerRequest, opts ...grpc.CallOption) (*BannerResponse, error) {
	out := new(BannerResponse)
	err := c.cc.Invoke(ctx, "/libplugin.SshPiperPlugin/Banner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sshPiperPluginClient) VerifyHostKey(ctx context.Context, in *VerifyHostKeyRequest, opts ...grpc.CallOption) (*VerifyHostKeyResponse, error) {
	out := new(VerifyHostKeyResponse)
	err := c.cc.Invoke(ctx, "/libplugin.SshPiperPlugin/VerifyHostKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sshPiperPluginClient) PipeCreateErrorNotice(ctx context.Context, in *PipeCreateErrorNoticeRequest, opts ...grpc.CallOption) (*PipeCreateErrorNoticeResponse, error) {
	out := new(PipeCreateErrorNoticeResponse)
	err := c.cc.Invoke(ctx, "/libplugin.SshPiperPlugin/PipeCreateErrorNotice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sshPiperPluginClient) PipeStartNotice(ctx context.Context, in *PipeStartNoticeRequest, opts ...grpc.CallOption) (*PipeStartNoticeResponse, error) {
	out := new(PipeStartNoticeResponse)
	err := c.cc.Invoke(ctx, "/libplugin.SshPiperPlugin/PipeStartNotice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sshPiperPluginClient) PipeErrorNotice(ctx context.Context, in *PipeErrorNoticeRequest, opts ...grpc.CallOption) (*PipeErrorNoticeResponse, error) {
	out := new(PipeErrorNoticeResponse)
	err := c.cc.Invoke(ctx, "/libplugin.SshPiperPlugin/PipeErrorNotice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SshPiperPluginServer is the server API for SshPiperPlugin service.
// All implementations must embed UnimplementedSshPiperPluginServer
// for forward compatibility
type SshPiperPluginServer interface {
	Logs(*StartLogRequest, SshPiperPlugin_LogsServer) error
	ListCallbacks(context.Context, *ListCallbackRequest) (*ListCallbackResponse, error)
	NewConnection(context.Context, *NewConnectionRequest) (*NewConnectionResponse, error)
	NextAuthMethods(context.Context, *NextAuthMethodsRequest) (*NextAuthMethodsResponse, error)
	NoneAuth(context.Context, *NoneAuthRequest) (*NoneAuthResponse, error)
	PasswordAuth(context.Context, *PasswordAuthRequest) (*PasswordAuthResponse, error)
	PublicKeyAuth(context.Context, *PublicKeyAuthRequest) (*PublicKeyAuthResponse, error)
	KeyboardInteractiveAuth(SshPiperPlugin_KeyboardInteractiveAuthServer) error
	UpstreamAuthFailureNotice(context.Context, *UpstreamAuthFailureNoticeRequest) (*UpstreamAuthFailureNoticeResponse, error)
	Banner(context.Context, *BannerRequest) (*BannerResponse, error)
	VerifyHostKey(context.Context, *VerifyHostKeyRequest) (*VerifyHostKeyResponse, error)
	PipeCreateErrorNotice(context.Context, *PipeCreateErrorNoticeRequest) (*PipeCreateErrorNoticeResponse, error)
	PipeStartNotice(context.Context, *PipeStartNoticeRequest) (*PipeStartNoticeResponse, error)
	PipeErrorNotice(context.Context, *PipeErrorNoticeRequest) (*PipeErrorNoticeResponse, error)
	mustEmbedUnimplementedSshPiperPluginServer()
}

// UnimplementedSshPiperPluginServer must be embedded to have forward compatible implementations.
type UnimplementedSshPiperPluginServer struct {
}

func (UnimplementedSshPiperPluginServer) Logs(*StartLogRequest, SshPiperPlugin_LogsServer) error {
	return status.Errorf(codes.Unimplemented, "method Logs not implemented")
}
func (UnimplementedSshPiperPluginServer) ListCallbacks(context.Context, *ListCallbackRequest) (*ListCallbackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCallbacks not implemented")
}
func (UnimplementedSshPiperPluginServer) NewConnection(context.Context, *NewConnectionRequest) (*NewConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewConnection not implemented")
}
func (UnimplementedSshPiperPluginServer) NextAuthMethods(context.Context, *NextAuthMethodsRequest) (*NextAuthMethodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NextAuthMethods not implemented")
}
func (UnimplementedSshPiperPluginServer) NoneAuth(context.Context, *NoneAuthRequest) (*NoneAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NoneAuth not implemented")
}
func (UnimplementedSshPiperPluginServer) PasswordAuth(context.Context, *PasswordAuthRequest) (*PasswordAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PasswordAuth not implemented")
}
func (UnimplementedSshPiperPluginServer) PublicKeyAuth(context.Context, *PublicKeyAuthRequest) (*PublicKeyAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublicKeyAuth not implemented")
}
func (UnimplementedSshPiperPluginServer) KeyboardInteractiveAuth(SshPiperPlugin_KeyboardInteractiveAuthServer) error {
	return status.Errorf(codes.Unimplemented, "method KeyboardInteractiveAuth not implemented")
}
func (UnimplementedSshPiperPluginServer) UpstreamAuthFailureNotice(context.Context, *UpstreamAuthFailureNoticeRequest) (*UpstreamAuthFailureNoticeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpstreamAuthFailureNotice not implemented")
}
func (UnimplementedSshPiperPluginServer) Banner(context.Context, *BannerRequest) (*BannerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Banner not implemented")
}
func (UnimplementedSshPiperPluginServer) VerifyHostKey(context.Context, *VerifyHostKeyRequest) (*VerifyHostKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyHostKey not implemented")
}
func (UnimplementedSshPiperPluginServer) PipeCreateErrorNotice(context.Context, *PipeCreateErrorNoticeRequest) (*PipeCreateErrorNoticeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PipeCreateErrorNotice not implemented")
}
func (UnimplementedSshPiperPluginServer) PipeStartNotice(context.Context, *PipeStartNoticeRequest) (*PipeStartNoticeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PipeStartNotice not implemented")
}
func (UnimplementedSshPiperPluginServer) PipeErrorNotice(context.Context, *PipeErrorNoticeRequest) (*PipeErrorNoticeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PipeErrorNotice not implemented")
}
func (UnimplementedSshPiperPluginServer) mustEmbedUnimplementedSshPiperPluginServer() {}

// UnsafeSshPiperPluginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SshPiperPluginServer will
// result in compilation errors.
type UnsafeSshPiperPluginServer interface {
	mustEmbedUnimplementedSshPiperPluginServer()
}

func RegisterSshPiperPluginServer(s grpc.ServiceRegistrar, srv SshPiperPluginServer) {
	s.RegisterService(&SshPiperPlugin_ServiceDesc, srv)
}

func _SshPiperPlugin_Logs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StartLogRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SshPiperPluginServer).Logs(m, &sshPiperPluginLogsServer{stream})
}

type SshPiperPlugin_LogsServer interface {
	Send(*Log) error
	grpc.ServerStream
}

type sshPiperPluginLogsServer struct {
	grpc.ServerStream
}

func (x *sshPiperPluginLogsServer) Send(m *Log) error {
	return x.ServerStream.SendMsg(m)
}

func _SshPiperPlugin_ListCallbacks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCallbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SshPiperPluginServer).ListCallbacks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libplugin.SshPiperPlugin/ListCallbacks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SshPiperPluginServer).ListCallbacks(ctx, req.(*ListCallbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SshPiperPlugin_NewConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SshPiperPluginServer).NewConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libplugin.SshPiperPlugin/NewConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SshPiperPluginServer).NewConnection(ctx, req.(*NewConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SshPiperPlugin_NextAuthMethods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NextAuthMethodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SshPiperPluginServer).NextAuthMethods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libplugin.SshPiperPlugin/NextAuthMethods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SshPiperPluginServer).NextAuthMethods(ctx, req.(*NextAuthMethodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SshPiperPlugin_NoneAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoneAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SshPiperPluginServer).NoneAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libplugin.SshPiperPlugin/NoneAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SshPiperPluginServer).NoneAuth(ctx, req.(*NoneAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SshPiperPlugin_PasswordAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SshPiperPluginServer).PasswordAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libplugin.SshPiperPlugin/PasswordAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SshPiperPluginServer).PasswordAuth(ctx, req.(*PasswordAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SshPiperPlugin_PublicKeyAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublicKeyAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SshPiperPluginServer).PublicKeyAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libplugin.SshPiperPlugin/PublicKeyAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SshPiperPluginServer).PublicKeyAuth(ctx, req.(*PublicKeyAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SshPiperPlugin_KeyboardInteractiveAuth_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SshPiperPluginServer).KeyboardInteractiveAuth(&sshPiperPluginKeyboardInteractiveAuthServer{stream})
}

type SshPiperPlugin_KeyboardInteractiveAuthServer interface {
	Send(*KeyboardInteractiveAuthMessage) error
	Recv() (*KeyboardInteractiveAuthMessage, error)
	grpc.ServerStream
}

type sshPiperPluginKeyboardInteractiveAuthServer struct {
	grpc.ServerStream
}

func (x *sshPiperPluginKeyboardInteractiveAuthServer) Send(m *KeyboardInteractiveAuthMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sshPiperPluginKeyboardInteractiveAuthServer) Recv() (*KeyboardInteractiveAuthMessage, error) {
	m := new(KeyboardInteractiveAuthMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SshPiperPlugin_UpstreamAuthFailureNotice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpstreamAuthFailureNoticeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SshPiperPluginServer).UpstreamAuthFailureNotice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libplugin.SshPiperPlugin/UpstreamAuthFailureNotice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SshPiperPluginServer).UpstreamAuthFailureNotice(ctx, req.(*UpstreamAuthFailureNoticeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SshPiperPlugin_Banner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BannerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SshPiperPluginServer).Banner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libplugin.SshPiperPlugin/Banner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SshPiperPluginServer).Banner(ctx, req.(*BannerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SshPiperPlugin_VerifyHostKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyHostKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SshPiperPluginServer).VerifyHostKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libplugin.SshPiperPlugin/VerifyHostKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SshPiperPluginServer).VerifyHostKey(ctx, req.(*VerifyHostKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SshPiperPlugin_PipeCreateErrorNotice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PipeCreateErrorNoticeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SshPiperPluginServer).PipeCreateErrorNotice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libplugin.SshPiperPlugin/PipeCreateErrorNotice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SshPiperPluginServer).PipeCreateErrorNotice(ctx, req.(*PipeCreateErrorNoticeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SshPiperPlugin_PipeStartNotice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PipeStartNoticeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SshPiperPluginServer).PipeStartNotice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libplugin.SshPiperPlugin/PipeStartNotice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SshPiperPluginServer).PipeStartNotice(ctx, req.(*PipeStartNoticeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SshPiperPlugin_PipeErrorNotice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PipeErrorNoticeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SshPiperPluginServer).PipeErrorNotice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libplugin.SshPiperPlugin/PipeErrorNotice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SshPiperPluginServer).PipeErrorNotice(ctx, req.(*PipeErrorNoticeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SshPiperPlugin_ServiceDesc is the grpc.ServiceDesc for SshPiperPlugin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SshPiperPlugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "libplugin.SshPiperPlugin",
	HandlerType: (*SshPiperPluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCallbacks",
			Handler:    _SshPiperPlugin_ListCallbacks_Handler,
		},
		{
			MethodName: "NewConnection",
			Handler:    _SshPiperPlugin_NewConnection_Handler,
		},
		{
			MethodName: "NextAuthMethods",
			Handler:    _SshPiperPlugin_NextAuthMethods_Handler,
		},
		{
			MethodName: "NoneAuth",
			Handler:    _SshPiperPlugin_NoneAuth_Handler,
		},
		{
			MethodName: "PasswordAuth",
			Handler:    _SshPiperPlugin_PasswordAuth_Handler,
		},
		{
			MethodName: "PublicKeyAuth",
			Handler:    _SshPiperPlugin_PublicKeyAuth_Handler,
		},
		{
			MethodName: "UpstreamAuthFailureNotice",
			Handler:    _SshPiperPlugin_UpstreamAuthFailureNotice_Handler,
		},
		{
			MethodName: "Banner",
			Handler:    _SshPiperPlugin_Banner_Handler,
		},
		{
			MethodName: "VerifyHostKey",
			Handler:    _SshPiperPlugin_VerifyHostKey_Handler,
		},
		{
			MethodName: "PipeCreateErrorNotice",
			Handler:    _SshPiperPlugin_PipeCreateErrorNotice_Handler,
		},
		{
			MethodName: "PipeStartNotice",
			Handler:    _SshPiperPlugin_PipeStartNotice_Handler,
		},
		{
			MethodName: "PipeErrorNotice",
			Handler:    _SshPiperPlugin_PipeErrorNotice_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Logs",
			Handler:       _SshPiperPlugin_Logs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "KeyboardInteractiveAuth",
			Handler:       _SshPiperPlugin_KeyboardInteractiveAuth_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "plugin.proto",
}
