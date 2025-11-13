package main

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	file, err := os.ReadFile("server.json")
	if err != nil {
		log.Println(err)
		return
	}
	var servers []Server
	err = json.Unmarshal(file, &servers)
	if err != nil {
		log.Println(err)
		return
	}

	controller := ContainerController{
		Servers: servers,
		wg:      &sync.WaitGroup{},
	}

	handler := NewContainerHanlder(controller)

	c := gin.Default()
	c.GET("/containers", LimitRequestGet(), ValidateToken(), handler.QueryContainers)

	containers := c.Group("/containers")
	containers.Use(LimitRequestPost(), ValidateToken())
	{
		containers.PUT("/start", handler.StartContainers)
		containers.PUT("/stop", handler.StopContainers)
	}

	c.Run(":7788")
}
