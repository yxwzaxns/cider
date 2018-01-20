package config

type Config struct {
	ListenIp   string
	ListenPort string
}

func initConfig() *Config {
	c := new(Config)
	c.ListenIp = "127.0.0.1"
	c.ListenPort = "8080"
	return c
}

func GetConfig() *Config {
	return initConfig()
}
