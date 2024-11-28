package srvenv

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
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

func NewConfig(logger *zap.SugaredLogger) (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")

	if err := viper.ReadInConfig(); err != nil {
		logger.Errorf("Error reading config file: %v", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		logger.Errorf("Error unmarshalling config: %v", err)

	}

	return &config, nil
}

func (c *Config) GetDatabaseDsn(logger *zap.SugaredLogger) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(c.DatabaseConfig.Dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error opening database: %v", err)
	}

	return db, nil
}
