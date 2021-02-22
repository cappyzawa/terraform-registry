package main

import (
	"os"

	"github.com/cappyzawa/terraform-registry/internal/server"
)

type cli struct {
	port       string
	configFile string
	pidFile    string
}

var (
	version = "dev"
)

func (c *cli) Run(args []string) {

	if len(args) != 0 && args[0] == "version" {
		print(version)
		return
	}

	if c.port == "" {
		c.port = "8080"
	}
	if c.configFile == "" {
		c.configFile = "./config.yaml"
	}

	opts := &server.Opt{
		Port:       c.port,
		ConfigPATH: c.configFile,
		PIDPATH:    c.pidFile,
	}

	server.Run(opts)
}

func main() {
	c := &cli{
		port:       os.Getenv("PORT"),
		configFile: os.Getenv("CONFIG_FILE"),
		pidFile:    os.Getenv("PID_FILE"),
	}
	c.Run(os.Args[1:])
}
