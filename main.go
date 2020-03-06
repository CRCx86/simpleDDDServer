package main

import (
	config2 "simpleDDDServer/config"
	"simpleDDDServer/repository"
	server2 "simpleDDDServer/server"
	"simpleDDDServer/service"
)

func main() {
	config := config2.NewConfig()

	db, err := config.Connect(config)

	if err != nil {
		panic(err)
	}

	userRepository := repository.NewUserRepository(db)

	userService := service.NewUserService(config, userRepository)

	server := server2.NewServer(config, userService)

	server.Run()
}