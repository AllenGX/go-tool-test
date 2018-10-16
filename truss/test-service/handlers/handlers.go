package handlers

import (
	"context"

	pb "test/truss/pb"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.TestServer {
	return testService{}
}

type testService struct{}

// Hello implements Service.
func (s testService) Hello(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	var resp pb.Response
	resp = pb.Response{
		// Msg:
		Msg: "nihao",
	}
	return &resp, nil
}
