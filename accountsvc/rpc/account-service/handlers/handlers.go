package handlers

import (
	"context"

	pb "test/accountsvc/rpc/pb"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.AccountServer {
	return accountService{}
}

type accountService struct{}

// RegisterLocalAccount implements Service.
func (s accountService) RegisterLocalAccount(ctx context.Context, in *pb.RegisterLocalAccountReq) (*pb.RegisterLocalAccountRsp, error) {
	var resp pb.RegisterLocalAccountRsp
	resp = pb.RegisterLocalAccountRsp{
		// Common:
		// AccountId:
	}
	return &resp, nil
}

// RegisterLDAPAccount implements Service.
func (s accountService) RegisterLDAPAccount(ctx context.Context, in *pb.RegisterLDAPAccountReq) (*pb.RegisterLDAPAccountRsp, error) {
	var resp pb.RegisterLDAPAccountRsp
	resp = pb.RegisterLDAPAccountRsp{
		// Common:
		// AccountId:
	}
	return &resp, nil
}

// UnRegisterAccount implements Service.
func (s accountService) UnRegisterAccount(ctx context.Context, in *pb.UnRegisterAccountReq) (*pb.UnRegisterAccountRsp, error) {
	var resp pb.UnRegisterAccountRsp
	resp = pb.UnRegisterAccountRsp{
		// Common:
	}
	return &resp, nil
}

// GetAccountInfo implements Service.
func (s accountService) GetAccountInfo(ctx context.Context, in *pb.GetAccountInfoReq) (*pb.GetAccountInfoRsp, error) {
	var resp pb.GetAccountInfoRsp
	resp = pb.GetAccountInfoRsp{
		// Common:
		// AccountInfo:
	}
	return &resp, nil
}

// GetAccountCount implements Service.
func (s accountService) GetAccountCount(ctx context.Context, in *pb.GetAccountCountReq) (*pb.GetAccountCountRsp, error) {
	var resp pb.GetAccountCountRsp
	resp = pb.GetAccountCountRsp{
		// Common:
		// Count:
	}
	return &resp, nil
}

// GetAccountList implements Service.
func (s accountService) GetAccountList(ctx context.Context, in *pb.GetAccountListReq) (*pb.GetAccountListRsp, error) {
	var resp pb.GetAccountListRsp
	resp = pb.GetAccountListRsp{
		// Common:
		// AccountInfoList:
	}
	return &resp, nil
}

// Login implements Service.
func (s accountService) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginRsp, error) {
	var resp pb.LoginRsp
	resp = pb.LoginRsp{
		// Common:
		// Token:
	}
	return &resp, nil
}

// Logout implements Service.
func (s accountService) Logout(ctx context.Context, in *pb.LogoutReq) (*pb.LogoutRsp, error) {
	var resp pb.LogoutRsp
	resp = pb.LogoutRsp{
		// Common:
	}
	return &resp, nil
}

// GetSessionInfo implements Service.
func (s accountService) GetSessionInfo(ctx context.Context, in *pb.GetSessionInfoReq) (*pb.GetSessionInfoRsp, error) {
	var resp pb.GetSessionInfoRsp
	resp = pb.GetSessionInfoRsp{
		// Common:
		// AccountInfo:
	}
	return &resp, nil
}

// SetSessionUserData implements Service.
func (s accountService) SetSessionUserData(ctx context.Context, in *pb.SetSessionUserDataReq) (*pb.SetSessionUserDataRsp, error) {
	var resp pb.SetSessionUserDataRsp
	resp = pb.SetSessionUserDataRsp{
		// Common:
	}
	return &resp, nil
}

// GetSessionUserData implements Service.
func (s accountService) GetSessionUserData(ctx context.Context, in *pb.GetSessionUserDataReq) (*pb.GetSessionUserDataRsp, error) {
	var resp pb.GetSessionUserDataRsp
	resp = pb.GetSessionUserDataRsp{
		// Common:
		// UserData:
	}
	return &resp, nil
}

// RegisterApp implements Service.
func (s accountService) RegisterApp(ctx context.Context, in *pb.RegisterAppReq) (*pb.RegisterAppRsp, error) {
	var resp pb.RegisterAppRsp
	resp = pb.RegisterAppRsp{
		// Common:
	}
	return &resp, nil
}

// UnRegisterApp implements Service.
func (s accountService) UnRegisterApp(ctx context.Context, in *pb.UnRegisterAppReq) (*pb.UnRegisterAppRsp, error) {
	var resp pb.UnRegisterAppRsp
	resp = pb.UnRegisterAppRsp{
		// Common:
	}
	return &resp, nil
}
