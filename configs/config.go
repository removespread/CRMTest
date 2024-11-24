package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ServerConfig struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	DatabaseConfig struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"db_name"`
		Dsn      string `yaml:"dsn"`
	} `yaml:"database"`
	JWTConfig struct {
		SecretKey  string `yaml:"secret_key"`
		Expiration int64  `yaml:"expiration"`
	} `yaml:"jwt"`
}

func NewConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) GetDatabaseDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Asia%%2FShanghai", c.DatabaseConfig.User, c.DatabaseConfig.Password, c.DatabaseConfig.Host, c.DatabaseConfig.Port, c.DatabaseConfig.DBName)
}
