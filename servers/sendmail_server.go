package servers

import (
	"context"
	"fmt"
	"net/smtp"

	"github.com/echo-marche/sendmail-api/config"
	pb "github.com/echo-marche/sendmail-api/proto/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SendmailServer struct{}

type unencryptedAuth struct {
	smtp.Auth
}

func (a unencryptedAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	s := *server
	s.TLS = true
	return a.Auth.Start(&s)
}

func (server *SendmailServer) SendSample(ctx context.Context, req *pb.EmailRequest) (*pb.EmailResponse, error) {
	from := req.FromAddress
	to := req.ToAddress

	auth := unencryptedAuth{smtp.PlainAuth("", from, "password", config.GetEnv("SMTP_HOST"))}

	msg := []byte("" +
		"From: " + req.FromUserName + " <" + from + ">\r\n" +
		"To: " + to + "\r\n" +
		"Subject: 件名 " + req.Subject + " \r\n" +
		"\r\n" +
		req.Msg +
		"\r\n")

	err := smtp.SendMail(
		config.GetEnv("SMTP_HOST")+":"+config.GetEnv("SMTP_PORT"),
		auth,
		from,
		[]string{to},
		[]byte(msg),
	)
	if err != nil {
		fmt.Println(err)
		return nil, status.Errorf(codes.Unavailable,
			err.Error())
	}
	return &pb.EmailResponse{Status: "sendmail ok"}, nil
}
