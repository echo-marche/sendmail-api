package servers

import (
	"context"

	"google.golang.org/grpc/codes"
	health "google.golang.org/grpc/health/grpc_health_v1" // here
	"google.golang.org/grpc/status"
)

type HealthServer struct {
}

// here
func (h *HealthServer) Check(context.Context, *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{
		Status: health.HealthCheckResponse_SERVING,
	}, nil
}

// here
func (h *HealthServer) Watch(*health.HealthCheckRequest, health.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "watch is not implemented.")
}
