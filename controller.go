package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

var urlPattern = "containers"

type ContainerController struct {
	Servers []Server
	wg      *sync.WaitGroup
}

func (c *ContainerController) GetContainers() ([]Server, error) {
	conChan := make(chan Server, len(c.Servers))

	for _, server := range c.Servers {
		s := server
		c.wg.Add(1)

		url := fmt.Sprintf("http://%v:%v/%v/%v", s.Host, s.Port, s.Name, urlPattern)
		go func(s Server, url string) {
			defer c.wg.Done()

			containers, err := fetch(url)
			if err != nil {
				log.Println(err)
				return
			}

			result := Server{
				Host:      s.Host,
				Port:      s.Port,
				Name:      s.Name,
				Container: containers,
			}

			conChan <- result

		}(s, url)
	}

	c.wg.Wait()
	close(conChan)

	var servers []Server
	for s := range conChan {
		servers = append(servers, s)

	}

	return servers, nil
}

func fetch(url string) ([]Container, error) {
	client := http.Client{}
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

	var containers map[string][]Container
	err = json.NewDecoder(res.Body).Decode(&containers)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res.Body.Close()

	return containers["containers"], nil
}

func (c *ContainerController) StartContainer(s Server, id string) (map[string]string, error) {
	url := fmt.Sprintf("http://%s:%v/%s/%s/start/", s.Host, s.Port, s.Name, urlPattern)

	client := http.Client{}
	req, err := http.NewRequest("PUT", url+id, nil)
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

	var result map[string]string
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res.Body.Close()

	return result, nil
}

func (c *ContainerController) StopContainer(s Server, id string) (map[string]string, error) {
	url := fmt.Sprintf("http://%s:%v/%s/%s/stop/", s.Host, s.Port, s.Name, urlPattern)

	client := http.Client{}
	req, err := http.NewRequest("PUT", url+id, nil)
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

	var result map[string]string
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res.Body.Close()

	return result, nil
}
