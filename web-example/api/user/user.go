package user

import (
	"context"
	"web-example/api"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"web-example/global"
	reponse "web-example/global/response"
	"web-example/proto"
)

func GetUserList(ctx *gin.Context) {
	// 获取请求参数
	pn := ctx.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	pSize := ctx.DefaultQuery("psize", "10")
	pSizeInt, _ := strconv.Atoi(pSize)

	// grpc远程调用，传递ginContext
	rsp, err := global.UserSrvClient.GetUserList(context.WithValue(context.Background(), "ginContext", ctx), &proto.PageInfo{
		Pn:    uint32(pnInt),
		PSize: uint32(pSizeInt),
	})

	if err != nil {
		zap.S().Errorw("[GetUserList] 查询 【用户列表】失败")
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}

	// 组装返回结果
	reMap := gin.H{
		"total": rsp.Total,
	}
	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		user := reponse.UserResponse{
			Id:       value.Id,
			NickName: value.NickName,
			Birthday: reponse.JsonTime(time.Unix(int64(value.BirthDay), 0)),
			Gender:   value.Gender,
			Mobile:   value.Mobile,
		}
		result = append(result, user)
	}

	reMap["data"] = result
	ctx.JSON(http.StatusOK, reMap)
}