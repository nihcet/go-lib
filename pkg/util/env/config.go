package configutil

import (
	"github.com/jessevdk/go-flags"
)

var serverConfig ServerConfig

// "env" use to read from .json
// "long" use to read from .env
type ServerConfig struct {
	ProjectId   string `long:"PROJECT_ID" env:"PROJECT_ID"`
	ServiceName string `long:"SERVICE_NAME" env:"SERVICE_NAME"`
	ServiceHost string `long:"SERVICE_HOST" env:"SERVICE_HOST"`
	ServicePort string `long:"SERVICE_PORT" env:"SERVICE_PORT"`
}

func LoadConfig() {
	_, err := flags.Parse(&serverConfig)
	if err != nil {
		panic(err)
	}

}
