package main

import "github.com/Go-Golang-Training/microservices-with-go/pkg/server"

func main() {
	s := server.NewServer()
	s.Start()
}
