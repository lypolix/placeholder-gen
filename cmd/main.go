package main

import (
	"PlaceholderGen/configs"
	"PlaceholderGen/internal/server"
	"flag"
	"log"
)

var confPath = flag.String("conf-path", "./configs/.env", "Path to congig env.")


func main() {
	conf, err := configs.New(*confPath)
	if err != nil {
		log.Fatalln(err)
	}
	server.Run(conf)
}