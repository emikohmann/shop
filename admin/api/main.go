package main

import (
	"api/docker"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	// create server
	server := gin.Default()

	// cors
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "X-Requested-With", "Content-Type", "Accept"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// docker client
	client, err := docker.NewClient()
	if err != nil {
		panic(fmt.Sprintf("error initializing docker client: %v", err))
	}

	// docker handlers
	dockerHandlers := docker.NewHandlers(client)

	// routes
	server.GET("/versions", dockerHandlers.ServiceVersions)
	server.GET("/docker/containers", dockerHandlers.ListContainers)
	server.POST("/docker/build", dockerHandlers.Build)
	server.POST("/docker/start", dockerHandlers.Start)
	server.POST("/docker/stop", dockerHandlers.Stop)

	// run server
	if err := server.Run(":9999"); err != nil {
		panic(fmt.Sprintf("err running server: %v", err))
	}
}
