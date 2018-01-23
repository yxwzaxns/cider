package config

// Config xx
type Config struct {
	ListenIP   string
	ListenPort string
}

func initConfig() *Config {
	c := new(Config)
	c.ListenIP = "127.0.0.1"
	c.ListenPort = "8080"
	return c
}

// GetConfig xx
func GetConfig() *Config {
	return initConfig()
}
