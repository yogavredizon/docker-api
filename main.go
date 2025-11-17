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
	serverFile, err := os.ReadFile("server.json")
	if err != nil {
		log.Println(err)
		return
	}
	var servers map[string]Server
	err = json.Unmarshal(serverFile, &servers)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(servers)
	mitraFile, err := os.ReadFile("mitra.json")
	if err != nil {
		log.Println(err)
		return
	}

	var ListMitra map[string]string
	err = json.Unmarshal(mitraFile, &ListMitra)

	controller := ContainerController{
		Servers: servers,
		Mitra:   ListMitra,
		wg:      &sync.WaitGroup{},
	}

	log.Println(controller.Mitra)

	handler := NewContainerHanlder(controller)

	c := gin.Default()

	containers := c.Group("/containers")
	containers.GET("", LimitRequestGet(), ValidateToken(), handler.QueryContainers)
	containers.Use(LimitRequestPost(), ValidateToken())
	{
		containers.PUT("/start", handler.StartContainers)
		containers.PUT("/stop", handler.StopContainers)
	}

	c.Run(":7788")
}
