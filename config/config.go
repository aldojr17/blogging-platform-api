package config

import (
	log "blogging-platform-api/logger"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Database *Database `mapstructure:"db"`
	Redis    *Redis    `mapstructure:"redis"`
}

var (
	instance *Config
	once     sync.Once
	mutex    sync.Mutex
)

func initConfig(env string) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	viper.Set("env", env)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed read file config : %s", err.Error())
	}

	if err := viper.UnmarshalKey(viper.GetString("env"), &instance); err != nil {
		log.Fatalf("unable to decode into config struct, %v", err)
	}
}

func GetConfig(env string) *Config {
	once.Do(func() {
		initConfig(env)
	})
	mutex.Lock()
	defer mutex.Unlock()
	return instance
}
