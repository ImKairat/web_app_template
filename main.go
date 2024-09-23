package main

import (
	"fmt"
	"log"
	"ses_back/internal/service"
)

func main() {
	//apps.App()
	err := service.Connect()
	if err != nil {
		log.Fatal(err)
	}

	res, err3 := service.DB.Query("SELECT * FROM bulls;")
	if err3 != nil {
		panic(err3)
	} else {
		fmt.Println(res.Columns())
	}

	err2 := service.Disconnect()
	if err2 != nil {
		log.Fatal(err2)
	}
}
