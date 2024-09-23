package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hi(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		gin.H{"message": "Hi"},
	)
}

type user struct {
	name string
	age  int32
}

func GetUsers(ctx *gin.Context) {
	users := map[string]user{
		"user1": {
			name: "Kairat", age: 25,
		},
		"user2": {
			name: "Baiysh", age: 32,
		},
		"user3": {
			name: "Guldaana", age: 24,
		},
		"user4": {
			name: "Kalys", age: 30,
		},
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{"Users": gin.MIMEJSON, "user": users["user4"].name},
	)
}
