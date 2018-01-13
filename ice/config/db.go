package config

import "fmt"

type DatabaseManagerConfig struct {
	Default   string           `yaml:"default"`
	Enabled   []string         `yaml:"enabled"`
	Databases []DatabaseConfig `yaml:"databases"`
}

type DatabaseConfig struct {
	Name     string `yaml:"name"`
	Driver   string `yaml:"driver"`
	DSN      string `yaml:"dsn"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

func (c DatabaseConfig) GetDSN() string {
	if c.DSN != "" {
		return c.DSN
	} else {
		return fmt.Sprintf("host=%s port=%d user=%s passowrd=%s name=%s", c.Host, c.Port, c.User, c.Password, c.DBName)
	}
}

func (c DatabaseConfig) String() string {
	return fmt.Sprintf("name %s driver %s dsn %s", c.Name, c.Driver, c.GetDSN())
}
