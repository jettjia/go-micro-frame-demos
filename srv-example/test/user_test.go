package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/jettjia/go-micro-frame/core/register/nacos"

	"srv-example/proto"
)

func init() {
	InitClient()
}

func Test_GetUserList(t *testing.T) {
	rsp, err := GrpcClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    1,
		PSize: 5,
	})
	if err != nil {
		panic(err)
	}
	for _, user := range rsp.Data {
		fmt.Println(user.Mobile, user.NickName, user.PassWord)
		if err != nil {
			panic(err)
		}
	}

	defer ClientConn.Close()
}

func Test_GetUserById(t *testing.T) {
	rsp, err := GrpcClient.GetUserById(context.Background(), &proto.IdRequest{Id: 1})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp)

	defer ClientConn.Close()
}

// nacos方式，负载均衡获取服务
func Test_Discovery(t *testing.T) {
	client := nacos.NewRegistryClient("10.4.7.71", 8848, "b79d1500-d143-447d-873a-4d545e3d186c", "nacos", "nacos")

	ins, err := client.Discovery("gomicrom-srv", "dev")
	if err != nil {
		panic(err)
	}

	userClient := proto.NewUserClient(ins)
	rsp, err := userClient.GetUserById(context.Background(), &proto.IdRequest{Id: 1})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp)

}
