package controller

import (
	"Alice/model"
	"Alice/util"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

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
		util.Success(ctx, nil, "<img src=\"https://db.t1y.net/alice.jpg\"><br/>我叫<span style=\"color: yellow;\">爱丽丝</span>。是您的私人助手，我会尽力帮助您完成各种任务。请问有什么我可以帮助您的吗？<a href=\"https://github.com/wwwAngHua/Alice\"><br/><br/>进一步了解我们！</a>")
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
		/*** 接入GPT3 ***/
		// 创建问题
		id, err := createQuestion(data.Message)
		if err != nil {
			util.Fail(ctx, nil, "AI大模型连接失败！")
			return
		}

		// 查询问题回复结果
		message, err := queryQuestionResult(id)
		if err != nil {
			util.Fail(ctx, nil, "AI大模型处理超时！")
			return
		}
		util.Success(ctx, nil, message)
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
	util.Success(ctx, nil, "主人，上传成功啦!该功能只是一个实例，目前没有实际作用～")
}

// 创建问题（GPT3）
func createQuestion(msg string) (string, error) {
	url := "https://api.takomo.ai/6e82a570-f79e-4918-bff5-deb58d7c03fd"
	token := "Bearer tk_fe27c5cc1131de0fe83dd38fc2e839bc29ea95fb4497bf3d495791eee0c98fe1a4aff59973f5e43eba91672111c3d600"

	payload := map[string]string{
		"message": msg,
	}
	payloadBytes, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", err
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("authorization", token)
	req.Header.Set("content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var response map[string]interface{}
	fmt.Println(response)
	json.NewDecoder(resp.Body).Decode(&response)

	id := response["id"].(string)
	return id, nil
}

// 查询问题回复结果
func queryQuestionResult(id string) (string, error) {
	url := fmt.Sprintf("https://api.takomo.ai/inferences/%s", id)
	token := "Bearer tk_fe27c5cc1131de0fe83dd38fc2e839bc29ea95fb4497bf3d495791eee0c98fe1a4aff59973f5e43eba91672111c3d600"

	for i := 0; i < 10; i++ {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return "", err
		}

		req.Header.Set("accept", "application/json")
		req.Header.Set("authorization", token)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()

		var response map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&response)
		fmt.Println(response)

		if resp.StatusCode == http.StatusOK && response["status"].(string) == "successful" {
			message := response["data"].(map[string]interface{})["message"].(string)
			return message, nil
		}

		time.Sleep(time.Second)
	}

	return "", fmt.Errorf("AI大模型回复超时。")
}
