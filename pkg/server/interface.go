package server

import "github.com/gin-gonic/gin"

type IServer interface {
	Run() error
	Stop() error
	GinEngine() *gin.Engine
}

type GinServer struct {
	addrs  []string
	engine *gin.Engine
}

func (g *GinServer) Run() error {
	if err := g.engine.Run(g.addrs...); err != nil {
		return err
	}
	return nil
}

func (g *GinServer) Stop() error {
	return nil
}

func (g *GinServer) GinEngine() *gin.Engine {
	return g.engine
}

func NewGinServer(addr string) IServer {
	return &GinServer{
		addrs:  []string{addr},
		engine: gin.Default(),
	}
}
