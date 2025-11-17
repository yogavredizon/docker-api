package main

type Container struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	ServerID  string `json:"serverId"`
	MitraName string `json:"mitraName"`
}

type RequestContainer struct {
	ContainerId string `json:"containerId"`
	ServerId    string `json:"serverId"`
}

type ResponseContainer struct {
	Code       int         `json:"code"`
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	Containers []Container `json:"containers"`
}

type Server struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}
