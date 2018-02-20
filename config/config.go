package config

import "os"

// Config xx
type Config struct {
	ListenIP                  string
	ListenPort                string
	ContinuousIntegrationPath string
	// ContinuousDistribution
	AppDbPath string
}

func (c *Config) Init() {
	c.ListenIP = "127.0.0.1"
	c.ListenPort = "8080"
	c.ContinuousIntegrationPath = "/var/"
	c.AppDbPath = "/etc/cider/cider.db"
}
func (c *Config) InitEnvironmentVariable() {
	os.Setenv("ContinuousIntegrationPath", c.ContinuousIntegrationPath)
}
