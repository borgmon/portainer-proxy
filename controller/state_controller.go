package controller

import (
	"github.com/gin-gonic/gin"
)

type StateController interface {
	GetState(*gin.Context)
}
