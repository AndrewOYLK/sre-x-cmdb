package main

import (
	"github.com/AndrewOYLK/ou-cmdb/handler"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/AndrewOYLK/ou-cmdb/docs"
)

func main() {
	engine := gin.Default()
	engine.GET("/apidoc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := engine.Group("/api/v1")
	// Model
	v1.POST("/model", handler.CreateModel)
	v1.DELETE("/model/:id", handler.DeleteModel)
	v1.PUT("/model/:id", handler.UpdateModel)
	v1.GET("/model/:id", handler.GetModel)
	v1.GET("/models", handler.ListModels)

	// Attribute
	v1.POST("/attribute", handler.CreateAttribute)
	v1.DELETE("/attribute/:id", handler.DeleteAttribute)
	v1.PUT("/attribute/:id", handler.UpdateAttribute)
	v1.GET("/attribute/:id", handler.GetAttribute)
	v1.GET("/attributes", handler.ListAttribute)

	// UniqueAttrs
	v1.POST("/uniqueattrs", handler.CreateUniqueAttrs)
	v1.DELETE("/uniqueattrs/:id", handler.DeleteUniqueAttrs)
	v1.PUT("/uniqueattrs/:id", handler.UpdateUniqueAttrs)
	v1.GET("/uniqueattrs", handler.ListUniqueAttrs)

	// Entity(Entity & Value)
	v1.POST("/entity", handler.CreateEntity)
	v1.DELETE("/entity/:id", handler.DeleteEntity)
	v1.GET("/entities", handler.ListEntity)

	// LinkType
	v1.POST("/linktype", handler.CreateLinkType)
	v1.DELETE("/linktype/:id", handler.DeleteLinkType)
	v1.PUT("/linktype/:id", handler.UpdateLinkType)
	v1.GET("/linktypes", handler.ListLinkTypes)

	// LinkModel
	v1.POST("/linkmodel", handler.CreateLinkModel)
	v1.DELETE("/linkmodel/:id", handler.DeleteLinkModel)
	v1.GET("/linkmodels", handler.ListLinkModels)

	// LinkEntity
	v1.POST("/linkentity", handler.CreateLinkEntity)
	v1.DELETE("/linkentity/:id", handler.DeleteLinkEntity)

	engine.Run("0.0.0.0:8080")
}
