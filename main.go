package main

import (
	"app/db"
	"app/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Welcome to my blog api hehehe!")
	db.ConnectTODB()
	router := gin.Default()
	routes.RegistedRoute(router)
	router.Run(":8000")
}
