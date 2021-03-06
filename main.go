package main

import (
	"log"

	config "github.com/Alibi-S/todoProject/configs"
	routes "github.com/Alibi-S/todoProject/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()
	router := gin.Default()

	routes.Routes(router)

	log.Fatal(router.Run(":8080"))

}
