package config

import "fmt"

type Database struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"db_name"`
}

func (d *Database) Config() string {
	if d.Password == "" {
		return fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
			d.Host, d.Port, d.Username, d.DbName)
	}

	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		d.Host, d.Port, d.Username, d.Password, d.DbName)
}

func (d *Database) ConfigInfo() string {
	return fmt.Sprintf("%+v", d)
}
