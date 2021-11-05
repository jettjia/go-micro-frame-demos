package handler

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"srv-example/domain/model"
	"srv-example/domain/service"
	"srv-example/global"
	"srv-example/proto"
)

type UserServer struct {
	proto.UnimplementedUserServer
	UserService service.IUserService
}

func ModelToResponse(user *model.User) proto.UserInfoResponse {
	//在grpc的message中字段有默认值，你不能随便赋值nil进去，容易出错
	//这里要搞清， 哪些字段是有默认值
	userInfoRsp := proto.UserInfoResponse{
		Id:       user.ID,
		PassWord: user.Password,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
		Mobile:   user.Mobile,
	}
	if user.Birthday != nil {
		userInfoRsp.BirthDay = uint64(user.Birthday.Unix())
	}
	return userInfoRsp
}

// 获取用户列表
func (s *UserServer) GetUserList(ctx context.Context, req *proto.PageInfo) (*proto.UserListResponse, error) {
	fmt.Println("我被调用了")
	fmt.Println("我调用的数据库配置是", global.ServerConfig.MysqlInfo)

	var users []model.User
	result := global.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	global.DB.Scopes(Paginate(int(req.Pn), int(req.PSize))).Find(&users)

	rsp := &proto.UserListResponse{}
	rsp.Total = int32(result.RowsAffected)

	for _, user := range users {
		userInfoRsp := ModelToResponse(&user)
		rsp.Data = append(rsp.Data, &userInfoRsp)
	}

	return rsp, nil
}

// 这里本来是根据用户id 获取用户信息的
// 在这里增加一个 分布式的 redis 锁，来测试；注意这里只是模拟，实际业务一般是可以在控制库存 加减的时候
func (s *UserServer) GetUserById(ctx context.Context, req *proto.IdRequest) (*proto.UserInfoResponse, error) {
	user, err := s.UserService.FindUserByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "未找到编号为%d的用户", req.Id)
	}

	userInfoRsp := ModelToResponse(user)
	return &userInfoRsp, nil
}
