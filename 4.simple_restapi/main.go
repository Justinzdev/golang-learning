package main

import (
	"myrestapi/4.simple_restapi/database"
	"myrestapi/4.simple_restapi/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	database.InitializeDB(dbUser, dbPassword, dbHost, dbPort, dbName)
	
	r := gin.Default()
	routerV1 := r.Group("/api/v1")
	routes.SetupPingRoutes(routerV1)
	routes.SetupUserRoutes(routerV1)
	r.Run(":4000")
}
