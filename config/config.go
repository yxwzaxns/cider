package config

import (
	"os"

	"github.com/yxwzaxns/cider/utils"
)

var Conf *Config

// Config xx
type Config struct {
	ListenIP                  string
	ListenPort                string
	ContinuousIntegrationPath string
	// ContinuousDistribution
	Debug     bool
	AppDbPath string
	LogPath   string

	AuthKey      string
	TokenTimeout int
}

func (c *Config) setConf() {
	if os.Getenv("CIDER_SERVER_IP") == "" {
		c.ListenIP = "127.0.0.1"
		os.Setenv("CIDER_SERVER_IP", c.ListenIP)
	} else {
		c.ListenIP = os.Getenv("CIDER_SERVER_IP")
	}

	if os.Getenv("CIDER_SERVER_PORT") == "" {
		c.ListenPort = "8080"
		os.Setenv("CIDER_SERVER_PORT", c.ListenPort)
	} else {
		c.ListenPort = os.Getenv("CIDER_SERVER_PORT")
	}

	if os.Getenv("CIDER_AUTH_KEY") == "" {
		c.AuthKey = "admin"
		os.Setenv("CIDER_AUTH_KEY", utils.Base64Encode(c.AuthKey))
	} else {
		c.AuthKey = os.Getenv("CIDER_AUTH_KEY")
	}

	if os.Getenv("CIDER_TOKEN_TIMEOUT") == "" {
		c.TokenTimeout = 10
		os.Setenv("CIDER_TOKEN_TIMEOUT", utils.Itoa(c.TokenTimeout))
	} else {
		c.TokenTimeout = utils.Atoi(os.Getenv("CIDER_TOKEN_TIMEOUT"))
	}

	if os.Getenv("CIDER_DB_PATH") == "" {
		// c.AppDbPath = "/usr/local/cider/cider.db"
		c.AppDbPath = "/tmp/cider/cider.db"
		os.Setenv("CIDER_DB_PATH", c.AppDbPath)
	} else {
		c.AppDbPath = os.Getenv("CIDER_DB_PATH")
	}

	if os.Getenv("CIDER_DEBUG") == "false" {
		c.Debug = false
		os.Setenv("CIDER_DEBUG", "false")
	} else {
		c.Debug = true
		os.Setenv("CIDER_DEBUG", "true")
	}

	if os.Getenv("CIDER_LOGPATH") == "" {
		// c.LogPath = "/var/log/cider/cider.log"
		c.LogPath = "/tmp/cider/cider.log"
		os.Setenv("CIDER_LOGPATH", c.LogPath)
	} else {
		c.LogPath = os.Getenv("CIDER_LOGPATH")
	}

	c.ContinuousIntegrationPath = "/var/"

}

func Init() {
	Conf = new(Config)
	Conf.setConf()
	LogInit()
}
