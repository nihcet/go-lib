package configutil

import (
	"github.com/jessevdk/go-flags"
)

var serverConfig ServerConfig

type ServerConfig struct {
	ProjectId   string `env:"PROJECT_ID"`
	ServiceName string `env:"SERVICE_NAME"`
	ServiceHost string `env:"SERVICE_HOST"`
	ServicePort string `env:"SERVICE_PORT"`
}

func LoadConfig() {
	_, err := flags.Parse(&serverConfig)
	if err != nil {
		panic(err)
	}

}
