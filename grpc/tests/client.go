package tests

import (
	"go-micro-frame/proto"
	"google.golang.org/grpc"
)

var GrpcClient proto.UserClient
var ClientConn *grpc.ClientConn

func InitClient() {
	var err error
	ClientConn, err = grpc.Dial("10.4.7.71:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	GrpcClient = proto.NewUserClient(ClientConn)
}
