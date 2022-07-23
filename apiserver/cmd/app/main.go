package main

import (
	"apiserver/config"
	"apiserver/internal/app"
)

func main() {
	conf := config.ReadConfig()
	app.Run(&conf)
}