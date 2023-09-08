package controller

import (
	"Alice/model"
	"Alice/util"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func Chat(ctx *gin.Context) {
	var data model.Chat
	if err := ctx.ShouldBindJSON(&data); err != nil {
		util.Fail(ctx, nil, "请求体错误！")
		return
	}

	switch {
	case strings.HasPrefix(data.Message, "/start"):
		util.Success(ctx, nil, "<img src=\"https://cdn.jsdelivr.net/gh/wwwanghua/Alice@main/image/alice.jpg\"><br/>我叫<span style=\"color: yellow;\">爱丽丝</span>。是您的私人助手，我会尽力帮助您完成各种任务。请问有什么我可以帮助您的吗？<a href=\"https://cdn.jsdelivr.net/gh/wwwanghua/Alice@main/image/alice.jpg/\"><br/><br/>进一步了解我们！</a>")
		return
	case strings.HasPrefix(data.Message, "/setting"):
		message := strings.TrimLeft(data.Message[len("/setting"):], " ")
		// TODO: 功能回复
		if message == "getUID" {
			// UID为用户唯一标识信息
			util.Success(ctx, nil, data.Uid)
			return
		}
		util.Success(ctx, nil, message)
		return
	default:
		util.Fail(ctx, nil, "抱歉，我不明白您的意思...(•̀⌓•́)シ")
	}
}

func Upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		util.Fail(ctx, nil, "主人，上传失败啦："+err.Error())
		return
	}

	// TODO: 文件相关处理

	workDir, _ := os.Getwd()
	filePath := workDir + "/upload/" + file.Filename
	ctx.SaveUploadedFile(file, filePath)
	util.Success(ctx, nil, "主人，上传成功啦!")
}
