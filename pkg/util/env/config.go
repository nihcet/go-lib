package configutil

import (
	"github.com/jessevdk/go-flags"
)

var serverConfig ServerConfig

type ServerConfig struct {
	ProjectId   string `env:"project_id"`
	ServiceName string `env:"service_name"`
	ServiceHost string `env:"service_host"`
	ServicePort string `env:"service_port"`
}

func LoadConfig() {
	_, err := flags.Parse(&serverConfig)
	if err != nil {
		panic(err)
	}

}
