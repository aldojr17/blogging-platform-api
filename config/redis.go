package config

import (
	"fmt"
	"time"
)

type Redis struct {
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	DefaultTTL int    `mapstructure:"default_ttl"`
}

func (r *Redis) Addr() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}

func (r *Redis) Index() int {
	return 0
}

func (r *Redis) ConfigInfo() string {
	return fmt.Sprintf("%+v", r)
}

func (r *Redis) GetDefaultTTL() time.Duration {
	return time.Duration(r.DefaultTTL) * time.Minute
}
