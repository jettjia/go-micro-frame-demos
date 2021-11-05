package ceshi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"web-example/amqp/producer"
)

func SendMq(ctx *gin.Context)  {
	//goods_id := ctx.DefaultQuery("goods_id", "0")

	//producer.TestGoods(goods_id)

	ctx.JSON(http.StatusOK, nil)
}
