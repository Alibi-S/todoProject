package main

import (
	"log"

	config "github.com/Alibi-S/todoProject/config"
	routes "github.com/Alibi-S/todoProjectroutes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()
	router := gin.Default()

	routes.Routes(router)

	log.Fatal(router.Run(":4747"))

}
