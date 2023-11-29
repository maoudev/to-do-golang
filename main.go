package main

import (
	"github.com/maoudev/todo/internal/config"
	"github.com/maoudev/todo/internal/infraestructure/api"
)

func main() {
	config.GetEnv()
	api.RunServer()

}
