package servers

import (
	"context"

	pb "github.com/echo-marche/sendmail-api/proto/pb"
)

type SendmailServer struct{}

func (server *SendmailServer) SendUserRegistration(ctx context.Context, req *pb.EmailRequest) (*pb.EmailResponse, error) {
	// systemCode := req.GetSystemCode
	return &pb.EmailResponse{Status: "ok"}, nil
}
