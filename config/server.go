package config

import "fmt"

type ServerConfig struct {
	Host string `required:"true"`
	Port int    `required:"true"`
}

func (s *ServerConfig) GetAddress() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
