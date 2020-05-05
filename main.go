package main

import (
	"log"
	"net"

	"github.com/echo-marche/sendmail-api/config"
	pb "github.com/echo-marche/sendmail-api/proto/pb"
	"github.com/echo-marche/sendmail-api/servers"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
)

func main() {
	// init gRPC server
	listenPort, err := net.Listen("tcp", ":"+config.GetEnv("API_PORT"))
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_validator.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
		)),
	)
	sendmailServer := &servers.SendmailServer{}
	pb.RegisterSendmailServer(server, sendmailServer)
	if err := server.Serve(listenPort); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
