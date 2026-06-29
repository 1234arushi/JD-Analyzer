package main

import (
	//Go does not build imports from the repository’s outer folder structure. It starts from the folder containing go.mod.
	"log"

	"github.com/1234arushi/JD-Analyzer/database"
	"github.com/1234arushi/JD-Analyzer/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.ConnectDB()
	router := gin.Default()
	jdscope := router.Group("/jdscope")
	routes.LoadServices(jdscope)
	router.Run(":8080")

}
