package server

import (
	"fmt"
	"net"

	loggerbin "github.com/salesforceanton/grpc-logger-bin/pkg/domain"
	"google.golang.org/grpc"
)

type Server struct {
	grpcSrv *grpc.Server
	handler loggerbin.LoggerbinServiceServer
}

func New(server loggerbin.LoggerbinServiceServer) *Server {
	return &Server{
		grpcSrv: grpc.NewServer(),
		handler: server,
	}
}

func (s *Server) ListenAndServe(port int) error {
	addr := fmt.Sprintf(":%d", port)

	// Initilialize listener on target port
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	// Start cycle as a http.ListenAndServe
	loggerbin.RegisterLoggerbinServiceServer(s.grpcSrv, s.handler)
	if err := s.grpcSrv.Serve(lis); err != nil {
		return err
	}

	return nil
}
