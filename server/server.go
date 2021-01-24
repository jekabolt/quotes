package server

import (
	"encoding/json"
)

type Server struct {
	*Config
	JWTSecretKey []byte
}

type Config struct {
	Port         string `env:"SERVER_PORT" envDefault:"8080"`
	JWTSecretKey string `env:"JWT_SECRET_KEY" envDefault:"kek"`
	Debug        bool   `env:"DEBUG" envDefault:"true"`
}

func (c *Config) String() string {
	bs, _ := json.Marshal(c)
	return string(bs)
}

func (c *Config) InitServer() (*Server, error) {
	s := &Server{
		Config:       c,
		JWTSecretKey: []byte(c.JWTSecretKey),
	}
	return s, nil
}
