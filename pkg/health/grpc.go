package health

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
//	"runtime"
)

type healthServer struct {
	grpc_health_v1.UnimplementedHealthServer
}

func (s *healthServer) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

// func main() {
//	go StartGrpcHealthServer()

//	runtime.Gosched()
// }

func StartGrpcHealthServer() {

	defer func() {
		if r := recover(); r != nil {
			log.Printf("grpc health server error : %v", r)
			return
		}
	}()
	for usePort := 50051; usePort < 2<<15; usePort += 100 {
		lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", usePort))
		if err != nil {
			//log.Fatalf("failed to listen: %v", err)
			continue
		}
		s := grpc.NewServer()
		grpc_health_v1.RegisterHealthServer(s, &healthServer{})
		log.Printf("Starting GRPC Health Server on %s", lis.Addr().String())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
			return
		}
	}
}
