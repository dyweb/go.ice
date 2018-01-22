package config

import "fmt"

type DatabaseManagerConfig struct {
	Default   string           `yaml:"default"`
	Databases []DatabaseConfig `yaml:"databases"`
}

var EmptyDatabaseConfig = DatabaseConfig{}

type DatabaseConfig struct {
	Name     string `yaml:"name"`
	Adapter  string `yaml:"adapter"`
	DSN      string `yaml:"dsn"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

func (c DatabaseConfig) String() string {
	return fmt.Sprintf("name=%s adapter=%s host=%s port=%d user=%s dbname=%s",
		c.Name, c.Adapter, c.Host, c.Port, c.User, c.DBName)
}
