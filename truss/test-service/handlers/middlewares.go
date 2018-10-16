package handlers

import (
	"context"
	"fmt"
	pb "test/truss/pb"
	"test/truss/test-service/svc"
	"time"

	"github.com/go-kit/kit/metrics"
)

// LabeledMiddleware will get passed the endpoint name when passed to
// WrapAllLabeledExcept, this can be used to write a generic metrics
// middleware which can send the endpoint name to the metrics collector.
type LabeledMiddleware func(svc.Endpoints) svc.Endpoints

type InstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	CountResult    metrics.Histogram
}

// ErrorCounter is a LabeledMiddleware, when applied with WrapAllLabeledExcept name will be populated with the endpoint name, and such this middleware will
// report errors to the metric provider with the endpoint name. Feel free to
// copy this example middleware to your service.
func ErrorCounter(prometheus InstrumentingMiddleware) LabeledMiddleware {
	return func(in svc.Endpoints) svc.Endpoints {
		return svc.Endpoints{
			HelloEndpoint: func(ctx context.Context, req interface{}) (rep interface{}, err error) {
				defer func(begin time.Time) {
					lvs := []string{"method", "uppercase", "error", fmt.Sprint(err != nil)}
					prometheus.RequestCount.With(lvs...).Add(1)
					prometheus.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
				}(time.Now())
				fmt.Println(prometheus.RequestLatency)
				fmt.Println(prometheus.RequestCount)
				resp, err := in.HelloEndpoint(ctx, req)
				return resp, err
			},
			//....
		}
	}
}

// WrapEndpoints accepts the service's entire collection of endpoints, so that a
// set of middlewares can be wrapped around every middleware (e.g., access
// logging and instrumentation), and others wrapped selectively around some
// endpoints and not others (e.g., endpoints requiring authenticated access).
// Note that the final middleware wrapped will be the outermost middleware
// (i.e. applied first)
func WrapEndpoints(in svc.Endpoints) svc.Endpoints {

	// Pass a middleware you want applied to every endpoint.
	// optionally pass in endpoints by name that you want to be excluded
	// e.g.
	// in.WrapAllExcept(authMiddleware, "Status", "Ping")

	// Pass in a svc.LabeledMiddleware you want applied to every endpoint.
	// These middlewares get passed the endpoints name as their first argument when applied.
	// This can be used to write generic metric gathering middlewares that can
	// report the endpoint name for free.
	// github.com/tuneinc/truss/_example/middlewares/labeledmiddlewares.go for examples.
	// in.WrapAllLabeledExcept(errorCounter(statsdCounter), "Status", "Ping")

	// How to apply a middleware to a single endpoint.
	// in.ExampleEndpoint = authMiddleware(in.ExampleEndpoint)
	fmt.Println("middle")
	return in
}

func WrapService(in pb.TestServer) pb.TestServer {
	fmt.Println("middle2")
	return in
}
