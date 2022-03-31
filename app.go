package main

import (
	"os"

	"borgmon.me/state-proxy/controller"
	"borgmon.me/state-proxy/docker"
	"borgmon.me/state-proxy/helper"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	helper.PanicIfError(err)

	dockerService := &docker.DockerServiceImpl{
		Client: cli,
	}

	stateHandler := &controller.StateHandlerImpl{
		DockerService: dockerService,
		Whitelist:     helper.ParseCSVSlice(os.Getenv("WHITELIST")),
	}

	r.GET("/state/:name", stateHandler.GetState)

	r.SetTrustedProxies(nil)
	r.Run(helper.GetPort(os.Getenv("PORT")))
}
