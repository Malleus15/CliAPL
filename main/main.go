package main

import (
	"awesomeProject2/api"
	"net/http"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Duration(50) * time.Second}
	comm:= api.NewCommunication(c)
	MenuHome(&comm)

}