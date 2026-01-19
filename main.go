package main

import "fmt"

type Server struct {
	Port int
	Env  string
}

func (s Server) IsProduction() bool {
	// return s.Env == "production"
	return s.Env == "staging"
}

func (s *Server) SetPort(port int) {
	s.Port = port
}

func main() {
	server := Server{}
	server.SetPort(8080)
	server.Env = "production"
	fmt.Println(server.IsProduction())
	fmt.Println(server.Port)
}
