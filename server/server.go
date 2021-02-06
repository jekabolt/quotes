package server

import (
	"encoding/json"
)

type Server struct {
	*Config
	JWTSecretKey  []byte
	AuthSecretKey []byte
}

type Config struct {
	Port          string `env:"SERVER_PORT" envDefault:"8080"`
	JWTSecretKey  string `env:"JWT_SECRET_KEY" envDefault:"kek"`
	AuthSecretKey string `env:"JWT_SECRET_KEY" envDefault:"kekjejcipher1337"`
	Debug         bool   `env:"DEBUG" envDefault:"true"`
}

func (c *Config) String() string {
	bs, _ := json.Marshal(c)
	return string(bs)
}

func (c *Config) InitServer() (*Server, error) {
	s := &Server{
		Config:        c,
		JWTSecretKey:  []byte(c.JWTSecretKey),
		AuthSecretKey: []byte(c.AuthSecretKey),
	}
	return s, nil
}
