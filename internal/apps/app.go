package apps

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"ses_back/internal/service"
)

const ServerPort = ":8080"

func App() {
	server := gin.Default()

	// Services:
	go server.GET(
		"/",
		service.Hi,
	)

	go server.GET(
		"/user1",
		service.GetUsers,
	)

	go server.POST(
		"/user",
		service.GetUsers,
	)

	//---------------------------------------------------------------//
	fmt.Printf("Server run on http://127.0.0.1%s\n", ServerPort)

	err := server.Run(ServerPort)
	if err != nil {
		log.Fatal("Server run error: ", err)
	}
}
