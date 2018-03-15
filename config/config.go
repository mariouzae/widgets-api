package config

import (
	"log"
	"os"
	"path"
	"time"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Database Database
}

// Represents database server and credentials
type Database struct {
	ServerAddr string
	DbName     string
	Timeout    time.Duration
	Username   string
	Password   string
}

// Read and parse the configuration file
func (c *Config) Read() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dirPath := path.Join(dir, "../src/widgets-api/config/config.toml")
	if _, err := toml.DecodeFile(dirPath, &c); err != nil {
		log.Fatal(err)
	}
}
