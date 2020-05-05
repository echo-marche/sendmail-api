module github.com/echo-marche/sendmail-api

go 1.13

require (
	github.com/golang/protobuf v1.4.0-rc.4.0.20200313231945-b860323f09d0
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/jinzhu/gorm v1.9.12 // indirect
	github.com/mwitkow/go-proto-validators v0.3.0
	golang.org/x/sys v0.0.0-20200501145240-bc7a7d42d5c3 // indirect
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.21.0
)

// for realize
replace gopkg.in/urfave/cli.v2 => github.com/urfave/cli/v2 v2.2.0
