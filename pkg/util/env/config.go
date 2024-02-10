package configutil

import (
	"fmt"

	"github.com/jessevdk/go-flags"
)

type ServerConfig struct {
	ProjectId   string `env:"project_id"`
	ServiceName string `env:"service_name"`
	ServiceHost string `env:"service_host"`
	ServicePort string `env:"service_port"`
}

func LoadConfig(config any) {
	_, err := flags.Parse(&config)
	fmt.Printf("config: %+v\n", config)
	if err != nil {
		panic(err)
	}
}
