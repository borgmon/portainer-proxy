package controller

import (
	"borgmon.me/state-proxy/docker"
	"borgmon.me/state-proxy/helper"
	"github.com/gin-gonic/gin"
)

type StateHandlerImpl struct {
	DockerService docker.DockerService
	Whitelist     []string
}

func (stateCtl *StateHandlerImpl) GetState(c *gin.Context) {
	name := c.Param("name")
	if !helper.Contains(stateCtl.Whitelist, name) {
		c.JSON(404, gin.H{
			"message": "Container not found",
		})
		return
	}

	container, err := stateCtl.DockerService.GetContainerByName(name)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"state": container.State,
	})
}
