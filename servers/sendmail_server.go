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

func (server *SendmailServer) SendSample(ctx context.Context, req *pb.EmailRequest) (*pb.EmailResponse, error) {
	from := req.FromAddress
	to := req.ToAddress

	// func PlainAuth(identity, username, password, host string) Auth
	auth := smtp.PlainAuth("", from, "secret", config.GetEnv("SMTP_HOST"))

	msg := []byte("" +
		"From: 送信した人 <" + from + ">\r\n" +
		"To: " + to + "\r\n" +
		"Subject: 件名 " + req.Subject + " \r\n" +
		"\r\n" +
		req.Msg +
		"\r\n")

	// func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
	err := smtp.SendMail(
		config.GetEnv("SMTP_HOST")+":"+config.GetEnv("SMTP_PORT"),
		auth,
		from,
		[]string{to},
		msg,
	)
	if err != nil {
		if err != nil {
			fmt.Println(err)
			return nil, status.Errorf(codes.Unavailable,
				err.Error())
		}
	}

	fmt.Print("success")
	return &pb.EmailResponse{Status: "sendmail ok"}, nil
}
