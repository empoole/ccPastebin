package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Save(context *gin.Context) {
	fmt.Println("hi from controller land")
}