// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 7dc4d5d85c
// Version Date: Mon May 28 22:12:59 UTC 2018

package svc

// This file provides server-side bindings for the gRPC transport.
// It utilizes the transport/grpc.Server.

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"

	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "test/accountsvc/rpc/pb"
)

// MakeGRPCServer makes a set of endpoints available as a gRPC AccountServer.
func MakeGRPCServer(endpoints Endpoints) pb.AccountServer {
	serverOptions := []grpctransport.ServerOption{
		grpctransport.ServerBefore(metadataToContext),
	}
	return &grpcServer{
		// account

		registerlocalaccount: grpctransport.NewServer(
			endpoints.RegisterLocalAccountEndpoint,
			DecodeGRPCRegisterLocalAccountRequest,
			EncodeGRPCRegisterLocalAccountResponse,
			serverOptions...,
		),
		registerldapaccount: grpctransport.NewServer(
			endpoints.RegisterLDAPAccountEndpoint,
			DecodeGRPCRegisterLDAPAccountRequest,
			EncodeGRPCRegisterLDAPAccountResponse,
			serverOptions...,
		),
		unregisteraccount: grpctransport.NewServer(
			endpoints.UnRegisterAccountEndpoint,
			DecodeGRPCUnRegisterAccountRequest,
			EncodeGRPCUnRegisterAccountResponse,
			serverOptions...,
		),
		getaccountinfo: grpctransport.NewServer(
			endpoints.GetAccountInfoEndpoint,
			DecodeGRPCGetAccountInfoRequest,
			EncodeGRPCGetAccountInfoResponse,
			serverOptions...,
		),
		getaccountcount: grpctransport.NewServer(
			endpoints.GetAccountCountEndpoint,
			DecodeGRPCGetAccountCountRequest,
			EncodeGRPCGetAccountCountResponse,
			serverOptions...,
		),
		getaccountlist: grpctransport.NewServer(
			endpoints.GetAccountListEndpoint,
			DecodeGRPCGetAccountListRequest,
			EncodeGRPCGetAccountListResponse,
			serverOptions...,
		),
		login: grpctransport.NewServer(
			endpoints.LoginEndpoint,
			DecodeGRPCLoginRequest,
			EncodeGRPCLoginResponse,
			serverOptions...,
		),
		logout: grpctransport.NewServer(
			endpoints.LogoutEndpoint,
			DecodeGRPCLogoutRequest,
			EncodeGRPCLogoutResponse,
			serverOptions...,
		),
		getsessioninfo: grpctransport.NewServer(
			endpoints.GetSessionInfoEndpoint,
			DecodeGRPCGetSessionInfoRequest,
			EncodeGRPCGetSessionInfoResponse,
			serverOptions...,
		),
		setsessionuserdata: grpctransport.NewServer(
			endpoints.SetSessionUserDataEndpoint,
			DecodeGRPCSetSessionUserDataRequest,
			EncodeGRPCSetSessionUserDataResponse,
			serverOptions...,
		),
		getsessionuserdata: grpctransport.NewServer(
			endpoints.GetSessionUserDataEndpoint,
			DecodeGRPCGetSessionUserDataRequest,
			EncodeGRPCGetSessionUserDataResponse,
			serverOptions...,
		),
		registerapp: grpctransport.NewServer(
			endpoints.RegisterAppEndpoint,
			DecodeGRPCRegisterAppRequest,
			EncodeGRPCRegisterAppResponse,
			serverOptions...,
		),
		unregisterapp: grpctransport.NewServer(
			endpoints.UnRegisterAppEndpoint,
			DecodeGRPCUnRegisterAppRequest,
			EncodeGRPCUnRegisterAppResponse,
			serverOptions...,
		),
	}
}

// grpcServer implements the AccountServer interface
type grpcServer struct {
	registerlocalaccount grpctransport.Handler
	registerldapaccount  grpctransport.Handler
	unregisteraccount    grpctransport.Handler
	getaccountinfo       grpctransport.Handler
	getaccountcount      grpctransport.Handler
	getaccountlist       grpctransport.Handler
	login                grpctransport.Handler
	logout               grpctransport.Handler
	getsessioninfo       grpctransport.Handler
	setsessionuserdata   grpctransport.Handler
	getsessionuserdata   grpctransport.Handler
	registerapp          grpctransport.Handler
	unregisterapp        grpctransport.Handler
}

// Methods for grpcServer to implement AccountServer interface

func (s *grpcServer) RegisterLocalAccount(ctx context.Context, req *pb.RegisterLocalAccountReq) (*pb.RegisterLocalAccountRsp, error) {
	_, rep, err := s.registerlocalaccount.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.RegisterLocalAccountRsp), nil
}

