package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	RestPort      string `required:"true" envconfig:"REST_PORT"`
	GrpcPort      string `required:"true" envconfig:"GRPC_PORT"`
	JWTSignature  string `required:"true" envconfig:"JWT_SIGNATURE"`
	TokenDuration int    `required:"true" envconfig:"TOKEN_DURATION"`
}

var Global Env = Env{}

func Load() error {
	if err := envconfig.Process("", &Global); err != nil {
		return err
	}
	return nil
}
