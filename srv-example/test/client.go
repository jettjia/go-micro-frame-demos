package test

import (
	"google.golang.org/grpc"

	"srv-example/proto"
)

var GrpcClient proto.UserClient
var ClientConn *grpc.ClientConn

func InitClient() {
	var err error
	ClientConn, err = grpc.Dial("192.168.106.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	GrpcClient = proto.NewUserClient(ClientConn)
}
