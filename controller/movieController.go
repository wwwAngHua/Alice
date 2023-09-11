package controller

import (
	"Alice/model"
	"Alice/util"
	"strings"

	"github.com/gin-gonic/gin"
)

func MovieChat(ctx *gin.Context) {
	var data model.Chat
	if err := ctx.ShouldBindJSON(&data); err != nil {
		util.Fail(ctx, nil, "请求体错误！")
		return
	}

	switch {
	case strings.HasPrefix(data.Message, "/start"):
		util.Success(ctx, nil, "您好，我是一个电影解析机器人。支持解析爱奇艺、腾讯视频、哔哩哔哩等电影，请向我发送一个电影链接。")
		return
	case strings.HasPrefix(data.Message, "https://"):
		util.Success(ctx, nil, "综合/B站：https://jx.jsonplayer.com/player/?url="+data.Message+"<br/><br/>"+
			"CK：https://www.ckplayer.vip/jiexi/?url="+data.Message+"<br/><br/>"+
			"YT：https://jx.yangtu.top/?url="+data.Message+"<br/><br/>"+
			"Player-JY：https://jx.playerjy.com/?url="+data.Message+"<br/><br/>"+
			"yparse：https://jx.yparse.com/index.php?url="+data.Message+"<br/><br/>"+
			"8090：https://www.8090g.cn/?url="+data.Message+"<br/><br/>"+
			"剖元：https://www.pouyun.com/?url="+data.Message+"<br/><br/>"+
			"虾米：https://jx.xmflv.com/?url="+data.Message+"<br/><br/>"+
			"全民：https://43.240.74.102:4433?url="+data.Message+"<br/><br/>"+
			"OK：https://api.okjx.cc:3389/jx.php?url="+data.Message+"<br/><br/>"+
			"OKJX：https://okjx.cc/?url="+data.Message+"<br/><br/>"+
			"m1907：https://im1907.top/?jx="+data.Message+"<br/><br/>"+
			"爱豆：https://jx.aidouer.net/?url="+data.Message+"<br/><br/>"+
			"猪蹄：https://jx.iztyy.com/Bei/?url="+data.Message+"<br/><br/>"+
			"夜幕：https://www.yemu.xyz/?url="+data.Message+"<br/><br/>"+
			"MAO：https://www.mtosz.com/m3u8.php?url="+data.Message+"<br/><br/>"+
			"M3U8TV：https://jx.m3u8.tv/jiexi/?url="+data.Message+"<br/><br/>"+
			"铭人云：https://parse.123mingren.com/?url="+data.Message+"<br/><br/>"+
			"4kdv：https://jx.4kdv.com/?url="+data.Message+"<br/><br/>"+
			"1717：https://ckmov.ccyjjd.com/ckmov/?url="+data.Message+"<br/><br/>"+
			"qianqi：https://api.qianqi.net/vip/?url="+data.Message+"<br/><br/>"+
			"laobandq：https://vip.laobandq.com/jiexi.php?url="+data.Message+"<br/><br/>"+
			"playm3u8：https://www.playm3u8.cn/jiexi.php?url="+data.Message+"<br/><br/>"+
			"无名小站：https://www.administratorw.com/video.php?url="+data.Message+"<br/><br/>"+
			"盘古：https://go.yh0523.cn/y.cy?url="+data.Message+"<br/><br/>"+
			"Blbo：https://jx.blbo.cc:4433/?url="+data.Message)
		return
	default:
		util.Fail(ctx, nil, "电影链接格式不正确！")
	}
}
