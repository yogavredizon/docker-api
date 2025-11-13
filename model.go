package main

type Container struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type ResponseContainer struct {
	Code       int      `json:"code"`
	Status     string   `json:"status"`
	Message    string   `json:"message"`
	Containers []Server `json:"containers"`
}

type Server struct {
	Host      string      `json:"host"`
	Port      int         `json:"port"`
	Name      string      `json:"name"`
	Container []Container `json:"container"`
}