func (s *grpcServer) RegisterLDAPAccount(ctx context.Context, req *pb.RegisterLDAPAccountReq) (*pb.RegisterLDAPAccountRsp, error) {
	_, rep, err := s.registerldapaccount.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.RegisterLDAPAccountRsp), nil
}

func (s *grpcServer) UnRegisterAccount(ctx context.Context, req *pb.UnRegisterAccountReq) (*pb.UnRegisterAccountRsp, error) {
	_, rep, err := s.unregisteraccount.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UnRegisterAccountRsp), nil
}

func (s *grpcServer) GetAccountInfo(ctx context.Context, req *pb.GetAccountInfoReq) (*pb.GetAccountInfoRsp, error) {
	_, rep, err := s.getaccountinfo.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetAccountInfoRsp), nil
}

func (s *grpcServer) GetAccountCount(ctx context.Context, req *pb.GetAccountCountReq) (*pb.GetAccountCountRsp, error) {
	_, rep, err := s.getaccountcount.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetAccountCountRsp), nil
}

func (s *grpcServer) GetAccountList(ctx context.Context, req *pb.GetAccountListReq) (*pb.GetAccountListRsp, error) {
	_, rep, err := s.getaccountlist.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetAccountListRsp), nil
}

func (s *grpcServer) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRsp, error) {
	_, rep, err := s.login.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.LoginRsp), nil
}

func (s *grpcServer) Logout(ctx context.Context, req *pb.LogoutReq) (*pb.LogoutRsp, error) {
	_, rep, err := s.logout.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.LogoutRsp), nil
}

func (s *grpcServer) GetSessionInfo(ctx context.Context, req *pb.GetSessionInfoReq) (*pb.GetSessionInfoRsp, error) {
	_, rep, err := s.getsessioninfo.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetSessionInfoRsp), nil
}

func (s *grpcServer) SetSessionUserData(ctx context.Context, req *pb.SetSessionUserDataReq) (*pb.SetSessionUserDataRsp, error) {
	_, rep, err := s.setsessionuserdata.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SetSessionUserDataRsp), nil
}

func (s *grpcServer) GetSessionUserData(ctx context.Context, req *pb.GetSessionUserDataReq) (*pb.GetSessionUserDataRsp, error) {
	_, rep, err := s.getsessionuserdata.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetSessionUserDataRsp), nil
}

func (s *grpcServer) RegisterApp(ctx context.Context, req *pb.RegisterAppReq) (*pb.RegisterAppRsp, error) {
	_, rep, err := s.registerapp.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.RegisterAppRsp), nil
}

func (s *grpcServer) UnRegisterApp(ctx context.Context, req *pb.UnRegisterAppReq) (*pb.UnRegisterAppRsp, error) {
	_, rep, err := s.unregisterapp.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UnRegisterAppRsp), nil
}

// Server Decode

// DecodeGRPCRegisterLocalAccountRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC registerlocalaccount request to a user-domain registerlocalaccount request. Primarily useful in a server.
func DecodeGRPCRegisterLocalAccountRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.RegisterLocalAccountReq)
	return req, nil
}

// DecodeGRPCRegisterLDAPAccountRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC registerldapaccount request to a user-domain registerldapaccount request. Primarily useful in a server.
func DecodeGRPCRegisterLDAPAccountRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.RegisterLDAPAccountReq)
	return req, nil
}

// DecodeGRPCUnRegisterAccountRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC unregisteraccount request to a user-domain unregisteraccount request. Primarily useful in a server.
func DecodeGRPCUnRegisterAccountRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.UnRegisterAccountReq)
	return req, nil
}

// DecodeGRPCGetAccountInfoRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC getaccountinfo request to a user-domain getaccountinfo request. Primarily useful in a server.
func DecodeGRPCGetAccountInfoRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetAccountInfoReq)
	return req, nil
}

// DecodeGRPCGetAccountCountRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC getaccountcount request to a user-domain getaccountcount request. Primarily useful in a server.
func DecodeGRPCGetAccountCountRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetAccountCountReq)
	return req, nil
}

// DecodeGRPCGetAccountListRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC getaccountlist request to a user-domain getaccountlist request. Primarily useful in a server.
func DecodeGRPCGetAccountListRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetAccountListReq)
	return req, nil
}

// DecodeGRPCLoginRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC login request to a user-domain login request. Primarily useful in a server.
func DecodeGRPCLoginRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.LoginReq)
	return req, nil
}

// DecodeGRPCLogoutRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC logout request to a user-domain logout request. Primarily useful in a server.
func DecodeGRPCLogoutRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.LogoutReq)
	return req, nil
}

// DecodeGRPCGetSessionInfoRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC getsessioninfo request to a user-domain getsessioninfo request. Primarily useful in a server.
func DecodeGRPCGetSessionInfoRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetSessionInfoReq)
	return req, nil
}

// DecodeGRPCSetSessionUserDataRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC setsessionuserdata request to a user-domain setsessionuserdata request. Primarily useful in a server.
func DecodeGRPCSetSessionUserDataRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.SetSessionUserDataReq)
	return req, nil
}

// DecodeGRPCGetSessionUserDataRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC getsessionuserdata request to a user-domain getsessionuserdata request. Primarily useful in a server.
func DecodeGRPCGetSessionUserDataRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetSessionUserDataReq)
	return req, nil
}

// DecodeGRPCRegisterAppRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC registerapp request to a user-domain registerapp request. Primarily useful in a server.
func DecodeGRPCRegisterAppRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.RegisterAppReq)
	return req, nil
}

// DecodeGRPCUnRegisterAppRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC unregisterapp request to a user-domain unregisterapp request. Primarily useful in a server.
func DecodeGRPCUnRegisterAppRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.UnRegisterAppReq)
	return req, nil
}

// Server Encode

// EncodeGRPCRegisterLocalAccountResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain registerlocalaccount response to a gRPC registerlocalaccount reply. Primarily useful in a server.
func EncodeGRPCRegisterLocalAccountResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.RegisterLocalAccountRsp)
	return resp, nil
}

// EncodeGRPCRegisterLDAPAccountResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain registerldapaccount response to a gRPC registerldapaccount reply. Primarily useful in a server.
func EncodeGRPCRegisterLDAPAccountResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.RegisterLDAPAccountRsp)
	return resp, nil
}

// EncodeGRPCUnRegisterAccountResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain unregisteraccount response to a gRPC unregisteraccount reply. Primarily useful in a server.
func EncodeGRPCUnRegisterAccountResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.UnRegisterAccountRsp)
	return resp, nil
}

// EncodeGRPCGetAccountInfoResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain getaccountinfo response to a gRPC getaccountinfo reply. Primarily useful in a server.
func EncodeGRPCGetAccountInfoResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.GetAccountInfoRsp)
	return resp, nil
}

// EncodeGRPCGetAccountCountResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain getaccountcount response to a gRPC getaccountcount reply. Primarily useful in a server.
func EncodeGRPCGetAccountCountResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.GetAccountCountRsp)
	return resp, nil
}

// EncodeGRPCGetAccountListResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain getaccountlist response to a gRPC getaccountlist reply. Primarily useful in a server.
func EncodeGRPCGetAccountListResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.GetAccountListRsp)
	return resp, nil
}

// EncodeGRPCLoginResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain login response to a gRPC login reply. Primarily useful in a server.
func EncodeGRPCLoginResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.LoginRsp)
	return resp, nil
}

// EncodeGRPCLogoutResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain logout response to a gRPC logout reply. Primarily useful in a server.
func EncodeGRPCLogoutResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.LogoutRsp)
	return resp, nil
}

// EncodeGRPCGetSessionInfoResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain getsessioninfo response to a gRPC getsessioninfo reply. Primarily useful in a server.
func EncodeGRPCGetSessionInfoResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.GetSessionInfoRsp)
	return resp, nil
}

// EncodeGRPCSetSessionUserDataResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain setsessionuserdata response to a gRPC setsessionuserdata reply. Primarily useful in a server.
func EncodeGRPCSetSessionUserDataResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.SetSessionUserDataRsp)
	return resp, nil
}

// EncodeGRPCGetSessionUserDataResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain getsessionuserdata response to a gRPC getsessionuserdata reply. Primarily useful in a server.
func EncodeGRPCGetSessionUserDataResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.GetSessionUserDataRsp)
	return resp, nil
}

// EncodeGRPCRegisterAppResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain registerapp response to a gRPC registerapp reply. Primarily useful in a server.
func EncodeGRPCRegisterAppResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.RegisterAppRsp)
	return resp, nil
}

// EncodeGRPCUnRegisterAppResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain unregisterapp response to a gRPC unregisterapp reply. Primarily useful in a server.
func EncodeGRPCUnRegisterAppResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.UnRegisterAppRsp)
	return resp, nil
}

// Helpers

func metadataToContext(ctx context.Context, md metadata.MD) context.Context {
	for k, v := range md {
		if v != nil {
			// The key is added both in metadata format (k) which is all lower
			// and the http.CanonicalHeaderKey of the key so that it can be
			// accessed in either format
			ctx = context.WithValue(ctx, k, v[0])
			ctx = context.WithValue(ctx, http.CanonicalHeaderKey(k), v[0])
		}
	}

	return ctx
}