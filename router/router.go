package router

import (
	"github.com/gin-gonic/gin"
	"im/middlewares"
	"im/service"
)

func Router() *gin.Engine {
	r := gin.Default()

	// 用户登录
	r.POST("/login", service.Login)
	// 发送验证码
	r.POST("/send/code", service.SendCode)
	// 用户注册
	r.POST("/register", service.Register)

	auth := r.Group("/u", middlewares.AuthCheck())

	// UserDetail
	auth.GET("/user/detail", service.UserDetail)
	// 查询指定用户的个人信息
	auth.GET("/usre/query", service.UserQuery)
	// 发送、接受消息
	auth.GET("/websocket/message", service.WebsocketMessage)
	// 聊天纪录列表
	auth.GET("/chat/list", service.ChatList)

	// 添加用户
	auth.POST("/user/add", service.UserAdd)
	// 删除用户
	auth.POST("/user/delete", service.UserDelete)

	return r
}
