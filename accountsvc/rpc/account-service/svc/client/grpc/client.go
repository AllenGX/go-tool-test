// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 7dc4d5d85c
// Version Date: Mon May 28 22:12:59 UTC 2018

// Package grpc provides a gRPC client for the Account service.
package grpc

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	"test/accountsvc/rpc/account-service/svc"
	pb "test/accountsvc/rpc/pb"
)

// New returns an service backed by a gRPC client connection. It is the
// responsibility of the caller to dial, and later close, the connection.
func New(conn *grpc.ClientConn, options ...ClientOption) (pb.AccountServer, error) {
	var cc clientConfig

	for _, f := range options {
		err := f(&cc)
		if err != nil {
			return nil, errors.Wrap(err, "cannot apply option")
		}
	}

	clientOptions := []grpctransport.ClientOption{
		grpctransport.ClientBefore(
			contextValuesToGRPCMetadata(cc.headers)),
	}
	var registerlocalaccountEndpoint endpoint.Endpoint
	{
		registerlocalaccountEndpoint = grpctransport.NewClient(
			conn,
			"pb.Account",
			"RegisterLocalAccount",
			EncodeGRPCRegisterLocalAccountRequest,
			DecodeGRPCRegisterLocalAccountResponse,
			pb.RegisterLocalAccountRsp{},
			clientOptions...,
		).Endpoint()
	}

	var registerldapaccountEndpoint endpoint.Endpoint
	{
		registerldapaccountEndpoint = grpctransport.NewClient(
			conn,
			"pb.Account",
			"RegisterLDAPAccount",
			EncodeGRPCRegisterLDAPAccountRequest,
			DecodeGRPCRegisterLDAPAccountResponse,
			pb.RegisterLDAPAccountRsp{},
			clientOptions...,
		).Endpoint()
	}

	var unregisteraccountEndpoint endpoint.Endpoint
	{
		unregisteraccountEndpoint = grpctransport.NewClient(
			conn,
			"pb.Account",
			"UnRegisterAccount",
			EncodeGRPCUnRegisterAccountRequest,
			DecodeGRPCUnRegisterAccountResponse,
			pb.UnRegisterAccountRsp{},
			clientOptions...,
		).Endpoint()
	}

	var getaccountinfoEndpoint endpoint.Endpoint
	{
		getaccountinfoEndpoint = grpctransport.NewClient(
			conn,
			"pb.Account",
			"GetAccountInfo",
			EncodeGRPCGetAccountInfoRequest,
			DecodeGRPCGetAccountInfoResponse,
			pb.GetAccountInfoRsp{},
			clientOptions...,
		).Endpoint()
	}

	var getaccountcountEndpoint endpoint.Endpoint
	{
		getaccountcountEndpoint = grpctransport.NewClient(
			conn,
			"pb.Account",
			"GetAccountCount",
			EncodeGRPCGetAccountCountRequest,
			DecodeGRPCGetAccountCountResponse,
			pb.GetAccountCountRsp{},
			clientOptions...,
		).Endpoint()
	}

	var getaccountlistEndpoint endpoint.Endpoint
	{
		getaccountlistEndpoint = grpctransport.NewClient(
			conn,
			"pb.Account",
			"GetAccountList",
			EncodeGRPCGetAccountListRequest,
			DecodeGRPCGetAccountListResponse,
			pb.GetAccountListRsp{},
			clientOptions...,
		).Endpoint()
	}

	var loginEndpoint endpoint.Endpoint
	{
		loginEndpoint = grpctransport.NewClient(
			conn,
			"pb.Account",
			"Login",
			EncodeGRPCLoginRequest,
			DecodeGRPCLoginResponse,
			pb.LoginRsp{},
			clientOptions...,
		).Endpoint()
	}

	var logoutEndpoint endpoint.Endpoint
	{
		logoutEndpoint = grpctransport.NewClient(
			conn,
			"pb.Account",
			"Logout",
			EncodeGRPCLogoutRequest,
			DecodeGRPCLogoutResponse,
			pb.LogoutRsp{},
			clientOptions...,
		).Endpoint()
	}

	var getsessioninfoEndpoint endpoint.Endpoint
	{
		getsessioninfoEndpoint = grpctransport.NewClient(
			conn,
			"pb.Account",
			"GetSessionInfo",
			EncodeGRPCGetSessionInfoRequest,
			DecodeGRPCGetSessionInfoResponse,
			pb.GetSessionInfoRsp{},
			clientOptions...,
		).Endpoint()
	}

	var setsessionuserdataEndpoint endpoint.Endpoint
	{
		setsessionuserdataEndpoint = grpctransport.NewClient(
			conn,
			"pb.Account",
			"SetSessionUserData",
			EncodeGRPCSetSessionUserDataRequest,
			DecodeGRPCSetSessionUserDataResponse,
			pb.SetSessionUserDataRsp{},
			clientOptions...,
		).Endpoint()
	}

	var getsessionuserdataEndpoint endpoint.Endpoint
	{
		getsessionuserdataEndpoint = grpctransport.NewClient(
			conn,
			"pb.Account",
			"GetSessionUserData",
			EncodeGRPCGetSessionUserDataRequest,
			DecodeGRPCGetSessionUserDataResponse,
			pb.GetSessionUserDataRsp{},
			clientOptions...,
		).Endpoint()
	}

	var registerappEndpoint endpoint.Endpoint
	{
		registerappEndpoint = grpctransport.NewClient(
			conn,
			"pb.Account",
			"RegisterApp",
			EncodeGRPCRegisterAppRequest,
			DecodeGRPCRegisterAppResponse,
			pb.RegisterAppRsp{},
			clientOptions...,
		).Endpoint()
	}

	var unregisterappEndpoint endpoint.Endpoint
	{
		unregisterappEndpoint = grpctransport.NewClient(
			conn,
			"pb.Account",
			"UnRegisterApp",
			EncodeGRPCUnRegisterAppRequest,
			DecodeGRPCUnRegisterAppResponse,
			pb.UnRegisterAppRsp{},
			clientOptions...,
		).Endpoint()
	}

	return svc.Endpoints{
		RegisterLocalAccountEndpoint: registerlocalaccountEndpoint,
		RegisterLDAPAccountEndpoint:  registerldapaccountEndpoint,
		UnRegisterAccountEndpoint:    unregisteraccountEndpoint,
		GetAccountInfoEndpoint:       getaccountinfoEndpoint,
		GetAccountCountEndpoint:      getaccountcountEndpoint,
		GetAccountListEndpoint:       getaccountlistEndpoint,
		LoginEndpoint:                loginEndpoint,
		LogoutEndpoint:               logoutEndpoint,
		GetSessionInfoEndpoint:       getsessioninfoEndpoint,
		SetSessionUserDataEndpoint:   setsessionuserdataEndpoint,
		GetSessionUserDataEndpoint:   getsessionuserdataEndpoint,
		RegisterAppEndpoint:          registerappEndpoint,
		UnRegisterAppEndpoint:        unregisterappEndpoint,
	}, nil
}

