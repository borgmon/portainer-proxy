package main

import (
	"context"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

func getContainers() ([]types.Container, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	return cli.ContainerList(context.Background(), types.ContainerListOptions{})
}

func findStateByName(name string, containers []types.Container, whitelist []string) string {
	for _, container := range containers {
		for _, containerName := range container.Names {
			if containerName[1:] == name && Contains(whitelist, name) {
				return container.State
			}
		}
	}
	return ""
}

func Contains[T comparable](s []T, e T) bool {
    for _, v := range s {
        if v == e {
            return true
        }
    }
    return false
}

func ParseCSVSlice(str string) []string {
	m := []string{}
	p := strings.Split(str, ",")
	for _, e := range p {
		e = strings.TrimSpace(e)
		m = append(m, e)
	}
	if len(m) == 1 && m[0] == "" {
		return nil
	}
	return m
}

func main() {
	r := gin.Default()
	whitelist := ParseCSVSlice(os.Getenv("WHITELIST"))

	r.GET("/state/:name", func(c *gin.Context) {

		containers, err := getContainers()
		if err != nil {
			c.JSON(500, gin.H{
				"error": err,
			})
			return
		}
		name := c.Param("name")
		if state := findStateByName(name, containers, whitelist); state == "" {
			c.JSON(404, gin.H{
				"message": "Container not found",
			})
		} else {
			c.JSON(200, gin.H{
				"state": state,
			})

		}

	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
