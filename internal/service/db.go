package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetEvents(ctx *gin.Context) {
	data, err := DBselect("")
	if err != nil {
		log.Panic(err)
	}
	ctx.IndentedJSON(http.StatusOK, data)
}
