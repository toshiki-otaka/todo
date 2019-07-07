package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

// Config struct
var Config struct {
	Env    string
	Server struct {
		Host string `toml:"host"`
		Port string `toml:"port"`
	}
	Db struct {
		Host string `toml:"host"`
		Port string `toml:"port"`
		Name string `toml:"name"`
	}
}

// Init func
func init() {
	env := os.Getenv("ENV")
	_, err := toml.DecodeFile("config/"+env+".toml", &Config)
	if err != nil {
		panic(err)
	}
	Config.Env = env
}
