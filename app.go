package main

import (
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

var (
	client          = resty.New()
	portainerClient = &PortainerClient{
		Client:   *client.SetHostURL(os.Getenv("PORTAINER_URL")),
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
	}
	filterList = os.Getenv("FILTER")
)

func main() {
	router := gin.Default()

	router.GET("/status", GetStatus)
	router.Run()
}

func GetStatus(c *gin.Context) {
	m, err := portainerClient.GetStatus(false)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "can't get status"})
	} else {
		var filtered []string

		for k, v := range m {
			if strings.Contains(filterList, k) && v == "running" {
				filtered = append(filtered, k)
			}
		}

		c.JSON(200, gin.H{"online": filtered})
	}

}
