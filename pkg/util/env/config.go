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

func LoadConfig[T any]() *T {
	var config T
	_, err := flags.Parse(&config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("config: %+v\n", config)
	return &config
}
