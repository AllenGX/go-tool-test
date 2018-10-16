package svc

import (
	"context"
)

// @microgen middleware, logging, grpc, http, recovering, error-logging, tracing, caching, metrics
// @protobuf testmicrogen/pb
type Service interface {
	Hello(ctx context.Context, name string) (a context.Context, b string, c error)
}

type TestRequest struct {
	Name string
}

type TestResponse struct {
	Ctx context.Context
	Msg string
}
