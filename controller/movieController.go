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
		util.Success(ctx, nil, "综合/B站：<a href='https://jx.jsonplayer.com/player/?url="+data.Message+"'>https://jx.jsonplayer.com/player/?url="+data.Message+"<br/><br/>"+
			"CK：<a href='https://www.ckplayer.vip/jiexi/?url="+data.Message+"'>https://www.ckplayer.vip/jiexi/?url="+data.Message+"<br/><br/>"+
			"YT：<a href='https://jx.yangtu.top/?url="+data.Message+"'>https://jx.yangtu.top/?url="+data.Message+"<br/><br/>"+
			"Player-JY：<a href='https://jx.playerjy.com/?url="+data.Message+"'>https://jx.playerjy.com/?url="+data.Message+"<br/><br/>"+
			"yparse：<a href='https://jx.yparse.com/index.php?url="+data.Message+"'>https://jx.yparse.com/index.php?url="+data.Message+"<br/><br/>"+
			"8090：<a href='https://www.8090g.cn/?url="+data.Message+"'>https://www.8090g.cn/?url="+data.Message+"<br/><br/>"+
			"剖元：<a href='https://www.pouyun.com/?url="+data.Message+"'>https://www.pouyun.com/?url="+data.Message+"<br/><br/>"+
			"虾米：<a href='https://jx.xmflv.com/?url="+data.Message+"'>https://jx.xmflv.com/?url="+data.Message+"<br/><br/>"+
			"全民：<a href='https://43.240.74.102:4433?url="+data.Message+"'>https://43.240.74.102:4433?url="+data.Message+"<br/><br/>"+
			"OK：<a href='https://api.okjx.cc:3389/jx.php?url="+data.Message+"'>https://api.okjx.cc:3389/jx.php?url="+data.Message+"<br/><br/>"+
			"OKJX：<a href='https://okjx.cc/?url="+data.Message+"'>https://okjx.cc/?url="+data.Message+"<br/><br/>"+
			"m1907：<a href='https://im1907.top/?jx="+data.Message+"'>https://im1907.top/?jx="+data.Message+"<br/><br/>"+
			"爱豆：<a href='https://jx.aidouer.net/?url="+data.Message+"'>https://jx.aidouer.net/?url="+data.Message+"<br/><br/>"+
			"猪蹄：<a href='https://jx.iztyy.com/Bei/?url="+data.Message+"'>https://jx.iztyy.com/Bei/?url="+data.Message+"<br/><br/>"+
			"夜幕：<a href='https://www.yemu.xyz/?url="+data.Message+"'>https://www.yemu.xyz/?url="+data.Message+"<br/><br/>"+
			"MAO：<a href='https://www.mtosz.com/m3u8.php?url="+data.Message+"'>https://www.mtosz.com/m3u8.php?url="+data.Message+"<br/><br/>"+
			"M3U8TV：<a href='https://jx.m3u8.tv/jiexi/?url="+data.Message+"'>https://jx.m3u8.tv/jiexi/?url="+data.Message+"<br/><br/>"+
			"铭人云：<a href='https://parse.123mingren.com/?url="+data.Message+"'>https://parse.123mingren.com/?url="+data.Message+"<br/><br/>"+
			"4kdv：<a href='https://jx.4kdv.com/?url="+data.Message+"'>https://jx.4kdv.com/?url="+data.Message+"<br/><br/>"+
			"1717：<a href='https://ckmov.ccyjjd.com/ckmov/?url="+data.Message+"'>https://ckmov.ccyjjd.com/ckmov/?url="+data.Message+"<br/><br/>"+
			"qianqi：<a href='https://api.qianqi.net/vip/?url="+data.Message+"'>https://api.qianqi.net/vip/?url="+data.Message+"<br/><br/>"+
			"laobandq：<a href='https://vip.laobandq.com/jiexi.php?url="+data.Message+"'>https://vip.laobandq.com/jiexi.php?url="+data.Message+"<br/><br/>"+
			"playm3u8：<a href='https://www.playm3u8.cn/jiexi.php?url="+data.Message+"'>https://www.playm3u8.cn/jiexi.php?url="+data.Message+"<br/><br/>"+
			"无名小站：<a href='https://www.administratorw.com/video.php?url="+data.Message+"'>https://www.administratorw.com/video.php?url="+data.Message+"<br/><br/>"+
			"盘古：<a href='https://go.yh0523.cn/y.cy?url="+data.Message+"'>https://go.yh0523.cn/y.cy?url="+data.Message+"<br/><br/>"+
			"Blbo：<a href='https://jx.blbo.cc:4433/?url="+data.Message+"'>https://jx.blbo.cc:4433/?url="+data.Message+"<br/><br/>")
		return
	default:
		util.Fail(ctx, nil, "电影链接格式不正确！")
	}
}
