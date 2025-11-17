package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var urlPattern = "containers"

type ContainerController struct {
	Servers map[string]Server
	Mitra   map[string]string
	wg      *sync.WaitGroup
}

func (c *ContainerController) GetContainers() ([]Container, error) {
	conChan := make(chan []Container, len(c.Servers))
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	for _, server := range c.Servers {
		s := server
		c.wg.Add(1)

		go func(s Server, mitra map[string]string) {
			defer c.wg.Done()
			containers, err := fetch(&client, s, mitra)
			if err != nil {
				log.Println(err)
				return
			}

			conChan <- containers

		}(s, c.Mitra)
	}

	c.wg.Wait()
	close(conChan)

	var containers []Container
	for s := range conChan {
		containers = append(containers, s...)
	}

	return containers, nil
}

func fetch(client *http.Client, s Server, mitra map[string]string) ([]Container, error) {
	url := fmt.Sprintf("http://%v:%v/%v", s.Host, s.Port, urlPattern)
	req, err := http.NewRequest("GET", url, nil)

	req.Header.Set("X-AUTH-TOKEN", os.Getenv("X_AUTH_TOKEN"))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var result map[string][]Container
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var containers []Container
	for _, c := range result["containers"] {
		if mitraName, ok := mitra[c.Name]; ok {
			c.MitraName = mitraName
		}

		containers = append(containers, c)
	}

	return containers, nil
}

func (c *ContainerController) StartContainer(container RequestContainer) (map[string]string, error) {
	// check if server has a given container
	server, ok := c.Servers[container.ServerId]
	if !ok {
		return nil, errors.New("server id tidak valid")
	}

	// make a request to the container server
	client := http.Client{}
	url := fmt.Sprintf("http://%s:%v/%s/start/", server.Host, server.Port, urlPattern)
	req, err := http.NewRequest("PUT", url+container.ContainerId, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header.Set("X-AUTH-TOKEN", os.Getenv("X_AUTH_TOKEN"))
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var result map[string]string
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res.Body.Close()

	return result, nil
}

func (c *ContainerController) StopContainer(container RequestContainer) (map[string]string, error) {

	// check if server has a given container
	server, ok := c.Servers[container.ServerId]
	if !ok {
		return nil, errors.New("server id tidak valid")
	}

	// make a request to the container server
	client := http.Client{}
	url := fmt.Sprintf("http://%s:%v/%s/stop/", server.Host, server.Port, urlPattern)
	req, err := http.NewRequest("PUT", url+container.ContainerId, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header.Set("X-AUTH-TOKEN", os.Getenv("X_AUTH_TOKEN"))
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var result map[string]string
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res.Body.Close()

	return result, nil
}
