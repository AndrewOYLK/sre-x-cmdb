package api

import (
	"github.com/AndrewOYLK/ou-cmdb/pkg/server"
	"github.com/gin-gonic/gin"
)

type Server struct {
	// 服务接口
	server.IServer
	// Service功能集合（模型、属性）
}

func NewServer(iserver server.IServer) *Server {
	server := &Server{
		iserver,
	}

	gRouter := server.GinEngine().Group("/api/v1apha1")
	gRouter.GET("/model", server.getModel)

	return server
}

func (api *Server) getModel(ctx *gin.Context) {
}
