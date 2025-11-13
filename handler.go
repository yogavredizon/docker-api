package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ContainerHandler struct {
	ContainerController ContainerController
}

func NewContainerHanlder(c ContainerController) *ContainerHandler {
	return &ContainerHandler{ContainerController: c}
}

func (cc *ContainerHandler) QueryContainers(c *gin.Context) {
	containers, err := cc.ContainerController.GetContainers()
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, ResponseContainer{
			Code:       http.StatusInternalServerError,
			Status:     "Internal Server Error",
			Message:    "Gagal mendapatkan container",
			Containers: nil,
		})
		return
	}

	c.IndentedJSON(http.StatusOK, ResponseContainer{
		Code:       http.StatusOK,
		Status:     "OK",
		Message:    "Berhasil mendapatkan container",
		Containers: containers,
	})
}

func (cc *ContainerHandler) StartContainers(c *gin.Context) {
	var request map[string]string
	err := json.NewDecoder(c.Request.Body).Decode(&request)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, ResponseContainer{
			Code:       http.StatusInternalServerError,
			Status:     "Internal Server Error",
			Message:    "Gagal mendapatkan container",
			Containers: nil,
		})
		return
	}

	port, err := strconv.Atoi(request["port"])
	if err != nil {
		log.Println(err)
		return
	}
	_, err = cc.ContainerController.StartContainer(Server{
		Host: request["host"],
		Port: port,
		Name: request["name"],
	}, request["containerId"])

	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, ResponseContainer{
			Code:       http.StatusInternalServerError,
			Status:     "Internal Server Error",
			Message:    "Gagal mendapatkan container",
			Containers: nil,
		})
		return
	}

	c.IndentedJSON(http.StatusOK, ResponseContainer{
		Code:       http.StatusOK,
		Status:     "OK",
		Message:    "Berhasil menjalankan container container",
		Containers: nil,
	})

}

func (cc *ContainerHandler) StopContainers(c *gin.Context) {
	var request map[string]string
	err := json.NewDecoder(c.Request.Body).Decode(&request)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, ResponseContainer{
			Code:       http.StatusInternalServerError,
			Status:     "Internal Server Error",
			Message:    "Gagal mendapatkan container p" + err.Error(),
			Containers: nil,
		})
		return
	}

	port, err := strconv.Atoi(request["port"])
	if err != nil {
		log.Println(err)
		return
	}

	_, err = cc.ContainerController.StopContainer(Server{
		Host: request["host"],
		Port: port,
		Name: request["name"],
	}, request["containerId"])

	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, ResponseContainer{
			Code:       http.StatusInternalServerError,
			Status:     "Internal Server Error",
			Message:    "Gagal mendapatkan container",
			Containers: nil,
		})
		return
	}

	c.IndentedJSON(http.StatusOK, ResponseContainer{
		Code:       http.StatusOK,
		Status:     "OK",
		Message:    "Berhasil menghentikan container container",
		Containers: nil,
	})
}
