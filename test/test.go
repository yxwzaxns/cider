package main

import "fmt"

type Config struct {
	listenIp   string
	listenPort string
}

func initConfig() *Config {
	c := new(Config)
	c.listenIp = "127.0.0.1"
	c.listenPort = "8080"
	return c
}

func GetConfig() *Config {
	c := initConfig()
	fmt.Println(c.listenIp)
	return c
}

func main() {
	c := GetConfig()
	println(c.listenIp)
}
