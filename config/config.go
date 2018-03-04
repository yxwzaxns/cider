package config

import (
	"cider/utils"
	"os"
)

// Config xx
type Config struct {
	ListenIP                  string
	ListenPort                string
	ContinuousIntegrationPath string
	// ContinuousDistribution
	AppDbPath string
}

func (c *Config) Init() {
	if os.Getenv("CIDER_SERVER_IP") == "" {
		c.ListenIP = "127.0.0.1"
		os.Setenv("CIDER_SERVER_IP", "127.0.0.1")
	}
	if os.Getenv("CIDER_SERVER_PORT") == "" {
		c.ListenPort = "8080"
		os.Setenv("CIDER_SERVER_PORT", "8080")
	}
	if os.Getenv("CIDER_AUTH_KEY") == "" {
		os.Setenv("CIDER_AUTH_KEY", utils.Base64Encode("admin"))
	}

	c.ContinuousIntegrationPath = "/var/"
	c.AppDbPath = "/etc/cider/cider.db"
}
