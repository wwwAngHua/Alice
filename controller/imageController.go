package controller

import (
	"Alice/model"
	"Alice/util"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var msg = "您好，我是一个免费图床机器人！您可以点击右下角的上传文件按钮，上传图片，我会给您一个永久有效的图片URL地址。<br/><br/><a href=\"https://img.t1ykf.com/\">进一步了解我们！</a>"

func ImageChat(ctx *gin.Context) {
	var data model.Chat
	if err := ctx.ShouldBindJSON(&data); err != nil {
		util.Fail(ctx, nil, "请求体错误！")
		return
	}

	switch {
	case strings.HasPrefix(data.Message, "/start"):
		util.Success(ctx, nil, msg)
		return
	default:
		util.Fail(ctx, nil, msg)
	}
}

func ImageUpload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		util.Fail(ctx, nil, "图片上传失败："+err.Error())
		return
	}

	// 生成唯一的文件名，保留原始扩展名
	ext := filepath.Ext(file.Filename)
	randomFileName := generateRandomFileName(ext)

	// 保存上传的文件到本地
	workDir, _ := os.Getwd()
	filePath := workDir + "/upload/image/" + randomFileName
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		util.Fail(ctx, nil, "保存图片失败："+err.Error())
		return
	}

	// 发送文件到指定接口
	url := "https://img.t1ykf.com/upload"
	result, err := sendFile(url, filePath)
	if err != nil {
		util.Fail(ctx, nil, "图片上传失败："+err.Error())
		return
	}

	util.Success(ctx, nil, "图片上传成功！图片URL为：<a href=\"https://img.t1ykf.com"+result+"\">https://img.t1ykf.com"+result+"</a>")
}

func generateRandomFileName(extension string) string {
	uuidString := uuid.New().String()
	// 移除中间的横杠，确保文件名合法
	uuidString = strings.ReplaceAll(uuidString, "-", "")
	return uuidString + extension
}

func sendFile(url string, filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	part, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		return "", err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return "", err
	}

	writer.Close()

	req, err := http.NewRequest("POST", url, &body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode == http.StatusOK {
		var response []struct {
			Src string `json:"src"`
		}
		if err := json.Unmarshal(responseBody, &response); err != nil {
			return "", err
		}
		if len(response) > 0 {
			return response[0].Src, nil
		}
	} else {
		var errorResponse struct {
			Error string `json:"error"`
		}
		if err := json.Unmarshal(responseBody, &errorResponse); err != nil {
			return "", err
		}
		return "", errors.New(errorResponse.Error)
	}

	return "", errors.New("unexpected response")
}
