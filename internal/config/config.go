package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/ryutah/oapi-codegen-sample/internal/xerror"
)

type Config struct {
	Port string
}

func Load() (*Config, error) {
	var c Config
	if err := envconfig.Process("hello", &c); err != nil {
		return nil, xerror.New(xerror.Internal, "failed to load config", xerror.WithCause(err))
	}
	return &c, nil
}