// GRPC Client Decode

// DecodeGRPCRegisterLocalAccountResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC registerlocalaccount reply to a user-domain registerlocalaccount response. Primarily useful in a client.
func DecodeGRPCRegisterLocalAccountResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.RegisterLocalAccountRsp)
	return reply, nil
}

// DecodeGRPCRegisterLDAPAccountResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC registerldapaccount reply to a user-domain registerldapaccount response. Primarily useful in a client.
func DecodeGRPCRegisterLDAPAccountResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.RegisterLDAPAccountRsp)
	return reply, nil
}

// DecodeGRPCUnRegisterAccountResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC unregisteraccount reply to a user-domain unregisteraccount response. Primarily useful in a client.
func DecodeGRPCUnRegisterAccountResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.UnRegisterAccountRsp)
	return reply, nil
}

// DecodeGRPCGetAccountInfoResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC getaccountinfo reply to a user-domain getaccountinfo response. Primarily useful in a client.
func DecodeGRPCGetAccountInfoResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.GetAccountInfoRsp)
	return reply, nil
}

// DecodeGRPCGetAccountCountResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC getaccountcount reply to a user-domain getaccountcount response. Primarily useful in a client.
func DecodeGRPCGetAccountCountResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.GetAccountCountRsp)
	return reply, nil
}

// DecodeGRPCGetAccountListResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC getaccountlist reply to a user-domain getaccountlist response. Primarily useful in a client.
func DecodeGRPCGetAccountListResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.GetAccountListRsp)
	return reply, nil
}

// DecodeGRPCLoginResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC login reply to a user-domain login response. Primarily useful in a client.
func DecodeGRPCLoginResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.LoginRsp)
	return reply, nil
}

// DecodeGRPCLogoutResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC logout reply to a user-domain logout response. Primarily useful in a client.
func DecodeGRPCLogoutResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.LogoutRsp)
	return reply, nil
}

// DecodeGRPCGetSessionInfoResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC getsessioninfo reply to a user-domain getsessioninfo response. Primarily useful in a client.
func DecodeGRPCGetSessionInfoResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.GetSessionInfoRsp)
	return reply, nil
}

// DecodeGRPCSetSessionUserDataResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC setsessionuserdata reply to a user-domain setsessionuserdata response. Primarily useful in a client.
func DecodeGRPCSetSessionUserDataResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.SetSessionUserDataRsp)
	return reply, nil
}

// DecodeGRPCGetSessionUserDataResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC getsessionuserdata reply to a user-domain getsessionuserdata response. Primarily useful in a client.
func DecodeGRPCGetSessionUserDataResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.GetSessionUserDataRsp)
	return reply, nil
}

// DecodeGRPCRegisterAppResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC registerapp reply to a user-domain registerapp response. Primarily useful in a client.
func DecodeGRPCRegisterAppResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.RegisterAppRsp)
	return reply, nil
}

// DecodeGRPCUnRegisterAppResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC unregisterapp reply to a user-domain unregisterapp response. Primarily useful in a client.
func DecodeGRPCUnRegisterAppResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.UnRegisterAppRsp)
	return reply, nil
}

// GRPC Client Encode

// EncodeGRPCRegisterLocalAccountRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain registerlocalaccount request to a gRPC registerlocalaccount request. Primarily useful in a client.
func EncodeGRPCRegisterLocalAccountRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.RegisterLocalAccountReq)
	return req, nil
}

// EncodeGRPCRegisterLDAPAccountRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain registerldapaccount request to a gRPC registerldapaccount request. Primarily useful in a client.
func EncodeGRPCRegisterLDAPAccountRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.RegisterLDAPAccountReq)
	return req, nil
}

// EncodeGRPCUnRegisterAccountRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain unregisteraccount request to a gRPC unregisteraccount request. Primarily useful in a client.
func EncodeGRPCUnRegisterAccountRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UnRegisterAccountReq)
	return req, nil
}

// EncodeGRPCGetAccountInfoRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain getaccountinfo request to a gRPC getaccountinfo request. Primarily useful in a client.
func EncodeGRPCGetAccountInfoRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetAccountInfoReq)
	return req, nil
}

// EncodeGRPCGetAccountCountRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain getaccountcount request to a gRPC getaccountcount request. Primarily useful in a client.
func EncodeGRPCGetAccountCountRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetAccountCountReq)
	return req, nil
}

// EncodeGRPCGetAccountListRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain getaccountlist request to a gRPC getaccountlist request. Primarily useful in a client.
func EncodeGRPCGetAccountListRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetAccountListReq)
	return req, nil
}

// EncodeGRPCLoginRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain login request to a gRPC login request. Primarily useful in a client.
func EncodeGRPCLoginRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.LoginReq)
	return req, nil
}

// EncodeGRPCLogoutRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain logout request to a gRPC logout request. Primarily useful in a client.
func EncodeGRPCLogoutRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.LogoutReq)
	return req, nil
}

// EncodeGRPCGetSessionInfoRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain getsessioninfo request to a gRPC getsessioninfo request. Primarily useful in a client.
func EncodeGRPCGetSessionInfoRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetSessionInfoReq)
	return req, nil
}

// EncodeGRPCSetSessionUserDataRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain setsessionuserdata request to a gRPC setsessionuserdata request. Primarily useful in a client.
func EncodeGRPCSetSessionUserDataRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.SetSessionUserDataReq)
	return req, nil
}

// EncodeGRPCGetSessionUserDataRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain getsessionuserdata request to a gRPC getsessionuserdata request. Primarily useful in a client.
func EncodeGRPCGetSessionUserDataRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetSessionUserDataReq)
	return req, nil
}

// EncodeGRPCRegisterAppRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain registerapp request to a gRPC registerapp request. Primarily useful in a client.
func EncodeGRPCRegisterAppRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.RegisterAppReq)
	return req, nil
}

// EncodeGRPCUnRegisterAppRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain unregisterapp request to a gRPC unregisterapp request. Primarily useful in a client.
func EncodeGRPCUnRegisterAppRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UnRegisterAppReq)
	return req, nil
}

type clientConfig struct {
	headers []string
}

// ClientOption is a function that modifies the client config
type ClientOption func(*clientConfig) error

func CtxValuesToSend(keys ...string) ClientOption {
	return func(o *clientConfig) error {
		o.headers = keys
		return nil
	}
}

func contextValuesToGRPCMetadata(keys []string) grpctransport.ClientRequestFunc {
	return func(ctx context.Context, md *metadata.MD) context.Context {
		var pairs []string
		for _, k := range keys {
			if v, ok := ctx.Value(k).(string); ok {
				pairs = append(pairs, k, v)
			}
		}

		if pairs != nil {
			*md = metadata.Join(*md, metadata.Pairs(pairs...))
		}

		return ctx
	}
}
