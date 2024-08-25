package initialize

import (
	"blogging-platform-api/config"
	log "blogging-platform-api/logger"
	"blogging-platform-api/utils"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Application struct {
	Config   *config.Config
	Database *gorm.DB
	Redis    *redis.Client
}

func InitApp() *Application {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env := utils.GetEnv("ENV", "development")

	cfg := config.GetConfig(env)

	app := &Application{
		Config:   cfg,
		Database: initializeDB(cfg),
		Redis:    initializeRedis(cfg),
	}

	return app
}

func initializeDB(cfg *config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.Database.Config()), &gorm.Config{})

	if err != nil {
		log.Fatalf("error initialize database : %s", err.Error())
	}

	log.Info("Database:\n",
		"config", cfg.Database.ConfigInfo(),
	)

	return db
}

func initializeRedis(cfg *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Addr(),
		DB:   cfg.Redis.Index(),
	})

	log.Info("Redis:\n",
		"config", cfg.Redis.ConfigInfo(),
	)

	return client
}
