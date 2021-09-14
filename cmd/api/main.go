package main

import (
	"fmt"

	"github.com/AndrewOYLK/ou-cmdb/pkg/server"
	"github.com/AndrewOYLK/ou-cmdb/pkg/server/api"
	"github.com/AndrewOYLK/ou-cmdb/pkg/utils"
)

var (
	metricsHost       = "0.0.0.0"
	metricsPort int32 = 8383
)

func main() {
	serverHost := utils.Getenv("server_host", "0.0.0.0")
	serverPort := utils.Getenv("server_port", "8080")

	ginServer := server.NewGinServer(fmt.Sprintf("%s:%s", serverHost, serverPort))
	apiServer := api.NewServer(ginServer)
	if err := apiServer.Run(); err != nil {
		panic(err)
	}
}
