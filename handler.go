package main

import (
	"encoding/json"
	"log"
	"net/http"

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
	var container RequestContainer
	err := json.NewDecoder(c.Request.Body).Decode(&container)
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

	_, err = cc.ContainerController.StartContainer(container)

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
	var container RequestContainer
	err := json.NewDecoder(c.Request.Body).Decode(&container)
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

	_, err = cc.ContainerController.StopContainer(container)

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
