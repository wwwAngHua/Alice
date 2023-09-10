package controller

import (
	"Alice/model"
	"Alice/util"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func LeanChat(ctx *gin.Context) {
	var data model.Chat
	if err := ctx.ShouldBindJSON(&data); err != nil {
		util.Fail(ctx, nil, "请求体错误！")
		return
	}

	switch {
	case strings.HasPrefix(data.Message, "/start"):
		util.Success(ctx, nil, "您好，Alice需要您的加入，如果您想开发自己的机器人，请参考爱丽丝 Alice编码规范：<a href=\"https://github.com/wwwAngHua/Alice\">github.com/wwwAngHua/Alice</a><br/><br/>QQ交流群：<a href=\"http://qm.qq.com/cgi-bin/qm/qr?_wv=1027&k=Rp_HOO7VyCU6eG4bsDGhkO111pSeqz9P&authKey=PWafMxRNkFzIUB8oZHOu0fCerGnEhQRuOK%2B%2BaqmSY1Cd%2FzSnblgl8UoChyE%2FKjv0&noverify=0&group_code=671810985\">671810985</a><br/><br/>您可以使用以下命令：<br/>/setting getUID -> 获取您的UID<br/>/setting getVersion -> 获取您的客户端版本<br/>/setting test -> 返回test处指定文本")
		return
	case strings.HasPrefix(data.Message, "/setting"):
		message := strings.TrimLeft(data.Message[len("/setting"):], " ")
		// TODO: 功能回复
		if message == "getUID" {
			// UID 为用户唯一标识信息
			util.Success(ctx, nil, "您的UID为："+data.Uid)
			return
		} else if message == "getVersion" {
			// Version 为用户客户端版本
			util.Success(ctx, nil, "您当前使用的Alice版本为："+strconv.FormatFloat(float64(data.Version), 'f', -1, 32))
			return
		}
		util.Success(ctx, nil, message)
		return
	default:
		util.Fail(ctx, nil, "抱歉，我正在完善中，未能帮助您解决问题，深感抱歉！")
	}
}
