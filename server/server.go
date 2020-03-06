package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	config2 "simpleDDDServer/config"
	"simpleDDDServer/service"
)

type Server struct {
	config *config2.Config
	userService *service.UserService
}

func NewServer(config *config2.Config, service *service.UserService) *Server {
	return &Server{
		config:      config,
		userService: service,
	}
}

func (server *Server) Run() {
	httpServer := &http.Server{
		Addr:    server.config.Port,
		Handler: server.Handler(),
	}

	fmt.Println("Listen at port: " + server.config.Port)
	err := httpServer.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func (server *Server) Handler() http.Handler  {
	mux := http.NewServeMux()

	mux.HandleFunc("/users", server.users)

	return mux
}

func (server *Server) users(writer http.ResponseWriter, request *http.Request) {
	users := server.userService.FindAll()
	bytes, _ := json.Marshal(users)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(bytes)
}
