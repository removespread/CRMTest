package configs

import (
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func (c *Config) GetDatabaseDsn() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(c.DatabaseConfig.Dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
